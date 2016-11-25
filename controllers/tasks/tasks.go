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
		new(modelTasks.TaskGroupMap),
		new(modelTasks.TaskAccountMap))
}

type TaskCtrl struct {
	beego.Controller
}

func (e *TaskCtrl) GetTasks() {
	var (
		tasksCount int64
		tasks      modelTasks.Tasks
		userTasks  []*modelTasks.Task
		tgMap      []*modelTasks.TaskGroupMap
		taMap      []*modelTasks.TaskAccountMap
		err        error
		num        int64
		user       []*modelAccounts.Account
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	// Create query string for task table
	qsTask := o.QueryTable("task")

	// Create query string for taskgroupmap table
	qsTaskGroupMap := o.QueryTable("taskgroupmap")

	// Create query string for taskaccountmap table
	qsTaskAccountMap := o.QueryTable("taskaccountmap")

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
		qsTaskGroupMap = qsTaskGroupMap.SetCond(condGroupsAssigned)
		qsTaskGroupMap = qsTaskGroupMap.SetCond(condStatus)
	}

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskAccountMap = qsTaskAccountMap.SetCond(condAccountsAssigned)
		qsTaskAccountMap = qsTaskAccountMap.SetCond(condStatus)
	}

	_, err = qsTaskGroupMap.Filter("Created_by__exact", user[0].Account_id).All(&tgMap)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error TaskGroupMap. Unable to get the tasks.")
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

	_, err = qsTaskAccountMap.Filter("Created_by__exact", user[0].Account_id).All(&taMap)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error TaskAccountMap. Unable to get the tasks.")
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

	// Get tasks
	tasksCount, _ = qsTask.Count()
	_, err = qsTask.Filter("Created_by__exact", user[0].Account_id).All(&userTasks)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	tasks.Populate(userTasks)

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
		err  error
		num  int64
		user []*modelAccounts.Account
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

	tgMap := new(modelTasks.TaskGroupMap)

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
	}

	taMap := new(modelTasks.TaskAccountMap)

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
		tgMap       []*modelTasks.TaskGroupMap
		taMap       []*modelTasks.TaskAccountMap
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
	qsTaskGroupMap := o.QueryTable("taskgroupmap")
	condTaskGroupMap := orm.NewCondition()
	qsTaskAccountMap := o.QueryTable("taskaccountmap")
	condTaskAccountMap := orm.NewCondition()

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
		condTaskGroupMap = condTaskGroupMap.And("Task_id__exact", userTask.Task_id).And("Group_id__exact", groupId)
		qsTaskGroupMap = qsTaskGroupMap.SetCond(condTaskGroupMap)
		num, err = qsTaskGroupMap.All(&tgMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error TaskGroupMap. Unable to find the taskgroupmap.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num == 0 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error TaskGroupMap. Unable to find the taskgroupmap.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		num, err = qsTaskGroupMap.Update(orm.Params{
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
		condTaskAccountMap = condTaskAccountMap.And("Task_id__exact", userTask.Task_id).And("Account_id__exact", accountId)
		qsTaskAccountMap = qsTaskAccountMap.SetCond(condTaskAccountMap)
		num, err = qsTaskAccountMap.All(&taMap)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error TaskAccountMap. Unable to find the taskaccountmap.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		if num == 0 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error TaskAccountMap. Unable to find the taskaccountmap.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		num, err = qsTaskAccountMap.Update(orm.Params{
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
		tgMap []*modelTasks.TaskGroupMap
		taMap []*modelTasks.TaskAccountMap
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

	// Create query string for taskgroupmap table
	qsTaskGroupMap := o.QueryTable("taskgroupmap")

	// Create query string for taskaccountmap table
	qsTaskAccountMap := o.QueryTable("taskaccountmap")

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
		qsTaskGroupMap = qsTaskGroupMap.SetCond(condGroupsAssigned)
		qsTaskGroupMap = qsTaskGroupMap.SetCond(condStatus)
	}

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskAccountMap = qsTaskAccountMap.SetCond(condAccountsAssigned)
		qsTaskAccountMap = qsTaskAccountMap.SetCond(condStatus)
	}

	_, err = qsTaskGroupMap.Filter("Created_by__exact", user[0].Account_id).All(&tgMap)
	// Delete from taskgroupmap
	num, err = qsTaskGroupMap.Delete()
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

	_, err = qsTaskAccountMap.Filter("Created_by__exact", user[0].Account_id).All(&taMap)
	// Delete from taskaccountmap
	num, err = qsTaskAccountMap.Delete()
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
