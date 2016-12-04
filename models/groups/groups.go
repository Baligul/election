package groups

import "time"

/* user_group
{
    "group_id":1,
    "title":"Group Bommanahalli",
    "description":"",
    "group_lead_id":1,
    "updated_by":2,
    "created_by":2,
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00"
}
*/

type Usergroup struct {
	Group_id      int       `form:"-" orm:"auto" json:"group_id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	Group_lead_id int       `json:"group_lead_id,omitempty"`
	Updated_by    int       `json:"updated_by,omitempty"`
	Created_by    int       `json:"created_by,omitempty"`
	Updated_on    time.Time `orm:"auto_now;type(datetime)" json:"updated_on,omitempty"`
	Created_on    time.Time `orm:"auto_now_add;type(datetime)" json:"created_on,omitempty"`
    Account_id    []int     `orm:"-" json:"Account_id,omitempty"`
}

type GroupQuery struct {
	GroupId     []int `json:"group_id,omitempty"`
	CreatedBy   []int `json:"created_by,omitempty"`
	GroupLeadId []int `json:"group_lead_id,omitempty"`
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
            "group_lead_id":1,
            "updated_by":2,
            "created_by":2,
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        },
        {
            "group_id":2,
            "title":"Group 2",
            "description":"description goes here.",
            "group_lead_id":2,
            "updated_by":2,
            "created_by":2,
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00"
        }
    ]
}
*/

type Groups struct {
	Total  int64       `json:"total,omitempty"`
	Groups []Usergroup `json:"groups,omitempty"`
}

func (groups *Groups) Populate(usergroups []*Usergroup) {
	for _, group := range usergroups {
		groups.Groups = append(groups.Groups, *group)
	}
}
