/*
   GET VOTERS
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/voters

   GET STATISTIC
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://localhost:80/api/statistic

   GET STATISTICS
   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[19,20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://localhost:80/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://localhost:80/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":["Muslim"],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://localhost:80/api/statistics

   GET OTP
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9564783954}' http://localhost:80/api/otp

   Register
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9343352734, "otp":23435}' http://localhost:80/api/register

   Get List
   curl -X POST -H "Content-Type: application/json" -d '{"districts": [19,20], "acs":[34, 43]}' http://localhost:80/api/list

   Set Voter
   curl -X POST -H "Content-Type: application/json" -d '{"district":"Moradabad", "voter_id": [12345,20045], "vote":1, "email":"example@example.com","mobile_no":9456732819,"image":"jsdjsd22ksndsndsnk22knknlcxx"}' http://localhost:80/api/voter

   Read Json
   curl -X POST -H "Content-Type: application/json" -d @json/data.json http://localhost:80/api/read/json
*/

package controllers

import (
	token "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	modelVoters "github.com/Baligul/election/models/voters"
	modelTasks "github.com/Baligul/election/models/tasks"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/craigmj/gototp"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//orm.RegisterDataBase("default", "postgres", "postgres://member:hu123*Member@http://104.197.6.26:5432/election")
	orm.RegisterDataBase("default", "postgres", "user=member dbname=election sslmode=disable")
	orm.RegisterModel(new(modelVoters.Account), 
					  new(modelVoters.Voter), 
					  new(modelVoters.Voter_19), 
					  new(modelVoters.Voter_20), 
					  new(modelVoters.Voter_21), 
					  new(modelTasks.Task))
}

type ElectionController struct {
	beego.Controller
}

func (e *ElectionController) GetVoters() {
	var (
		votersCountRampur    int64
		votersCountMoradabad int64
		votersCountBijnor    int64
		votersCountBangalore int64
		votersCountHubli     int64
		voters               modelVoters.Voters
		votersMoradabad      []*modelVoters.Voter_19
		votersRampur         []*modelVoters.Voter_20
		votersBijnor         []*modelVoters.Voter_21
		votersBangalore      []*modelVoters.Voter
		votersHubli          []*modelVoters.Voter
		num                  int64
		user                 []*modelVoters.Account
		err                  error
	)

	limit, _ := e.GetInt("limit")
	offset, _ := e.GetInt("offset")
	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if limit == 0 {
		limit = -1
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelVoters.Query)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	qsRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	cond := orm.NewCondition()
	condVoterId := orm.NewCondition()
	condAcNumber := orm.NewCondition()
	condPartNumber := orm.NewCondition()
	condSectionNumber := orm.NewCondition()
	condSerialNumberInPart := orm.NewCondition()
	condNameEnglish := orm.NewCondition()
	condNameHindi := orm.NewCondition()
	condRelationNameEnglish := orm.NewCondition()
	condRelationNameHindi := orm.NewCondition()
	condGender := orm.NewCondition()
	condIdCardNumber := orm.NewCondition()
	condAcNameEnglish := orm.NewCondition()
	condAcNameHindi := orm.NewCondition()
	condSectionNameEnglish := orm.NewCondition()
	condSectionNameHindi := orm.NewCondition()
	condReligionEnglish := orm.NewCondition()
	condReligionHindi := orm.NewCondition()
	condAge := orm.NewCondition()
	condVote := orm.NewCondition()
	condEmail := orm.NewCondition()
	condMobileNo := orm.NewCondition()
	condImage := orm.NewCondition()

	// Apply filters for each query string
	// Voter Id
	for _, voterId := range query.VoterID {
		if voterId > 0 {
			condVoterId = condVoterId.Or("Voter_id__exact", voterId)
		}
	}

	// Ac Number
	for _, acNumber := range query.AcNumber {
		if acNumber > 0 {
			condAcNumber = condAcNumber.Or("Ac_number__exact", acNumber)
		}
	}

	// Part Number
	for _, partNumber := range query.PartNumber {
		if partNumber > 0 {
			condPartNumber = condPartNumber.Or("Part_number__exact", partNumber)
		}
	}

	// Section Number
	for _, sectionNumber := range query.SectionNumber {
		if sectionNumber > 0 {
			condSectionNumber = condSectionNumber.Or("Section_number__exact", sectionNumber)
		}
	}

	// Serial Number In Part
	for _, serialNumberInPart := range query.SerialNumberInPart {
		if serialNumberInPart > 0 {
			condSerialNumberInPart = condSerialNumberInPart.Or("Serial_number_in_part__exact", serialNumberInPart)
		}
	}

	// Name English
	for _, nameEnglish := range query.NameEnglish {
		if len(strings.TrimSpace(nameEnglish)) > 0 {
			condNameEnglish = condNameEnglish.Or("Name_english__icontains", nameEnglish)
		}
	}

	// Name Hindi
	for _, nameHindi := range query.NameHindi {
		if len(strings.TrimSpace(nameHindi)) > 0 {
			condNameHindi = condNameHindi.Or("Name_hindi__icontains", nameHindi)
		}
	}

	// Relation Name English
	for _, relationNameEnglish := range query.RelationNameEnglish {
		if len(strings.TrimSpace(relationNameEnglish)) > 0 {
			condRelationNameEnglish = condRelationNameEnglish.Or("Relation_name_english__icontains", relationNameEnglish)
		}
	}

	// Relation Name Hindi
	for _, relationNameHindi := range query.RelationNameHindi {
		if len(strings.TrimSpace(relationNameHindi)) > 0 {
			condRelationNameHindi = condRelationNameHindi.Or("Relation_name_hindi__icontains", relationNameHindi)
		}
	}

	// Gender
	for _, gender := range query.Gender {
		if gender == "M" || gender == "F" {
			condGender = condGender.Or("Gender__exact", gender)
		}
	}

	// ID Card Number
	for _, idCardNumber := range query.NameHindi {
		if len(strings.TrimSpace(idCardNumber)) > 0 {
			condIdCardNumber = condIdCardNumber.Or("Id_card_number__exact", idCardNumber)
		}
	}

	// Ac Name English
	for _, acNameEnglish := range query.AcNameEnglish {
		if len(strings.TrimSpace(acNameEnglish)) > 0 {
			condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__exact", acNameEnglish)
		}
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__exact", acNameHindi)
		}
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__exact", sectionNameEnglish)
		}
	}

	// Section Name Hindi
	for _, sectionNameHindi := range query.SectionNameHindi {
		if len(strings.TrimSpace(sectionNameHindi)) > 0 {
			condSectionNameHindi = condSectionNameHindi.Or("Section_name_hindi__ilike", sectionNameHindi)
		}
	}

	//Religion English
	for _, religionEnglish := range query.ReligionEnglish {
		if len(strings.TrimSpace(religionEnglish)) > 0 {
			condReligionEnglish = condReligionEnglish.Or("Religion_english__exact", religionEnglish)
		}
	}

	// Religion Hindi
	for _, religionHindi := range query.ReligionHindi {
		if len(strings.TrimSpace(religionHindi)) > 0 {
			condReligionHindi = condReligionHindi.Or("Religion_hindi__exact", religionHindi)
		}
	}

	// Age
	for _, age := range query.Age {
		if age > 0 {
			condAge = condAge.Or("Age__exact", age)
		}
	}

	// Vote
	for _, vote := range query.Vote {
		if vote == 0 || vote == 1 {
			condVote = condVote.Or("Vote__exact", vote)
		}
	}

	// Email
	for _, email := range query.Email {
		if len(strings.TrimSpace(email)) > 0 {
			condEmail = condEmail.Or("Email__exact", email)
		}
	}

	// Mobile Number
	for _, mobileNo := range query.MobileNo {
		if mobileNo > 0 {
			condMobileNo = condMobileNo.Or("MobileNo__exact", mobileNo)
		}
	}

	// Image
	for _, image := range query.Image {
		if len(strings.TrimSpace(image)) > 0 {
			condImage = condImage.Or("Image__exact", image)
		}
	}

	if condVoterId != nil && !condVoterId.IsEmpty() {
		cond = condVoterId
	}

	if condAcNumber != nil && !condAcNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNumber)
		} else {
			cond = condAcNumber
		}
	}

	if condPartNumber != nil && !condPartNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condPartNumber)
		} else {
			cond = condPartNumber
		}
	}

	if condSectionNumber != nil && !condSectionNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNumber)
		} else {
			cond = condSectionNumber
		}
	}

	if condSerialNumberInPart != nil && !condSerialNumberInPart.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSerialNumberInPart)
		} else {
			cond = condSerialNumberInPart
		}
	}

	if condNameEnglish != nil && !condNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condNameEnglish)
		} else {
			cond = condNameEnglish
		}
	}

	if condNameHindi != nil && !condNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condNameHindi)
		} else {
			cond = condNameHindi
		}
	}

	if condRelationNameEnglish != nil && !condRelationNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condRelationNameEnglish)
		} else {
			cond = condRelationNameEnglish
		}
	}

	if condRelationNameHindi != nil && !condRelationNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condRelationNameHindi)
		} else {
			cond = condRelationNameHindi
		}
	}

	if condGender != nil && !condGender.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condGender)
		} else {
			cond = condGender
		}
	}

	if condIdCardNumber != nil && !condIdCardNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condIdCardNumber)
		} else {
			cond = condIdCardNumber
		}
	}

	if condAcNameEnglish != nil && !condAcNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNameEnglish)
		} else {
			cond = condAcNameEnglish
		}
	}

	if condAcNameHindi != nil && !condAcNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNameHindi)
		} else {
			cond = condAcNameHindi
		}
	}

	if condSectionNameEnglish != nil && !condSectionNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNameEnglish)
		} else {
			cond = condSectionNameEnglish
		}
	}

	if condSectionNameHindi != nil && !condSectionNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNameHindi)
		} else {
			cond = condSectionNameHindi
		}
	}

	if condReligionEnglish != nil && !condReligionEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condReligionEnglish)
		} else {
			cond = condReligionEnglish
		}
	}

	if condReligionHindi != nil && !condReligionHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condReligionHindi)
		} else {
			cond = condReligionHindi
		}
	}

	if condAge != nil && !condAge.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAge)
		} else {
			cond = condAge
		}
	}

	if condVote != nil && !condVote.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condVote)
		} else {
			cond = condVote
		}
	}

	if condEmail != nil && !condEmail.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condEmail)
		} else {
			cond = condEmail
		}
	}

	if condMobileNo != nil && !condMobileNo.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condMobileNo)
		} else {
			cond = condMobileNo
		}
	}

	if condImage != nil && !condImage.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condImage)
		} else {
			cond = condImage
		}
	}

	qsRampur = qsRampur.SetCond(cond)
	qsMoradabad = qsMoradabad.SetCond(cond)
	qsBijnor = qsBijnor.SetCond(cond)

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
			if state == 27 {
				votersCountRampur, _ = qsRampur.Count()
				_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
				if err != nil {
					responseStatus := modelVoters.NewResponseStatus()
					responseStatus.Response = "error"
					responseStatus.Message = fmt.Sprintf("Db Error Rampur. Unable to get the voters.")
					responseStatus.Error = err.Error()
					e.Data["json"] = &responseStatus
					e.ServeJSON()
				}
				votersCountMoradabad, _ = qsMoradabad.Count()
				_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
				if err != nil {
					responseStatus := modelVoters.NewResponseStatus()
					responseStatus.Response = "error"
					responseStatus.Message = fmt.Sprintf("Db Error Moradabad. Unable to get the voters.")
					responseStatus.Error = err.Error()
					e.Data["json"] = &responseStatus
					e.ServeJSON()
				}
				votersCountBijnor, _ = qsBijnor.Count()
				_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
				if err != nil {
					responseStatus := modelVoters.NewResponseStatus()
					responseStatus.Response = "error"
					responseStatus.Message = fmt.Sprintf("Db Error Bijnor. Unable to get the voters.")
					responseStatus.Error = err.Error()
					e.Data["json"] = &responseStatus
					e.ServeJSON()
				}
			}

			if state == 12 {
				votersCountBangalore, _ = qsBangalore.Count()
				_, err = qsBangalore.Limit(limit, offset).All(&votersBangalore)
				if err != nil {
					responseStatus := modelVoters.NewResponseStatus()
					responseStatus.Response = "error"
					responseStatus.Message = fmt.Sprintf("Db Error Bangalore. Unable to get the voters.")
					responseStatus.Error = err.Error()
					e.Data["json"] = &responseStatus
					e.ServeJSON()
				}
				votersCountHubli, _ = qsHubli.Count()
				_, err = qsHubli.Limit(limit, offset).All(&votersHubli)
				if err != nil {
					responseStatus := modelVoters.NewResponseStatus()
					responseStatus.Response = "error"
					responseStatus.Message = fmt.Sprintf("Db Error Hubli. Unable to get the voters.")
					responseStatus.Error = err.Error()
					e.Data["json"] = &responseStatus
					e.ServeJSON()
				}
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Moradabad. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Rampur. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Bijnor. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Moradabad. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Rampur. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Bijnor. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Moradabad. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if district == 20 {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Rampur. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if district == 21 {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Db Error Bijnor. Unable to get the voters.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	voters.Populate_19(votersMoradabad)
	voters.Populate_20(votersRampur)
	voters.Populate_21(votersBijnor)

	totalVotersCount := votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli + votersCountBijnor

	if votersCountRampur > 0 || votersCountMoradabad > 0 || votersCountBijnor > 0 || votersCountBangalore > 0 || votersCountHubli > 0 {
		voters.Total = totalVotersCount
		e.Data["json"] = voters
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No voters found with this criteria."
		if err != nil {
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.ServeJSON()
}

func (e *ElectionController) GetStatistic() {
	var (
		votersCount                      int64
		votersCountRampur                int64
		votersCountMoradabad             int64
		votersCountBijnor                int64
		votersCountBangalore             int64
		votersCountHubli                 int64
		muslimVotersCount                int64
		OthersVotersCount                int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		OthersMaleVotersCount            int64
		OthersFemaleVotersCount          int64
		muslimVotersP                    float64
		OthersVotersP                    float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		OthersMaleVotersP                float64
		OthersFemaleVotersP              float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		OthersVotersCountRampur          int64
		OthersVotersCountMoradabad       int64
		OthersVotersCountBijnor          int64
		OthersVotersCountBangalore       int64
		OthersVotersCountHubli           int64
		maleVotersCountRampur            int64
		maleVotersCountMoradabad         int64
		maleVotersCountBijnor            int64
		maleVotersCountBangalore         int64
		maleVotersCountHubli             int64
		femaleVotersCountRampur          int64
		femaleVotersCountMoradabad       int64
		femaleVotersCountBijnor          int64
		femaleVotersCountBangalore       int64
		femaleVotersCountHubli           int64
		muslimMaleVotersCountRampur      int64
		muslimMaleVotersCountMoradabad   int64
		muslimMaleVotersCountBijnor      int64
		muslimMaleVotersCountBangalore   int64
		muslimMaleVotersCountHubli       int64
		muslimFemaleVotersCountRampur    int64
		muslimFemaleVotersCountMoradabad int64
		muslimFemaleVotersCountBijnor    int64
		muslimFemaleVotersCountBangalore int64
		muslimFemaleVotersCountHubli     int64
		OthersMaleVotersCountRampur      int64
		OthersMaleVotersCountMoradabad   int64
		OthersMaleVotersCountBijnor      int64
		OthersMaleVotersCountBangalore   int64
		OthersMaleVotersCountHubli       int64
		OthersFemaleVotersCountRampur    int64
		OthersFemaleVotersCountMoradabad int64
		OthersFemaleVotersCountBijnor    int64
		OthersFemaleVotersCountBangalore int64
		OthersFemaleVotersCountHubli     int64
		num                              int64
		user                             []*modelVoters.Account
		err                              error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelVoters.Query)
	statistic := new(modelVoters.Statistic)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = "Invalid Json. Unable to parse."
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	qsRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	cond := orm.NewCondition()
	condAcNumber := orm.NewCondition()
	condPartNumber := orm.NewCondition()
	condSectionNumber := orm.NewCondition()
	condAcNameEnglish := orm.NewCondition()
	condAcNameHindi := orm.NewCondition()
	condSectionNameEnglish := orm.NewCondition()
	condSectionNameHindi := orm.NewCondition()
	condAge := orm.NewCondition()
	condVote := orm.NewCondition()

	// Apply filters for each query string
	// Ac Number
	for _, acNumber := range query.AcNumber {
		if acNumber > 0 {
			condAcNumber = condAcNumber.Or("Ac_number__exact", acNumber)
		}
	}

	// Part Number
	for _, partNumber := range query.PartNumber {
		if partNumber > 0 {
			condPartNumber = condPartNumber.Or("Part_number__exact", partNumber)
		}
	}

	// Section Number
	for _, sectionNumber := range query.SectionNumber {
		if sectionNumber > 0 {
			condSectionNumber = condSectionNumber.Or("Section_number__exact", sectionNumber)
		}
	}

	// Ac Name English
	for _, acNameEnglish := range query.AcNameEnglish {
		if len(strings.TrimSpace(acNameEnglish)) > 0 {
			condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__exact", acNameEnglish)
		}
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__exact", acNameHindi)
		}
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__exact", sectionNameEnglish)
		}
	}

	// Section Name Hindi
	for _, sectionNameHindi := range query.SectionNameHindi {
		if len(strings.TrimSpace(sectionNameHindi)) > 0 {
			condSectionNameHindi = condSectionNameHindi.Or("Section_name_hindi__ilike", sectionNameHindi)
		}
	}

	// Age
	for _, age := range query.Age {
		if age > 0 {
			condAge = condAge.Or("Age__exact", age)
		}
	}

	// Vote
	for _, vote := range query.Vote {
		if vote == 0 || vote == 1 {
			condVote = condVote.Or("Vote__exact", vote)
		}
	}

	if condAcNumber != nil && !condAcNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNumber)
		} else {
			cond = condAcNumber
		}
	}

	if condPartNumber != nil && !condPartNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condPartNumber)
		} else {
			cond = condPartNumber
		}
	}

	if condSectionNumber != nil && !condSectionNumber.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNumber)
		} else {
			cond = condSectionNumber
		}
	}

	if condAcNameEnglish != nil && !condAcNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNameEnglish)
		} else {
			cond = condAcNameEnglish
		}
	}

	if condAcNameHindi != nil && !condAcNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAcNameHindi)
		} else {
			cond = condAcNameHindi
		}
	}

	if condSectionNameEnglish != nil && !condSectionNameEnglish.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNameEnglish)
		} else {
			cond = condSectionNameEnglish
		}
	}

	if condSectionNameHindi != nil && !condSectionNameHindi.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condSectionNameHindi)
		} else {
			cond = condSectionNameHindi
		}
	}

	if condAge != nil && !condAge.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAge)
		} else {
			cond = condAge
		}
	}

	if condVote != nil && !condVote.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condVote)
		} else {
			cond = condVote
		}
	}

	qsRampur = qsRampur.SetCond(cond)
	qsMoradabad = qsMoradabad.SetCond(cond)
	qsBijnor = qsBijnor.SetCond(cond)

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsRampur.Count()
				votersCountMoradabad, err = qsMoradabad.Count()
				votersCountBijnor, err = qsBijnor.Count()
				muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCount > 0 {
		statistic.Total = votersCount

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistic.Muslim = muslimVotersCount

		OthersVotersCount = OthersVotersCountRampur + OthersVotersCountMoradabad + OthersVotersCountBijnor + OthersVotersCountBangalore + OthersVotersCountHubli
		statistic.Others = OthersVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistic.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistic.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistic.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistic.MuslimFemale = muslimFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistic.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistic.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCount)) * 100
		statistic.MuslimP = muslimVotersP

		OthersVotersP = (float64(OthersVotersCount) / float64(votersCount)) * 100
		statistic.OthersP = OthersVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCount)) * 100
		statistic.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCount)) * 100
		statistic.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCount)) * 100
		statistic.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCount)) * 100
		statistic.MuslimFemaleP = muslimFemaleVotersP

		OthersMaleVotersP = (float64(OthersMaleVotersCount) / float64(votersCount)) * 100
		statistic.OthersMaleP = OthersMaleVotersP

		OthersFemaleVotersP = (float64(OthersFemaleVotersCount) / float64(votersCount)) * 100
		statistic.OthersFemaleP = OthersFemaleVotersP

		e.Data["json"] = statistic
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No statistic found with this criteria. Please try again!"
		if err != nil {
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	e.ServeJSON()
}

func (e *ElectionController) GetStatistics() {
	var (
		votersCount                      int64
		votersResult                     int64
		votersCountQuery                 int64
		votersCountScope                 int64
		votersCountRampur                int64
		votersCountMoradabad             int64
		votersCountBijnor                int64
		votersCountBangalore             int64
		votersCountHubli                 int64
		muslimVotersCount                int64
		OthersVotersCount                int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		OthersMaleVotersCount            int64
		OthersFemaleVotersCount          int64
		muslimVotersP                    float64
		OthersVotersP                    float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		OthersMaleVotersP                float64
		OthersFemaleVotersP              float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		OthersVotersCountRampur          int64
		OthersVotersCountMoradabad       int64
		OthersVotersCountBijnor          int64
		OthersVotersCountBangalore       int64
		OthersVotersCountHubli           int64
		maleVotersCountRampur            int64
		maleVotersCountMoradabad         int64
		maleVotersCountBijnor            int64
		maleVotersCountBangalore         int64
		maleVotersCountHubli             int64
		femaleVotersCountRampur          int64
		femaleVotersCountMoradabad       int64
		femaleVotersCountBijnor          int64
		femaleVotersCountBangalore       int64
		femaleVotersCountHubli           int64
		muslimMaleVotersCountRampur      int64
		muslimMaleVotersCountMoradabad   int64
		muslimMaleVotersCountBijnor      int64
		muslimMaleVotersCountBangalore   int64
		muslimMaleVotersCountHubli       int64
		muslimFemaleVotersCountRampur    int64
		muslimFemaleVotersCountMoradabad int64
		muslimFemaleVotersCountBijnor    int64
		muslimFemaleVotersCountBangalore int64
		muslimFemaleVotersCountHubli     int64
		OthersMaleVotersCountRampur      int64
		OthersMaleVotersCountMoradabad   int64
		OthersMaleVotersCountBijnor      int64
		OthersMaleVotersCountBangalore   int64
		OthersMaleVotersCountHubli       int64
		OthersFemaleVotersCountRampur    int64
		OthersFemaleVotersCountMoradabad int64
		OthersFemaleVotersCountBijnor    int64
		OthersFemaleVotersCountBangalore int64
		OthersFemaleVotersCountHubli     int64
		num                              int64
		user                             []*modelVoters.Account
		err                              error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	queries := new(modelVoters.Queries)
	statistics := new(modelVoters.Statistics)

	err = json.Unmarshal(inputJson, &queries)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = "Invalid Json. Unable to parse."
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	// Query
	qsRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	// Scope
	qsScopeRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsScopeMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsScopeBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	qsScopeBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	qsScopeHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	condQuery := orm.NewCondition()
	condScope := orm.NewCondition()
	condAcNumberQuery := orm.NewCondition()
	condAcNumberScope := orm.NewCondition()
	condPartNumberQuery := orm.NewCondition()
	condPartNumberScope := orm.NewCondition()
	condSectionNumberQuery := orm.NewCondition()
	condSectionNumberScope := orm.NewCondition()
	condAcNameEnglishQuery := orm.NewCondition()
	condAcNameEnglishScope := orm.NewCondition()
	condAcNameHindiQuery := orm.NewCondition()
	condAcNameHindiScope := orm.NewCondition()
	condSectionNameEnglishQuery := orm.NewCondition()
	condSectionNameEnglishScope := orm.NewCondition()
	condSectionNameHindiQuery := orm.NewCondition()
	condSectionNameHindiScope := orm.NewCondition()
	condAgeQuery := orm.NewCondition()
	condAgeScope := orm.NewCondition()
	condVoteQuery := orm.NewCondition()
	condVoteScope := orm.NewCondition()

	// Apply filters for each query string
	// Ac Number Query
	for _, acNumber := range queries.Query.AcNumber {
		if acNumber > 0 {
			condAcNumberQuery = condAcNumberQuery.Or("Ac_number__exact", acNumber)
		}
	}

	// Apply filters for each query string
	// Ac Number Scope
	for _, acNumber := range queries.Scope.AcNumber {
		if acNumber > 0 {
			condAcNumberScope = condAcNumberScope.Or("Ac_number__exact", acNumber)
		}
	}

	// Part Number Query
	for _, partNumber := range queries.Query.PartNumber {
		if partNumber > 0 {
			condPartNumberQuery = condPartNumberQuery.Or("Part_number__exact", partNumber)
		}
	}

	// Part Number Scope
	for _, partNumber := range queries.Scope.PartNumber {
		if partNumber > 0 {
			condPartNumberScope = condPartNumberScope.Or("Part_number__exact", partNumber)
		}
	}

	// Section Number Query
	for _, sectionNumber := range queries.Query.SectionNumber {
		if sectionNumber > 0 {
			condSectionNumberQuery = condSectionNumberQuery.Or("Section_number__exact", sectionNumber)
		}
	}

	// Section Number Scope
	for _, sectionNumber := range queries.Scope.SectionNumber {
		if sectionNumber > 0 {
			condSectionNumberScope = condSectionNumberScope.Or("Section_number__exact", sectionNumber)
		}
	}

	// Ac Name English Query
	for _, acNameEnglish := range queries.Query.AcNameEnglish {
		if len(strings.TrimSpace(acNameEnglish)) > 0 {
			condAcNameEnglishQuery = condAcNameEnglishQuery.Or("Ac_name_english__exact", acNameEnglish)
		}
	}

	// Ac Name English Scope
	for _, acNameEnglish := range queries.Scope.AcNameEnglish {
		if len(strings.TrimSpace(acNameEnglish)) > 0 {
			condAcNameEnglishScope = condAcNameEnglishScope.Or("Ac_name_english__exact", acNameEnglish)
		}
	}

	// Ac Name Hindi Query
	for _, acNameHindi := range queries.Query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindiQuery = condAcNameHindiQuery.Or("Ac_name_hindi__exact", acNameHindi)
		}
	}

	// Ac Name Hindi Scope
	for _, acNameHindi := range queries.Scope.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindiScope = condAcNameHindiScope.Or("Ac_name_hindi__exact", acNameHindi)
		}
	}

	// Section Name English Query
	for _, sectionNameEnglish := range queries.Query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglishQuery = condSectionNameEnglishQuery.Or("Section_name_english__exact", sectionNameEnglish)
		}
	}

	// Section Name English Scope
	for _, sectionNameEnglish := range queries.Scope.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglishScope = condSectionNameEnglishScope.Or("Section_name_english", sectionNameEnglish)
		}
	}

	// Section Name Hindi Query
	for _, sectionNameHindi := range queries.Query.SectionNameHindi {
		if len(strings.TrimSpace(sectionNameHindi)) > 0 {
			condSectionNameHindiQuery = condSectionNameHindiQuery.Or("Section_name_hindi__ilike", sectionNameHindi)
		}
	}

	// Section Name Hindi Scope
	for _, sectionNameHindi := range queries.Scope.SectionNameHindi {
		if len(strings.TrimSpace(sectionNameHindi)) > 0 {
			condSectionNameHindiScope = condSectionNameHindiScope.Or("Section_name_hindi__ilike", sectionNameHindi)
		}
	}

	// Age Query
	for _, age := range queries.Query.Age {
		if age > 0 {
			condAgeQuery = condAgeQuery.Or("Age__exact", age)
		}
	}

	// Age Scope
	for _, age := range queries.Scope.Age {
		if age > 0 {
			condAgeScope = condAgeScope.Or("Age__exact", age)
		}
	}

	// Vote Query
	for _, vote := range queries.Query.Vote {
		if vote > 0 {
			condVoteQuery = condVoteQuery.Or("Vote__exact", vote)
		}
	}

	// Vote Scope
	for _, vote := range queries.Scope.Vote {
		if vote > 0 {
			condVoteScope = condVoteScope.Or("Vote__exact", vote)
		}
	}

	if condAcNumberQuery != nil && !condAcNumberQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condAcNumberQuery)
		} else {
			condQuery = condAcNumberQuery
		}
	}

	if condAcNumberScope != nil && !condAcNumberScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condAcNumberScope)
		} else {
			condScope = condAcNumberScope
		}
	}

	if condPartNumberQuery != nil && !condPartNumberQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condPartNumberQuery)
		} else {
			condQuery = condPartNumberQuery
		}
	}

	if condPartNumberScope != nil && !condPartNumberScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condPartNumberScope)
		} else {
			condScope = condPartNumberScope
		}
	}

	if condSectionNumberQuery != nil && !condSectionNumberQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condSectionNumberQuery)
		} else {
			condQuery = condSectionNumberQuery
		}
	}

	if condSectionNumberScope != nil && !condSectionNumberScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condSectionNumberScope)
		} else {
			condScope = condSectionNumberScope
		}
	}

	if condAcNameEnglishQuery != nil && !condAcNameEnglishQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condAcNameEnglishQuery)
		} else {
			condQuery = condAcNameEnglishQuery
		}
	}

	if condAcNameEnglishScope != nil && !condAcNameEnglishScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condAcNameEnglishScope)
		} else {
			condScope = condAcNameEnglishScope
		}
	}

	if condAcNameHindiQuery != nil && !condAcNameHindiQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condAcNameHindiQuery)
		} else {
			condQuery = condAcNameHindiQuery
		}
	}

	if condAcNameHindiScope != nil && !condAcNameHindiScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condAcNameHindiScope)
		} else {
			condScope = condAcNameHindiScope
		}
	}

	if condSectionNameEnglishQuery != nil && !condSectionNameEnglishQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condSectionNameEnglishQuery)
		} else {
			condQuery = condSectionNameEnglishQuery
		}
	}

	if condSectionNameEnglishScope != nil && !condSectionNameEnglishScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condSectionNameEnglishScope)
		} else {
			condScope = condSectionNameEnglishScope
		}
	}

	if condSectionNameHindiQuery != nil && !condSectionNameHindiQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condSectionNameHindiQuery)
		} else {
			condQuery = condSectionNameHindiQuery
		}
	}

	if condSectionNameHindiScope != nil && !condSectionNameHindiScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condSectionNameHindiScope)
		} else {
			condScope = condSectionNameHindiScope
		}
	}

	if condAgeQuery != nil && !condAgeQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condAgeQuery)
		} else {
			condQuery = condAgeQuery
		}
	}

	if condAgeScope != nil && !condAgeScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condAgeScope)
		} else {
			condScope = condAgeScope
		}
	}

	if condVoteQuery != nil && !condVoteQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condVoteQuery)
		} else {
			condQuery = condVoteQuery
		}
	}

	if condVoteScope != nil && !condVoteScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condVoteScope)
		} else {
			condScope = condVoteScope
		}
	}

	qsRampur = qsRampur.SetCond(condQuery)
	qsMoradabad = qsMoradabad.SetCond(condQuery)
	qsBijnor = qsBijnor.SetCond(condQuery)
	qsScopeRampur = qsScopeRampur.SetCond(condScope)
	qsScopeMoradabad = qsScopeMoradabad.SetCond(condScope)

	// CALCULATIONS FOR QUERY
	// Get voters for each state
	if len(queries.Query.DistrictNameEnglish) == 0 && len(queries.Query.DistrictNameHindi) == 0 && len(queries.Query.DistrictNumber) == 0 {
		for _, state := range queries.Query.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsRampur.Count()
				votersCountMoradabad, err = qsMoradabad.Count()
				muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				votersCountBijnor, err = qsBijnor.Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCountQuery = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountQuery > 0 {
		statistics.Query.Total = votersCountQuery

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Query.Muslim = muslimVotersCount

		OthersVotersCount = OthersVotersCountRampur + OthersVotersCountMoradabad + OthersVotersCountBijnor + OthersVotersCountBangalore + OthersVotersCountHubli
		statistics.Query.Others = OthersVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Query.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Query.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Query.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Query.MuslimFemale = muslimFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistics.Query.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistics.Query.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimP = muslimVotersP

		OthersVotersP = (float64(OthersVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OthersP = OthersVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimFemaleP = muslimFemaleVotersP

		OthersMaleVotersP = (float64(OthersMaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OthersMaleP = OthersMaleVotersP

		OthersFemaleVotersP = (float64(OthersFemaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OthersFemaleP = OthersFemaleVotersP
	}

	//Reset all counters
	votersCountRampur = 0
	votersCountMoradabad = 0
	muslimVotersCountRampur = 0
	muslimVotersCountMoradabad = 0
	OthersVotersCountRampur = 0
	OthersVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	muslimVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	OthersVotersCountBangalore = 0
	OthersVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	OthersMaleVotersCountBangalore = 0
	OthersMaleVotersCountHubli = 0
	OthersFemaleVotersCountBangalore = 0
	OthersFemaleVotersCountHubli = 0

	// CALCULATIONS FOR SCOPE
	// Get voters for each state
	if len(queries.Scope.DistrictNameEnglish) == 0 && len(queries.Scope.DistrictNameHindi) == 0 && len(queries.Scope.DistrictNumber) == 0 {
		for _, state := range queries.Scope.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsScopeRampur.Count()
				votersCountMoradabad, err = qsScopeMoradabad.Count()
				muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				votersCountBijnor, err = qsScopeBijnor.Count()
				muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsScopeBangalore.Count()
				votersCountHubli, err = qsScopeHubli.Count()
				muslimVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Muslim").Count()
				OthersVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				OthersMaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Scope.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Scope.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCountScope = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountScope > 0 {
		statistics.Scope.Total = votersCountScope

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Scope.Muslim = muslimVotersCount

		OthersVotersCount = OthersVotersCountRampur + OthersVotersCountMoradabad + OthersVotersCountBijnor + OthersVotersCountBangalore + OthersVotersCountHubli
		statistics.Scope.Others = OthersVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Scope.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Scope.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Scope.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Scope.MuslimFemale = muslimFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistics.Scope.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistics.Scope.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimP = muslimVotersP

		OthersVotersP = (float64(OthersVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OthersP = OthersVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimFemaleP = muslimFemaleVotersP

		OthersMaleVotersP = (float64(OthersMaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OthersMaleP = OthersMaleVotersP

		OthersFemaleVotersP = (float64(OthersFemaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OthersFemaleP = OthersFemaleVotersP
	}

	condNameEnglishQuery := orm.NewCondition()
	condNameEnglishScope := orm.NewCondition()
	condNameHindiQuery := orm.NewCondition()
	condNameHindiScope := orm.NewCondition()
	condRelationNameEnglishQuery := orm.NewCondition()
	condRelationNameEnglishScope := orm.NewCondition()
	condRelationNameHindiQuery := orm.NewCondition()
	condRelationNameHindiScope := orm.NewCondition()
	condGenderQuery := orm.NewCondition()
	condGenderScope := orm.NewCondition()
	condReligionEnglishQuery := orm.NewCondition()
	condReligionEnglishScope := orm.NewCondition()
	condReligionHindiQuery := orm.NewCondition()
	condReligionHindiScope := orm.NewCondition()

	// Apply filters for each query string
	// Name English Query
	for _, nameEnglish := range queries.Query.NameEnglish {
		if len(strings.TrimSpace(nameEnglish)) > 0 {
			condNameEnglishQuery = condNameEnglishQuery.Or("Name_english__icontains", nameEnglish)
		}
	}

	// Name English Scope
	for _, nameEnglish := range queries.Scope.NameEnglish {
		if len(strings.TrimSpace(nameEnglish)) > 0 {
			condNameEnglishScope = condNameEnglishScope.Or("Name_english__icontains", nameEnglish)
		}
	}

	// Apply filters for each query string
	// Name Hindi Query
	for _, nameHindi := range queries.Query.NameHindi {
		if len(strings.TrimSpace(nameHindi)) > 0 {
			condNameHindiQuery = condNameHindiQuery.Or("Name_hindi__icontains", nameHindi)
		}
	}

	// Name Hindi Scope
	for _, nameHindi := range queries.Scope.NameHindi {
		if len(strings.TrimSpace(nameHindi)) > 0 {
			condNameHindiScope = condNameHindiScope.Or("Name_hindi__icontains", nameHindi)
		}
	}

	// Relation Name English Query
	for _, relationNameEnglish := range queries.Query.RelationNameEnglish {
		if len(strings.TrimSpace(relationNameEnglish)) > 0 {
			condRelationNameEnglishQuery = condRelationNameEnglishQuery.Or("Relation_name_english__icontains", relationNameEnglish)
		}
	}

	// Relation Name English Scope
	for _, relationNameEnglish := range queries.Scope.RelationNameEnglish {
		if len(strings.TrimSpace(relationNameEnglish)) > 0 {
			condRelationNameEnglishScope = condRelationNameEnglishScope.Or("Relation_name_english__icontains", relationNameEnglish)
		}
	}

	// Relation Name Hindi Query
	for _, relationNameHindi := range queries.Query.RelationNameHindi {
		if len(strings.TrimSpace(relationNameHindi)) > 0 {
			condRelationNameHindiQuery = condRelationNameHindiQuery.Or("Relation_name_hindi__icontains", relationNameHindi)
		}
	}

	// Relation Name Hindi Scope
	for _, relationNameHindi := range queries.Scope.RelationNameHindi {
		if len(strings.TrimSpace(relationNameHindi)) > 0 {
			condRelationNameHindiScope = condRelationNameHindiScope.Or("Relation_name_hindi__icontains", relationNameHindi)
		}
	}

	// Gender Query
	for _, gender := range queries.Query.Gender {
		if gender == "M" || gender == "F" {
			condGenderQuery = condGenderQuery.Or("Gender__exact", gender)
		}
	}

	// Gender Scope
	for _, gender := range queries.Scope.Gender {
		if gender == "M" || gender == "F" {
			condGenderScope = condGenderScope.Or("Gender__exact", gender)
		}
	}

	//Religion English Query
	for _, religionEnglish := range queries.Query.ReligionEnglish {
		if len(strings.TrimSpace(religionEnglish)) > 0 {
			condReligionEnglishQuery = condReligionEnglishQuery.Or("Religion_english__exact", religionEnglish)
		}
	}

	//Religion English Scope
	for _, religionEnglish := range queries.Scope.ReligionEnglish {
		if len(strings.TrimSpace(religionEnglish)) > 0 {
			condReligionEnglishScope = condReligionEnglishScope.Or("Religion_english__exact", religionEnglish)
		}
	}

	// Religion Hindi Query
	for _, religionHindi := range queries.Query.ReligionHindi {
		if len(strings.TrimSpace(religionHindi)) > 0 {
			condReligionHindiQuery = condReligionHindiQuery.Or("Religion_hindi__exact", religionHindi)
		}
	}

	// Religion Hindi Scope
	for _, religionHindi := range queries.Scope.ReligionHindi {
		if len(strings.TrimSpace(religionHindi)) > 0 {
			condReligionHindiScope = condReligionHindiScope.Or("Religion_hindi__exact", religionHindi)
		}
	}

	if condNameEnglishQuery != nil && !condNameEnglishQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condNameEnglishQuery)
		} else {
			condQuery = condNameEnglishQuery
		}
	}

	if condNameEnglishScope != nil && !condNameEnglishScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condNameEnglishScope)
		} else {
			condScope = condNameEnglishScope
		}
	}

	if condNameHindiQuery != nil && !condNameHindiQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condNameHindiQuery)
		} else {
			condQuery = condNameHindiQuery
		}
	}

	if condNameHindiScope != nil && !condNameHindiScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condNameHindiScope)
		} else {
			condScope = condNameHindiScope
		}
	}

	if condNameEnglishQuery != nil && !condNameEnglishQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condNameEnglishQuery)
		} else {
			condQuery = condNameEnglishQuery
		}
	}

	if condRelationNameEnglishScope != nil && !condRelationNameEnglishScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condRelationNameEnglishScope)
		} else {
			condScope = condRelationNameEnglishScope
		}
	}

	if condRelationNameHindiQuery != nil && !condRelationNameHindiQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condRelationNameHindiQuery)
		} else {
			condQuery = condRelationNameHindiQuery
		}
	}

	if condRelationNameHindiScope != nil && !condRelationNameHindiScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condRelationNameHindiScope)
		} else {
			condScope = condRelationNameHindiScope
		}
	}

	if condGenderQuery != nil && !condGenderQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condGenderQuery)
		} else {
			condQuery = condGenderQuery
		}
	}

	if condGenderScope != nil && !condGenderScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condGenderScope)
		} else {
			condScope = condGenderScope
		}
	}

	if condReligionEnglishQuery != nil && !condReligionEnglishQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condReligionEnglishQuery)
		} else {
			condQuery = condReligionEnglishQuery
		}
	}

	if condReligionEnglishScope != nil && !condReligionEnglishScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condReligionEnglishScope)
		} else {
			condScope = condReligionEnglishScope
		}
	}

	if condReligionHindiQuery != nil && !condReligionHindiQuery.IsEmpty() {
		if condQuery != nil && !condQuery.IsEmpty() {
			condQuery = condQuery.AndCond(condReligionHindiQuery)
		} else {
			condQuery = condReligionHindiQuery
		}
	}

	if condReligionHindiScope != nil && !condReligionHindiScope.IsEmpty() {
		if condScope != nil && !condScope.IsEmpty() {
			condScope = condScope.AndCond(condReligionHindiScope)
		} else {
			condScope = condReligionHindiScope
		}
	}

	qsRampur = qsRampur.SetCond(condQuery)
	qsMoradabad = qsMoradabad.SetCond(condQuery)
	qsBijnor = qsBijnor.SetCond(condQuery)
	qsScopeRampur = qsScopeRampur.SetCond(condScope)
	qsScopeMoradabad = qsScopeMoradabad.SetCond(condScope)

	//Reset all counters
	votersCountRampur = 0
	votersCountMoradabad = 0
	muslimVotersCountRampur = 0
	muslimVotersCountMoradabad = 0
	OthersVotersCountRampur = 0
	OthersVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	OthersVotersCountBangalore = 0
	OthersVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	OthersMaleVotersCountBangalore = 0
	OthersMaleVotersCountHubli = 0
	OthersFemaleVotersCountBangalore = 0
	OthersFemaleVotersCountHubli = 0

	// CALCULATIONS FOR QUERY
	// Get voters Count for each state
	if len(queries.Query.DistrictNameEnglish) == 0 && len(queries.Query.DistrictNameHindi) == 0 && len(queries.Query.DistrictNumber) == 0 {
		for _, state := range queries.Query.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsRampur.Count()
				votersCountMoradabad, err = qsMoradabad.Count()
				votersCountBijnor, err = qsBijnor.Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountRampur, err = qsRampur.Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
		}
	}

	votersResult = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli
	statistics.Result = votersResult

	//Reset all counters
	votersCountRampur = 0
	votersCountMoradabad = 0
	muslimVotersCountRampur = 0
	muslimVotersCountMoradabad = 0
	OthersVotersCountRampur = 0
	OthersVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	OthersVotersCountBangalore = 0
	OthersVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	OthersMaleVotersCountBangalore = 0
	OthersMaleVotersCountHubli = 0
	OthersFemaleVotersCountBangalore = 0
	OthersFemaleVotersCountHubli = 0

	// CALCULATIONS FOR SCOPE
	// Get voters count for each state
	if len(queries.Scope.DistrictNameEnglish) == 0 && len(queries.Scope.DistrictNameHindi) == 0 && len(queries.Scope.DistrictNumber) == 0 {
		for _, state := range queries.Scope.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsScopeRampur.Count()
				votersCountMoradabad, err = qsScopeMoradabad.Count()
				votersCountBijnor, err = qsScopeBijnor.Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsScopeBangalore.Count()
				votersCountHubli, err = qsScopeHubli.Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Scope.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsScopeRampur.Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountRampur, err = qsScopeRampur.Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Scope.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsScopeRampur.Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsScopeBijnor.Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
		}

		if district == 20 {
			votersCountRampur, err = qsScopeRampur.Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsScopeBijnor.Count()
		}
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli
	statistics.Total = votersCount

	if votersCount > 0 && votersResult > 0 {
		e.Data["json"] = statistics
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No statistics found with this criteria. Please try again!"
		if err != nil {
			responseStatus.Error = err.Error()
		} else {
			responseStatus.Error = "No Error"
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	e.ServeJSON()
}

func (e *ElectionController) Home() {
	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = "This API is up and running!"
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}

func (e *ElectionController) OTP() {
	var (
		otp       int32
		num       int64
		recipient []*modelVoters.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")
	qsRecipient := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(modelVoters.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	exist := qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid mobile number or mobile number does not exists in our database. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	otp, err = generateOTP()
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't generate the otp. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Update(orm.Params{
		"OTP": otp,
	})

	if err != nil || num != 1 {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Error occur while updating the otp. Please contact electionubda.com team for assistance.")
		if err != nil {
			responseStatus.Error = err.Error()
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsRecipient.Filter("Mobile_no__exact", account.Mobile_no).All(&recipient)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't send the otp. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(recipient) > 0 {
		err = sendOTP(recipient[0].Otp, recipient[0].Email, recipient[0].Display_name, recipient[0].Mobile_no)

		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't send the otp. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = fmt.Sprintf("One time password has been generated and sent to %s successfully.", recipient[0].Email)
		e.Data["json"] = &responseStatus
		e.ServeJSON()

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't send the otp. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
}

func generateOTP() (int32, error) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	secret := gototp.RandomSecret(10, rnd)
	otp, err := gototp.New(secret)
	if err != nil {
		return 0, err
	}
	return otp.Now(), nil
}

func sendOTP(otp int, email string, displayName string, mobileNumber int64) error {
	// Set up authentication information.
	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		"electionubda",
		"hu123*ElectionUBDA",
		smtpServer,
	)

	from := mail.Address{"ElectionUBDA", "electionubda@gmail.com"}
	to := mail.Address{displayName, email}
	toCC1 := mail.Address{"Baligul Hasan", "baligcoup8@gmail.com"}
	toCC2 := mail.Address{"Iftekhar Khan", "iiiftekhar@gmail.com"}
	title := strconv.Itoa(otp) + " is your One Time Password - " + strconv.FormatInt(mobileNumber, 10)

	body := "Hi " + displayName + "!\n\nWelcome to Election UBDA.\n\nThanks & Regards,\nElectionUBDA Team"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address, toCC1.Address, toCC2.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		return err
	}

	return nil
}

func randToken() string {
	b := make([]byte, 8)
	token.Read(b)
	return fmt.Sprintf("%x", b)
}

func (e *ElectionController) Register() {
	var (
		token string
		num   int64
		users []*modelVoters.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(modelVoters.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	exist := qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid mobile number or mobile number does not exists in our database. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).All(&users)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't generate the token. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(users) > 0 {
		if users[0].Otp != account.Otp || users[0].Otp == 0 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Invalid Otp. Please try with a correct otp sent to you or contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		token = randToken()

		num, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Filter("Otp__exact", account.Otp).Update(orm.Params{
			"TOKEN": token,
			"OTP":   0,
		})

		if err != nil || num != 1 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Error occur while updating the token. Please contact electionubda.com team for assistance.")
			if err != nil {
				responseStatus.Error = err.Error()
			}
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		user := new(modelVoters.Account)
		users[0].Token = token
		user = users[0]
		e.Data["json"] = &user
		e.ServeJSON()
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't generate the token. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
}

func (e *ElectionController) GetList() {
	var (
		num  int64
		user []*modelVoters.Account
		err  error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	list := new(modelVoters.List)

	err = json.Unmarshal(inputJson, &list)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(list.Acs) > 0 {
		sections := getApprovedSections(user[0].Approved_acs, list.Acs)
		e.Data["json"] = &sections
		e.ServeJSON()
	}

	if len(list.Districts) > 0 {
		acs := getApprovedAcs(user[0].Approved_districts, list.Districts)
		e.Data["json"] = &acs
		e.ServeJSON()
	}

	approvedDistricts := strings.Split(user[0].Approved_districts, ",")
	e.Data["json"] = &approvedDistricts
	e.ServeJSON()
}

func (e *ElectionController) ReadJson() {
	inputJson := e.Ctx.Input.RequestBody
	readJsons := new(modelVoters.ReadJsons)

	acNameSectionName := make(map[string]string)

	err := json.Unmarshal(inputJson, &readJsons)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	for _, val := range readJsons.ReadJsons {
		acNameSectionName[val.AcName] = val.SectionName
	}
}

func getApprovedSections(csvAcs string, acs []string) []string {
	approvedAcs := strings.Split(csvAcs, ",")
	approvedSections := getSections(approvedAcs)
	sections := getSections(acs)

	return getCommonItems(sections, approvedSections)
}

func getApprovedAcs(csvDistricts string, districts []string) []string {
	approvedDistricts := strings.Split(csvDistricts, ",")
	approvedAcs := getAcs(approvedDistricts)
	acs := getAcs(districts)

	return getCommonItems(acs, approvedAcs)
}

func getSections(acs []string) []string {
	var sections []string

	for _, ac := range acs {
		switch ac {
		case "Najibabad":
			sections = append(sections, "HAKUMATPUR KESHO", "JHAKKAKI", "MOHD ALIPUR DWARKA", "MOHD. ALIPUR HIRDAY AN0", "RAILWAY COLONY RAHUKHERI KORA", "MEERAMPUR BEGA", "MOTA DHAK", "MOJJAMPUR SADAT", "DHANSINI AN0", "PURANPUR GARHI AN0", "MUNEERGANJ AN0", "DAUDPUR HAJI", "RAMPUR CHATHA", "KALHERI", "MANOHARWALA", "MEDUWALA", "PUNDARI KHURD", "KAMRAJPUR AN0", "CHANDGOYALA", "GOSPUR", "KHAIRULLAPUR AN0", "PARWATPUR MAKHDUMPUR AN0", "MAYAPURI", "AURANGPUR BASANTA AN0", "SAIDPUR MIRA", "MOHD. AMI KHANPUR", "AKBARPUR ANWLA", "BIJORI", "RANIPUR", "NARAYANPUR RATAN AN0", "SHEKHPURA ALAM", "SHRAWANPUR", "SIKKAMPUR", "MANDAWALI AN0", "FAZALPUR HABIB AN0", "DUNGERPUR", "HIMPUR PACHATPURA", "MUBARAKPUR RATHE", "MALIPUR", "ZABTAGANJ AN0", "SAIDPURI", "JEETPUR KHAS", "SADAT NAGAR B.A.", "CHAMROLLA", "HAKIMPUR DURG", "MOHD. ALIPUR VEERBHAN A", "HAKIMAN", "AHMADPUR SADAT", "DARODGRAN", "IBRAHIMPUR BAWAN", "MEMRAN", "SHAREEFPUR BANGAR A.", "MIRAMPUR DURG", "HAKIMPUR KAJI", "BANJARAN", "PREM NAGAR", "HASANPUR", "SAIWALA", "NAWABPURA", "HARSHWARA AN0", "MAHAWATPUR BILLOCH", "QURESHIYAN AN0", "KUTUBPUR", "MIRZAPUR SAID", "RAJARAMPUR TULSI", "AMRITSARI", "RAM SAHAYWALI", "MATHURAPUR MOR AN0", "ISSEPUR AN0", "THAPAL", "MOCHIPURA", "IBRAHIMPUR RAJU", "LALPUR SHOJIMAL A.", "MOHD. ALAMPUR", "PASBAN", "MALUKWALI", "SABNIGRAN", "KHERPUR JALIKA", "HUMAYUPUR IDDU", "RAMPURA AN0", "MAHEMSAPUR", "CHAR BAGH", "SHEKH SARAI", "MAHAWATPUR NATHA", "PURANPUR GARHI A", "JAHANABAD", "KAROLI", "MIRZAPUR", "REHMAPUR", "NAWABGANJ", "LALPUR BIKKU", "HARCHANDPUR AN0", "DARIYAPUR", "JWALI LALA AN0", "SARAI ALAM AN0", "MUBARAKPUR AN0", "CHANDOK AN0", "SADULLAPUR", "MOHD. ALIPUR VEERBHAN", "JAMHIRI", "RAJARAMPUR FAZIL", "IBRAHIMPUR BAHAUDDIN", "GADARIYAN", "KAMALPUR", "KALLUGANJ", "BHAGWAN WALA", "RAMDASWALI AN0", "SHAHJAHANPUR", "CHANDPURA", "JABALPUR GUDAR AN0", "GANGUWALA", "VISHANPUR", "KACHHERI SARAI AN0", "NANGLA PITHORA", "KHERA", "KAJIYAN", "FAZALPUR TABELA", "TELINAGALI", "KALHERI A", "FAZALPUR MAN", "KISHANPUR", "DINORA", "NADE HIND", "VIRUWALA", "DAUDPUR NANHERA", "RAGHUNATHPUR", "BASERI", "SEWARAM AN0", "NANGAL AN0", "ASADULLAPUR AN0", "PREAMPUR", "BAIBALWALA", "BHAWAN", "BHAGUWALA AN0", "IBRAHIMPUR JAMALUDDIN", "RASOOLPUR SAID AN0", "RAFIPUR MAJARA", "SHEKHPUR GARHU", "SHYAMIWALA AN0", "RAIPUR SUMALKERI", "FATEHULLAPUR DURG", "JAFARABAD", "KOTAWALI", "MAJIDGANJ", "MALIYAN", "CHAH SALOONO", "MURSHADPUR", "HALDUKHATA", "TAKSAL", "DANIYALPUR", "PURUSHAOTTAMPUR", "BARAMPUR AN0", "SAIFABAD", "PRATHVIPUR", "RAHUKHEDI KAURA", "KARMASKHERI", "AJAMPUR MOHD.ALI URF KHANPUR", "HAREWALI", "RAMDASWALI", "HAJIPUR", "RAMPUR BANWARI A.", "DARBARASHAH", "RAHATPUR KHURD AN0", "HIMAUPUR RAI", "SANIYAN AN-", "SABALPUR AN0", "KESHOPUR", "SHEKHPUR LALA", "MUSVI KHANPUR", "DHARMA GARHI", "NAJIMPUR", "QUILA", "SADATPUR", "RAFIPUR MOHAN", "SHAHZAHAPUR JASRATH", "MUSSEPUR", "BIJARKHATA", "AURANGPUR BHIKKU", "JAFARPUR", "SAJAWALPUR", "MUNGARPUR", "RAJPUR NAVADA", "AURANGABAD AN0", "SANIYAN AN0", "KACHAHARI SARAI AN0", "SAHANPUR SANTOKI", "TAHARPUR ISHAK AN0", "KATRA CHETRAM", "KHANPUR", "GARHMALPUR", "JWALACHANDI", "MOHD PUR ATA", "CHATRUWALA", "MUGLUSHAH AN0", "SADULLA NAGAR", "SHAKARPUR BAHADAR", "RAILWAY COLONY AN0", "MIRZAPUR RANGILA", "MOHD. ALIPUR PARMA", "SIKANDARPUR BASI", "GAJRULA", "MOHD.ALIPUR ALIMUDDIN", "MAHAWATPUR DALPAT", "ALAWALPUR NAINU", "ASLAMPUR JHOJA", "BULCHANDPUR", "SHERAWALA", "LALPUR RAI", "TATARPUR LALU AN0", "GURHA", "GULZAR NAGAR", "SARAI SAID GHORA", "AHMADPUR MAJEED", "KANAKPUR AN0", "REHMAN NAGAR", "AGUPURA PYARA", "ARAFPUR KHAJURI", "BHOGPUR", "GAJIPUR AN0", "GULALWALI", "RAMNAGAR", "RANIKOTA", "OOTWAN", "JASWANTPUR AN0", "KALYANPUR", "ALIPURA AN0", "ANSARIYAN", "SAMIPUR AN0", "RAWAPURI AN0", "PADARATHPUR", "PADHAN", "DAHIRPUR AN0", "NAGLA HARDAS", "FAZALPUR HAIBAT", "AURANGPUR FATTE KHAN", "MOJJAMPUR TULSI", "NAGAL AN0", "MEHNDI BAGH", "RAMPURI", "HASAN ALIPUR", "KUSHHALPUR MARKA", "AKBARPUR CHOGANWA AN0", "MOHD. TAHARPUR", "KUTUBPUR NANGLI", "GOYALA B.A.", "DHANORA", "BISATIYAN", "AJAMPUR GAJI", "SHAHZADPUR AN0", "JATAN", "SHAHARYARPUR BANGER", "JADOPUR", "RAIPUR KHAS AN0", "AURANGPUR BASANTA", "PURANPUR NAROTTAM", "HAKIMPUR DISONDI", "KHALILPUR", "SHAREEFPUR BANGAR B.A.", "SANTOMALAN AN0", "SUNGERPUR", "MANUPURA", "KALALAN", "BANSFODAN AN0", "JALPUR AN0", "BEGAMPUR BHONAWALA", "GULISTAN", "CHOWK BAZAR", "DHARAMPUR BHOJA URF PUNDRI KALA", "SAHANPUR NANU", "SIKRODDA AN0", "HARNATHPUR", "HUSAINPUR MAKHANA", "JAGDEESHPUR", "AMANULLAPUR", "MOTHALA", "TODA", "BALAKRAM", "FAZALPUR FATEHULLA", "TAJPUR", "VEERPUR", "NOORAMPUR", "KAMGARPUR AN0", "MAJIDGANJ AN0", "RAMPUR MANGAL", "SHYAMLI", "SURSENA", "TAYYABPUR GORVA", "BHADOLA", "RAMPUR BANWARI", "HARIJAN BASTI", "BARKHURDARPUR GOPAL", "BHAGUWALA", "FAZALPUR KHAS", "ABDULLAPUR", "SHAHPUR MIRA", "SABALGARH", "PATHANPURA AN0", "MAHAL SARAI AN0", "GOYALA", "JATPURA KHAS", "FARAZPUR", "WAHID NAGAR AN0", "FIELD", "NAYAMATPUR", "BAKARPUR", "MUKTESHWAR MAHADEV", "SALEMPUR", "LAHAK KALA", "MOHD PUR RAVA", "NADDAFAN", "VANSH GOPALPUR", "ZULFUKARPUR GARHI", "TISOTRA AN0", "LALPUR SHOJIMAL", "JAGDISHPUR", "JASPUR", "SHERPUR ABHI", "JARUFSAJAN", "JALABPUR GUDAR AN0", "SADATNAGAER A0", "SANWALDAS", "MUNEERGANJ", "MUSTAFAPUR ANSU", "KISHORPUR", "DHARAMDAS", "MAHAL SARAI", "PAIBAGH", "RAILWAY COLONY ANO", "BHATOLI", "SOFATPUR", "KANWALNAINPUR", "RASOOLPUR SAID", "MUNEER GANJ AN0", "VIJAYPUR", "MOHD. AMIPUR", "NOORPUR DEHRA", "LAHAK KHURD", "SHAHALIPUR NAKI", "SHAHABPURA UMRAV SINGH", "SHERPUR HARSARAN", "MAQBARA AN0", "MAHARAIPUR SEKH URF DINORI", "KASHIRAMPUR", "BOREKI", "MOJJAMPUR TULSI AN0", "NARAYANPUR ICHCHA", "NANGLA UBBAN", "GIRDAWA SAHANPUR URF KHALAPAR", "LALPUR MAN", "DIWAN PARMANAND", "BARKATPUR", "BASHIRPU", "PRATAPPUR", "JATPURA BONDA AN0", "KORIYA")
		case "Nagina":
			sections = append(sections, "MASURI", "AHMADKHEL AN.H.NO 138-END", "BHAGWANPUR", "BHUDDI AN.H.NO 1-142", "GAZIPUR AN", "VISHNOI SARAI AN", "PALDI", "HASALIPUR MUTHRA", "AURANGJEBPUR CHANDA GAIR ABAD", "AURANGJEBPUR SAMBHAL", "KHUSHHALPUR MATHERI H.NO 1-101", "TAKIPUR HARVANSH", "PAKAHANPUR", "HUSAINPUR MEERA", "JAGANNATHPUR", "BHARAIKI", "PUNJABIYAN", "AKBARABAD  AN0", "TUKHMAPUR HARVANSH AN.", "CHAMARAN", "KHUSHALNAGAR GARHI", "HASANLIPUR BHOGAN", "SABHACHANDPUR MOHAN URF MAKHWADA", "NARULLAPUR BADLI GAIR ABAD", "BUDGARI AN.H.NO.1-498", "KADARPUR TAYYAB", "GALIBALIPUR MALUK", "MEMAN SADAT AN.H.NO 1-372", "SINGHADIYAN", "KAMALPUR  GAIR ABAD", "PIPALSANA", "PURSHOTAMPUR URF THEY PUR", "SAINAWALA", "HASANLIPUR KHANDU", "HARGAON CHANDAN AN.H.NO 1-130", "AMANNAGAR", "LUHARI SARAI H.NO 181-END", "ROSHANPUR CHIJROLI", "CHOHANAN AN.", "GOSPUR RAI GAIR ABAD", "SARAI MEER AN H.NO 117-END", "LOHARAN", "SAHLIPUR AB.SATTAR", "GANORA", "HAKIKATPUR VERCHAND", "MANIHARI SARAI AN H.NO 147 -END", "LOHARI SARAI AN.H.NO 1-180", "VALIPURA AN.H.NO 1-149", "MUBARAKPUR MEERA AN", "RASULPUR DAUD", "SAHZAHIR AN.H.NO 1-END", "MOTHEPUR KIRAT", "NURULDEHARPUR URF MIRZAPUR", "MUKANDPUR RAMU", "NANGLA ISLAM", "RAMPUR ASHA", "MUSSEPUR", "NURMOHDPUR", "TAKIPUR TULSI NG.CH.", "BANKHALA", "KASBA KOTRA AN. H.NO .", "AHMAD KHEL AN.", "MOHSANPUR", "KAJIPUR PACHDU", "ALAMPUR GANGA", "NALBANDAN AN.H.NO 1-141/2", "DAYALPU GYANPUR", "HAIBATPUR PEERU", "SARAI SHARIF NAGAR GAIR ABAD", "LAL SARAI AN.H.NO 1-102", "UMRI AN.H.NO 1-157", "KAYASTH SARAI AN 151-END", "SHADIPUR AN H.NO 56-END", "SHARIFULMALAKPUR", "KOTWALI AN", "SIKRI", "KALAN AN. H.NO 351-END", "DHANORI KUNWAR", "LUHARI SARAI AN.", "SHEKHSARAI AN", "KHATAI", "AJUPURACHOHAN", "ALIPURMAN", "HAMIDPUR VIDHICHAND GAIR ABAD", "GUJARPUR ASU GAIR ABAD", "KITHODA RAI", "DHARAMPUR BHAGWAN GAIR ABAD", "SARKADA KHERI", "NIJAMPURDEVSI", "MEMAN SADAT AN.H.NO 1-263", "KASSABAN", "LUKMANPURA", "NURALPIURBHAGWANT", "GOPALPUR", "KISHANPUR", "RAMPUR DAS AN.", "MOJAMPUR HARVANSH", "JAMALPUR BANGAR", "BARHAPUR", "DOODHLI", "CHANDUPURA NASEBPUR KANNA", "AURANGJEBPURHARDAS", "ASHRAFPUR", "NOORPUR HATTI", "KAYRI FUL GAIR ABAD", "MUBARAKPUR KHOSA", "SAHZAHIR AN.H.NO 121-END", "KHAKROBAN UTTARI", "SHAKARPURI", "GAJIPUR HIDAYAT URF GARHIBAN", "SAHLIPUR HEERA", "KALALAN AN.H.NO 1-350", "AKBARABAD  AN0 H.NO 549-END", "TARIKAMPUR GAIR ABAD", "SARAI JEEVAN", "MALAKPUR DEHRI", "SAYADWARA", "PREMPUR", "MUGLAN PASCHIMI", "MUFTIYAN", "DHAKI SADHO", "AMBEDKAR NAGAR", "NAKIMPUR RAMRAI GAIR ABAD", "MEMANSADAT AN.H.NO0264-END", "KAJIYAN AN.H.NO 1-127", "MOJIPUR DHARMA", "VAISHPUR KUDIYA", "BHUDDI AN.H.NO 143-END", "ANSARIYAN AN.", "SAHBAZPUR", "MEHMUDPUR NARAYANAN.H.NO 103-END", "SARAIMEER AN H.NO 144-END", "GOVINDPUR", "MUBARAKPUR MEERA AN.", "SULTANPUR SABHACHAND", "KAMALPURPURAINI", "JAFARPUR PARASRAM", "HASANPURA", "AKABRAN", "JHALRI", "AHMADKHEL AN.H.NO 1-137", "ALEYARPUR URF HAJIPUR", "MALAKPUR LAKHMAN", "JATI AN H.NO 1-45", "MANIRAMPUR", "BEGAMPUR RUPCHAND", "LOHARI SARAI PU.AN.H.NO 11/2-180/4", "MILKIYAN", "BUDGARI AN.H.NO.499-END", "KARONDA CHODHAR AN.H.NO.1-130", "CHATRIYANAGAR H.NO.325-877", "ANANDKHERI", "NARAYANPUR", "SHADIPUR AN", "BILASPUR GAIR ABAD", "BHOJPUR", "VISHNOI SARAI AN.H.NO 181-END", "RAMPUR BISHNA", "KOTWALI AN H.NO 477-END", "MOJAMPURDAYAL", "BASEDA", "SARAI MEER AN H.NO 1-265", "LAL SARAI AN.H.NO 356-END", "TAKIPUR MANOHAR", "ISLAMPUR HATTU", "KHUSHALPUR MATHERI H.NO 102-END", "KAKADPUR", "NAIMPUR GOVARDHAN", "AKBARABAD AN.H.NO 1-548", "BHOGLI AN", "MIRZAPUR", "MEHMUDPUR BHAVTA AN.H.NO 66-END", "RAJOPUR SADAT AN", "CHANDAK TURK", "SHEKH SARAI", "DELPURA KANNA", "KHURRAMPUR KHADAK", "RASULPUR JAGAN", "HAIJARPUR VEERCHAND", "MOHD PUR TAKI", "SABHACHANDPUR KESHOAN", "NASRULLAPUR", "DHARAMPUR DHANSI GAIR ABAD", "AURANGPUR HIRDAY", "KASBA", "JATPURA", "CHAKNARANGI", "CHITAWAR AN", "KALALAN H.NO 172-END", "AMIRPUR MAIDAS GAIR ABAD", "ABULFAJALPUR KHAS", "SAIDPURA GAJJU", "SHISHGRAN AN. H.NO.1-121", "SHADIPUR IMMAR", "AURANGJEBPUR MEHMUD URF MANGOLPURA", "SAHZAHIR AN. H.NO 1-120", "HAMIDPUR GANGA GAIR ABAD", "BEHADA CHOHAN", "BUDHPUR MUFTI", "KAJIPUR IMMA", "SAHABPURA RATAN SINGH", "RAJARAM PUR PRATAP", "MALAKPUR GAIR ABAD", "AURANGJEBPUR GULAL", "HEEMPUR MANAK", "NEKPUR", "CHANDAK JATT", "JHANDA AN", "RAFIKPUR", "KOTWALI AN HNO -1-123", "KALAPUR BUJURG", "ISLAMPUR SAHU", "KAYASTH SARAI AN.H.NO 1-150", "DHARAMSHANANGLI", "SAIDPURI", "ISLAMPUR VISHNOI", "KILA", "AFGANAN", "PADLI", "SAHJANPUR ROSHAN", "AFGANAN AN.", "RATANPUR", "HARBANSHPUR DHARAM", "HARDASPUR MADE", "KASBA KOTRA AN.H.NO 101-212", "NAJIBPURSADAT", "PAHADIDDARVAJA", "ALAMPUR", "SULTANKA GAIR ABAD", "BADSHAPUR MAIDAS", "SHADIPUR H.NO1-55", "ARGUPURA GAIR ABAD", "BHAVANPUR KALA", "NIJAMPURPADARATH", "LUHARI SARAI AN.H.NO 1-275", "MATHERA CHOHAN", "RAJOPURSADAT AN H.NO 126-END", "BAGHALA", "HAKIKATPUR GANGWALI AN.H.NO.1-209", "MALKAN", "MEHMUDPUR NARAYAN AN.1-102", "ISLAMPUR SADAT", "BEGAMPUR HARREY", "BEGAMPUR CHAIMAL", "AURANGJEBPUR FAJIL", "MANIHARI SARAI AN0 H.NO 1-146", "MUKIMPUR PADARATH", "MUKIMPUR DUNIYA GAIR ABAD", "ANWARPUR CHATAR", "ALAMPURASU G.ABD", "NOROJPUR GAIR ABAD", "SABHACHANDPUR", "AZMABAD", "SULTANPUR SADAT", "LUHARI SARAI AN.H.NO 276-END", "KAJIYAN AN.H.NO 128-END", "SAMASPUR VEERBHAN", "NEJOWALI GAWDI", "MOHD ASGARPUR", "SHEKHPURA NOABAD", "AURANGPUR FATTA", "ARJANIPUR", "MEHMUDPUR BHAVTA AN.H.NO 1-65", "KALAKHEDI", "MAHAPUR MOHD ALI", "SARAI IMMA", "NANDPUR", "RASULPUR ALIMUDDIN", "BEGAMPUR SHADI AN0", "MAKSUDNAGAR DEVIDAS", "RAIPUR MOJAMPUR NARAYAN", "MO ALIPUR SHEKH", "KOTWALI AN H.NO 1-412", "SARAI BAHAUDDIN", "AURANPUR NANDLAL", "MUBARAKPUR HADSA", "SHERPUR JAMAL GAIR ABAD", "HADIPUR SADRUDDIN", "KOTLA", "MO.ALIPUR SUKHANAND", "TARIKAMPUR DARGO", "MEMAN SADAT AN.H.NO 373-END", "PITTHANHERI JIYA AN.H.NO 114-END", "KAJISARAI AN.H.NO 1-73", "KALALAN AN.", "KHAKROBAN", "LAL SARAI AN.H.NO 1-150 TAK", "NAINPURA AN.", "MOMINPUR LALU", "BAKARPUR JAIRAM", "YARMOHDPUR PRITHVI", "ALHADADPUR", "MUSTAFAPURBHURAPURCHAKSAIDJAMAL", "CHODHRANA AN H.N0 1-124", "KALA NANGLI GAIR ABAD", "PADLA AN", "CHODHRANA AN0H.NO 125-END", "KHASOR", "JAMALPUR KHOKO", "MEHMUDPUR", "SUFIPUR ANGAD", "KIRTO NANGLI", "BURHPUR NAIN SINGH", "LALSARAI AN.H.NO 151-END", "JALPURA", "MANPUR", "LUHARI SARAI PURV AN.H.NO 1-283", "MUSLIM KATERA", "KARONDA PACHDU AN.H.NO 1-200", "KARONDA PACHDU AN.", "TAKIPUR NAROTTAM", "VISHNOI SARAI AN.", "KHAKROBAN ST.", "MANAKCHAND", "SALEMPUR GADHELI", "SHEKHPURA TURK H.NO 68-END", "LAL SARAI AN.H.NO 103-355", "ASDULLAPUR PRITHVI", "JATI AN H.NO 46-END", "BUDGARA AN.", "HINDU KATERA", "SAHUWAN", "MAHESHWARI JAT AN0H.NO 131-END", "HAKIKATPUR MUTHRA", "HARGAON CHANDAN AN.H.NO 131-END", "SAIDPURIMAHICHAND", "PITHANHERI JHOJHA KALA", "KARONDAPAHDU AN.H.NO 201-END", "NALBANDAN AN.H.NO 142-END", "SAHMUZAFARPUR", "MAHAJANAN", "LADPURA AN.", "MAHALKI", "SAIDKHERI", "HAKIKATPUR PRAYAG", "DARGOHPUR", "RAMPURLASHKARI", "KUMHERA", "BIHARIPUR", "RAMKAYRI", "MOHIUDDINPUR", "SULEMASIKHOPUR", "UGRSAINPUR", "UMRI AN.H.NO 158-END", "KAJISARAI  AN.H.NO 1-270", "KHURRAM ALI SARAIH.NO 1-115TAK", "MANJHLETA", "NAWADA", "MAHESHWARI JAT AN H.NO 1-130", "AJAMPUR MOHAN", "GOKALPUR SUNDAR", "SARAI JALAL GAIR ABAD", "HALWAIYAN", "BANWARIPUR", "VALIPUR AN H.NO 150-END", "SHISHGRAN AN.H.NO 122-END", "KAJISARAI AN.", "KALALAN AN.H.NO 1-171", "HAKIKATPURGANGWALI AN.H.NO210-END", "KIRATPUR", "JATAN", "SAHABPUR URF HINDUPUR H.N 1-150", "KHAIRULLAPUR", "MOHD ASHIKPURKAMALNAIN", "KOTWALIAN H.NO 413-END", "HAMIDPUR MAKHAN", "RAMPURCHAJMAL", "FATEHPUR RAJARAM", "CHIPIWADA", "DAYAMNANGLA", "SHARIFPUR KHORAJ", "ACHARJAN", "GHANGHERI", "PURSHOTTAMPUR BHAGWAN", "ABDIPUR KALYAN", "RANGDAN", "UMARPUR", "JEVANPUR", "METHA SAID", "JAFRULLAPUR", "KITHODA MUFTI", "TIGRI GAIR ABAD", "ISLAMPUR MEERA", "KAJISARAI AN.H.NO 271-END", "HAKIKATPUR GOVIND", "BAGWANAN", "KASBA KOTRA AN.H.NO 1-100", "ALIPUR GANGA", "MALAKPUR SEHSU", "PITHANHERIJHOJHA KHURD", "SAMASPUR KHORA URF KHODPURA", "MUBARAKPUR SAHARAN", "PITTANHERI JIYA AN.H.NO 1-113", "DULHAPUR GAIR ABAD", "ANSARIYAN  AN.", "TUKHMAPURHARVANSH AN0", "TODARPUR", "MUTHRAPUR URF KALUWALA", "RAIPUR GAIR ABAD", "BIRDO NANGLI", "CHAJUPURA", "HUSAINPUR SULTAN", "GARABPUR", "ASDULLAPUR KALYAN", "BRAHAMJITPUR CHANDA", "JALALPUR SULTAN", "SARAI MEER AN0 H.NO01-116", "ROSHANPUR PRATAP", "CHATARBHOJPUR KUSHAL", "KHURRAMALI SARAI AN.H.NO 116-END", "KAMALPUR BHOGA", "RASULPUR SAHBUDDIN", "KAMKARAN AN.", "HAKIKATPURGANGWALI AN.", "MUGLAN PURVI", "ज्ञKOTWALI AN H.NO 124-END", "BHANERA", "KHANPUR", "MIRZALIPUR BHARA", "AMBEDKAR NAGAR H.NO 233-324", "KARONDA CHODHAR AN.H.NO 131-END", "SAHLIPUR ASHA", "BAHADARPUR SHARFUDIN HUSAIN AN.", "HAKIKATPUR SEHSU", "NAKIPUR BAMNOLIGAIR ABAD", "MUSTAFAPUR GAIR ABAD", "MIRDGAN", "AJAMPUR YAR MOHD", "MOMINPUR DASU", "DUDHLA", "KHURRAMPUR DALLU", "MAMRAJPUR", "ALEYALIPUR", "GUNIYAPUR", "CHAHRONAK", "KOTWALI AN H.NO 1-476", "KHIJARPURJAGGU", "BUDHAWALA", "BUDHUPADA", "VISHNOI SARAI AN.H.N0.1-180", "NASIRPUR MITHARI", "RAMPUR MURAR", "SHERNAGAR NARAINI", "SARAI MEER AN H.NO 266-END", "MALAKPUR ABDULLA", "SARAI MEER AN H-NO 1-143", "AKBARABAD AN0", "SHEKHPURATURK AN H.NO 1-67", "SAHABPUR URF HINDU PUR H.NO 151-END", "KAYASTHAN", "KAJISARAI AN.H.NO 74-END", "SAHPUR SUKKHA", "GHANSURPURAMROLI", "MAHESHPUR", "PADHAN", "PATERI", "ISLAMPUR SAHLI", "ISLAMPUR BEGA", "ABDULAJIJPUR", "DELPURA NAINU", "LALPUR")
		case "Barhapur":
			sections = append(sections, "MUTAAZ HUSAIN AN0", "KAYAMNAGARDODHRAJPUR", "GOSPUR SADAT", "PAKKABAG", "BENIPUR", "TANDAMAIDAS AN0 1 TO 452", "PATHANAN", "MANKHANDEPUR GARHI", "ALIPURMOLAD", "NOMI MEDIUM", "NEMATPYRISKUPURA", "ALINAGAR", "NEJO SARAI", "ABHAYRAJPUR", "AANANDIPUR", "AJAMPURMANNI", "BHAJDAWALAKHALSA", "JAGANNATHPURGOKAL", "TELIPADA", "BHABANIPUR", "RAIPUR SADAT AN H.NO 201", "KHANAWALUDALLAWALI", "MUKANDOURRAHMAL", "RASULPUR METHEY", "HIDAYATPUR AN0", "ANOOPNAGAR", "KALLU NAWADA", "SIRYABALI", "ALAUDEENPUR", "AALAMPUR GABRI AN0", "TAYABPUR", "GOHAR ALI KHAN", "MACHHMAR", "BHOGPUR HNO549 TO END", "ADREHMANPURMEHJANG", "BHAGWANPURPRATAP", "NOAABAD JAGEER", "SHANAGAR KURALI", "HUSAINABAD AN. H.NO.1-277", "TANDAMAIDAS AN0 1 TO 794", "MOHDPUR ROJORE HNO 255 TO END", "MOHD. ALIPUR CHOHAR", "SHAHALIPUR JAMAN", "HAJIMODPUR KOTKADAR AN0 H.NO 1 TO 326", "RAMPUR", "SULEMASEKOPUR", "AGWANPUR", "VEERBHANWALA", "JATPURA", "ISLAMABAD AN0 1 TO END", "PUKHAY WALA", "BASUWALA", "JAHANABAD", "KADARPURNANU", "SALABATPUR GIRDHAR", "TANDA BAIRAGI", "KHERABAD", "SEERBASHU CHAND AN0", "KHURALLAPUR DESH", "MODPURTIRLOK AN HNO 1 TO 163", "INAYATPUR H.NO 1 TO END", "MAHSANPUR", "HARVANSGWALA", "TAKIPUR BEGA", "MUSSAYPUR", "BUNDKI", "RASOOLPURMUZAFAR H.NO 1 TO 954", "KHOKRA", "MAJHEDASAKRU HNO 228 TO END", "JAGANNATHPURMEHRU", "LAL SARAI HNO 1 TO 100", "HARKISANPUR JANULABDEEN", "TAILIYAN", "NARAYANPUR", "KUAKHERA H NO177 TO END", "NEJO SARAI AN0 H N0 1 TO 70", "SIKKAWALA", "ISMAILPURSALI BA", "HAJIMODPUR KOTKADAR H.NO 440 TO END", "GAJIPUR BHAGWAN", "AJIMULLANAGAR AN0", "PREMPURI", "ABILFAJALPURBANI", "KASHIWALA", "NEJO SARAI AN0  H NO 71 TO END", "FIROZPUR", "DAHLAWALA", "IMRATPUR", "HAMJAPURKATHAUR", "DULICHANDPUR", "MEERPUR GHASI", "ASLAMPURBHULLAN", "JHADPURA", "BHAGWANPUR  HN0 1 TO 125", "ABDULLAPURKURASHI", "KUAKHERA H NO 1 TO 176", "HASANPUR", "DHAKYA", "JAGATCHANDPUR", "MANNABARPUR CHAK", "TANDAMAIDAS AN0 1 TO 775", "MUNIR NAGAR", "RAM NAGAR", "CHAMRUNAWADA", "KASHRIPUR", "BANIRAMDASS", "MOMIN NAGAR", "AFGALPUR HARRAY", "CHHAJMAL WALA", "GOVINDPUR", "FAZALPUR", "BINJHAHEDI 1 TO 100", "JHILMILA AN.H.NO 1-85", "JOJHIYAN", "JANULABEDDIN AN0", "MOHD ASHIKPUR BHUREY", "ABDULREHMANPUR JEVAN", "PALPUR", "DHARMPUR", "CHANDPUR", "MADHOWALA", "QILA H NO 81 TO END", "BHAWANIPUR AN0", "BHOGPUR HNO 1 TO 548", "MODPURTIRLOK AN HNO 164 TO END", "BAWAN SARAI", "CHAKUDAYCHAND", "ASDULLAPUR KHALSA", "HARKISHANPUR", "ANWARPUR CHANDIKA", "MEDPURASULTAN", "KASSABAN", "DALLAWALA", "KHUSHALPUR DEWAN SINGH", "AJAB NAGAR", "MAHARATPURKALA AN0 H.NO 148 TO END", "CHIRANJI LALA", "THAKURAN", "RAI BHAN WALA", "AARAJI KANDALA", "MAHABATPUR", "HUSAINABAD AN.H.NO 1-277", "GADRIYAN", "SEEPAH", "NARAYANWALA", "ALHAYPUR MOHKAM", "RAIPURSADAT AN0 445 TO END", "DAUDPUR", "PAGAMPUR URF BAGIR GANG", "CHAKCHOHADMAL URF BHAJRHAWALA", "MAKHDOOMPUR", "LAL SARAI HNO101 TO END", "SAIDPUR HARIAWALA", "KUDDUSPUR", "NAIWALA", "CHAKARPANPUR", "HAJIMOHD PUR KOTKADAR  AN 1 TO 110", "BAZAR", "ALHAHEDI", "ALIYARPUR", "SHAHLIPUR GADDU", "TARAGANG", "KHEJARPUR", "MAHAPUR", "HASANLIPUR HEERA", "KUJAITA", "BHAGWANPUR  HN0 126 TO END", "PAYBAND KHERI", "MAHARATPURKALA AN0 H.NO 685 TO END", "BINJHAHEDI H.NO 101-END", "nawabpura", "MOHDPUR ROJORE HNO 1 TO 159", "SULTAN AJAMPURBANI GNESH", "KASAMPUR GARHI AN0", "HEERAPUR MALI", "JAFARPUR", "DESONDIWALA DESH", "DOARKADASPUR", "SEERWASHU CHAND AN0", "ALHAYPUR BEERAN", "FAJULLAPURMAHESH", "MAGHPUR", "HASANLIPURDHARMA", "BABLI", "JAMALPUR DHIKLI", "MAJHOLI H NO 51 TO END", "SHAHLIPURKOTRA  H NO 1 TO 178", "HEERAPUR GUKAL", "NOMI SOUTH", "NOABAD DESH", "BHATT PURA", "RASULPUR ABAD", "MOHIDINPUR", "MITHOPUR", "DHARMAWALA", "HAJIMODPUR KOTKADAR 111 TO END", "CHAMARAN", "RAIPUR SADAT H.NO1 TO 444  END", "ALIPUR MITTHAN", "NOIABAD JANGAL", "BADSHAPUR", "IMAMPURTIGRI", "HASANLIPURMAAN", "REHMAPURUSMAPUR", "HAREBALI AN0", "JASSUJOT", "PARMANANDPUR GAWDI", "SHIVPURI", "SEERBASHU CHAND", "NANGLA LADDHA", "MURTUZABAD", "MH. PUR BHAJJA", "CHAKMANI RAGHUNATHPUR", "ISMAILPURDAMI", "SUAWALA AN0 H.NO 1 TO 408", "MANAVARPYURSAID HNO TO 1 TO 110", "AFJALPIRBHAU", "BARKHERA", "AJIMULLA NAGAR AN0", "MUSHTAFFABAD GARHI", "MAJHEDASAKRU HNO 37 TO 454", "JAHANGURPURKHASS", "TURATPUR", "KHANPURMANAK", "SULTANPUR", "NOORPURARAB", "SUNDARWALI", "SHALIPUR BA", "MIYAZI MOKHA", "BEDARBAKTHPURJADO", "RASOOLPURMUZAFAR H.NO 9 TO 276", "MAHARATPURKALA AN0 H.NO 1 TO 684", "TAMURABAD", "PADARATHPUR", "MOINUDDINPUR", "DLEEPUR", "HUSAINABAD AN.H-NO 278-END", "RASULPUR ABAD AN0", "JAMANPUR", "NARAYANPUR RAMJI", "ISLAMABAD AN0 20-394", "HAJIMODPUR KOTKADAR H.NO 1 TO 923", "SULTANPUR KHAS", "KALLUWALA AN0", "JALALPUR", "SHANKARPUR", "MOHDPUR ROJORE HNO 1 TO 254", "MURADNAGAR", "MEERAPUR VEERAN", "GHASSIWALA", "KAHRIPUR JANGAL", "MAJRA BHARATPUR", "SAMASPUR NASEEB H.NO 1 TO 150", "NATHEYWALI", "HARPUR", "SARAI DADUMBAR", "SHADIPUR", "AURANGJEBPURSAHLI", "GADLA", "GHASIWALA", "HARGANPUR H NO 142 TO END", "KRISHNA NAGAR H.N0 493TO 512", "GOPIWALA", "ALIPURA JATT", "SABDALPURDEJAHGIR", "BEGAM SARIA", "CHAPRHA", "KANAKPUR", "MURTJAPUR URF PUNAPUR AN H.NO 1 TO 220", "TIRLOKPUR", "HAJIMODPUR KOTKADAR H.NO 1 TO 439", "RANI NAGAL", "RAMPUR KALA", "JALALPUR SULTAN", "MAHARATPURKALA AN0 H.NO 1 TO 147", "BAHEDI", "DHARMPUR URF JOGIWALA", "AJAMPURPARMA", "KADARGANJ", "RAWALHERIKHAJIRI", "SHAHALIPUR KOTRA AN0", "MEERAPUR MODIWALA AN0", "RAMNAGARGOSAI", "RAMPUR GARHI", "TANDAMAIDAS AN0 151 TO END", "FAGLABAD PARMANANDPUR", "ABABKARPUR", "RAFAYATPUR HULLAS", "SARFUDEENNAGAR AN", "WAJIDPURLALA", "SAHUWALAGOSAIWALA  HNO 106 TO END", "JAYNAGAR", "RAIPUR SADAT AN", "JAFRABAD", "UMARPURBERKHERA", "MORGHANDI", "WANSIJOT", "HARGANPUR H NO 1 TO 142", "SHAHJAHAN PUR GARHI", "BAHADAR NAGAR", "ABDIPURHARVANSH", "CHOHADHPUR NANU", "RAJUPURA H.NO 1 TO 239", "NOMI UTTARI", "FATHAY ULLAH NAGAR", "KIRATPUR", "DAULLATPUR", "HARISINGHKABHOGLA", "MANAVARPYURSAID HNO 111 TO END", "SHARGHAR AN0", "KALYANPUR H.NO 1 TO END", "REHAR AN0", "SHAJADPUR AN0", "JAY SINGH JOT", "SHYAMPURSARVANKA", "AASFABAD CHAMAN", "AMAN NAGAR", ".JATNANGLA", "SAHUWALAGOSAIWALA  HNO 1 TO 105", "RASULPUR VEERAN", "RAMPUR ASAF", "SHAHALIPUR", "CHAK", "SAMASPURSADDO", "SUAWALA AN0 H.NO 251 TO 594", "FATTUWALA", "KANHADAWALA", "JASPUR", "SARDARPUR", "TARAPUR", "RAIPURSADAT  H.NO 1 TO 200", "SAEED NAGAR", "WAJIDPUR", "SHAHLIPUR", "BHAGIJOT", "UMARPUR NATTHAN", "SHAHJAHAPUR", "NEMATPUR", "GOWARDHANPUR AN0", "MADPURI", "AURANGABAD", "SADATPURGADI", "MAJHOLI H NO 1 TO 50", "LALPURI", "UDHO JOT", "NAYAK SARAI", "DALLIWALA", "FAGLABAD PARMANANDPUR ANO", "SHAHPUR JAMAL A", "SHAJADPUR", "KAMRUDINNAGAR", "AJUPURARANI", "SAHARAWALA", "MIRZALIPUR CHOHAD", "JABTA NAGAR", "HERWELA", "PAGAMBERPUR", "MAKSHOODABAD", "HARIJAN BASTI", "LAKHIWALA", "MURTJAPUR URF PUNAPUR AN H.NO  221", "KOTKADAR HAJI MOHDPUR H.NO 327 TO END", "SEER SHIVRAJ SINGH", "ALLAHDADPURKHAJWA", "SARAI PURAINI", "CHANDLIKA", "RAIPURI", "FATEPUR A", "RAJPURKOT", "CHATARPUR", "KALLUWALA", "KISANPURKUNDA", "MANSHA", "PREMNAGAR", "SUAWALA AN0 H.NO 1 TO458", "PHOOLBAG", "BARKATA", "JHANABAD AN 1 TO 875", "NANJIWALA URF RAMJIWALA", "HALLOWALI", "MURADBAKSHPUR", "AMEERPUR MANJU", "FATEPUR DHARA", "SHARGHAR AN0 HN0 1 TO 376", "SATTIYAN", "SADA SHIVPUR GABRI", "TANDAMAIDAS AN0 1 TO 150", "DEEVANANDPUR GARHI", "MAJHEDASAKRU HNO 1 TO 227", "RATUPURA", "GUSAIWALA", "AJIMULLA NAGAR", "GANGARAMWALA", "JHULA", "MH. PUR URF JOGIPURA", "HEEMPUR MAKHDUM JAHAN", "D.S.I.L DWARIKESH NAGAR", "pipli", "MOHD ALIPUR PATTAGHAT", "MAHARATPURKALA", "SHAHLIPURKOTRA  H NO 179 TO END", "SAMASPURNASEEB H.NO 151 TO END", "TIGRIANOOP SINGH", "QILA 1 TO 80", "MAHENDER NAGAR", "CHAMPATPURCHAKLA", "JHILMILA AN.H.NO 86-END", "ASAFPURSAIDPEER", "ALAMPUR", "NURULLAPUR", "ALIGANG", "E", "GAJIPURSABHACHAND", "PATPADA KESHO", "SABUWALA", "LADPUR", "ISLAM NAGAR AN0", "KHANPUR", "MIRZALIPUR BHARA", "AMBEDKAR NAGAR H.NO 233-324", "KARONDA CHODHAR AN.H.NO 131-END", "SAHLIPUR ASHA", "BAHADARPUR SHARFUDIN HUSAIN AN.", "HAKIKATPUR SEHSU", "NAKIPUR BAMNOLIGAIR ABAD", "MUSTAFAPUR GAIR ABAD", "MIRDGAN", "AJAMPUR YAR MOHD", "MOMINPUR DASU", "DUDHLA", "KHURRAMPUR DALLU", "MAMRAJPUR", "ALEYALIPUR", "GUNIYAPUR", "CHAHRONAK", "KOTWALI AN H.NO 1-476", "KHIJARPURJAGGU", "BUDHAWALA", "BUDHUPADA", "VISHNOI SARAI AN.H.N0.1-180", "NASIRPUR MITHARI", "RAMPUR MURAR", "SHERNAGAR NARAINI", "SARAI MEER AN H.NO 266-END", "MALAKPUR ABDULLA", "SARAI MEER AN H-NO 1-143", "AKBARABAD AN0", "SHEKHPURATURK AN H.NO 1-67", "SAHABPUR URF HINDU PUR H.NO 151-END", "KAYASTHAN", "KAJISARAI AN.H.NO 74-END", "SAHPUR SUKKHA", "GHANSURPURAMROLI", "MAHESHPUR", "PADHAN", "PATERI", "ISLAMPUR SAHLI", "ISLAMPUR BEGA", "ABDULAJIJPUR", "DELPURA NAINU", "LALPUR")
		case "Dhampur":
			sections = append(sections, "VAJIRAPURLAL AN.", "CHANARAN", "NANGALA JAJAN", "FATEHAPUR LAL B.A", "PIR KA BAJAR AN.", "BUAPUR HARAPAL AN.", "MUKRAPURI AN.", "JAMALAPUR MAHIPAT", "RASULAPUR MOHD KULI AN.", "IMAMABAD AN", "MOHD.ALIPUR ABHYACHAND", "BASEDA KHADAR", "RAGHUNATHAPUR", "NINDUDU KHAS AN.", "KHUJISTANAGAR", "LATIFAPUR URF SIPAHIVALA AN.", "MIJARPUR PALLA", "NAVADA SAIDAPUR JAMAL URF BHOOTPURI", "UMARPUR BANGAR B.A", "SHAURAMAPUR", "ASAFABAD BHAWAN URF GANVDI", "NIRULLAPUR CHIMMA", "CHAKARAJMAL", "HUMAYUMPUR CHANDRAPAL", "BUAPUR KHEM", "SADATABAD", "RANAVLI", "BAJAR KALAN", "HALAVAIYAN", "RAFIPUR", "RASULAPUR MOHD KULI . AN.", "BAJIRAPUR JANGIR AN.", "NAVADA KESHO", "TAKI SARAY", "MANDAUNRI AN.", "JAMAPUR", "JAGUPURA", "SHERAPUR", "SHERAPUR KALIYA", "HAKIMAN AN.", "BURHANNAGAR", "JAMALAPUR UDYACHAND AN.", "MUBARAKPUR TAPPA BHARTI B..A", "NANGALA GANNA", "NAI SARAY", "KODUPURA", "SADARNAPUR", "TARAI AN.", "AMAKHEDA SHAJRAPUR AN.", "PALKI EMAN", "NANDAGANVA", "AURANGASHAHAPUR NARAYAN", "SARKATHAL SAHANI AN.", "CHAK MOHD NAGAR", "NANGLA NANSAI", "SAIDAPUR BIRU", "RASULPUR MAOH KULI . AN.", "MANAPUR SHIVAPURI", "TARAIR AN.", "AFAGANAN AN.", "NAUDHNA AN.", "SHEKHPUR THATH", "BALKISHNAPUR", "MOHD.ALIPUR LALMAN URF RAMSAHAYVALA", "SALAVA", "ALAVALPUR", "LAIDERPUR", "MOHD.ALIPUR INAYAT", "MILAK BENIRAM", "AJITAPUR DASI", "HUSAINPUR HAMEED", "HARRA AHAMADPUR JALAL", "BUAPUR NATTHU AN.", "RAYAPUR", "DHAMPUR HUSAINPUR AN.", "HAFIJABAD SHAKI AN.", "BAKAR KASSABAN", "KHANUJAT", "MOHD.PUR SULTAN", "SUHAGAPUR AN.", "NATHADOI", "KASBA AN.", "DATANGAR", "SHAHAALIPUR UDYACHAND", "SADULLAPUR KHANAPUR", "PURNAPUR", "SHARIFAPUR", "LOHIYAN AN.", "PAHADI DARAVAJA AN.", "NANGALA BUDHVA", "SALARAPUR  B.A", "MAUJAMPUR JAITRA AN.", "DINADARAPUR", "NANGALA BHAJJA", "NIN‍DOODU KHAS AN.", "SAMANASARAY AN.", "MAUHADI", "DHURARA", "UMRAPUR KHADAR", "SADADOBAI AN.", "NURAPUR CHHIBRI AN.", "NAYAK SARAY", "TARAKAULI MANDN", "GOVLI B.A", "DITNAPUR AN.", "SHEKHAN AN.", "GAJUPURA", "GURADASAPUR HIRA", "KOTRA AN.", "NURALLAPUR UDYACHAND", "SHERAPUR RAINI", "MANDI MARKAM GANJ", "PALNAPUR", "NAYA GANVA MAJRA", "KOPA", "WAJIRPUR BHAGWAN", "NANGALA GUNGA", "KADARABAD KHURD", "MAUJMAPUR SURAJ", "RAMPUR DULLI", "BAKALAN", "HINDU CHAUDHARI", "ISLAMNAGAR AN.", "KAYASTHAN AN.", "MOHD.PUR JAMAL", "MITTHEPUR", "ALHAIPUR AN.", "MU.CHAUDHARIYAN", "FATAHANGAR AN.", "MOHD.PUR SADA", "GUJARATIYAN", "KAHARAN AN.", "MAUHADA AN.", "SADATNAGAR", "TAKI SARAY AN.", "HAKIMPUR SHAKARGANJ AN.", "SHEKHAPUR BHAVDA", "ACHARAJAN", "BADASHAHAPUR LAXHMI SEN", "FATAHANGA AN.", "PATAVARIYAN AN.", "KHANAPURSIRIYA", "SARAKATHAL MADHO", "MAIHARI SARAY AN.", "SALEM SARAY", "BAMANAULI", "BADAVAN AN.", "RAMPUR BHOLA", "HARIYANA", "JUMERAT KA BAJAR.", "JAMALAPUR ALAM", "MOHD.ALIPUR BHIKAN", "VIRTHALA", "KHALILAPUR", "HAYATNAGAR", "GAJARAULA B.A", "MILAKIYAN AN.", "AURANGSHAHPUR BAUVARI", "AJMAPUR JAMANIMAN", "HAJIPUR", "KAJI SARAY", "FAJULLAPUR AN.", "BERAKHEDA TANDA AN.", "SHAFIYABAD AN.", "THATJAT AN.", "SABADLAPUR AN.", "KAHARANA AN.", "CHAK SHAHJANI", "HASNAPUR PALKI", "MU.CHAUDHARIYAN AN.", "SHAHAJADAPUR TARU", "KESHOPUR", "BADASHAPUR SHAH MAUHM‍MAD", "CHAUDHARIYAN", "ABUNASRAPUR", "SARKADA CHATRU AN.", "JASAMAUR", "RAMATHERA", "NASIRAPUR BANAVARI", "UMARAPUR ASHA AN.", "AMIRAPUR", "NIKAMMASHAN AN.", "MUSATAFAPUR TAIYAB", "KHATRIYAN AN.", "SAHUVAN", "FATAHAPUR JAMAL AN0", "BAKHSHANAPUR AN.", "SALARABAD", "IBRAHIMAPUR NARAYAN", "HAKEEMPUR MEGHA", "BHAJDI SARAY", "CHANCHAL PURAYAN", "JOTAHIMMA AN.", "MOHD.PUR VEERU", "NANGLI LADAN AN.", "NOORPUR CHHIBRI AN.", "ALADINPUR BHATPPURA", "SADADOBAIR AN.", "BUNDUKACHIYAN AN.", "KHURADA AN.", "TEEVRI AN0", "SHAHALIPUR NICHAL", "RASULPUR IMMA", "KAJIYAN", "WAJIDAPUR", "BAJIRAPUR URF MADHIYA")
		case "Nehtaur":
			sections = append(sections, "WAJIDPUR", "SURPUR AANSU GAIR AB.", "NARAYAN KHAIDI", "AMRABAD", "AMIRPUR GANGU GAIR AB.", "KASHMIRI ABUNASARPUR", "PALIJAT", "JOSHIYAN", "MALAKPUR", "RAISAN U. 1,2", "LATEEFPUR", "SAIDPURA", "PEER SHAHEED KALA", "NANGLA SARI", "SHAFIPUR HEERA", "KUMRARA", "SHAMSABAD", "PADLI MANDU", "TAHARPUR ASKARAN", "BALAPUR", "FARIDPUR MALHU", "BASEDA KHEMCHAND", "FARIDPUR DARGO", "DAKKHANA KHAS SHUMALI", "RUGDI CHETRAM", "DOULATPUR SUKHKHA", "TARKOULI EMMA", "SIJOULI", "CHOUDHRIYAN PU.", "RAZANAGAR", "FULSANDA HEERA", "FARIDPUR SADHIRAN", "YAKUBPUR GAIR AB.", "CHUDYAKHEDA GAIR AB.AD", "FARIDPUR MAN", "AHMADPUR CHANDRU URF GDANA", "ISLAMPUR CHAMRA", "AKBARPUR GARAV", "SULTANPUR ABAD", "KOKAPUR", "ISLAMPUR LALU", "GULI TALAB", "AMBARPUR ZAFAR", "KADRABAD", "SAHANPUR", "KAREEMPUR MUBARAK", "ELAHABAD", "KAZIPURA", "SIHORA GIRDHAR", "CHOUDHRIYAN P.", "SIKRI BUZURG", "KAKRALA", "ISMAILPUR", "DAKOULA", "CHHAJUPURA SAID", "SIKANDARPUR BUDH SINGH ABAD GAIR AB.", "SHAHKARAMPUR GILARA", "GINDOURI", "MANDORI", "SHEKHAN SHUMALI", "SHEESHGARAN", "KHIJARPUR", "PANCHAYTI MANDIR", "MIMLA", "AMHEDI GAIR AB.", "TAKIPURA", "CHOHARPUR GAIR AB.", "BAGWAN CHAPEGARAN", "NARGADI", "HARBALLABHPUR GAIR AB.AD", "RAYPUR BEGA", "HAKEEMAN", "MOHSANPUR KALYAN", "SADKABAD", "KIRARKHEDI", "MEHARPUR", "PEEPJAN", "KOTRA TAPPA KESHO", "NURUDEENPUR", "KHANMPUR CHAK GAIR AB.AD", "ABDULPUR MUNNA", "FARIDPUR NIZAM", "BANGADPUR KISHNA GAIR. AN.AD", "IBRAHIMPUR SAHDO", "JAMALPUR", "PALI TEKCHAND", "FULSANDA GANGADAS", "ARAZI GOPAL JOT", "UCHA", "BALDIYA", "AMHEDA", "RUSTAMPUR", "GOWARDHANPUR NANGLI", "AKBARPUR RADHE GAIR AB.", "GUJARPUR JASPAL", "PAHARPUR", "NIZATPUR", "BASEDA UBAR", "BALKISHANPUR GAIR AB.", "MANGU CHARKHI", "TEERGARAN", "PAHARPUR CHATAR", "KUNDA BAGAIN", "BHARAKHERI EMMA", "MUKARRABPUR", "HASANPUR JAT", "MUSTFABAD", "CHANDA NANGLI", "HARDASPUR GAIR AB.AD", "TAHARPUR TAPPA UMRI GAIR AB.AD", "DHELI GUJJAR", "PRATHVIPUR BANWARI", "MUZAFFARPUR DEVIDAS", "TAPROULA", "KHEDA D. 2", "FARIDPUR SANSARU URF DHOKALPUR", "MOMINPUR GULARI", "JYOTHI", "BISATH", "SAMASPUR HUSSAINPUR", "HOLI", "AKHERA", "KHAJURA JAT", "MOH.PUR PARMA", "BHATYANA KHUSHHALPUR", "PADLA", "DHELA GUJJAR", "CHOHARPUR", "BILASPUR", "ATAPUR URF KHATAPUR", "TAKIPURA PURAN", "JALALPUR HASNA", "KHAIRABAD", "RUSTAMPUR WAJID", "TAHARPUR TAPPA NANGAL", "AASPUR NAWADA", "ENAYATPUR", "NASIRPUR NAIN SINGH", "BASEDA KHURD", "RAISAN D. 5", "ISSAPUR", "FARIDPUR ALAM GAIR AB.", "SHAHBAD", "BASAWANPUR", "RANGREZAN", "MIRDGAN", "CHACK GOWARDHAN", "SHURPUR GAIR AB.AD", "HARPUR", "SHADIPUR", "PUTTHA 2", "DABTHALA", "SIKERA", "ILAYACHIPUR KHADAGU", "EIDGAHA", "SALLHAPUR 1", "KHALILPUR", "MARUFPUR", "FATEHPUR ASAL", "MUZAFFARPUR", "CHAK ABDUL REHMAN", "TAKIYA GARI PASHIMI", "CHACK MUSTAFABAD GAIR AB.", "LATIFULLAPUR", "PAHARPUR BAST GAIR AB.", "MO.PUR LAL", "GANGODA JAT", "GARAVPUR", "FARHA SAN SHUMALI", "MAHU NANGLI", "RUSTAMPUR SHER", "DABKHEDI SALAR", "ISHAKPUR", "SAHANPUR NAWADA", "FIROZPUR UGRSAIN", "SAIDPURI GAIR AB.AD", "MATAPUR GAIR AB.AD", "SARAY QAZIYAN", "BAKAR NANGLA", "IBRAHIMABAD", "BAIRMABAD GADI", "HAMEEDPUR", "IBRAHEEMPUR LAL", "CHAPEGARAN D.", "MEMUDPUR MILAK", "DAWARPUR GAIR AB.AD", "AMINABAD", "AOURANGSHAHPUR GANGDHAR", "SARAY TALE WALI", "SETHPUR DHANESAR", "MAHUAA", "BAKARNAGAR", "QAZIYAN", "SALEMPUR RUKHADIO", "ALHEDADPUR MUBARAK", "RAISAN D. 3", "BHARAKHAIDI URF DUGRU", "HALWAIYAN", "KHEDA U. 1", "PUTTHA 1", "DAKKHANA KHAS JUNUBI", "MATOURA DURGA", "PUTTHA", "DUWARKAPUR", "IBRAHIMPUR KHANDSAL", "KHANDSAL", "RAJPUR CHOMELA GAIR AB.AD", "SHAHPUR SAISU", "RAGHUNATHPUR", "MAZHARPUR", "NAWADA SAHANPUR", "NASIRPUR BUZURG", "ASADGAR", "SNUPTA", "FALOUDI", "CHAPEGARAN U.", "GOPALPUR", "GOVINDPUR GAIR AB.AD", "GYASPUR", "RAWANPUR", "PAPSARA", "KASAMPUR AHMAD", "SARAY RAJAB ALI", "KUKDA ISLAMPUR", "KINAN URF KALYAN", "RAISAN D. 8", "DARBAR SADAT", "RAJDEV NANGLI", "SARAY ASNARA", "SADRUDDINNAGAR", "MOLVIYAN JUNUBI", "LADANPUR", "CHAK KUKDA GAIR AN.AD", "AHIRPUR GAR.AN.", "KHOSPURA", "JAMALPUR SHAIKH", "TAKIYA GARI PU.", "BASEDA NARAYAN", "JALALPUR TURK", "NAWAJISPUR AHMAD", "MUNEEMPUR", "SLAPUR SHAFKATPUR", "RASULDARAN", "MANKUAA", "NARAYANPUR GAIR AB.AD", "RAGHORAMPUR", "RUKANPUR", "FAREEDNAGAR", "RAISAN D. 4", "SHADIPUR KALA", "FULSANDA KHAKAM", "KHURDI", "HARDASPUR GADI ZAFRABAD PRATHAM", "BERKHERA CHOUHAN", "MOLVIYAN SHUMALI", "MILAK MUKEEMPUR", "CHOUHARA", "ALAWALPUR UDDA", "PANCHAYTI MANDIR UTTARI", "ISLAMPUR", "RAISAN D. 6", "SHAIKHPURA PITTHA", "HARGANPUR", "BAIRAMNAGAR", "BHAWANIPUR TARKOULA", "FATTANPUR", "GOGALI", "HAKEEMPUR NARAYAN", "SULTANPUR VIRAN GAIR AB.", "TARIKAMPUR", "MO.ALIPUR MUKTA", "ALINAGAR PALNI", "AAKU", "DHINGARPUR", "ZRIFPUR CHATAR", "RAYPUR MALIYABAD", "HAKEEMPUR CHANDAN", "FARHASAN JUNUBI", "FARIDABAD", "CHOHADPUR", "KANDKHEDI", "NANGAL JAT", "RUGDI SAID", "INAMPURA", "MUSSEPUR", "SIJOULA", "FARIDPUR DULLI", "BEGRAJPUR", "MEHMUDPUR BUZURG", "MOMINPUR LAL KUWAR", "DHNORA", "TAHARPUR SAID", "MURSHADPUR", "KHEDA U.2", "KHAIDIJAT", "DAKOULI", "MANDI BAZAR", "KHEDKI TAPPA NANGAL", "SALLAHPUR 2", "NOORPUR KHAIDKI", "JOGIDASPUR GAIR AB.", "BARUKI", "MO.KULIPUR", "AMIRPUR GANGU", "ABDUL KHALIKPUR BALRAM", "KISHANPU BHOGAN", "PRATAP URF KHERPUR", "PIPALSANA", "NAUDHA", "KASAMPUR KRIPARAM URF SUNDARPUR", "HAKEEMPUR DULLA G.AB.", "SIKRI KHURD", "SIKANDARPUR", "SHAIKHPUR LALA", "NASEEBPUR", "DHANORI", "CHAJUPURA GANDHU", "HATHI MANDIR", "RASULPUR VIRAN GAIR AB.", "FALOUDA", "MATOURAMAN", "ATHAI SHAIKH", "IBRAHIMPUR URF KUMHARPURA", "ALADEENPUR BHOUGI", "BAGWAN NAKTA", "RAJPUR", "DHAKKA KARAMCHAND", "NEZE SARAY", "AFGANAN", "RAYPUR LAKRA", "MILAK GANGODA JAT AN.", "KHETAPUR", "KALALI", "SHAHNAGAR", "ISLAMPUR THAMBU CHAU", "UMRI", "SHEKHPURA", "SARAY JOKHA SINGH", "RAISAN D. 1", "NANHEDA", "LAKHARAN", "MAHAKMA JUNUBI", "KANHA NANGLA", "SAIDA", "BHAUKHEDA GAIR ABAD", "RAISAN D. 2", "KHEDA U. 3", "SEDI", "SAMASPUR URF MANAKPUR", "BURPUR", "NAKIBPUR", "SHAPPAR SHIKOHPUR", "DAKSHIN SARAY", "ALIPUR DAMODAR", "RAMPUR", "SADAT SAHADRI", "MILAK JAHAGIRABAD", "RAYPUR MULAK", "PAHARPUR KAMALPUR", "KASBA", "GANGU NAGLI GAR. AB.AD", "BUDIKA GAIR AB.AD", "SHERPUR", "DHUNPURA", "GAZIPUR", "BAIBALPUR", "MAHMUDPUR JAGMAL", "FATEHPUR", "KAYSTHAN", "MEHMODPUR KAMIL", "HAJJAMAN", "SADATNAGAR", "MO.HUSAINPUR GAR ABAD", "FAIZPUR RAMANAND", "BAIRAMPUR KHAJURI", "FAZALALIPUR", "TURABNAGAR", "SHAIKHAN DAKSHIR", "MANDORA VIP", "AKBARPUR DEVMAL", "TARKOULA", "RAISAN D. 7", "AMHEDA 2", "FIROZPUR", "KASAMPUR LEKHRAJ", "MURADPUR GAIR AB.", "JALALPUR AASRA", "MIRZAPUR", "SALAMATABAD", "ANWARPUR ADIL GAIR ABAD", "CHAKARPUR", "KHEDA D. 1", "BARKAPUR GAIR AB.AD", "MITHAN KUWAR PRATAP SINGH", "BAMNOLI", "JASMOURA", "MEMRAN", "SHAHPUR SAIDU", "KESHODASPUR GAIR AB.", "MEHAR ALIPUR", "MEHAKMA UTTARI", "FAZALPUR", "GANGODA SHEKH", "DARBARPUR", "GADAL", "BAHADA GAIR AB.AD", "MUZAHIDPUR", "MEMUDPUR BHIKKA", "MUKRANDANPUR MANAK", "FARIDPUR DALLU", "RAISAN U. 2", "NAIKOFAL URF BILAI", "MUZAFFARABAD", "SHERPUR BALLA")
		case "Bijnor":
			sections = append(sections, "AHMADPUR  MOHIUDDINPUR", "AKBARPUR DEVIDASWALA", "FIROZPUR TAREEKAM", "SHAHPUR LAL", "BRAHMNAN", "ALEDADPUR", "RAMBAG", "TAREKAMPUR PARAS", "AHAMADPUR BHARTA", "GANGOI KHADAR", "BHARUKA", "RASULPUR PITTANKA", "MIRZAPUR BANGAR", "CHHAKDA", "KALYANPUR", "CHAUHANPURA", "MUKARPUR GUJAR", "CHANDRBHANPUR KISHOR", "RAIPUR BERYSAL", "AMIPUR BEGA", "MIRDGAN", "RATANPUR RIYAYA", "MEEROPUR AGRI", "MUSHARFABAD", "RAGHDAN", "ULAKPUR", "SHADIPUR", "KHWAJGIPUR", "AMIPUR SUDHA", "HAKIMPUR", "NAGLI", "FATEHPUR KHURD", "MIRGIPUR", "QAZIPARA JANUBI", "JEETPURA", "KHATRIYAN B-13", "BAZAR SHAMBA", "ALLAUDDINPUR", "CHOUDHARYAN", "MANDRO JATT", "NASRPURAN NAJAMUDDIN", "FIROZPUR MANDU", "KHEDKI HEMRAJ COLLONY", "TEEP", "SULTANPUR TAPPA NANGAL", "MURTAJAPUR BULAKI URF PEDA", "SIMLI", "UMARPUR", "SAIFPUR BANGAR", "MODHIA DHANSHI", "DARANAGAR", "FARIEDPUR MIRA", "TAREEKAMPUR ROOPCHAND", "MOHMADPUR MANDAWALI", "SADAKPUR URF BILASPUR", "JHAKDI BANGAR", "MODHIA BHOGAN", "MOHAN PUR", "MUKEEMPUR DHARU", "KANUNGOYAN", "MIRDAGAN", "JALALPUR KAZI", "GAJRAULA SHIV", "ZAFARPUR ANSHU", "QAZIPARA JANUBI B", "SADUPURA", "BADSHAHPUR", "ACHARJAN", "BHATAN", "FARIDPUR KAZI AN", "NASIRI", "TAIMOORPUR", "KHADAK", "FIROZPUR MUBARAK URF NAYA GAOI", "ALIPUR SUKHANANDPUR URF LALPUR", "NAWADA", "RAMSAHAYWALA", "NAI BASTI B-24", "HIMMATPUR BELA", "SHAHBAZPUR KHANA", "MACHKI", "JULAHAN", "RAFIULNAGAR URF RAWALI", "BURHANUDDINPUR", "SWAHEDI KHURD", "JAHAGERPUR", "MAHAJANAN", "DAIVALGARH", "MANSHAPUR", "SALMABAD", "MOHIUDDINPUR", "GHUDIYAPUR", "AKBARPUR ANGAKHEDI", "MANDAWLI SAIDU", "FARASHTOLI", "BHATAN AN", "PEDI", "BHAGIN", "SHAJADPUR", "KAZIYAN", "BHADARPUR KHURD", "VASIHPUR MAN", "AURANGBAD SHAKURPUR", "MUNDALA", "GAUSPUR TOPARI", "NANUPURA", "BAZDARAN", "PADAMPUR", "NAWALPUR", "SHABAJPUR", "ISLAMPUR NEEMDAS", "AWAS VIKAS COLONY", "KHAKROWAN", "CHANDPUR FERU", "AHMADPUR SULEMAN", "BARKALA", "RAJARAMPUR KHADAR", "DHARAMPURA", "JAHANGIRPUR", "RAFILUL NAGAR URF RAWALI", "SHADIPUR DALLA", "AGRI", "BARKATPUR", "MAUZAMPUR SUJAN", "SIRDHANE BANGAR", "FATEPUR KALA", "SIMLA KALA", "MUKEEMPUR BALU", "BHOGPUR PATTI HARSUKK", "BAZAR KALA", "SIHORA NIZAMABAD", "BATAPURA", "BEGAMPUR", "PAHADPUR CHANDRASAIN", "MUBARKPUR TALAN", "CHAHSHEERI B B-21", "CIVIL LINE FIRST", "SOTIYAN", "JHANDAPUR", "AURANGPUR HAZI", "QAZIPARA  JANUBI  B", "KUTUBPUR GARHI", "CIVIL LINE  SECOND", "AURANG TARA", "HUSAINPUR", "MAHAGALI", "SAIFPUR KHADAR", "JAMALPUR PATHANI", "MUKEEMPUR DHARMASI", "MAHMOODPUR", "RAMLELA PURV", "SHAHWILLYAT", "GAJRAULA ACHPAL", "RAMPUR THAKRA", "KARIMPUR BAMNAULI", "KALPUR", "JAMALPUR", "KUWAR BAL GOVIND", "CHAHSHEERI B B-24", "AFGANAN", "BARKATPUR GARHI", "MOMINPUR DARGO", "JAHANGIRPUR LALU", "SARIYAPUR", "ADAMPUR", "KHEDA", "RAMPUR BANGAR", "HASANPUR KAZI", "KHANJAHANPUR BHADAR", "FARIDPUR BHARTA", "RANIPUR", "SIDHARUWALA", "AMIPUR URF DHARMNAGRI", "JALALPUR BUCHA", "JAHENABAD", "FATEHPUR CHAK", "DHOMANPUR", "KADARPUR", "RASOOLPUR PRITHI", "ISLAMPUR DAS", "TAIIBPUR QAZI", "SHUKKHAPUR MARKANDEYAMPURAM", "SHIKANDARPUR LALMAN", "PERZADAGAN", "SHUMAL KHEDI", "KINAN URF MADI", "KHANA PATTI", "DEWAPUR", "MADSUDANPUR BHADHAR", "GANGA KHEDI", "MUKARPUR KHEMA", "CHAHSHEERI B B-22", "BAZAR MANGAL", "HIRDERAMPUR", "NIZAMPUR KHORA B A", "MAHESHWARI", "MIRZAPUR MAHESH", "ALIPUR NAGLA", "DAULATPUR", "KARAUHIDI", "PURUSHOTTAMPUR", "BADSHAHPUR TARIKAM", "FIROZPUR NAROTAM", "HUSAINPUR SANKAT SINGH", "KHANPUR MADHO", "ISLAMPUR DEEPA", "SHAHJADPUR", "KISHANWAS", "SWAHEDI BUJURG", "SALEMPUR MATHNA URF  PURANPUR", "MOLHARPUR", "JAHANABAD", "AJIJPURA", "PAPAWAR KHURD", "INCHAWALA", "SULTANPUR BANGAR", "RAMPUR NOUBAD", "MUKEEMPUR JAMAL URF INAMPUR AN", "MITHANPUR SHOBHARAM", "ITAWA", "HASHAMPUR", "QAZIWALA", "NIZAMPUR KHORA  A", "CHAUKPURI", "FARSHTOLI", "KHARI", "JATAN B-4", "JAISINGH PUR", "KHUDAHERI", "MUKARANDPUR", "NILOHA", "KADAPUR", "GAIBLIPUR", "MUDALA", "SHAHPUR", "JHALRI", "CHANDPUR NAUBAD", "LADAPURA", "RAMPUR BAKALI", "FATEHPUR NOUABAD", "NAI BASTI B-15", "FARIDPUR UDDA", "RAIPUR PIYAGI", "BHARAIRA", "PATWARIYAN", "CHANDANPURA", "RAMJIWALA", "BANJARAN", "ABU SAIDPUR", "FAZALPUR", "ALIPUR MAKHAN", "KAMBHAUR", "DEDANAGALA NAGALA", "FARIDPUR BHOGAN", "MUBARAKPUR URF GADAI", "JATAN B-3", "HUKAMPUR MACHIKI", "SHAHBAJPUR", "GOGALPUR KASAM", "QAZIPARA  SHUMALI", "TATARPUR", "HUSAINPUR TAPPA NANGAL", "KASSABAN", "RAGHUNATHPUR", "JATAN B-2", "MOHAMADPUR LAKHU", "ISLAMPUR LALA", "BHOGPUR PATTI ABDULLA", "GAJIPUR", "FAREEDPUR BHOGI", "RASHIDPUR GARHI", "NAI BASTI B-23", "TITARWALA", "KARIMNAGAR URF ULEDA", "CHAHSHEERI B B B-21", "BAHADARPUR BUJURG", "GOPALPUR", "KISHANPUR", "MOHAMADPUR DEOMAL", "MIRDGAN B-10", "MUQAITPUR", "MEERAPUR KHADAR", "NOORALAPUR HAFIZ", "YUSUFPUR HAMID", "FARIDPUR QAZI AN", "KAMALPUR", "MUZAFFERPUR KESHO", "RASOOLPUR", "FATEPUR SABHACHAND", "NOORPUR DALO", "SADULLAPUR", "PRHTHVIPUR", "KUNDANPUR", "SHUKANANDPUR", "KHATRIYAN B-12", "KACHHPURA", "BAHADARPUR JATT", "MEERAPUR RAZA", "FARIDPUR BHORO", "DAYALWALA", "MAQSUDANPUR HAFIZ", "RAMPUR VEERAN", "KHALIULLAPUR", "RAIPUR BARYSAL", "JHAL", "SEKHPURA", "KABULPUR", "BASIRUHASI", "PAMARGANG", "MANGOLPURA", "AURANGPUR BIBI", "ALAMPUR NILA", "MOHSANPUR CHAMRA", "SADAT", "AMIPUR URF NARYANPUR", "MANWALA", "ABUL KHAIRPUR BANGAR", "RAFATPUR FATEHABAD", "BHOJPUR BHOPATPUR", "KHANPUR DULLI", "KOHARPUR", "CHHACHHARI TEEP", "SEWAARAMPUR SAKAT", "NIJAMATPURA GANJ", "SHAHWILLAYAT", "RAGHAVRAMPUR  SARNGPUR", "GAVRI BUJURG", "MOLVIYAN", "SHIKANDRABAD", "SHAWILATE", "WAZIDPUR", "MIRDGAN B-11", "CHAHSHEERI B-24", "BHAWANIPUR", "GANGADASPUR", "TIVARI URF KADARPUR JASWANT", "GHASIWALA COLLONY", "DWARIKAPURI", "MADSUDANPUR NAND URF JHALRA", "REHADWA", "DHARAMNAGARI COLONY", "KANONGOYAN", "KSSABAN")
		case "Chandpur":
			sections = append(sections, "KIRATPUR", "SALHAPUR", "MERJAPUR KHADAR", "KAFOORPUR", "MUSTAFABAD URF GHDANPURA", "BHAISA", "BAGAMPUR", "RAVANA", "MUZAFFARPUR KHADER", "PIPLIJAT", "NAWADA", "RAMPUR NAJRANA", "AZAMPUR URF ADDOPUR A", "AKBARPUR TIGRI  A.", "SHAPUR BASODHI A.", "DHOLANPUR", "KHANPUR RAYDASS G.A", "SHAREEFPUR", "SAIDPUR", "MUBARAEKPUR A", "MUBARAEKPUR KALA", "MURAHAT A.", "MARUFPUR", "BABANPURA", "MEERAPUR SEEKRI", "KUHUTUBPUR GAWRI", "CHAHSANG A.", "REHRA A.", "MAHU", "HALLANAGLA", "KHANPUR KHADER A.", "YUSUFA", "MUBARAEKPUR A.", "KAZEESORA", "HARPUR", "BALIPUR GAWRI", "BEBIPURA", "SINGH A.", "RASULPUR PUNA", "KAJISARAY A.", "SHADIPUR", "LATEEFPUR CHUKHRI", "SHRAIRAFI  A.", "SEMLA", "MADAMAHDUD", "SINGHA", "BASANTPUR", "PEETANONDHA", "RAIPUR KHADER", "SIAU A.", "SULTANPUR KHADER", "BAMNOLA", "MAHU A.", "BARKHADA", "RATANPUR KUHRD", "HAZARPUR KASWA", "RAJPUR PARSU", "PADLA", "PRATHBIPUR", "DARWARA A.", "CHAHSANG", "BHAGOORA", "AKONDA A.", "MERJAPUR BALA A.", "JAHGEERPUR M.GAYNPUR URF ROOPPUR", "ATHI", "BHAWANIPUR GADDO", "TIRPUDI", "LAKMIDHAR", "BHARKARA URF AZZAMNAGAR", "DARWAR A.", "AZAMPUR URF ADDOPUR A.", "SHRAYRAFI  A.", "OLIYAPUR", "HUSANPUR KHASA PURV", "AKONDA", "KHANDSHAL", "JOINPUR", "RUSTAMPUR", "PATEYPARA A.", "RAVTI", "MAHMOODPUR", "AKLASPUR", "DHARUPUR A.", "GHANSURPUR", "MATLABPUR", "SEDPURA URF NAIPURA", "SUNGHAR", "MO.PUR KHADER", "ASALATPUR URF TOFAPUR", "JAFARPUR COT", "KALANPUR", "MANZOLA GURJJAR A.", "SISONA A.", "JAFARABAD KURAI", "MUFTISARAY A", "SAINDWAR", "AJDEV", "HIRNAKHERI", "AJIMULLHANAGAR", "CHEEMIN A.", "SHADIPUR HUSAINPUR", "MUBARAEKPUR POTA", "ISMAILPUR", "ROONIYA", "KUNDA", "KAKRALA", "DHARMUPOOTA", "THARPUR MADAD IMAM", "BAHADARPUR", "KAYASTAN A.", "PATYAPARA A.", "FOLADPUR", "ALABALPUR", "KULCHANA", "KHANPUR MAJRA GANSURPUR", "MUSTAFABAD", "MANPUR", "PAHULI", "BAMNOOLI", "UMRIPEER A.", "SIRAKTHAL", "HUSANPUR KHASA P.", "KHARPUR JAGEER", "HASAMPUR", "TANDA", "IMALIA", "GANDHOR A.", "SHAHUVAN A", "KOSHALY", "CHAMROOLA", "JHUJHAPURA URF  NAIPURA", "LINDERPUR A.", "BAHTOLA", "NASARPUR  MANSURPUR", "BAGARPUR", "MOHANPUR", "FATHULLAHPUR POTTA", "THURALA", "SAINDWAR A", "RAVANA DARGU", "CHIMMAN  A.", "PAVTI A.", "SICKRONDHA", "TAKANPURI", "HELALPUR", "BHAVANIPUR GANGRAIN", "HEEMPUR BUJURG A.", "BASODA", "SABDALPUR A", "SHAPUR BASODI A.", "SHEKPURI CHOHARD", "KARAMPUR URF GAERI", "NAJARPUR", "BASTA A.", "MANOTA", "KAYSTAN A.", "LANGPURI", "REHRA A", "SABDALPUR KHURDH", "WAZEEDPUR", "SOHARABNAGAR", "JATOOLA", "MALAKPUR", "SAKTALPUR MILAK", "ARAGI BHASA", "FAZIPUR KADER", "SHUJATPUR KHADER", "SEEMLI", "TAKHATPUR", "MAHBULAPURDHAKI", "SARAY SHAK HABEB", "RUKANPUR A.", "JOGIONDA", "BADIWALA", "PEPALSANA", "SARAY SIAU", "KHADA G.B", "AJUPURA JAT", "ALIPURA", "ISMAYILPUR A", "JAMALLUDDINPUR", "SANIPURA", "RAMPUR KHADER", "BEERAMPUR", "AHROLA", "SAHACHANDAN A.", "SHAHPUR DHAMDI", "JAGANPUR URF JAGANBALA", "ZUZHAILA A.", "PIPLIJAT  A", "PHEENA A.", "KARANPUR", "P.V. PAHADPURKALA URF MALASJYA", "BAAJAR", "CHAKORANGABAD", "BUCHNAGAL", "AJJUNAGLI", "ORAINGPUR JOGIPURA URF NORANGPUR", "MAHMOODPUR KASWA", "ORANGABAD", "DUARIKAPURI", "MILAK SAKTALPUR", "PAHADPUR KHURD", "ISMAYILPUR", "BAHMANSORA", "KAMNA UMRA G.A", "FAJALPUR POUTA", "UMRIPEER", "KOTLA A.", "LADUPURA", "NALPURA", "HESAMPUR", "ABIDPUR", "GANGU NAGLA A.", "LATRA", "AZAMGAR URF RATANGAR A.", "SANSARPUR", "MUKARPUR", "EBRAHIMPUR", "HEEMPUR DEEPA", "CHAKRUSTAMPUR", "DHUNDLI", "BOHRA", "DHAKI", "MEERAPUR", "BUNDRA KHURD", "TAHARPUR GULAM EMAMEN", "AKBARPUR URF CHOTAPUR", "JAINULAUDDINPUR", "CHAHLI", "MANZOLA GUJJRA A.", "NASEERPUR SHEK", "MEHMUDA KHADER", "SUDNIPUR", "KAMALPUR A.", "MERJAPUR BALA", "SIHALI", "RAOPUR BAHAMAN A.", "DHEVERPURA A.", "RASOLPUR NANGLA A.", "AKBARPUR", "BABARPUR", "SOOTKHADI", "SUNDRA", "RAVANA HABAT", "MASEET A.", "SHARAY", "TRAPUDA", "BHOOGPUR", "NAWABPURA URF SHAKUPURA", "CHAUNDHRI", "KUTQUIE", "PILLANA", "HATAMPUR", "CHAMANPURA", "TABEBPUR", "DATTYANA A.", "SABDALPUR A.", "AKBARPUR  ZOZA", "GOVLI A.", "THAT", "MUDHAL", "AKBARPUR MAJRA SAINDVAR", "JALALPUR KHADAR", "SIKANDER NAGLA", "SHAKHPURI MEENA", "COT", "KAZIJADGAAN A.", "KAMALUDDINPUR HUSANPUR", "BHIDEYKHERA", "BEBIPURA A.", "KAMALPUR", "JAFARHUSENPUR", "MEERZAPUR BAKENA", "BASADI", "KHERKI A.", "RAMPUR PUNA", "TUNGRI", "RUKANPUR", "FATHPUR", "NANDNOR", "RAMOROOPPUR", "FAZIPUR KASWA", "BASADA", "LUHARPURA", "BALIYANAGLI", "DARWAR", "MITTANPUR", "TALIBPUR", "RAOPUR BAHAMAN A", "BAGWANTPUR", "BIRAL", "KATARMAL A.", "TANGROLA", "KARAL", "BUNDRA KALA", "JALALPUR A.", "SALAMPUR A.", "SADABAD", "MOHD. HUSANPUR MAJRA ORANGABAD", "HUSANPUR KALA")
		case "Noorpur":
			sections = append(sections, "MIRJAPUR DHIKLI,", "RAWANA SHIKARPUR", "SHARFUDDINPUR KAIMRY", "DAULATPUR BILLOCH", "BUDA NANGLA", "ROSHANPUR JAGEER A.", "AULIYAPUR", "GOVINDNAGAR MUHAMMADNAGAR", "JAMALPUR KIRAT A.", "MANGOLPURA", "MUSSEPUR", "VIJAI NANGLA", "RAHU NANGLI A.", "HAKUMATPUR", "JAMALUDDINPUR G.A.", "ATHAI JAMRUDDIN H.N. 1 TO 120", "ABDULLANAGAR URF DANDA", "BER", "CHAKCHAND G.A.", "NYAMTABAD", "SULTANPUR", "MEWA NAWADA A.", "KANHEDI", "JAMALPUR KIRAT", "SELA A.", "JAHIRABAD G.A.", "BAMNAULA,", "KASAMPUR BILLOCH", "SELA KHEDI", "SHAYAMABAD", "MAGHPUR UDAYCHAND", "MORNA A.", "JUJHELA", "HASUPURA HARKISHANPUR A. H.N. 61 TO 243", "MADHUPURA", "SIDIYAWALI", "MOLVIYAN A.", "HUSENPUR KALA", "SURANANGLA", "RAHTA BILLOCH", "NAWADA RAWANA URF TELIPURA", "HEMPUR", "MAHAMDABAD A.", "BHAGWANPUR RAINI", "SHIWALA A.M.N. 1 TO 124", "SAIDPUR MAFI", "GOPALPUR", "GAJIPUR", "RAMKHEDA", "GOHAWER JAIT A0", "ASGARIPUR A0", "CHEHLA A.", "ISLAMNAGAR A.M.N. 1", "PITTHAPUR A.", "ROSHANPUR", "ALAUDDINPUR", "SELPURA", "KASAMPUR DHANNA PUR MAFI", "KABIRNAGAR A.", "KOTRA MOLVIYAN", "SHAIKHAN UTTARI AN.", "CHAJJUPURA", "AFGANAN UTTARI", "SEH", "BHAWANIPUR KALIYA", "REHTOLI", "ABDULLPUR DAHANA", "RASULPUR KASBA A.", "PANIYALA A.", "JALALPUR DIPA", "NOORPUR DEHAT", "FATEHABAD", "KHEDI", "BAHLOLPUR", "MANDORA", "SHAIKHAN UTTARI", "ALIYARPUR", "IBRAHEEMPUR,", "BURHANPUR", "KHUKUNANGLA", "CHANGIPUR", "MAHDUDNASHO", "AFGANAN PURVI A.", "SAKLI", "SADAFAL A.", "BIRBALPUR", "ISLAMNAGAR A.M.N. 1 TO 119", "MORLA", "CHODHRIYAN", "BASANTPUR MAFI", "BAGWADA", "GANDHINAGAR", "DOSTPUR GAIR A.", "PITTHUPURA", "HAIGERPUR BHATT", "HAMA NANGLI", "JAINABAD", "AKBERPUR ASHA URF HAROLI", "HASUPURA HARKISHANPUR A. H.N. 244 TO END", "NAYAK NANGLA", "SHAIKHAN TARAI", "HALDUA MAFI", "HAKEEPURA DAKSHIRIN", "SUMALKHEDI", "JAMALPUR MAN", "JABERDASTPUR G.A.", "GOVINDPUR", "SHOLAN SHEKH", "ALIPURA", "DOLTABAD", "RAMNAGAR SHAHIDNAGAR", "MORMAKDUMPUR", "BANJARAN", "BUDANPUR A.", "CHAKJAMAL G.A.", "FATEHPUR BULANDI", "MADARIPUR KAKRALA", "MUNDA KHEDI", "ASGARIPUR A0 H.S. NO. 1 TO 133", "LAMBIYA", "PURANPUR NANGLA", "RAMNAGAR A0", "SATVAI", "KASAMPUR VIRU KHALSA", "SHAHPUR KHEDI", "ADAMPUR KAKWADA", "AFGANAN PASHCHIM A.", "DHAILA AHIR", "BALLA NANGLA", "GUNIYAKHEDI", "DHAMROLA", "MAJHOLA BILLOCH A0", "MEERAPUR", "FAJALPUR DHAKI A. H.N. 1 TO 114", "SHAHNAZARPUR KOTT", "ALADINPUR A.", "SATTO NANGLI A.", "HIMPUR PIRTHIYA", "POTA", "MUJAHIDPUR", "MUBARAKPUR NAWADA", "SALEMPUR JHILLA", "ANESA NANGLI", "GOHAWER HALLU A0", "TAJPUR A. H.S. NO. 1 TO 796", "ISLAMNAGAR A.", "KHANPUR BILLOCH", "RAMPUR", "SARAEKALA", "AFJALPUR BALDANI", "BHUDHPUR,", "JHIRAN", "SHAHBAZPUR", "PUWENA", "SHUJATPUR TIKAR A.", "RANANAGLA", "MAIHAWATPUR HEMRAJ", "RASOOLPUR DHAKI", "PURANPUR SAYANA", "ALAMPURI", "MEWLA MAFI", "BAGWADA A.", "NANGLI JAJU", "HASANPUR BILLOCH", "GOVEDHANPUR", "IDALPUR URF JOGI PURA", "RASHEEPUR", "RAMNAGAR A.", "MALWA", "MUKSUDPUR A.", "KAZIPURA", "RAMPUR KISHNA", "SANTO NANGLI A.", "KIWAD A.", "SARKADI", "BRAHAMPURI", "NAMDARPUR CHANDAN", "UMRI BADI", "DHINGERPUR", "KHASPURA A.", "TAJPUR A.", "PANDIYA", "BERKHEDA", "ALAUDDINPUR DHANRAJ", "NICHALPUR", "SABDALPURI G.A.", "TAHAERPUR", "RAWANA SHIKARPUR A.", "MORNA A0", "JALALPUR URF SHAIKPURA", "ALAMPUR", "PAINDAPUR", "PIPLA JAGIR A.", "JAIRAMPUR", "ISLAMNAGAR A.M.N. 82 TO END", "MAKANPUR", "SHERPUR MAJRA", "AFGANAN PASHCHMI  A.", "NOGRA", "AFGANAN UTTRI A.", "BISHANPURA URF BIDRA A0", "ABIDNAGAR URF DHUDLI", "SHIWALA A.M.N.125 TO END", "GANDA JUD", "TAJPUR A. H.S. NO. 797 TAK.", "MOHD. TAKIPUR GHASI", "MALAKPUR", "BAJHRABAD URF KHATAI", "MAKSUDPUR A.", "MAHAMDABAD", "KAJAMPUR", "ALAMPUR ADWA", "HASUPURI G.A.", "MAHMUDPUR A.", "BUDPUR NAWADA", "LUFTIPUR,", "KHASPURA A,.", "MURTJANAGAR URF GANGADHAR", "DHOLAGAD", "KHAWAJA AHMADPUR JALAL URF PAITIYA", "NAINU NANGAL", "ATHAI JAMRUDDIN H.N. 121 TO END", "PURENA ABDUL REHMANPUR A.", "BEDA,", "NASIRPUR MANSUKH", "FAJALPUR DHAKI A. H.N. 115 TO END", "MITHAI", "CHOUDHRIYAN", "KHAI KHEDA", "RUPPUR", "JAMALPUR JATT", "SHAPUR HARRAE", "MUKAM", "LAMBAKHERA A.", "CHELLAPUR", "TANGROLI", "DARIYAPUR A0", "MAHUPURA", "SADHARANPUR", "BASEDA KOTT", "ISLAMNAGAR A.M.N. 120 TO END", "GOSPUR KHALSA", "RAHPANPUR", "HAJRATNAGAR A.", "HASUPURA HARKISHANPUR A. H.N. 1 TO 220", "RASULPUR G.A.", "MUJAHIPUR", "POTI", "RAVIDASNAGAR", "AASRA KHERA", "UMARPUR", "SHAIKHAN DAKSHIRAIN", "NAZARPUR BILLOCH", "SAHADPUR GULAL", "MEWA JATT A.", "TENDERA", "JATT NANGLA", "KASHMIRPUR GADI", "GOVINDNAGAR", "SHEKHAN UTTRI A.", "DARIYAPUR,", "LAKHIPURA", "DHAILY AHIR", "AMINABAD", "DAWAL KHEDI", "RATANPURA", "IBRAHIMABAD", "KASAMPUR MANGAL KHEDA", "SHIWALA A.", "MUKARPUR AHIR.", "HASUPURA HARKISHANPUR A. H.N. 221 TO END", "PURENI DURWESHPUR", "RAEPUR BIJJU", "AMANATPUR", "HAKEEMPURA", "KOLASAGAR", "PATHWARI", "CHAKRAMAIYA G.A.", "SULTANPUR MEV", "LODIPUR MILAK", "YAKUBPUR BILLOCH", "KHAJURI", "MOHD. ALIPUR INAYAT A.", "KASMABAD", "MAJHOLI", "BHOGPUR", "CHACK MEHDUD SANI", "KUNDA KHURD", "MUBARAKPUR MADHU URF GADI", "BASANTPUR MAFI G.A.", "BISHANPURA", "HASUPURA HARKISHANPUR", "AKBARABAD G.A.", "MADHUSUDANPUR URF JALALPUR", "ASGARIPUR A0 H.S. NO.134 TO END", "UMRI BUJURG", "RAMPUR VIDAR", "KURI BANGER", "KHANPUR", "ASADPUR DHAMROLI", "PAIJANIYA A.", "PAIJANIYA A. H.N. 81 TO END", "REHTI JAGIR", "ATHAI AHIR", "KAMALA", "SADANPUR", "DERABULANDI", "HIRANPURA", "FAIZPUR", "SHAHPUR BADWA")
		case "Kanth":
			sections = append(sections, "MO. TEKEDARAN OMRI KLA", "SARUA DHARAMPUR", "PILAKPUR SHYORAM", "MEHDOOD KALMI", "PRITHABI GANJ PASHMI", "GOVINDPUR", "VISHANPURA", "MALEHPUR KHAIYYA", "DULHEPUR NIKAT CHANGERI", "RAMSARAY", "MEHMOODPUR MAFI", "HASANGHADI", "CHOOKBAJAR UTTRI", "RAYPUR KHURD ORF SYOHARA", "MUSTAFAPUR DITIYA", "GHYANPUR", "MUJJAFFARPUR TANDA", "BADODE FATEEHULLAPUR", "BERAMPUR", "CHAJLET", "KUMHARIYA JUBLA", "KASAMPUR NIKAT KANTH", "JAHAGEERPUR CHEKFERI", "NAKSANDA BAD", "GHAKKARPUR", "SULTANPUR NIKAT ESAPUR", "MILAK AMAWATI", "AHMADPUR ANANDPUR", "GHADI", "SIDARUNAJARPUR", "SADAKPUR MALOOK ORF KHICDHI", "BAJIDPUR", "MUNDALA MOHAMMAD JMAPUR", "SALEMSARAY", "SAHUPUR", "GADHI", "DIDORA", "AADAMPUR", "SIKANDARPUR", "GAJGOLA NANALBADI", "KHANPUR MUJAFFARPUR", "RUSTAMPUR", "GULADIYA MAFI", "PAKBDA", "AABI HAFEEJPUR", "BASABANPUR", "FALJABAD", "MANKUA MAKSOODPUR", "KAJIKHEDA", "MANGUPURA", "BAKENIYA", "MO. AVIDPURA OMRI KLA", "PAHADMAU", "SADUPURA", "KAJIPURA", "SADARPUR MATLABPUR", "GUTABLI MUO", "MUDIYA MOHINUDDINPUR", "SHERPUR ETMADPUR", "BHIKANPUR ASDALPUR", "DODRAJPUR", "PRATHVI GANJ DAKSHNI", "CHAJJANANGLA", "KUAKHEDA MAFI", "RAMNANGLA", "DIDORI", "NANGLA JOGRAJ", "KURI RAWANA AANSHIK", "MADHUBA KHALSA", "BARIPUR", "SIKANDRABAD", "SIRSA THAT", "SHAHPUR ABDULBARI", "PRATHVI GANJ PURBI", "SAMADPUR BUNIYADPUR", "SHAHALAMPUR", "ISMAILPUR", "SIHALI KHADDAR", "CHANDKHEDI", "RANA NANGLA", "MO. CHODHRIYAN OMRI KLA", "MOHDI HAJRATPUR", "FAKEERGANJ PASHMI", "BEGAMPUR", "PATTIBALA PASHMI", "AKBARPUR CHEDRY", "MIRJAPUR KARIMUDDIN", "MEHLAKPUR NIJAMPUR", "DEHRA NIKAT OMRI KLA", "NEMTULLA NAGAR ORF MITHANPUR", "JEBDA", "AHMADPUR NINGU NANGLA", "CHATKALI", "GAKKHARPUR", "SADARPUR", "MANPUR MUJAFFARPUR", "RAMNAGAR ORF RAMPURA", "AGWANNPUR", "ALIPURA MAJRA KUI", "NAJRANA", "NEHTORA", "SAMATHAL", "ALIPURA KHALSA", "GAJROLA SED", "BARKHEDA BASANT PUR ORF DYANATHPUR", "KHADANA", "MANOHARPUR", "RASOOLPUR GUJAR DHYAN SINGH", "SALABA", "LADAWI", "PAIGAMBARPUR SUKHBASI LAL", "DARIYAPUR", "AGWANPUR", "SALEMPUR KHAS", "MISRIPUR", "HARINOORPUR", "PATEGANJ", "SALEMPUR BANGAR", "GIDODA", "MO. FAREEDGANJ OMRI KLA", "KAJIPURA KHALSA", "NASEERPUR", "BHESLI JAMALPUR", "LODHIPUR RAJPOOT", "MANPUR SABIT", "RASOOLPUR SUNBARI", "GOLAPANDEY ORF SANDLIPUR", "CHAJJUPURA DOYAM", "AKBARPUR SIHALI", "PATTIBALA UTTARI", "GHOOSIPURA", "MAGHPURI ORF INAYATPUR", "KHANKHANAPUR ORF BICHPURI", "MADHUPURI", "MUSTAFFAPUR ABBAL", "DEHRI JUMMAN", "BAKARPUR ATYAN", "SHAHPUR MUBARAKPUR", "SIKRI", "MO. BIDATTSHAH OMRI KLA", "UTTAMPUR BEHLOLPUR", "MANSOORPUR", "ANYARI ORF ALINAGAR", "KUDAMEERPUR", "DHANUPURA", "DUDELA", "KIRATPURI", "GOPALPUR NATHTHA NANGLA ORF KOKARPUR", "PHOOLPUR MITHANPUR", "KURIRAWANA AANSHIK", "BUDHANPUR KHALSA", "MEBLA", "SANJARPUL SULTANPUR", "AJDHAPUR", "SEHDOLY", "BHUDHANPUR ORF BILAYATNAGAR", "MISHRIPUR", "FOLDA PATTI", "CHEDRY AKBARPUR", "ADHAMPUR", "NANGLA BANVEER", "GHOOSIPURA PURBI", "PALLUPURA GHOSI", "DHADI MEHMOODPUR", "KHALILPUR KADDIM", "MO. ANSARIYAN OMRI KLA", "EDELPUR JOGIYA SIHALI", "MO. NASERABAD OMRI KLA", "SARAY KHAJOOR", "MAHENDRI SIKANDARPUR", "PATTIBALA DASHNI", "HASAMPUR GOPAL", "CHODHARPUR", "DUMHELA KHALSA", "KHATAPUR", "KURI RAWANA", "JAMNA AZAM", "DHARAK NANGLA", "LADAWLI", "KISHANPUR", "MAJHOLI", "MOHAMMADPUR DHYAN SINGH", "FAKEERGANJ PURBI", "MEHLAKPUR MAFI", "SEEMLA", "MUKHTYAR PUR NAWADA", "KHLILPUR AMRU", "PATHANGI", "BAGIYA SAGAR", "DHAKI", "BAHADARPUR KHADDAR", "PELI VISHNOI", "GOPALPUR NATHATHA NANGLA ORF KOKARPUR", "PATTI MODHA", "POTA", "GINNOR DEHMAFI", "FATTEHPUR VISHNOI", "JAMANIYA KHURD", "NAJARPUR", "MO. PACHEYAN ORMI KLA", "CHANGERI AANSHIK", "RASOOLPUR CHOHRA", "BADHERA", "HIMAYUPUR", "GURETHA", "BIBIPUR", "MEHMOODPURA PACHMI", "KARANPUR HARKISHANPUR", "JAMALPUR MUNDA NANGLA", "MODHA", "MATHANA", "FAREEDPUR BHENDI", "BAGADPUR", "PACHOKRA KHANPUR", "MO. NANGLIYAN OMRI KLA", "FAJALPUR", "BAHEDI BRAHAMNAN", "NIDHI", "UMRI", "MAHESHPUR")
		case "Thakurdwara":
			sections = append(sections, "RAMANAVALA", "ADAL PUR", "VALAPUR", "DAULAVALA", "RAMAPUR GHOGHAR", "SARKADAKAREEM", "KARNAPUR", "MASUMAPUR", "PADIYA NAGLA", "RATUPURA 2", "HONSPURA", "SHAHABAJAPUR KALAN", "TANDA AALAM", "FAREEDPUR KASAM", "BAHADUR NAGAR", "JAHIDAPUR SIKMAPUR", "MEHMOODPUR KESHO", "JAMNAVALA ANSIK", "CHANDAKHEDI", "MIRAPUR MOHAN CHAK J0 MU0", "VATHUAKHEDA MU.", "SAHASAPURI", "MANPUR DATTRAM MU.", "JALALPUR KHALSA", "MADARPUR", "MAHMOODPUR LAL", "BITHUATHER", "DILARI CHAGERI", "BAHADURPUR PATTI", "PILKAPUR GUMANI", "BHANAPUR GAJAROLA", "BHARATAVALA", "MASTALLIPUR", "SAUDASPUR", "PAVARTAPUR MAHANAND", "TUMADIYAKALA", "SARKADAPRAMAPUR MAFI", "DHAKIYAPIRU", "SULTANAPUR KHADDAR MUS", "ALAMPUR", "AMANATABAD", "RAJAPUR KALA", "UGHO PURA", "ALLEHPUR", "FATEHAULLA GANJ MANIHARAN", "TAH MADAN", "BHAYAPUR", "FAIJULLAGANJ", "VARAHILALPUR MUST0", "TAGALI", "TAGALA", "NAGALIYA NARAYAN", "SUNDAR NAGAR", "MADH PURI", "BAHADURGAJJU", "FATEHULLAGANJ ANSARIYAN", "KACHANAL", "BAIJANATHAPUR", "BABUPURA GHOSIPURA", "SHERAPUR BAHALIN", "MUNDIYAKHRDA", "FATAN PUR", "PANUVALA", "KARANAVALA KHALSA", "GULADIYA MURAD", "JUL DHAKIYA", "KALAVALA", "VIRAPUR FATEHAULLAPUR", "LALPUR CHAUHAN", "SARKADA KHAS", "MASOOMPUR", "KUNDESRA", "CHATARPUR NAYAK", "NAGLA TAHAR", "MANGA VALA", "ADONGALI", "DULHAPUR PATTI CHAUHAN", "PASIYAPURA PADARATH", "FAULADPUR", "FATEHAULLAGANJ BANJARAN", "ALIYABAD", "BAHADUR GANJ", "FATEHAULLAGANJ DHOBIYAN", "KARANAVALA JABTI", "RATUPURA", "PIPAL GANV", "LALAVALA", "KHAIRULLA PUR", "PANDITAPUR", "UDAVALA", "IRSAPUR", "RUPAPUR TANDOLA", "DULHAPUR PATTI JAT", "AKBARAPUR", "SAHABGANJ", "KAMALAPURI KHALSA", "FONDAPUR", "MALHAPUR LAKSHMIPUR", "NRAYANPUR CHANGA", "TUMADIYAKALAN", "JOGIPURA", "KURI", "GADUBALA", "DHINGARPUR", "THAKURADVARA TALI", "NABADA", "PADANLA", "KANKARKHERA", "LAKHNAPUR LADAPUR", "DARAPUR", "CHAUPURA", "USAMANPUR", "SULATANPUR DOST", "INAYATNAGAR", "RAMNAGAR KHAGGUVALA", "GAJARAULA JAY SINGH", "ASDULLAPUR", "NIRMALAPUR", "BANKAVALA", "CHANDAPUR", "KISHNAPUR GAVDI", "DHAKIYAJAT", "FATEHAULLAGANJ NAIBASTI", "FAREEDNAGAR", "NANHUVALA MU.", "MUHIUDDIN PUR", "LALPUR PUROHIT", "HANSUPURA", "MUNIMAPUR", "DAULAPURI BAMANIYA", "BHAVANIPUR", "BHATGAVAN", "ASALEMAPUR", "BUDHANPUR MU.", "RAJUPUR MILAK", "LONGIKALAN", "BAIRMAPUR", "SAIJNI MUSTAHAKAM", "JAMASHEDAPUR", "SATTI KHEDA", "SHARIF NAGAR", "ALIYABAD MUU0 SABIK", "BAMANIYAPATTI", "THAKURADVARA ARYANAGAR", "IGRAH", "RAJPUR KESARIYAVALA", "ABDULLAPUR LEDA", "RANI NAGAL", "RAMUVALA SHEKHU", "ALMAGIRAPUR", "RAMNAGAR KHAGUVALA", "NARENDRAPUR", "CHAKLOHARA", "HARIRAM ROAD VARD-13", "RADHUVALA", "KHAI KHEDA", "NURAPUR", "FAIJULLANGAR", "FATEHULLAGANJ JATVAN MANDIR", "MANAVALA", "PRITHVIPUR GABDI", "JAHANGIRPUR", "FATEHAULLAGANJ SAIFIYAN", "JATPURA", "FAREEDPUR HAJI", "PIPLI GHANSYAM", "FATEHAULLAGANJ CHALACHITRA ROAD", "SHERAPUR PATTI", "SABALPUR", "BUDHNAPUR", "ILAR", "BAHEDI", "RAIHATAMAFI", "BOVAD VALA", "KARIYANAGLA MUST.", "KAMALAPUR", "CHANDPUR", "BHUD", "KALYANAPUR", "HAJI NAGLA", "JAFARABAD", "PIPLI AHIR MU.", "TANDA AFJAL MU", "LONGI KHURD", "LOGI KALA", "MALKAPUR SEMLI", "RAMAPUR BALBHADRA  JAHEED", "SURJAN NAGAR", "SALAR PUR", "VIJAY RAMAPUR", "GOPIVALA", "TATARPUR", "KOTLA NAGLA", "KHVAJAPUR DHANTALA", "PAINDA PUR", "FATEHAULLAGANJ KURESHIYAN", "SAKAT PUR", "KALA JHANDA", "PIPALIUMRAPUR", "KUKURJHUNDI", "ROSHANPUR", "MOHIUDDINPUR KHALSA", "VAHAPUR", "SHIVNAGAR", "NARAYNAPUR TEJU", "SULTANAPUR MUNDA", "GAJHEDA ALAM", "TARFADLAPAT", "VILAVALA", "RAGHUNATHPUR", "BHAGTAPUR TANDA", "BUDDHANGAR", "NAKHUNKA", "SHUMALKHEDA", "KUNA KHEDA KHALSA", "UDMA VALA", "BHAGIYAVALA", "SIDHAVLI", "FATEHULLAGANJ TAHIRA ROD", "SULTANPUR MUNDA", "LALAPUR PIPLASANA", "LADHUPURA", "BAHAPUR")
		case "Moradabad Rural":
			sections = append(sections, "BAGIAWALA", "TAMBAKUWALAN ANSHIK", "CIVIL LINES PURVI", "RAMGANGA COLONY", "MEDNIPUR", "PARWATPUR DAMBU", "MUFTI TOLA ANSHIK", "SHUKLAN", "FATTEHPURI ANSHIK", "TITUS SCHOOL", "THATHERA MOHALLA", "TRILOKPUR", "IMAMBARA", "MOHALLA DANG", "BERKHEDA CHAK", "MUGALPURA PRATHAM ANSHIK", "BAGH GULABRAI UTTARI", "KAFIABAD", "BAZAR MUFTI ANSHIK", "NAYA GAON", "MANIHARAN ANSHIK", "MADARSA ZYARAT SHAH BULAKI ANSHIK", "HARTHALA ANSHIK", "DEHRIA ANSHIK", "GOONGANAGLA", "NAGPHANI UTTARI ANSHIK", "RAMGANGA BIHAR-2", "JHABBU KA NALA", "BAHORANPUR KALA ANSHIK", "SAIHAL MUSTAHKAM", "THEEKRI ANSHIK", "DUPTY GANJ ANSHIK", "SABJI MANDI HARTHALA", "SAIHAL AITMALI (GAIRA.)", "NAGPHANI DAKSHINI ANSHIK", "MAHAL SARAI", "RAMNAGAR MAJHRA", "KASDI TOLA", "SARAI SHEKH MAHMOOD ANSHIK", "DINDARPURA", "LAKDI WALAN PASHCHIMI", "MILAK HAJIRO BALI", "SUNARO WALA", "GODHI ANSHIK", "DUMDAMA JIGAR COLONY", "CHOWKI HASAN KHAN PURVI", "LALBAGH KHALLE RAMGANGA COLONY", "ASHIYANA-2", "DUNGARPUR", "DADUPUR PAINDAPUR AITMALI (GAIRA)", "MOHALLA SAHU", "RATANPURA", "SARAI KHALSA ANSHIK", "CHAK BEEJNA", "BANGLA GAON ANSHIK", "MOHALLA SAINI BASTI", "BADI MANDI", "DEENDAYAL NAGAR PHASE-1", "CHAKARPUR", "KOHNA MUGALPURA ANSHIK", "SHAHI MASJID", "GORA BAZAR", "RAMGANGA BIHAR ANSHIK", "KISHROL PASHCHIMI ANSHIK", "KHANPUR KASBA (GAIRA.)", "KHAIYA KHADDAR", "HATHAT", "BEEJNA ANSHIK", "LALBAGH KHALLE ANSHIK", "PADLIBAJE", "MILAK SIRASWA HARCHAND", "NAVEEN NAGAR ANSHIK", "NAYEE BASTI PURVI", "AFRASHIYAN GANJ", "AKKA FATTU HAFIJPUR (GAIRA)", "SAMDA CHATURBHUJ", "KISROL PURVI ANSHIK", "NIWADKHAS ANSHIK", "KANOON GOYAN ANSHIK", "BARADARI", "KISRAUL PURVI ANSHIK", "KANKARKHEDA", "RANI NANGAL ANSHIK", "BELWADA", "SIHALI ANSHIK", "SIRASWA GAUD ANSHIK", "BADA SHAH SHAFA ANSHIK", "CHAUDHARIYAN", "PASIYAPURA", "MUKARRABPUR ANSHIK", "PEERGAIB", "RAMGANGA BIHAR COLONY", "MOHABBATPUR BHAGWANTPUR AITMALI (GAIRA)", "PEEPALSANA SHAHI MASJID", "BANWATA GANJ", "KHAWADIAGHAT", "DEVIPURA (ANSHIK)", "RUSTAMPUR TIGRI ANSHIK", "RAMGANGA BIHAR-1", "JALPUR", "GANESHPUR DEVI  AITMALI (GAIRA)", "CIVIL LINE ANSHIK", "CHAURASI GHANTA", "TABELA ANSHIK", "BARU BHUDA", "CIVIL LINES PURVI ANSHIK", "GUJRATI", "MUKARRABPUR MANDI BANS", "KAIRAN", "WAJIDPUR TIGRI", "PAKKA BAGH ANSHIK", "NAVEEN NAGAR", "MILAK MOHABBATPUR", "MUFTI TOLA", "SARAI GULJARIMAL", "RAFATPUR", "JEELAL MANDI BANS", "JHADEWALA", "AVANTIKA COLONY", "DEENDAYAL NAGAR PHASE-2", "MAJHRA-BERKHEDA", "GHASMANDI", "SARAI KISHAN LAL", "RAMGANGA BIHAR", "KASAIBADA", "FAIJGANJ PASHCHIMI", "BAWARIYAN", "NAWABPURA ANSHIK", "MUDIA", "CIVIL LINE HOSPITAL", "GAUHARPUR SULTANPUR", "KORBAKU", "GURHATTI", "KAZI TOLA", "LAXMIPUR KATTAI", "MANDI BANS", "SIRASWA HARCHAND", "JIGAR COLONY MORADABAD", "FEELKHANA ANSHIK", "SHAHPUR ANSHIK", "ASHIYANA", "DHARAMPUR", "SAIDPUR KHADDAR (GAIRA.)", "MILAK MUDIA", "BHAGATPUR RATAN", "PEEPALSANA ANSHIK", "LALUWALA ANSHIK", "AMBEDKAR COLONY", "DEENDAYAL NAGAR ANSHIK", "MALHUPURA HARDODANDI", "NAVPUR", "BAGIA", "HARTHALA", "DAULATBAGH NAGPHANI UTTARI ANSHIK", "BELDARAN", "BASANTPUR RAMRAI", "HARTHALA SONAKPUR MORADABAD", "KARIYANAGLA SANI", "NAKATPURI KHURD", "HARVOLA WALA", "MEHMOODPUR TIGRI ANSHIK", "NAYA GAON ANSHIK", "ISLAMNAGAR RAMPURA", "BARWARA MAJHRA ANSHIK", "MALWADA URF MANPUR", "KUCHA SHEKH BACHAI", "DAULATBAGH ANSHIK", "SAILATHAN", "DANDI DURJAN", "AKKA BHIKANPUR AITMALI (GAIRA)", "BADHPURA MAHESHPUR KHEM", "SARAI HUSAINI BEGUM ANSHIK", "DARIBAPAN", "TAJPUR MAFI ANSHIK", "SHAUKATBAGH PASHCHIMI", "CHOWKI HASAN  KHAN PASHCHIMI", "MILAK KAMRU", "MOTI MASJID", "LALBAGH UTTARI ANSHIK", "NAYEE BASTI", "RAIPUR SAMDA", "SIHALI AITMALI (GAIRA)", "ASHIYANA.", "TAMBOLI", "MANIHARAN", "KATRA NAJ", "MAHESHPUR KHEM ANSHIK", "VILAKUDAN", "GULADIA", "DAULATBAGH BALMIKI BASTI", "HARTHLA UTTARI ANSHIK", "JAGATPUR RAMRAI", "RANI NAGAL ANSHIK", "NEKPUR", "HARDASPUR", "DILRA RAIPUR", "JAMA MASJID", "BAZAR DIWAN DAKSHINI", "BEGUMPUR", "CHAK AMBEDKAR", "MILAK PEEPLILAL", "SHYAMPUR HADIPUR", "TAKIAWALA ANSHIK")
		case "Moradabad Nagar":
			sections = append(sections, "SHAHPUR TIGRI ANSHIK", "WARSI NAGAR GALI NO. 2 JAMA MASJID", "SURAJ NAGAR", "SHIV MANDIR GULABWADI", "PACHPEDA KATGHAR ANSHIK", "YUSUF COLONY REHMAT NAGAR", "PAIPATPURA MANDI SAMITI", "MEHBULLA GANJ PURVI ANSHIK", "ADARSH COLONY", "DURGESH NAGAR ANSHIK SITAPURI", "JOGANPUR GAYATRI NAGAR", "NAYA GAON ANSHIK GANGAN", "SHIVAJI NAGAR LINEPAR", "SEETAPURI DAS SARAI ANSHIK", "BRAHAMPURI  ANSHIK JAYANTIPUR", "ADARSH NAGAR", "GALI MIYAN SAHAB", "MAKBARA ROAD PEERZADA", "BUDDHI VIHAR AVAS VIKAS", "GUIYAN BAGH ANSHIK", "POLICE LINE", "ROJEY WALI PARIA PEERZADA", "PEERZADA MADHYA", "AVID NAGAR TEELA ANSHIK", "ASALATPURA MADHYA ANSHIK", "HATHI KHANA", "MAJHOLA PURVI", "MANPUR PURVI ANSHIK", "PEERZADA MADHYA ANSHIK", "KHOKRAN UTTRI", "AKHBAR FACTORY DHEEMRI", "DHEEMRI ANSHIK", "BABULAL KA FATAK PATPAT SARAI", "KATAR SHAHEED", "DARGAH MAKBARA", "KATABAGH ANSHIK", "GANDHI NAGAR ANSHIK", "EIDGAH AWADI ANSHIK", "DHAKKA ANSHIK", "GYANI WALI BASTI", "KUMAR KUNJ KOTHI SAHU RAMESH KUMAR", "PREET VIHAR ANSHIK MAJHOLA", "WARSI NAGAR ANSHIK", "BALMIKI MOHALLA CHANDRA NAGAR", "RAMESHWAR COLONY MAJHOLA", "MACHHUAPURA KATGHAR", "RAJO GALI ANSHIK", "RAILWAY SATTLEMENT SECTOR-4 ANSHIK", "SARASWATI VIHAR ANSHIK", "MANPUR NARAYANPUR ANSHIK", "ASALATPURA PULIA ANSHIK", "GAYATRI NAGAR ANSHIK", "MAKBARA PRATHAM", "MAU DEHAT ANSHIK", "GOYAL NAGAR MANPUR", "MADHUBANI ANSHIK", "PRAKASH NAGAR", "ASALATPURA DEHRA ANSHIK", "SONAKPUR DEHAT", "SARAI KOTHIWALAN ANSHIK", "BANK COLONY SIHORA GOVIND", "BARWALAN UTTRI ANSHIK", "PREM NAGAR ANSHIK KANTH ROAD", "KATAR SHAHEED ANSHIK", "INDRA COLONY GULABWADI", "LANGDE KI PULIA ASALATPURA", "JAMA MASJID UTTRI ANSHIK", "MORA NAYEE AWADI", "REHMAT NAGAR SOCIETY  ANSHIK", "GYANI WALI BASTI ANSHIK", "POLICE CHOWKI KE PICHHE LAJPAT NAGAR", "BHUDE KA CHAURAHA ASALATPURA", "GALI NO. 1 JAMA MASJID", "MEHBULLA GANJ ANSHIK", "NURANI MASJID JAYANTIPUR", "JATAV BASTI ASALATPURA DAKSHIN", "PEERZADA DAKSHINI ANSHIK", "PUJERI GALI", "EDAM EIBZE COLLEGE ROAD PEERZADA DAKSHIN", "REHMAT NAGAR GALI NO. 1", "MALIYO WALI GALI KANJRI SARAI", "RAILWAY THIRD GATE", "CHAMDE KA GODAM ASALATPURA", "BUDDHI VIHAR", "LAL NAGRI DHEEMRI", "RAILWAY SATTLEMENT SECTOR-5 ANSHIK", "PAWAR HOUSE MAJHOLA", "SAKKA AVVAL MUGALPURA FIRST", "BHATAWALI ANSHIK", "KHWAZA NAGAR DHEEMRI ANSHIK", "MAKBARA DAKSHINI", "BUDH VIHAR MAJHOLA", "C.P.H. KANTH ROAD", "MAJHRA SHAHPUR", "ZERO GALI WARSI NAGAR", "KATGHAR BEECH ANSHIK", "MATA MANDIR LINEPAR", "SHIV NAGAR", "ASALATPURA UTTRI", "SAGAR SARAI ANSHIK", "IBRAHIM MARKET", "CHHATRE WALI GALI SARAI KOTHIWALAN", "SECTOR - 4B BUDDHI VIHAR", "EKTA COLONY LINEPAR", "REHMAT NAGAR ANSHIK", "MADAIYA BHOGPUR MITHONI", "ASALAT GANJ CHHOTI MANDI", "MILAK BHOLA SINGH SONAKPUR", "BUDDHI VIHAR, SECTOR", "ISMAIL ROAD ASALATPURA", "JIVAN KI SARAI", "SHIVPURI", "GULABWADI ANSHIK", "MANSURI COLONY JAYANTIPUR", "DHEEMRI KHWAZA NAGAR ANSHIK", "UNCHA TEELA JAYANTIPUR", "LAJPAT NAGAR ANSHIK", "GAGAN KI MILAK", "CHAND WALI MASJID BHADORA DEHAT", "BARWALAN ANSHIK", "GOVIND NAGAR ANSHIK", "RAILWAY SATTLEMENT NO. A. SECTOR-1", "SEEDHI SARAI PASHCHIMI ANSHIK", "RAILWAY SATTLEMENT SECTOR-2", "FAKEERPURA ANSHIK", "MAKBARA PRATHAM ANSHIK", "MILAN VIHAR", "SHAKTI NAGAR LINEPAR", "KASHIRAM NAGAR BLOCK -", "MANSOORI COLONY ANSHIK", "PACHPEDA KATGHAR", "SAINI WALI GALI BHADORA", "MANPUR PASHCHIMI", "GALSHAHEED PURVI ANSHIK", "MEHANDI KA TIRAHA ASALATPURA", "LOCOSHED ANSHIK", "HIMGIRI COLONY ANSHIK", "SEEDHI SARAI PASHCHIMI", "HAJI YUSUF WALI GALI PEERZADA", "MEENA NAGAR", "KACHODAN DHAKKA", "SINGHMAN HAZARI ANSHIK", "VEERSHAH HAZARI ANSHIK", "CHIDIA TOLA LINEPAR ANSHIK", "KHWAZA NAGAR", "HAJI KA KUA ASALATPURA", "NETA COLONY JAYANTIPUR", "KHUSHHALPUR ANSHIK", "LAKADI ANSHIK", "MEERPUR MAJHOLI DEHAT", "MAJISTRET COLONY KANTH ROAD", "MAJHOLA ANSHIK", "DARJI GALI MAKBARA DAKSHINI", "NEW CIVIL LINES", "PREM NAGAR ANSHIK", "FAJALPUR ANSHIK", "PRAKASH ENCLAVE", "KAZIPURA ANSHIK", "MAKBARA SABJI MANDI", "LAJPAT NAGAR", "VIJAY NAGAR SHIVPURI", "ABID NAGAR TEELA JAYANTIPUR", "PREM NAGAR LINEPAR ANSHIK", "MENATHER ANSHIK", "BHEEMATHER ANSHIK", "LAKSHMANPURI LINEPAR", "MANSAROVAR COLONY", "KATGHAR BEECH HOLI KA MEDAN", "SHIV MANDIR SAGAR SARAI", "BHADORA DEHAT AMBEDKAR NAGAR", "KUNDANPUR ANSHIK", "VIDYA NAGAR", "SARAI PUKHTA/PAKKI SARAI ANSHIK", "CHANDRA NAGAR ANSHIK", "MEENA NGAR JAYANTIPUR", "HULSAN GANJ LAKDI FAJALPUR", "DEHRI GAON ANSHIK", "SARAI KAZI ANSHIK", "KATRA PURAN JAT", "HARPAL NAGAR", "LAKDIWALAN PURVI", "BODDH VIHAR MAJHOLA", "CHAUHANO WALI MILAK LAKDI", "ZAHID NAGAR KARULA", "BARWALAN DAKSHINI ANSHIK", "KOTHIWAL NAGAR ANSHIK", "MUGALPURA SAKKA DOYAM", "RAILWAY SUTILMENTE SECTOR-7 LINEPAR", "MEHARSHI DAYANAND SURAJ NAGAR", "KACHCHI BASTI SURAJ NAGAR", "JAYANTIPUR ANSHIK", "BHATTI STREET", "GUIYAN BAGH", "BUDDHI VIHAR ANSHIK", "RAM TALAIYA ANSHIK", "ISLAM NAGAR ANSHIK DHEEMRI", "HAYAT NAGAR KATGHAR ANSHIK", "ISMAIL ROAD BEKRI WALI GALI", "ANSARI PARK KATABAGH", "MEENA NAGAR ANSHIK JAYANTIPUR", "SURAJ NAGAR ANSHIK", "PANDIT NAGLA ANSHIK", "BHADORA DEHAT ANSHIK", "PUTLIGHAR  LINEPAR ANSHIK", "MORA MUSTAHKAM PURANI AWADI", "MILAK KALYANPUR", "AJAD NAGAR ANSHIK", "DAN SAHAY KI MILAK", "SHIVAJI NAGAR ANSHIK", "PATEL NAGAR BHOGPUR MITHONI ANSHIK", "BALLAM STREET", "HOLI KA MAIDAN LINEPAR", "BAGIA MAKBARA", "NAYA GAON ANSHIK MAU", "PEETAL BASTI ANSHIK", "CHAU KI BASTI ANSHIK", "MUGALPURA SECOND", "HANUMAN NAGAR LINEPAR", "CHUDI WALI GALI DEHRI", "ADARSH COLONY ANSHIK", "SHIVAJI NAGAR LINEPAR ANSHIK", "SOCIETY REHMAT NAGAR  ANSHIK", "KHUSHHAL NAGAR", "AVID NAGAR TEELA", "BARWALAN DAKSHINI", "JHAJHANPUR KANTH ROAD", "GANDHI ASHRAM", "BHOGPUR MITHONI ANSHIK", "SIRKOI ANSHIK", "KALA PYADA", "YUSUF PARK REHMAT NAGAR", "BLOCK - C GOVIND NAGAR", "P.T.C. FIRST", "KARBALA REHMAT NAGAR", "WALA KI SARAI", "SURYA NAGAR ANSHIK", "BHUDA ASALATPURA", "RAILWAY SATTLEMENT", "SARITA NAGAR LINEPAR", "BALDEVPURI ANSHIK", "SHASTRI NAGAR", "RAILWAY HARTHALA COLONY", "UDPURA KATGHAR ANSHIK", "AJAD NAGAR BHOGPUR MITHONI ANSHIK", "LATPAT NAGAR ANSHIK", "BAKRI KA AHATA DEHRA ASALATPURA", "ZAHID NAGAR ANSHIK", "LODIPUR JAWAHAR NAGAR ANSHIK", "PEERGAIB BAJIGRAN", "SINGHMAN HAZARI", "RAM TALAIYA LINEPAR", "VISHISHT VIHAR", "SURAJ NAGAR GULABWADI", "GADI KHANA ANSHIK KATGHAR", "FRIENDS COLONY", "SHAHPUR ANSHIK", "SEETAPURI DAS SARAI", "VIKAS NAGAR LINEPAR ANSHIK", "SAMBHLI GATE", "TALAB WALI MASJID PEERZADA", "MAJHOLI DEHAT ANSHIK", "PRAKASH NAGAR ANSHIK", "DAS SARAI", "MUGALPURA SAKKA AVVAL ANSHIK", "DEHRI GHAT", "LAL NAGRI DHEEMRI ANSHIK", "JAWAHAR NAGAR", "SHAMSI COLONY DHEEMRI", "DURGESH NAGAR DOBLE FATAK", "MANSOORI AVID NAGAR", "HAJI MADNI KI GALI ASALATPURA", "MILAN VIHAR ANSHIK", "VEERSHAH HAZARI KATGHAR", "KANJARI SARAI JAIL KE PICHHE", "KHUSHHAR NAGAR ANSHIK", "ATAI MOHALLA", "P.T.C. SECOND", "BARWALAN UTTRI CHAMUNDA GALI ANSHIK", "SARAI KISHAN LAL", "SIHORA GOVIND ANSHIK", "KATA BAGH ASALATPURA", "GADIKHANA KATGHAR", "P.A.C.  24 BATALIAN", "NAND COLONY", "SAMRAT ASHOK NAGAR", "DARGAH ANSHIK MAKBARA DAKSHINI")
		case "Kundarki":
			sections = append(sections, "FARIDAPUR HAMIR", "VISART NAGAR URF MAJHRA", "MAUHAMMAD JAMAPUR", "NAUSAN ASHEKHAPUR  AANSHIK", "KULAVADA", "SARKADA KHAS  AANSHIK", "SIRASKHEDHA ANSHIK", "KHAVDIYABHOOD ANSHIK", "HASAN GANJ", "RUSTAMSARAY", "NURULLA  MADHYA", "KUTUBAPUR AJJU", "VAKIPUR JATNI", "MALIPUR", "NAGALA HASHA", "KONDRI", "YOOSUFPUR NAGALIYA MAJRA KONDRI", "DEVAPUR ANSHIK", "GANESH GHAT AI0", "LALAVARA  AANSHIK", "CHAN‍DANPUR ISAPUR", "MILAK GOVIN‍DPUR KHURD", "LODHIPUR VASOO AI0 GAIRA0", "NAGALA  SALAR", "BIKNAPUR", "HARIYANA", "MILK DHOBI VALI", "AK‍KA DILARI ANSHIK", "SAK‍TOO NAGLA AANSHIK", "NURULLA MADHYA", "LALAPUR TITRI", "NAUSNA SHEKHPUR AANSHIK", "ASDAPUR", "JAITPUR PATTI AANSHIK", "MEHM‍DINGAR", "RATANPUR KALA AANSHIK", "KUTUBPUR RAIB GAIRA0", "RASULAPUR HAMIR", "MAUSMAPUR", "SALEMAPUR", "SHEKHOOPUR KHAS", "BASERAKHAS", "DOMGHAR", "BHAJANPURI", "IMRATAPUR FAKHARUDDINAPUR", "N‍YAMTPUR IK‍ROTIYA ANSHIK", "ALLAPUR BHIKAN", "MASEVI RASULAPUR  AANSHIK", "JAFRAPUR", "ALIRAJAPUR", "MUDIYA MALOOKPUR AANSHIK", "SAB‍JIPUR AANSHIK", "BHADASNA", "MADHUPURI", "TAKHTAPUR HASHA", "SAIFAPUR CHITTU", "KHANPUR MAJHRA MAU0 GAIRA0", "NAIYMATPUR BOOCHA GAIRA0", "CHAMRAU ANSHIK", "SIKRI", "MOHNPUR AANSHIK", "LALPUR GANGVARI AANSHIK", "SAIFAPUR PALLA", "RUSTAMAPUR BADHAMAR AANSHIK", "ILAR RASULABAD", "RAMAPUR MAIGAN", "AHAROLA", "CHITTUPUR", "MOHAMMADPUR BASTAUR  AANSHIK", "GADIPUR", "PARASUPURA VAJE", "BASERA PATTI GAIRA.", "MUNIMAPUR", "SHIVUPRI", "CHUHA NGALA", "DHAKIYA JAMAPUR", "NAKTAPURI KALA", "PANDIYA AANSHIK", "TAH NAYAK", "MATHAIYYA MAJHRA", "TUTIPUR GAIRA", "BHAYPUR AANSHIK", "MACHHARIYA  ANSHIK", "MAHESHAPUR BHILA", "MAJHRA JANRAILVALA", "NARAKHEDA", "MAJHOLI KHAS AANSHIK", "AK‍KA PAN‍DE BHOJPUR ANSHIK", "VAHPUR AANSHIK", "MAHMOODANGLA GAILA0", "KHANPUR MAJHRA GAIRA0", "NAHTAURA GAIRA.", "KAYASTHAN MADHYA", "TAKHTPUR ALLA URF NANKAR AANSHIK", "AHAMAD NAGAR JAITVADHA AANSHIK", "BHAYPUR GAIRA0", "NAJRPUR ANSHIK", "DUPEDHA", "JAITAPUR", "CHHIRAVLI", "MAJHRA GHOSIPURA", "SIRAS KHENDA ANSHIK", "KHAIKHEDHA ANSHIK", "FAT‍TEHPUR KHAS", "DHEAVI VALI MILAK", "SUNARI", "BAHADURAPUR RAJAPUT", "BHAISIYA", "CHAK HAMIRPUR GAIRA0", "KHARGAPUR JAGTAPUR", "KAMALPUR KUNWARSAIN GAIRA", "GHOSIYAN", "RUPAPUR", "KHAIRKHATA  ANSHIK", "JAITIYASADULLAPUR AANSHIK", "BHITAKHEDA", "BHAISIYA ANSHIK", "SIKMAPUR PANDE", "MANKARA", "SADAT DAKSHINI AANSHIK", "JATPURA AANSHIK", "KHURRAMANAGAR BICHAULA", "CHANDAPUR MANGOL", "MILK KHARAGPUR BAJE", "LADPUR GAIRA", "RONDA AANSHIK", "GOT ANSHIK", "BARAITHA KHIJRAPUR", "HADAYPUR", "RAMAPUR BHILA AANSHIK", "BHAYPUR ANSHIK", "NANAPUR", "KABIR NAGAR KALONI", "MADANAPUR", "ANVALAGHAT", "MUNDHA PANDE  ANSHIK", "TAKHTPUR ALLA URF NANKAR  AANSHIK", "JAGARAMPURA", "KISVA NAGLA", "NURULLA DAKSHIN AANSHIK", "MAJHRA MILK (DILARI)", "AHAMDAPUR", "SAMADARAM SAHAY", "MOHAMMADAPUR", "BARAVARA KHAS ANSHIK", "DAULARI ANSHIK", "HARSAINPUR", "LAKHANPUR", "MANPUR PATTI AI0 GAIRA0", "UNCHAKANI", "KADRAPUR MASTI", "NAIYAMATPUR BUCHA GAIRA.", "CHAKFAJLPUR AANSHIK", "SIRSA INAYTAPUR ANSHIK", "VUJHPUR MAN", "MALKAPUR", "SIHORA BAJE  ANSHIK", "SAUNDA", "CHATPUR NAKTAKHEDA", "KAYASTHAN PURVI", "ROOS‍TAMPUR PARMAL GAIRA0", "SAHRIYA AANSHIK", "RANIYATHER", "LAITHER", "KURI", "DEVAPUR", "RAJODA KANNAR DEV", "AJAMATPUR GAIRA0", "BHANDARA", "DHATURA MEGHA NAGLA  AI0", "TEVRAPATTI URF KAJIPURA AANSHIK", "GARDAI KHEDA", "VUJHPUR AASHA AANSHIK", "KOHNKOO", "KHANAPUR LAKKHI", "JAITIYA SADULLAPUR  AANSHIK", "SULTANAPUR", "DALAPATPUR ANSHIK", "DHATURA MEGHA NAGLA AANSHIK", "VINAVALA", "JAITPUR BISAHT AANSHIK", "HAMIRAPUR", "GOVARDHANPUR", "VARVARA KHAS AI0", "VIRPUR THAN ANSHIK", "SADAT UT‍TRI", "GILAPURA", "SIKANDARAPUR PATTI", "LALATIKAR MAHESH NAGLI", "SADAT PASH‍CHIMI AANSHIK", "MOHM‍MADPUR BAS‍TAUR", "MANSURAPUR", "KAMALAPUR FATEHABAD", "LALPUR HAMIR AANSHIK", "LALAPUR HAMIR", "KARIMAPUR JABTI", "GURER AANSHIK", "NABBA NAGLA", "TAHRPUR AB‍BAL AANSHIK", "BHIKMAPUR KULAVADA", "JAITIYA FIROJAPUR", "GANESH GHAT  AANSHIK", "HIRAN KHEDA  AI0", "YAKUTPUR CHHAPARRA", "ARIFPUR KUN‍DARKI GAIRA0", "KARNAPUR", "SAIJNA", "KAKAR GHATE", "DINGARPUR AANSHIK", "NAJRPUR PURVI", "NURULLA PASH‍CHIMI", "GOVIN‍DPUR KALA", "IMARATPUR UDHO0 AANSHIK", "HUSAINPUR HAMIR", "MATIPUR URF MAINI AANSHIK", "SARDAR NAGAR AANSHIK", "VIRAPUR VARIYAR URF KHARAG  AANSHIK", "VIRPUR GAIRA0", "SAHU NAGLA", "AK‍KA DILARI", "CHAUDHRI PATTI GAIRA.", "MASEBI RASULAPUR", "KAYASTHAN PASH‍CHIMI AANSHIK", "DEVAPUR  AI0", "HISAMAPUR", "BHATAPURAKUNDARKI", "RASULAPUR NAGLI", "MILK MEVATI", "KHARAGPUR BAJE ANSHIK HADAYPUR GAIRABAD", "HALA NGALA", "BHANDRI AANSHIK", "GHONDA AANSHIK")
		case "Bilari":
			sections = append(sections, "UDAYPUR NARAULI", "ALHAPUR URF RAMPUR", "MOHAMMAD HAYATAPUR", "ABHANPUR NARAULI", "BHIDVARI", "BHAISOD", "BAHORANPUR NARAULI", "KHITABPUR URF KHANUPURA", "BAHADARPUR", "MAHAUDPUR MAFI", "HATHIPUR CHANDU", "BAHATA SARTHAL", "JIGANIYA", "PIPLI", "GORASHAHAGADH", "NURUDDINPUR URF GANJ", "MUDIYA RAJA", "AJAMGAR-CHOPDA", "GWALKHEDA", "MAHAMUDAPUR MAFI", "ASALATPUR  MAJARABILARI", "KALAL KHEDA URF SHAHAJADAKHEDA", "VICHAULA DHUKI", "JANAKPUR", "DHANNUPURATUR  KHEDA", "RAMNAGAR GANGPUR", "ANSARIYAN D0 WARD-12", "SANAI", "MITHANPUR BALLU URF NAGLA", "SAFEELPUR", "DHARMPUR KALAN", "GAJUPUR", "JHAKDA", "SAIFPUR JAGNA", "NASEERPUR BHIKKA", "CHANGERI", "UDRANAPUR CHAK URF VIRMAPUR", "LUDHPURA", "DEVRI", "DAULTAPUR BHARAKAU", "KUDI MANAK", "NARAYANPUR DEVA", "DOHARI", "NASIRPUR KUNDARKI", "CHANDPUR GANESH", "SIHARI NANDA", "SAIFAPUR MALHU", "NIYAMATPUR", "DEVIPUR URF NAGLA", "MUNDIYA RAJA", "MAHMUDPUR MAFI", "GVARAU", "NAIYAKHEDA", "SARAY PANJU", "NASEERPUR DHANNA", "HATHIPUR CHITTU", "ANSARIYAN U0", "RUSTAMPUR KHAS", "NIMRI", "JWALAPURI", "IBRAHIMAPUR", "KHAJRA", "ABUPUR KHURD", "SISAUNA", "JAMALPUR", "MITHANPUR MAHESH", "BHUD MARESI", "DANDA", "TEVRAKHAS", "RUSTAMNAGAR SAHASAPUR", "MALHAPUR JANNU", "UGHANPUR", "SHIKARPUR MAJHRA MITHANPUR MAHESH", "DINAURA", "SHAHPUR", "MANGUPURA", "HAMAJAPUR", "CHITERI", "HAJIPUR", "SANDALPUR", "TARAPUR SYODARA", "SIHARIMALA", "MAINATHER", "ANSARIYAN P0 WARD -6", "CHIDIYATHER", "SEELPUR", "MALHIPUR MAHMUDA NAGLA", "THAKURAN PURVI", "AMARPUR KASHI", "SYODARA", "MUNDIYA BHIKAM", "SATARAN", "ROJA", "HAYATPUR", "HARAURA", "KUAN KHEDA", "MAKANPUR", "ASALAT NAGAR (KALIJAPUR)", "SIHARI LADDA", "DHAKIYANRU", "KORIYAN WARD  N0-4", "UMARAGOPALPUR", "KARSARA", "ABDULLA MADHYA", "NAVADASHAHAPUR", "PALIYA", "MANAKPUR KUNDARKI", "THAKURAN", "ALAHADADPUR", "SURAJPUR KALYAN", "ALINAGAR", "KARAVAR", "NAMAINI UDAIYA", "NAUSNASYODARA", "TEVARKHAS", "KHANDUVA", "SUJANPUR", "KORIYAN WARD -4", "KHABRI GANDU", "HATHIPUR BAHAUDDIN", "NAGALAKAMAL", "SAIDNAGAR URF  NAYAGANVA", "NAGLA NASSU", "UDRANAPUR CHAK URF  VIRMAPUR", "DAMMU NAGLA URF DAMPURA", "MILAK AHLADADPUR KARAR", "MADHVA", "SAMASPUR", "BARAULI RUSTAMAPUR", "MILAK KATHAIR ASALATNAGAR BAGHA", "BAJAR PASHCHAMI WARD NO-1", "ANSARIYAN D0 WARD -12", "HUSAINAPUR PACHATAUR", "PEETPUR", "RUSTAMNAGAR SAHASPUR", "AKBARPUR KHAS", "BHIKANPURBAGAH", "PIPALICHAK", "NAGALIYA MASHMULA", "AHLADPUR KHEM URF RAIPUR", "MAUHAMMAD IBRAHEEMPUR", "RAMAPUR PATTI GUJAR", "MURADPUR", "DHATRARA", "SONAKPUR", "MUKARRABPUR URF KUAKHEDA", "NAGALIYASHAHPUR", "NAGALAGUJAR", "THANVALA", "KORIYAN WARD NO-4", "NARUKHEDA", "MAKRANDPUR", "THAKURAN PASCHIMI WARD -7", "NARAUDA", "ABDULLADAKSHINI", "JASRATH PUR", "ABDULLAPURVI", "TISAVA", "MOHD SADIKPUR", "BAGRAUA", "BAJAR PASCHIMI ANSHIK", "SHERPUR MAFI", "JARGANV", "MAHAMUDPUR MAFI", "SARTHAL", "KHATA", "BICHAULA- KUNDARKI", "TANDAAMRAPUR", "JALALPUR KHAS", "KANOVI", "PALANPUR", "JATAPURA JHADU", "ISAPUR", "BHUDABAS", "BHUDA", "AJMAT NAGAR JYODERA", "BAKAINIYACHANDAPUR", "MANKULA", "IMRATPUR SYODARA", "MILAK AZAM BALI (BHIKANPUR BAGHA)", "KHANPUR", "FATEHAPUR SALAVAT NAGAR", "SAMATHAL", "RAMPUR", "BAJAR PURVI", "THAKURAN DAKSHINI", "CHIDIYABHAVAN", "MAHALAULI", "UMRARA", "ARIKHEDA", "ASALAT NAGAR VAGHA", "SHYAMSINGHPUR URF BHUDI", "AHAMDABAD KASAURA", "SHERAPUR MAFI URF DHANIYAKHEDA", "GVAR KHEDA", "ABUSAIDAPUR", "KHASEPUR", "NAMAINI GADDI", "NAGALIYAJAT", "MIRAPUR MAFI", "ALAHADADAPUR KARAR", "MAHAJANAN PURVI", "MUDIYAJAIN", "MALHAPUR SINDHARI", "HASANPUR RUP", "NARSARA", "SUNDARPUR", "VAGHI GOVARDHANPUR", "KHAVRI AVVAL", "BACHHAL BHUD", "CHIDIYABHVAN", "AHARAULA", "BAJAR PASCHIMI WARD NO-1")
		case "Suar":
			sections = append(sections, "ALI  NAGAR JAGEER-2", "SHIVNAGAR NEAR LOHARI", "FAJALPUR NEAR DHANOORY", "CHAK NO-9(2) TANDA", "RATANPURA-2", "KHUSHHALPUR", "CHHITARIYA GAJIR-2", "SEDU KA MAZHARA", "MASWASI-3", "CHAK NO-14 TANDA", "DADHIYAL EHATMALI-4", "ZALPUR", "RATUANAGALA-2", "JALAFNAGLA-1", "PIPLI NAYAK-3", "CHAK SUAR-6", "DADHIYAL EHATMALI-6", "MUBANA", "LADPUR NEAR SEMRA-1", "BIJARKHATA UTTRI-3", "SHIKARPUR", "MAZRA TRILOK PUR", "DUNDAWALA AEHATMALI , MUSTEHKAM 5", "CHAK NO-2 TANDA", "FAZALPUR PUSWADA", "MALGOSHA", "SEMRA NEAR LADPUR-1", "AGLAGA -1", "KAITHOLA", "IMRATPUR", "DUNDAWALA  AEHATMALI , MUSTEHAQAM-1", "GADDINAGLI-1", "RAMNAGAR LATIFPUR-3", "LAMBAKHEDA-3", "NAVAB GANJ PARSAL", "CHAK NO-15 TANDA", "ASALATPUR NEAR  SHIV NAGAR", "BIJARKHATA UTTRI-1", "RAMNAGAR LATIFPUR-2", "DHANOORI-1", "PIPLI NAYAK-1", "CHAK NO-25 (3)TANDA", "CHAK NO-6 TANDA", "MILAK DUNDI", "ALI  NAGAR JAGEER-1", "CHHITARIYA GAJIR-1", "SIKANDARABAD-1", "SENTAKHEDA-1", "MILAK MOHMMAD SHAH KHAN", "LOHARI", "CHAK NO-22(1)TANDA", "CHAK KHARDIYA", "RAYPUR-2", "RAYPUR-3", "PARBATPUR", "SENTAKHEDA-5", "LADPUR NEAR SEMRA-2", "LAMBAKHEDA-2", "MADHUPURA-1", "AHMADNAGAR NEAR BHAGWANT NAGAR", "DADHIYAL EHATMALI-1", "KHEMPUR-1", "MAJHRA  KALINAGAR  LACHHIWALA MAJHRA KHEMWALA -1", "RASOOLPUR FAREEDPUR-2", "CHAKJEEM", "CHAK NO-26TANDA", "KASHIPUR -2", "NARPATNAGAR-6", "SITLA", "CHAK NO-22(2) TANDA", "RUSTAMNAGAR NEAR RAYPUR", "MUKUTPUR", "NAANKAR RANI-4", "CHAUPURA-3", "NARPATNAGAR-7", "DHANUPURA-1", "BHUVRA AHETMALI-1", "KARKHEDA-2", "SONAKPUR-1", "BHUVRA  MUSTAHAQAM-2", "TODIPURA", "SAHARIYA JAWHAR", "CHAK NO-16TANDA-2", "CHANDUPURA SHIKAMPUR-2", "KHERA PARSAL", "BHAUPURA", "CHAK NO-18 (1)TANDA-1", "KUNDESRA", "ABBASNAGAR TANDA", "GOSIPUR", "CHAK NO-13 TANDA", "DUNDAWALA  AEHATMALI , MUSTEHAQAM-2", "KAISONAGALI TANDA", "MAJHRA  KALINAGAR  LACHHIWALA MAJHRA KHEMWALA -2", "CHANDUPURI", "RAJPURA TANDA", "SUAR KHAS UTTARI", "CHAK NO-25(1)-2TANDA", "PATTI KALAN-2", "LADPUR NEAR BATHUAKHERA", "MEWALA DHARU", "NARPATNAGAR-5", "AGLAGA -2", "MOHABBAT NAGAR KUNDESRA", "CHAUPURA-2", "KHARDIYA", "KASHIPUR -1", "NAVAB NAGAR PADPURI", "RAIKANAGALA", "CHAK NO-25(3) TANDA-2", "KASIYA KUNDA-2", "AZEEM NAGAR", "CHAK NO-11 TANDA", "BHUVRA  MUSTAHAQAM-1", "MANDBA HASANPUR", "MILAK MOHMMAD NAKI", "DEVIPURA", "CHAK SUAR -3", "SAID NAGAR TANDA", "PATTIKLA-1", "RATUANAGALA-3", "CHAK SUAR -4", "DARAV NAGAR", "SIKANDARABAD-2", "MILAK HAFIZ BALLI", "CHANDUPURA SHIKAMPUR-1", "KHANPUR UTTARI-2", "KAREEMPUR", "MUNDIYA RASOOLPUR", "MILAK NOUKHAREED", "MILAK QAZI", "AARSAL PARSAL-1", "SAMODIYA-1", "KARKHEDA-1", "SHAHPURA-3", "NARAYANPUR", "DADHIYAL MUSTEKAM -3", "CHAK NO-9- TANDA(1)", "PASIYAPURA", "CHAK NO-20(2) TANDA", "ABBASPUR", "CHAPARRA", "DHIMARKHEDA", "MASWASI-2", "MANDBA HASANPUR-2", "SHIV NAGAR NEAR ASALATPUR", "CHAK NO-1 TANDA", "DAKKANAGALIYA", "SENTAKHEDA-3", "KHOD KALA-1", "KHAZUYA KHEDA", "MAHUAKHEDA TANDA", "AKHBARABAD", "CHHITARIYA GAJIR-4", "LADPUR BIB", "LAKHMIPUR", "CHAK NO-21TANDA", "MILAK MUGLA URF MUGLPIUR", "CHAK NO-18 (2) TANDA", "CHAK NO-7 TANDA", "RAMNAGAR LATIFPUR-1", "TANDA GOLU", "SENTAKHEDA-2", "MILAK GULAM KHAN", "MAHARAJ PUR", "MILAK KHANPUR", "SHAHDRA NIKAT DHANPUR", "RATUANAGALA-1", "LODHIPUR NAYAK-2", "DADHIYAL MUSTEKAM-6", "CHAK NO-10(2)TANDA", "NARPATNAGAR-8", "CHAK NO-4 TANDA", "MADHUPURA-2", "IMARATI NIKAT AKBRABAD", "CHAK NO-19 TANDA", "RATANPURA-1", "CHAK NO-12 TANDA", "DHAKPURA", "KESHONGLI SUAR", "MILAK KHOOD", "AHMADABAD URF TELIPURA", "CHAK NO-25 TANDA", "CHAK NO-25 (2)TANDA", "SEMRA NEAR LADPUR-2", "CHAK NO-10(2) TANDA", "MIRZAPUR SUAR", "DADHIYAL EHATMALI-2", "MASWASI 5", "ALI  NAGAR UTTARI", "CHAK SUAR-7", "MANPUR UTTARI-1", "CHAK NO-27TANDA", "CHAK NO-28 TANDA", "SHAMSHABAD  URF KHABRIYA", "KHANPUR KHARVI", "DHARMPUR GARVI", "SIRKAPATTI KUNDESARI", "MILAK ASAD KHAN", "KUNDESRI", "BUDI DARIYAL", "KASIYA KUNDA-1", "JATPURA", "NAYA GANV NAZIBABAD", "GADDI NAGLI- 2", "DADHIYAL MUSTEKAM -2", "GANGAN NAGLI", "HUSAINGANJ GADDINAGLI", "RASOOLPUR  -1", "CHAK LADPUR", "CHAK HARDASPUR", "NARPATNAGAR-4", "RASOOLPUR  -4", "BATHUAKHERA", "SUAR KHAS DAKSHNI -2", "HAKIMGANJ", "CHAK SUAR -1", "WAHADUR KS MAZJHAR", "SIRKAPATTI RUPAPUR", "SUAR KHAS DAKSHNI -3", "DUNDAWALA AEHATMALI , MUSTEHKAM 4", "LOHARA TANDA", "DADHIYAL EHATMALI-3", "ZAMALGANJ", "DADHIYAL EHATMALI-5", "SHAHPURA-1", "AARSAL PARSAL-2", "SIHORA", "GULAD PEEPALSANA-2", "PARSHUPURA", "KARKHEDI", "KHEDA GAJROLA", "DADHIYAL MUSTEKAM -1", "NAANKAR RANI-1", "RUPAPUR", "GAZROLA", "CHAK KUNDESARI", "DADHIYAL MUSTEKAM-5", "SHEKHUPURA", "NEWVILLAGE NIKAT AKBARABAD", "SARTHAL", "MANPUR UTTARI-2", "PUSWADA-1", "DHANUPURA-2", "RAMPUR DHAMMAN", "DHARAMPUR UTTARI", "CHAK NO-26 TANDA-2", "GOVINDPURA NEAR RAYPUR", "MOHABBATNAGAR NEAR SUAR", "SUAR KHAS DAKSHNI -1", "BADALI-2", "CHHITARIYA GAJIR-3", "RUSTAMNAGAR NEAR CHAPRRA", "MUKRMPUR", "RASOOLPUR  -5", "CHANDUPURA SHIKAMPUR-3", "RAFTPUR", "RASOOLPUR  -3", "BHATIKHEDA", "AARSAL PARSAL-3", "MANPUR CHHAPAT", "BHUVRA AHETMALI-2", "NAUGJA", "MASWASI-1", "DHANOURI-2", "RASOOLPUR  -6", "BADALI-1", "MEWLA KALA", "DADHIYAL MUSTEKAM-4", "LALPUR", "DUNDAWALA AEHATMALI , MUSTEHKAM 3", "CHAK NO-18 TANDA", "SONAKPUR-2", "LODHIPUR NAYAK-1", "DHANPUR NIKAT  SHAHDRA", "CHAK NO-3", "CHANDPUR", "RAYPUR -1", "JAHANGEERPUR", "LAKHMAN NAGALA", "CHAK NO-23 TANDA", "GULAD PEEPALSANA-1", "SHAHPURA-2", "CHAK NO-17 TANDA", "KHANDI KHEDA", "SAMODIYA-2", "AHMADNAGAR URF MASAIYAN MOUHAMMDI", "FAJULANAGAR", "KISHANPUR DULLANAGALA", "LAMBAKHEDA-4", "BHOLANAGALA", "HASAN PUR UTTARI", "ALLAPUR", "LOHARRA INAYATGANJ", "DHANOORA", "NABAB NAGAR TANDA", "SENTAKHEDA-4", "CHAK NO-16 TANDA -1", "DHANUPURA-3", "DADHIYAL MUSTEKAM-7", "CHAK NO-20(1) TANDA", "CHAUPURA-4", "RASOOLPUR  -2", "JALAFNAGLA-2", "LAMBAKHEDA-1", "PUSWADA-2", "NAGALIYA", "MASWASI-4", "DHIRAZNAGAR", "SED NAGAR NEAR ASALATPUR", "RASOOLPUR FAREEDPUR-1", "CHAUPURA-1", "CHAK NO-5 TANDA", "CHAK SUAR-5", "BHAGWANT NAGAR -2", "SURAZPUR", "NAANKAR RANI-2", "JAMIN GANJ", "KHANPUR UTTARI-1", "BIJARKHATA UTTRI-2", "KALYANPUR", "CHAK SUAR -2", "NAANKAR RANI-3", "SONAKPUR-3", "KHOD KALA-2", "NARPATNAGAR-1", "NARPATNAGAR-3", "KISHANPUR MOLLAGARH", "CHAK NO-8 TANDA", "SENTAKHEDA-6", "CHUKHANDI", "CHAK NO-24TANDA", "PIPLI NAYAK-2", "SAHRIYA LAKKKHA", "KHEMPUR-2", "CHAK GAZROLA", "NOORPUR HAZOORPUR", "BHAGWANT NAGAR-1", "PADPURI SUAR", "NARPATNAGAR-2", "MOHBBATGANJ TANDA")
		case "Chamraua":
			sections = append(sections, "HASHMINAGAR", "SARKADI -1", "AHMADABAD", "ALIGANJ BENJEER-2", "HAKIMGANJ", "TALABPUR-1", "MILAK KHANAM", "KHAUD-2", "SANAYYA JATT", "SANAYYA CHAK", "ISWARPUR", "MALHAPURA", "JUTHIYA-3", "NAGLA GANESH", "JITHANYA SHARKI", "NAVI GANJ BATHUAKHEDA", "DARSHANPUR", "PAINDA NAGAR", "AZEEMNAGAR", "WAHADURPUR", "DEWRANIYA SHARKI", "MOHANPUR", "PATTHAR KHEDA", "BAGADKHA-2", "HASANPUR BILASPUR", "DILPURA-1", "KHAUD-1", "AHAMAD NAGAR PADMPUR", "SIGANKHEDA-4", "PARSHUPURA", "RAZA NAGAR-1", "MILAL AFZAL KHAN", "DEVRANIYA SHUMALI 1", "GOVIND PURA NEAR PADAMPUR", "SEJNINANKAR-1", "JUTHIYA-1", "DAUKPURI TANDA-4", "SEDULLANAGAR", "DURNAGLA", "BAIZANI", "SARABA-1", "LALPUR KALA-4", "KISROL", "HAMIRPUR-2", "IMRATA RAY", "MIRZAPUR BILASPUR", "IMRATA-3", "SIGANI", "KUMHARIYA KLA", "AHMAD NAGAR PAHADI", "MANUNAGAR", "NARKHEDA", "RAMPURA", "MAHUNAGAR -3", "NURPUR", "MEERAPUR MEERGANJ-1", "DHANUPURA", "PARCHAI", "BAZABALA-1", "RAJPURA SUAR", "KOYLA-4", "MUSTAFABAD URF TAKLABAD", "SAIDNAGAR BAZARPTTI", "SIGANKHEDA-2", "SUAR KHURD", "SEJNINANKAR-6", "PAHADPUR", "AHMADNAGAR KA A", "CHHATRAPUR", "MILAK YAKOOB ALI KHAN", "PEGA", "PARSIPURA", "NAGLIYA AAQIL-4", "KASHIPUR-2", "SHAHAJAD NAGAR", "MILAK HASHAM", "KAREEMPUR SHARKI", "PIPALIYA JAT", "AHAMDNAGAR NIKAT FAIZGANJ", "KOYLA-1", "LALPUR KALA -3", "IMRATA-4", "KHIMOTIYA KHEDA", "SIGANKHEDA-3", "MURSAINA", "BHABANIPUR", "KHEDATANDA-1", "BHOT BAKKAL -1", "MILAK WADULLA", "BAMNA", "SAKATPURA", "NAGLIYA AAQIL-7", "SARKADI -2", "HAMEERPUR", "MUSTAFABAD KLA", "MIRAPUR MEERGANJ -2", "KHANDIA-1", "MIRAPUR BILASPUR", "WAJIDPUR", "BAGRAWWA-1", "BADPURA SHUMALI", "BHAMROA-3", "PAHADI", "BAGI-2", "AHMADNAGAR NEAR AGA", "HARNAGLA", "SURATSINGHPUR URF NEW VILLAGE-2", "KUNVARPUR PADAMAPUR", "HUSAIN GANJ BHUDA", "RAZA NAGAR-2", "HOSPUR", "WAHADUR GARH", "KOYLA-3", "GAMMANAPURA", "BADHPURA SHARKI", "SURATSINGHPUR URF NEW VILLAGE-1", "MILAK BHURE KHAN", "SAIDNAGAR URF MADAIYAN PUSE", "SHADINAGAR HARDASPUR", "KASHIPUR-7", "KUKDI KHEDA", "DARIYA GADH", "ALINAGAR SHUMALI", "MUSTAFABAD KHURD", "KALIYA NAGALA", "JATPURA-1", "IMRATA-2", "BAGADKHA-1", "MENHDINAGAR", "SEJNINANKAR-5", "NAGALIYA KASAMGANJ", "SALVE NAGAR", "MAHUA KHERA SUAR", "INDRI", "JHURAKJHUNDI", "ISAMAIL NAGAR NEAR AHMADNAGAR", "MILAK HASHAM IBRAHEEM KHAN", "NABABGANJ SHUMALI", "SIGANKHEDA-6", "MOHAMMAD NAGAR", "NORANGPUR", "MUNDIYA SEDNAGAR", "MUSTAFABAD DAUNKPURI", "NABIGANJ", "DILARI", "BANSNAGLI", "SHIV PURI", "BHOT-2", "BIJHDA", "SEDABAD", "KISANPUR PANCHAKKI", "DALELNAGAR-1", "MADAIYAN BHAJJAN", "PAIGAMBAR PUR", "SAID ALIGANJ BANZARIYA", "JATPURA-2", "BAGADKHA-3", "IMRATA-1", "KASHIPUR-1", "MAHUNAGAR -1", "NASRAT NAGAR", "JUTHIYA-2", "KUNVARAPUR NANAKAR-2", "NASIMGANJ", "SAHRIYA NARPAT", "HARDASPUR KOTRA-2", "DARA NAGAR", "BHOT", "KHANDIA-3", "KASHIPUR-6", "NANKAREN", "HAMIR PUR -3", "SHADI NAGAR HAJIRA-2", "DHAKKA HAJINAGAR 1", "HAYATNAGAR DAKSHNI", "LALPUR KALA -2", "NABI GANJ PIPLI", "PAHADPUR BILASPUR", "AHROLA", "MILAK MIRZA FAYYAZ", "KHEDATANDA-4", "INDRA-1", "MOMINPUR AHAMDABAD-1", "MILAK ABBU KHAN", "SHAHDRA NEAR  BHOT", "PRITHVI NAGAR", "KHEDATANDA-3", "GAJUPURA", "ALIGANJ BENJEER-1", "SHAFAYAT NAGAR", "SEJNINANKAR-4", "TALABPUR-2", "ANDKHEDA", "MOHD PUR SHUMALI", "MILAK NAGALI", "BHOT BAKKAL -4", "DILPURA-2", "KHANDIA-2", "MILAK AARIF", "SEDNAGAR MUNDIYA", "NAWADA", "ASHOKPUR", "NAGLIYA AAQIL-2", "SHADINAGAR HAZIRA 1", "SAHRIYA DRAJ", "KHODPURA 1", "MUNDIYA", "MEHNDINAGAR SHUMALI", "MILAK TAJ KHAN", "MILAK  SANYYA", "SIGANKHEDA-1", "PIPALGAUN-2", "BUDANPUR", "KOYLA-2", "PASIYAPURA SHUMALI", "HARETA-1", "BHOT BAKKAL -3", "SEJNINANKAR-2", "RATANPURA SHUMALI 1", "HARDASPUR KOTRA-1", "KARIMPUR GARVI", "KUMHARIYA-1", "LALPUR KALA -1", "CHAMROUA-4", "DAUKPURI TANDA-3", "CHAMROUA-3", "MARGHATI", "HAMIRPUR-4", "LALUNAGLA-1", "MAHUNAGAR -2", "DEVRANIYA SHUMALI 2", "MOMINPUR AHAMDABAD-2", "IMARTI", "PIPALGAUN-1", "KUMHARIYA KHURD", "AJAYPUR", "MANGUPURA", "KHODPURA 2", "JITHANIYA JAGIR", "KHEDATANDA-2", "FAZGANJ", "KESARPUR", "DAUKPURI TANDA -1", "CHAMROUA-6", "MATKHEDA-1", "BHOT-3", "KHUDAGANJ", "KASHIPUR-8", "GAGAN NAGALA", "BHAMROA-2", "BAGRAWWA-2", "KASHIPUR-3", "BIZPURI", "BIHARI NAGAR", "BAHADARGANJ", "MUTIYAPURA NIKAT BHOT", "CHAMROUA-1", "PIPLI", "NAGLIYA AAQIL-5", "PATTI ASHOKPUR", "NAGLIYA AAQIL-6", "KAJARHAI", "BHATPURA", "ABBAS NAGAR SUAR", "HAMIRPUR-6", "KOYLI", "NAGLIYA AAQIL-8", "KARANPUR", "HARETA-2", "MILAK MALANG GAAN", "KHIZARPUR", "KHURSHEED NAGAR", "NAGLA BANSNAGLI", "SHADINAGAR SAKATPURA", "MAJHRA AMEER KHAN", "KHAIRULLA PUR", "BAGI", "BAZABALA", "KUMHARIYA-2", "PADMAPUR", "BAHADUR GANJ", "AHAMDNAGAR NEAR ISMAIL NAGAR", "AAGA-2", "NAGLIYA AAQIL 1", "MATKHEDA-2", "SURJAPUR", "BAIRAM NAGAR", "DHAKKA HAJINAGAR-2", "SEJNINANKAR-3", "SIGANKHEDA-7", "KUCHITA", "HAMIRPUR-5", "SANIYA AHMAD", "SARKADI", "ZAGATPUR", "BAZANA-2", "PRANPUR", "SANAYYA SUKH", "SHADARAUFR THUNAPUR", "PAIMPUR", "PADPURI NEAR KALYANPUR", "BHOT BAKKAL -2", "KASHIPUR-5", "BAINDU KHEDA", "AHMADPUR", "KASHIPUR-4", "IMRATA-5", "SHIKARUR", "KISANPUR ATRIYA", "AAGA-1", "CHAMROUA-5", "KHIMOTIYA BAKHTI", "MANULLAPUR", "NABAB  NAGAR TAH", "LODHIPUR GULADIYA", "DAUKPURI TANDA-2", "SENPUR", "MISRINAGAR SHUMALI", "HASANPUR SHARKI", "DALELNAGAR-2", "RANUA NAGALA", "RABNNA", "LALUNAGLA2", "BHAMROA-1", "CHAMROUA-2", "SIGANKHEDA-5", "ATRIYA", "KUNVARAPUR NANAKAR-1", "TAHKLA", "MUTIYAPURA BAZARPATTI", "MANAKPUR ANZARIYA", "HAMIRPUR-1", "AHMADNAGAR TARANA", "ALI NAGAR TAH", "REHPURA", "SARABA-2", "BAGI-1", "RATANPURA SHUMALI 2")
		case "Bilaspur":
			sections = append(sections, "SIHARI-1", "MAJULLANAGAR-1", "PURNAPUR", "GIRDHARPUR", "DHANELI UTTARI-2", "BHATTI TOLA-2", "KHANPUR", "ISWARPUR", "CHICHOLI", "HAMIDABAD-2", "BANSKHERA", "JIYARAT SHIRIMIYAN-4", "SIRRA", "SUJATGANJ KHERA", "CHANDELA", "HARAIYA KURD", "HINGANAGLA", "HIRANKHERA", "BHAWARKI", "BOSENA", "MUDIYA NEAR MANKARA", "RAIPUR", "GOKAL NAGRI", "SINGRA-1", "AHRAULA", "MUBARAKPUR-2", "NAWABGANJ", "RAMNAGRIYA", "GADAIYA NIKAT HARSONAGLA", "KARSAULA", "SAEDABAD", "[DHURYAI", "JAMAPUR", "MANKARA", "RAMPURA BUJURGA-2", "DHARAMPUR", "AHRO-4", "KAPNERI", "UDAYPUR JAGIR", "SHIVNAGAR NEAR RAMNAGAR", "AHRO-3", "BAJAR KALAN-3", "SIHORA", "BIDWA NAGALA", "KHANDIYA", "SIKHADIYAN -4", "KHANPUR JADID-2", "JYOHRA-2", "DHAWANI BUJURG-1", "REHSENA-2", "RATHONDA-1", "CHAMARAN -2", "KURTHIYA", "MUDIYA KHURD", "GODHI-2", "PANWADIYA", "MAHESHPURA", "KHAMRI", "KHATA", "CHANDWA NAGALA", "PATTI BASANTPUR", "DANDIYA NAYAMATGANG", "MAINI", "CHANDPURA JADID", "BAMANPURA-1", "KHEMPUR", "KEWALPUR", "GULADIYA BHAT", "SAHUKARA-1", "LAKHMIPUR", "UDHAMPUR", "BALKHEDA", "KANAKPUR-1", "JIYARAT SHIRIMIYAN-3", "FULSUNGHI", "BAJAR KALAN-1", "NAGRIYA KHURD", "PANJAWNAGAR", "JIWAI JADID", "BACHHIYAI", "GULALPUR", "MANUNAGAR", "KASAMGANJ", "PASHUPURA", "PATNA", "SWALEPUR", "MILAK PIPALSANA", "JASMOLI-2", "AHRO-1", "MUSARFGANJ", "HARJIPUR", "PEGAMWARPUR", "DIBDIBA-1", "DOHAIYA", "PIPALIYA GOPAL-2", "GUJRAILA", "BADPURA", "MANPUR OJHA-1", "JASMOLI-1", "HURMATNAGAR", "BAKNOARI", "KOTRA", "GODHI-1", "TAKIYA -1", "IMAMBADA-3", "MANGADPUR", "GADIYA NASEEMGANJ", "AIMI-1", "NAGLA UDAI-2", "SIKHADIYAN -3", "BAJAR KALAN-4", "HAMIDNAGAR", "AINJANKHEDA", "RATHONDA-2", "CHAKIYA HAYATNAGAR-1", "KOTA ALINAGAR", "TIRAH", "PIPALIYA AHALA", "RAFATPUR", "SAHUKARA-3", "LODHIPUR", "BAJAR KALAN-2", "BARA KHAS", "FAIJGANJ", "MANIHAR KHEDA", "GEHLUIYA", "AMBARPUR", "REHSENA-1", "SILAI BARA GAON-1", "KANAKPUR-2", "SUJANPUR", "ROOPPUR", "SISAUNA-2", "CHAKFERI", "ALINAGAR KOTA", "JIYARAT SHIRIMIYAN-2", "BAKENIYA BHAT", "SHADINAGAR", "FOOLPUR", "ALBANAGALA", "ADILNAGAR", "PIPALIYA NAV", "SIHARI-2", "PIPALIYA VIJAYNAGAR", "BAJAR KALAN-8", "BIDPURA", "RAMNAGARIYA", "HURMATNAGAR TANDA -1", "DIBDIBA-4", "MANSURPUR-2", "SUNARKHEDA", "KOTHA JAGEER", "NAGARIYA KALAN -2", "BIJLI", "HARDUA", "BEGMABAD", "ALIPUR THEKA", "BHATTI TOLA-1", "KASWA RAJPUR", "SADARAKHEDA", "DANKARA", "NARSUA", "DHAWANI BUJURG-2", "BHOLAPUR KADEEM", "SHIV NAGAR NEAR GULADIYA", "GULADIYA BIJLA", "JAURASI", "MANSURPUR-1", "TALMAHAWAR", "JYOHRA-1", "RAMNAGAR", "HARDASPUR", "IMAMBADA -4", "MUDIYA KALAN", "KUIYA", "DHAKI", "DEVRI KHURD", "LALA NAGLA", "SAHUNAGLA", "BHOGPUR", "BERKHEDI FEJABAD", "MUDI", "PIPALIYA RAIJADA", "MUBARAKPUR-1", "SARBARNAGAR", "TAH KHURD", "JAMANPUR", "KAUSHALGANJ", "BALBHADRAPUR", "JIYARAT SHIRIMIYAN-5", "JAGDISHPUR", "SHAMSHABAD", "PAJWA-2", "HUSAINGANJ", "SWAR KHURD-1", "BHAWARKA-1", "MAHTOSH-1", "SHRI NAGAR", "KHANPUR JADID-1", "NAGLI BHAGWANT", "BAHADURGANJ", "LALPUR", "MOHAMMADNAGAR NANKAR", "BHATTI TOLA-3", "PAIPURA", "HALUNAAGAR", "BIDAU", "CHANDPUR", "SUHAG NAGLA", "AAKILPUR", "UMRI", "VISHANPUR JAGEER", "FAJALPUR", "KAYAMGANJ", "SWAR KHURD-2", "DHANELI UTTARI-1", "KHUNTA KHERA", "PSIYAPURA", "KAYASTHAN-1", "BAJAR KALAN-7", "KACHNAL", "BHATPURA TARAN-1", "CHANDPURA KADIM", "KHAMARIYA-2", "KHAJURIYA KALAN", "AIMI-2", "BAHPUR GANGAPUR", "BABOORA", "HAMIDABAD-1", "AADPUR", "MEGHANAGLA JADID", "BAJAR KALAN-6", "GULADIYA TEULA", "FULSUNGHA", "JAGANNATHPUR", "MAJULLANAGAR-3", "HARSONAGLA", "CHANDAYAN", "KALAYANPUR", "SARAY KADEEM", "KHUBIYA NAGALA", "MILAK MOHAMMAD BAKSH", "MAHTOSH-2", "DHALLIYA", "PAJABA", "ABDULLANAGAR", "HURMATNAGAR TANDA -2", "CHAKIYA HAYATNAGAR-2", "MUKRABPUR", "BHUKSAURA", "SINGRA-2", "MADEYAN SENDOALI", "SAHUKARA-6", "SIMARIYA-1", "PAJIYA", "ORANGNAGAR KHEDA", "KHAJURIYA KHURD", "MUNDIYA KALAN", "JIYARAT SHIRIMIYAN-1", "JIYARAT SHIRIMIYAN-6", "GANGAPUR JADID", "RATANPURA", "GAGNPUR", "BAMANPURA-2", "PAJWA-1", "KHATA CHINTAMAN", "DIBDIBA-3", "SIMRA", "SISAUNA-1", "PIPALIYA MEHTO", "MUHAMMADPUR BAS KHADIYA", "BHAWARKA-2", "VISHARATNAGAR-2", "SIRSA", "MENDIPUR", "SAHUKARA-4", "MANPUR OJHA-2", "LAKHIMPUR VISHNU", "KADRIGANJ", "IMAMBADA -5", "SIRASKHERA", "NAGLA UDAI-1", "MADANAGLA", "MUJHIYANA", "TEMRA", "KHAUNDALPUR-2", "LAKHNAKHEDA", "GANGAPUR KALN", "DIBDIBA-2", "UDAYEPUR", "RUTAMPUR", "DHANORA", "PURENIYA JADID", "JANUNAGAR", "MANAONA", "RAJPURA", "MULLAKHERA", "DHIMRI", "PIPALSANA", "DULICHANDPUR", "RAVANAGHASI", "PADPURI", "AKONDA", "SILAI BARA GAON-2", "NARKHEDI", "BHARATPUR", "RAWANALALA &amp; VISHANPURI", "GANGAPUR KADIM-2", "IMAMBADA -1", "KUAKHERA", "MOHAMMADPUR KADIM", "BEDPUR", "RUSTAMPUR", "AGAPUR", "PAREWA", "BIBRA", "JHUNAIYA", "KHAUNDALPUR-1", "KHANPUR KADIM", "SARDHARPUR", "SIKRAURA", "SIKHADIYAN -1", "SENDOALI", "NARAYAN NAGALA", "SAHUKARA-5", "MANPUR OJHA-3", "NARKHEDA-2", "KAGANAGLA", "MALANKHEDA", "CHAINPUR", "INDARPUR", "MILAK MAFI PAIPURA", "BIHAT", "SIHARIYA", "SAKATUBA", "BARADARI", "BHOAJIPURA", "TAKIYA-2", "PRITHVIPUR &amp; CHIDIYAKHEDA", "SIHOR", "KAUTALKHEDA", "HYATNAGAR", "KAJIYAPURA", "INAYATPUR", "CHAMARAN-1", "SENJNA", "GANGAPUR KADIM-1", "KARSAULI", "FAIJNAGAR", "MALKI KAMALPUR", "HAFIJNAGAR", "NAYA GANAO", "SAHUKARA-2", "BHATPURA TARAN-2", "GULAMGANJ", "ANKHERA", "SITAURA", "DHAWNI HASANPUR-2", "SHIVNAGAR NIKAT GANGAPUR", "GAJRAULA", "MIRJAPUR CHAKARPUR", "NIPANIYA-1", "KALYANPUR-2", "MILAK MUNDI", "MANPUR OJHA-4", "SYONDRA", "VISHARATNAGAR-1", "DHAWNI HASANPUR-1", "BHATPURA CHAKRPAN", "SHYAMPUR", "JAFRABAD", "DANKARI", "HARIYA KALAN", "MAJULLANAGAR-2", "BAJAR KALAN-5", "KALYANPUR-1", "TEHRI KHAWJA-1", "DURJANPUR", "BHENSIYA JUALAPUR-2", "KHAMARIYA-1", "AHRO-2", "SIMARIYA-2", "CHAUKONI", "MAJULLANAGAR-4", "TEHRI KHAWJA-2", "PIPALIYA GOPAL", "DIBDIBA-5", "LAKHIMPUR BHIKA", "DALKI", "HIMMATGANJ", "RASULA", "BHENSIYA JUALAPUR-1", "SANKARA", "CHAKARPUR", "NAGARIYA KALAN -1", "MUBARAKPUR", "HAJRATPUR", "KAMUA NAGALA", "TUMADIYA", "ANWARIYA TALIWABAD", "SIKHADIYAN -2", "KAYASTHAN-2", "GADAIYA NIKAT HINGANAGLA", "RASDANDIYA", "PIPALIYA MISRA", "GADA", "DEVRI BUJURG", "IMAMBADA -2", "BERKHEDA", "RAMPURA BUJURGA-1", "NARKHEDA-1", "NIPANIYA-2", "JAGATPUR", "RASULPUR", "DEVIPURA", "SUANAGLA", "ITANGA BERAMNAGAR")
		case "Rampur":
			sections = append(sections, "VIJIYA", "MILAK DYUDI INAYAT RASOOL KHAN", "MADAIRYAN KALI", "BAKENIYA", "KHATKAN PASIYAN V TALAV MASTU KHAN", "CHHIPIYAN MAY BADHEYAN -1", "JVALANGAR C 0 OII0 D0 V LEBAR KALONI-3", "PARTAPUR HARDASPUR", "KOTAVALI KOHANA", "STATION ROAD-1", "GHER SALAVAT KHAN", "MUNDIA KHEDA", "RAJADARA NAVAB GET V MASJID PILU-4", "MASJID JASAUDI PRADHAN", "MAJAR BAGADADI SAHAB", "TAKIYA MUBARAK SHAH KHA", "MASJID KALLASH MAY GUIYA TALAB-2", "BAHAPURA-2", "FUTAMAHAL MAY JAIN MANDIR", "KUNCHA BHAGMAL", "TAKIYA FATEH SHAH KHAN", "AGAPUR-3", "GHER BECHA KHAN", "PUL PUKHTA", "CHARAKH VALI MASJID", "CHAH KHAJAN KHAN", "FILAKHANA JADID", "GHAIR MUNSHIVALA MAY  CHAH SOTIYAN", "MILETRI LAIN-2", "GHAIR MILKIYAN", "DEHALI DARAVAJA MAY GAUKHANA KAIMP", "GHER PURAN SINGH", "CHAHASATAI", "KAKROA-4", "RAJADARA NAVAB GET V MASJID PILU-1", "LAL KABAR-1", "AVAS VIKAS GANGAPUR", "PATTI KALYANPUR", "ZIARAT KHURMEWALI 2", "AJEETPUR-5", "RAIPUR-1", "SATUNE SANG -2", "SARAY SAHADAT YAR KHAN-2", "AGAPUR-6", "MISHRINAGAR JUNUBI", "DUKAN SONA BINNA", "KUNDA 1", "DUNGARPUR", "BAJODI TOLA- 4", "HAJINAGAR", "MANDEYAN BALLU", "GHER MIAN KHAN-1", "BAGICHA AIMNA", "CHAH KHANASAMA", "GHER YUSUF KHAN", "KHAJURIYA", "BELADARAN", "CHHIPIYAN MAY BADHEYAN -2", "DOMAHALA ROD MAY GHER GULAM NASIR KHAN", "KAKROA-2", "JVALANGAR C0 OII0 D0 LEVAR KALONI-6", "AGAPUR-5", "DARAKHT KAITH", "SARAYE JAHANGIR -1", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-17", "MOLANA MOHAMMAD ALI JOHAR MARG MAY GANGAPUR HATS-9", "SADAR SAFAKHANA-1", "CHAH KHAIYATAN", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-15", "MEHEL KUMEDAN", "DUNDAI-1", "STATION ROAD-2", "BAJARIYA JHABBU KHAN MAY SILEYAN-1", "MILETRI LAIN-1", "DHARMASHALA MAY MAI KA THAN", "NAEEM GANJ", "KAKROA-1", "CHIKTI RAMNAGAR-2", "KAKROA-5", "GHER USMAN KHA", "MORI GET-1", "NALAPAR MAY KUCHASAUDAGARAN-1", "VAMNAPURI GET-1", "GHER KALANDAR KHAN-1", "JVALANGAR C0 OII0 D0 LEVAR KALONI-7", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-12", "PATLE V MASJID KHATOTI", "ASTABAL KAIMP MAY THANA GANJ", "FAIZNAGAR", "CHAH JATTA", "AHMAD NAGAR JAGEER-1", "MADEYAN ODEYRAJ", "MAGARMAU", "CHAH SHOR-1", "MANDEYAN RAMI", "GHER MIYAN KHEL", "GHER MOHABBAT KHA", "AJITAPUR-9", "ZIARAT KHURMEWALI-1", "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 2", "TICKET GANJ", "GHER KALANDAR KHAN-2", "BAGICHA CHHOTE MIAN 2", "GHOR SAKHI", "PANVADIYA KVATARS 2", "SARAYE PUKHTA", "NAYA GAON", "DEENPUR", "GHER PIPALWALA-2", "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 1", "FULAVAD", "BAGICHA JOKHIRAM", "TALAB MAJAR BAGADADI SAHAB", "CHAKASHADINGAR", "RELVE STESH N ERIYA", "HAJRATPUR", "LAL KABAR-2", "THOTHAR MAY SABJI MANDI -2", "SHOKAT ALI MARG -3", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-11", "JAULAPUR", "AKAB MOTI MASJID", "BAJARIYA FATEH ALI KHA-1", "GHER BAJ KHA MAY FILAVALAN-3", "BAJODI TOLA-3", "RAJJAD", "PILA TALAB", "PANJABNAGAR-2", "FAIZULLANAGAR-2", "CHAMARPURA", "KHARADIYAN", "JVALANGAR C0 OII0 D0 LEVAR KALONI-9", "CHAH SHOR-2", "MANSUR PUR", "GHER REHMAT KHAN 2", "GHER GAUS MOHAMMAD KHA", "GHER AJIM KHA MAY AKHUNAJADA", "JINA INAYAT KHA", "MASJID MOLANA JAHURUDDIN KHAN 2", "LAL MASJID-2", "MASJID KALE KHAN", "JVALANGAR C0OII0D0 V LEVAR KALONI -8", "GHER KATE BAJ KHAN", "MISTAN GANJ AKAB MISTAN GANJ", "GHER SAIFUDADIN KHAN", "LADORA NARAYANPUR", "KHUSRO BAG", "IMLI JHULEWALI", "BHAIYANGALA", "NAUGAVAN-1", "FAIZULLANAGAR-1", "SIKARAUL", "GHAIR AJAM KHAN", "BERIYAN AKAB MAHAL SARAY 2", "THANATIN.5", "GHER SARABDAL KHAN", "SHOKAT ALI MARG -1", "BILASAPUR GATE", "MAJAR CHUPASHAH MIYA", "GHER HASSAN KHA", "MASJID MOLANA JAHURUDDIN  KHAN-1", "PANJABIYAN", "TASHKA -1", "MILAK NIBBI SINGH", "JVALANGAR AI0 D0 OII0 0 LE0KA0-5", "KUNCHA KASSAVAKHANA", "CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-1", "AJEETPUR-6", "THANATIN-1", "KUNCHA FIRANGAN", "MANDHOLI-1", "GHER MIAN KHAN -2", "AJEETPUR-8", "THEGA", "VENAJIRPURA  URF GHATAMPUR 2", "BAG PUKHTA", "WAZEERNAGAR", "AFARIDIYAN-1", "CHIRAN", "KILA FORT", "GHER BAJ KHA MAY FILAVALAN-2", "AHMAD NAGAR JAGEER-2", "VISHNU VIHAR-1", "BAMNAPURI -1", "BAGICHA GAJI MUJFAFAR KHA", "MADAIYAN JAULAPUR", "TAKIYA MEMARAN", "AHMAD NAGAR KHEDA", "MAJAR VAJIR ALI SAHAB", "BHANDPURA-2", "AJEETPUR-4", "KUNCHA NAKKALAN", "MASJID DAROGA MAHABUB JAN", "IMLI BATANIYA", "PAHADI GET", "GHER SAIFUDADIN KHAN 2", "KUNCHA LANGARAKHANA", "GHER TONGA-3", "MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HATS-7", "MILAK SHADI NAGAR", "KALGHAR-1", "SHAZAD NAGAR", "IMLI ASMAT KHAN", "LADORI", "GHER MUNAVVAR KHAN", "GHAIR MULLA MALUK", "MANDOLI-2", "SHADINGAR NI0AHAMADNAGAR JAGIR", "GHER JIYAUNNAVI KHA", "AJEETPUR-10", "BAJARIYA KADHU", "GHER INAYAT KHAN", "GHER BAJ KHA MAY FILAVALAN-1", "THAKURDUWARA", "TALAB MULLA AIRAM", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-14", "PATTI TOLA", "TALAB KELE VALA", "LALPUR PATTI KHURD", "RAMNATH KALONI MAY GOVAN ASPATAL", "NAWAB GAUNJ JUNUBI", "FADASHEKH MAY BANS MANDI", "TALAB NIHALUDDIN MAY GHER HASSAN KHAN", "KATRA JALALUDIN MAY GHER HASSAN KHAN", "BHANDPURA-1", "GHER MULLA GAINRAT", "KALKATTA", "AJEETPUR-7", "PANJABNAGAR-3", "NAUGAVAN-2", "KAKKROA-3", "BUDHIYA KI DUKAN", "BAJODI TOLA -2", "NALAPAR MAY KUCHA SAUDAGARAN-2", "SAIDNAGAR JANUBI", "AJEETPUR-1", "DANYAPUR SHANKARPUR", "JHANDA BADE PIR SAHAB-1", "NASRATNAGAR", "GHER REHMAT KHAN 1", "GAUKHANA KADIM", "RASOOLPUR", "ANGUR VALI MASJID", "HARIYALE", "SHOKAT ALI MARG -5", "JVALANGAR C0 OII0 D0 V LEBAR KALONI-2", "BAHAPURI", "SARAY DARAVAJA-1", "MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HARTS 8", "KALRAKH", "HASNAPUR", "PESH KOTAVALI", "CHAH JADIYAN V JARGARAN", "GHAIR SHAH MOHAMMAD KHAN", "THOTHAR MAY SABZI MANDI - 1", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-16", "KHSNDSAR KOHNA", "GHAIR GILJIYAN", "AKHUN KHE LAN", "GOHARPUR", "CHAH KESRA SIH", "KUNDA 3", "SHUTRAKHANA MAY JYOTISHYAN-2", "BAJARIYA FATEH ALI KHA 2", "SHOKAT ALI MARG -4", "BAJAR SAFDARGANJ MAY PAN DARIBA", "DABKA-1", "VAMNAPURI GET-2", "ALINAGAR JUNUBI", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-13", "BERIYAN AKAB MAHAL SARAY 1", "GHER MIAN KHAN -3", "MAJAR SHAH DARAGAHI SAHAB", "ATTA ALLANUR", "VISHNU VIHAR-2", "AFARIDIYAN-2", "TATATVALAN MAY MAJAR TATASHAH MIYAN", "BAJARIYA MULLA JARIF", "KOTAVALAN", "REHAT GANJ", "GHER MEERBAZ KHAN MEY GHER HASSN KHAN-1", "MANKARA", "BULAND KVATARS", "MORI GET-2", "KOTHI KHAS BAG-1", "JYARAT HALKE VALI", "MILAK ICHHCARAM", "SHAHABAD GET-1", "KATAKUIYA MAY FILAKHANA KADIM-1", "BAJODI TOLA -1", "MASJID KALLASH MAY GUIYA TALAB-1", "GUJAR TOLA", "AGAPUR-2", "CHIKTI RAMNAGAR-1", "AGAPUR-1", "TOPAKHANA ROAD", "SHAHABAD GET-2", "GHER TONGA-1", "AKAB POST AFIS", "FATTEPUR", "SARAYE JAHANGIR -2", "BAMNAPURI -2", "GHAIR KUTVI MIYAN", "KUNCHA NALVANDAN", "BANDUKACHIYAN", "THANATIN.4", "THANATIN.6", "MANDEYAN SHADI-1", "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-10", "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 3", "AFZALPUR PATTI", "GHER NAJJU KHAN", "BANGALA KASAM ALI KHAN", "SARAY KALAN-1", "RAJA RAMPUR", "SHOKAT ALI MARG -2", "KUNCHA SEEKALGARANN", "SHIV VIHAR", "MANDEYAN SHADI-2", "AFISARS BANGALE", "PASIYAPURA JANOOBI", "BISRA-1", "MASJID KAITHAVALI", "LALPUR PATTI QADAR KHAN", "SADAR SAFAKHANA -3", "GHAIR ALLI KHA", "KALGHAR-2", "GHER FATEH KHAN", "PIPAL TOLA", "BISRI", "SHUTRAKHANA MAY JYOTISHYAN-1", "MILAK SIKARAUL", "MAHAMUD PUR", "DUDAI-2", "GHER TONGA-2", "GHAIR SHAIRFUDDIN KHAN", "GHAIR DARIYA KHA", "LALPUR PATTI SAEED GANJ", "LAL MASJID-1", "CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-2", "BHAISAKHANA", "BAJARIYA JHABBU KHAN MAY SILEYAN-2", "GHER FARUKHSHAH KHAN", "MEHANDI NAGAR JUNUBI", "MILAK TAHABBAR ALI KHAN", "KITAPLAI", "RAJADARA NAVAB GET V MASJID PILU-3", "GANJAKADIM -1", "CHAHASHIRI", "KUCHA DEVIDAS", "SARAY KALAN-2", "JHANDA BADE PIR SAHAB-2", "BAJARIYA KHANASAMA", "SATUNE SANG -4", "BEHPURA- 1", "PULIS LAIN", "GHER JANAS KHAN", "KOTHI KHAS BAG-2", "THANATIN.3", "RAIPUR-2", "SATUNE SANG -5", "RAJADARA NAVAB GET V MASJID PILU-2", "GHER MEERBAZ KHAN MEY GHER HASSN KHAN-2", "AHMADNAGAR THEGA", "SARAY SAHADAT YAR KHAN-1", "SATUNE SANG -3", "KATAKUIYA MAY FILAKHANA KADIM-2", "AJEETPUR-2", "THANATIN.2", "GHAIR SAIFAL KHAN", "PANVADIYA KVATARS-1", "PANJABNAGAR-1", "VENAJIRPURA URF GHATAMPUR-1", "CHIKNA", "BAJAR ABDULLA GANJ", "GHER MUHAB KHA", "GHAIR BAKHSHI", "BAJODI TOLA- 5", "GHER CHANDAN KHAN MAY JEL K VATARS", "STATION ROAD-3 MAY GOVAN COLONY", "KUCHA PARAMESHVARI DAS", "GANJ KADIM -2", "GHER MUBARAK SHAH KHAN", "BANGALA AJAD KHA", "GHER PIPALWALA-1", "KUNCHA LALA MIYAN", "KALYANPUR", "GHANASHYAMAPUR", "MADRASA ALIYA", "TUMADIYA-1", "KUNCHA KAJI", "CHAH INCHHARAM", "TAKIYA SARVAR SHAH KHAN", "DABKA-2", "PIPLATOLA NIKAT GUJRATOLA-1", "GHER ALAF KHAN", "TALAB CHAMARAN", "CHAH MOTE KALLAN", "JVALANGAR C0 OII0 D 0 LEVAR KALONI-1", "PIPLATOLA NIKAT GUJRATOLA 2", "BAGICHA CHHOTE MIAN-1", "MADRASA KOHANA", "ANYA FAIKTARI ERIYA", "VISHNU VIHAR-3", "SARAY DARAVAJA-2", "TASHKA- 2", "ANGURI BAG", "ABBAS NAGAR", "PANVADIYA", "SHUTRAKHANA MAY JYOTISHYAN-3", "MORI GET-3", "AKHADA MALLI KHAN", "DARAKHT KHINNI", "TUMADIYA-2", "KUNCHA ATISH BAJAN", "BESRA-2", "LALPUR PATTI KUNDAN", "SADAR SAFAKHANA-2", "MAU. BARADRI MAHAMUD KHAN", "GHER MARDAN KHAN", "RAJADARA NAVAB GET V MASJID PILU-5", "BAJAR NASARULLA KHAN", "NYU AVAS VIKAS", "SATUNE SANG -1", "AJEETPUR-3", "JVALANGAR C 0 OII0 D0 V LEBAR KALONI-4", "AGAPUR-4", "BAJARIYA HIMMAT KHAN")
		case "Milak":
			sections = append(sections, "SAIFNI -11", "MATHURAPUR 01", "SAIFNI -10", "DAKURIYA", "KAMORA-1", "SADAT 02", "MADARPUR", "LALWARA -1", "BOODPUR", "NABIGANJ JADID", "KHATA KALAN-2", "GALPURA", "MITTARPUR AHROULA 02", "LEHTORA", "CHITOUNI-1", "NEW SAGARPUR, BISOULI -02", "DHAKIYA 03", "NABBA NAGLA", "KESHARPUR NEAR RAYPUR", "SHEHZADNAGAR", "RUSTAMPUR", "MASJID QAZI", "MILAK MAJBOOTA", "BHARATPUR", "PATARIYA", "AWADI RAILWAY STATION", "RUSTAMNAGAR", "SALABATPUR", "MOHAMMAD NAGAR KHUTIA", "BARDA GAUN 02", "JADOANPUR-2", "MANGOLI 01", "MILAK ASDULLAPUR-2", "MILAK BANKERALI", "KARAITHI -1", "PATWAI -2", "KASSHABAN", "NAVADIYA", "MADHUPARI -01", "KARAITHI-2", "GADMAR PATTI TIKA SINGH", "KRIMCHA-1", "PACHTOUR", "KHATA KALAN-1", "NARKHERA", "BARKHERA", "CHOUKONI", "KAMORA-2", "TALIBABAD", "SARKADA", "HIMMATPUR 02", "GAJEJJA", "MITTARPUR AHROULA 01", "KHATA KALAN-4", "JIWAI KADIM-1", "LADPUR", "NARKHERI", "BICHPURI SANGRAMPUR", "PATWAI -1", "DHAMORA-3", "DHNORA", "DANIAPUR", "MADHUKAR 01", "CHANDPURA KALAN-1", "CHAMRA", "NISVI", "KAMRUDDIN NAGAR", "NRENDARPUR 01", "AHMADNAGAR NIKAT KUNDANPUR", "NISWA", "FARRAASHAN 04", "AINCHORA-2", "DUGANPUR", "BRIJPUR", "PARAM-1", "BARBALAN", "NEW SAGARPUR, BISOULI -01", "DHOLSAR", "SAIFNI -09", "CHANDPURA KHURD", "SAIFNI -03", "MOHAMMADNAGAR", "BARA NIKAT GAJEJA", "DHANOURA", "ASHOKPUR", "CHITOUNI-2", "PARAM-2", "MOHAMMADGUNJ", "ROARA KHURD-2", "SAIFNI -02", "NANDGAUN", "ATAI NAGAR", "NRENDARPUR 02", "BEHTA", "KHARSOUL 01", "RAVANA -02", "PATIYA", "ISAKHERA", "HAREYYA KA MAJHRA", "KRIPYA HAPPU", "KRIMCHA-2", "LOHAPATTI BHAGIRATH", "BEDAN", "KOOP -2", "PRANPUR", "BHENSODI-3", "MILAK ASDULLAPUR-4", "NAGARIYA - 2", "TAJPUR LAKHAN", "DHAKIYA 1", "BUDHPUR", "DHANELI POORVI-2", "DIVIYANAGLA", "NAANKAR", "PATWAI 04", "DEVIPURA", "KEORAR-1", "AINCHORA-1", "AHMAD NAGAR NIKAT JIWAI KADIM", "AFGANAN 01", "PROTA -2", "MAHEVA", "MATWALI -2", "MADHUKAR 02", "SAIFNI -05", "SHAHDARA", "MEHNDIPUR", "KHARSOUL 02", "MIRJAPUR PARAM", "MAU", "KARIMGANJ", "SURAJPUR", "MILAK TURKHEDA", "KHANDELI 02", "REHPURA", "MDEAYAN GOR", "YUSUF NAGAR", "REVERI KALAN 01", "MADAIYAN TULSI", "JALIF NAGLA-3", "BHUDASI 02", "BAMANPURI", "LALWARA 02", "MILAK GULAM ALI", "PAIGAMBARPUR", "POHAPPUR", "JANAKPUR", "PAIGUPURA", "PALPURA", "CHAKARPUR BHUD", "KRIPYA PANDE", "NEW SAGARPUR, BISOULI -03", "JILEDARAN 02", "MOHAMMADPUR JADID", "KHUTIA", "BEGMABAD NIKAT KRIMCHA", "BANDAR 02", "LOHAPATTI BHOLANATH-2", "KHIRKA", "LOHAPATTI BHOLANATH-1", "KANOONGOYAN 01", "NADNAU", "BANDAR 01", "ASALATPUR", "BHOANAKPUR", "NAIMGANJ", "OSI 02", "MADAIYAN VADE", "BARUA -02", "NAGARIYA-1", "SAHAVIAN KHURD", "PASUPURA", "BEGMABAD NIKAT ASHOKPUR", "JALIF NAGLA-2", "MATHURAPUR -2", "KIRA 03", "KARANPUR", "BAJHANPUR-02", "TAJPUR BEHTA", "MEHNDINAGAR", "PURENIYA KALAN", "GADMARPATTI TIKA SINGH", "MDEYAN BUDHPUR", "KANOONGOYAN 02", "FARRAASHAN 01", "BIJPURI LALJI", "BHITAR GAUN 02", "SAIFNI -04", "LASHKAR GANJ", "TARAUA", "GHARAMPUR", "RIVDI KHURD", "JADONPUR-1", "JMALPUR", "GULARIYA KALAN", "DHAKIYA 02", "FAKEERGANJ URF GHOSIPURA", "MATWALI --1", "KAMALPUR", "MAJHRA", "BHITAR GAUN 01", "KAREEMGUNJ", "KHATPUR", "BEHTRA", "TANDA 02", "NABABGUNJ", "MIAN GANJ", "RASOOLPUR", "BHOURKI JADID", "RAMPURA", "MILAK NASIRABAD-2", "NADNA", "FARRAASHAN 02", "BARDA GAUN 03", "MANSOORPUR", "SHAHPUR DEV 01", "SADAT 01", "001 CHAKFERI", "ANWA 02", "MADEYAM DHOLSAR", "UNCHA GAUN 02", "JILEDARAN 01", "TELIPURA", "AHMAD NAGAR NEAR BHAGWATIPUR", "LODHIPUR", "RAVANA -01", "PIPLA SHIV NAGAR-2", "CHANDPURA KALAN -02", "SAIFNI -01", "ROARA KHURD-3", "NAWABNAGAR NEAR PATWAI", "2 AMEER ALI KHAN 1", "KASHAVNAGLA", "BALUPURA", "BHOURKA", "BANKRABAD", "BARUA -01", "MADHUPARI -02", "KHUDANAGAR", "DOHARIYA", "GUJRAILA", "KHERA", "SHAHPUR DEV 02", "AHMADGANJ URF MUKATPUR", "KARINGA", "PAGAMBERPUR", "PARAM-3", "KALKATTA", "JAITOLI 01", "GOUSAMPUR -2", "BHOUPATPUR", "UNCHA GAUN 01", "MILAK FEJULLANAGAR", "PIPALIYA BHOJ", "MAMURPUR BADOULI", "PAROUTA 01", "KHANDELI", "DHARAMPURA", "HUSAINGANJ NEAR JANAKPUR", "SAIFNI -07", "RAJOURA", "BARDA GAUN 01", "CHATARPUR", "MOTIPURA", "JATPURA", "HIMMATPUR 01", "CHAKARPUR KADEEM", "MOHALIYA", "ANWA 01", "KRIMCHA-3", "AMEER ALI KHAN 02", "VEERPUR", "ROARA KHURD-1", "TAJPUR", "BHUDASI 01", "DIVAAN", "SUJAYATGUNJ", "HATA KALAN-3", "GANGAPURSHARKI", "SHEKHUPURA", "THIRIYA JADID", "JAITOLI 02", "BHENSORI-1", "SADDIKNAGAR", "JALIF NAGLA-1", "KISHNPUR", "HAKIMAAN 01", "AFGANAN 02", "KOOP -3", "MILAK ASDULLAPUR-3", "CHAMROL", "HUSHIANGUNJ NEAR MADARPUR", "CHAKARPUR KADEEM 02", "JAGESAR", "GEHNI", "KHANJIPURA", "RAWANIPATTI JBHAHAR", "BHENSORI-2", "MADHUKAR-3", "NAVIGUJ KADEEM", "SOOPA", "KHANUPURA", "RAIPUR-01", "KOOP-1", "NAVAVPURA", "DHAMORA-1", "GOUSAMPUR 01", "UDAIPUR JAGIR", "TURKHERA", "RAIPUR-02", "HAKIMAAN 02", "JAIDOLI", "ROARA KALAN-2", "KESARPUR NEAR KAREEMGANJ", "TANDA 01", "DUABAT", "BHAGWATIPUR", "KIRA 01", "MANGOLI 02", "CHAKASHI", "PURENIYA KHURD", "ZAHIDPUR", "NASRAT NAGAR", "KANOONGOYAN 03", "DHURIYAI", "NANKAR", "FARRAASHAN 03", "KEORAR-2", "DHARMPUR", "MADIYAN JHAU", "DHAKIYA 04", "PATWAI -3", "OSI 01", "NISHVI", "VIKRAMPUR", "SAHAVIYAN KALAN", "ROARA KALAN-1", "PURENA", "SARAY MAHESH -2", "AHAMDABAD", "SUHAVA", "KIRA 02", "THIRIA VISHNU", "MAHUNAGAR 02", "PIPALWALA", "NADARGANJ", "ROOPPUR", "SAIFNI -08", "SAIFNI -06", "CHANDPURA SALIS", "NORAHA", "MILAK ASDULLAPUR-5", "BHOURKI KADEEM", "SARAI IMAM", "PARAM-4", "SARAY MAHESH -1", "RAWANIPATTI UDA", "MAHUNAGAR 01", "THIRIYA SALEPUR", "MEERAPUR", "MEGHANAGLA KADIM", "REVERI KALAN 02", "MILAK NASIRABAD-1", "MILAK ASDULLAPUR-1", "SAEED NAGAR", "MILAK NASIRABAD-3", "JIWAI KADIM-2", "RAMNAGAR", "NIYAMAT NAGAR", "MATHURAPUR -3", "ALAFGUNJ", "SOHANA", "GOVINDPURA", "DHANELI POORVI-1", "PATTI FAZILABAD", "DHAMORA-2", "BAJHANPUR-01", "NALAPAR", "BHAGBANTPUR", "PIPLA SHIVNAGAR-1", "HAREYYA")
		}
	}
	return sections
}

func getAcs(districts []string) []string {
	var acs []string

	for _, district := range districts {
		switch district {
		case "Moradabad":
			acs = append(acs,
				"Kanth",
				"Thakurdwara",
				"Moradabad Rural",
				"Moradabad Nagar",
				"Kundarki",
				"Bilari")
		case "Rampur":
			acs = append(acs,
				"Suar",
				"Chamraua",
				"Bilaspur",
				"Rampur",
				"Milak")
		case "Bijnor":
			acs = append(acs,
				"Najibabad",
				"Nagina",
				"Barhapur",
				"Dhampur",
				"Nehtaur",
				"Bijnor",
				"Chandpur",
				"Noorpur")
		}
	}
	return acs
}

func getCommonItems(list1 []string, list2 []string) []string {
	var items []string

	for _, item := range list1 {
		if contains(list2, item) {
			items = append(items, item)
		}
	}

	return items
}

func contains(sliceString []string, item string) bool {
	for _, val := range sliceString {
		if val == item {
			return true
		}
	}
	return false
}

func (e *ElectionController) UpdateVoter() {
	var (
		num  int64
		user []*modelVoters.Account
		err  error
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for each and every district
	qsRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	//qsBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	//qsHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	condVoterId := orm.NewCondition()

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	vote := new(modelVoters.UpdateVote)

	err = json.Unmarshal(inputJson, &vote)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Apply filters for each query string
	// Voter Id
	for _, voterId := range vote.VoterID {
		if voterId > 0 {
			condVoterId = condVoterId.Or("Voter_id__exact", voterId)
		}
	}

	if vote.District == "Moradabad" {
		qsMoradabad = qsMoradabad.SetCond(condVoterId)
		if vote.Vote == 1 || vote.Vote == 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
			"vote": vote.Vote,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(vote.Email)) > 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
			"email": vote.Email,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if vote.MobileNo > 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
			"mobile_no": vote.MobileNo,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	if vote.District == "Rampur" {
		qsRampur = qsRampur.SetCond(condVoterId)
		if vote.Vote == 1 || vote.Vote == 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
			"vote": vote.Vote,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(vote.Email)) > 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
			"email": vote.Email,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if vote.MobileNo > 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
			"mobile_no": vote.MobileNo,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
		
	}

	if vote.District == "Bijnor" {
		qsBijnor = qsBijnor.SetCond(condVoterId)
		if vote.Vote == 1 || vote.Vote == 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
			"vote": vote.Vote,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(vote.Email)) > 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
			"email": vote.Email,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the email.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if vote.MobileNo > 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
			"mobile_no": vote.MobileNo,
			})
			if err != nil {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the mobile number.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
	}

	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = fmt.Sprintf("The vote value has been set to %d.", vote.Vote)
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}

/*
func (e *ElectionController) CreateTask() {
	var (
		num  int64
		user []*modelVoters.Account
		err  error
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for each and every district
	qsRampur := o.QueryTable(modelVoters.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(modelVoters.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(modelVoters.GetTableName("Bijnor"))
	//qsBangalore := o.QueryTable(modelVoters.GetTableName("Bangalore"))
	//qsHubli := o.QueryTable(modelVoters.GetTableName("Hubli"))

	condVoterId := orm.NewCondition()

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	vote := new(modelVoters.Vote)

	err = json.Unmarshal(inputJson, &vote)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Apply filters for each query string
	// Voter Id
	for _, voterId := range vote.VoterID {
		if voterId > 0 {
			condVoterId = condVoterId.Or("Voter_id__exact", voterId)
		}
	}

	if vote.District == "Moradabad" {
		qsMoradabad = qsMoradabad.SetCond(condVoterId)
		updatedRows, err := qsMoradabad.Update(orm.Params{
			"vote": vote.Vote,
		})
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		if updatedRows < 1 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = "The voter id(s) provided is/are not valid."
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	if vote.District == "Rampur" {
		qsRampur = qsRampur.SetCond(condVoterId)
		qsRampur = qsRampur.SetCond(condVoterId)
		updatedRows, err := qsRampur.Update(orm.Params{
			"vote": vote.Vote,
		})
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		if updatedRows < 1 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = "The voter id(s) provided is/are not valid."
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	if vote.District == "Bijnor" {
		qsBijnor = qsBijnor.SetCond(condVoterId)
		qsBijnor = qsBijnor.SetCond(condVoterId)
		qsBijnor = qsBijnor.SetCond(condVoterId)
		updatedRows, err := qsBijnor.Update(orm.Params{
			"vote": vote.Vote,
		})
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		if updatedRows < 1 {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Unable to update the vote value.")
			responseStatus.Error = "The voter id(s) provided is/are not valid."
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	}

	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = fmt.Sprintf("The vote value has been set to %d.", vote.Vote)
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}
*/