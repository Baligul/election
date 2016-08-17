/*
   GET VOTERS
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/voters

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/voters

   GET STATISTIC
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[2882],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":["Rampur"],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["रामपुर"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":["मुरादाबाद"],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}' http://localhost:8080/api/statistic

   GET STATISTICS
   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[] },"scope":{"state_number":[],"district_number":[19,20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}}' http://localhost:8080/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[] },"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}}' http://localhost:8080/api/statistics

   curl -X POST -H "Content-Type: application/json" -d '{"query":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":["Muslim"],"religion_hindi":[],"age":[] },"scope":{"state_number":[],"district_number":[20],"voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":["M"],"id_card_number":[],"district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":[],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[]}}' http://localhost:8080/api/statistics

   GET OTP
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9564783954}' http://localhost:8080/api/otp

   Register
   curl -X POST -H "Content-Type: application/json" -d '{"mobile_no": 9343352734, "otp":23435}' http://localhost:8080/api/register

   Get List
   curl -X POST -H "Content-Type: application/json" -d '{"districts": [19,20], "acs":[34, 43]}' http://localhost:8080/api/list

   Read Json
   curl -X POST -H "Content-Type: application/json" -d @json/data.json http://localhost:8080/api/read/json
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

	models "github.com/Baligul/election/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/craigmj/gototp"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://ggxssikrsehequ:sQElIpN-CHqcFFNAx7mJO31Y3v@ec2-54-225-93-34.compute-1.amazonaws.com:5432/da6obv8tnlvcev")
	//orm.RegisterDataBase("default", "postgres", "user=member dbname=election sslmode=disable")
	orm.RegisterModel(new(models.Account))
	orm.RegisterModel(new(models.Voter_21))
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
		voters               []*models.Voter_21
		votersRampur         []*models.Voter_21
		votersMoradabad      []*models.Voter_21
		votersBijnor         []*models.Voter_21
		votersBangalore      []*models.Voter_21
		votersHubli          []*models.Voter_21
		num                  int64
		user                 []*models.Account
		err                  error
	)

	votersWithTotal := new(models.Voters)

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
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if limit == 0 {
		limit = -1
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(models.Query)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	qsRampur := o.QueryTable(models.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(models.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(models.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(models.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(models.GetTableName("Hubli"))

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
			condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__icontains", acNameEnglish)
		}
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__icontains", acNameHindi)
		}
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__icontains", sectionNameEnglish)
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

	qsRampur = qsRampur.SetCond(cond)
	qsMoradabad = qsMoradabad.SetCond(cond)
	qsBijnor = qsBijnor.SetCond(cond)

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
			if state == 27 {
				votersCountRampur, _ = qsRampur.Count()
				_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
				votersCountMoradabad, _ = qsMoradabad.Count()
				_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
				votersCountBijnor, _ = qsBijnor.Count()
				_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
			}

			if state == 12 {
				votersCountBangalore, _ = qsBangalore.Count()
				_, err = qsBangalore.Limit(limit, offset).All(&votersBangalore)
				votersCountHubli, _ = qsHubli.Count()
				_, err = qsHubli.Limit(limit, offset).All(&votersHubli)
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, _ = qsMoradabad.Count()
			_, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}

		if district == 20 {
			votersCountRampur, _ = qsRampur.Count()
			_, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}

		if district == 21 {
			votersCountBijnor, _ = qsBijnor.Count()
			_, err = qsBijnor.Limit(limit, offset).All(&votersBijnor)
		}
	}

	voters = append(votersRampur, votersMoradabad...)
	voters = append(votersRampur, votersBijnor...)
	voters = append(voters, votersBangalore...)
	voters = append(voters, votersHubli...)

	totalVotersCount := votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli + votersCountBijnor

	if votersCountRampur > 0 || votersCountMoradabad > 0 || votersCountBijnor > 0 || votersCountBangalore > 0 || votersCountHubli > 0 {
		votersWithTotal.Total = totalVotersCount
		votersWithTotal.Voters = voters
		e.Data["json"] = votersWithTotal
	} else {
		responseStatus := models.NewResponseStatus()
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
		otherVotersCount                 int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		otherMaleVotersCount             int64
		otherFemaleVotersCount           int64
		muslimVotersP                    float64
		otherVotersP                     float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		otherMaleVotersP                 float64
		otherFemaleVotersP               float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		otherVotersCountRampur           int64
		otherVotersCountMoradabad        int64
		otherVotersCountBijnor           int64
		otherVotersCountBangalore        int64
		otherVotersCountHubli            int64
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
		otherMaleVotersCountRampur       int64
		otherMaleVotersCountMoradabad    int64
		otherMaleVotersCountBijnor       int64
		otherMaleVotersCountBangalore    int64
		otherMaleVotersCountHubli        int64
		otherFemaleVotersCountRampur     int64
		otherFemaleVotersCountMoradabad  int64
		otherFemaleVotersCountBijnor     int64
		otherFemaleVotersCountBangalore  int64
		otherFemaleVotersCountHubli      int64
		num                              int64
		user                             []*models.Account
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
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(models.Query)
	statistic := new(models.Statistic)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = "Invalid Json. Unable to parse."
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	qsRampur := o.QueryTable(models.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(models.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(models.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(models.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(models.GetTableName("Hubli"))

	cond := orm.NewCondition()
	condAcNumber := orm.NewCondition()
	condPartNumber := orm.NewCondition()
	condSectionNumber := orm.NewCondition()
	condAcNameEnglish := orm.NewCondition()
	condAcNameHindi := orm.NewCondition()
	condSectionNameEnglish := orm.NewCondition()
	condSectionNameHindi := orm.NewCondition()
	condAge := orm.NewCondition()

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
			condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__icontains", acNameEnglish)
		}
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__icontains", acNameHindi)
		}
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__icontains", sectionNameEnglish)
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
				otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCount > 0 {
		statistic.Total = votersCount

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistic.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBijnor + otherVotersCountBangalore + otherVotersCountHubli
		statistic.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistic.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistic.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistic.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistic.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBijnor + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistic.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBijnor + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
		statistic.OtherFemale = otherFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCount)) * 100
		statistic.MuslimP = muslimVotersP

		otherVotersP = (float64(otherVotersCount) / float64(votersCount)) * 100
		statistic.OtherP = otherVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCount)) * 100
		statistic.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCount)) * 100
		statistic.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCount)) * 100
		statistic.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCount)) * 100
		statistic.MuslimFemaleP = muslimFemaleVotersP

		otherMaleVotersP = (float64(otherMaleVotersCount) / float64(votersCount)) * 100
		statistic.OtherMaleP = otherMaleVotersP

		otherFemaleVotersP = (float64(otherFemaleVotersCount) / float64(votersCount)) * 100
		statistic.OtherFemaleP = otherFemaleVotersP

		e.Data["json"] = statistic
	} else {
		responseStatus := models.NewResponseStatus()
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
		otherVotersCount                 int64
		maleVotersCount                  int64
		femaleVotersCount                int64
		muslimMaleVotersCount            int64
		muslimFemaleVotersCount          int64
		otherMaleVotersCount             int64
		otherFemaleVotersCount           int64
		muslimVotersP                    float64
		otherVotersP                     float64
		maleVotersP                      float64
		femaleVotersP                    float64
		muslimMaleVotersP                float64
		muslimFemaleVotersP              float64
		otherMaleVotersP                 float64
		otherFemaleVotersP               float64
		muslimVotersCountRampur          int64
		muslimVotersCountMoradabad       int64
		muslimVotersCountBijnor          int64
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		otherVotersCountRampur           int64
		otherVotersCountMoradabad        int64
		otherVotersCountBijnor           int64
		otherVotersCountBangalore        int64
		otherVotersCountHubli            int64
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
		otherMaleVotersCountRampur       int64
		otherMaleVotersCountMoradabad    int64
		otherMaleVotersCountBijnor       int64
		otherMaleVotersCountBangalore    int64
		otherMaleVotersCountHubli        int64
		otherFemaleVotersCountRampur     int64
		otherFemaleVotersCountMoradabad  int64
		otherFemaleVotersCountBijnor     int64
		otherFemaleVotersCountBangalore  int64
		otherFemaleVotersCountHubli      int64
		num                              int64
		user                             []*models.Account
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
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	queries := new(models.Queries)
	statistics := new(models.Statistics)

	err = json.Unmarshal(inputJson, &queries)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = "Invalid Json. Unable to parse."
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// Create query string for each and every district
	// Query
	qsRampur := o.QueryTable(models.GetTableName("Rampur"))
	qsMoradabad := o.QueryTable(models.GetTableName("Moradabad"))
	qsBijnor := o.QueryTable(models.GetTableName("Bijnor"))
	qsBangalore := o.QueryTable(models.GetTableName("Bangalore"))
	qsHubli := o.QueryTable(models.GetTableName("Hubli"))

	// Scope
	qsScopeRampur := o.QueryTable(models.GetTableName("Rampur"))
	qsScopeMoradabad := o.QueryTable(models.GetTableName("Moradabad"))
	qsScopeBijnor := o.QueryTable(models.GetTableName("Bijnor"))
	qsScopeBangalore := o.QueryTable(models.GetTableName("Bangalore"))
	qsScopeHubli := o.QueryTable(models.GetTableName("Hubli"))

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
			condAcNameEnglishQuery = condAcNameEnglishQuery.Or("Ac_name_english__icontains", acNameEnglish)
		}
	}

	// Ac Name English Scope
	for _, acNameEnglish := range queries.Scope.AcNameEnglish {
		if len(strings.TrimSpace(acNameEnglish)) > 0 {
			condAcNameEnglishScope = condAcNameEnglishScope.Or("Ac_name_english__icontains", acNameEnglish)
		}
	}

	// Ac Name Hindi Query
	for _, acNameHindi := range queries.Query.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindiQuery = condAcNameHindiQuery.Or("Ac_name_hindi__icontains", acNameHindi)
		}
	}

	// Ac Name Hindi Scope
	for _, acNameHindi := range queries.Scope.AcNameHindi {
		if len(strings.TrimSpace(acNameHindi)) > 0 {
			condAcNameHindiScope = condAcNameHindiScope.Or("Ac_name_hindi__icontains", acNameHindi)
		}
	}

	// Section Name English Query
	for _, sectionNameEnglish := range queries.Query.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglishQuery = condSectionNameEnglishQuery.Or("Section_name_english__icontains", sectionNameEnglish)
		}
	}

	// Section Name English Scope
	for _, sectionNameEnglish := range queries.Scope.SectionNameEnglish {
		if len(strings.TrimSpace(sectionNameEnglish)) > 0 {
			condSectionNameEnglishScope = condSectionNameEnglishScope.Or("Section_name_english__icontains", sectionNameEnglish)
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
				otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				votersCountBijnor, err = qsBijnor.Count()
				muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Count()
				votersCountHubli, err = qsHubli.Count()
				muslimVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBangalore, _ = qsBangalore.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountHubli, _ = qsHubli.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBangalore, _ = qsBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountHubli, _ = qsHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
			muslimVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsBijnor.Count()
			muslimVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	votersCountQuery = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountQuery > 0 {
		statistics.Query.Total = votersCountQuery

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Query.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBijnor + otherVotersCountBangalore + otherVotersCountHubli
		statistics.Query.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Query.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Query.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Query.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Query.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBijnor + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistics.Query.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBijnor + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
		statistics.Query.OtherFemale = otherFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimP = muslimVotersP

		otherVotersP = (float64(otherVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OtherP = otherVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.MuslimFemaleP = muslimFemaleVotersP

		otherMaleVotersP = (float64(otherMaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OtherMaleP = otherMaleVotersP

		otherFemaleVotersP = (float64(otherFemaleVotersCount) / float64(votersCountQuery)) * 100
		statistics.Query.OtherFemaleP = otherFemaleVotersP
	}

	//Reset all counters
	votersCountRampur = 0
	votersCountMoradabad = 0
	muslimVotersCountRampur = 0
	muslimVotersCountMoradabad = 0
	otherVotersCountRampur = 0
	otherVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	otherMaleVotersCountRampur = 0
	otherMaleVotersCountMoradabad = 0
	otherFemaleVotersCountRampur = 0
	otherFemaleVotersCountMoradabad = 0
	muslimVotersCountBijnor = 0
	otherVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	otherMaleVotersCountBijnor = 0
	otherFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	otherVotersCountBangalore = 0
	otherVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	otherMaleVotersCountBangalore = 0
	otherMaleVotersCountHubli = 0
	otherFemaleVotersCountBangalore = 0
	otherFemaleVotersCountHubli = 0

	// CALCULATIONS FOR SCOPE
	// Get voters for each state
	if len(queries.Scope.DistrictNameEnglish) == 0 && len(queries.Scope.DistrictNameHindi) == 0 && len(queries.Scope.DistrictNumber) == 0 {
		for _, state := range queries.Scope.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsScopeRampur.Count()
				votersCountMoradabad, err = qsScopeMoradabad.Count()
				muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
				maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
				femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
				femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				votersCountBijnor, err = qsScopeBijnor.Count()
				muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
				femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

			if state == 12 {
				votersCountBangalore, err = qsScopeBangalore.Count()
				votersCountHubli, err = qsScopeHubli.Count()
				muslimVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Muslim").Count()
				muslimVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Muslim").Count()
				otherVotersCountBangalore, _ = qsScopeBangalore.Filter("Religion_english__exact", "Other").Count()
				otherVotersCountHubli, _ = qsScopeHubli.Filter("Religion_english__exact", "Other").Count()
				maleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Count()
				maleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Count()
				femaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Count()
				femaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Count()
				muslimMaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimMaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				muslimFemaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
				otherMaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherMaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountBangalore, _ = qsScopeBangalore.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
				otherFemaleVotersCountHubli, _ = qsScopeHubli.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range queries.Scope.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameHindi == "बिजनौर" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// District Name English
	for _, districtNameEnglish := range queries.Scope.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if districtNameEnglish == "Bijnor" || districtNameEnglish == "bijnor" {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
		if district == 19 {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
			muslimVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Count()
			femaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountMoradabad, _ = qsScopeMoradabad.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 20 {
			votersCountRampur, err = qsScopeRampur.Count()
			muslimVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountRampur, _ = qsScopeRampur.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Count()
			femaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountRampur, _ = qsScopeRampur.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}

		if district == 21 {
			votersCountBijnor, err = qsScopeBijnor.Count()
			muslimVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Muslim").Count()
			otherVotersCountBijnor, _ = qsScopeBijnor.Filter("Religion_english__exact", "Other").Count()
			maleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Count()
			femaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Count()
			muslimMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Muslim").Count()
			muslimFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Muslim").Count()
			otherMaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "M").Filter("Religion_english__exact", "Other").Count()
			otherFemaleVotersCountBijnor, _ = qsScopeBijnor.Filter("Gender__exact", "F").Filter("Religion_english__exact", "Other").Count()
		}
	}

	votersCountScope = votersCountRampur + votersCountMoradabad + votersCountBijnor + votersCountBangalore + votersCountHubli

	if votersCountScope > 0 {
		statistics.Scope.Total = votersCountScope

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBijnor + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Scope.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBijnor + otherVotersCountBangalore + otherVotersCountHubli
		statistics.Scope.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBijnor + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Scope.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBijnor + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Scope.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBijnor + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Scope.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBijnor + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Scope.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBijnor + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistics.Scope.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBijnor + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
		statistics.Scope.OtherFemale = otherFemaleVotersCount

		muslimVotersP = (float64(muslimVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimP = muslimVotersP

		otherVotersP = (float64(otherVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OtherP = otherVotersP

		maleVotersP = (float64(maleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MaleP = maleVotersP

		femaleVotersP = (float64(femaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.FemaleP = femaleVotersP

		muslimMaleVotersP = (float64(muslimMaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimMaleP = muslimMaleVotersP

		muslimFemaleVotersP = (float64(muslimFemaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.MuslimFemaleP = muslimFemaleVotersP

		otherMaleVotersP = (float64(otherMaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OtherMaleP = otherMaleVotersP

		otherFemaleVotersP = (float64(otherFemaleVotersCount) / float64(votersCountScope)) * 100
		statistics.Scope.OtherFemaleP = otherFemaleVotersP
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
	otherVotersCountRampur = 0
	otherVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	otherMaleVotersCountRampur = 0
	otherMaleVotersCountMoradabad = 0
	otherFemaleVotersCountRampur = 0
	otherFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	otherVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	otherMaleVotersCountBijnor = 0
	otherFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	otherVotersCountBangalore = 0
	otherVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	otherMaleVotersCountBangalore = 0
	otherMaleVotersCountHubli = 0
	otherFemaleVotersCountBangalore = 0
	otherFemaleVotersCountHubli = 0

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
	otherVotersCountRampur = 0
	otherVotersCountMoradabad = 0
	maleVotersCountRampur = 0
	maleVotersCountMoradabad = 0
	femaleVotersCountRampur = 0
	femaleVotersCountMoradabad = 0
	muslimMaleVotersCountRampur = 0
	muslimMaleVotersCountMoradabad = 0
	muslimFemaleVotersCountRampur = 0
	muslimFemaleVotersCountMoradabad = 0
	otherMaleVotersCountRampur = 0
	otherMaleVotersCountMoradabad = 0
	otherFemaleVotersCountRampur = 0
	otherFemaleVotersCountMoradabad = 0
	votersCountBijnor = 0
	muslimVotersCountBijnor = 0
	otherVotersCountBijnor = 0
	maleVotersCountBijnor = 0
	femaleVotersCountBijnor = 0
	muslimMaleVotersCountBijnor = 0
	muslimFemaleVotersCountBijnor = 0
	otherMaleVotersCountBijnor = 0
	otherFemaleVotersCountBijnor = 0
	votersCountBangalore = 0
	votersCountHubli = 0
	muslimVotersCountBangalore = 0
	muslimVotersCountHubli = 0
	otherVotersCountBangalore = 0
	otherVotersCountHubli = 0
	maleVotersCountBangalore = 0
	maleVotersCountHubli = 0
	femaleVotersCountBangalore = 0
	femaleVotersCountHubli = 0
	muslimMaleVotersCountBangalore = 0
	muslimMaleVotersCountHubli = 0
	muslimFemaleVotersCountBangalore = 0
	muslimFemaleVotersCountHubli = 0
	otherMaleVotersCountBangalore = 0
	otherMaleVotersCountHubli = 0
	otherFemaleVotersCountBangalore = 0
	otherFemaleVotersCountHubli = 0

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
		responseStatus := models.NewResponseStatus()
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
	responseStatus := models.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = "This API is up and running!"
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}

func (e *ElectionController) OTP() {
	var (
		otp       int32
		num       int64
		recipient []*models.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")
	qsRecipient := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(models.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	exist := qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Exist()
	if !exist {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid mobile number or mobile number does not exists in our database. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	otp, err = generateOTP()
	if err != nil {
		responseStatus := models.NewResponseStatus()
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
		responseStatus := models.NewResponseStatus()
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
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't send the otp. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(recipient) > 0 {
		err = sendOTP(recipient[0].Otp, recipient[0].Email, recipient[0].Display_name, recipient[0].Mobile_no)

		if err != nil {
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Couldn't send the otp. Please contact electionubda.com team for assistance.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "ok"
		responseStatus.Message = fmt.Sprintf("One time password has been generated and sent to %s successfully.", recipient[0].Email)
		e.Data["json"] = &responseStatus
		e.ServeJSON()

	} else {
		responseStatus := models.NewResponseStatus()
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
		users []*models.Account
	)

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for account table
	qsAccount := o.QueryTable("account")

	inputJson := e.Ctx.Input.RequestBody
	account := new(models.Account)

	err := json.Unmarshal(inputJson, &account)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	exist := qsAccount.Filter("Mobile_no__exact", account.Mobile_no).Exist()
	if !exist {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid mobile number or mobile number does not exists in our database. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	_, err = qsAccount.Filter("Mobile_no__exact", account.Mobile_no).All(&users)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't generate the token. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if len(users) > 0 {
		if users[0].Otp != account.Otp || users[0].Otp == 0 {
			responseStatus := models.NewResponseStatus()
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
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Error occur while updating the token. Please contact electionubda.com team for assistance.")
			if err != nil {
				responseStatus.Error = err.Error()
			}
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
		user := new(models.Account)
		users[0].Token = token
		user = users[0]
		e.Data["json"] = &user
		e.ServeJSON()
	} else {
		responseStatus := models.NewResponseStatus()
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
		user []*models.Account
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
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	num, err = qsAccount.Filter("Mobile_no__exact", mobileNo).All(&user)

	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if num > 0 {
		if user[0].Token != token {
			responseStatus := models.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("You are not authorised for this request. Please contact electionubda.com team for assistance.")
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

	} else {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Couldn't serve your request at this time. Please contact electionubda.com team for assistance.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	inputJson := e.Ctx.Input.RequestBody
	list := new(models.List)

	err = json.Unmarshal(inputJson, &list)
	if err != nil {
		responseStatus := models.NewResponseStatus()
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
	readJsons := new(models.ReadJsons)

	acNameSectionName := make(map[string]string)

	err := json.Unmarshal(inputJson, &readJsons)
	if err != nil {
		responseStatus := models.NewResponseStatus()
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
			sections = append(sections, "AHMADPUR MAJEED", "REHMAN NAGAR", "KALYANPUR", "RAMNAGAR", "GAJIPUR AN0", "HAKUMATPUR KESHO", "MANOHARWALA", "PUNDARI KHURD", "KAMRAJPUR AN0", "SAMIPUR AN0", "CHANDGOYALA", "FAZALPUR HAIBAT", "MEERAMPUR BEGA", "MOTA DHAK", "SAHANPUR NANU", "DHANSINI AN0", "CHOWK BAZAR", "DAUDPUR HAJI", "KALHERI", "RAMPUR CHATHA", "AKBARPUR ANWLA", "KACHAHARI SARAI AN0", "SHEKHPURA ALAM", "NARAYANPUR RATAN AN0", "SHRAWANPUR", "SIKKAMPUR", "GOSPUR", "KHAIRULLAPUR AN0", "PARWATPUR MAKHDUMPUR AN0", "SADULLA NAGAR", "SAIDPUR MIRA", "SAIDPURI", "SIKANDARPUR BASI", "JEETPUR KHAS", "SADAT NAGAR B.A.", "KUTUBPUR NANGLI", "MANDAWALI AN0", "GURHA", "FAZALPUR HABIB AN0", "HIMPUR PACHATPURA", "MUBARAKPUR RATHE", "MALIPUR", "MEMRAN", "GULALWALI", "MIRAMPUR DURG", "HAKIMPUR KAJI", "HAKIMPUR DURG", "AGUPURA PYARA", "DARODGRAN", "HARSHWARA AN0", "QURESHIYAN AN0", "DHARAMPUR BHOJA URF PUNDRI KALA", "SUNGERPUR", "RAM SAHAYWALI", "PASBAN", "AURANGPUR FATTE KHAN", "MALUKWALI", "HASAN ALIPUR", "MATHURAPUR MOR AN0", "NAWABGANJ", "KATRA CHETRAM", "THAPAL", "MOCHIPURA", "MEHNDI BAGH", "LALPUR SHOJIMAL A.", "SHEKH SARAI", "PURANPUR GARHI A", "JAHANABAD", "KAROLI", "BISATIYAN", "MOHD. TAHARPUR", "REHMAPUR", "KHERPUR JALIKA", "CHAR BAGH", "CHANDOK AN0", "NAYAMATPUR", "SADULLAPUR", "NOORPUR DEHRA", "RAJARAMPUR FAZIL", "SHAHALIPUR NAKI", "NAWABGANJ", "MOHD PUR RAVA", "MUKTESHWAR MAHADEV", "SALEMPUR", "THAPAL", "MUBARAKPUR AN0", "VISHANPUR", "KACHHERI SARAI AN0", "VANSH GOPALPUR", "GADARIYAN", "KAMALPUR", "SHAHJAHANPUR", "MAHAL SARAI", "VEERPUR", "VIRUWALA", "TAJPUR", "DAUDPUR NANHERA", "RAGHUNATHPUR", "BASERI", "BALAKRAM", "ASADULLAPUR AN0", "JAGDEESHPUR", "SANWALDAS", "KAMGARPUR AN0", "KISHANPUR", "SHEKHPUR GARHU", "RAIPUR SUMALKERI", "FATEHULLAPUR DURG", "PREAMPUR", "BHAWAN", "FAZALPUR KHAS", "SABALGARH", "RAFIPUR MAJARA", "MUNEER GANJ AN0", "KANWALNAINPUR", "HALDUKHATA", "SHRAWANPUR", "JATPURA KHAS", "MALIYAN", "SAIDPUR MIRA", "AJAMPUR MOHD.ALI URF KHANPUR", "RAMDASWALI", "SHERPUR HARSARAN", "MAQBARA AN0", "PRATAPPUR", "PURUSHAOTTAMPUR", "DIWAN PARMANAND", "RAHUKHEDI KAURA", "PRATHVIPUR", "MUNEERGANJ", "MUSTAFAPUR ANSU", "SADATNAGAER A0", "SHYAMLI", "HIMAUPUR RAI", "KISHORPUR", "SANIYAN AN0", "RAFIPUR MOHAN", "SHAHZAHAPUR JASRATH", "SAJAWALPUR", "NOORPUR DEHRA", "MOHD PUR ATA", "MAYAPURI", "MUGLUSHAH AN0", "CHATRUWALA", "SHAHALIPUR NAKI", "RAILWAY COLONY AN0", "SAHANPUR SANTOKI", "THAPAL", "KATRA CHETRAM", "KHANPUR", "GARHMALPUR", "ALAWALPUR NAINU", "ASLAMPUR JHOJA", "SHERAWALA", "LALPUR RAI", "MIRZAPUR RANGILA", "MOHD. ALIPUR PARMA", "GAJRULA", "MOHD.ALIPUR ALIMUDDIN", "MAHAWATPUR DALPAT", "ARAFPUR KHAJURI", "BHOGPUR", "GULALWALI", "RANIKOTA", "MAHAL SARAI", "OOTWAN", "JASWANTPUR AN0", "ALIPURA AN0", "GULZAR NAGAR", "SANWALDAS", "SARAI SAID GHORA", "AGUPURA PYARA", "PADHAN", "RAWAPURI AN0", "IBRAHIMPUR RAJU", "SHRAWANPUR", "KUSHHALPUR MARKA", "JATPURA KHAS", "SAIDPUR MIRA", "NAGAL AN0", "AJAMPUR GAJI", "SHAHZADPUR AN0", "JATAN", "SHAHARYARPUR BANGER", "DHANORA", "GOYALA B.A.", "KUTUBPUR NANGLI", "KHALILPUR", "HAKIMAN", "SANTOMALAN AN0", "SHAREEFPUR BANGAR B.A.", "MOHD. ALIPUR VEERBHAN A", "JADOPUR", "RAIPUR KHAS AN0", "AURANGPUR BASANTA", "PURANPUR NAROTTAM", "AMRITSARI", "RAJARAMPUR TULSI", "CHOWK BAZAR", "SAHANPUR NANU", "KUTUBPUR", "MAJIDGANJ AN0", "TELINAGALI", "FAZALPUR MAN", "SHYAMLI", "TAYYABPUR GORVA", "BHADOLA", "RAMPUR BANWARI", "HUSAINPUR MAKHANA", "AMANULLAPUR", "MOTHALA", "TODA", "SEWARAM AN0", "DINORA", "NADE HIND", "NOORAMPUR", "IBRAHIMPUR JAMALUDDIN", "TAKSAL", "RASOOLPUR SAID AN0", "SHAHPUR MIRA", "PATHANPURA AN0", "BAIBALWALA", "MAHAL SARAI AN0", "HARIJAN BASTI", "KOTAWALI", "JAFARABAD", "BHAGUWALA", "CHOWK BAZAR", "HARCHANDPUR AN0", "DARIYAPUR", "LAHAK KALA", "IBRAHIMPUR BAHAUDDIN", "GOYALA", "JATPURA KHAS", "MOHD. ALIPUR VEERBHAN", "WAHID NAGAR AN0", "FIELD", "AJAMPUR MOHD.ALI URF KHANPUR", "BHAGWAN WALA", "KALLUGANJ", "JASPUR", "NADDAFAN", "KHERA", "TISOTRA AN0", "ZULFUKARPUR GARHI", "JAGDISHPUR", "MUSVI KHANPUR", "DHARMA GARHI", "PAIBAGH", "MAHAL SARAI", "DARBARASHAH", "ASADULLAPUR AN0", "SHERPUR ABHI", "QUILA", "SANWALDAS", "AGUPURA PYARA", "BHATOLI", "MUNGARPUR", "AURANGPUR BHIKKU", "JAFARPUR", "AURANGABAD AN0", "RAILWAY COLONY ANO", "NOORPUR DEHRA", "LAHAK KHURD", "CHAH SALOONO", "SHAHALIPUR NAKI", "NAWABGANJ", "SOFATPUR", "MURSHADPUR", "VIJAYPUR", "TAKSAL", "BASHIRPU", "KORIYA", "HAJIPUR", "BOREKI", "MAHARAIPUR SEKH URF DINORI", "RAMPUR BANWARI A.", "NARAYANPUR ICHCHA", "NANGLA UBBAN", "GIRDAWA SAHANPUR URF KHALAPAR", "MOHD. ALIPUR HIRDAY AN0", "AGUPURA PYARA", "RAILWAY COLONY RAHUKHERI KORA", "GULZAR NAGAR", "JASWANTPUR AN0", "OOTWAN", "GULALWALI", "RANIKOTA", "JHAKKAKI", "MOHD ALIPUR DWARKA", "MEDUWALA", "MOJJAMPUR SADAT", "PADHAN", "PURANPUR GARHI AN0", "MUNEERGANJ AN0", "KHANPUR", "KATRA CHETRAM", "MOHD. AMI KHANPUR", "BIJORI", "RANIPUR", "AURANGPUR FATTE KHAN", "RAILWAY COLONY AN0", "MAYAPURI", "MOHD PUR ATA", "AURANGPUR BASANTA AN0", "ZABTAGANJ AN0", "MAHAWATPUR DALPAT", "MOHD.ALIPUR ALIMUDDIN", "CHOWK BAZAR", "MIRZAPUR RANGILA", "LALPUR RAI", "SHERAWALA", "DUNGERPUR", "TAKSAL", "ASLAMPUR JHOJA", "GAJIPUR AN0", "SHAREEFPUR BANGAR A.", "BANJARAN", "CHAMROLLA", "AHMADPUR MAJEED", "MOHD. ALIPUR VEERBHAN A", "HAKIMAN", "AHMADPUR SADAT", "SANTOMALAN AN0", "IBRAHIMPUR BAWAN", "CHOWK BAZAR", "MAHAWATPUR BILLOCH", "KUTUBPUR", "SAHANPUR NANU", "MIRZAPUR SAID", "RAJARAMPUR TULSI", "AMRITSARI", "PREM NAGAR", "HASANPUR", "SAIWALA", "NAWABPURA", "NAGAL AN0", "MOHD. ALAMPUR", "SABNIGRAN", "AGUPURA PYARA", "ISSEPUR AN0", "IBRAHIMPUR RAJU", "MAHAWATPUR NATHA", "KUTUBPUR NANGLI", "MIRZAPUR", "JATAN", "HUMAYUPUR IDDU", "RAMPURA AN0", "MAHEMSAPUR", "FIELD", "MOHD. ALIPUR VEERBHAN", "JAMHIRI", "SANWALDAS", "JATPURA KHAS", "IBRAHIMPUR BAHAUDDIN", "KANWALNAINPUR", "LALPUR BIKKU", "MAHAL SARAI", "DARIYAPUR", "HARCHANDPUR AN0", "JWALI LALA AN0", "SARAI ALAM AN0", "GANGUWALA", "ZULFUKARPUR GARHI", "NANGLA PITHORA", "KHERA", "KALLUGANJ", "JASPUR", "SHERPUR HARSARAN", "BHAGWAN WALA", "RAMDASWALI AN0", "JABALPUR GUDAR AN0", "CHANDPURA", "DINORA", "NADE HIND", "TODA", "NANGAL AN0", "SEWARAM AN0", "MOTHALA", "NOORPUR DEHRA", "KAJIYAN", "RAMPUR BANWARI", "SHYAMLI", "FAZALPUR TABELA", "TELINAGALI", "SHAHALIPUR NAKI", "KALHERI A", "FAZALPUR MAN", "MAJIDGANJ AN0", "SHYAMIWALA AN0", "JAFARABAD", "KOTAWALI", "HARIJAN BASTI", "PATHANPURA AN0", "BAIBALWALA", "BHAGUWALA AN0", "IBRAHIMPUR JAMALUDDIN", "RASOOLPUR SAID AN0", "MURSHADPUR", "MOHD PUR RAVA", "MUKTESHWAR MAHADEV", "MAJIDGANJ", "SHAHALIPUR NAKI", "NAYAMATPUR", "CHAH SALOONO", "NOORPUR DEHRA", "KARMASKHERI", "HAREWALI", "NARAYANPUR ICHCHA", "HAJIPUR", "RAMPUR BANWARI A.", "BOREKI", "BASHIRPU", "TAKSAL", "DANIYALPUR", "SAIFABAD", "BARAMPUR AN0", "SANWALDAS", "QUILA", "NAJIMPUR", "JAGDEESHPUR", "DARBARASHAH", "RAHATPUR KHURD AN0", "SANIYAN AN-", "TAJPUR", "VEERPUR", "MAHAL SARAI", "SABALPUR AN0", "KESHOPUR", "MUSVI KHANPUR", "SHEKHPUR LALA", "DHARMA GARHI", "AURANGABAD AN0", "RAILWAY COLONY ANO", "SADATPUR", "MUSSEPUR", "BIJARKHATA", "AURANGPUR BHIKKU", "JAFARPUR", "MUNGARPUR", "RAJPUR NAVADA", "JWALACHANDI", "SAIDPUR MIRA", "SHAKARPUR BAHADAR", "SADULLA NAGAR", "GOSPUR", "KHAIRULLAPUR AN0", "NARAYANPUR RATAN AN0", "SHEKHPURA ALAM", "SHRAWANPUR", "KACHAHARI SARAI AN0", "TAHARPUR ISHAK AN0", "MALIPUR", "BULCHANDPUR", "HIMPUR PACHATPURA", "GURHA", "TATARPUR LALU AN0", "SADAT NAGAR B.A.", "SHERPUR HARSARAN", "SAIDPURI", "SIKANDARPUR BASI", "HAKUMATPUR KESHO", "GAJIPUR AN0", "RAMNAGAR", "KALYANPUR", "ANSARIYAN", "AHMADPUR MAJEED", "REHMAN NAGAR", "KANAKPUR AN0", "RAMPUR CHATHA", "KUTUBPUR", "DHANSINI AN0", "DAHIRPUR AN0", "NAGLA HARDAS", "FAZALPUR HAIBAT", "CHANDGOYALA", "SAMIPUR AN0", "PADARATHPUR", "MEHNDI BAGH", "RAMPURI", "THAPAL", "MATHURAPUR MOR AN0", "HASAN ALIPUR", "AKBARPUR CHOGANWA AN0", "AURANGPUR FATTE KHAN", "MOJJAMPUR TULSI", "NAYAMATPUR", "MOHD. TAHARPUR", "BISATIYAN", "SAIDPUR MIRA", "MIRAMPUR DURG", "HAKIMPUR DISONDI", "SHRAWANPUR", "MEMRAN", "SIKRODDA AN0", "HARNATHPUR", "SUNGERPUR", "MANUPURA", "KALALAN", "BANSFODAN AN0", "JALPUR AN0", "BEGAMPUR BHONAWALA", "GULISTAN", "DHARAMPUR BHOJA URF PUNDRI KALA", "KAMGARPUR AN0", "RAMPUR MANGAL", "SURSENA", "ASADULLAPUR AN0", "JAGDEESHPUR", "RAGHUNATHPUR", "BALAKRAM", "FAZALPUR FATEHULLA", "TAJPUR", "VIRUWALA", "VEERPUR", "RAFIPUR MAJARA", "FAZALPUR KHAS", "ABDULLAPUR", "SABALGARH", "JAFARPUR", "BARKHURDARPUR GOPAL", "MUNGARPUR", "RAIPUR SUMALKERI", "BAKARPUR", "MUKTESHWAR MAHADEV", "SALEMPUR", "MOHD PUR RAVA", "NAWABGANJ", "FARAZPUR", "NAYAMATPUR", "KAMALPUR", "TAKSAL", "VANSH GOPALPUR", "LALPUR SHOJIMAL", "VISHANPUR", "KISHORPUR", "DHARAMDAS", "NAWABGANJ", "HIMAUPUR RAI", "JARUFSAJAN", "SADATNAGAER A0", "JALABPUR GUDAR AN0", "AHMADPUR MAJEED", "MUNEERGANJ", "MUSTAFAPUR ANSU", "SANIYAN AN0", "MALIYAN", "SHAHABPURA UMRAV SINGH", "HALDUKHATA", "KANWALNAINPUR", "RASOOLPUR SAID", "MUNEER GANJ AN0", "MOHD. AMIPUR", "DIWAN PARMANAND", "RAHUKHEDI KAURA", "PRATHVIPUR", "BARKATPUR", "PURUSHAOTTAMPUR", "JATPURA BONDA AN0", "PRATAPPUR", "SHERPUR HARSARAN", "KASHIRAMPUR", "MAQBARA AN0", "MOJJAMPUR TULSI AN0", "LALPUR MAN", "AJAMPUR MOHD.ALI URF KHANPUR")
		case "Nagina":
			sections = append(sections, "SAIDKHERI", "RAMPURLASHKARI", "CHAMARAN", "MOHIUDDINPUR", "KHURRAM ALI SARAIH.NO 1-115TAK", "KAJISARAI  AN.H.NO 1-270", "UMRI AN.H.NO 158-END", "UGRSAINPUR", "SULEMASIKHOPUR", "PIPALSANA", "ASDULLAPUR PRITHVI", "SHADIPUR AN H.NO 56-END", "BUDGARA AN.", "HAIBATPUR PEERU", "PAKAHANPUR", "HARGAON CHANDAN AN.H.NO 131-END", "MAHESHWARI JAT AN0H.NO 131-END", "AHMAD KHEL AN.", "SAHUWAN", "MOHSANPUR", "PITHANHERI JHOJHA KALA", "SAIDPURIMAHICHAND", "LUHARI SARAI PURV AN.H.NO 1-283", "SARAI JEEVAN", "JALPURA", "KARONDA PACHDU AN.", "KARONDA PACHDU AN.H.NO 1-200", "SAHLIPUR HEERA", "NOORPUR HATTI", "RAMPUR ASHA", "KHAKROBAN ST.", "TAKIPUR TULSI NG.CH.", "BAKARPUR JAIRAM", "MOMINPUR LALU", "LAL SARAI AN.H.NO 1-150 TAK", "HASANLIPUR KHANDU", "CHODHRANA AN H.N0 1-124", "MOJAMPUR HARVANSH", "ALHADADPUR", "PADLA AN", "KASSABAN", "KALA NANGLI GAIR ABAD", "SUFIPUR ANGAD", "SAHLIPUR AB.SATTAR", "JAMALPUR KHOKO", "KHASOR", "CHODHRANA AN0H.NO 125-END", "SARAI IMMA", "MO ALIPUR SHEKH", "RAIPUR MOJAMPUR NARAYAN", "MAKSUDNAGAR DEVIDAS", "RASULPUR ALIMUDDIN", "BEGAMPUR SHADI AN0", "KOTLA", "KHUSHALNAGAR GARHI", "AURANPUR NANDLAL", "KOTWALI AN H.NO 1-412", "KHAKROBAN", "AJUPURACHOHAN", "MO.ALIPUR SUKHANAND", "BHARAIKI", "SABHACHANDPUR", "NOROJPUR GAIR ABAD", "HUSAINPUR MEERA", "NEJOWALI GAWDI", "SAMASPUR VEERBHAN", "KAJIPUR PACHDU", "HASALIPUR MUTHRA", "SHEKHPURA NOABAD", "MAHAPUR MOHD ALI", "KALAKHEDI", "BEGAMPUR CHAIMAL", "ISLAMPUR SADAT", "NANGLA ISLAM", "ANWARPUR CHATAR", "MUKIMPUR DUNIYA GAIR ABAD", "MUBARAKPUR MEERA AN", "KHAKROBAN UTTARI", "MOTHEPUR KIRAT", "MUKIMPUR PADARATH", "MANIHARI SARAI AN0 H.NO 1-146", "HAKIKATPUR VERCHAND", "ALAMPURASU G.ABD", "KASBA KOTRA AN.H.NO 101-212", "HARDASPUR MADE", "ALAMPUR", "GOPALPUR", "MATHERA CHOHAN", "LUHARI SARAI AN.H.NO 1-275", "CHOHANAN AN.", "PURSHOTAMPUR URF THEY PUR", "SAINAWALA", "JHALRI", "BEHADA CHOHAN", "ABDULAJIJPUR", "ISLAMPUR SAHLI", "BUDHPUR MUFTI", "DELPURA NAINU", "LALPUR", "CHATRIYANAGAR H.NO.325-877", "SHADIPUR IMMAR", "BUDHUPADA", "AMIRPUR MAIDAS GAIR ABAD", "HARDASPUR MADE", "KHIJARPURJAGGU", "ABULFAJALPUR KHAS", "NASIRPUR MITHARI", "KASBA", "CHAKNARANGI", "KAYASTHAN", "SHEKHPURATURK AN H.NO 1-67", "SARAI MEER AN H-NO 1-143", "AURANGPUR HIRDAY", "KAJISARAI AN.H.NO 74-END", "NAIMPUR GOVARDHAN", "NAKIPUR BAMNOLIGAIR ABAD", "HAKIKATPUR SEHSU", "KILA", "SAHJANPUR ROSHAN", "PADLI", "AFGANAN", "DELPURA KANNA", "CHAHRONAK", "HAIJARPUR VEERCHAND", "SAIDPURI", "MAMRAJPUR", "RASULPUR SAHBUDDIN", "KAMALPUR BHOGA", "CHATARBHOJPUR KUSHAL", "ISLAMPUR SAHU", "ROSHANPUR PRATAP", "JALALPUR SULTAN", "ASDULLAPUR KALYAN", "MUGLAN PURVI", "HAKIKATPURGANGWALI AN.", "KAMKARAN AN.", "AMBEDKAR NAGAR H.NO 233-324", "CHANDAK JATT", "KARONDA CHODHAR AN.H.NO 131-END", "HEEMPUR MANAK", "ANSARIYAN  AN.", "MILKIYAN", "PITTANHERI JIYA AN.H.NO 1-113", "MUBARAKPUR SAHARAN", "BIRDO NANGLI", "MUTHRAPUR URF KALUWALA", "LOHARI SARAI PU.AN.H.NO 11/2-180/4", "TUKHMAPURHARVANSH AN0", "JATI AN H.NO 1-45", "HUSAINPUR SULTAN", "JAFARPUR PARASRAM", "JAFRULLAPUR", "GOVINDPUR", "DHARAMPUR DHANSI GAIR ABAD", "NASRULLAPUR", "MOMINPUR LALU", "UMARPUR", "KAMALPURPURAINI", "HAKIKATPUR GOVIND", "KAJISARAI AN.H.NO 271-END", "ISLAMPUR MEERA", "ALIPUR GANGA", "MOJIPUR DHARMA", "CHITAWAR AN", "SAMASPUR KHORA URF KHODPURA", "MALAKPUR SEHSU", "MUGLAN PASCHIMI", "KHAIRULLAPUR", "SABHACHANDPUR KESHOAN", "KIRATPUR", "RAMPURCHAJMAL", "HAMIDPUR MAKHAN", "SHEKH SARAI", "RATANPUR", "FATEHPUR RAJARAM", "ABDIPUR KALYAN", "GHANGHERI", "AJAMPUR MOHAN", "MAHESHWARI JAT AN H.NO 1-130", "KAKADPUR", "BANWARIPUR", "TAKIPUR MANOHAR", "KALALAN AN.H.NO 1-171", "BILASPUR GAIR ABAD", "RAMPUR BISHNA", "MANAKCHAND", "VISHNOI SARAI AN.", "DOODHLI", "SALEMPUR GADHELI", "KAYRI FUL GAIR ABAD", "ASHRAFPUR", "AURANGJEBPURHARDAS", "GAJIPUR HIDAYAT URF GARHIBAN", "SHAKARPURI", "KHAKROBAN UTTARI", "MANPUR", "AKBARABAD  AN0 H.NO 549-END", "KALALAN AN.H.NO 1-350", "MUSLIM KATERA", "MEHMUDPUR", "SARKADA KHERI", "NIJAMPURDEVSI", "HARDASPUR MADE", "DHARAMPUR BHAGWAN GAIR ABAD", "KITHODA RAI", "NURALPIURBHAGWANT", "MEMAN SADAT AN.H.NO 1-263", "JAMALPUR BANGAR", "YARMOHDPUR PRITHVI", "MATHERA CHOHAN", "RAMPUR DAS AN.", "GOPALPUR", "NAINPURA AN.", "NAIMPUR GOVARDHAN", "BIHARIPUR", "KOTWALI AN", "MO ALIPUR SHEKH", "RAMKAYRI", "HAKIKATPUR PRAYAG", "GUJARPUR ASU GAIR ABAD", "MO.ALIPUR SUKHANAND", "AJUPURACHOHAN", "MAHAJANAN", "KHATAI", "SHEKHSARAI AN", "LUHARI SARAI AN.", "DHANORI KUNWAR", "KALAN AN. H.NO 351-END", "SIKRI", "MAHALKI", "NALBANDAN AN.H.NO 142-END", "KAJIPUR PACHDU", "HINDU KATERA", "HAKIKATPUR MUTHRA", "LAL SARAI AN.H.NO 1-102", "SARAI SHARIF NAGAR GAIR ABAD", "ALAMPUR GANGA", "KAYASTH SARAI AN 151-END", "UMRI AN.H.NO 1-157", "JALPURA", "LOHARI SARAI AN.H.NO 1-180", "AURANGJEBPUR FAJIL", "SAHZAHIR AN.H.NO 1-END", "RASULPUR DAUD", "BEGAMPUR HARREY", "MUKANDPUR RAMU", "TAKIPUR TULSI NG.CH.", "NURMOHDPUR", "MUSSEPUR", "MALKAN", "RAMPUR ASHA", "HARGAON CHANDAN AN.H.NO 1-130", "HASANLIPUR KHANDU", "MOMINPUR LALU", "ROSHANPUR CHIJROLI", "ARGUPURA GAIR ABAD", "ALHADADPUR", "NIJAMPURPADARATH", "AMANNAGAR", "BADSHAPUR MAIDAS", "SAHLIPUR AB.SATTAR", "LOHARAN", "CHITAWAR AN", "PITTHANHERI JIYA AN.H.NO 114-END", "PUNJABIYAN", "SARAI BAHAUDDIN", "CHAMARAN", "KADARPUR TAYYAB", "MOHIUDDINPUR", "HASANLIPUR BHOGAN", "PIPALSANA", "KAMALPUR  GAIR ABAD", "SINGHADIYAN", "AHMADKHEL AN.H.NO 138-END", "MASURI", "PAKAHANPUR", "TAKIPUR HARVANSH", "AURANGJEBPUR SAMBHAL", "AURANGJEBPUR CHANDA GAIR ABAD", "PALDI", "AZMABAD", "SULTANPUR SADAT", "DUDHLA", "HAMIDPUR MAKHAN", "KHURRAMPUR DALLU", "ISLAMPUR VISHNOI", "RATANPUR", "MUSTAFAPUR GAIR ABAD", "AFGANAN AN.", "AJAMPUR YAR MOHD", "MOMINPUR DASU", "MIRDGAN", "PIPALSANA", "HARBANSHPUR DHARAM", "SAHLIPUR ASHA", "AJAMPUR MOHAN", "MIRZALIPUR BHARA", "NEKPUR", "KALAPUR BUJURG", "BHANERA", "KOTWALI AN HNO -1-123", "RAFIKPUR", "JHANDA AN", "DHARAMSHANANGLI", "KAYASTH SARAI AN.H.NO 1-150", "SHISHGRAN AN. H.NO.1-121", "PADHAN", "HAMIDPUR GANGA GAIR ABAD", "MALAKPUR GAIR ABAD", "MAHESHPUR", "DHARAMPUR DHANSI GAIR ABAD", "NASRULLAPUR", "GHANSURPURAMROLI", "SHERNAGAR NARAINI", "RAMPUR MURAR", "SAIDPURA GAJJU", "KALALAN H.NO 172-END", "CHITAWAR AN", "SHARIFPUR KHORAJ", "MEHMUDPUR BHAVTA AN.H.NO 66-END", "ACHARJAN", "MIRZAPUR", "BHOGLI AN", "RANGDAN", "NAIMPUR GOVARDHAN", "PURSHOTTAMPUR BHAGWAN", "DAYAMNANGLA", "RAJOPUR SADAT AN", "RASULPUR JAGAN", "KHURRAMPUR KHADAK", "DELPURA KANNA", "MOHD PUR TAKI", "HAIJARPUR VEERCHAND", "VISHNOI SARAI AN.H.NO 181-END", "BHOJPUR", "ROSHANPUR PRATAP", "HALWAIYAN", "SARAI MEER AN H.NO 1-265", "MUGLAN PURVI", "MOJAMPURDAYAL", "KHUSHALPUR MATHERI H.NO 102-END", "ISLAMPUR HATTU", "AMBEDKAR NAGAR H.NO 233-324", "NAWADA", "MANJHLETA", "MALAKPUR LAKHMAN", "ALEYARPUR URF HAJIPUR", "AHMADKHEL AN.H.NO 1-137", "JHALRI", "MANIRAMPUR", "CHAJUPURA", "TODARPUR", "BEGAMPUR RUPCHAND", "RAIPUR GAIR ABAD", "CHATRIYANAGAR H.NO.325-877", "KARONDA CHODHAR AN.H.NO.1-130", "BUDGARI AN.H.NO.499-END", "MUFTIYAN", "PITHANHERIJHOJHA KHURD", "PREMPUR", "SAYADWARA", "MALAKPUR DEHRI", "BAGWANAN", "NASIRPUR MITHARI", "ANSARIYAN AN.", "KITHODA MUFTI", "TIGRI GAIR ABAD", "SULTANPUR SABHACHAND", "JEVANPUR", "SAHBAZPUR", "MAHALKI", "KHATAI", "SHEKHSARAI AN", "LADPURA AN.", "MAHAJANAN", "AJUPURACHOHAN", "SAHMUZAFARPUR", "KUMHERA", "DARGOHPUR", "HAKIKATPUR PRAYAG", "KHUSHALNAGAR GARHI", "RAMKAYRI", "BIHARIPUR", "JATI AN H.NO 46-END", "LAL SARAI AN.H.NO 103-355", "SHEKHPURA TURK H.NO 68-END", "HAKIKATPUR MUTHRA", "KAJIPUR PACHDU", "HINDU KATERA", "NALBANDAN AN.H.NO 142-END", "KARONDAPAHDU AN.H.NO 201-END", "BHARAIKI", "MUSLIM KATERA", "HAKIKATPUR VERCHAND", "MANPUR", "LALSARAI AN.H.NO 151-END", "BURHPUR NAIN SINGH", "MUBARAKPUR MEERA AN", "KHAKROBAN UTTARI", "MOTHEPUR KIRAT", "AURANGJEBPURHARDAS", "ASHRAFPUR", "SALEMPUR GADHELI", "DOODHLI", "MANAKCHAND", "VISHNOI SARAI AN.", "TAKIPUR NAROTTAM", "NAINPURA AN.", "MUSTAFAPURBHURAPURCHAKSAIDJAMAL", "GOPALPUR", "RAMPUR DAS AN.", "JAMALPUR BANGAR", "YARMOHDPUR PRITHVI", "KITHODA RAI", "KIRTO NANGLI", "MEHMUDPUR", "PIPALSANA", "SINGHADIYAN", "KAMALPUR  GAIR ABAD", "HASANLIPUR BHOGAN", "MOMINPUR DASU", "KADARPUR TAYYAB", "MUSTAFAPUR GAIR ABAD", "NANDPUR", "HADIPUR SADRUDDIN", "SHERPUR JAMAL GAIR ABAD", "MUBARAKPUR HADSA", "CHAMARAN", "SARAI BAHAUDDIN", "KALALAN AN.", "KAJISARAI AN.H.NO 1-73", "PITTHANHERI JIYA AN.H.NO 114-END", "MEMAN SADAT AN.H.NO 373-END", "TARIKAMPUR DARGO", "KAJIYAN AN.H.NO 128-END", "AHMAD KHEL AN.", "LUHARI SARAI AN.H.NO 276-END", "SULTANPUR SADAT", "AZMABAD", "MEHMUDPUR BHAVTA AN.H.NO 1-65", "PALDI", "ARJANIPUR", "PAKAHANPUR", "AURANGPUR FATTA", "MOHD ASGARPUR", "AURANGJEBPUR SAMBHAL", "MASURI", "RAMPUR ASHA", "MALKAN", "HAKIKATPUR GANGWALI AN.H.NO.1-209", "TAKIPUR TULSI NG.CH.", "MAHESHPUR", "MUKANDPUR RAMU", "BEGAMPUR HARREY", "MEHMUDPUR NARAYAN AN.1-102", "RASULPUR DAUD", "SAHLIPUR HEERA", "SAHZAHIR AN.H.NO 1-END", "AURANGJEBPUR FAJIL", "NAJIBPURSADAT", "SAHLIPUR AB.SATTAR", "LOHARAN", "BADSHAPUR MAIDAS", "SULTANKA GAIR ABAD", "PAHADIDDARVAJA", "KASSABAN", "NIJAMPURPADARATH", "BHAVANPUR KALA", "ARGUPURA GAIR ABAD", "SHADIPUR H.NO1-55", "ROSHANPUR CHIJROLI", "BAGHALA", "HASANLIPUR KHANDU", "RAJOPURSADAT AN H.NO 126-END", "JAFARPUR PARASRAM", "MAHESHPUR", "JATI AN H.NO 1-45", "ISLAMPUR BEGA", "PATERI", "PADHAN", "LOHARI SARAI PU.AN.H.NO 11/2-180/4", "MILKIYAN", "BUDHAWALA", "CHITAWAR AN", "SAIDPURA GAJJU", "MUGLAN PASCHIMI", "KOTWALI AN H.NO 1-476", "VISHNOI SARAI AN.H.N0.1-180", "BADSHAPUR MAIDAS", "SAHABPUR URF HINDU PUR H.NO 151-END", "AKBARABAD AN0", "MALAKPUR ABDULLA", "SARAI MEER AN H.NO 266-END", "SHERNAGAR NARAINI", "RAMPUR MURAR", "GHANSURPURAMROLI", "DHARAMPUR DHANSI GAIR ABAD", "NASRULLAPUR", "SAHPUR SUKKHA", "HARBANSHPUR DHARAM", "BAHADARPUR SHARFUDIN HUSAIN AN.", "MOMINPUR DASU", "AJAMPUR YAR MOHD", "MIRDGAN", "RATANPUR", "MUSTAFAPUR GAIR ABAD", "AFGANAN AN.", "KHURRAMPUR DALLU", "DUDHLA", "GUNIYAPUR", "ALEYALIPUR", "KHURRAMALI SARAI AN.H.NO 116-END", "SARAI MEER AN0 H.NO01-116", "BRAHAMJITPUR CHANDA", "GARABPUR", "BHANERA", "KALAPUR BUJURG", "ज्ञKOTWALI AN H.NO 124-END", "JHANDA AN", "RAFIKPUR", "SULTANPUR SADAT", "MIRZALIPUR BHARA", "KHANPUR", "SAHLIPUR ASHA", "KAKADPUR", "SHADIPUR IMMAR", "DULHAPUR GAIR ABAD", "CHATRIYANAGAR H.NO.325-877", "BEGAMPUR RUPCHAND", "RAIPUR GAIR ABAD", "TODARPUR", "BEHADA CHOHAN", "MANIRAMPUR", "BUDHPUR MUFTI", "CHAJUPURA", "SALEMPUR GADHELI", "JHALRI", "JEVANPUR", "METHA SAID", "SULTANPUR SABHACHAND", "ANSARIYAN AN.", "TIGRI GAIR ABAD", "KITHODA MUFTI", "KASBA KOTRA AN.H.NO 1-100", "BAGWANAN", "MALAKPUR DEHRI", "AMIRPUR MAIDAS GAIR ABAD", "SAYADWARA", "PITHANHERIJHOJHA KHURD", "PREMPUR", "HAIJARPUR VEERCHAND", "SAHABPUR URF HINDUPUR H.N 1-150", "JATAN", "MOHD PUR TAKI", "DELPURA KANNA", "HAKIKATPUR PRAYAG", "KOTWALIAN H.NO 413-END", "MOHD ASHIKPURKAMALNAIN", "RAMKAYRI", "DAYAMNANGLA", "CHIPIWADA", "BIHARIPUR", "NAIMPUR GOVARDHAN", "RANGDAN", "PURSHOTTAMPUR BHAGWAN", "ACHARJAN", "MEHMUDPUR BHAVTA AN.H.NO 66-END", "SHARIFPUR KHORAJ", "BHOGLI AN", "NAWADA", "MANJHLETA", "GOKALPUR SUNDAR", "VALIPUR AN H.NO 150-END", "HALWAIYAN", "SARAI JALAL GAIR ABAD", "ISLAMPUR SAHU", "HAKIKATPURGANGWALI AN.H.NO210-END", "BHOJPUR", "KAJISARAI AN.", "SHISHGRAN AN.H.NO 122-END", "CHANDUPURA NASEBPUR KANNA", "KHAKROBAN ST.", "MALKAN", "BARHAPUR", "NOORPUR HATTI", "BEGAMPUR HARREY", "AURANGJEBPUR FAJIL", "SAHLIPUR HEERA", "SAHZAHIR AN.H.NO 121-END", "KARONDA PACHDU AN.", "MUBARAKPUR KHOSA", "JALPURA", "SARAI JEEVAN", "TARIKAMPUR GAIR ABAD", "JAMALPUR KHOKO", "CHODHRANA AN0H.NO 125-END", "KALA NANGLI GAIR ABAD", "LUKMANPURA", "KASSABAN", "PADLA AN", "BADSHAPUR MAIDAS", "MOJAMPUR HARVANSH", "ALHADADPUR", "KISHANPUR", "MOMINPUR LALU", "BAKARPUR JAIRAM", "SULEMASIKHOPUR", "SHARIFULMALAKPUR", "UMRI AN.H.NO 158-END", "MOHIUDDINPUR", "RAMPURLASHKARI", "HAMIDPUR VIDHICHAND GAIR ABAD", "ALIPURMAN", "PITHANHERI JHOJHA KALA", "SAIDPURIMAHICHAND", "MAHESHWARI JAT AN0H.NO 131-END", "AZMABAD", "SULTANPUR SADAT", "MOHSANPUR", "AHMAD KHEL AN.", "SAHUWAN", "KASBA KOTRA AN. H.NO .", "BANKHALA", "HAIBATPUR PEERU", "DAYALPU GYANPUR", "NALBANDAN AN.H.NO 1-141/2", "SHADIPUR AN H.NO 56-END", "ASDULLAPUR PRITHVI", "VALIPURA AN.H.NO 1-149", "MANIHARI SARAI AN H.NO 147 -END", "HAKIKATPUR VERCHAND", "NURULDEHARPUR URF MIRZAPUR", "MUKIMPUR PADARATH", "MOTHEPUR KIRAT", "MUBARAKPUR MEERA AN", "NANGLA ISLAM", "BEHADA CHOHAN", "SALEMPUR GADHELI", "SAINAWALA", "PURSHOTAMPUR URF THEY PUR", "CHOHANAN AN.", "YARMOHDPUR PRITHVI", "MATHERA CHOHAN", "LUHARI SARAI H.NO 181-END", "GOSPUR RAI GAIR ABAD", "ALAMPUR", "GANORA", "HARDASPUR MADE", "SARAI MEER AN H.NO 117-END", "AMIRPUR MAIDAS GAIR ABAD", "MO.ALIPUR SUKHANAND", "AKBARABAD  AN0", "KHAKROBAN", "HAKIKATPUR PRAYAG", "AURANPUR NANDLAL", "KHUSHALNAGAR GARHI", "TUKHMAPUR HARVANSH AN.", "BUDGARI AN.H.NO.1-498", "NARULLAPUR BADLI GAIR ABAD", "SABHACHANDPUR MOHAN URF MAKHWADA", "BIHARIPUR", "MO ALIPUR SHEKH", "RAMKAYRI", "MAKSUDNAGAR DEVIDAS", "MEMAN SADAT AN.H.NO 1-372", "GALIBALIPUR MALUK", "MAHAPUR MOHD ALI", "VISHNOI SARAI AN", "GAZIPUR AN", "BHUDDI AN.H.NO 1-142", "BHAGWANPUR", "KHUSHHALPUR MATHERI H.NO 1-101", "SHEKHPURA NOABAD", "HASALIPUR MUTHRA", "JAGANNATHPUR", "HUSAINPUR MEERA", "SAMASPUR VEERBHAN", "SABHACHANDPUR", "ISLAMPUR SAHU", "BHARAIKI", "MAMRAJPUR", "SAIDPURI", "CHAHRONAK", "SAHJANPUR ROSHAN", "AFGANAN", "PADLI", "KILA", "NAWADA", "HEEMPUR MANAK", "CHANDAK JATT", "AMBEDKAR NAGAR H.NO 233-324", "KAMKARAN AN.", "MUGLAN PURVI", "ROSHANPUR PRATAP", "ISLAMPUR SAHU", "SHADIPUR IMMAR", "AURANGJEBPUR MEHMUD URF MANGOLPURA", "TODARPUR", "DELPURA NAINU", "BUDHPUR MUFTI", "ISLAMPUR SAHLI", "BEHADA CHOHAN", "SAHZAHIR AN. H.NO 1-120", "AURANGJEBPUR GULAL", "RAJARAM PUR PRATAP", "SAHABPURA RATAN SINGH", "KAJIPUR IMMA", "AURANGPUR HIRDAY", "KAYASTHAN", "CHAKNARANGI", "JATPURA", "NASIRPUR MITHARI", "KASBA", "ABULFAJALPUR KHAS", "BUDHUPADA", "AMIRPUR MAIDAS GAIR ABAD", "ABDIPUR KALYAN", "AKBARABAD AN.H.NO 1-548", "SHEKH SARAI", "CHANDAK TURK", "MUSTAFAPUR GAIR ABAD", "MOMINPUR DASU", "HAMIDPUR MAKHAN", "RAMPURCHAJMAL", "SABHACHANDPUR KESHOAN", "KOTWALI AN H.NO 477-END", "RAMPUR BISHNA", "BILASPUR GAIR ABAD", "SHADIPUR AN", "NARAYANPUR", "LAL SARAI AN.H.NO 356-END", "TAKIPUR MANOHAR", "BASEDA", "BANWARIPUR", "BHANERA", "KAKADPUR", "AJAMPUR MOHAN", "MAHESHPUR", "AKABRAN", "HASANPURA", "JAFARPUR PARASRAM", "PADHAN", "HUSAINPUR SULTAN", "JATI AN H.NO 1-45", "SAHLIPUR HEERA", "LOHARI SARAI PU.AN.H.NO 11/2-180/4", "MUTHRAPUR URF KALUWALA", "ANANDKHERI", "MUBARAKPUR SAHARAN", "MILKIYAN", "MUGLAN PASCHIMI", "SAMASPUR KHORA URF KHODPURA", "MOJIPUR DHARMA", "KAJIYAN AN.H.NO 1-127", "MEMANSADAT AN.H.NO0264-END", "NAKIMPUR RAMRAI GAIR ABAD", "ALIPUR GANGA", "AMBEDKAR NAGAR", "DHAKI SADHO", "RAMPUR MURAR", "BHUDDI AN.H.NO 143-END", "ISLAMPUR MEERA", "VAISHPUR KUDIYA", "HAKIKATPUR GOVIND", "KAMALPURPURAINI", "UMARPUR", "MUBARAKPUR MEERA AN.", "GOVINDPUR", "JAFRULLAPUR", "SARAIMEER AN H.NO 144-END", "MEHMUDPUR NARAYANAN.H.NO 103-END")
		case "Barhapur":
			sections = append(sections, "CHATARPUR", "KALLUWALA", "RAJPURKOT", "NOABAD DESH", "ANWARPUR CHANDIKA", "MEDPURASULTAN", "KASSABAN", "AJAB NAGAR", "DHARMAWALA", "HASANLIPURDHARMA", "RAIPURI", "ALHAYPUR BEERAN", "NARAYANWALA", "CHANDLIKA", "RAI BHAN WALA", "HEERAPUR GUKAL", "LAKHIWALA", "MADHOWALA", "JANULABEDDIN AN0", "DESONDIWALA DESH", "AJUPURARANI", "PAGAMBERPUR", "MODPURTIRLOK AN HNO 164 TO END", "JABTA NAGAR", "KASAMPUR GARHI AN0", "SHAHLIPUR GADDU", "NAYAK SARAI", "UDHO JOT", "TARAGANG", "SHAHPUR JAMAL A", "MANKHANDEPUR GARHI", "AASFABAD CHAMAN", "BHAGIJOT", "SHAHLIPUR", "REHMAPURUSMAPUR", "UMARPUR NATTHAN", "WAJIDPUR", "JASSUJOT", "MH. PUR BHAJJA", "AURANGABAD", "NEMATPUR", "SARDARPUR", "DAUDPUR", "PAYBAND KHERI", "FATHAY ULLAH NAGAR", "FATTUWALA", "KIRATPUR", "SABUWALA", "FAJULLAPURMAHESH", "ISLAM NAGAR AN0", "LADPUR", "TARAGANG", "PATPADA KESHO", "JAMALPUR DHIKLI", "MITHOPUR", "NURULLAPUR", "ALIGANG", "PAYBAND KHERI", "TIGRIANOOP SINGH", "CHAKUDAYCHAND", "TAMURABAD", "ASDULLAPUR KHALSA", "MAHENDER NAGAR", "MOHD ASHIKPUR BHUREY", "ABDULREHMANPUR JEVAN", "DHARMPUR", "SHAHLIPURKOTRA  H NO 179 TO END", "HEEMPUR MAKHDUM JAHAN", "CHAKMANI RAGHUNATHPUR", "DHARMPUR URF JOGIWALA", "BAHEDI", "GANGARAMWALA", "D.S.I.L DWARIKESH NAGAR", "MAJHEDASAKRU HNO 1 TO 227", "DEEVANANDPUR GARHI", "CHAPRHA", "SATTIYAN", "RAI BHAN WALA", "RANI NAGAL", "SAIDPUR HARIAWALA", "GADLA", "HALLOWALI", "NOIABAD JANGAL", "CHAKARPANPUR", "SHARGHAR AN0 HN0 1 TO 376", "MADHOWALA", "AMEERPUR MANJU", "FATEPUR DHARA", "PAGAMPUR URF BAGIR GANG", "BARKATA", "NATHEYWALI", "GANGARAMWALA", "SIKKAWALA", "DHARMPUR URF JOGIWALA", "BAHEDI", "NEMATPYRISKUPURA", "PREMPURI", "NEJO SARAI", "FIROZPUR", "KASHIWALA", "TELIPADA", "GOVINDPUR", "DEEVANANDPUR GARHI", "KHANAWALUDALLAWALI", "BUNDKI", "MURTJAPUR URF PUNAPUR AN H.NO 1 TO 220", "MUSSAYPUR", "KANAKPUR", "RASULPUR METHEY", "CHAPRHA", "TAILIYAN", "JALALPUR SULTAN", "RAMPUR KALA", "JAGANNATHPURGOKAL", "RANI NAGAL", "GHASIWALA", "AURANGJEBPURSAHLI", "SHADIPUR", "KAYAMNAGARDODHRAJPUR", "GADLA", "GOPIWALA", "HARGANPUR H NO 142 TO END", "MAJRA BHARATPUR", "TARAGANG", "NATHEYWALI", "MANKHANDEPUR GARHI", "AFGALPUR HARRAY", "SHANKARPUR", "MOMIN NAGAR", "MOHDPUR ROJORE HNO 1 TO 254", "JALALPUR", "MOHDPUR ROJORE HNO 255 TO END", "MURADNAGAR", "ALIGANG", "KASHRIPUR", "RAMPUR", "NARAYANPUR RAMJI", "AJIMULLANAGAR AN0", "AGWANPUR", "SIRYABALI", "TAMURABAD", "RASULPUR ABAD AN0", "JAMANPUR", "MOINUDDINPUR", "PADARATHPUR", "KALLU NAWADA", "NOORPURARAB", "BHAGWANPURPRATAP", "BEDARBAKTHPURJADO", "MEERPUR GHASI", "BHAJDAWALAKHALSA", "HARKISANPUR JANULABDEEN", "SAHUWALAGOSAIWALA  HNO 1 TO 105", "SAMASPURSADDO", "TAKIPUR BEGA", "CHAK", "SHAHALIPUR", "RASOOLPURMUZAFAR H.NO 1 TO 954", "ALINAGAR", "AJIMULLANAGAR AN0", "AGWANPUR", "AASFABAD CHAMAN", "JAY SINGH JOT", "RAMPUR", ".JATNANGLA", "AMAN NAGAR", "ISMAILPURSALI BA", "MANAVARPYURSAID HNO 111 TO END", "KALLU NAWADA", "DAULLATPUR", "SHAJADPUR AN0", "SHARGHAR AN0", "REHAR AN0", "KHURALLAPUR DESH", "FATHAY ULLAH NAGAR", "FATTUWALA", "KIRATPUR", "BHAGWANPURPRATAP", "TANDA BAIRAGI", "KHERABAD", "CHAMRUNAWADA", "SHAHJAHAN PUR GARHI", "SULEMASEKOPUR", "NOMI UTTARI", "RAJUPURA H.NO 1 TO 239", "RAIPURI", "UMARPURBERKHERA", "GOVINDPUR", "FAZALPUR", "HARGANPUR H NO 1 TO 142", "LAKHIWALA", "RAIPUR SADAT AN", "JAYNAGAR", "NOAABAD JAGEER", "DHAKYA", "MANKHANDEPUR GARHI", "MANNABARPUR CHAK", "HIDAYATPUR AN0", "SARFUDEENNAGAR AN", "RAMNAGARGOSAI", "RAMPUR GARHI", "SIKKAWALA", "SHAHJAHAPUR", "GANGARAMWALA", "SAEED NAGAR", "KASHIWALA", "HASANLIPUR HEERA", "PREMPURI", "KHEJARPUR", "ALHAHEDI", "MUSSAYPUR", "BAZAR", "HAJIMOHD PUR KOTKADAR  AN 1 TO 110", "DEEVANANDPUR GARHI", "LALPURI", "DALLIWALA", "JASPUR", "SAIDPUR HARIAWALA", "KUDDUSPUR", "NAIWALA", "KANHADAWALA", "CHAKARPANPUR", "PAGAMPUR URF BAGIR GANG", "HUSAINABAD AN.H.NO 1-277", "MAHABATPUR", "FATEPUR A", "MOMIN NAGAR", "RAIPURSADAT AN0 445 TO END", "ALHAYPUR MOHKAM", "SEEPAH", "MH. PUR URF JOGIPURA", "ALAMPUR", "MANSHA", "SUAWALA AN0 H.NO 1 TO458", "PHOOLBAG", "GUSAIWALA", "DALLAWALA", "ALIGANG", "KISANPURKUNDA", "CHAKUDAYCHAND", "QILA H NO 81 TO END", "BHOGPUR HNO 1 TO 548", "BHAWANIPUR AN0", "NANJIWALA URF RAMJIWALA", "SAHARAWALA", "SEER SHIVRAJ SINGH", "ABDULREHMANPUR JEVAN", "MOHD ASHIKPUR BHUREY", "SARAI PURAINI", "HARIJAN BASTI", "CHANDPUR", "MEERPUR GHASI", "PALPUR", "DHARMPUR", "AFJALPIRBHAU", "BHAJDAWALAKHALSA", "SULTANPUR", "KHANPURMANAK", "TURATPUR", "SADA SHIVPUR GABRI", "JASSUJOT", "HAREBALI AN0", "ALINAGAR", "REHMAPURUSMAPUR", "GUSAIWALA", "JHULA", "MH. PUR URF JOGIPURA", "SHIVPURI", "PARMANANDPUR GAWDI", "MH. PUR BHAJJA", "MURTUZABAD", "NANJIWALA URF RAMJIWALA", "HASANLIPURMAAN", "FATTUWALA", "CHAMARAN", "HAJIMODPUR KOTKADAR 111 TO END", "MURADBAKSHPUR", "BADSHAPUR", "RAIPUR SADAT H.NO1 TO 444  END", "ALIPUR MITTHAN", "NOABAD DESH", "RAJPURKOT", "RASULPUR ABAD", "NOMI SOUTH", "DHARMAWALA", "MOHIDINPUR", "SHAHJAHAPUR", "CHANDLIKA", "LALPURI", "BABLI", "RAIPURI", "ALHAYPUR BEERAN", "SEERWASHU CHAND AN0", "DOARKADASPUR", "HASANLIPURDHARMA", "MAGHPUR", "HEERAPUR GUKAL", "HEERAPUR MALI", "KANHADAWALA", "JAFARPUR", "LAKHIWALA", "pipli", "MOHD ALIPUR PATTAGHAT", "NOAABAD JAGEER", "DESONDIWALA DESH", "JASPUR", "MORGHANDI", "BHAGWANPUR  HN0 126 TO END", "KUJAITA", "HIDAYATPUR AN0", "KASAMPUR GARHI AN0", "SULTAN AJAMPURBANI GNESH", "nawabpura", "BINJHAHEDI H.NO 101-END", "PAGAMBERPUR", "ISLAMABAD AN0 20-394", "SULEMASEKOPUR", "ABDIPURHARVANSH", "HAJIMODPUR KOTKADAR AN0 H.NO 1 TO 326", "DHARMAWALA", "VEERBHANWALA", "ISLAMABAD AN0 1 TO END", "MEERAPUR VEERAN", "SHAHALIPUR JAMAN", "MOHD. ALIPUR CHOHAR", "TANDAMAIDAS AN0 1 TO 794", "ALHAYPUR BEERAN", "GHASSIWALA", "HEERAPUR GUKAL", "RASOOLPURMUZAFAR H.NO 9 TO 276", "GOHAR ALI KHAN", "NOAABAD JAGEER", "SUNDARWALI", "SHALIPUR BA", "BHOGPUR HNO549 TO END", "MIYAZI MOKHA", "DLEEPUR", "ANOOPNAGAR", "HIDAYATPUR AN0", "TAYABPUR", "ALAUDEENPUR", "AALAMPUR GABRI AN0", "AJAMPURMANNI", "BHAJDAWALAKHALSA", "MUKANDOURRAHMAL", "MORGHANDI", "BHABANIPUR", "RAIPUR SADAT AN H.NO 201", "ABHAYRAJPUR", "JASSUJOT", "AASFABAD CHAMAN", "KADARGANJ", "ALINAGAR", "AANANDIPUR", "PATHANAN", "HARPUR", "GOSPUR SADAT", "NOMI MEDIUM", "LALPURI", "ALIPURMOLAD", "SHAHJAHAPUR", "MUTAAZ HUSAIN AN0", "KRISHNA NAGAR H.N0 493TO 512", "SABDALPURDEJAHGIR", "FATHAY ULLAH NAGAR", "KIRATPUR", "MORGHANDI", "WANSIJOT", "MOMIN NAGAR", "SHANKARPUR", "CHHAJMAL WALA", "AFGALPUR HARRAY", "JOJHIYAN", "JHILMILA AN.H.NO 1-85", "MURADNAGAR", "BINJHAHEDI 1 TO 100", "RAM NAGAR", "CHOHADHPUR NANU", "BAHADAR NAGAR", "BANIRAMDASS", "KASHRIPUR", "TAMURABAD", "KUAKHERA H NO 1 TO 176", "CHAKUDAYCHAND", "HASANPUR", "ABABKARPUR", "BHAGWANPUR  HN0 1 TO 125", "JHADPURA", "ABDULREHMANPUR JEVAN", "MOHD ASHIKPUR BHUREY", "WAJIDPURLALA", "MEERPUR GHASI", "DULICHANDPUR", "ASLAMPURBHULLAN", "DHARMPUR", "BAHEDI", "DHARMPUR URF JOGIWALA", "SIKKAWALA", "ABILFAJALPURBANI", "SHYAMPURSARVANKA", "DAHLAWALA", "KASHIWALA", "GAJIPUR BHAGWAN", "HAJIMODPUR KOTKADAR H.NO 440 TO END", "PREMPURI", "MUSSAYPUR", "KANAKPUR", "RAMPUR ASAF", "CHAPRHA", "BUNDKI", "RANI NAGAL", "GHASSIWALA", "MAJHEDASAKRU HNO 228 TO END", "KHOKRA", "TAILIYAN", "RASULPUR VEERAN", "RAMPUR KALA", "GADLA", "SEERBASHU CHAND AN0", "SALABATPUR GIRDHAR", "SAIDPUR HARIAWALA", "MH. PUR URF JOGIPURA", "MODPURTIRLOK AN HNO 1 TO 163", "KADARPURNANU", "PAGAMPUR URF BAGIR GANG", "KALYANPUR H.NO 1 TO END", "BASUWALA", "NATHEYWALI", "HARISINGHKABHOGLA", "DLEEPUR", "JASSUJOT", "KISANPURKUNDA", "BAHADAR NAGAR", "PHOOLBAG", "SUAWALA AN0 H.NO 1 TO458", "PREMNAGAR", "DALLAWALA", "MANSHA", "SEEPAH", "FAJULLAPURMAHESH", "JAMALPUR DHIKLI", "FATEPUR A", "HUSAINABAD AN.H.NO 1-277", "PALPUR", "DHARMPUR", "MURTJAPUR URF PUNAPUR AN H.NO  221", "HARIJAN BASTI", "ALLAHDADPURKHAJWA", "MOHD ASHIKPUR BHUREY", "SARAI PURAINI", "KOTKADAR HAJI MOHDPUR H.NO 327 TO END", "SEER SHIVRAJ SINGH", "ABDULREHMANPUR JEVAN", "MIRZALIPUR CHOHAD", "SAHARAWALA", "ASDULLAPUR KHALSA", "KAMRUDINNAGAR", "BHAWANIPUR AN0", "MAKSHOODABAD", "PAYBAND KHERI", "HERWELA", "CHAKUDAYCHAND", "LALPURI", "RASULPUR VEERAN", "MAJHOLI H NO 1 TO 50", "SADATPURGADI", "DALLIWALA", "ALHAYPUR BEERAN", "SHAJADPUR", "FAGLABAD PARMANANDPUR ANO", "NOABAD DESH", "KHEJARPUR", "SAEED NAGAR", "MADPURI", "GOWARDHANPUR AN0", "SHAHJAHAPUR", "CHAKMANI RAGHUNATHPUR", "TARAPUR", "RAIPURSADAT  H.NO 1 TO 200", "PAGAMPUR URF BAGIR GANG", "CHAKARPANPUR", "KANHADAWALA", "SUAWALA AN0 H.NO 251 TO 594", "SAIDPUR HARIAWALA", "NOIABAD JANGAL", "JASPUR", "RAI BHAN WALA", "GAJIPURSABHACHAND", "HEERAPUR GUKAL", "E", "GHASSIWALA", "ALHAYPUR BEERAN", "HASANLIPURDHARMA", "ANWARPUR CHANDIKA", "JHILMILA AN.H.NO 86-END", "DHARMAWALA", "ASAFPURSAIDPEER", "NOABAD DESH", "PAGAMPUR URF BAGIR GANG", "KASAMPUR GARHI AN0", "SAMASPURNASEEB H.NO 151 TO END", "CHAMPATPURCHAKLA", "AFGALPUR HARRAY", "QILA 1 TO 80", "DLEEPUR", "DESONDIWALA DESH", "MOHD ALIPUR PATTAGHAT", "SAIDPUR HARIAWALA", "pipli", "MAHARATPURKALA", "HEERAPUR MALI", "MADHOWALA", "MH. PUR URF JOGIPURA", "JHULA", "MURTUZABAD", "MH. PUR BHAJJA", "GUSAIWALA", "AJIMULLA NAGAR", "REHMAPURUSMAPUR", "JASSUJOT", "TANDAMAIDAS AN0 1 TO 150", "SADA SHIVPUR GABRI", "RATUPURA", "TARAGANG", "ALIPUR MITTHAN", "ABDULREHMANPUR JEVAN", "MOHD ASHIKPUR BHUREY", "MURADBAKSHPUR", "HARPUR", "JHANABAD AN 1 TO 875", "NANJIWALA URF RAMJIWALA", "AANANDIPUR", "MAHARATPURKALA AN0 H.NO 1 TO 147", "GUSAIWALA", "AJAMPURPARMA", "AJIMULLANAGAR AN0", "RAWALHERIKHAJIRI", "ALINAGAR", "KADARGANJ", "TIRLOKPUR", "MUKANDOURRAHMAL", "MOMIN NAGAR", "AFGALPUR HARRAY", "BEGAM SARIA", "BHAJDAWALAKHALSA", "HAJIMODPUR KOTKADAR H.NO 1 TO 439", "KHERABAD", "SARAI DADUMBAR", "TANDA BAIRAGI", "MEERPUR GHASI", "SABDALPURDEJAHGIR", "ALIPURA JATT", "KHURALLAPUR DESH", "KRISHNA NAGAR H.N0 493TO 512", "MUTAAZ HUSAIN AN0", "SAMASPUR NASEEB H.NO 1 TO 150", "ALIPURMOLAD", "KAHRIPUR JANGAL", "NANJIWALA URF RAMJIWALA", "HARPUR", "GOSPUR SADAT", "FAZALPUR", "GOVINDPUR", "GHASSIWALA", "MEERAPUR VEERAN", "SHAHALIPUR JAMAN", "KALLUWALA AN0", "SULTANPUR KHAS", "ISLAMABAD AN0 20-394", "HAJIMODPUR KOTKADAR H.NO 1 TO 923", "SULEMASEKOPUR", "MAHARATPURKALA AN0 H.NO 1 TO 684", "TAYABPUR", "MANNABARPUR CHAK", "HIDAYATPUR AN0", "DHAKYA", "DLEEPUR", "HUSAINABAD AN.H-NO 278-END", "MIYAZI MOKHA", "SUNDARWALI", "NOAABAD JAGEER", "pipli", "SHALIPUR BA", "RASOOLPURMUZAFAR H.NO 9 TO 276", "KHOKRA", "JAGANNATHPURGOKAL", "RASULPUR VEERAN", "LALPURI", "TAILIYAN", "KHANAWALUDALLAWALI", "TELIPADA", "RASULPUR METHEY", "MUSSAYPUR", "BUNDKI", "GAJIPUR BHAGWAN", "PREMPURI", "SHYAMPURSARVANKA", "ABILFAJALPURBANI", "KASHIWALA", "FIROZPUR", "SIKKAWALA", "HARISINGHKABHOGLA", "MANKHANDEPUR GARHI", "KADARPURNANU", "KALYANPUR H.NO 1 TO END", "KANHADAWALA", "JASPUR", "KASHRIPUR", "RAMPUR", "BAHADAR NAGAR", "BANIRAMDASS", "AGWANPUR", "CHOHADHPUR NANU", "SUAWALA AN0 H.NO 1 TO458", "MOHDPUR ROJORE HNO 255 TO END", "JAFRABAD", "JOJHIYAN", "BHAJDAWALAKHALSA", "MOMIN NAGAR", "AFGALPUR HARRAY", "MORGHANDI", "WANSIJOT", "SAHUWALAGOSAIWALA  HNO 106 TO END", "WAJIDPURLALA", "MEERPUR GHASI", "ASLAMPURBHULLAN", "NOABAD DESH", "BHAGWANPURPRATAP", "SHAHALIPUR KOTRA AN0", "SAHARAWALA", "MEERAPUR MODIWALA AN0", "KALLU NAWADA", "ABABKARPUR", "RAFAYATPUR HULLAS", "FAGLABAD PARMANANDPUR", "TANDAMAIDAS AN0 151 TO END", "GUSAIWALA", "AURANGABAD", "MH. PUR URF JOGIPURA", "GADLA", "ALIGANG", "SHAHLIPUR", "AJIMULLANAGAR AN0", "MAHAPUR", "ALIYARPUR", "TARAGANG", "SHAHLIPUR GADDU", "NAYAK SARAI", "UDHO JOT", "MURADBAKSHPUR", "TANDA BAIRAGI", "LAL SARAI HNO101 TO END", "FATTUWALA", "KHURALLAPUR DESH", "CHAKCHOHADMAL URF BHAJRHAWALA", "DAUDPUR", "MAKHDOOMPUR", "NANJIWALA URF RAMJIWALA", "DEEVANANDPUR GARHI", "FAZALPUR", "GADRIYAN", "GOVINDPUR", "AARAJI KANDALA", "RAI BHAN WALA", "NARAYANWALA", "CHANDLIKA", "RAIPURI", "AJAB NAGAR", "KHUSHALPUR DEWAN SINGH", "MAHARATPURKALA AN0 H.NO 148 TO END", "ANWARPUR CHANDIKA", "HARKISHANPUR", "GANGARAMWALA", "KASSABAN", "MEDPURASULTAN", "RAJPURKOT", "THAKURAN", "CHIRANJI LALA", "PAGAMBERPUR", "BAWAN SARAI", "MODPURTIRLOK AN HNO 164 TO END", "MANNABARPUR CHAK", "DHAKYA", "pipli", "MOHD ALIPUR PATTAGHAT", "JANULABEDDIN AN0", "LAKHIWALA", "MADHOWALA", "BARKHERA", "RAIPURI", "MANAVARPYURSAID HNO TO 1 TO 110", "JAGANNATHPURGOKAL", "LALPURI", "JAHANGURPURKHASS", "RASULPUR METHEY", "MUSHTAFFABAD GARHI", "DEEVANANDPUR GARHI", "AJIMULLA NAGAR AN0", "KHANAWALUDALLAWALI", "MAJHEDASAKRU HNO 37 TO 454", "D.S.I.L DWARIKESH NAGAR", "CHAKMANI RAGHUNATHPUR", "SHAHJAHAPUR", "SUAWALA AN0 H.NO 1 TO 408", "ISMAILPURDAMI", "GANGARAMWALA", "NANGLA LADDHA", "SEERBASHU CHAND", "MANKHANDEPUR GARHI", "BARKATA", "IMAMPURTIGRI", "SHARGHAR AN0 HN0 1 TO 376", "LAKHIWALA", "KANHADAWALA", "JASPUR", "HALLOWALI", "NOIABAD JANGAL", "BHATT PURA", "AGWANPUR", "RAMPUR", "NURULLAPUR", "ALIGANG", "ALAMPUR", "MITHOPUR", "SUAWALA AN0 H.NO 1 TO458", "PHOOLBAG", "PATPADA KESHO", "FAJULLAPURMAHESH", "SHAHLIPURKOTRA  H NO 1 TO 178", "MAJHOLI H NO 51 TO END", "HARIJAN BASTI", "BHAGWANPURPRATAP", "KALLU NAWADA", "MAHENDER NAGAR", "SAHARAWALA", "TIGRIANOOP SINGH", "PAYBAND KHERI", "MOHDPUR ROJORE HNO 1 TO 159", "MAHARATPURKALA AN0 H.NO 685 TO END", "NARAYANPUR RAMJI", "BAHADAR NAGAR", "AGWANPUR", "RAMPUR", "JATPURA", "MURADNAGAR", "MOHDPUR ROJORE HNO 255 TO END", "MORGHANDI", "JALALPUR", "SHANKARPUR", "MACHHMAR", "SHANAGAR KURALI", "BHAGWANPURPRATAP", "ADREHMANPURMEHJANG", "KIRATPUR", "HUSAINABAD AN. H.NO.1-277", "KALLU NAWADA", "JAMANPUR", "TAMURABAD", "SIRYABALI", "RAIPURI", "PAYBAND KHERI", "ABABKARPUR", "RANI NAGAL", "RAMPUR KALA", "JALALPUR SULTAN", "RASULPUR VEERAN", "JAGANNATHPURGOKAL", "RASULPUR METHEY", "RAMPUR ASAF", "CHAPRHA", "KANAKPUR", "TELIPADA", "KHANAWALUDALLAWALI", "NEJO SARAI", "NEMATPYRISKUPURA", "BAHEDI", "DHARMPUR URF JOGIWALA", "CHAKMANI RAGHUNATHPUR", "NATHEYWALI", "MANKHANDEPUR GARHI", "TANDAMAIDAS AN0 1 TO 452", "BENIPUR", "PAKKABAG", "MAJRA BHARATPUR", "GOPIWALA", "AURANGJEBPURSAHLI", "KAYAMNAGARDODHRAJPUR", "GADLA", "GHASIWALA", "GOVINDPUR", "FAZALPUR", "RAI BHAN WALA", "CHAPRHA", "MEERAPUR VEERAN", "RANI NAGAL", "GHASSIWALA", "MUNIR NAGAR", "RAJUPURA H.NO 1 TO 239", "SHAHJAHAN PUR GARHI", "CHAMRUNAWADA", "ABDIPURHARVANSH", "ABDULLAPURKURASHI", "RAMPUR GARHI", "SARFUDEENNAGAR AN", "TANDAMAIDAS AN0 1 TO 775", "DLEEPUR", "JAGATCHANDPUR", "NATHEYWALI", "DHAKYA", "MANNABARPUR CHAK", "IMRATPUR", "ALIGANG", "GADLA", "MADHOWALA", "HAMJAPURKATHAUR", "NEJO SARAI AN0 H N0 1 TO 70", "KUAKHERA H NO177 TO END", "AMAN NAGAR", "ISMAILPURSALI BA", "NEJO SARAI AN0  H NO 71 TO END", "KADARGANJ", "AASFABAD CHAMAN", "AJIMULLANAGAR AN0", "TAKIPUR BEGA", "SHAHALIPUR", "RASOOLPURMUZAFAR H.NO 1 TO 954", "CHAK", "HARVANSGWALA", "MAHSANPUR", "NARAYANPUR", "JAGANNATHPURMEHRU", "HARKISANPUR JANULABDEEN", "LAL SARAI HNO 1 TO 100", "TANDA BAIRAGI", "KHERABAD", "KIRATPUR", "INAYATPUR H.NO 1 TO END", "KHURALLAPUR DESH", "FATHAY ULLAH NAGAR", "PUKHAY WALA", "JAHANABAD", "DAULLATPUR", "HARPUR", "DEEVANANDPUR GARHI")
		case "Dhampur":
			sections = append(sections, "GAJARAULA B.A", "MAUHADI", "SAMANASARAY AN.", "NURAPUR CHHIBRI AN.", "DHURARA", "SADADOBAI AN.", "BHAJDI SARAY", "UMRAPUR KHADAR", "NOORPUR CHHIBRI AN.", "LOHIYAN AN.", "SALARAPUR  B.A", "PAHADI DARAVAJA AN.", "BAJIRAPUR URF MADHIYA", "ABUNASRAPUR", "MOHD.ALIPUR INAYAT", "LAIDERPUR", "HAJIPUR", "RAMATHERA", "HARRA AHAMADPUR JALAL", "HAFIJABAD SHAKI AN.", "SHAFIYABAD AN.", "BAKAR KASSABAN", "MOHD.PUR SULTAN", "SAIDAPUR BIRU", "RASULPUR MAOH KULI . AN.", "MANAPUR SHIVAPURI", "KESHOPUR", "MU.CHAUDHARIYAN AN.", "SALAVA", "SHAHAJADAPUR TARU", "MOHD.ALIPUR LALMAN URF RAMSAHAYVALA", "MUBARAKPUR TAPPA BHARTI B..A", "JAMALAPUR UDYACHAND AN.", "NAI SARAY", "NANGALA GANNA", "GUJARATIYAN", "AMAKHEDA SHAJRAPUR AN.", "SADARNAPUR", "SARKATHAL SAHANI AN.", "NAYA GANVA MAJRA", "PALKI EMAN", "AURANGASHAHAPUR NARAYAN", "NANDAGANVA", "HUMAYUMPUR CHANDRAPAL", "SADATABAD", "NAVADA KESHO", "RASULAPUR MOHD KULI . AN.", "MITTHEPUR", "MANDAUNRI AN.", "TAKI SARAY", "VIRTHALA", "BAJAR KALAN", "SHERAPUR KALIYA", "JAMAPUR", "JAGUPURA", "JUMERAT KA BAJAR.", "RAGHUNATHAPUR", "NINDUDU KHAS AN.", "TARAKAULI MANDN", "GOVLI B.A", "MIJARPUR PALLA", "LATIFAPUR URF SIPAHIVALA AN.", "GURADASAPUR HIRA", "NIRULLAPUR CHIMMA", "SHAURAMAPUR", "MANDI MARKAM GANJ", "SHERAPUR RAINI", "ASAFABAD BHAWAN URF GANVDI", "CHANARAN", "KADARABAD KHURD", "NANGALA GUNGA", "BUAPUR HARAPAL AN.", "FATAHANGAR AN.", "JAMALAPUR MAHIPAT", "MUKRAPURI AN.", "ALADINPUR BHATPPURA", "DATANGAR", "KHURADA AN.", "SHARIFAPUR", "SADULLAPUR KHANAPUR", "PURNAPUR", "WAJIDAPUR", "NIKAMMASHAN AN.", "UMARAPUR ASHA AN.", "NASIRAPUR BANAVARI", "RAMATHERA", "MAUJAMPUR JAITRA AN.", "DINADARAPUR", "KHATRIYAN AN.", "HAKEEMPUR MEGHA", "MITTHEPUR", "IBRAHIMAPUR NARAYAN", "BAKHSHANAPUR AN.", "SALARABAD", "FATAHAPUR JAMAL AN0", "NANGLI LADAN AN.", "MOHD.PUR VEERU", "CHANCHAL PURAYAN", "JOTAHIMMA AN.", "SABADLAPUR AN.", "KAHARANA AN.", "HASNAPUR PALKI", "CHAK MOHD NAGAR", "BADASHAPUR SHAH MAUHM‍MAD", "ABUNASRAPUR", "SARKADA CHATRU AN.", "CHAUDHARIYAN", "AFAGANAN AN.", "MILAKIYAN AN.", "GAJARAULA B.A", "AURANGSHAHPUR BAUVARI", "AJMAPUR JAMANIMAN", "HAYATNAGAR", "FAJULLAPUR AN.", "MILAK BENIRAM", "BERAKHEDA TANDA AN.", "KHANUJAT", "NATHADOI", "RAYAPUR", "RANAVLI", "HARIYANA", "BUAPUR KHEM", "CHAKARAJMAL", "BUAPUR HARAPAL AN.", "SHERAPUR", "RAFIPUR", "HALAVAIYAN", "KHALILAPUR", "MOHD.ALIPUR BHIKAN", "BAJIRAPUR JANGIR AN.", "RAGHUNATHAPUR", "SADATNAGAR", "KAHARAN AN.", "SHEKHAPUR BHAVDA", "TAKI SARAY AN.", "KHANAPURSIRIYA", "BADASHAHAPUR LAXHMI SEN", "PATAVARIYAN AN.", "FATAHANGA AN.", "BAMANAULI", "MAIHARI SARAY AN.", "SARAKATHAL MADHO", "SALEM SARAY", "RAMPUR DULLI", "HUMAYUMPUR CHANDRAPAL", "NANGALA JAJAN", "BAKALAN", "VAJIRAPURLAL AN.", "MITTHEPUR", "ISLAMNAGAR AN.", "HINDU CHAUDHARI", "KAYASTHAN AN.", "MOHD.PUR JAMAL", "PIR KA BAJAR AN.", "BASEDA KHADAR", "MOHD.ALIPUR ABHYACHAND", "IMAMABAD AN", "SHEKHAN AN.", "NURALLAPUR UDYACHAND", "GAJUPURA", "KOTRA AN.", "NAYA GANVA MAJRA", "UMARPUR BANGAR B.A", "NANDAGANVA", "KHATRIYAN AN.", "NANGALA BHAJJA", "DINADARAPUR", "MAUJAMPUR JAITRA AN.", "RAMATHERA", "HAJIPUR", "CHANCHAL PURAYAN", "NIN‍DOODU KHAS AN.", "KASBA AN.", "DATANGAR", "SHAHAALIPUR UDYACHAND", "NANGALA BUDHVA", "KESHOPUR", "PURNAPUR", "SADULLAPUR KHANAPUR", "SHAHAJADAPUR TARU", "SHARIFAPUR", "MILAK BENIRAM", "ALAVALPUR", "HAYATNAGAR", "BUAPUR NATTHU AN.", "AJITAPUR DASI", "GAJARAULA B.A", "HUSAINPUR HAMEED", "DHAMPUR HUSAINPUR AN.", "RAYAPUR", "NATHADOI", "SUHAGAPUR AN.", "KHANUJAT", "CHAK MOHD NAGAR", "NANGLA NANSAI", "AFAGANAN AN.", "CHAUDHARIYAN", "TARAIR AN.", "ABUNASRAPUR", "SHEKHPUR THATH", "NAUDHNA AN.", "BADASHAPUR SHAH MAUHM‍MAD", "BALKISHNAPUR", "GOVLI B.A", "HAKIMAN AN.", "BURHANNAGAR", "SHEKHAPUR BHAVDA", "SADATNAGAR", "TARAI AN.", "KODUPURA", "KHANAPURSIRIYA", "CHAKARAJMAL", "BUAPUR KHEM", "RANAVLI", "RAYAPUR", "BAJIRAPUR JANGIR AN.", "HALAVAIYAN", "KHALILAPUR", "RAFIPUR", "SHERAPUR", "MOHD.ALIPUR ABHYACHAND", "IMAMABAD AN", "BASEDA KHADAR", "KHUJISTANAGAR", "NAVADA SAIDAPUR JAMAL URF BHOOTPURI", "UMARPUR BANGAR B.A", "NAYA GANVA MAJRA", "VAJIRAPURLAL AN.", "BAKALAN", "NANGALA JAJAN", "RAMPUR DULLI", "PIR KA BAJAR AN.", "FATEHAPUR LAL B.A", "ISLAMNAGAR AN.", "RASULAPUR MOHD KULI AN.", "MITTHEPUR", "SADADOBAIR AN.", "NOORPUR CHHIBRI AN.", "SHAHALIPUR NICHAL", "RASULPUR IMMA", "TEEVRI AN0", "BUNDUKACHIYAN AN.", "KAJIYAN", "BAJIRAPUR URF MADHIYA", "SALARAPUR  B.A", "AMIRAPUR", "MUSATAFAPUR TAIYAB", "BHAJDI SARAY", "RAYAPUR", "SAHUVAN", "NATHADOI", "KHANUJAT", "CHAK SHAHJANI", "SHAHAJADAPUR TARU", "MU.CHAUDHARIYAN AN.", "SADULLAPUR KHANAPUR", "KESHOPUR", "JASAMAUR", "HAJIPUR", "KAJI SARAY", "BAKAR KASSABAN", "SHAFIYABAD AN.", "THATJAT AN.", "VAJIRAPURLAL AN.", "BADAVAN AN.", "RAMPUR BHOLA", "HAJIPUR", "HUMAYUMPUR CHANDRAPAL", "JUMERAT KA BAJAR.", "JAMALAPUR ALAM", "BAJAR KALAN", "VIRTHALA", "TAKI SARAY", "MAUHADA AN.", "MOHD.PUR SADA", "GUJARATIYAN", "ACHARAJAN", "HAKIMPUR SHAKARGANJ AN.", "NANDAGANVA", "KESHOPUR", "SARKATHAL SAHANI AN.", "SADARNAPUR", "BUAPUR KHEM", "MAUJMAPUR SURAJ", "KADARABAD KHURD", "NANGALA GUNGA", "WAJIRPUR BHAGWAN", "CHANARAN", "MUKRAPURI AN.", "FATAHANGAR AN.", "ALHAIPUR AN.", "BUAPUR HARAPAL AN.", "MU.CHAUDHARIYAN", "NAYAK SARAY", "DITNAPUR AN.", "TARAKAULI MANDN", "GOVLI B.A", "RAGHUNATHAPUR", "ASAFABAD BHAWAN URF GANVDI", "SHAURAMAPUR", "MANDI MARKAM GANJ", "SHERAPUR RAINI", "GURADASAPUR HIRA", "KOPA", "PALNAPUR")
		case "Nehtaur":
			sections = append(sections, "DARBARPUR", "TAKIYA GARI PASHIMI", "MUZAFFARPUR", "FATEHPUR ASAL", "MEHAKMA UTTARI", "SALLHAPUR 1", "KHALILPUR", "PAHARPUR BAST GAIR AB.", "CHACK MUSTAFABAD GAIR AB.", "LAKHARAN", "RAISAN U. 2", "RUSTAMPUR SHER", "HAJJAMAN", "FIROZPUR UGRSAIN", "SAHANPUR NAWADA", "HAMEEDPUR", "RAISAN D. 7", "TARKOULA", "CHAPEGARAN D.", "IBRAHEEMPUR LAL", "FIROZPUR", "SARAY QAZIYAN", "SALEMPUR RUKHADIO", "BHARAKHAIDI URF DUGRU", "FAZALALIPUR", "HALWAIYAN", "MEMUDPUR MILAK", "DAWARPUR GAIR AB.AD", "SETHPUR DHANESAR", "MAHUAA", "SHAHNAGAR", "AFGANAN", "RAYPUR LAKRA", "KHEDA U. 1", "KHEDA D. 1", "IBRAHIMPUR KHANDSAL", "DUWARKAPUR", "DAKKHANA KHAS SHUMALI", "ANWARPUR ADIL GAIR ABAD", "PUTTHA", "MATOURA DURGA", "PEEPJAN", "MEHAR ALIPUR", "HASANPUR JAT", "PAHARPUR KAMALPUR", "RAYPUR MULAK", "MEMRAN", "PALI TEKCHAND", "DHELI GUJJAR", "TAHARPUR TAPPA UMRI GAIR AB.AD", "DAKSHIN SARAY", "PRATHVIPUR BANWARI", "JYOTHI", "MEHMODPUR KAMIL", "BISATH", "MOMINPUR GULARI", "KHEDA D. 2", "BHATYANA KHUSHHALPUR", "GANGU NAGLI GAR. AB.AD", "PADLA", "AKHERA", "BAIBALPUR", "HOLI", "CHOUDHRIYAN P.", "ALADEENPUR BHOUGI", "JALALPUR HASNA", "RAJPUR", "BAGWAN NAKTA", "ISSAPUR", "FARIDPUR ALAM GAIR AB.", "CHAJUPURA GANDHU", "RAISAN D. 5", "BASEDA KHURD", "NASIRPUR NAIN SINGH", "ENAYATPUR", "DHANORI", "BASAWANPUR", "SHAIKHPUR LALA", "ATHAI SHAIKH", "RUSTAMPUR WAJID", "SHURPUR GAIR AB.AD", "HARPUR", "MIRDGAN", "LAKHARAN", "NANHEDA", "PUTTHA 2", "SHEKHPURA", "MILAK GANGODA JAT AN.", "KAKRALA", "AMIRPUR GANGU", "MO.KULIPUR", "CHOUDHRIYAN P.", "BARUKI", "KAZIPURA", "TAHARPUR SAID", "DHNORA", "CHHAJUPURA SAID", "DAKOULA", "MUNEEMPUR", "RASULDARAN", "NAUDHA", "SHEESHGARAN", "GINDOURI", "NARGADI", "RAISAN D. 6", "PANCHAYTI MANDIR", "CHOHARPUR GAIR AB.", "BAGWAN CHAPEGARAN", "KHURDI", "MIMLA", "PEEPJAN", "KOTRA TAPPA KESHO", "MEHARPUR", "MAZHARPUR", "RAGHUNATHPUR", "SADKABAD", "DHINGARPUR", "PALI TEKCHAND", "UCHA", "BALDIYA", "FATTANPUR", "ARAZI GOPAL JOT", "KHANMPUR CHAK GAIR AB.AD", "IBRAHIMPUR SAHDO", "BANGADPUR KISHNA GAIR. AN.AD", "SIJOULA", "MUSSEPUR", "RUSTAMPUR", "AKBARPUR RADHE GAIR AB.", "AMHEDA", "KASAMPUR AHMAD", "MOMINPUR LAL KUWAR", "FARIDABAD", "PAHARPUR CHATAR", "SARAY ASNARA", "MOLVIYAN JUNUBI", "RUGDI SAID", "BASEDA UBAR", "NIZATPUR", "NAWAJISPUR AHMAD", "JALALPUR TURK", "KASHMIRI ABUNASARPUR", "RAGHORAMPUR", "DAKOULI", "MANDI BAZAR", "RAISAN U. 1,2", "BASEDA NARAYAN", "JOSHIYAN", "CHOUHARA", "MOLVIYAN SHUMALI", "SHAFIPUR HEERA", "NANGLA SARI", "HARGANPUR", "PEER SHAHEED KALA", "ISLAMPUR", "PANCHAYTI MANDIR UTTARI", "FAREEDNAGAR", "DAKKHANA KHAS SHUMALI", "FARIDPUR DARGO", "BASEDA KHEMCHAND", "FARIDPUR MALHU", "SHAMSABAD", "TAHARPUR ASKARAN", "PADLI MANDU", "DOULATPUR SUKHKHA", "GOPALPUR", "SIJOULI", "FARIDPUR SADHIRAN", "FULSANDA HEERA", "NASIRPUR BUZURG", "ISLAMPUR LALU", "KOKAPUR", "RAJDEV NANGLI", "FARIDPUR MAN", "CHUDYAKHEDA GAIR AB.AD", "CHOHADPUR", "AKBARPUR GARAV", "ISLAMPUR CHAMRA", "FARHASAN JUNUBI", "RAWANPUR", "KADRABAD", "GYASPUR", "ELAHABAD", "KAREEMPUR MUBARAK", "DARBAR SADAT", "RAISAN D. 8", "MANDORA VIP", "SARAY TALE WALI", "RAISAN D. 3", "BAIRAMPUR KHAJURI", "TURABNAGAR", "QAZIYAN", "BAKARNAGAR", "KASAMPUR LEKHRAJ", "ENAYATPUR", "MUNEEMPUR", "RASULDARAN", "SALAMATABAD", "JALALPUR AASRA", "BAMNOLI", "BARKAPUR GAIR AB.AD", "JASMOURA", "DAKKHANA KHAS JUNUBI", "CHAKARPUR", "MEHAR ALIPUR", "MEMRAN", "MO.PUR LAL", "GANGODA JAT", "KESHODASPUR GAIR AB.", "SHAHPUR SAIDU", "MAZHARPUR", "BAHADA GAIR AB.AD", "CHAK ABDUL REHMAN", "FAZALPUR", "ISHAKPUR", "FARIDPUR DALLU", "MUKRANDANPUR MANAK", "JYOTHI", "MUZAHIDPUR", "KASAMPUR AHMAD", "NAIKOFAL URF BILAI", "MUZAFFARABAD", "MAHU NANGLI", "FARHA SAN SHUMALI", "SARAY ASNARA", "FALOUDA", "RASULPUR VIRAN GAIR AB.", "MATOURAMAN", "SHAHBAD", "NASEEBPUR", "HATHI MANDIR", "MANDI BAZAR", "AASPUR NAWADA", "KALALI", "SHAHNAGAR", "KHETAPUR", "SHADIPUR", "EIDGAHA", "DHAKKA KARAMCHAND", "AFGANAN", "RAYPUR LAKRA", "SIKERA", "MAHAKMA JUNUBI", "IBRAHIMPUR KHANDSAL", "RANGREZAN", "SARAY JOKHA SINGH", "CHACK GOWARDHAN", "RAISAN D. 1", "MUZAFFARPUR DEVIDAS", "SHAPPAR SHIKOHPUR", "HARDASPUR GAIR AB.AD", "KANHA NANGLA", "KHALILPUR", "SAMASPUR URF MANAKPUR", "BHAUKHEDA GAIR ABAD", "KHEDA U. 3", "RAISAN D. 2", "MUKARRABPUR", "KASBA", "ALIPUR DAMODAR", "RAMPUR", "CHANDA NANGLI", "SADAT SAHADRI", "BUDIKA GAIR AB.AD", "GAZIPUR", "BILASPUR", "CHOHARPUR", "FARHASAN JUNUBI", "FARIDPUR SANSARU URF DHOKALPUR", "FAIZPUR RAMANAND", "DAKKHANA KHAS JUNUBI", "MO.HUSAINPUR GAR ABAD", "KAYSTHAN", "FATEHPUR", "MAHMUDPUR JAGMAL", "SADATNAGAR", "HAJJAMAN", "ALINAGAR PALNI", "MO.ALIPUR MUKTA", "FARIDPUR NIZAM", "BAIRAMNAGAR", "GOGALI", "JAMALPUR", "RAYPUR MALIYABAD", "MOHSANPUR KALYAN", "NURUDEENPUR", "FULSANDA HEERA", "MANGU CHARKHI", "NANGAL JAT", "PAHARPUR", "INAMPURA", "BHARAKHERI EMMA", "FARHASAN JUNUBI", "HAKEEMPUR CHANDAN", "TEERGARAN", "CHOHADPUR", "MEHMUDPUR BUZURG", "KADRABAD", "BEGRAJPUR", "FARIDPUR DULLI", "NOORPUR KHAIDKI", "MANDI BAZAR", "KHEDKI TAPPA NANGAL", "DAKOULI", "CHOUDHRIYAN PU.", "ABDUL KHALIKPUR BALRAM", "ISMAILPUR", "TAKIPURA", "KHIJARPUR", "HAKEEMAN", "RAYPUR BEGA", "HAKEEMPUR DULLA G.AB.", "MANDORI", "DAKKHANA KHAS SHUMALI", "SIKANDARPUR", "SIKRI KHURD", "TAHARPUR ASKARAN", "SHEKHAN SHUMALI", "NAWADA SAHANPUR", "FALOUDI", "PEEPJAN", "KHANDSAL", "MAZHARPUR", "RAGHUNATHPUR", "SHAHPUR SAISU", "YAKUBPUR GAIR AB.", "ARAZI GOPAL JOT", "TARKOULI EMMA", "UCHA", "PALI TEKCHAND", "RUGDI CHETRAM", "CHAPEGARAN U.", "RAZANAGAR", "CHOUDHRIYAN PU.", "SARAY RAJAB ALI", "KINAN URF KALYAN", "GOVINDPUR GAIR AB.AD", "PAPSARA", "AMHEDA", "SAHANPUR", "KASAMPUR AHMAD", "AMBARPUR ZAFAR", "GULI TALAB", "SARAY ASNARA", "MOLVIYAN JUNUBI", "SADRUDDINNAGAR", "KHOSPURA", "CHOUDHRIYAN P.", "JAMALPUR SHAIKH", "LADANPUR", "NARAYANPUR GAIR AB.AD", "MANKUAA", "RUKANPUR", "WAJIDPUR", "NARAYAN KHAIDI", "AMRABAD", "RASULDARAN", "MUNEEMPUR", "RAISAN D. 4", "SHEESHGARAN", "SHADIPUR KALA", "BALAPUR", "GINDOURI", "RAISAN D. 6", "SHAIKHPURA PITTHA", "SAIDPURA", "LATEEFPUR", "KHURDI", "ALAWALPUR UDDA", "CHOUDHRIYAN PU.", "FAZALPUR", "MARUFPUR", "DAKSHIN SARAY", "CHAK ABDUL REHMAN", "GANGODA JAT", "MO.PUR LAL", "GARAVPUR", "RAYPUR MULAK", "MEMRAN", "MEHAR ALIPUR", "LATIFULLAPUR", "BAIBALPUR", "DABKHEDI SALAR", "MAHU NANGLI", "FARHA SAN SHUMALI", "NAIKOFAL URF BILAI", "SAIDPURI GAIR AB.AD", "MUZAHIDPUR", "MUKRANDANPUR MANAK", "ISHAKPUR", "BAIRMABAD GADI", "IBRAHIMABAD", "BAKAR NANGLA", "KASAMPUR LEKHRAJ", "MATAPUR GAIR AB.AD", "BAGWAN NAKTA", "QAZIYAN", "RAJPUR", "BAKARNAGAR", "RAISAN D. 3", "ALHEDADPUR MUBARAK", "SARAY TALE WALI", "AOURANGSHAHPUR GANGDHAR", "AMINABAD", "DAKKHANA KHAS JUNUBI", "PUTTHA 1", "BAMNOLI", "NANHEDA", "LAKHARAN", "SALAMATABAD", "CHANDA NANGLI", "MUSTFABAD", "RAMPUR", "MUKARRABPUR", "SAMASPUR URF MANAKPUR", "KHEDA U. 3", "RAISAN D. 2", "TAPROULA", "HARDASPUR GAIR AB.AD", "MEHAKMA UTTARI", "MUZAFFARPUR DEVIDAS", "HAJJAMAN", "SAMASPUR HUSSAINPUR", "FATEHPUR", "FARIDPUR SANSARU URF DHOKALPUR", "FAIZPUR RAMANAND", "MOH.PUR PARMA", "CHOHARPUR", "BILASPUR", "DHELA GUJJAR", "GAZIPUR", "KHAJURA JAT", "ATAPUR URF KHATAPUR", "TAKIPURA PURAN", "FAZALALIPUR", "AASPUR NAWADA", "TAHARPUR TAPPA NANGAL", "SHAHBAD", "TARKOULA", "RAISAN D. 7", "KHAIRABAD", "RAISAN D. 1", "CHACK GOWARDHAN", "SARAY JOKHA SINGH", "RANGREZAN", "RAYPUR LAKRA", "AFGANAN", "SIKERA", "EIDGAHA", "HAKEEMAN", "ILAYACHIPUR KHADAGU", "SHADIPUR", "DABTHALA", "KHETAPUR", "KALALI", "SHAHNAGAR", "SIKRI BUZURG", "ISMAILPUR", "SIHORA GIRDHAR", "DAKOULI", "MANDI BAZAR", "RAGHORAMPUR", "SIKANDARPUR BUDH SINGH ABAD GAIR AB.", "NOORPUR KHAIDKI", "NAWAJISPUR AHMAD", "SHEKHAN SHUMALI", "RANGREZAN", "SHAHKARAMPUR GILARA", "MANDORI", "HARBALLABHPUR GAIR AB.AD", "RAYPUR BEGA", "HAKEEMAN", "KHIJARPUR", "TAKIPURA", "AMHEDI GAIR AB.", "MOLVIYAN SHUMALI", "NURUDEENPUR", "RAYPUR MALIYABAD", "KIRARKHEDI", "MOHSANPUR KALYAN", "JAMALPUR", "FULSANDA GANGADAS", "FARIDPUR NIZAM", "ABDULPUR MUNNA", "FARIDPUR DULLI", "RAISAN D. 8", "GUJARPUR JASPAL", "GOWARDHANPUR NANGLI", "MEHMUDPUR BUZURG", "CHOHADPUR", "TEERGARAN", "FARHASAN JUNUBI", "BHARAKHERI EMMA", "HAKEEMPUR CHANDAN", "KUNDA BAGAIN", "PAHARPUR", "INAMPURA", "BALKISHANPUR GAIR AB.", "MANGU CHARKHI", "RASULDARAN", "SURPUR AANSU GAIR AB.", "AMRABAD", "NARAYAN KHAIDI", "MUNEEMPUR", "AMIRPUR GANGU GAIR AB.", "WAJIDPUR", "TAHARPUR SAID", "MALAKPUR", "MO.KULIPUR", "JAMALPUR SHAIKH", "KHOSPURA", "PALIJAT", "ALAWALPUR UDDA", "KUMRARA", "DAKKHANA KHAS JUNUBI", "KHURDI", "SAIDPURA", "LATEEFPUR", "RAISAN D. 6", "BALAPUR", "NAUDHA", "RAZANAGAR", "CHOUDHRIYAN PU.", "CHAPEGARAN U.", "RUGDI CHETRAM", "FATTANPUR", "TARKOULI EMMA", "SHAHPUR SAISU", "MAZHARPUR", "RAGHUNATHPUR", "YAKUBPUR GAIR AB.", "MOLVIYAN JUNUBI", "SARAY ASNARA", "SULTANPUR ABAD", "AMBARPUR ZAFAR", "GULI TALAB", "FARHA SAN SHUMALI", "MAHU NANGLI", "AHMADPUR CHANDRU URF GDANA", "KASAMPUR AHMAD", "SAHANPUR", "MOMINPUR LAL KUWAR", "MUSSEPUR", "AKBARPUR DEVMAL", "SETHPUR DHANESAR", "FAZALALIPUR", "HALWAIYAN", "SHAIKHAN DAKSHIR", "SALEMPUR RUKHADIO", "SHAHBAD", "AMHEDA 2", "FIROZPUR", "SARAY QAZIYAN", "CHAPEGARAN D.", "HAMEEDPUR", "TARKOULA", "RAISAN D. 7", "PUTTHA", "MIRZAPUR", "ANWARPUR ADIL GAIR ABAD", "MURADPUR GAIR AB.", "IBRAHIMPUR KHANDSAL", "RANGREZAN", "DUWARKAPUR", "EIDGAHA", "KHEDA D. 1", "KHEDA U. 1", "MITHAN KUWAR PRATAP SINGH", "KHALILPUR", "GADAL", "MEHAKMA UTTARI", "MUZAFFARPUR DEVIDAS", "DARBARPUR", "GANGODA SHEKH", "FIROZPUR UGRSAIN", "MEMUDPUR BHIKKA", "SHERPUR BALLA", "RAISAN U. 2", "RUSTAMPUR WAJID", "ATHAI SHAIKH", "DHANORI", "SHAIKHPUR LALA", "MATAPUR GAIR AB.AD", "ENAYATPUR", "CHAJUPURA GANDHU", "BAKARNAGAR", "BAGWAN NAKTA", "RAJPUR", "IBRAHIMPUR URF KUMHARPURA", "ALADEENPUR BHOUGI", "DAKKHANA KHAS JUNUBI", "MILAK GANGODA JAT AN.", "SHEKHPURA", "UMRI", "ISLAMPUR THAMBU CHAU", "NEZE SARAY", "LAKHARAN", "NANHEDA", "MIRDGAN", "CHAK ABDUL REHMAN", "DAKSHIN SARAY", "NAKIBPUR", "SAIDA", "SEDI", "BURPUR", "RAYPUR MULAK", "PAHARPUR KAMALPUR", "HASANPUR JAT", "MILAK JAHAGIRABAD", "BAIBALPUR", "AKHERA", "HOLI", "SHERPUR", "DHUNPURA", "PADLA", "GANGU NAGLI GAR. AB.AD", "FARHA SAN SHUMALI", "MAHU NANGLI", "MUSSEPUR", "JYOTHI", "MEHMODPUR KAMIL", "IBRAHIMPUR SAHDO", "BANGADPUR KISHNA GAIR. AN.AD", "CHOUDHRIYAN PU.", "AAKU", "FATTANPUR", "BHAWANIPUR TARKOULA", "UCHA", "ARAZI GOPAL JOT", "HAKEEMPUR NARAYAN", "PALI TEKCHAND", "TARIKAMPUR", "SULTANPUR VIRAN GAIR AB.", "ZRIFPUR CHATAR", "DHINGARPUR", "SADKABAD", "KOTRA TAPPA KESHO", "PEEPJAN", "MEHARPUR", "BASEDA UBAR", "RUGDI SAID", "AMBARPUR ZAFAR", "DAKKHANA KHAS SHUMALI", "PAHARPUR CHATAR", "KANDKHEDI", "FARIDABAD", "MOMINPUR LAL KUWAR", "RUSTAMPUR", "SIJOULA", "MUSSEPUR", "SALLAHPUR 2", "DAKOULA", "NARAYAN KHAIDI", "MURSHADPUR", "TAHARPUR SAID", "DHNORA", "KHAIDIJAT", "KHEDA U.2", "KAZIPURA", "BAGWAN NAKTA", "BARUKI", "JOGIDASPUR GAIR AB.", "AMIRPUR GANGU", "CHOUDHRIYAN P.", "MO.KULIPUR", "BAGWAN CHAPEGARAN", "MIMLA", "KISHANPU BHOGAN", "PRATAP URF KHERPUR", "PANCHAYTI MANDIR", "LATEEFPUR", "NARGADI", "LAKHARAN", "GINDOURI", "PIPALSANA", "SHEESHGARAN", "KASAMPUR KRIPARAM URF SUNDARPUR", "NAUDHA", "NASIRPUR BUZURG", "SNUPTA", "ASADGAR", "NURUDEENPUR", "FULSANDA HEERA", "RAJPUR CHOMELA GAIR AB.AD", "SIJOULI", "CHOUDHRIYAN P.", "JAMALPUR", "GOPALPUR", "FARIDPUR NIZAM", "KUKDA ISLAMPUR", "DARBAR SADAT", "RAISAN D. 8", "ELAHABAD", "GYASPUR", "KADRABAD", "RAWANPUR", "ISLAMPUR CHAMRA", "BHARAKHERI EMMA", "TEERGARAN", "MANGU CHARKHI", "RAJDEV NANGLI", "KOKAPUR", "PAHARPUR", "AHIRPUR GAR.AN.", "BASEDA NARAYAN", "TAKIYA GARI PU.", "RAISAN U. 1,2", "FAZALALIPUR", "CHAK KUKDA GAIR AN.AD", "RAGHORAMPUR", "NAWAJISPUR AHMAD", "JALALPUR TURK", "SLAPUR SHAFKATPUR", "TAHARPUR ASKARAN", "FULSANDA KHAKAM", "DAKKHANA KHAS SHUMALI", "FARIDPUR MALHU", "FARIDPUR DARGO", "FAREEDNAGAR", "RAYPUR BEGA", "HAKEEMAN", "HARGANPUR", "NANGLA SARI", "ISLAMPUR", "PANCHAYTI MANDIR UTTARI", "PEER SHAHEED KALA", "BERKHERA CHOUHAN", "MOLVIYAN SHUMALI", "MILAK MUKEEMPUR", "AMHEDI GAIR AB.", "HARDASPUR GADI ZAFRABAD PRATHAM", "CHOUHARA")
		case "Bijnor":
			sections = append(sections, "JEETPURA", "ALLAUDDINPUR", "KADARPUR", "FATEHPUR CHAK", "HAKIMPUR", "NAGLI", "INCHAWALA", "SULTANPUR BANGAR", "AJIJPURA", "MIRDGAN", "AMIPUR BEGA", "JAHANABAD", "ADAMPUR", "MUSHARFABAD", "MEEROPUR AGRI", "JAHANGIRPUR LALU", "RAGHDAN", "SARIYAPUR", "ULAKPUR", "RAMPUR NOUBAD", "KARAUHIDI", "CHAUHANPURA", "PURUSHOTTAMPUR", "NIZAMPUR KHORA B A", "CHANDRBHANPUR KISHOR", "CHHAKDA", "BADSHAHPUR TARIKAM", "FIROZPUR NAROTAM", "RAMPUR THAKRA", "TAREKAMPUR PARAS", "MAHAGALI", "SAIFPUR KHADAR", "AHAMADPUR BHARTA", "BHARUKA", "DEWAPUR", "BRAHMNAN", "ALEDADPUR", "HIMMATPUR BELA", "NAI BASTI B-24", "MACHKI", "RAFIULNAGAR URF RAWALI", "JULAHAN", "SHAJADPUR", "ITAWA", "KHARI", "TAIMOORPUR", "RAMPUR BANGAR", "FIROZPUR MUBARAK URF NAYA GAOI", "KHADAK", "NIZAMPUR KHORA  A", "NAWADA", "FARIDPUR BHARTA", "ACHARJAN", "BHATAN", "FARIDPUR KAZI AN", "KHEDA", "MODHIA BHOGAN", "MOHAN PUR", "MUKEEMPUR DHARU", "JALALPUR KAZI", "GAJRAULA SHIV", "ZAFARPUR ANSHU", "SADUPURA", "HUSAINPUR SANKAT SINGH", "FARIEDPUR MIRA", "RAMJIWALA", "TAREEKAMPUR ROOPCHAND", "MOHMADPUR MANDAWALI", "JHAKDI BANGAR", "SULTANPUR TAPPA NANGAL", "ABU SAIDPUR", "MURTAJAPUR BULAKI URF PEDA", "SIMLI", "MODHIA DHANSHI", "FIROZPUR MANDU", "MADSUDANPUR BHADHAR", "MUDALA", "TEEP", "MANDRO JATT", "NAI BASTI B-15", "NASRPURAN NAJAMUDDIN", "CHANDPUR FERU", "RAJARAMPUR KHADAR", "KHANPUR DULLI", "SHADIPUR DALLA", "GOPALPUR", "ISLAMPUR NEEMDAS", "AWAS VIKAS COLONY", "RASHIDPUR GARHI", "HUKAMPUR MACHIKI", "GAUSPUR TOPARI", "NANUPURA", "PADAMPUR", "ALIPUR MAKHAN", "TATARPUR", "VASIHPUR MAN", "FARASHTOLI", "BHATAN AN", "PEDI", "MACHKI", "BHAGIN", "NOORPUR DALO", "SHAJADPUR", "GHUDIYAPUR", "FARIDPUR BHORO", "KHADAK", "GANGADASPUR", "MOHAN PUR", "SEWAARAMPUR SAKAT", "JALALPUR KAZI", "MUZAFFERPUR KESHO", "SHAWILATE", "ZAFARPUR ANSHU", "MAHAJANAN", "BHOGPUR PATTI ABDULLA", "KISHANPUR", "KUWAR BAL GOVIND", "CHANDRBHANPUR KISHOR", "KOHARPUR", "RAGHUNATHPUR", "HUSAINPUR TAPPA NANGAL", "MOHAMADPUR LAKHU", "GAJRAULA ACHPAL", "RAMPUR THAKRA", "KARIMPUR BAMNAULI", "SEKHPURA", "MAHAGALI", "RAMPUR VEERAN", "RAIPUR BARYSAL", "JAMALPUR PATHANI", "SAIFPUR KHADAR", "RAMLELA PURV", "MUKEEMPUR DHARMASI", "AURANG TARA", "ALAMPUR NILA", "AURANGPUR BIBI", "MUBARAKPUR URF GADAI", "PAMARGANG", "SOTIYAN", "JHANDAPUR", "AURANGPUR HAZI", "DWARIKAPURI", "QAZIPARA  JANUBI  B", "SADULLAPUR", "SHUKANANDPUR", "REHADWA", "MUKEEMPUR BALU", "SHIKANDRABAD", "BAZAR KALA", "KAMALPUR", "FATEPUR SABHACHAND", "BATAPURA", "RAGHAVRAMPUR  SARNGPUR", "AGRI", "TIVARI URF KADARPUR JASWANT", "BHAWANIPUR", "SIRDHANE BANGAR", "NOORALAPUR HAFIZ", "WAZIDPUR", "KALYANPUR", "HUSAINPUR SANKAT SINGH", "KHANPUR MADHO", "ISLAMPUR DEEPA", "SHAHJADPUR", "RAFATPUR FATEHABAD", "KHANPUR DULLI", "KISHANWAS", "SULTANPUR TAPPA NANGAL", "HIRDERAMPUR", "RAIPUR BERYSAL", "MIRZAPUR MAHESH", "ALIPUR NAGLA", "KHANA PATTI", "FIROZPUR TAREEKAM", "MADSUDANPUR BHADHAR", "GANGA KHEDI", "MUKARPUR KHEMA", "GANGOI KHADAR", "SHUMAL KHEDI", "KINAN URF MADI", "RAMBAG", "JAHENABAD", "RASOOLPUR PRITHI", "ISLAMPUR DAS", "MACHKI", "RAMPUR BANGAR", "BAZAR SHAMBA", "HASANPUR KAZI", "FARIDPUR BHARTA", "FATEHPUR KHURD", "MIRGIPUR", "SIDHARUWALA", "AMIPUR URF DHARMNAGRI", "MOMINPUR DARGO", "KHEDA", "CHAHSHEERI B B-24", "AFGANAN", "ZAFARPUR ANSHU", "KHWAJGIPUR", "JALALPUR KAZI", "DARANAGAR", "BHOGPUR PATTI ABDULLA", "UMARPUR", "FAZALPUR", "CHANDRBHANPUR KISHOR", "RAIPUR PIYAGI", "FARIDPUR UDDA", "CHANDANPURA", "MOHAMADPUR LAKHU", "RAMPUR BAKALI", "JATAN B-4", "MUKARANDPUR", "KHUDAHERI", "NILOHA", "SHAHPUR", "KHEDKI HEMRAJ COLLONY", "BRAHMNAN", "RAMSAHAYWALA", "ALIPUR SUKHANANDPUR URF LALPUR", "SADULLAPUR", "SHAHBAZPUR KHANA", "HAKIMPUR", "HASHAMPUR", "MUKEEMPUR JAMAL URF INAMPUR AN", "FATEHPUR KHURD", "MIRDAGAN", "MITHANPUR SHOBHARAM", "FATEPUR SABHACHAND", "NASIRI", "MUSHARFABAD", "MOLHARPUR", "JAHANABAD", "AJIJPURA", "PAPAWAR KHURD", "MUQAITPUR", "SULTANPUR BANGAR", "INCHAWALA", "RAMPUR NOUBAD", "SOTIYAN", "BAHADARPUR JATT", "MEERAPUR RAZA", "MANDAWLI SAIDU", "SADULLAPUR", "PRHTHVIPUR", "KUNDANPUR", "KHATRIYAN B-12", "KACHHPURA", "HASHAMPUR", "SHUKANANDPUR", "KAMALPUR", "JAHAGERPUR", "MITHANPUR SHOBHARAM", "RASOOLPUR", "BURHANUDDINPUR", "BATAPURA", "FATEPUR SABHACHAND", "MOHAMADPUR DEOMAL", "SULTANPUR BANGAR", "MIRDGAN B-10", "MANSHAPUR", "MUQAITPUR", "MEERAPUR KHADAR", "DAIVALGARH", "NOORALAPUR HAFIZ", "BHOGPUR PATTI ABDULLA", "FAREEDPUR BHOGI", "NAI BASTI B-23", "TITARWALA", "KARIMNAGAR URF ULEDA", "CHAHSHEERI B B B-21", "BAHADARPUR BUJURG", "KISHANPUR", "HUSAINPUR TAPPA NANGAL", "RAFILUL NAGAR URF RAWALI", "KASSABAN", "RAGHUNATHPUR", "JATAN B-2", "DHARAMPURA", "RAMPUR THAKRA", "MOHAMADPUR LAKHU", "BARKALA", "AHMADPUR SULEMAN", "ISLAMPUR LALA", "AURANGBAD SHAKURPUR", "BHADARPUR KHURD", "QAZIPARA  SHUMALI", "SAIFPUR KHADAR", "KAMBHAUR", "BAZDARAN", "FARIDPUR BHOGAN", "MUBARAKPUR URF GADAI", "JATAN B-3", "FARASHTOLI", "KANONGOYAN", "SHAJADPUR", "KSSABAN", "MUBARKPUR TALAN", "GHASIWALA COLLONY", "FARIDPUR BHARTA", "JAHAGERPUR", "SIMLA KALA", "CHAHSHEERI B-24", "FATEPUR KALA", "MAUZAMPUR SUJAN", "GANGADASPUR", "SEWAARAMPUR SAKAT", "NIJAMATPURA GANJ", "SIHORA NIZAMABAD", "SHAHWILLAYAT", "BEGAMPUR", "GAVRI BUJURG", "SHAWILATE", "BHOGPUR PATTI HARSUKK", "RAFATPUR FATEHABAD", "KHANPUR DULLI", "AMIPUR URF NARYANPUR", "MANWALA", "KALPUR", "DHARAMPURA", "MANGOLPURA", "GAUSPUR TOPARI", "NANUPURA", "PADAMPUR", "HUSAINPUR", "MAHMOODPUR", "MAQSUDANPUR HAFIZ", "KHALIULLAPUR", "KABULPUR", "NIZAMPUR KHORA  A", "SIDHARUWALA", "MIRGIPUR", "FATEHPUR KHURD", "FARIDPUR BHARTA", "QAZIPARA JANUBI", "BAZAR SHAMBA", "KHATRIYAN B-13", "RAMPUR BANGAR", "AMIPUR SUDHA", "RASOOLPUR PRITHI", "SHADIPUR", "BHOGPUR PATTI HARSUKK", "KHWAJGIPUR", "PAHADPUR CHANDRASAIN", "KHEDA", "RATANPUR RIYAYA", "HASHAMPUR", "ABU SAIDPUR", "MUKARPUR GUJAR", "HIRDERAMPUR", "RAIPUR BERYSAL", "RASULPUR PITTANKA", "MIRZAPUR BANGAR", "KALYANPUR", "RAMJIWALA", "HUSAINPUR SANKAT SINGH", "KINAN URF MADI", "RAMBAG", "GANGOI KHADAR", "HUSAINPUR", "AKBARPUR DEVIDASWALA", "MADSUDANPUR BHADHAR", "FIROZPUR TAREEKAM", "SHAHPUR LAL", "KHANA PATTI", "HASHAMPUR", "SHAHBAZPUR KHANA", "ALIPUR SUKHANANDPUR URF LALPUR", "RAMSAHAYWALA", "BADSHAHPUR", "RAMPUR NOUBAD", "PAPAWAR KHURD", "AJIJPURA", "INCHAWALA", "SULTANPUR BANGAR", "NASIRI", "JAHANABAD", "KANUNGOYAN", "MIRDAGAN", "JAHAGERPUR", "QAZIPARA JANUBI B", "AHMADPUR SULEMAN", "DHARAMPURA", "CHANDANPURA", "BADSHAHPUR TARIKAM", "SADAKPUR URF BILASPUR", "FARIDPUR UDDA", "RAIPUR PIYAGI", "NIZAMPUR KHORA B A", "PURUSHOTTAMPUR", "FAZALPUR", "KARAUHIDI", "SAIFPUR BANGAR", "UMARPUR", "DARANAGAR", "DEWAPUR", "SHAHPUR", "KHEDKI HEMRAJ COLLONY", "NILOHA", "JATAN B-4", "MUKARANDPUR", "CHOUDHARYAN", "KHAKROWAN", "AHMADPUR SULEMAN", "BARKALA", "MOHAMADPUR LAKHU", "DHARAMPURA", "RAGHUNATHPUR", "JAHANGIRPUR", "HUSAINPUR TAPPA NANGAL", "RAFILUL NAGAR URF RAWALI", "NAWALPUR", "SHABAJPUR", "KISHANPUR", "BAHADARPUR BUJURG", "BHOGPUR PATTI ABDULLA", "MUBARAKPUR URF GADAI", "JATAN B-3", "BAZDARAN", "ALAMPUR NILA", "KAZIYAN", "BHADARPUR KHURD", "SEKHPURA", "AURANGBAD SHAKURPUR", "RAIPUR BARYSAL", "MUNDALA", "SHUKANANDPUR", "SADULLAPUR", "AKBARPUR ANGAKHEDI", "MANDAWLI SAIDU", "BAHADARPUR JATT", "MEERAPUR RAZA", "NOORALAPUR HAFIZ", "DAIVALGARH", "MEERAPUR KHADAR", "TIVARI URF KADARPUR JASWANT", "MANSHAPUR", "SALMABAD", "MOHIUDDINPUR", "MUQAITPUR", "BHAWANIPUR", "RAGHAVRAMPUR  SARNGPUR", "FATEPUR SABHACHAND", "SWAHEDI KHURD", "BURHANUDDINPUR", "JAHAGERPUR", "KALPUR", "MUKARPUR GUJAR", "MANWALA", "JAMALPUR", "SHAHWILLYAT", "KHANPUR DULLI", "KHALIULLAPUR", "MAHMOODPUR", "CIVIL LINE  SECOND", "ALIPUR MAKHAN", "HUSAINPUR", "HUKAMPUR MACHIKI", "FATEHPUR KHURD", "GHASIWALA COLLONY", "KUTUBPUR GARHI", "MUBARKPUR TALAN", "MUQAITPUR", "KANONGOYAN", "CHAHSHEERI B B-21", "CIVIL LINE FIRST", "SHAWILATE", "BHOGPUR PATTI HARSUKK", "SIHORA NIZAMABAD", "BEGAMPUR", "PAHADPUR CHANDRASAIN", "SEWAARAMPUR SAKAT", "NIJAMATPURA GANJ", "GANGADASPUR", "BARKATPUR", "MAUZAMPUR SUJAN", "FATEPUR KALA", "SIMLA KALA", "FIROZPUR NAROTAM", "BADSHAHPUR TARIKAM", "CHAHSHEERI B B-22", "BAZAR MANGAL", "CHANDRBHANPUR KISHOR", "NIZAMPUR KHORA B A", "MAHESHWARI", "DAULATPUR", "KARAUHIDI", "CHAUHANPURA", "PURUSHOTTAMPUR", "BRAHMNAN", "ALEDADPUR", "DEWAPUR", "ALAMPUR NILA", "SHIKANDARPUR LALMAN", "PERZADAGAN", "AHAMADPUR BHARTA", "SEKHPURA", "TAREKAMPUR PARAS", "HAKIMPUR", "FATEHPUR CHAK", "DHOMANPUR", "KADARPUR", "SHUKKHAPUR MARKANDEYAMPURAM", "TAIIBPUR QAZI", "MOHAN PUR", "KHANJAHANPUR BHADAR", "RANIPUR", "SALEMPUR MATHNA URF  PURANPUR", "ALIPUR SUKHANANDPUR URF LALPUR", "JALALPUR BUCHA", "ULAKPUR", "JAHANGIRPUR LALU", "RAGHDAN", "SARIYAPUR", "ADAMPUR", "MUSHARFABAD", "TIVARI URF KADARPUR JASWANT", "BARKATPUR GARHI", "GAJIPUR", "ABU SAIDPUR", "SIMLI", "SULTANPUR TAPPA NANGAL", "BHARAIRA", "PATWARIYAN", "RAMJIWALA", "FARIEDPUR MIRA", "KALYANPUR", "BANJARAN", "LADAPURA", "FATEHPUR NOUABAD", "NAI BASTI B-15", "MANDRO JATT", "JAISINGH PUR", "KADAPUR", "GAIBLIPUR", "MUDALA", "JHALRI", "HUKAMPUR MACHIKI", "CHANDPUR NAUBAD", "FIROZPUR TAREEKAM", "QAZIWALA", "MIRGIPUR", "NIZAMPUR KHORA  A", "CHAUKPURI", "KHARI", "FARSHTOLI", "BAZAR SHAMBA", "ITAWA", "MACHKI", "HIMMATPUR BELA", "KHWAJGIPUR", "GAJRAULA SHIV", "ZAFARPUR ANSHU", "JALALPUR KAZI", "MOHAN PUR", "SWAHEDI BUJURG", "ACHARJAN", "NIZAMPUR KHORA  A", "FARIDPUR BHORO", "NOORPUR DALO", "SHAJADPUR", "BHAGIN", "BHATAN AN", "FARASHTOLI", "MAHAJANAN", "BHOGPUR PATTI HARSUKK", "MUZAFFERPUR KESHO", "SIHORA NIZAMABAD", "FARIDPUR QAZI AN", "YUSUFPUR HAMID", "GAJIPUR", "MAHAGALI", "RASHIDPUR GARHI", "ABU SAIDPUR", "GOPALPUR", "AHMADPUR  MOHIUDDINPUR", "SHAHBAJPUR", "GOGALPUR KASAM", "TATARPUR", "HUSAINPUR", "ALIPUR MAKHAN", "GAUSPUR TOPARI", "NANUPURA", "DEDANAGALA NAGALA", "HUKAMPUR MACHIKI", "REHADWA", "DHARAMNAGARI COLONY", "MANDAWLI SAIDU", "DWARIKAPURI", "JHANDAPUR", "SOTIYAN", "MADSUDANPUR NAND URF JHALRA", "JALALPUR BUCHA", "DAIVALGARH", "WAZIDPUR", "MIRDGAN B-11", "BHAWANIPUR", "TIVARI URF KADARPUR JASWANT", "BURHANUDDINPUR", "RAGHAVRAMPUR  SARNGPUR", "BATAPURA", "MOLVIYAN", "SHIKANDRABAD", "RAMPUR THAKRA", "ABUL KHAIRPUR BANGAR", "AHMADPUR SULEMAN", "BHOJPUR BHOPATPUR", "BARKALA", "RAFILUL NAGAR URF RAWALI", "KOHARPUR", "CHHACHHARI TEEP", "KUWAR BAL GOVIND", "MOHSANPUR CHAMRA", "SADAT", "NIZAMPUR KHORA B A", "PAMARGANG", "AURANGPUR BIBI", "ALAMPUR NILA", "BAZDARAN", "AURANG TARA", "SHIKANDARPUR LALMAN", "SAIFPUR KHADAR", "DAYALWALA", "RAMPUR VEERAN", "MAHAGALI", "RAIPUR BARYSAL", "JHAL", "BASIRUHASI", "SEKHPURA")
        case "Chandpur":
			sections = append(sections, "KAFOORPUR", "NAWADA", "AZAMPUR URF ADDOPUR A", "SOHARABNAGAR", "KARANPUR", "FAZIPUR KADER", "AKBARPUR TIGRI  A.", "KHANPUR RAYDASS G.A", "DHOLANPUR", "ORANGABAD", "BADIWALA", "PEPALSANA", "KAMNA UMRA G.A", "FAJALPUR POUTA", "BABANPURA", "MURAHAT A.", "KUHUTUBPUR GAWRI", "CHAHSANG A.", "RAMPUR KHADER", "BEERAMPUR", "MEERAPUR SEEKRI", "YUSUFA", "HALLANAGLA", "CHAKRUSTAMPUR", "HARPUR", "SINGH A.", "BEBIPURA", "LATEEFPUR CHUKHRI", "LATRA", "MUBARAEKPUR POTA", "KAJISARAY A.", "RASULPUR PUNA", "MADAMAHDUD", "JAINULAUDDINPUR", "AKBARPUR URF CHOTAPUR", "KUNDA", "BASANTPUR", "SINGHA", "BOHRA", "CHAKORANGABAD", "PEETANONDHA", "BARKHADA", "SIAU A.", "HAZARPUR KASWA", "RATANPUR KUHRD", "RAJPUR PARSU", "PAHADPUR KHURD", "PRATHBIPUR", "PADLA", "MAHMOODPUR KASWA", "AKONDA A.", "BHAGOORA", "TIRPUDI", "ALIPURA", "MERJAPUR BALA A.", "HESAMPUR", "LAKMIDHAR", "DARWAR A.", "PATEYPARA A.", "RUSTAMPUR", "DHARUPUR A.", "SHARAY", "BHOOGPUR", "AKBARPUR MAJRA SAINDVAR", "MATLABPUR", "ASALATPUR URF TOFAPUR", "MUFTISARAY A", "JAFARABAD KURAI", "MANZOLA GURJJAR A.", "AJDEV", "MUBARAEKPUR POTA", "CHEEMIN A.", "THARPUR MADAD IMAM", "KAKRALA", "DHARMUPOOTA", "KUNDA", "BABARPUR", "KAYASTAN A.", "AKBARPUR", "ALABALPUR", "FOLADPUR", "MUSTAFABAD", "KULCHANA", "PAHULI", "DHOLANPUR", "MANPUR", "SIRAKTHAL", "MITTANPUR", "KHARPUR JAGEER", "CHAMROOLA", "SHAHUVAN A", "MOHD. HUSANPUR MAJRA ORANGABAD", "LINDERPUR A.", "HUSANPUR KALA", "BAGARPUR", "MEERAPUR SEEKRI", "SAINDWAR A", "HELALPUR", "BHAVANIPUR GANGRAIN", "PAVTI A.", "TAKANPURI", "COT", "SHAPUR BASODI A.", "BASTA A.", "REHRA A", "KAYSTAN A.", "FAZIPUR KASWA", "DHEVERPURA A.", "SOHARABNAGAR", "SABDALPUR KHURDH", "FATHPUR", "FAZIPUR KADER", "ARAGI BHASA", "SAKTALPUR MILAK", "RAMPUR PUNA", "TAKHATPUR", "SARAY SHAK HABEB", "JOGIONDA", "MUBARAEKPUR A", "RUKANPUR A.", "SAIDPUR", "ALIPURA", "AJUPURA JAT", "REHRA A.", "JAMALLUDDINPUR", "ISMAYILPUR A", "AHROLA", "MUZAFFARPUR KHADER", "SAHACHANDAN A.", "MERJAPUR KHADAR", "PIPLIJAT  A", "SALHAPUR", "ZUZHAILA A.", "BAGAMPUR", "PHEENA A.", "SHAPUR BASODHI A.", "BAAJAR", "CHAKORANGABAD", "ORANGABAD", "KAMNA UMRA G.A", "BAHMANSORA", "ISMAYILPUR", "DUARIKAPURI", "FAJALPUR POUTA", "NALPURA", "OLIYAPUR", "ABIDPUR", "BHAWANIPUR GADDO", "MUKARPUR", "SANSARPUR", "EBRAHIMPUR", "BALIPUR GAWRI", "MUBARAEKPUR A.", "MAHU A.", "BUNDRA KHURD", "MEERAPUR", "SUDNIPUR", "MEHMUDA KHADER", "KAMALPUR A.", "SIHALI", "KHANPUR MAJRA GANSURPUR", "RASOLPUR NANGLA A.", "SOHARABNAGAR", "DHEVERPURA A.", "SOOTKHADI", "RAVANA HABAT", "PAHULI", "TRAPUDA", "MASEET A.", "ORANGABAD", "NAWABPURA URF SHAKUPURA", "KHARPUR JAGEER", "CHAUNDHRI", "KHANDSHAL", "HUSANPUR KHASA PURV", "KUTQUIE", "HATAMPUR", "KALANPUR", "SABDALPUR A.", "TABEBPUR", "GOVLI A.", "THAT", "SIKANDER NAGLA", "NAJARPUR", "JALALPUR KHADAR", "SHAKHPURI MEENA", "KARAMPUR URF GAERI", "BASODA", "BEBIPURA A.", "KAMALPUR", "BHIDEYKHERA", "SICKRONDHA", "MUBARAEKPUR POTA", "MEERZAPUR BAKENA", "KUNDA", "DHARMUPOOTA", "SEEMLI", "TUNGRI", "MALAKPUR", "BASADI", "JATOOLA", "RUKANPUR", "BASADA", "NANDNOR", "JHUJHAPURA URF  NAIPURA", "TANDA", "HASAMPUR", "IMALIA", "RUSTAMPUR", "DARWAR", "BALIYANAGLI", "DHARUPUR A.", "KATARMAL A.", "RAVANA DARGU", "FATHULLAHPUR POTTA", "JALALPUR A.", "KARAL", "SALAMPUR A.", "MOHANPUR", "MANZOLA GURJJAR A.", "BAHTOLA", "BHAISA", "BAGAMPUR", "MUSTAFABAD URF GHDANPURA", "MERJAPUR KHADAR", "KIRATPUR", "SALHAPUR", "RAVANA", "CHAKRUSTAMPUR", "PIPLIJAT", "MUZAFFARPUR KHADER", "CHAKORANGABAD", "RAMPUR NAJRANA", "SHAPUR BASODHI A.", "AKBARPUR URF CHOTAPUR", "SAIDPUR", "SHAREEFPUR", "JOGIONDA", "RUKANPUR A.", "MUBARAEKPUR A", "MARUFPUR", "MUBARAEKPUR KALA", "MOHANPUR", "KHANPUR KHADER A.", "MAHU", "REHRA A.", "ALIPURA", "BALIPUR GAWRI", "MUBARAEKPUR A.", "KAZEESORA", "EBRAHIMPUR", "SANSARPUR", "SHADIPUR", "SHRAIRAFI  A.", "SEMLA", "KARANPUR", "RAIPUR KHADER", "MAHU A.", "BAMNOLA", "SULTANPUR KHADER", "FAJALPUR POUTA", "DARWARA A.", "ORANGABAD", "CHAHSANG", "BHAWANIPUR GADDO", "ATHI", "JAHGEERPUR M.GAYNPUR URF ROOPPUR", "BHARKARA URF AZZAMNAGAR", "SHRAYRAFI  A.", "OLIYAPUR", "AZAMPUR URF ADDOPUR A.", "AKONDA", "HUSANPUR KHASA PURV", "JOINPUR", "KHANDSHAL", "MAHMOODPUR", "RAVTI", "MASEET A.", "AKLASPUR", "GHANSURPUR", "BHAWANIPUR GADDO", "SUNGHAR", "SEDPURA URF NAIPURA", "OLIYAPUR", "JAFARPUR COT", "KALANPUR", "MO.PUR KHADER", "MOHD. HUSANPUR MAJRA ORANGABAD", "SISONA A.", "SAINDWAR", "SIHALI", "HIRNAKHERI", "COT", "MEHMUDA KHADER", "SHADIPUR HUSAINPUR", "AJIMULLHANAGAR", "ROONIYA", "ISMAILPUR", "RAVANA HABAT", "BAHADARPUR", "DHEVERPURA A.", "RASOLPUR NANGLA A.", "FATHPUR", "PATYAPARA A.", "KHANPUR MAJRA GANSURPUR", "UMRIPEER A.", "BAMNOOLI", "HUSANPUR KHASA P.", "IMALIA", "BALIYANAGLI", "HASAMPUR", "TANDA", "JHUJHAPURA URF  NAIPURA", "KOSHALY", "GANDHOR A.", "NASARPUR  MANSURPUR", "BAHTOLA", "SALAMPUR A.", "MOHANPUR", "FATHULLAHPUR POTTA", "CHIMMAN  A.", "RAVANA DARGU", "AKBARPUR MAJRA SAINDVAR", "THURALA", "SICKRONDHA", "BHIDEYKHERA", "SABDALPUR A", "BASODA", "HEEMPUR BUJURG A.", "KARAMPUR URF GAERI", "SHEKPURI CHOHARD", "NAJARPUR", "LANGPURI", "MANOTA", "JATOOLA", "WAZEEDPUR", "TUNGRI", "MALAKPUR", "MAHBULAPURDHAKI", "SEEMLI", "MEERZAPUR BAKENA", "SHUJATPUR KHADER", "PEPALSANA", "BADIWALA", "KHADA G.B", "DHOLANPUR", "SARAY SIAU", "BHAWANIPUR GADDO", "MEERAPUR SEEKRI", "BEERAMPUR", "OLIYAPUR", "SANIPURA", "RAMPUR KHADER", "BABANPURA", "SHAHPUR DHAMDI", "JAGANPUR URF JAGANBALA", "KAFOORPUR", "P.V. PAHADPURKALA URF MALASJYA", "KARANPUR", "FATHPUR", "AJJUNAGLI", "BUCHNAGAL", "BHAGOORA", "ORAINGPUR JOGIPURA URF NORANGPUR", "MAHMOODPUR KASWA", "PAHADPUR KHURD", "MILAK SAKTALPUR", "LADUPURA", "KOTLA A.", "UMRIPEER", "DARWAR A.", "HESAMPUR", "GANGU NAGLA A.", "TIRPUDI", "RASULPUR PUNA", "AZAMGAR URF RATANGAR A.", "BAGAMPUR", "LATRA", "DHUNDLI", "HEEMPUR DEEPA", "CHAKRUSTAMPUR", "TAHARPUR GULAM EMAMEN", "PEETANONDHA", "DHAKI", "BOHRA", "BASANTPUR", "JAINULAUDDINPUR", "AKBARPUR URF CHOTAPUR", "MADAMAHDUD", "CHAHLI", "MANZOLA GUJJRA A.", "MUBARAEKPUR POTA", "NASEERPUR SHEK", "KARAMPUR URF GAERI", "MERJAPUR BALA", "RAOPUR BAHAMAN A.", "FOLADPUR", "SUNDRA", "KAYASTAN A.", "BABARPUR", "AKBARPUR", "KUNDA", "SEEMLI", "DHARMUPOOTA", "BHOOGPUR", "SHARAY", "TANDA", "RUSTAMPUR", "PILLANA", "CHAMANPURA", "MANZOLA GURJJAR A.", "AKBARPUR  ZOZA", "DATTYANA A.", "MOHANPUR", "FATHULLAHPUR POTTA", "AKBARPUR MAJRA SAINDVAR", "MATLABPUR", "MUDHAL", "SHAPUR BASODI A.", "KAMALUDDINPUR HUSANPUR", "KAZIJADGAAN A.", "COT", "TAKANPURI", "PAVTI A.", "BHAVANIPUR GANGRAIN", "JAFARHUSENPUR", "BAHADARPUR", "RAMPUR PUNA", "KHERKI A.", "FAZIPUR KADER", "SOHARABNAGAR", "SABDALPUR KHURDH", "FATHPUR", "FAZIPUR KASWA", "RAMOROOPPUR", "CHAMROOLA", "LUHARPURA", "KHANDSHAL", "KHARPUR JAGEER", "MITTANPUR", "TALIBPUR", "PAHULI", "RAOPUR BAHAMAN A", "TANGROLA", "BIRAL", "BAGWANTPUR", "BUNDRA KALA", "SADABAD", "RAMPUR KHADER", "HUSANPUR KALA", "MOHD. HUSANPUR MAJRA ORANGABAD")
        case "Noorpur":
			sections = append(sections, "ATHAI AHIR", "HIRANPURA", "FAIZPUR", "ALADINPUR A.", "MEERAPUR", "UMRI BUJURG", "PAIJANIYA A. H.N. 81 TO END", "BASANTPUR MAFI", "RASOOLPUR DHAKI", "RANANAGLA", "HASUPURA HARKISHANPUR", "MADHUSUDANPUR URF JALALPUR", "ALAUDDINPUR DHANRAJ", "SHOLAN SHEKH", "KASMABAD", "ALAMPUR", "RAEPUR BIJJU", "HAKEEMPURA", "POTA", "JAINABAD", "KOLASAGAR", "SULTANPUR MEV", "LODIPUR MILAK", "CHODHRIYAN", "DARIYAPUR,", "AMINABAD", "IBRAHIMABAD", "BIRBALPUR", "KASAMPUR VIRU KHALSA", "NAZARPUR BILLOCH", "MUNDA KHEDI", "SAHADPUR GULAL", "TENDERA", "CHAKJAMAL G.A.", "RASULPUR G.A.", "MUJAHIPUR", "DOLTABAD", "ALIPURA", "RAVIDASNAGAR", "UMARPUR", "SHAIKHAN DAKSHIRAIN", "CHELLAPUR", "TANGROLI", "BHAWANIPUR KALIYA", "MAHUPURA", "SADHARANPUR", "KASAMPUR DHANNA PUR MAFI", "SELPURA", "GOSPUR KHALSA", "BEDA,", "CHOUDHRIYAN", "RUPPUR", "HASUPURI G.A.", "SAKLI", "LUFTIPUR,", "BUDPUR NAWADA", "AFGANAN PURVI A.", "MAHDUDNASHO", "SULTANPUR MEV", "SULTANPUR", "ALIYARPUR", "MALAKPUR", "JALALPUR DIPA", "JAIRAMPUR", "MAKANPUR", "NOGRA", "RAHU NANGLI A.", "MADHUSUDANPUR URF JALALPUR", "BISHANPURA URF BIDRA A0", "TAJPUR A. H.S. NO. 797 TAK.", "SHIWALA A.M.N.125 TO END", "ALAUDDINPUR DHANRAJ", "RAWANA SHIKARPUR A.", "ALAMPUR", "PAINDAPUR", "PIPLA JAGIR A.", "MAHAMDABAD A.", "SANTO NANGLI A.", "KIWAD A.", "BRAHAMPURI", "NAMDARPUR CHANDAN", "TAJPUR A.", "KHASPURA A.", "DHINGERPUR", "MORNA A.", "BERKHEDA", "MEWLA MAFI", "NANGLI JAJU", "IDALPUR URF JOGI PURA", "RASHEEPUR", "MALWA", "NYAMTABAD", "RAMPUR KISHNA", "ABDULLANAGAR URF DANDA", "AKBARABAD G.A.", "MAIHAWATPUR HEMRAJ", "BASANTPUR MAFI G.A.", "MUBARAKPUR MADHU URF GADI", "KUNDA KHURD", "TAJPUR A. H.S. NO. 1 TO 796", "GOHAWER HALLU A0", "MAJHOLI", "BHUDHPUR,", "AFJALPUR BALDANI", "SARAEKALA", "SHAHPUR BADWA", "POTA", "MUBARAKPUR NAWADA", "SULTANPUR MEV", "SALEMPUR JHILLA", "ANESA NANGLI", "ADAMPUR KAKWADA", "DARIYAPUR,", "SULTANPUR", "BALLA NANGLA", "DHAILA AHIR", "AFGANAN PASHCHIM A.", "GUNIYAKHEDI", "DHAMROLA", "MAJHOLA BILLOCH A0", "RAMPUR VIDAR", "FAJALPUR DHAKI A. H.N. 1 TO 114", "SATTO NANGLI A.", "JATT NANGLA", "FATEHPUR BULANDI", "MADHUSUDANPUR URF JALALPUR", "ASGARIPUR A0 H.S. NO. 1 TO 133", "SATVAI", "SUMALKHEDI", "AASRA KHERA", "SHOLAN SHEKH", "BANJARAN", "RAMNAGAR SHAHIDNAGAR", "DOSTPUR GAIR A.", "HAIGERPUR BHATT", "PITTHUPURA", "AMANATPUR", "HASUPURA HARKISHANPUR A. H.N. 244 TO END", "NAYAK NANGLA", "PURENI DURWESHPUR", "MUKARPUR AHIR.", "SADAFAL A.", "MORLA", "LAKHIPURA", "BAGWADA", "BASANTPUR MAFI", "IBRAHEEMPUR,", "KHUKUNANGLA", "CHANGIPUR", "MURTJANAGAR URF GANGADHAR", "KAJAMPUR", "KHEDI", "NOORPUR DEHAT", "MAHAMDABAD", "BAHLOLPUR", "SHAIKHAN UTTARI", "MAKANPUR", "ROSHANPUR", "BASEDA KOTT", "SHAIKHAN UTTARI AN.", "CHAJJUPURA", "SEH", "JAMALPUR JATT", "SHIWALA A.M.N. 1 TO 124", "ALAUDDINPUR DHANRAJ", "GOPALPUR", "GAJIPUR", "MITHAI", "RAMKHEDA", "SHOLAN SHEKH", "ASGARIPUR A0", "GOHAWER JAIT A0", "ISLAMNAGAR A.M.N. 1", "ALAMPUR", "CHEHLA A.", "JUJHELA", "HASUPURA HARKISHANPUR A. H.N. 61 TO 243", "POTA", "SIDIYAWALI", "MADHUPURA", "HUSENPUR KALA", "SURANANGLA", "SARKADI", "BHAGWANPUR RAINI", "HEMPUR", "BER", "KAZIPURA", "KANHEDI", "JAMALPUR KIRAT", "BAMNAULA,", "HASANPUR BILLOCH", "SELA KHEDI", "MAGHPUR UDAYCHAND", "AULIYAPUR", "MUSSEPUR", "JAMALPUR KIRAT A.", "AFGANAN UTTRI A.", "MAHUPURA", "HAKUMATPUR", "SHERPUR MAJRA", "JAMALUDDINPUR G.A.", "MIRJAPUR DHIKLI,", "JALALPUR URF SHAIKPURA", "SHARFUDDINPUR KAIMRY", "DAULATPUR BILLOCH", "TAHAERPUR", "REHTI JAGIR", "KAMALA", "SADANPUR", "SALEMPUR JHILLA", "DERABULANDI", "MUBARAKPUR NAWADA", "JAINABAD", "PANDIYA", "SHAHPUR BADWA", "KHANPUR", "KURI BANGER", "RAMPUR VIDAR", "ASADPUR DHAMROLI", "PAIJANIYA A.", "AFGANAN PASHCHIM A.", "ADAMPUR KAKWADA", "CHACK MEHDUD SANI", "MUBARAKPUR MADHU URF GADI", "KUNDA KHURD", "BISHANPURA", "BASANTPUR MAFI G.A.", "MAIHAWATPUR HEMRAJ", "AKBARABAD G.A.", "ASGARIPUR A0 H.S. NO.134 TO END", "YAKUBPUR BILLOCH", "BHUDHPUR,", "KHAJURI", "MOHD. ALIPUR INAYAT A.", "ALIPURA", "BHOGPUR", "CHAKJAMAL G.A.", "MAJHOLI", "AMANATPUR", "MURTJANAGAR URF GANGADHAR", "PATHWARI", "CHAKRAMAIYA G.A.", "KASHMIRPUR GADI", "BASANTPUR MAFI", "GOVINDNAGAR", "SHEKHAN UTTRI A.", "DHAILY AHIR", "MORLA", "KASAMPUR MANGAL KHEDA", "RATANPURA", "DAWAL KHEDI", "MUKARPUR AHIR.", "SHIWALA A.", "MEERAPUR", "HASUPURA HARKISHANPUR A. H.N. 221 TO END", "PURENI DURWESHPUR", "MEWA JATT A.", "JATT NANGLA", "FATEHPUR BULANDI", "RAHPANPUR", "HAJRATNAGAR A.", "HASUPURA HARKISHANPUR A. H.N. 1 TO 220", "SHOLAN SHEKH", "POTI", "AASRA KHERA", "JAMALPUR JATT", "DARIYAPUR A0", "SHAIKHAN UTTARI AN.", "BASEDA KOTT", "ISLAMNAGAR A.M.N. 120 TO END", "NAINU NANGAL", "ATHAI JAMRUDDIN H.N. 121 TO END", "PURENA ABDUL REHMANPUR A.", "NASIRPUR MANSUKH", "RAMKHEDA", "FAJALPUR DHAKI A. H.N. 115 TO END", "MITHAI", "KHAI KHEDA", "MUKAM", "SHAPUR HARRAE", "LAMBAKHERA A.", "JAMALPUR JATT", "KHASPURA A,.", "MORNA A.", "MAHMUDPUR A.", "MURTJANAGAR URF GANGADHAR", "DHOLAGAD", "KHAWAJA AHMADPUR JALAL URF PAITIYA", "MAHAMDABAD A.", "IBRAHEEMPUR,", "LAKHIPURA", "MOHD. TAKIPUR GHASI", "CHAKCHAND G.A.", "MAKSUDPUR A.", "BAJHRABAD URF KHATAI", "MAHAMDABAD", "KASAMPUR MANGAL KHEDA", "JAHIRABAD G.A.", "KAJAMPUR", "ALAMPUR ADWA", "JAMALUDDINPUR G.A.", "ISLAMNAGAR A.M.N. 82 TO END", "AFGANAN PASHCHMI  A.", "SHERPUR MAJRA", "HAKUMATPUR", "AKBARABAD G.A.", "AFGANAN UTTRI A.", "MUSSEPUR", "KASAMPUR DHANNA PUR MAFI", "BHAWANIPUR KALIYA", "ABIDNAGAR URF DHUDLI", "AULIYAPUR", "GANDA JUD", "NICHALPUR", "SABDALPURI G.A.", "TAHAERPUR", "DAULATPUR BILLOCH", "MORNA A0", "JALALPUR URF SHAIKPURA", "HEMPUR", "SARKADI", "UMRI BADI", "MADHUPURA", "PANDIYA", "BAGWADA A.", "ALAMPURI", "SELA KHEDI", "HASANPUR BILLOCH", "BAMNAULA,", "GOVEDHANPUR", "RAMNAGAR A.", "MUKSUDPUR A.", "KAZIPURA", "SULTANPUR", "ALIYARPUR", "PUWENA", "MADHUSUDANPUR URF JALALPUR", "RANANAGLA", "SHUJATPUR TIKAR A.", "HASUPURA HARKISHANPUR", "RASOOLPUR DHAKI", "PURANPUR SAYANA", "ISLAMNAGAR A.", "KHANPUR BILLOCH", "RAMPUR", "JHIRAN", "SHAHBAZPUR", "HIMPUR PIRTHIYA", "MORNA A.", "MUJAHIDPUR", "MAHAMDABAD A.", "SHAHPUR KHEDI", "LAKHIPURA", "KASAMPUR MANGAL KHEDA", "ALADINPUR A.", "MEERAPUR", "SHAHNAZARPUR KOTT", "MUNDA KHEDI", "MADARIPUR KAKRALA", "LAMBIYA", "AKBARABAD G.A.", "RAMNAGAR A0", "PURANPUR NANGLA", "BHAWANIPUR KALIYA", "KASAMPUR VIRU KHALSA", "HAKEEPURA DAKSHIRIN", "SHAIKHAN DAKSHIRAIN", "JAMALPUR MAN", "UMARPUR", "JABERDASTPUR G.A.", "RAVIDASNAGAR", "GOVINDPUR", "DOLTABAD", "MORMAKDUMPUR", "BUDANPUR A.", "CHAKJAMAL G.A.", "GANDHINAGAR", "SULTANPUR MEV", "HAMA NANGLI", "JAINABAD", "AKBERPUR ASHA URF HAROLI", "SHAHPUR BADWA", "HALDUA MAFI", "SHAIKHAN TARAI", "BIRBALPUR", "ISLAMNAGAR A.M.N. 1 TO 119", "IBRAHIMABAD", "CHODHRIYAN", "BURHANPUR", "JAINABAD", "MAHDUDNASHO", "AFGANAN PURVI A.", "LUFTIPUR,", "BUDPUR NAWADA", "PANDIYA", "SAKLI", "ABDULLPUR DAHANA", "REHTOLI", "JALALPUR DIPA", "PANIYALA A.", "RASULPUR KASBA A.", "FATEHABAD", "MANDORA", "ALIYARPUR", "GOSPUR KHALSA", "PITTHAPUR A.", "SHERPUR MAJRA", "SELPURA", "ALAUDDINPUR", "KASAMPUR DHANNA PUR MAFI", "KOTRA MOLVIYAN", "AFGANAN UTTARI", "MAHUPURA", "KABIRNAGAR A.", "CHELLAPUR", "TANGROLI", "BHAWANIPUR KALIYA", "SAIDPUR MAFI", "ALIPURA", "BEDA,", "CHAKJAMAL G.A.", "MORNA A.", "MURTJANAGAR URF GANGADHAR", "KHASPURA A.", "DHINGERPUR", "MOLVIYAN A.", "RAHTA BILLOCH", "NAWADA RAWANA URF TELIPURA", "MAHAMDABAD A.", "SANTO NANGLI A.", "NYAMTABAD", "CHAKCHAND G.A.", "ABDULLANAGAR URF DANDA", "RASHEEPUR", "MEWA NAWADA A.", "MAHAMDABAD", "JAHIRABAD G.A.", "SELA A.", "KASAMPUR BILLOCH", "SHAYAMABAD", "MEERAPUR", "KAJAMPUR", "ROSHANPUR JAGEER A.", "MANGOLPURA", "GOVINDNAGAR MUHAMMADNAGAR", "VIJAI NANGLA", "RAHU NANGLI A.", "MAKANPUR", "ATHAI JAMRUDDIN H.N. 1 TO 120", "JAIRAMPUR", "PAINDAPUR", "ALAMPUR", "RAWANA SHIKARPUR", "BUDA NANGLA", "JAMALPUR JATT", "ALAUDDINPUR DHANRAJ")
		}
	}
	return sections
}

func getAcs(districts []string) []string {
	var acs []string

	for _, district := range districts {
		switch district {
		case "Moradabad":
			acs = append(acs, "Moradabad Nagar", "Kanth", "Moradabad Rural", "Thakurdwara")
		case "Rampur":
			acs = append(acs, "Suar", "Chamraua", "Bilaspur", "Milak", "Rampur")
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
