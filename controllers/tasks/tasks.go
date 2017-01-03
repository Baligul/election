/*
   GET Created TASKS
   curl -X POST -H "Content-Type: application/json" -d '{"groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"completed"}' http://107.178.208.219:80/api/tasks/created
   curl -X POST -H "Content-Type: application/json" -d '{"groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"completed"}' "http://107.178.208.219:80/api/tasks/created?mobile_no=9343352734&token=f8a220f5e8d1741d"

   GET My TASKS
   curl -X POST -H "Content-Type: application/json" -d '{"status":"completed"}' http://107.178.208.219:80/api/tasks/my
   curl -X POST -H "Content-Type: application/json" -d '{"status":"completed"}' "http://107.178.208.219:80/api/tasks/my?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Update TASK
   curl -X PUT -H "Content-Type: application/json" -d '{"task_id": 2, "title":"updated title", "description":"updated description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"in process"}' http://107.178.208.219:80/api/task
   curl -X PUT -H "Content-Type: application/json" -d '{"task_id": 2, "title":"updated title", "description":"updated description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"in process"}' "http://107.178.208.219:80/api/task?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Create TASK
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4]}' http://107.178.208.219:80/api/task
   curl -X POST -H "Content-Type: application/json" -d '{"title":"new title", "description":"new description", "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4]}' "http://107.178.208.219:80/api/task?mobile_no=9343352734&token=f8a220f5e8d1741d"

   Delete TASK
   curl -X DELETE -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"completed", "updated_by":[2,3,4], "created_by":[2,3,4]}' http://107.178.208.219:80/api/task
   curl -X DELETE -H "Content-Type: application/json" -d '{"task_id":[1,2,4], "groups_assigned":[2,3,4], "accounts_assigned":[2,3,4], "status":"completed", "updated_by":[2,3,4], "created_by":[2,3,4]}' "http://107.178.208.219:80/api/task?mobile_no=9343352734&token=f8a220f5e8d1741d"

   GET TASK DETAILS
   curl -X POST -H "Content-Type: application/json" -d '{"task_id":1}' http://107.178.208.219:80/api/taskdetail
   curl -X POST -H "Content-Type: application/json" -d '{"task_id":1}' "http://107.178.208.219:80/api/taskdetail?mobile_no=9343352734&token=f8a220f5e8d1741d"
*/

package tasks

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelGroups "github.com/Baligul/election/models/groups"
	modelTasks "github.com/Baligul/election/models/tasks"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/logs"
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

func (e *TaskCtrl) GetCreatedTasks() {
	var (
		tasksCount int64
		tasks      modelTasks.Tasks
		userTasks  modelTasks.ByTitle
		userGroups modelGroups.ByTitle
		groups     modelGroups.Groups
		users      modelAccounts.ByDisplayName
		accounts   modelAccounts.Accounts
		tgMap      []*modelTasks.Taskgroupmap
		taMap      []*modelTasks.Taskaccountmap
		err        error
		num        int64
		user       []*modelAccounts.Account
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
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Get Tasks API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelTasks.TaskQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId := orm.NewCondition()
	condGroupsAssigned := orm.NewCondition()
	condAccountsAssigned := orm.NewCondition()

	// Apply filters for each query string
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

	if condGroupsAssigned != nil && !condGroupsAssigned.IsEmpty() {
		qsTaskgroupmap = qsTaskgroupmap.SetCond(condGroupsAssigned)
		// Status
		if query.Status == "new" || query.Status == "in process" || query.Status == "completed" {
			_, err = qsTaskgroupmap.Filter("Status__exact", query.Status).All(&tgMap)
		} else {
			_, err = qsTaskgroupmap.All(&tgMap)
		}
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskgroupmap. Unable to get the tasks.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		for _, tg := range tgMap {
			condTaskId = condTaskId.Or("Task_id__exact", tg.Task_id)
		}
	}

	if condAccountsAssigned != nil && !condAccountsAssigned.IsEmpty() {
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condAccountsAssigned)
		// Status
		if query.Status == "new" || query.Status == "in process" || query.Status == "completed" {
			_, err = qsTaskaccountmap.Filter("Status__exact", query.Status).All(&taMap)
		} else {
			_, err = qsTaskaccountmap.All(&taMap)
		}

		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Taskaccountmap. Unable to get the tasks.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		for _, ta := range taMap {
			condTaskId = condTaskId.Or("Task_id__exact", ta.Task_id)
		}
	}

	if (len(query.AccountsAssigned) > 0 || len(query.GroupsAssigned) > 0) && (condTaskId == nil || condTaskId.IsEmpty()) {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "Groups or Accounts passed has no tasks assigned to them."
		responseStatus.Error = "No Error"
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if condTaskId != nil && !condTaskId.IsEmpty() {
		qsTask = qsTask.SetCond(condTaskId)
	}

	// Get tasks
	tasksCount, _ = qsTask.Filter("Created_by__exact", user[0].Account_id).Count()
	_, err = qsTask.Filter("Created_by__exact", user[0].Account_id).All(&userTasks)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if tasksCount > 0 {
		// Get accounts
		accountsCount, _ := qsAccount.Filter("Leader_id__exact", user[0].Account_id).Count()
		_, err = qsAccount.Filter("Leader_id__exact", user[0].Account_id).All(&users)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to get the accounts.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		accounts.Populate(users)
		accounts.Total = accountsCount
		sort.Sort(modelAccounts.ByDisplayName(accounts.Accounts))

		// Get groups
		groupsCount, _ := qsGroup.Filter("Created_by__exact", user[0].Account_id).Count()
		_, err = qsGroup.Filter("Created_by__exact", user[0].Account_id).All(&userGroups)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Groups. Unable to get the groups.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		groups.Populate(userGroups)
		groups.Total = groupsCount
		sort.Sort(modelGroups.ByTitle(groups.Groups))

		for i := range userTasks {
			users = nil
			// Get array of accounts
			_, err = o.Raw("SELECT DISTINCT tam.account_id, a.display_name FROM taskaccountmap AS tam LEFT OUTER JOIN account AS a ON tam.account_id = a.account_id WHERE tam.task_id=?", userTasks[i].Task_id).QueryRows(&users)

			if err != nil {
				// Log the error
				_ = logs.WriteLogs("Get Tasks API: " + err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the assigned accounts to task.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			userTasks[i].Accounts = users

			userGroups = nil
			// Get array of groups
			_, err = o.Raw("SELECT DISTINCT tgm.group_id, g.title AS group_title FROM taskgroupmap AS tgm LEFT OUTER JOIN usergroup AS g ON tgm.group_id = g.group_id WHERE tgm.task_id=?", userTasks[i].Task_id).QueryRows(&userGroups)

			if err != nil {
				// Log the error
				_ = logs.WriteLogs("Get Tasks API: " + err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the assigned groups to task.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			userTasks[i].Groups = userGroups
		}

		tasks.Populate(userTasks)
		tasks.Total = tasksCount
		sort.Sort(modelTasks.ByTitle(tasks.Tasks))
		tasks.Accounts = accounts
		tasks.Groups = groups
		e.Data["json"] = tasks
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No tasks found with this criteria."
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.ServeJSON()
}

func (e *TaskCtrl) GetMyTasks() {
	var (
		tasksCount int64
		tasks      modelTasks.Tasks
		userTasks  modelTasks.ByTitle
		taMap      []*modelTasks.Taskaccountmap
		err        error
		num        int64
		user       []*modelAccounts.Account
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

	// Create query string for task table
	qsTask := o.QueryTable("task")

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
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Get Tasks API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelTasks.TaskQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	fmt.Println("Query is ", query)

	fmt.Println("Status is ", query.Status)

	// Status
	if query.Status == "new" || query.Status == "in process" || query.Status == "completed" {
		fmt.Println("Query is set here")
		_, err = qsTaskaccountmap.Filter("Account_id__exact", user[0].Account_id).Filter("Status__exact", query.Status).All(&taMap)
	} else {
		_, err = qsTaskaccountmap.Filter("Account_id__exact", user[0].Account_id).All(&taMap)
	}

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Tasks API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Taskgroupmap. Unable to get the tasks.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	fmt.Println("TasksAccountMap ids found to be ", len(taMap))
	



	condTaskId := orm.NewCondition()
	condTaskId = nil

	fmt.Println("condTaskId is ", condTaskId)

	for _, am := range taMap {
		condTaskId = condTaskId.Or("Task_id__exact", am.Task_id)
	}

	fmt.Println("condTaskId becomes ", condTaskId)

	if condTaskId != nil && !condTaskId.IsEmpty() {
		qsTask = qsTask.SetCond(condTaskId)

		// Get tasks
		tasksCount, _ = qsTask.Count()

		fmt.Println("Here we got the tasks count as ", tasksCount)
		_, err = qsTask.All(&userTasks)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to get the tasks.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		tasks.Populate(userTasks)
	}
fmt.Println("tasksCount ", tasksCount)
	if tasksCount > 0 {
		tasks.Total = tasksCount
		sort.Sort(modelTasks.ByTitle(tasks.Tasks))
		e.Data["json"] = tasks
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No tasks found with this criteria."
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Get Tasks API: " + err.Error())
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.ServeJSON()
}

func (e *TaskCtrl) GetTaskDetail() {
	var (
		tasks          []*modelTasks.Task
		taskDetail     modelTasks.TaskDetail
		accountDetails []*modelTasks.AccountDetails
		groups         []*modelGroups.Usergroup
		accountDN      []*modelTasks.AccountDisplayName
		err            error
		num            int64
		user           []*modelAccounts.Account
		userGroup      []*modelGroups.Usergroup
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

	// Create query string for task table
	qsTask := o.QueryTable("task")

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
		_ = logs.WriteLogs("Get Task Details API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Get Task Details API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelTasks.Task)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Get Task Details API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	condTaskId := orm.NewCondition()

	// Apply filters for each query string
	// Task Id
	if query.Task_id > 0 {
		condTaskId = condTaskId.And("Task_id__exact", query.Task_id)
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid task id. Unable to get the task details.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	qsTask = qsTask.SetCond(condTaskId)
	qsTask = qsTask.OrderBy("Title")

	// Get task details
	num, err = qsTask.All(&tasks)
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

	if num > 0 {
		taskDetail.Populate(tasks[0])

		// Get array of account details
		_, err = o.Raw("SELECT DISTINCT tam.account_id, tam.status, tam.updated_by AS status_updated_by, tam.updated_on AS status_updated_on, tam.created_by AS task_assigned_by, tam.created_on AS task_assigned_on, a.display_name, a.last_login, g.group_id, g.title AS group_title FROM taskaccountmap AS tam LEFT OUTER JOIN account AS a ON tam.account_id = a.account_id LEFT OUTER JOIN usergroup AS g ON a.group_id = g.group_id WHERE tam.task_id=?", query.Task_id).QueryRows(&accountDetails)
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

		for i := range accountDetails {
			accountDN = nil
			_, err = o.Raw("SELECT display_name FROM account WHERE account_id=?", accountDetails[i].Status_updated_by).QueryRows(&accountDN)
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
			accountDetails[i].Status_updated_by_display_name = accountDN[0].Display_name

			accountDN = nil
			_, err = o.Raw("SELECT display_name FROM account WHERE account_id=?", accountDetails[i].Task_assigned_by).QueryRows(&accountDN)
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
			accountDetails[i].Task_assigned_by_display_name = accountDN[0].Display_name

			accountDN = nil
			_, err = o.Raw("SELECT display_name FROM account WHERE account_id=?", accountDetails[i].Account_id).QueryRows(&accountDN)
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
			accountDetails[i].Display_name = accountDN[0].Display_name

			userGroup = nil
			_, err = o.Raw("SELECT title FROM usergroup WHERE group_id=?", accountDetails[i].Group_id).QueryRows(&userGroup)
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("Get Accounts API: " + err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Accounts. Unable to get the accounts.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if len(userGroup) > 0 {
				accountDetails[i].Group_title = userGroup[0].Title
			}
			if accountDetails[i].Group_title == "" {
				accountDetails[i].Group_title = "N/A"
			}
		}

		accountDN = nil
		_, err = o.Raw("SELECT display_name FROM account WHERE account_id=?", taskDetail.Created_by).QueryRows(&accountDN)
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
		taskDetail.Created_by_display_name = accountDN[0].Display_name

		accountDN = nil
		_, err = o.Raw("SELECT display_name FROM account WHERE account_id=?", taskDetail.Updated_by).QueryRows(&accountDN)
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
		taskDetail.Updated_by_display_name = accountDN[0].Display_name

		// Get array of groups
		_, err = o.Raw("SELECT DISTINCT tgm.group_id, g.title AS group_title FROM taskgroupmap AS tgm LEFT OUTER JOIN usergroup AS g ON tgm.group_id = g.group_id WHERE tgm.task_id=?", query.Task_id).QueryRows(&groups)

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

		taskDetail.Groups = groups
		taskDetail.AccountDetails = accountDetails
		e.Data["json"] = taskDetail
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "Invalid task id. Unable to get the task details."
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
		_ = logs.WriteLogs("Create Task API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Create Task API: " + err.Error())
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
		// Log the error
		_ = logs.WriteLogs("Create Task API: " + err.Error())
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
		// Log the error
		_ = logs.WriteLogs("Create Task API: " + err.Error())
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
			// Log the error
			_ = logs.WriteLogs("Create Task API: " + err.Error())
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		accounts = nil
		_, err = qsAccount.Filter("Group_id__exact", groupId).All(&accounts)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("Create Task API: " + err.Error())
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
				// Log the error
				_ = logs.WriteLogs("Create Task API: " + err.Error())
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
			// Log the error
			_ = logs.WriteLogs("Create Task API: " + err.Error())
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
		arrTgMap    []*modelTasks.Taskgroupmap
		arrTaMap    []*modelTasks.Taskaccountmap
		title       string
		description string
		accounts    []*modelAccounts.Account
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

	// Create query string for task table
	qsTask := o.QueryTable("task")
	condTaskId := orm.NewCondition()
	qsTaskgroupmap := o.QueryTable("Taskgroupmap")
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
		// Log the error
		_ = logs.WriteLogs("Update Task API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Update Task API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	userTask := new(modelTasks.TaskCreateDelete)
	err = json.Unmarshal(inputJson, &userTask)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Update Task API: " + err.Error())
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
		// Log the error
		_ = logs.WriteLogs("Update Task API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Db Error Tasks. Unable to find the task.")
		responseStatus.Error = err.Error()
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
		// Log the error
		_ = logs.WriteLogs("Update Task API: " + err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if userTask.Status == "in process" || userTask.Status == "completed" || userTask.Status == "new" {
		// Update the status of the current user
		condTaskaccountmap = nil
		condTaskaccountmap = condTaskaccountmap.And("Task_id__exact", userTask.Task_id).And("Account_id__exact", user[0].Account_id)
		qsTaskaccountmap = qsTaskaccountmap.SetCond(condTaskaccountmap)
		num, err = qsTaskaccountmap.Update(orm.Params{
			"Status":     userTask.Status,
			"Updated_by": user[0].Account_id,
			"Updated_on": time.Now(),
		})
	} else {
		// Change the groups assigned or accounts assigned or both
		if len(userTask.Groups_assigned) > 0 {
			// Get all the groups assigned for a task
			qsTaskgroupmap.Filter("Task_id__exact", userTask.Task_id).All(&arrTgMap)
			for _, tgMap := range arrTgMap {
				if !contains(userTask.Groups_assigned, tgMap.Group_id) {
					// Delete from Taskgroupmap
					_, err = qsTaskgroupmap.Filter("Task_id__exact", userTask.Task_id).Filter("Group_id__exact", tgMap.Group_id).Delete()
				}
			}

			for _, groupId := range userTask.Groups_assigned {
				if !qsTaskgroupmap.Filter("Task_id__exact", userTask.Task_id).Filter("Group_id__exact", groupId).Exist() {
					tgMap := new(modelTasks.Taskgroupmap)
					tgMap.Task_id = int(userTask.Task_id)
					tgMap.Group_id = groupId
					tgMap.Status = "new"
					tgMap.Updated_by = user[0].Account_id
					tgMap.Created_by = user[0].Account_id
					_, err := o.Insert(tgMap)
					if err != nil {
						// Log the error
						_ = logs.WriteLogs("Create Task API: " + err.Error())
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}
					accounts = nil
					_, err = qsAccount.Filter("Group_id__exact", groupId).All(&accounts)
					if err != nil {
						// Log the error
						_ = logs.WriteLogs("Create Task API: " + err.Error())
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}

					taMap := new(modelTasks.Taskaccountmap)

					for _, account := range accounts {
						taMap.Task_id = int(userTask.Task_id)
						taMap.Account_id = account.Account_id
						taMap.Status = "new"
						taMap.Updated_by = user[0].Account_id
						taMap.Created_by = user[0].Account_id
						_, err := o.Insert(taMap)
						if err != nil {
							// Log the error
							_ = logs.WriteLogs("Create Task API: " + err.Error())
							responseStatus := modelVoters.NewResponseStatus()
							responseStatus.Response = "error"
							responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
							responseStatus.Error = err.Error()
							e.Data["json"] = &responseStatus
							e.ServeJSON()
						}
					}
				}
			}
		}

		if len(userTask.Accounts_assigned) > 0 {
			// Get all the accounts assigned for a task
			qsTaskaccountmap.Filter("Task_id__exact", userTask.Task_id).All(&arrTaMap)
			for _, taMap := range arrTaMap {
				if !contains(userTask.Accounts_assigned, taMap.Account_id) {
					// Delete from Taskaccountmap
					_, err = qsTaskaccountmap.Filter("Task_id__exact", userTask.Task_id).Filter("Account_id__exact", taMap.Account_id).Delete()
				}
			}

			for _, accountId := range userTask.Accounts_assigned {
				if !qsTaskaccountmap.Filter("Task_id__exact", userTask.Task_id).Filter("Account_id__exact", accountId).Exist() {
					taMap := new(modelTasks.Taskaccountmap)
					taMap.Task_id = int(userTask.Task_id)
					taMap.Account_id = accountId
					taMap.Status = "new"
					taMap.Updated_by = user[0].Account_id
					taMap.Created_by = user[0].Account_id
					_, err := o.Insert(taMap)
					if err != nil {
						// Log the error
						_ = logs.WriteLogs("Create Task API: " + err.Error())
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}
				}
			}
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
		// Log the error
		_ = logs.WriteLogs("Delete Task API: " + err.Error())
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
		_ = logs.WriteLogs("Update Last Login in Delete Task API: " + err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelTasks.TaskQuery)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Task API: " + err.Error())
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
	if query.Status == "new" || query.Status == "in process" || query.Status == "completed" {
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
	num, err = qsTaskgroupmap.Filter("Created_by__exact", user[0].Account_id).Delete()
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Task API: " + err.Error())
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
	num, err = qsTaskaccountmap.Filter("Created_by__exact", user[0].Account_id).Delete()
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("Delete Task API: " + err.Error())
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
		// Log the error
		_ = logs.WriteLogs("Delete Task API: " + err.Error())
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
