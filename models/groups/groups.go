package groups

import "time"

/* user_group
{
    "group_id":1,
    "title":"Group Bommanahalli",
    "description":"",
    "updated_by":2,
    "created_by":2,
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00"
}
*/

type UserGroup struct {
	Group_id    int       `form:"-" orm:"pk" json:"group_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Updated_by  int       `json:"updated_by"`
    Created_by  int       `json:"created_by"`  
    Updated_on  time.Time `json:"updated_on"`
	Created_on  time.Time `json:"created_on"`
}

type GroupQuery struct {
	GroupId     []int     `json:"group_id"`
    CreatedBy   []int     `json:"created_by"`
}

/*GroupQuery
{
    "group_id":[1,2,4],
    "created_by":[2,3,4]
}
*/

/* Groups
{
    total:2
    [
        {
            "group_id":1,
            "title":"Group 1",
            "description":"The people of Bommanahalli",
            "updated_by":2,
            "created_by":2,
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        },
        {
            "group_id":2,
            "title":"Group 2",
            "description":"description goes here.",
            "updated_by":2,
            "created_by":2,            
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        }
    ]
}
*/

type Groups struct {
	Total  int64       `json:"total"`
	Groups []UserGroup `json:"groups"`
}