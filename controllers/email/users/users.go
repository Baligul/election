/*
   Create and Send Pdf
   curl -X POST -H "Content-Type: application/json" -d '{"account_id":[], "group_id":[], "leader_id":[2]}' "http://104.197.6.26:8080/api/email/users?mobile_no=9343352734&token=0d4f93e68cdf26ff"
*/

package users

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"os/exec"
	"strconv"
	"strings"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/lib/html"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/scorredoira/email"
)

func init() {
	//cleanup()
}

type UsersCtrl struct {
	beego.Controller
}

func (e *UsersCtrl) CreateAndEmailPdf() {
	var (
		accountsCount int64
		accounts      modelAccounts.Accounts
		userAccounts  []*modelAccounts.Account
		err           error
		num           int64
		user          []*modelAccounts.Account
		filepath      string
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for account table
	qsUserAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelAccounts.AccountQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	filepath = createFilePath(query)
	if _, err := os.Stat(filepath); os.IsNotExist(err) || filepath == "Downloads/users_list" {
		cond := orm.NewCondition()
		condAccountId := orm.NewCondition()
		condGroupId := orm.NewCondition()
		condLeaderId := orm.NewCondition()

		// Apply filters for each query string

		// Account Id
		for _, accountId := range query.AccountId {
			if accountId > 0 {
				condAccountId = condAccountId.Or("Account_id__exact", accountId)
			}
		}

		// Group Id
		for _, groupId := range query.GroupId {
			if groupId > 0 {
				condGroupId = condGroupId.Or("Group_id__exact", groupId)
			}
		}

		// Leader Id
		for _, leaderId := range query.LeaderId {
			if leaderId > 0 {
				condLeaderId = condLeaderId.Or("Leader_id__exact", leaderId)
			}
		}

		if condAccountId != nil && !condAccountId.IsEmpty() {
			cond = condAccountId
		}

		if condGroupId != nil && !condGroupId.IsEmpty() {
			if cond != nil && !cond.IsEmpty() {
				cond = cond.AndCond(condGroupId)
			} else {
				cond = condGroupId
			}
		}

		if condLeaderId != nil && !condLeaderId.IsEmpty() {
			if cond != nil && !cond.IsEmpty() {
				cond = cond.AndCond(condLeaderId)
			} else {
				cond = condLeaderId
			}
		}

		qsUserAccount = qsUserAccount.SetCond(cond)

		// Get accounts
		accountsCount, _ = qsUserAccount.Count()
		_, err = qsUserAccount.Filter("Leader_id__exact", user[0].Account_id).All(&userAccounts)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to get the accounts.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		accounts.Populate(userAccounts)

		if accountsCount > 0 {
			accounts.Total = accountsCount
		} else {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "ok"
			responseStatus.Message = "No accounts found with this criteria."
			if err != nil {
				responseStatus.Error = err.Error()
			} else {
				responseStatus.Error = "No Error"
			}
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		// PDF creation code start here
		err = html.GenerateHtmlFile("templates/id_cards.html.tmpl", accounts, filepath+".html")
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		err = createPdf(filepath)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	err = sendEmailWithAttachment(user[0].Email, user[0].Display_name, filepath+".pdf")
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = fmt.Sprintf("The file has been sent Successfully.")
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}

func createPdf(filepath string) error {
	// PDF creation code start here
	// converts ".html" file to ".pdf"
	err := exec.Command("wkhtmltox/bin/wkhtmltopdf", filepath+".html", filepath+".pdf").Run()
	if err != nil {
		return err
	}

	return nil
}

func sendEmailWithAttachment(toEmail string, displayName string, filepath string) error {
	// compose the message
	m := email.NewMessage(strings.TrimPrefix(filepath, "Downloads/"), "Dear "+displayName+"!\n\nPlease find attached the required file.\n\nThanks & Regards,\nElectionUBDA Team")
	m.From = mail.Address{Name: "ElectionUBDA Team", Address: "electionubda@gmail.com"}
	m.To = []string{toEmail}

	// add attachments
	if err := m.Attach(filepath); err != nil {
		return err
	}

	// send it
	auth := smtp.PlainAuth("", "electionubda@gmail.com", "hu123*ElectionUBDA", "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}

	return nil
}

func createFilePath(query *modelAccounts.AccountQuery) string {
	var filepath string
	filepath = "Downloads/users_list"

	if len(query.LeaderId) == 1 {
		filepath = "Downloads/Leader_" + strconv.Itoa(query.LeaderId[0])
	}

	if len(query.GroupId) == 1 {
		if filepath == "Downloads/users_list" {
			filepath = "Downloads/Group_" + strconv.Itoa(query.GroupId[0])
		} else {
			filepath = filepath + "-Group_" + strconv.Itoa(query.GroupId[0])
		}
	}

	if len(query.AccountId) == 1 {
		if filepath == "Downloads/users_list" {
			filepath = "Downloads/Account_" + strconv.Itoa(query.AccountId[0])
		} else {
			filepath = filepath + "-Account_" + strconv.Itoa(query.AccountId[0])
		}
	}

	if filepath != "Downloads/users_list" {
		filepath = filepath + "-users_list"
	}

	return filepath
}
