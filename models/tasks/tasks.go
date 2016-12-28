package tasks

import "time"

/* task
{
    "task_id":1,
    "title":"Get voters Mobile numbers",
    "description":"",
    "groups_assigned":[],
    "acounts_assigned":[],
    "status":"New|In process|Complete",
    "updated_by":2,
    "created_by":2,
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00"
}
*/

type Task struct {
	Task_id     int       `form:"-" orm:"auto" json:"task_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Updated_by  int       `json:"updated_by,omitempty"`
	Created_by  int       `json:"created_by,omitempty"`
	Updated_on  time.Time `orm:"auto_now;type(datetime)" json:"updated_on,omitempty"`
	Created_on  time.Time `orm:"auto_now_add;type(datetime)" json:"created_on,omitempty"`
}

type Taskgroupmap struct {
	Task_group_map_id int       `form:"-" orm:"auto" json:"task_group_map_id,omitempty"`
	Task_id           int       `json:"task_id,omitempty"`
	Group_id          int       `json:"group_id,omitempty"`
	Status            string    `json:"status,omitempty"`
	Updated_by        int       `json:"updated_by,omitempty"`
	Created_by        int       `json:"created_by,omitempty"`
	Updated_on        time.Time `orm:"auto_now;type(datetime)" json:"updated_on,omitempty"`
	Created_on        time.Time `orm:"auto_now_add;type(datetime)" json:"created_on,omitempty"`
}

type Taskaccountmap struct {
	Task_account_map_id int       `form:"-" orm:"auto" json:"task_account_map_id,omitempty"`
	Task_id             int       `json:"task_id,omitempty"`
	Account_id          int       `json:"account_id,omitempty"`
	Status              string    `json:"status,omitempty"`
	Updated_by          int       `json:"updated_by,omitempty"`
	Created_by          int       `json:"created_by,omitempty"`
	Updated_on          time.Time `orm:"auto_now;type(datetime)" json:"updated_on,omitempty"`
	Created_on          time.Time `orm:"auto_now_add;type(datetime)" json:"created_on,omitempty"`
}

type TaskQuery struct {
	TaskId           []int  `json:"task_id,omitempty"`
	GroupsAssigned   []int  `json:"groups_assigned,omitempty"`
	AccountsAssigned []int  `json:"accounts_assigned,omitempty"`
	Status           string `json:"status,omitempty"`
	UpdatedBy        []int  `json:"created_by,omitempty"`
	CreatedBy        []int  `json:"created_by,omitempty"`
}

type TaskCreateDelete struct {
	Task_id           int       `json:"task_id,omitempty"`
	Title             string    `json:"title,omitempty"`
	Description       string    `json:"description,omitempty"`
	Groups_assigned   []int     `json:"groups_assigned,omitempty"`
	Accounts_assigned []int     `json:"accounts_assigned,omitempty"`
	Status            string    `json:"status,omitempty"`
	Updated_by        int       `json:"updated_by,omitempty"`
	Created_by        int       `json:"created_by,omitempty"`
	Updated_on        time.Time `json:"updated_on,omitempty"`
	Created_on        time.Time `json:"created_on,omitempty"`
}

type TaskDetail struct {
	Task_id                 int               `json:"task_id,omitempty"`
	Title                   string            `json:"title,omitempty"`
	Description             string            `json:"description,omitempty"`
	AccountDetails          []*AccountDetails `json:"account_details,omitempty"`
	Updated_by              int               `json:"updated_by,omitempty"`
	Created_by              int               `json:"created_by,omitempty"`
	Updated_by_display_name string            `json:"updated_by_display_name,omitempty"`
	Created_by_display_name string            `json:"created_by_display_name,omitempty"`
	Updated_on              time.Time         `json:"updated_on,omitempty"`
	Created_on              time.Time         `json:"created_on,omitempty"`
}

type AccountDetails struct {
	Account_id                     int       `json:"account_id,omitempty"`
	Status                         string    `json:"status,omitempty"`
	Status_updated_by              int       `json:"status_updated_by,omitempty"`
	Status_updated_by_display_name string    `json:"status_updated_by_display_name,omitempty"`
	Status_updated_on              time.Time `json:"status_updated_on,omitempty"`
	Task_assigned_by               int       `json:"task_assigned_by,omitempty"`
	Task_assigned_by_display_name  string    `json:"task_assigned_by_display_name,omitempty"`
	Task_assigned_on               time.Time `json:"task_assigned_on,omitempty"`
	Display_name                   string    `json:"display_name,omitempty"`
	Group_id                       int       `json:"group_id,omitempty"`
	Group_title                    string    `json:"group_title,omitempty"`
}

type AccountDisplayName struct {
	Display_name string `json:"display_name,omitempty"`
}

/*GroupQuery
{
    "group_id":[1,2,4],
    "created_by":[2,3,4]
}
*/

/* Tasks
{
    total:2
    [
        {
            "task_id":1,
            "title":"Get voters Mobile numbers",
            "description":"",
            "groups_assigned":[],
            "acounts_assigned":[],
            "status":"New|In process|Complete",
            "updated_by":2,
            "created_by":2,
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        },
        {
            "task_id":1,
            "title":"Get voters Mobile numbers",
            "description":"",
            "groups_assigned":[],
            "acounts_assigned":[],
            "status":"New|In process|Complete",
            "updated_by":2,
            "created_by":2,
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        }
    ]
}
*/

// ByTitle implements sort.Interface for []Task based on
// the Title field.
type ByTitle []Task

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return a[i].Title < a[j].Title }

type Tasks struct {
	Total int64   `json:"total,omitempty"`
	Tasks ByTitle `json:"tasks,omitempty"`
}

func (tasks *Tasks) Populate(tasksList ByTitle) {
	for _, task := range tasksList {
		tasks.Tasks = append(tasks.Tasks, task)
	}
}

func (task *Task) Transpose(taskCreateDelete *TaskCreateDelete) {
	task.Task_id = taskCreateDelete.Task_id
	task.Title = taskCreateDelete.Title
	task.Description = taskCreateDelete.Description
	task.Updated_by = taskCreateDelete.Updated_by
	task.Created_by = taskCreateDelete.Created_by
	task.Updated_on = taskCreateDelete.Updated_on
	task.Created_on = taskCreateDelete.Created_on
}

func (taskDetail *TaskDetail) Populate(task *Task) {
	taskDetail.Task_id = task.Task_id
	taskDetail.Title = task.Title
	taskDetail.Description = task.Description
	taskDetail.Updated_by = task.Updated_by
	taskDetail.Created_by = task.Created_by
	taskDetail.Updated_on = task.Updated_on
	taskDetail.Created_on = task.Created_on
}
