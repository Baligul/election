package accounts

import "time"

/* account
{
    "account_id":1,
    "otp":3344,
    "display_name":"amit",
    "email":"amit@gmail.com",
    "mobile_no":9344627378,
    "token":"a%4fFF$%jgds^&J",
    "approved_districts":["Moradabad"],
    "approved_acs":["Kanth","Bilari"],
    "last_login":"2014-05-16T08:28:06.801064-04:00",
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00",
    "role":"leader",
    "image":"",
    "group_id":5,
    "leader_id":2
}
*/

type Account struct {
	Account_id         int       `orm:"pk" json:"account_id"`
	Otp                int       `json:"otp"`
	Display_name       string    `json:"display_name"`
	Email              string    `json:"email"`
	Mobile_no          int64     `json:"mobile_no"`
	Token              string    `json:"token"`
	Approved_districts string    `json:"approved_districts"`
	Approved_acs       string    `json:"approved_acs"`
	Last_login         time.Time `json:"last_login"`
	Updated_on         time.Time `json:"updated_on"`
	Created_on         time.Time `json:"created_on"`
    Role               string    `json:"role"`
    Image              string    `json:"image"`
    Group_id           int       `json:"group_id"`
    Leader_id          int       `json:"leader_id"`
}

type AccountQuery struct {
    AccountId   []int     `json:"account_id"`
	GroupId     []int     `json:"group_id"`
    LeaderId    []int     `json:"leader_id"`
}

/*AccountQuery
{
    "account_id":[2,3,6],
    "group_id":[1,2,4],
    "leader_id":[1]
}
*/

/* Accounts
{
    total:2
    [
        {
            "account_id":1,
            "otp":3344,
            "display_name":"amit",
            "email":"amit@gmail.com",
            "mobile_no":9344627378,
            "token":"a%4fFF$%jgds^&J",
            "approved_districts":["Moradabad"],
            "approved_acs":["Kanth","Bilari"],
            "last_login":"2014-05-16T08:28:06.801064-04:00",
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00",
            "role":"leader",
            "image":"",
            "group_id":5,
            "leader_id":2
        },
        {
            "account_id":2,
            "otp":3344,
            "display_name":"amit",
            "email":"amit@gmail.com",
            "mobile_no":9344627378,
            "token":"a%4fFF$%jgds^&J",
            "approved_districts":["Moradabad"],
            "approved_acs":["Kanth","Bilari"],
            "last_login":"2014-05-16T08:28:06.801064-04:00",
            "updated_on":"2014-05-16T08:28:06.801064-04:00",
            "created_on":"2014-05-16T08:28:06.801064-04:00",
            "role":"leader",
            "image":"",
            "group_id":5,
            "leader_id":2
        }
    ]
}
*/

type Accounts struct {
	Total  int64       `json:"total"`
	Accounts []Account `json:"accounts"`
}

func (accounts *Accounts) Populate(useraccounts []*Account) {
	for _, account := range useraccounts {
		accounts.Accounts = append(accounts.Accounts, *account)
	}
}