/*
   GET TASKS
   curl -X POST -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"complete", "updated_by":[2,3,4], "created_by":[2,3,4]}' http://104.197.6.26:8080/api/tasks
   curl -X POST -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"complete", "updated_by":[2,3,4], "created_by":[2,3,4]}' "http://104.197.6.26:8080/api/tasks?mobile_no=9343352734&token=cc5b86572d1ad660"

   Update TASK
   curl -X PUT -H "Content-Type: application/json" -d '{"task_id": 2, "title":"updated title", "description":"updated description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"in process"}' http://104.197.6.26:8080/api/task
   curl -X PUT -H "Content-Type: application/json" -d '{"task_id": 2, "title":"updated title", "description":"updated description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"in process"}' "http://104.197.6.26:8080/api/task?mobile_no=9343352734&token=cc5b86572d1ad660"

   Create TASK
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4]}' http://104.197.6.26:8080/api/task
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4]}' "http://104.197.6.26:8080/api/task?mobile_no=9343352734&token=cc5b86572d1ad660"

   Delete TASK
   curl -X DELETE -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"complete", "updated_by":[2,3,4], "created_by":[2,3,4]}' http://104.197.6.26:8080/api/task
   curl -X DELETE -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"complete", "updated_by":[2,3,4], "created_by":[2,3,4]}' "http://104.197.6.26:8080/api/task?mobile_no=9343352734&token=cc5b86572d1ad660"
*/

package tasks

import (
	"encoding/json"
	"fmt"
	"strings"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelTasks "github.com/Baligul/election/models/tasks"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterModel(new(modelTasks.Task),
		new(modelTasks.Taskgroupmap),
		new(modelTasks.Taskaccountmap))
}

type TaskCtrl struct {
	beego.Controller
}

func (e *TaskCtrl) GetTasks() {
	var (
		tasksCount   int64
		tasks        modelTasks.Tasks
		tasksReturn  []*modelTasks.TaskReturn
		tgMap        []*modelTasks.Taskgroupmap
		taMap        []*modelTasks.Taskaccountmap
		err          error
		num          int64
		user         []*modelAccounts.Account
		taskIds      []int
		createdByIds []int
		updatedByIds []int
	)

	selectQuery := "SELECT t.task_id, t.title, t.description, t.updated_by, t.created_by, t.updated_on, t.created_on, tam.status AS account_status, tam.updated_by AS account_status_updated_by, tam.updated_on AS account_status_updated_on, a.account_id, a.display_name FROM task AS t LEFT OUTER JOIN taskaccountmap AS tam ON t.task_id = tam.task_id LEFT OUTER JOIN account AS a ON tam.account_id = a.account_id"

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for Taskgroupmap table
	qsTaskgroupmap := o.QueryTable("Taskgroupmap")

	// Create query string for Taskaccountmap table
	qsTaskaccountmap := o.QueryTable("Taskaccountmap")

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
	query := new(modelTasks.TaskQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId := ""
	condGroupsAssigned := orm.NewCondition()
	condAccountsAssigned := orm.NewCondition()
	condStatus := orm.NewCondition()
	condUpdatedBy := ""
	condCreatedBy := ""

	// Apply filters for each query string
	// Task Id
	for _, taskId := range query.TaskId {
		if taskId > 0 {
			taskIds = append(taskIds, taskId)
			condTaskId = condTaskId + " ?,"
		}
	}

	// Groups Assigned
	for _, groupAssigned := range query.GroupsAssigned {
		if groupAssigned > 0 {
			condGroupsAssigned = condGroupsAssigned.Or("Group_id__exact", groupAssigned)
		}
	}

	// Accounts Assigned
	for _, accountAssigned := range query.AccountsAssigned {
		if accountAssigned > 0 {
			condAccountsAssigned = condAccountsAssigned.Or("Account_id__exact", accountAssigned)
		}
	}

	// Status
	if query.Status == "new" || query.Status == "in process" || query.Status == "complete" {
		condStatus = condStatus.And("Status__exact", query.Status)
	}

	// Updated By
	for _, updatedBy := range query.UpdatedBy {
		if updatedBy > 0 {
			updatedByIds = append(updatedByIds, updatedBy)
			condUpdatedBy = condUpdatedBy + " ?,"
		}
	}

	// Created By
	for _, createdBy := range query.CreatedBy {
		if createdBy > 0 {
			createdByIds = append(createdByIds, createdBy)
			condCreatedBy = condCreatedBy + " ?,"
		}
	}

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condGroupsAssigned)
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condStatus)
	}

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condAccountsAssigned)
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condStatus)
	}

	_, err = qsTaskgroupmap.Filter("Created_by__exact", user[0].Account_id).All(&tgMap)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Taskgroupmap. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	for _, tg := range tgMap {
		taskIds = append(taskIds, tg.Task_id)
		condTaskId = condTaskId + "?,"
	}

	_, err = qsTaskaccountmap.Filter("Created_by__exact", user[0].Account_id).All(&taMap)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Taskaccountmap. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	for _, ta := range taMap {
		taskIds = append(taskIds, ta.Task_id)
		condTaskId = condTaskId + "?,"
	}

	// Get tasks
	condTaskId = strings.TrimRight(strings.TrimSpace(condTaskId), ",")
	if len(taskIds) > 0 && len(createdByIds) > 0 && len(updatedByIds) > 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.created_by IN (" + condCreatedBy + ") AND t.updated_by IN (" + condUpdatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, taskIds, createdByIds, updatedByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) > 0 && len(createdByIds) == 0 && len(updatedByIds) == 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, taskIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) > 0 && len(createdByIds) > 0 && len(updatedByIds) == 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.created_by IN (" + condCreatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, taskIds, createdByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) > 0 && len(createdByIds) == 0 && len(updatedByIds) > 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.updated_by IN (" + condUpdatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, taskIds, updatedByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) == 0 && len(createdByIds) > 0 && len(updatedByIds) > 0 {
		selectQuery = selectQuery + " WHERE t.created_by IN (" + condCreatedBy + ") AND t.updated_by IN (" + condUpdatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, createdByIds, updatedByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) == 0 && len(createdByIds) > 0 && len(updatedByIds) == 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.created_by IN (" + condCreatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, createdByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else if len(taskIds) == 0 && len(createdByIds) == 0 && len(updatedByIds) > 0 {
		selectQuery = selectQuery + " WHERE t.task_id IN (" + condTaskId + ") AND t.updated_by IN (" + condUpdatedBy + ") AND t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, updatedByIds, user[0].Account_id).QueryRows(&tasksReturn)
	} else {
		selectQuery = selectQuery + " WHERE t.created_by = ?"
		tasksCount, err = o.Raw(selectQuery, user[0].Account_id).QueryRows(&tasksReturn)
	}

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	tasks.Populate(tasksReturn)
	if tasksCount > 0 {
		tasks.Total = tasksCount
		e.Data["json"] = tasks
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No tasks found with this criteria."
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

func (e *TaskCtrl) CreateTask() {
	var (
		err      error
		num      int64
		user     []*modelAccounts.Account
		accounts []*modelAccounts.Account
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

	if user[0].Role != "Leader" && user[0].Role != "leader" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised to create task.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	userTask := new(modelTasks.TaskCreateDelete)
	task := new(modelTasks.Task)
	err = json.Unmarshal(inputJson, &userTask)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	task.Transpose(userTask)
	task.Updated_by = user[0].Account_id
	task.Created_by = user[0].Account_id
	taskId, err := o.Insert(task)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	tgMap := new(modelTasks.Taskgroupmap)

	for _, groupId := range userTask.Groups_assigned {
		tgMap.Task_id = int(taskId)
		tgMap.Group_id = groupId
		tgMap.Status = "new"
		tgMap.Updated_by = user[0].Account_id
		tgMap.Created_by = user[0].Account_id
		_, err := o.Insert(tgMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		accounts = nil
		_, err = qsAccount.All(&accounts)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		taMap := new(modelTasks.Taskaccountmap)

		for _, account := range accounts {
			taMap.Task_id = int(taskId)
			taMap.Account_id = account.Account_id
			taMap.Status = "new"
			taMap.Updated_by = user[0].Account_id
			taMap.Created_by = user[0].Account_id
			_, err := o.Insert(taMap)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	taMap := new(modelTasks.Taskaccountmap)

	for _, accountId := range userTask.Accounts_assigned {
		taMap.Task_id = int(taskId)
		taMap.Account_id = accountId
		taMap.Status = "new"
		taMap.Updated_by = user[0].Account_id
		taMap.Created_by = user[0].Account_id
		_, err := o.Insert(taMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	e.Data["json"] = &userTask
	e.ServeJSON()
}

func (e *TaskCtrl) UpdateTask() {
	var (
		err         error
		num         int64
		user        []*modelAccounts.Account
		userTasks   []*modelTasks.Task
		tgMap       []*modelTasks.Taskgroupmap
		taMap       []*modelTasks.Taskaccountmap
		title       string
		description string
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for task table
	qsTask := o.QueryTable("task")
	condTaskId := orm.NewCondition()
	qsTaskgroupmap := o.QueryTable("Taskgroupmap")
	condTaskgroupmap := orm.NewCondition()
	qsTaskaccountmap := o.QueryTable("Taskaccountmap")
	condTaskaccountmap := orm.NewCondition()

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
	userTask := new(modelTasks.TaskCreateDelete)
	err = json.Unmarshal(inputJson, &userTask)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId = condTaskId.And("Task_id__exact", userTask.Task_id).And("Created_by__exact", user[0].Account_id)
	qsTask = qsTask.SetCond(condTaskId)
	num, err = qsTask.All(&userTasks)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to find the task.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num == 0 {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to find the task.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if strings.TrimSpace(userTask.Title) != "" {
		title = strings.TrimSpace(userTask.Title)
	} else {
		title = strings.TrimSpace(userTasks[0].Title)
	}
	if strings.TrimSpace(userTask.Description) != "" {
		description = strings.TrimSpace(userTask.Description)
	} else {
		description = strings.TrimSpace(userTasks[0].Description)
	}
	num, err = qsTask.Update(orm.Params{
		"Title":       title,
		"Description": description,
		"Updated_by":  user[0].Account_id,
	})
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	for _, groupId := range userTask.Groups_assigned {
		condTaskgroupmap = condTaskgroupmap.And("Task_id__exact", userTask.Task_id).And("Group_id__exact", groupId)
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condTaskgroupmap)
		num, err = qsTaskgroupmap.All(&tgMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskgroupmap. Unable to find the Taskgroupmap.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num == 0 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskgroupmap. Unable to find the Taskgroupmap.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		num, err = qsTaskgroupmap.Update(orm.Params{
			"Status": userTask.Status,
		})
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	for _, accountId := range userTask.Accounts_assigned {
		condTaskaccountmap = condTaskaccountmap.And("Task_id__exact", userTask.Task_id).And("Account_id__exact", accountId)
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condTaskaccountmap)
		num, err = qsTaskaccountmap.All(&taMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskaccountmap. Unable to find the Taskaccountmap.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num == 0 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskaccountmap. Unable to find the Taskaccountmap.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		num, err = qsTaskaccountmap.Update(orm.Params{
			"Status": userTask.Status,
		})
		if err != nil {
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

func (e *TaskCtrl) DeleteTask() {
	var (
		tgMap []*modelTasks.Taskgroupmap
		taMap []*modelTasks.Taskaccountmap
		err   error
		num   int64
		user  []*modelAccounts.Account
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for task table
	qsTask := o.QueryTable("task")

	// Create query string for Taskgroupmap table
	qsTaskgroupmap := o.QueryTable("Taskgroupmap")

	// Create query string for Taskaccountmap table
	qsTaskaccountmap := o.QueryTable("Taskaccountmap")

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
	query := new(modelTasks.TaskQuery)

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
	condTaskId := orm.NewCondition()
	condGroupsAssigned := orm.NewCondition()
	condAccountsAssigned := orm.NewCondition()
	condStatus := orm.NewCondition()
	condUpdatedBy := orm.NewCondition()
	condCreatedBy := orm.NewCondition()

	// Apply filters for each query string
	// Task Id
	for _, taskId := range query.TaskId {
		if taskId > 0 {
			condTaskId = condTaskId.Or("Task_id__exact", taskId)
		}
	}

	// Groups Assigned
	for _, groupAssigned := range query.GroupsAssigned {
		if groupAssigned > 0 {
			condGroupsAssigned = condGroupsAssigned.Or("Group_id__exact", groupAssigned)
		}
	}

	// Accounts Assigned
	for _, accountAssigned := range query.AccountsAssigned {
		if accountAssigned > 0 {
			condAccountsAssigned = condAccountsAssigned.Or("Account_id__exact", accountAssigned)
		}
	}

	// Status
	if query.Status == "new" || query.Status == "in process" || query.Status == "complete" {
		condStatus = condStatus.And("Status__exact", query.Status)
	}

	// Updated By
	for _, updatedBy := range query.UpdatedBy {
		if updatedBy > 0 {
			condUpdatedBy = condUpdatedBy.Or("Updated_by__exact", updatedBy)
		}
	}

	// Created By
	for _, createdBy := range query.CreatedBy {
		if createdBy > 0 {
			condCreatedBy = condCreatedBy.Or("Created_by__exact", createdBy)
		}
	}

	if condTaskId != nil && !condTaskId.IsEmpty() {
		cond = condTaskId
	}

	if condUpdatedBy != nil && !condUpdatedBy.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condUpdatedBy)
		} else {
			cond = condUpdatedBy
		}
	}

	if condCreatedBy != nil && !condCreatedBy.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condCreatedBy)
		} else {
			cond = condCreatedBy
		}
	}

	qsTask = qsTask.SetCond(cond)

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condGroupsAssigned)
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condStatus)
	}

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condAccountsAssigned)
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condStatus)
	}

	_, err = qsTaskgroupmap.Filter("Created_by__exact", user[0].Account_id).All(&tgMap)
	// Delete from Taskgroupmap
	num, err = qsTaskgroupmap.Delete()
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete task based on task_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId = orm.NewCondition()

	for _, tg := range tgMap {
		condTaskId = condTaskId.Or("Task_id__exact", tg.Task_id)
	}

	if condTaskId != nil && !condTaskId.IsEmpty() {
		qsTask = qsTask.SetCond(condTaskId)
	}

	_, err = qsTaskaccountmap.Filter("Created_by__exact", user[0].Account_id).All(&taMap)
	// Delete from Taskaccountmap
	num, err = qsTaskaccountmap.Delete()
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete task based on task_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId = orm.NewCondition()

	for _, ta := range taMap {
		condTaskId = condTaskId.Or("Task_id__exact", ta.Task_id)
	}

	if condTaskId != nil && !condTaskId.IsEmpty() {
		qsTask = qsTask.SetCond(condTaskId)
	}

	// Delete tasks
	num, err = qsTask.Filter("Created_by__exact", user[0].Account_id).Delete()
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request to delete task based on task_id. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.Data["json"] = fmt.Sprintf("%d rows deleted.", num)
	e.ServeJSON()
}
