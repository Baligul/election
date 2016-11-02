/*
   GET ACCOUNTS
   curl -X POST -H "Content-Type: application/json" -d '{"account_id":[1,2,4], "group_id":[1,2,4], "leader_id":[1,2,4]}' http://localhost:8080/api/accounts
   curl -X POST -H "Content-Type: application/json" -d '{"account_id":[1,2,4], "group_id":[1,2,4], "leader_id":[1,2,4]}' "http://104.197.6.26:8080/api/accounts?mobile_no=9343352734&token=3964d6b3fb85f787"

   Update ACCOUNT
   curl -X PUT -H "Content-Type: application/json" -d '{"account_id":2, "display_name":"balig", "email":"balig@gmail.com", "mobile_no":9657432561, "approved_districts":"Moradabad,Rampur", "approved_acs":"Kanth,Bilaspur", "role":"group lead", "image":"sadsd&%^sd99(&*)", "approved_districts":"Moradabad, Rampur", "group_id":3}' http://localhost:8080/api/account
   curl -X PUT -H "Content-Type: application/json" -d '{"account_id":2, "display_name":"balig", "email":"balig@gmail.com", "mobile_no":9657432561, "approved_districts":"Moradabad,Rampur", "approved_acs":"Kanth,Bilaspur", "role":"group lead", "image":"sadsd&%^sd99(&*)", "approved_districts":"Moradabad, Rampur", "group_id":3}' "http://104.197.6.26:8080/api/account?mobile_no=9343352734&token=3964d6b3fb85f787"

   Create ACCOUNT
   curl -X POST -H "Content-Type: application/json" -d '{"display_name":"balig", "email":"balig@gmail.com", "mobile_no":9657432561, "approved_districts":"Moradabad,Rampur", "approved_acs":"Kanth,Bilaspur", "role":"group lead", "image":"sadsd&%^sd99(&*)", "approved_districts":"Moradabad, Rampur", "group_id":3}' http://localhost:8080/api/account
   curl -X POST -H "Content-Type: application/json" -d '{"display_name":"balig", "email":"balig@gmail.com", "mobile_no":9657432561, "approved_districts":"Moradabad,Rampur", "approved_acs":"Kanth,Bilaspur", "role":"group lead", "image":"sadsd&%^sd99(&*)", "approved_districts":"Moradabad, Rampur", "group_id":3}' "http://104.197.6.26:8080/api/account?mobile_no=9343352734&token=3964d6b3fb85f787"

   Delete ACCOUNT
   curl -X DELETE -H "Content-Type: application/json" -d '{"account_id":1, "group_id":1, "leader_id":2}' http://localhost:8080/api/account
   curl -X DELETE -H "Content-Type: application/json" -d '{"account_id":1, "group_id":1, "leader_id":2}' "http://104.197.6.26:8080/api/account?mobile_no=9343352734&token=3964d6b3fb85f787"
*/

package accounts

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(modelAccounts.Account))
}

type AccountCtrl struct {
	beego.Controller
}

func (e *AccountCtrl) GetAccounts() {
	var (
		accountsCount int64
		accounts      modelAccounts.Accounts
		userAccounts  []*modelAccounts.Account
		err           error
		num           int64
		user          []*modelAccounts.Account
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
	_, err = qsUserAccount.All(&userAccounts)
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
		e.Data["json"] = accounts
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

	e.ServeJSON()
}

func (e *AccountCtrl) CreateAccount() {
	var (
		err       error
		num       int64
		user      []*modelAccounts.Account
		//accountId int64
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

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

	if user[0].Role != "Leader" && user[0].Role != "leader" && user[0].Role != "group lead" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to create account.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userAccount := new(modelAccounts.Account)
	err = json.Unmarshal(inputJson, &userAccount)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	userAccount.Leader_id = user[0].Account_id
	userAccount.Approved_districts = user[0].Approved_districts
	userAccount.Approved_acs = user[0].Approved_acs
	//accountId, _ = o.QueryTable("account").Count()
	//accountId = accountId + 1
	//userAccount.Account_id = int(accountId)
	_, err = o.Insert(userAccount)
	//if err != nil && err.Error() != "no LastInsertId available" {
	if err != nil {		
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	//userAccount.Account_id = int(accountId)
	e.Data["json"] = &userAccount
	e.ServeJSON()
}

func (e *AccountCtrl) UpdateAccount() {
	var (
		err          error
		num          int64
		user         []*modelAccounts.Account
		userAccounts []*modelAccounts.Account
		displayName  string
		email        string
		mobNo        int64
		role         string
		image        string
		groupId      int
		leaderId     int
		age          int
		sex          string
		religion     string
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for account table
	qsUserAccount := o.QueryTable("account")
	condAccountId := orm.NewCondition()

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

	if user[0].Role != "Leader" && user[0].Role != "leader" && user[0].Role != "group lead" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to update any account.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userAccount := new(modelAccounts.Account)
	err = json.Unmarshal(inputJson, &userAccount)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condAccountId = condAccountId.And("Account_id__exact", userAccount.Account_id)
	qsUserAccount = qsUserAccount.SetCond(condAccountId)
	num, err = qsUserAccount.All(&userAccounts)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to find the account.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Account_id != userAccounts[0].Leader_id {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not the leader for this account.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to find the account.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	if strings.TrimSpace(userAccount.Display_name) != "" {
		displayName = strings.TrimSpace(userAccount.Display_name)
	} else {
		displayName = strings.TrimSpace(userAccounts[0].Display_name)
	}
	if strings.TrimSpace(userAccount.Email) != "" {
		email = strings.TrimSpace(userAccount.Email)
	} else {
		email = userAccounts[0].Email
	}
	if userAccount.Mobile_no != 0 {
		mobNo = userAccount.Mobile_no
	} else {
		mobNo = userAccounts[0].Mobile_no
	}
	if strings.TrimSpace(userAccount.Role) != "" {
		role = strings.TrimSpace(userAccount.Role)
	} else {
		role = strings.TrimSpace(userAccounts[0].Role)
	}
	if strings.TrimSpace(userAccount.Image) != "" {
		image = strings.TrimSpace(userAccount.Image)
	} else {
		image = userAccounts[0].Image
	}
	if userAccount.Group_id != 0 {
		groupId = userAccount.Group_id
	} else {
		groupId = userAccounts[0].Group_id
	}
	if userAccount.Leader_id != 0 {
		leaderId = userAccount.Leader_id
	} else {
		leaderId = userAccounts[0].Leader_id
	}
	if userAccount.Age != 0 {
		age = userAccount.Age
	} else {
		age = userAccounts[0].Age
	}
	if strings.TrimSpace(userAccount.Sex) != "" {
		sex = strings.TrimSpace(userAccount.Sex)
	} else {
		sex = userAccounts[0].Sex
	}
	if strings.TrimSpace(userAccount.Religion) != "" {
		religion = strings.TrimSpace(userAccount.Religion)
	} else {
		religion = userAccounts[0].Religion
	}
	num, err = qsUserAccount.Update(orm.Params{
		"Display_name": displayName,
		"Email":        email,
		"Mobile_no":    mobNo,
		"Role":         role,
		"Image":        image,
		"Group_id":     groupId,
		"Leader_id":    leaderId,
		"Age":          age,
		"Sex":          sex,
		"Religion":     religion,
		"Updated_on":   time.Now().Unix(),
	})
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.Data["json"] = fmt.Sprintf("%d rows updated.", num)
	e.ServeJSON()
}

func (e *AccountCtrl) DeleteAccount() {
	var (
		err          error
		num          int64
		user         []*modelAccounts.Account
		userAccounts []*modelAccounts.Account
		totalNum     int64
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for account table
	qsUserAccount := o.QueryTable("account")
	condAccountId := orm.NewCondition()

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

	if user[0].Role != "Leader" && user[0].Role != "leader" && user[0].Role != "group lead" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to delete any account.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userAccount := new(modelAccounts.Account)
	err = json.Unmarshal(inputJson, &userAccount)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if userAccount.Account_id != 0 {
		condAccountId = condAccountId.And("Account_id__exact", userAccount.Account_id)
		qsUserAccount = qsUserAccount.SetCond(condAccountId)
		num, err = qsUserAccount.All(&userAccounts)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to find the account.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num > 0 {
			if user[0].Account_id != userAccounts[0].Leader_id {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("You are not the leader for this account.")
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		} else {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to find the account.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}
	totalNum = 0

	if userAccount.Account_id != 0 {
		num, err = o.QueryTable("account").Filter("Account_id__exact", userAccount.Account_id).Delete()
	}

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete account based on account_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	if userAccount.Group_id != 0 {
		num, err = o.QueryTable("account").Filter("Group_id__exact", userAccount.Group_id).Delete()
	}

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete account based on group_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	if userAccount.Leader_id != 0 {
		num, err = o.QueryTable("account").Filter("Leader_id__exact", userAccount.Leader_id).Delete()
	}

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete account based on leader_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	e.Data["json"] = fmt.Sprintf("%d rows deleted.", totalNum)
	e.ServeJSON()
}
