/*
   GET VOTERS
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/voters

   GET STATISTIC
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' http://107.178.208.219:80/api/statistic

   GET STATISTICS
   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[19,20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://107.178.208.219:80/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://107.178.208.219:80/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":["Muslim"],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]},"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}}' http://107.178.208.219:80/api/statistics

   GET OTP
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9564783954}' http://107.178.208.219:80/api/otp

   Register
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9343352734, "otp":23435}' http://107.178.208.219:80/api/register

   Get List
   curl -X POST -H "Content-Type: application/json" -d '{"districts": ["Moradabad"], "acs":["Kanth", "Bilari"]}' http://107.178.208.219:80/api/list

   Set Voter
   curl -X POST -H "Content-Type: application/json" -d '{"district":"Moradabad", "voter_id": [12345,20045], "vote":1, "email":"example@example.com","mobile_no":9456732819,"image":"jsdjsd22ksndsndsnk22knknlcxx"}' http://107.178.208.219:80/api/voter

   Read Json
   curl -X POST -H "Content-Type: application/json" -d @json/data.json http://107.178.208.219:80/api/read/json
*/

package controllers

import (
	token "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/mail"
	"net/smtp"
	"sort"
	"strconv"
	"strings"
	"time"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/formattime"
	"github.com/Baligul/election/logs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/craigmj/gototp"
	_ "github.com/lib/pq"
	"github.com/scorredoira/email"
)

func init() {
	orm.RegisterModel(new(modelVoters.Voter_19),
		new(modelVoters.Voter_20),
		new(modelVoters.Voter_21))
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
		user                 []*modelAccounts.Account
		err                  error
	)

	limit, _ := e.GetInt("limit")
	offset, _ := e.GetInt("offset")
	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: GET/POST /api/voters, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: Mobile number or token is wrong.")
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	exist := qsAccount.Filter("Mobile_no__exact", mobileNo).Exist()
	if !exist {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: Mobile number does not exists.")
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: Token is not correct")
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: Cannot find account")
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Get Voters API: "+err.Error())
	}

	if limit == 0 {
		limit = -1
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelVoters.Query)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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

	if condAcNumber != nil && !condAcNumber.IsEmpty() {
		cond = condAcNumber
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

	if condVoterId != nil && !condVoterId.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condVoterId)
		} else {
			cond = condVoterId
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

	qsRampur = qsRampur.SetCond(cond)
	qsRampur = qsRampur.OrderBy("Name_english")
	qsMoradabad = qsMoradabad.SetCond(cond)
	qsMoradabad = qsMoradabad.OrderBy("Name_english")
	qsBijnor = qsBijnor.SetCond(cond)
	qsBijnor = qsBijnor.OrderBy("Name_english")

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
			if state == 27 {
				votersCountRampur, _ = qsRampur.Count()
				_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
				if err != nil {
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
		sort.Sort(modelVoters.ByName(voters.Voters))
		e.Data["json"] = voters
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = "No voters found with this criteria."
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Voters API: "+err.Error())
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
		dalitVotersCount                 int64
		OthersVotersCount                int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		dalitMaleVotersCount             int64
		dalitFemaleVotersCount           int64
		OthersMaleVotersCount            int64
		OthersFemaleVotersCount          int64
		muslimVotersP                    float64
		dalitVotersP                     float64
		OthersVotersP                    float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		dalitMaleVotersP                 float64
		dalitFemaleVotersP               float64
		OthersMaleVotersP                float64
		OthersFemaleVotersP              float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		dalitVotersCountRampur           int64
		dalitVotersCountMoradabad        int64
		dalitVotersCountBijnor           int64
		dalitVotersCountBangalore        int64
		dalitVotersCountHubli            int64
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
		dalitMaleVotersCountRampur       int64
		dalitMaleVotersCountMoradabad    int64
		dalitMaleVotersCountBijnor       int64
		dalitMaleVotersCountBangalore    int64
		dalitMaleVotersCountHubli        int64
		dalitFemaleVotersCountRampur     int64
		dalitFemaleVotersCountMoradabad  int64
		dalitFemaleVotersCountBijnor     int64
		dalitFemaleVotersCountBangalore  int64
		dalitFemaleVotersCountHubli      int64
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
		user                             []*modelAccounts.Account
		err                              error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/statistic, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistic API: "+err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Get Statistic API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelVoters.Query)
	statistic := new(modelVoters.Statistic)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistic API: "+err.Error())
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

	if condVote != nil && !condVote.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condVote)
		} else {
			cond = condVote
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

	if condAge != nil && !condAge.IsEmpty() {
		if cond != nil && !cond.IsEmpty() {
			cond = cond.AndCond(condAge)
		} else {
			cond = condAge
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
				dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Dalit").Count()
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
				dalitMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
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
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCount > 0 {
		statistic.Total = votersCount

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistic.Muslim = muslimVotersCount

		dalitVotersCount = dalitVotersCountRampur + dalitVotersCountMoradabad + dalitVotersCountBijnor + dalitVotersCountBangalore + dalitVotersCountHubli
		statistic.Dalit = dalitVotersCount

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

		dalitMaleVotersCount = dalitMaleVotersCountRampur + dalitMaleVotersCountMoradabad + dalitMaleVotersCountBijnor + dalitMaleVotersCountBangalore + dalitMaleVotersCountHubli
		statistic.DalitMale = dalitMaleVotersCount

		dalitFemaleVotersCount = dalitFemaleVotersCountRampur + dalitFemaleVotersCountMoradabad + dalitFemaleVotersCountBijnor + dalitFemaleVotersCountBangalore + dalitFemaleVotersCountHubli
		statistic.DalitFemale = dalitFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistic.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistic.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCount)) * 100
		statistic.MuslimP = muslimVotersP

		dalitVotersP = (float64(dalitVotersCount) / float64(votersCount)) * 100
		statistic.DalitP = dalitVotersP

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

		dalitMaleVotersP = (float64(dalitMaleVotersCount) / float64(votersCount)) * 100
		statistic.DalitMaleP = dalitMaleVotersP

		dalitFemaleVotersP = (float64(dalitFemaleVotersCount) / float64(votersCount)) * 100
		statistic.DalitFemaleP = dalitFemaleVotersP

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
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistic API: "+err.Error())
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
		dalitVotersCount                 int64
		OthersVotersCount                int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		dalitMaleVotersCount             int64
		dalitFemaleVotersCount           int64
		OthersMaleVotersCount            int64
		OthersFemaleVotersCount          int64
		muslimVotersP                    float64
		dalitVotersP                     float64
		OthersVotersP                    float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		dalitMaleVotersP                 float64
		dalitFemaleVotersP               float64
		OthersMaleVotersP                float64
		OthersFemaleVotersP              float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		dalitVotersCountRampur           int64
		dalitVotersCountMoradabad        int64
		dalitVotersCountBijnor           int64
		dalitVotersCountBangalore        int64
		dalitVotersCountHubli            int64
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
		dalitMaleVotersCountRampur       int64
		dalitMaleVotersCountMoradabad    int64
		dalitMaleVotersCountBijnor       int64
		dalitMaleVotersCountBangalore    int64
		dalitMaleVotersCountHubli        int64
		dalitFemaleVotersCountRampur     int64
		dalitFemaleVotersCountMoradabad  int64
		dalitFemaleVotersCountBijnor     int64
		dalitFemaleVotersCountBangalore  int64
		dalitFemaleVotersCountHubli      int64
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
		user                             []*modelAccounts.Account
		err                              error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/statistics, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistics API: "+err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Get Statistics API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	queries := new(modelVoters.Queries)
	statistics := new(modelVoters.Statistics)

	err = json.Unmarshal(inputJson, &queries)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistics API: "+err.Error())
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
	qsScopeBijnor = qsScopeBijnor.SetCond(condScope)

	// CALCULATIONS FOR QUERY
	// Get voters for each state
	if len(queries.Query.DistrictNameEnglish) == 0 && len(queries.Query.DistrictNameHindi) == 0 && len(queries.Query.DistrictNumber) == 0 {
		for _, state := range queries.Query.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsRampur.Count()
				votersCountMoradabad, err = qsMoradabad.Count()
				muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				votersCountBijnor, err = qsBijnor.Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Dalit").Count()
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
				dalitMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
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
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCountQuery = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountQuery > 0 {
		statistics.Query.Total = votersCountQuery

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Query.Muslim = muslimVotersCount

		dalitVotersCount = dalitVotersCountRampur + dalitVotersCountMoradabad + dalitVotersCountBijnor + dalitVotersCountBangalore + dalitVotersCountHubli
		statistics.Query.Dalit = dalitVotersCount

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

		dalitMaleVotersCount = dalitMaleVotersCountRampur + dalitMaleVotersCountMoradabad + dalitMaleVotersCountBijnor + dalitMaleVotersCountBangalore + dalitMaleVotersCountHubli
		statistics.Query.DalitMale = dalitMaleVotersCount

		dalitFemaleVotersCount = dalitFemaleVotersCountRampur + dalitFemaleVotersCountMoradabad + dalitFemaleVotersCountBijnor + dalitFemaleVotersCountBangalore + dalitFemaleVotersCountHubli
		statistics.Query.DalitFemale = dalitFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistics.Query.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistics.Query.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimP = muslimVotersP

		dalitVotersP = (float64(dalitVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.DalitP = dalitVotersP

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

		dalitMaleVotersP = (float64(dalitMaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.DalitMaleP = dalitMaleVotersP

		dalitFemaleVotersP = (float64(dalitFemaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.DalitFemaleP = dalitFemaleVotersP

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
	dalitVotersCountRampur = 0
	dalitVotersCountMoradabad = 0
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
	dalitMaleVotersCountRampur = 0
	dalitMaleVotersCountMoradabad = 0
	dalitFemaleVotersCountRampur = 0
	dalitFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	dalitVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	dalitMaleVotersCountBijnor = 0
	dalitFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	dalitVotersCountBangalore = 0
	dalitVotersCountHubli = 0
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
	dalitMaleVotersCountBangalore = 0
	dalitMaleVotersCountHubli = 0
	dalitFemaleVotersCountBangalore = 0
	dalitFemaleVotersCountHubli = 0
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
				dalitVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
				OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
				votersCountBijnor, err = qsScopeBijnor.Count()
				muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Dalit").Count()
				OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
				maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				dalitMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
				OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsScopeBangalore.Count()
				votersCountHubli, err = qsScopeHubli.Count()
				muslimVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Muslim").Count()
				dalitVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Dalit").Count()
				dalitVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Dalit").Count()
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
				dalitMaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitMaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
				dalitFemaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
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
			dalitVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Scope.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			dalitVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Dalit").Count()
			OthersVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Others").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			dalitMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Dalit").Count()
			dalitFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Dalit").Count()
			OthersMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Others").Count()
			OthersFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Others").Count()
		}
	}

	votersCountScope = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountScope > 0 {
		statistics.Scope.Total = votersCountScope

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Scope.Muslim = muslimVotersCount

		dalitVotersCount = dalitVotersCountRampur + dalitVotersCountMoradabad + dalitVotersCountBijnor + dalitVotersCountBangalore + dalitVotersCountHubli
		statistics.Scope.Dalit = dalitVotersCount

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

		dalitMaleVotersCount = dalitMaleVotersCountRampur + dalitMaleVotersCountMoradabad + dalitMaleVotersCountBijnor + dalitMaleVotersCountBangalore + dalitMaleVotersCountHubli
		statistics.Scope.DalitMale = dalitMaleVotersCount

		dalitFemaleVotersCount = dalitFemaleVotersCountRampur + dalitFemaleVotersCountMoradabad + dalitFemaleVotersCountBijnor + dalitFemaleVotersCountBangalore + dalitFemaleVotersCountHubli
		statistics.Scope.DalitFemale = dalitFemaleVotersCount

		OthersMaleVotersCount = OthersMaleVotersCountRampur + OthersMaleVotersCountMoradabad + OthersMaleVotersCountBijnor + OthersMaleVotersCountBangalore + OthersMaleVotersCountHubli
		statistics.Scope.OthersMale = OthersMaleVotersCount

		OthersFemaleVotersCount = OthersFemaleVotersCountRampur + OthersFemaleVotersCountMoradabad + OthersFemaleVotersCountBijnor + OthersFemaleVotersCountBangalore + OthersFemaleVotersCountHubli
		statistics.Scope.OthersFemale = OthersFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimP = muslimVotersP

		dalitVotersP = (float64(dalitVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.DalitP = dalitVotersP

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

		dalitMaleVotersP = (float64(dalitMaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.DalitMaleP = dalitMaleVotersP

		dalitFemaleVotersP = (float64(dalitFemaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.DalitFemaleP = dalitFemaleVotersP

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
	qsScopeBijnor = qsScopeBijnor.SetCond(condScope)

	//Reset all counters
	votersCountRampur = 0
	votersCountMoradabad = 0
	muslimVotersCountRampur = 0
	muslimVotersCountMoradabad = 0
	dalitVotersCountRampur = 0
	dalitVotersCountMoradabad = 0
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
	dalitMaleVotersCountRampur = 0
	dalitMaleVotersCountMoradabad = 0
	dalitFemaleVotersCountRampur = 0
	dalitFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	dalitVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	dalitMaleVotersCountBijnor = 0
	dalitFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	dalitVotersCountBangalore = 0
	dalitVotersCountHubli = 0
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
	dalitMaleVotersCountBangalore = 0
	dalitMaleVotersCountHubli = 0
	dalitFemaleVotersCountBangalore = 0
	dalitFemaleVotersCountHubli = 0
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
			votersCountBijnor, err = qsBijnor.Count()
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
	dalitVotersCountRampur = 0
	dalitVotersCountMoradabad = 0
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
	dalitMaleVotersCountRampur = 0
	dalitMaleVotersCountMoradabad = 0
	dalitFemaleVotersCountRampur = 0
	dalitFemaleVotersCountMoradabad = 0
	OthersMaleVotersCountRampur = 0
	OthersMaleVotersCountMoradabad = 0
	OthersFemaleVotersCountRampur = 0
	OthersFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	dalitVotersCountBijnor = 0
	OthersVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	dalitMaleVotersCountBijnor = 0
	dalitFemaleVotersCountBijnor = 0
	OthersMaleVotersCountBijnor = 0
	OthersFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	dalitVotersCountBangalore = 0
	dalitVotersCountHubli = 0
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
	dalitMaleVotersCountBangalore = 0
	dalitMaleVotersCountHubli = 0
	dalitFemaleVotersCountBangalore = 0
	dalitFemaleVotersCountHubli = 0
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
			votersCountBijnor, err = qsScopeBijnor.Count()
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
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Statistics API: "+err.Error())
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
		recipient []*modelAccounts.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Request: POST /api/otp, Json: %s", e.Ctx.Input.RequestBody))

	// Create query string for account table
	qsAccount := o.QueryTable("account")
	qsRecipient := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(modelAccounts.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get OTP API: "+err.Error())
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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get OTP API: "+err.Error())
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Get OTP API: "+err.Error())
	}

	num, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Update(orm.Params{
		"OTP": otp,
	})

	if err != nil || num != 1 {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Error occur while updating the otp. Please contact electionubda.com team for assistance.")
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get OTP API: "+err.Error())
			responseStatus.Error = err.Error()
		}
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsRecipient.Filter("Mobile_no__exact", account.Mobile_no).All(&recipient)

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get OTP API: "+err.Error())
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
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get OTP API: "+err.Error())
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

func sendOTP(otp int, toEmail string, displayName string, mobileNumber int64) error {
	// compose the message
	title := strconv.Itoa(otp) + " is your One Time Password - " + strconv.FormatInt(mobileNumber, 10)
	body := "Hi " + displayName + "!\n\nWelcome to Election UBDA.\n\nThanks & Regards,\nElectionUBDA Team"

	m := email.NewMessage(title, body)
	m.From = mail.Address{Name: "ElectionUBDA Team", Address: "electionubda@gmail.com"}
	toCC1 := "baligcoup8@gmail.com"
	toCC2 := "iiiftekhar@gmail.com"
	m.To = []string{toEmail, toCC1, toCC2}

	// send it
	auth := smtp.PlainAuth("", "electionubda@gmail.com", "hu123*ElectionUBDA", "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
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
		users []*modelAccounts.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Request: POST /api/register, Json: %s", e.Ctx.Input.RequestBody))

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(modelAccounts.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Register API: "+err.Error())
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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get Register API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Get Register API: "+err.Error())
				responseStatus.Error = err.Error()
			}
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		user := new(modelAccounts.Account)
		users[0].Token = token
		user = users[0]
		e.Data["json"] = &user
		e.ServeJSON()
	} else {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't generate the token. Please contact electionubda.com team for assistance.")
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Register API: "+err.Error())
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
}

func (e *ElectionController) GetList() {
	var (
		num  int64
		user []*modelAccounts.Account
		err  error
	)

	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/list, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get List API: "+err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Get Districts or Acs API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	list := new(modelVoters.List)

	err = json.Unmarshal(inputJson, &list)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get List API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(list.Acs) > 0 {
		sections := getApprovedSections(user[0].Approved_acs, list.Acs, user[0].Approved_sections)
		//sort.Strings(sections)
		e.Data["json"] = &sections
		e.ServeJSON()
	}

	if len(list.Districts) > 0 {
		acs := getApprovedAcs(user[0].Approved_districts, list.Districts, user[0].Approved_acs)
		sort.Strings(acs)
		e.Data["json"] = &acs
		e.ServeJSON()
	}

	approvedDistricts := strings.Split(strings.TrimSpace(user[0].Approved_districts), ",")
	sort.Strings(approvedDistricts)
	e.Data["json"] = &approvedDistricts
	e.ServeJSON()
}

func (e *ElectionController) ReadJson() {
	inputJson := e.Ctx.Input.RequestBody
	readJsons := new(modelVoters.ReadJsons)

	acNameSectionName := make(map[string]string)

	err := json.Unmarshal(inputJson, &readJsons)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Get List API: "+err.Error())
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

func getApprovedSections(csvAcs string, acs []string, csvApprovedSections string) []string {
	approvedAcs := strings.Split(strings.TrimSpace(csvAcs), ",")
	approvedSections := getSections(approvedAcs)
	sections := getSections(acs)

	return getCommonItems(sections, approvedSections, strings.Split(strings.TrimSpace(csvApprovedSections), ","))
}

func getApprovedAcs(csvDistricts string, districts []string, csvApprovedAcs string) []string {
	approvedDistricts := strings.Split(strings.TrimSpace(csvDistricts), ",")
	approvedAcs := getAcs(approvedDistricts)
	acs := getAcs(districts)

	return getCommonItems(acs, approvedAcs, strings.Split(strings.TrimSpace(csvApprovedAcs), ","))
}

func getSections(acs []string) []string {
	var (
		sections   []string
		dbSections []*modelVoters.Section
		err        error
		tableName  string
	)

	o := orm.NewOrm()
	o.Using("default")

	for _, ac := range acs {
		switch ac {
		case "Najibabad", "Nagina", "Barhapur", "Dhampur", "Nehtaur", "Bijnor", "Chandpur", "Noorpur":
			tableName = modelVoters.GetTableName("Bijnor")
		case "Kanth", "Thakurdwara", "Moradabad Rural", "Moradabad Nagar", "Kundarki", "Bilari":
			tableName = modelVoters.GetTableName("Moradabad")
		case "Suar", "Chamraua", "Bilaspur", "Rampur", "Milak":
			tableName = modelVoters.GetTableName("Rampur")
		}

		_, err = o.Raw("SELECT section_name_english AS section, religion_english AS religion, COUNT(*) AS count FROM "+tableName+" WHERE ac_name_english = ? GROUP BY section_name_english, religion_english ORDER BY religion_english, COUNT(*) DESC", ac).QueryRows(&dbSections)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Sections API: "+err.Error())
		}

		sectionName := make(map[string]int)
		for _, section := range dbSections {
			if _, ok := sectionName[section.Section]; ok {
				sectionName[section.Section] = sectionName[section.Section] + section.Count
			} else {
				sectionName[section.Section] = section.Count
			}
		}

		for _, section := range dbSections {
			if section.Religion == "Muslim" {
				total := sectionName[section.Section]
				percentage := (section.Count / total) * 100
				section.Section = section.Section + " -> " + strconv.Itoa(section.Count) + "|" + strconv.Itoa(percentage) + "%"
				sections = append(sections, section.Section)
			}
		}
	}

	return sections
}

func getAcs(districts []string) []string {
	var (
		acs   []string
		dbAcs []string
		err   error
	)

	o := orm.NewOrm()
	o.Using("default")

	for _, district := range districts {
		tableName := modelVoters.GetTableName(district)
		_, err = o.Raw("SELECT DISTINCT ac_name_english FROM " + tableName).QueryRows(&dbAcs)
		if err != nil {
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Get Acs API: "+err.Error())
		}
		for _, dbAc := range dbAcs {
			acs = append(acs, dbAc)
		}
	}
	return acs
}

func getCommonItems(list1 []string, list2 []string, list3 []string) []string {
	var items []string
	var itemsToReturn []string

	// Get the items common in first 2 lists
	for _, item := range list1 {
		if contains(list2, item) {
			items = append(items, item)
		}
	}

	if len(list3) == 1 && list3[0] == "" {
		return items
	}

	if len(list3) > 0 {
		// Get the items common from first 2 lists and 3rd list
		for _, item := range items {
			if contains(list3, strings.TrimSpace(strings.Split(item, "->")[0])) {
				itemsToReturn = append(itemsToReturn, item)
			}
		}
	} else {
		return items
	}

	return itemsToReturn
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
		user []*modelAccounts.Account
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

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/voter, Json: %s", mobileNo, e.Ctx.Input.RequestBody))

	if mobileNo == 0 || token == "" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

	_, err = qsAccount.Filter("Mobile_no__exact", mobileNo).Update(orm.Params{
		"Last_login": formattime.CurrentTime(),
	})
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Update Voter API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	voter := new(modelVoters.UpdateVoter)

	err = json.Unmarshal(inputJson, &voter)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Apply filters for each query string
	// Voter Id
	for _, voterId := range voter.VoterID {
		if voterId > 0 {
			condVoterId = condVoterId.Or("Voter_id__exact", voterId)
		}
	}

	if voter.District == "Moradabad" {
		qsMoradabad = qsMoradabad.SetCond(condVoterId)
		if voter.Vote == 1 || voter.Vote == 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Vote":       voter.Vote,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Email)) > 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Email":      voter.Email,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if voter.MobileNo > 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Mobile_no":  voter.MobileNo,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Image)) > 0 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Image":      voter.Image,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if voter.Age >= 18 {
			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Age":        voter.Age,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(voter.Religion)) > 0 {
			relHindi := ""
			if strings.TrimSpace(voter.Religion) == "Muslim" {
				relHindi = "मुसलमान"
			} else if strings.TrimSpace(voter.Religion) == "Others" {
				relHindi = "अन्य"
			} else if strings.TrimSpace(voter.Religion) == "Dalit" {
				relHindi = "दलित"
			}

			updatedRows, err := qsMoradabad.Update(orm.Params{
				"Religion_english": voter.Religion,
				"Religion_hindi":   relHindi,
				"Updated_on":       formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = fmt.Sprintf("The voter data has been successfully updated.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if voter.District == "Rampur" {
		qsRampur = qsRampur.SetCond(condVoterId)
		if voter.Vote == 1 || voter.Vote == 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
				"Vote":       voter.Vote,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Email)) > 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
				"Email":      voter.Email,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if voter.MobileNo > 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
				"Mobile_no":  voter.MobileNo,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Image)) > 0 {
			updatedRows, err := qsRampur.Update(orm.Params{
				"Image":      voter.Image,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if voter.Age >= 18 {
			updatedRows, err := qsRampur.Update(orm.Params{
				"Age":        voter.Age,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(voter.Religion)) > 0 {
			relHindi := ""
			if strings.TrimSpace(voter.Religion) == "Muslim" {
				relHindi = "मुसलमान"
			} else if strings.TrimSpace(voter.Religion) == "Others" {
				relHindi = "अन्य"
			} else if strings.TrimSpace(voter.Religion) == "Dalit" {
				relHindi = "दलित"
			}

			updatedRows, err := qsRampur.Update(orm.Params{
				"Religion_english": voter.Religion,
				"Religion_hindi":   relHindi,
				"Updated_on":       formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = fmt.Sprintf("The voter data has been successfully updated.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if voter.District == "Bijnor" {
		qsBijnor = qsBijnor.SetCond(condVoterId)
		if voter.Vote == 1 || voter.Vote == 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
				"Vote":       voter.Vote,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Email)) > 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
				"Email":      voter.Email,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if voter.MobileNo > 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
				"Mobile_no":  voter.MobileNo,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
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

		if len(strings.TrimSpace(voter.Image)) > 0 {
			updatedRows, err := qsBijnor.Update(orm.Params{
				"Image":      voter.Image,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the image.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if voter.Age >= 18 {
			updatedRows, err := qsBijnor.Update(orm.Params{
				"Age":        voter.Age,
				"Updated_on": formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the age.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}

		if len(strings.TrimSpace(voter.Religion)) > 0 {
			relHindi := ""
			if strings.TrimSpace(voter.Religion) == "Muslim" {
				relHindi = "मुसलमान"
			} else if strings.TrimSpace(voter.Religion) == "Others" {
				relHindi = "अन्य"
			} else if strings.TrimSpace(voter.Religion) == "Dalit" {
				relHindi = "दलित"
			}

			updatedRows, err := qsBijnor.Update(orm.Params{
				"Religion_english": voter.Religion,
				"Religion_hindi":   relHindi,
				"Updated_on":       formattime.CurrentTime(),
			})
			if err != nil {
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Update Voter API: "+err.Error())
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = err.Error()
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
			if updatedRows < 1 {
				responseStatus := modelVoters.NewResponseStatus()
				responseStatus.Response = "error"
				responseStatus.Message = fmt.Sprintf("Unable to update the religion.")
				responseStatus.Error = "The voter id(s) provided is/are not valid."
				e.Data["json"] = &responseStatus
				e.ServeJSON()
			}
		}
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = fmt.Sprintf("The voter data has been successfully updated.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}
	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "error"
	responseStatus.Message = fmt.Sprintf("The voter data didn't get updated. Please check your JSON sent as: %s", inputJson)
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}
