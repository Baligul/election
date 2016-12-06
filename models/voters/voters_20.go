package voters

import "time"

type Voter_20 struct {
	Voter_id              int       `form:"-" orm:"auto" json:"voter_id,omitempty"`
	Ac_number             int       `form:"acnumber" json:"ac_number,omitempty"`
	Part_number           int       `form:"partnumber" json:"part_number,omitempty"`
	Section_number        int       `form:"sectionnumber" json:"section_number,omitempty"`
	Serial_number_in_part int       `form:"seialnumberinpart json:"serial_number_in_part,omitempty"`
	Name_english          string    `form:"nameenglish json:"name_english,omitempty"`
	Name_hindi            string    `form:"namehindi json:"name_hindi,omitempty"`
	Relation_name_english string    `form:"relationnameenglish json:"relation_name_english,omitempty"`
	Relation_name_hindi   string    `form:"relationnamehindi json:"relation_name_hindi,omitempty"`
	Gender                string    `form:"gender json:"gender,omitempty"`
	Id_card_number        string    `form:"idcardnumber json:"id_card_number,omitempty"`
	District_name_hindi   string    `form:"districtnamehindi json:"district_name_hindi,omitempty"`
	District_name_english string    `form:"districtnameenglish json:"district_name_english,omitempty"`
	Ac_name_english       string    `form:"acnameenglish json:"ac_name_english,omitempty"`
	Ac_name_hindi         string    `form:"acnamehindi json:"ac_name_hindi,omitempty"`
	Section_name_english  string    `form:"sectionnameenglish json:"section_name_english,omitempty"`
	Section_name_hindi    string    `form:"sectionnamehindi json:"section_name_hindi,omitempty"`
	Religion_english      string    `form:"religionenglish json:"religion_english,omitempty"`
	Religion_hindi        string    `form:"religionhindi json:"religion_hindi,omitempty"`
	Age                   int       `form:"age json:"age,omitempty"`
	Vote                  int       `form:"vote json:"vote,omitempty"`
	Email                 string    `form:"email" json:"email,omitempty"`
	Mobile_no             int64     `form:"mobile_no" json:"mobile_no,omitempty"`
	Image                 string    `form:"image" json:"image,omitempty"`
	Updated_on            time.Time `orm:"auto_now;type(datetime)" json:"updated_on,omitempty"`
}

func (voter_20 *Voter_20) transpose() Voter {
	voter := new(Voter)
	voter.Voter_id = voter_20.Voter_id
	voter.Ac_number = voter_20.Ac_number
	voter.Part_number = voter_20.Part_number
	voter.Section_number = voter_20.Section_number
	voter.Serial_number_in_part = voter_20.Serial_number_in_part
	voter.Name_english = voter_20.Name_english
	voter.Name_hindi = voter_20.Name_hindi
	voter.Relation_name_english = voter_20.Relation_name_english
	voter.Relation_name_hindi = voter_20.Relation_name_hindi
	voter.Gender = voter_20.Gender
	voter.Id_card_number = voter_20.Id_card_number
	voter.District_name_hindi = voter_20.District_name_hindi
	voter.District_name_english = voter_20.District_name_english
	voter.Ac_name_english = voter_20.Ac_name_english
	voter.Ac_name_hindi = voter_20.Ac_name_hindi
	voter.Section_name_english = voter_20.Section_name_english
	voter.Section_name_hindi = voter_20.Section_name_hindi
	voter.Religion_english = voter_20.Religion_english
	voter.Religion_hindi = voter_20.Religion_hindi
	voter.Age = voter_20.Age
	voter.Vote = voter_20.Vote
	voter.Email = voter_20.Email
	voter.Mobile_no = voter_20.Mobile_no
	voter.Image = voter_20.Image

	return *voter
}

func (voters *Voters) Populate_20(voters_20 []*Voter_20) {
	for _, voter_20 := range voters_20 {
		voter := voter_20.transpose()
		voters.Voters = append(voters.Voters, voter)
	}
}
