/*
   GET GROUPS
   curl -X POST -H "Content-Type: application/json" -d '{"group_id":[1,2,4], "created_by":[2,3,4]}, "group_lead_id":[1,2,4]}' http://localhost:8080/api/groups
   curl -X POST -H "Content-Type: application/json" -d '{}' "http://107.178.208.219:80/api/groups?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Update GROUP
   curl -X PUT -H "Content-Type: application/json" -d '{"group_id":1, "title":"new title", "description":"new description", "group_lead_id":5, "account_id":[2]}' http://localhost:8080/api/group
   curl -X PUT -H "Content-Type: application/json" -d '{"group_id":1, "title":"new title", "description":"new description", "group_lead_id":5, "account_id":[2,3]}' "http://107.178.208.219:80/api/group?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Create GROUP
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "group_lead_id":5, "account_id":[2,3]}' http://localhost:8080/api/group
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "group_lead_id":5, "account_id":[2,3]}' "http://107.178.208.219:80/api/group?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Delete GROUP
   curl -X DELETE -H "Content-Type: application/json" -d '{"group_id":1, "created_by":2, "group_lead_id":5}' http://localhost:8080/api/group
   curl -X DELETE -H "Content-Type: application/json" -d '{"group_id":1, "created_by":2, "group_lead_id":5}' "http://107.178.208.219:80/api/group?mobile_no=9343352734&token=f8a220f5e8d1741d"
*/

package groups

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelGroups "github.com/Baligul/election/models/groups"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/logs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(modelGroups.Usergroup))
}

type GroupCtrl struct {
	beego.Controller
}

func (e *GroupCtrl) GetGroups() {
	var (
		groupsCount int64
		groups      modelGroups.Groups
		userGroups  modelGroups.ByTitle
		err         error
		num         int64
		user        []*modelAccounts.Account
		users       []*modelAccounts.Account
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for usergroup table
	qsGroup := o.QueryTable("usergroup")

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
		// Log the error
		_ = logs.WriteLogs("Get Groups API: " + err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": time.Now(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Last Login in Get Groups API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelGroups.GroupQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Groups API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	cond := orm.NewCondition()
	condGroupId := orm.NewCondition()
	condCreatedBy := orm.NewCondition()
	condGroupLeadId := orm.NewCondition()

	// Apply filters for each query string
	// Group Id
	for _, groupId := range query.GroupId {
		if groupId > 0 {
			condGroupId = condGroupId.Or("Group_id__exact", groupId)
		}
	}

	// Created By
	for _, createdBy := range query.CreatedBy {
		if createdBy > 0 {
			condCreatedBy = condCreatedBy.Or("Created_by__exact", createdBy)
		}
	}

	// Group Lead Id
	for _, groupLeadId := range query.GroupLeadId {
		if groupLeadId > 0 {
			condGroupLeadId = condGroupLeadId.Or("Group_lead_id__exact", groupLeadId)
		}
	}

	if condGroupId != nil && !condGroupId.IsEmpty() {
		cond = condGroupId
	}

	if condCreatedBy != nil && !condCreatedBy.IsEmpty() { 
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condCreatedBy)
		} else {
			cond = condCreatedBy
		}
	}

	if condGroupLeadId != nil && !condGroupLeadId.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condGroupLeadId)
		} else {
			cond = condGroupLeadId
		}
	}

	qsGroup = qsGroup.SetCond(cond)
	qsGroup = qsGroup.OrderBy("Title")

	// Get groups
	groupsCount, _ = qsGroup.Filter("Created_by__exact", user[0].Account_id).Count()
	_, err = qsGroup.Filter("Created_by__exact", user[0].Account_id).All(&userGroups)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Groups API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to get the groups.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	for i := range userGroups {
			_, err = o.Raw("SELECT * FROM account WHERE group_id=?", userGroups[i].Group_id).QueryRows(&users)
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("Get Task Details API: " + err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the task details.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			userGroups[i].Accounts = users
	}
	groups.Populate(userGroups)

	if groupsCount > 0 {
		groups.Total = groupsCount
		sort.Sort(modelGroups.ByTitle(groups.Groups))
		e.Data["json"] = groups
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No groups found with this criteria."
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Groups API: " + err.Error())
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.ServeJSON()
}

func (e *GroupCtrl) CreateGroup() {
	var (
		err  error
		num  int64
		user []*modelAccounts.Account
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		// Log the error
		_ = logs.WriteLogs("Create Group API: " + err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": time.Now(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Last Login in Create Group API: " + err.Error())
	}

	if user[0].Role != "Leader" && user[0].Role != "leader" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to create group.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userGroup := new(modelGroups.Usergroup)
	err = json.Unmarshal(inputJson, &userGroup)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Create Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	userGroup.Updated_by = user[0].Account_id
	userGroup.Created_by = user[0].Account_id
	newGroupId, err := o.Insert(userGroup)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Create Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Update all accounts for which this group is being created
	if len(userGroup.Account_id) > 0 {
		// Create query string for account table
		qsUserAccount := o.QueryTable("account")
		condAccountId := orm.NewCondition()

		// Account Id
		for _, accountId := range userGroup.Account_id {
			if accountId > 0 {
				condAccountId = condAccountId.Or("Account_id__exact", accountId)
			}
		}

		qsUserAccount = qsUserAccount.SetCond(condAccountId)

		_, err = qsUserAccount.Update(orm.Params{
			"Group_id": newGroupId,
		})
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Create Group API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	e.Data["json"] = &userGroup
	e.ServeJSON()
}

func (e *GroupCtrl) UpdateGroup() {
	var (
		err         error
		num         int64
		user        []*modelAccounts.Account
		userGroups  []*modelGroups.Usergroup
		title       string
		description string
		groupLeadId int
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for usergroup table
	qsGroup := o.QueryTable("usergroup")
	condGroupId := orm.NewCondition()

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
		// Log the error
		_ = logs.WriteLogs("Update Group API: " + err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": time.Now(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Last Login in Update Group API: " + err.Error())
	}

	if user[0].Role != "Leader" && user[0].Role != "leader" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to update any group.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userGroup := new(modelGroups.Usergroup)
	err = json.Unmarshal(inputJson, &userGroup)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condGroupId = condGroupId.And("Group_id__exact", userGroup.Group_id)
	qsGroup = qsGroup.SetCond(condGroupId)
	num, err = qsGroup.All(&userGroups)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to find the group.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Account_id != userGroups[0].Created_by {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not the owner of this group.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to find the group.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	if strings.TrimSpace(userGroup.Title) != "" {
		title = strings.TrimSpace(userGroup.Title)
	} else {
		title = strings.TrimSpace(userGroups[0].Title)
	}
	if strings.TrimSpace(userGroup.Description) != "" {
		description = strings.TrimSpace(userGroup.Description)
	} else {
		description = strings.TrimSpace(userGroups[0].Description)
	}
	if userGroup.Group_lead_id != 0 {
		groupLeadId = userGroup.Group_lead_id
	} else {
		groupLeadId = userGroups[0].Group_lead_id
	}
	num, err = qsGroup.Update(orm.Params{
		"Title":         title,
		"Description":   description,
		"Group_lead_id": groupLeadId,
		"Updated_by":    user[0].Account_id,
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Update all accounts for which this group is being updated
	if len(userGroup.Account_id) > 0 {
		// Create query string for account table
		qsUserAccount := o.QueryTable("account")
		condAccountId := orm.NewCondition()

		// Account Id
		for _, accountId := range userGroup.Account_id {
			if accountId > 0 {
				condAccountId = condAccountId.Or("Account_id__exact", accountId)
			}
		}

		qsUserAccount = qsUserAccount.SetCond(condAccountId)

		_, err = qsUserAccount.Update(orm.Params{
			"Group_id": userGroup.Group_id,
		})
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Update Group API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	e.Data["json"] = fmt.Sprintf("%d rows updated.", num)
	e.ServeJSON()
}

func (e *GroupCtrl) DeleteGroup() {
	var (
		err        error
		num        int64
		user       []*modelAccounts.Account
		userGroups []*modelGroups.Usergroup
		totalNum   int64
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for usergroup table
	qsGroup := o.QueryTable("usergroup")
	condGroupId := orm.NewCondition()

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
		// Log the error
		_ = logs.WriteLogs("Delete Group API: " + err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": time.Now(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Last Login in Delete Group API: " + err.Error())
	}

	if user[0].Role != "Leader" && user[0].Role != "leader" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to delete any group.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userGroup := new(modelGroups.Usergroup)
	err = json.Unmarshal(inputJson, &userGroup)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if userGroup.Group_id != 0 {
		condGroupId = condGroupId.And("Group_id__exact", userGroup.Group_id)
		qsGroup = qsGroup.SetCond(condGroupId)
		num, err = qsGroup.All(&userGroups)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Delete Group API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to find the group.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num > 0 {
			if user[0].Account_id != userGroups[0].Created_by {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("You are not the owner of this group.")
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		} else {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to find the group.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}
	totalNum = 0
	if userGroup.Group_id != 0 {
		num, err = o.QueryTable("usergroup").Filter("Group_id__exact", userGroup.Group_id).Delete()
	}

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete group based on group_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	if userGroup.Created_by != 0 {
		num, err = o.QueryTable("usergroup").Filter("Created_by__exact", userGroup.Created_by).Delete()
	}

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete group based on created_by. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	if userGroup.Group_lead_id != 0 {
		num, err = o.QueryTable("usergroup").Filter("Group_lead_id__exact", userGroup.Group_lead_id).Delete()
	}

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Group API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete group based on group_lead_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	} else {
		totalNum = totalNum + num
	}

	e.Data["json"] = fmt.Sprintf("%d rows deleted.", totalNum)
	e.ServeJSON()
}
