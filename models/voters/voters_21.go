package voters

type Voter_21 struct {
	Voter_id              int    `form:"-" orm:"pk" json:"voter_id"`
	Ac_number             int    `form:"acnumber" json:"ac_number"`
	Part_number           int    `form:"partnumber" json:"part_number"`
	Section_number        int    `form:"sectionnumber" json:"section_number"`
	Serial_number_in_part int    `form:"seialnumberinpart json:"serial_number_in_part"`
	Name_english          string `form:"nameenglish json:"name_english"`
	Name_hindi            string `form:"namehindi json:"name_hindi"`
	Relation_name_english string `form:"relationnameenglish json:"relation_name_english"`
	Relation_name_hindi   string `form:"relationnamehindi json:"relation_name_hindi"`
	Gender                string `form:"gender json:"gender"`
	Id_card_number        string `form:"idcardnumber json:"id_card_number"`
	District_name_hindi   string `form:"districtnamehindi json:"district_name_hindi"`
	District_name_english string `form:"districtnameenglish json:"district_name_english"`
	Ac_name_english       string `form:"acnameenglish json:"ac_name_english"`
	Ac_name_hindi         string `form:"acnamehindi json:"ac_name_hindi"`
	Section_name_english  string `form:"sectionnameenglish json:"section_name_english"`
	Section_name_hindi    string `form:"sectionnamehindi json:"section_name_hindi"`
	Religion_english      string `form:"religionenglish json:"religion_english"`
	Religion_hindi        string `form:"religionhindi json:"religion_hindi"`
	Age                   int    `form:"age json:"age"`
}

func (voter_21 *Voter_21) transpose() Voter {
	voter := new(Voter)
	voter.Voter_id = voter_21.Voter_id
	voter.Ac_number = voter_21.Ac_number
	voter.Part_number = voter_21.Part_number
	voter.Section_number = voter_21.Section_number
	voter.Serial_number_in_part = voter_21.Serial_number_in_part
	voter.Name_english = voter_21.Name_english
	voter.Name_hindi = voter_21.Name_hindi
	voter.Relation_name_english = voter_21.Relation_name_english
	voter.Relation_name_hindi = voter_21.Relation_name_hindi
	voter.Gender = voter_21.Gender
	voter.Id_card_number = voter_21.Id_card_number
	voter.District_name_hindi = voter_21.District_name_hindi
	voter.District_name_english = voter_21.District_name_english
	voter.Ac_name_english = voter_21.Ac_name_english
	voter.Ac_name_hindi = voter_21.Ac_name_hindi
	voter.Section_name_english = voter_21.Section_name_english
	voter.Section_name_hindi = voter_21.Section_name_hindi
	voter.Religion_english = voter_21.Religion_english
	voter.Religion_hindi = voter_21.Religion_hindi
	voter.Age = voter_21.Age

	return *voter
}

func (voters *Voters) Populate_21(voters_21 []*Voter_21) {
	for _, voter_21 := range voters_21 {
		voter := voter_21.transpose()
		voters.Voters = append(voters.Voters, voter)
	}
}
