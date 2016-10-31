package tasks

import "time"

/* task
{
    "task_id":1,
    "title":"Get voters Mobile numbers",
    "description":"",
    "assigned_to":23,
    "assigned_to":"completed",
    "mobile_no":9344627378,
    "token":"a%4fFF$%jgds^&J",
    "approved_districts":[19,20],
    "approved_acs":[203,204],
    "last_login":"2014-05-16T08:28:06.801064-04:00",
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00"
}
*/

type Task struct {
	Task_id           int       `form:"-" orm:"pk" json:"task_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Groups_assigned   string    `json:"groups_assigned"`
	Accounts_assigned string    `json:"tasks_assigned"`
	Status            string    `json:"status"`
	Updated_by        int       `json:"updated_by"`
	Created_by        int       `json:"created_by"`
	Updated_on        time.Time `json:"updated_on"`
	Created_on        time.Time `json:"created_on"`
}

/* Tasks
{
    total:2
    [
        {
            "task_id":1,
            "title":"Get voters Mobile numbers",
            "description":"Take the list of booth 1 and collect the voters mobile numbers",
            "assigned_to":23,
            "status":"new",
            "updated_by":2,
            "created_by":2,
            "updated_on":,
            "created_on":
        },
        {
            "task_id":2,
            "title":"Get voters Mobile numbers",
            "description":"Take the list of booth 2 and collect the voters mobile numbers",
            "assigned_to":24,
            "status":"new",
            "updated_by":2,
            "created_by":2,
            "updated_on":,
            "created_on":
        }
    ]
}
*/
type Tasks struct {
	Total int64  `json:"total"`
	Tasks []Task `json:"tasks"`
}
