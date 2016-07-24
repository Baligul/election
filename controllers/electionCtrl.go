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


*/

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	models "github.com/Baligul/election/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://ggxssikrsehequ:sQElIpN-CHqcFFNAx7mJO31Y3v@ec2-54-225-93-34.compute-1.amazonaws.com:5432/da6obv8tnlvcev")
	//orm.RegisterDataBase("default", "postgres", "user=member dbname=election sslmode=disable")
	orm.RegisterModel(new(models.Voter))
}

type ElectionController struct {
	beego.Controller
}

func (e *ElectionController) GetVoters() {
	var (
		votersCountRampur    int64
		votersCountMoradabad int64
		votersCountBangalore int64
		votersCountHubli     int64
		voters               []*models.Voter
		votersRampur         []*models.Voter
		votersMoradabad      []*models.Voter
		votersBangalore      []*models.Voter
		votersHubli          []*models.Voter
	)

	err := errors.New("Errors")
	err = nil

	limit, _ := e.GetInt("limit")
	offset, _ := e.GetInt("offset")

	if limit == 0 {
		limit = 100
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(models.Query)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := models.NewResponseStatus()
		responseStatus.Response = "error:"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for each and every district
	qsRampur := o.QueryTable("voter")
	qsMoradabad := o.QueryTable("voter")
	qsBangalore := o.QueryTable("voter")
	qsHubli := o.QueryTable("voter")

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
		condVoterId = condVoterId.Or("Voter_id__exact", voterId)
	}

	// Ac Number
	for _, acNumber := range query.AcNumber {
		condAcNumber = condAcNumber.Or("Ac_number__exact", acNumber)
	}

	// Part Number
	for _, partNumber := range query.PartNumber {
		condPartNumber = condPartNumber.Or("Part_number__exact", partNumber)
	}

	// Section Number
	for _, sectionNumber := range query.SectionNumber {
		condSectionNumber = condSectionNumber.Or("Section_number__exact", sectionNumber)
	}

	// Serial Number In Part
	for _, serialNumberInPart := range query.SerialNumberInPart {
		condSerialNumberInPart = condSerialNumberInPart.Or("Serial_number_in_part__exact", serialNumberInPart)
	}

	// Name English
	for _, nameEnglish := range query.NameEnglish {
		condNameEnglish = condNameEnglish.Or("Name_english__icontains", nameEnglish)
	}

	// Name Hindi
	for _, nameHindi := range query.NameHindi {
		condNameHindi = condNameHindi.Or("Name_hindi__icontains", nameHindi)
	}

	// Relation Name English
	for _, relationNameEnglish := range query.RelationNameEnglish {
		condRelationNameEnglish = condRelationNameEnglish.Or("Relation_name_english__icontains", relationNameEnglish)
	}

	// Relation Name Hindi
	for _, relationNameHindi := range query.RelationNameHindi {
		condRelationNameHindi = condRelationNameHindi.Or("Relation_name_hindi__icontains", relationNameHindi)
	}

	// Gender
	for _, gender := range query.Gender {
		condGender = condGender.Or("Gender__exact", gender)
	}

	// ID Card Number
	for _, idCardNumber := range query.NameHindi {
		condIdCardNumber = condIdCardNumber.Or("Id_card_number__exact", idCardNumber)
	}

	// Ac Name English
	for _, acNameEnglish := range query.AcNameEnglish {
		condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__icontains", acNameEnglish)
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__icontains", acNameHindi)
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__icontains", sectionNameEnglish)
	}

	// Section Name Hindi
	for _, sectionNameHindi := range query.SectionNameHindi {
		condSectionNameHindi = condSectionNameHindi.Or("Section_name_hindi__ilike", sectionNameHindi)
	}

	//Religion English
	for _, religionEnglish := range query.ReligionEnglish {
		condReligionEnglish = condReligionEnglish.Or("Religion_english__exact", religionEnglish)
	}

	// Religion Hindi
	for _, religionHindi := range query.ReligionHindi {
		condReligionHindi = condReligionHindi.Or("Religion_hindi__exact", religionHindi)
	}

	// Age
	for _, age := range query.Age {
		condAge = condAge.Or("Age__exact", age)
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

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
			if state == 27 {
				votersCountRampur, err = qsRampur.Limit(limit, offset).All(&votersRampur)
				votersCountMoradabad, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
			}

			if state == 12 {
				votersCountBangalore, err = qsBangalore.Limit(limit, offset).All(&votersBangalore)
				votersCountHubli, err = qsHubli.Limit(limit, offset).All(&votersHubli)
			}

		}
	}

	// District Name Hindi
	for _, districtNameHindi := range query.DistrictNameHindi {
		if districtNameHindi == "मुरादाबाद" {
			votersCountMoradabad, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}

		if districtNameHindi == "रामपुर" {
			votersCountRampur, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}
	}

	// District Name English
	for _, districtNameEnglish := range query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
		if district == 20 {
			votersCountRampur, err = qsRampur.Limit(limit, offset).All(&votersRampur)
		}

		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Limit(limit, offset).All(&votersMoradabad)
		}
	}

	voters = append(votersRampur, votersMoradabad...)
	voters = append(voters, votersBangalore...)
	voters = append(voters, votersHubli...)

	if votersCountRampur > 0 || votersCountMoradabad > 0 || votersCountBangalore > 0 || votersCountHubli > 0 {
		e.Data["json"] = voters
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
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		otherVotersCountRampur           int64
		otherVotersCountMoradabad        int64
		otherVotersCountBangalore        int64
		otherVotersCountHubli            int64
		maleVotersCountRampur            int64
		maleVotersCountMoradabad         int64
		maleVotersCountBangalore         int64
		maleVotersCountHubli             int64
		femaleVotersCountRampur          int64
		femaleVotersCountMoradabad       int64
		femaleVotersCountBangalore       int64
		femaleVotersCountHubli           int64
		muslimMaleVotersCountRampur      int64
		muslimMaleVotersCountMoradabad   int64
		muslimMaleVotersCountBangalore   int64
		muslimMaleVotersCountHubli       int64
		muslimFemaleVotersCountRampur    int64
		muslimFemaleVotersCountMoradabad int64
		muslimFemaleVotersCountBangalore int64
		muslimFemaleVotersCountHubli     int64
		otherMaleVotersCountRampur       int64
		otherMaleVotersCountMoradabad    int64
		otherMaleVotersCountBangalore    int64
		otherMaleVotersCountHubli        int64
		otherFemaleVotersCountRampur     int64
		otherFemaleVotersCountMoradabad  int64
		otherFemaleVotersCountBangalore  int64
		otherFemaleVotersCountHubli      int64
	)

	err := errors.New("Errors")
	err = nil

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

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for each and every district
	qsRampur := o.QueryTable("voter")
	qsMoradabad := o.QueryTable("voter")
	qsBangalore := o.QueryTable("voter")
	qsHubli := o.QueryTable("voter")

	cond := orm.NewCondition()
	condAcNumber := orm.NewCondition()
	condPartNumber := orm.NewCondition()
	condSectionNumber := orm.NewCondition()
	condAcNameEnglish := orm.NewCondition()
	condAcNameHindi := orm.NewCondition()
	condSectionNameEnglish := orm.NewCondition()
	condSectionNameHindi := orm.NewCondition()
	condAge := orm.NewCondition()

	// Ac Number
	for _, acNumber := range query.AcNumber {
		condAcNumber = condAcNumber.Or("Ac_number__exact", acNumber)
	}

	// Part Number
	for _, partNumber := range query.PartNumber {
		condPartNumber = condPartNumber.Or("Part_number__exact", partNumber)
	}

	// Section Number
	for _, sectionNumber := range query.SectionNumber {
		condSectionNumber = condSectionNumber.Or("Section_number__exact", sectionNumber)
	}

	// Ac Name English
	for _, acNameEnglish := range query.AcNameEnglish {
		condAcNameEnglish = condAcNameEnglish.Or("Ac_name_english__icontains", acNameEnglish)
	}

	// Ac Name Hindi
	for _, acNameHindi := range query.AcNameHindi {
		condAcNameHindi = condAcNameHindi.Or("Ac_name_hindi__icontains", acNameHindi)
	}

	// Section Name English
	for _, sectionNameEnglish := range query.SectionNameEnglish {
		condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__icontains", sectionNameEnglish)
	}

	// Section Name Hindi
	for _, sectionNameHindi := range query.SectionNameHindi {
		condSectionNameHindi = condSectionNameHindi.Or("Section_name_hindi__ilike", sectionNameHindi)
	}

	// Age
	for _, age := range query.Age {
		condAge = condAge.Or("Age__exact", age)
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

	// Get voters for each state
	if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
		for _, state := range query.StateNumber {
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
	}

	// Get voters for each district
	for _, district := range query.DistrictNumber {
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
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli

	if votersCount > 0 {
		statistic.Total = votersCount

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBangalore + muslimVotersCountHubli
		statistic.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBangalore + otherVotersCountHubli
		statistic.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBangalore + maleVotersCountHubli
		statistic.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBangalore + femaleVotersCountHubli
		statistic.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistic.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistic.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistic.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
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
		muslimVotersCountBangalore       int64
		muslimVotersCountHubli           int64
		otherVotersCountRampur           int64
		otherVotersCountMoradabad        int64
		otherVotersCountBangalore        int64
		otherVotersCountHubli            int64
		maleVotersCountRampur            int64
		maleVotersCountMoradabad         int64
		maleVotersCountBangalore         int64
		maleVotersCountHubli             int64
		femaleVotersCountRampur          int64
		femaleVotersCountMoradabad       int64
		femaleVotersCountBangalore       int64
		femaleVotersCountHubli           int64
		muslimMaleVotersCountRampur      int64
		muslimMaleVotersCountMoradabad   int64
		muslimMaleVotersCountBangalore   int64
		muslimMaleVotersCountHubli       int64
		muslimFemaleVotersCountRampur    int64
		muslimFemaleVotersCountMoradabad int64
		muslimFemaleVotersCountBangalore int64
		muslimFemaleVotersCountHubli     int64
		otherMaleVotersCountRampur       int64
		otherMaleVotersCountMoradabad    int64
		otherMaleVotersCountBangalore    int64
		otherMaleVotersCountHubli        int64
		otherFemaleVotersCountRampur     int64
		otherFemaleVotersCountMoradabad  int64
		otherFemaleVotersCountBangalore  int64
		otherFemaleVotersCountHubli      int64
	)

	err := errors.New("Errors")
	err = nil

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

	o := orm.NewOrm()
	o.Using("default")

	// Create query string for each and every district
	// Query
	qsRampur := o.QueryTable("voter")
	qsMoradabad := o.QueryTable("voter")
	qsBangalore := o.QueryTable("voter")
	qsHubli := o.QueryTable("voter")

	// Scope
	qsScopeRampur := o.QueryTable("voter")
	qsScopeMoradabad := o.QueryTable("voter")
	qsScopeBangalore := o.QueryTable("voter")
	qsScopeHubli := o.QueryTable("voter")

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
		condAcNumberQuery = condAcNumberQuery.Or("Ac_number__exact", acNumber)
	}

	// Apply filters for each query string
	// Ac Number Scope
	for _, acNumber := range queries.Scope.AcNumber {
		condAcNumberScope = condAcNumberScope.Or("Ac_number__exact", acNumber)
	}

	// Part Number Query
	for _, partNumber := range queries.Query.PartNumber {
		condPartNumberQuery = condPartNumberQuery.Or("Part_number__exact", partNumber)
	}

	// Part Number Scope
	for _, partNumber := range queries.Scope.PartNumber {
		condPartNumberScope = condPartNumberScope.Or("Part_number__exact", partNumber)
	}

	// Section Number Query
	for _, sectionNumber := range queries.Query.SectionNumber {
		condSectionNumberQuery = condSectionNumberQuery.Or("Section_number__exact", sectionNumber)
	}

	// Section Number Scope
	for _, sectionNumber := range queries.Scope.SectionNumber {
		condSectionNumberScope = condSectionNumberScope.Or("Section_number__exact", sectionNumber)
	}

	// Ac Name English Query
	for _, acNameEnglish := range queries.Query.AcNameEnglish {
		condAcNameEnglishQuery = condAcNameEnglishQuery.Or("Ac_name_english__icontains", acNameEnglish)
	}

	// Ac Name English Scope
	for _, acNameEnglish := range queries.Scope.AcNameEnglish {
		condAcNameEnglishScope = condAcNameEnglishScope.Or("Ac_name_english__icontains", acNameEnglish)
	}

	// Ac Name Hindi Query
	for _, acNameHindi := range queries.Query.AcNameHindi {
		condAcNameHindiQuery = condAcNameHindiQuery.Or("Ac_name_hindi__icontains", acNameHindi)
	}

	// Ac Name Hindi Scope
	for _, acNameHindi := range queries.Scope.AcNameHindi {
		condAcNameHindiScope = condAcNameHindiScope.Or("Ac_name_hindi__icontains", acNameHindi)
	}

	// Section Name English Query
	for _, sectionNameEnglish := range queries.Query.SectionNameEnglish {
		condSectionNameEnglishQuery = condSectionNameEnglishQuery.Or("Section_name_english__icontains", sectionNameEnglish)
	}

	// Section Name English Scope
	for _, sectionNameEnglish := range queries.Scope.SectionNameEnglish {
		condSectionNameEnglishScope = condSectionNameEnglishScope.Or("Section_name_english__icontains", sectionNameEnglish)
	}

	// Section Name Hindi Query
	for _, sectionNameHindi := range queries.Query.SectionNameHindi {
		condSectionNameHindiQuery = condSectionNameHindiQuery.Or("Section_name_hindi__ilike", sectionNameHindi)
	}

	// Section Name Hindi Scope
	for _, sectionNameHindi := range queries.Scope.SectionNameHindi {
		condSectionNameHindiScope = condSectionNameHindiScope.Or("Section_name_hindi__ilike", sectionNameHindi)
	}

	// Age Query
	for _, age := range queries.Query.Age {
		condAgeQuery = condAgeQuery.Or("Age__exact", age)
	}

	// Age Scope
	for _, age := range queries.Scope.Age {
		condAgeScope = condAgeScope.Or("Age__exact", age)
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
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
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
	}

	votersCountQuery = votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli

	if votersCountQuery > 0 {
		statistics.Query.Total = votersCountQuery

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Query.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBangalore + otherVotersCountHubli
		statistics.Query.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Query.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Query.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Query.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Query.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistics.Query.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
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
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
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
	}

	votersCountScope = votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli

	if votersCountScope > 0 {
		statistics.Scope.Total = votersCountScope

		muslimVotersCount = muslimVotersCountRampur + muslimVotersCountMoradabad + muslimVotersCountBangalore + muslimVotersCountHubli
		statistics.Scope.Muslim = muslimVotersCount

		otherVotersCount = otherVotersCountRampur + otherVotersCountMoradabad + otherVotersCountBangalore + otherVotersCountHubli
		statistics.Scope.Other = otherVotersCount

		maleVotersCount = maleVotersCountRampur + maleVotersCountMoradabad + maleVotersCountBangalore + maleVotersCountHubli
		statistics.Scope.Male = maleVotersCount

		femaleVotersCount = femaleVotersCountRampur + femaleVotersCountMoradabad + femaleVotersCountBangalore + femaleVotersCountHubli
		statistics.Scope.Female = femaleVotersCount

		muslimMaleVotersCount = muslimMaleVotersCountRampur + muslimMaleVotersCountMoradabad + muslimMaleVotersCountBangalore + muslimMaleVotersCountHubli
		statistics.Scope.MuslimMale = muslimMaleVotersCount

		muslimFemaleVotersCount = muslimFemaleVotersCountRampur + muslimFemaleVotersCountMoradabad + muslimFemaleVotersCountBangalore + muslimFemaleVotersCountHubli
		statistics.Scope.MuslimFemale = muslimFemaleVotersCount

		otherMaleVotersCount = otherMaleVotersCountRampur + otherMaleVotersCountMoradabad + otherMaleVotersCountBangalore + otherMaleVotersCountHubli
		statistics.Scope.OtherMale = otherMaleVotersCount

		otherFemaleVotersCount = otherFemaleVotersCountRampur + otherFemaleVotersCountMoradabad + otherFemaleVotersCountBangalore + otherFemaleVotersCountHubli
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
		condNameEnglishQuery = condNameEnglishQuery.Or("Name_english__icontains", nameEnglish)
	}

	// Name English Scope
	for _, nameEnglish := range queries.Scope.NameEnglish {
		condNameEnglishScope = condNameEnglishScope.Or("Name_english__icontains", nameEnglish)
	}

	// Apply filters for each query string
	// Name Hindi Query
	for _, nameHindi := range queries.Query.NameHindi {
		condNameHindiQuery = condNameHindiQuery.Or("Name_hindi__icontains", nameHindi)
	}

	// Name Hindi Scope
	for _, nameHindi := range queries.Scope.NameHindi {
		condNameHindiScope = condNameHindiScope.Or("Name_hindi__icontains", nameHindi)
	}

	// Relation Name English Query
	for _, relationNameEnglish := range queries.Query.RelationNameEnglish {
		condRelationNameEnglishQuery = condRelationNameEnglishQuery.Or("Relation_name_english__icontains", relationNameEnglish)
	}

	// Relation Name English Scope
	for _, relationNameEnglish := range queries.Scope.RelationNameEnglish {
		condRelationNameEnglishScope = condRelationNameEnglishScope.Or("Relation_name_english__icontains", relationNameEnglish)
	}

	// Relation Name Hindi Query
	for _, relationNameHindi := range queries.Query.RelationNameHindi {
		condRelationNameHindiQuery = condRelationNameHindiQuery.Or("Relation_name_hindi__icontains", relationNameHindi)
	}

	// Relation Name Hindi Scope
	for _, relationNameHindi := range queries.Scope.RelationNameHindi {
		condRelationNameHindiScope = condRelationNameHindiScope.Or("Relation_name_hindi__icontains", relationNameHindi)
	}

	// Gender Query
	for _, gender := range queries.Query.Gender {
		condGenderQuery = condGenderQuery.Or("Gender__exact", gender)
	}

	// Gender Scope
	for _, gender := range queries.Scope.Gender {
		condGenderScope = condGenderScope.Or("Gender__exact", gender)
	}

	//Religion English Query
	for _, religionEnglish := range queries.Query.ReligionEnglish {
		condReligionEnglishQuery = condReligionEnglishQuery.Or("Religion_english__exact", religionEnglish)
	}

	//Religion English Scope
	for _, religionEnglish := range queries.Scope.ReligionEnglish {
		condReligionEnglishScope = condReligionEnglishScope.Or("Religion_english__exact", religionEnglish)
	}

	// Religion Hindi Query
	for _, religionHindi := range queries.Query.ReligionHindi {
		condReligionHindiQuery = condReligionHindiQuery.Or("Religion_hindi__exact", religionHindi)
	}

	// Religion Hindi Scope
	for _, religionHindi := range queries.Scope.ReligionHindi {
		condReligionHindiScope = condReligionHindiScope.Or("Religion_hindi__exact", religionHindi)
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
	}

	// District Name English
	for _, districtNameEnglish := range queries.Query.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsMoradabad.Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsRampur.Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Query.DistrictNumber {
		if district == 20 {
			votersCountRampur, err = qsRampur.Count()
		}

		if district == 19 {
			votersCountMoradabad, err = qsMoradabad.Count()
		}
	}

	votersResult = votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli
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
	}

	// District Name English
	for _, districtNameEnglish := range queries.Scope.DistrictNameEnglish {
		if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
		}

		if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
			votersCountRampur, err = qsScopeRampur.Count()
		}
	}

	// Get voters for each district
	for _, district := range queries.Scope.DistrictNumber {
		if district == 20 {
			votersCountRampur, err = qsScopeRampur.Count()
		}

		if district == 19 {
			votersCountMoradabad, err = qsScopeMoradabad.Count()
		}
	}

	votersCount = votersCountRampur + votersCountMoradabad + votersCountBangalore + votersCountHubli
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
