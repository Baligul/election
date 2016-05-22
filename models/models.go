package models

type Voter struct {
    Voter_id                    int `form:"-" orm:"pk"`
    Ac_number                   int `form:"acnumber"`
    Part_number                 int `form:"partnumber"`
    Section_number              int `form:"sectionnumber"`
    Serial_number_in_part       int `form:"seialnumberinpart"`
    Name_english                string `form:"nameenglish"`
    Name_hindi                  string `form:"namehindi"`
    Relation_name_english       string `form:"relationnameenglish"`
    Relation_name_hindi         string `form:"relationnamehindi"`
    Gender                      string `form:"gender"`
    Id_card_number              string `form:"idcardnumber"`
    District_name_hindi         string `form:"districtnameenglish"`
    District_name_english       string `form:"districtnamehindi"`
    Ac_name_english             string `form:"acnameenglish"`
    Ac_name_hindi               string `form:"acnamehindi"`
    Section_name_english        string `form:"sectionnameenglish"`
    Section_name_hindi          string `form:"sectionnamehindi"`
}

func (a *Voter) TableName() string {
    return "voter"
}