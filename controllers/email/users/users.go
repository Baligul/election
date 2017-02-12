/*
   Create and Send Pdf
   curl -X POST -H "Content-Type: application/json" -d '{"account_id":[], "group_id":[], "leader_id":[2]}' "http://107.178.208.219:80/api/email/users?mobile_no=9343352734&token=b4704cf9a3dc3faa"
*/

package users

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelGroups "github.com/Baligul/election/models/groups"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/formattime"
	"github.com/Baligul/election/lib/html"
	"github.com/Baligul/election/logs"
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
		userAccounts  modelAccounts.ByDisplayName
		err           error
		num           int64
		user          []*modelAccounts.Account
		filepath      string
		userGroup     []*modelGroups.Usergroup
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/email/users, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact at info@humansystech.com for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact at info@humansystech.com for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact at info@humansystech.com for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact at info@humansystech.com for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact at info@humansystech.com for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Send Email Users API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelAccounts.AccountQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	filepath = createFilePath(query)
	if _, err := os.Stat(filepath + ".pdf"); os.IsNotExist(err) || filepath == "Downloads/users_list" {
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
		qsUserAccount = qsUserAccount.OrderBy("Display_name")

		// Get accounts
		accountsCount, _ = qsUserAccount.Count()
		_, err = qsUserAccount.Filter("Leader_id__exact", user[0].Account_id).All(&userAccounts)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to get the accounts.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		for i := range userAccounts {
			userGroup = nil
			userAccount := &userAccounts[i]
			_, err = o.Raw("SELECT title FROM usergroup WHERE group_id=?", userAccount.Group_id).QueryRows(&userGroup)
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Accounts API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to get the accounts.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if len(userGroup) > 0 {
				userAccount.Group_title = userGroup[0].Title
			}
		}
		accounts.Populate(userAccounts)

		if accountsCount > 0 {
			accounts.Total = accountsCount
			sort.Sort(modelAccounts.ByDisplayName(accounts.Accounts))
		} else {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "ok"
			responseStatus.Message = "No accounts found with this criteria."
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
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
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		err = createPdf(filepath)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Users API: "+err.Error())
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
	err := exec.Command("/usr/local/bin/wkhtmltopdf", filepath+".html", filepath+".pdf").Run()
	if err != nil {
		return err
	}

	return nil
}

func sendEmailWithAttachment(toEmail string, displayName string, filepath string) error {
	// compose the message
	m := email.NewMessage(strings.TrimPrefix(filepath, "Downloads/"), "Dear "+displayName+"!\n\nPlease find attached the required file.\n\nThanks & Regards,\nHumansys Technologies Team")
	m.From = mail.Address{Name: "Humansys Technologies Team", Address: "eubdaht@gmail.com"}
	m.To = []string{toEmail}

	// add attachments
	if err := m.Attach(filepath); err != nil {
		return err
	}

	// send it
	auth := smtp.PlainAuth("", "eubdaht@gmail.com", "Huma!2d7D2f3B", "smtp.gmail.com")
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
