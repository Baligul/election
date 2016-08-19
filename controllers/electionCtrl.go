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
	orm.RegisterModel(new(models.Voter))
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
		voters               []*models.Voter
		votersRampur         []*models.Voter
		votersMoradabad      []*models.Voter
		votersBijnor         []*models.Voter
		votersBangalore      []*models.Voter
		votersHubli          []*models.Voter
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
