package accounts

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
	Account_id         int       `orm:"auto" json:"account_id,omitempty"`
	Otp                int       `json:"otp,omitempty"`
	Display_name       string    `json:"display_name,omitempty"`
	Email              string    `json:"email,omitempty"`
	Mobile_no          int64     `json:"mobile_no,omitempty"`
	Token              string    `json:"token,omitempty"`
	Approved_districts string    `json:"approved_districts,omitempty"`
	Approved_acs       string    `json:"approved_acs,omitempty"`
	Last_login         string    `json:"last_login,omitempty"`
	Updated_on         string    `json:"updated_on,omitempty"`
	Created_on         string    `json:"created_on,omitempty"`
	Role               string    `json:"role,omitempty"`
	Image              string    `json:"image,omitempty"`
	Group_id           int       `json:"group_id,omitempty"`
	Leader_id          int       `json:"leader_id,omitempty"`
	Age                int       `json:"age,omitempty"`
	Sex                string    `json:"sex,omitempty"`
	Religion           string    `json:"religion,omitempty"`
	Father_name        string    `json:"father_name,omitempty"`
	Approved_sections  string    `json:"approved_sections,omitempty"`
	Approved_booths    string    `json:"approved_booths,omitempty"`
	Group_title        string    `json:"group_title,omitempty"`
}

type AccountQuery struct {
	AccountId []int `json:"account_id,omitempty"`
	GroupId   []int `json:"group_id,omitempty"`
	LeaderId  []int `json:"leader_id,omitempty"`
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

// ByDisplayName implements sort.Interface for []Account based on
// the Display_name field.
type ByDisplayName []Account

func (a ByDisplayName) Len() int           { return len(a) }
func (a ByDisplayName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDisplayName) Less(i, j int) bool { return a[i].Display_name < a[j].Display_name }

type Accounts struct {
	Total    int64         `json:"total,omitempty"`
	Accounts ByDisplayName `json:"accounts,omitempty"`
}

func (accounts *Accounts) Populate(useraccounts ByDisplayName) {
	for _, account := range useraccounts {
		accounts.Accounts = append(accounts.Accounts, account)
	}
}
