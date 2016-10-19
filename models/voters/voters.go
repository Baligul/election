package voters

import "time"

/* Voters
{
    total:1900234
    [
        {
            "voter_id":2,
            "ac_number":3,
            "part_number":5,
            "section_number":34,
            "serial_number_in_part":23,
            "name_english":"baligul hasan",
            "name_hindi":"बलिग",
            "relation_name_english":"Hasan",
            "relation_name_hindi":"हसन",
            "gender":"M",
            "id_card_number":"SJKJFF132JH",
            "district_name_hindi":"मुरादाबाद",
            "district_name_english":"Moradabad",
            "ac_name_english":"Moradabad Urban",
            "ac_name_hindi":"मुरादाबाद नगर",
            "section_name_english":"Civil Lines",
            "section_name_hindi":"सिविल लाइन्स",
            "religion_english":"Muslim",
            "religion_hindi":"मुसलमान",
            "age":34
            "vote":0
            "email":"abcd_example@test.com"
            "mobile_no":9898272676
            "image":"asdshku32%%7%ahssa*71212"
        }
    ]
}
*/

type Voter struct {
	Voter_id              int    `form:"-" orm:"pk" json:"voter_id"`
	Ac_number             int    `form:"acnumber" json:"ac_number"`
	Part_number           int    `form:"partnumber" json:"part_number"`
	Section_number        int    `form:"sectionnumber" json:"section_number"`
	Serial_number_in_part int    `form:"seialnumberinpart" json:"serial_number_in_part"`
	Name_english          string `form:"nameenglish" json:"name_english"`
	Name_hindi            string `form:"namehindi" json:"name_hindi"`
	Relation_name_english string `form:"relationnameenglish" json:"relation_name_english"`
	Relation_name_hindi   string `form:"relationnamehindi" json:"relation_name_hindi"`
	Gender                string `form:"gender" json:"gender"`
	Id_card_number        string `form:"idcardnumber" json:"id_card_number"`
	District_name_hindi   string `form:"districtnamehindi" json:"district_name_hindi"`
	District_name_english string `form:"districtnameenglish" json:"district_name_english"`
	Ac_name_english       string `form:"acnameenglish" json:"ac_name_english"`
	Ac_name_hindi         string `form:"acnamehindi" json:"ac_name_hindi"`
	Section_name_english  string `form:"sectionnameenglish" json:"section_name_english"`
	Section_name_hindi    string `form:"sectionnamehindi" json:"section_name_hindi"`
	Religion_english      string `form:"religionenglish" json:"religion_english"`
	Religion_hindi        string `form:"religionhindi" json:"religion_hindi"`
	Age                   int    `form:"age" json:"age"`
	Vote                  int    `form:"vote" json:"vote"`
    Email                 string `form:"email" json:"email"`
    Mobile_no             int    `form:"mobile_no" json:"mobile_no"`
    Image                 string `form:"image" json:"image"`
}

type Voters struct {
	Total  int64   `json:"total"`
	Voters []Voter `json:"voters"`
}

/* account
{
    "account_id":1,
    "otp":3344,
    "display_name":"amit",
    "email":"amit@gmail.com",
    "mobile_no":9344627378,
    "token":"a%4fFF$%jgds^&J",
    "approved_districts":[19,20],
    "approved_acs":[203,204],
    "last_login":"2014-05-16T08:28:06.801064-04:00",
    "updated_on":"2014-05-16T08:28:06.801064-04:00",
    "created_on":"2014-05-16T08:28:06.801064-04:00"
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
}

// Queries
/*{
    "query":{
        "state_number":[2,3],
        "district_number":[2,3],
        "voter_id":[1,2],
        "ac_number":[2,2],
        "part_number":[2,2],
        "section_number":[2,2],
        "serial_number_in_part":[2,2],
        "name_english":["baligul","iftekhar"],
        "name_hindi":["बलिग","इफ़्तेख़ार"],
        "relation_name_english":["Hasan","Khan"],
        "relation_name_hindi":["हसन","खान"],
        "gender":["M","F"],
        "id_card_number":["AD1213JGJ","SJKJFF132JH"],
        "district_name_hindi":["",""],
        "district_name_english":["",""],
        "ac_name_english":["",""],
        "ac_name_hindi":["",""],
        "section_name_english":["",""],
        "section_name_hindi":["",""],
        "religion_english":["",""],
        "religion_hindi":["",""],
        "age":[22,18,34],
        "vote":[0,1],
        "email":["abcd_example@test.com", "cdef_example@test.com"],
        "mobile_no":[9898272676, 8765432256],
        "image":["asdshku32%%7%ahssa*71212", "lkdfjlkd*&*nds8%hazd$hds11"]
    },
    "scope":{
        "state_number":[2,3],
        "district_number":[2,3],
        "voter_id":[1,2],
        "ac_number":[2,2],
        "part_number":[2,2],
        "section_number":[2,2],
        "serial_number_in_part":[2,2],
        "name_english":["baligul","iftekhar"],
        "name_hindi":["बलिग","इफ़्तेख़ार"],
        "relation_name_english":["Hasan","Khan"],
        "relation_name_hindi":["हसन","खान"],
        "gender":["M","F"],
        "id_card_number":["AD1213JGJ","SJKJFF132JH"],
        "district_name_hindi":["",""],
        "district_name_english":["",""],
        "ac_name_english":["",""],
        "ac_name_hindi":["",""],
        "section_name_english":["",""],
        "section_name_hindi":["",""],
        "religion_english":["",""],
        "religion_hindi":["",""],
        "age":[22,18,34],
        "vote":[0,1],
        "email":["abcd_example@test.com", "cdef_example@test.com"],
        "mobile_no":[9898272676, 8765432256],
        "image":["asdshku32%%7%ahssa*71212", "lkdfjlkd*&*nds8%hazd$hds11"]
    }
}*/

type Queries struct {
	Query Query `json:"query"`
	Scope Query `json:"scope"`
}

// Query
/*{
    "state_number":[1,2],
    "district_number":[2,3],
    "voter_id":[1,2],
    "ac_number":[2,2],
    "part_number":[2,2],
    "section_number":[2,2],
    "serial_number_in_part":[2,2],
    "name_english":["baligul","iftekhar"],
    "name_hindi":["बलिग","इफ़्तेख़ार"],
    "relation_name_english":["Hasan","Khan"],
    "relation_name_hindi":["हसन","खान"],
    "gender":["M","F"],
    "id_card_number":["AD1213JGJ","SJKJFF132JH"],
    "district_name_hindi":["",""],
    "district_name_english":["",""],
    "ac_name_english":["",""],
    "ac_name_hindi":["",""],
    "section_name_english":["",""],
    "section_name_hindi":["",""],
    "religion_english":["",""],
    "religion_hindi":["",""],
    "age":[22,18,34],
    "vote":[0,1],
    "email":["abcd_example@test.com", "cdef_example@test.com"],
    "mobile_no":[9898272676, 8765432256],
    "image":["asdshku32%%7%ahssa*71212", "lkdfjlkd*&*nds8%hazd$hds11"]
}*/

type Query struct {
	StateNumber         []int    `json:"state_number"`
	DistrictNumber      []int    `json:"district_number"`
	VoterID             []int    `json:"voter_id"`
	AcNumber            []int    `json:"ac_number"`
	PartNumber          []int    `json:"part_number"`
	SectionNumber       []int    `json:"section_number"`
	SerialNumberInPart  []int    `json:"serial_number_in_part"`
	NameEnglish         []string `json:"name_english"`
	NameHindi           []string `json:"name_hindi"`
	RelationNameEnglish []string `json:"relation_name_english"`
	RelationNameHindi   []string `json:"relation_name_hindi"`
	Gender              []string `json:"gender"`
	IDCardNumber        []string `json:"id_card_number"`
	DistrictNameHindi   []string `json:"district_name_hindi"`
	DistrictNameEnglish []string `json:"district_name_english"`
	AcNameEnglish       []string `json:"ac_name_english"`
	AcNameHindi         []string `json:"ac_name_hindi"`
	SectionNameEnglish  []string `json:"section_name_english"`
	SectionNameHindi    []string `json:"section_name_hindi"`
	ReligionEnglish     []string `json:"religion_english"`
	ReligionHindi       []string `json:"religion_hindi"`
	Age                 []int    `json:"age"`
	Vote                []int    `json:"vote"`
    Email               []string `json:"email"`
    MobileNo            []int    `json:"mobile_no"`
    Image               []string `json:"image"`
}

/* Statistic
{
    "total":213243,
    "muslim":23243,
    "other":2343,
    "male":2343,
    "female":2343,
    "muslim_male":2343,
    "muslim_female":2343,
    "other_male":2343,
    "other_female":2343,
    "muslim_p":10.89,
    "other_p":10.89,
    "male_p":10.89,
    "female_p":10.89,
    "muslim_male_p":10.89,
    "muslim_female_p":10.89,
    "other_male_p":10.89,
    "other_female_p":10.89
}
*/

type Statistic struct {
	Total         int64   `json:"total"`
	Muslim        int64   `json:"muslim"`
	Others        int64   `json:"other"`
	Male          int64   `json:"male"`
	Female        int64   `json:"female"`
	MuslimMale    int64   `json:"muslim_male"`
	MuslimFemale  int64   `json:"muslim_female"`
	OthersMale    int64   `json:"other_male"`
	OthersFemale  int64   `json:"other_female"`
	MuslimP       float64 `json:"muslim_p"`
	OthersP       float64 `json:"other_p"`
	MaleP         float64 `json:"male_p"`
	FemaleP       float64 `json:"female_p"`
	MuslimMaleP   float64 `json:"muslim_male_p"`
	MuslimFemaleP float64 `json:"muslim_female_p"`
	OthersMaleP   float64 `json:"other_male_p"`
	OthersFemaleP float64 `json:"other_female_p"`
}

type Statistics struct {
	Total  int64     `json:"total"`
	Result int64     `json:"result"`
	Query  Statistic `json:"query"`
	Scope  Statistic `json:"scope"`
}

type ResponseStatus struct {
	Response string `json:"response"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

type List struct {
	Districts []string `json:"districts"`
	Acs       []string `json:"acs"`
}

type UpdateVote struct {
	District string `json:"district"`
	VoterID  []int  `json:"voter_id"`
	Vote     int    `json:"vote"`
    Email    string `json:"email"`
	MobileNo int64  `json:"mobile_no"`
    Image    string `json:"image"`
}

func NewResponseStatus() *ResponseStatus {
	return new(ResponseStatus)
}

func GetTableName(districtName string) string {
	switch districtName {
	case "Rampur":
		return "voter_20"
	case "Moradabad":
		return "voter_19"
	case "Bangalore":
		return "voter_21"
		//return "voter"
	case "Bijnor":
		return "voter_21"
	case "Hubli":
		return "voter_21"
		//return "voter"
	default:
		return "voter_21"
		//return "voter"
	}
	return "voter_21"
	//return "voter"
}

type ReadJson struct {
	District    string `json:"district"`
	AcNum       int    `json:"ac_num"`
	AcName      string `json:"ac_name"`
	SectionName string `json:"section_name"`
}

type ReadJsons struct {
	ReadJsons []ReadJson
}
