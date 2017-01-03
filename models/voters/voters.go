package voters

/* Voters
{
    "total":1900234,
    "voters":[
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
            "age":34,
            "vote":0,
            "email":"abcd_example@test.com",
            "mobile_no":9898272676,
            "image":"asdshku32%%7%ahssa*71212",
            "updated_on":""
        }
    ]
}
*/

type Voter struct {
	Voter_id              int    `form:"-" orm:"auto" json:"voter_id,omitempty"`
	Ac_number             int    `form:"acnumber" json:"ac_number,omitempty"`
	Part_number           int    `form:"partnumber" json:"part_number,omitempty"`
	Section_number        int    `form:"sectionnumber" json:"section_number,omitempty"`
	Serial_number_in_part int    `form:"seialnumberinpart" json:"serial_number_in_part,omitempty"`
	Name_english          string `form:"nameenglish" json:"name_english,omitempty"`
	Name_hindi            string `form:"namehindi" json:"name_hindi,omitempty"`
	Relation_name_english string `form:"relationnameenglish" json:"relation_name_english,omitempty"`
	Relation_name_hindi   string `form:"relationnamehindi" json:"relation_name_hindi,omitempty"`
	Gender                string `form:"gender" json:"gender,omitempty"`
	Id_card_number        string `form:"idcardnumber" json:"id_card_number,omitempty"`
	District_name_hindi   string `form:"districtnamehindi" json:"district_name_hindi,omitempty"`
	District_name_english string `form:"districtnameenglish" json:"district_name_english,omitempty"`
	Ac_name_english       string `form:"acnameenglish" json:"ac_name_english,omitempty"`
	Ac_name_hindi         string `form:"acnamehindi" json:"ac_name_hindi,omitempty"`
	Section_name_english  string `form:"sectionnameenglish" json:"section_name_english,omitempty"`
	Section_name_hindi    string `form:"sectionnamehindi" json:"section_name_hindi,omitempty"`
	Religion_english      string `form:"religionenglish" json:"religion_english,omitempty"`
	Religion_hindi        string `form:"religionhindi" json:"religion_hindi,omitempty"`
	Age                   int    `form:"age" json:"age,omitempty"`
	Vote                  int    `form:"vote" json:"vote,omitempty"`
	Email                 string `form:"email" json:"email,omitempty"`
	Mobile_no             int64  `form:"mobile_no" json:"mobile_no,omitempty"`
	Image                 string `form:"image" json:"image,omitempty"`
    Updated_on            string `json:"updated_on,omitempty"`
}

// ByName implements sort.Interface for []Voter based on
// the Name_english field.
type ByName []Voter

func (v ByName) Len() int           { return len(v) }
func (v ByName) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByName) Less(i, j int) bool { return v[i].Name_english < v[j].Name_english }

type Voters struct {
	Total  int64  `json:"total,omitempty"`
	Voters ByName `json:"voters,omitempty"`
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
	Query Query `json:"query,omitempty"`
	Scope Query `json:"scope,omitempty"`
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
	StateNumber         []int    `json:"state_number,omitempty"`
	DistrictNumber      []int    `json:"district_number,omitempty"`
	VoterID             []int    `json:"voter_id,omitempty"`
	AcNumber            []int    `json:"ac_number,omitempty"`
	PartNumber          []int    `json:"part_number,omitempty"`
	SectionNumber       []int    `json:"section_number,omitempty"`
	SerialNumberInPart  []int    `json:"serial_number_in_part,omitempty"`
	NameEnglish         []string `json:"name_english,omitempty"`
	NameHindi           []string `json:"name_hindi,omitempty"`
	RelationNameEnglish []string `json:"relation_name_english,omitempty"`
	RelationNameHindi   []string `json:"relation_name_hindi,omitempty"`
	Gender              []string `json:"gender,omitempty"`
	IDCardNumber        []string `json:"id_card_number,omitempty"`
	DistrictNameHindi   []string `json:"district_name_hindi,omitempty"`
	DistrictNameEnglish []string `json:"district_name_english,omitempty"`
	AcNameEnglish       []string `json:"ac_name_english,omitempty"`
	AcNameHindi         []string `json:"ac_name_hindi,omitempty"`
	SectionNameEnglish  []string `json:"section_name_english,omitempty"`
	SectionNameHindi    []string `json:"section_name_hindi,omitempty"`
	ReligionEnglish     []string `json:"religion_english,omitempty"`
	ReligionHindi       []string `json:"religion_hindi,omitempty"`
	Age                 []int    `json:"age,omitempty"`
	Vote                []int    `json:"vote,omitempty"`
	Email               []string `json:"email,omitempty"`
	MobileNo            []int    `json:"mobile_no,omitempty"`
	Image               []string `json:"image,omitempty"`
}

/* Statistic
{
    "total":213243,
    "muslim":23243,
    "dalit":2434,
    "other":2343,
    "male":2343,
    "female":2343,
    "muslim_male":2343,
    "muslim_female":2343,
    "dalit_male":2343,
    "dalit_female":2343,
    "other_male":2343,
    "other_female":2343,
    "muslim_p":10.89,
    "dalit_p":10.89,
    "other_p":10.89,
    "male_p":10.89,
    "female_p":10.89,
    "muslim_male_p":10.89,
    "muslim_female_p":10.89,
    "dalit_male_p":10.89,
    "dalit_female_p":10.89,
    "other_male_p":10.89,
    "other_female_p":10.89
}
*/

type Statistic struct {
	Total         int64   `json:"total,omitempty"`
	Muslim        int64   `json:"muslim,omitempty"`
	Dalit         int64   `json:"dalit,omitempty"`
	Others        int64   `json:"other,omitempty"`
	Male          int64   `json:"male,omitempty"`
	Female        int64   `json:"female,omitempty"`
	MuslimMale    int64   `json:"muslim_male,omitempty"`
	MuslimFemale  int64   `json:"muslim_female,omitempty"`
	DalitMale     int64   `json:"dalit_male,omitempty"`
	DalitFemale   int64   `json:"dalit_female,omitempty"`
	OthersMale    int64   `json:"other_male,omitempty"`
	OthersFemale  int64   `json:"other_female,omitempty"`
	MuslimP       float64 `json:"muslim_p,omitempty"`
	DalitP        float64 `json:"dalit_p,omitempty"`
	OthersP       float64 `json:"other_p,omitempty"`
	MaleP         float64 `json:"male_p,omitempty"`
	FemaleP       float64 `json:"female_p,omitempty"`
	MuslimMaleP   float64 `json:"muslim_male_p,omitempty"`
	MuslimFemaleP float64 `json:"muslim_female_p,omitempty"`
	DalitMaleP    float64 `json:"dalit_male_p,omitempty"`
	DalitFemaleP  float64 `json:"dalit_female_p,omitempty"`
	OthersMaleP   float64 `json:"other_male_p,omitempty"`
	OthersFemaleP float64 `json:"other_female_p,omitempty"`
}

type Statistics struct {
	Total  int64     `json:"total,omitempty"`
	Result int64     `json:"result,omitempty"`
	Query  Statistic `json:"query,omitempty"`
	Scope  Statistic `json:"scope,omitempty"`
}

type ResponseStatus struct {
	Response string `json:"response,omitempty"`
	Message  string `json:"message,omitempty"`
	Error    string `json:"error,omitempty"`
}

type List struct {
	Districts []string `json:"districts,omitempty"`
	Acs       []string `json:"acs,omitempty"`
}

type UpdateVoter struct {
	District string `json:"district,omitempty"`
	VoterID  []int  `json:"voter_id,omitempty"`
	Vote     int    `json:"vote,omitempty"`
	Email    string `json:"email,omitempty"`
	MobileNo int64  `json:"mobile_no,omitempty"`
	Image    string `json:"image,omitempty"`
	Age      int    `json:"age,omitempty"`
	Religion string `json:"religion,omitempty"`
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
	District    string `json:"district,omitempty"`
	AcNum       int    `json:"ac_num,omitempty"`
	AcName      string `json:"ac_name,omitempty"`
	SectionName string `json:"section_name,omitempty"`
}

type ReadJsons struct {
	ReadJsons []ReadJson
}
