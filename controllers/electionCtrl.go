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
        votersCountRampur int64
        votersCountMoradabad int64
        votersCountBangalore int64
        votersCountHubli int64       
        voters []*models.Voter
        votersRampur []*models.Voter
        votersMoradabad []*models.Voter
        votersBangalore []*models.Voter
        votersHubli []*models.Voter
    )
    
    err := errors.New("Errors")
    err = nil
    
    inputJson := e.Ctx.Input.RequestBody
    query := new(models.Query)
    
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
    
    
    // Apply filters for each query string
    // Voter Id
    for _, voterId := range query.VoterID {
        qsRampur = qsRampur.Filter("Voter_id__exact", voterId)
        qsMoradabad = qsMoradabad.Filter("Voter_id__exact", voterId)
    }
    
    // Ac Number
    for _, acNumber := range query.AcNumber {
        qsRampur = qsRampur.Filter("Ac_number__exact", acNumber)
        qsMoradabad = qsMoradabad.Filter("Ac_number__exact", acNumber)
    }
     
    // Part Number
    for _, partNumber := range query.PartNumber {
        qsRampur = qsRampur.Filter("Part_number__exact", partNumber)
        qsMoradabad = qsMoradabad.Filter("Part_number__exact", partNumber)
    }   
    
    // Section Number
    for _, sectionNumber := range query.SectionNumber {
        qsRampur = qsRampur.Filter("Section_number__exact", sectionNumber)
        qsMoradabad = qsMoradabad.Filter("Section_number__exact", sectionNumber)
    }   
    
    // Serial Number In Part
    for _, serialNumberInPart := range query.SerialNumberInPart {
        qsRampur = qsRampur.Filter("Serial_number_in_part__exact", serialNumberInPart)
        qsMoradabad = qsMoradabad.Filter("Serial_number_in_part__exact", serialNumberInPart)
    }    
    
    // Name English
    for _, nameEnglish := range query.NameEnglish {
        qsRampur = qsRampur.Filter("Name_english__icontains", nameEnglish)
        qsMoradabad = qsMoradabad.Filter("Name_english__icontains", nameEnglish)
    }    
    
    // Name Hindi
    for _, nameHindi := range query.NameHindi {
        qsRampur = qsRampur.Filter("Name_hindi__icontains", nameHindi)
        qsMoradabad = qsMoradabad.Filter("Name_hindi__icontains", nameHindi)
    }    
    
    // Relation Name English
    for _, relationNameEnglish := range query.RelationNameEnglish {
        qsRampur = qsRampur.Filter("Relation_name_english__icontains", relationNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Relation_name_english__icontains", relationNameEnglish)
    }    
    
    // Relation Name Hindi
    for _, relationNameHindi := range query.RelationNameHindi {
        qsRampur = qsRampur.Filter("Relation_name_hindi__icontains", relationNameHindi)
        qsMoradabad = qsMoradabad.Filter("Relation_name_hindi__icontains", relationNameHindi)
    }
    
    // Gender
    for _, gender := range query.Gender {
        qsRampur = qsRampur.Filter("Gender__exact", gender)
        qsMoradabad = qsMoradabad.Filter("Gender__exact", gender)
    }
    
    // ID Card Number
    for _, idCardNumber := range query.NameHindi {
        qsRampur = qsRampur.Filter("Id_card_number__exact", idCardNumber)
        qsMoradabad = qsMoradabad.Filter("Id_card_number__exact", idCardNumber)
    }
    
    // Ac Name English
    for _, acNameEnglish := range query.AcNameEnglish {
        qsRampur = qsRampur.Filter("Ac_name_english__icontains", acNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Ac_name_english__icontains", acNameEnglish)
    }
    
    // Ac Name Hindi
    for _, acNameHindi := range query.AcNameHindi {
        qsRampur = qsRampur.Filter("Ac_name_hindi__icontains", acNameHindi)
        qsMoradabad = qsMoradabad.Filter("Ac_name_hindi__icontains", acNameHindi)
    }
    
    // Section Name English
    for _, sectionNameEnglish := range query.SectionNameEnglish {
        qsRampur = qsRampur.Filter("Section_name_english__icontains", sectionNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Section_name_english__ilike", sectionNameEnglish)
    }
    
    // Section Name Hindi
    for _, sectionNameHindi := range query.SectionNameHindi {
        qsRampur = qsRampur.Filter("Section_name_hindi__ilike", sectionNameHindi)
        qsMoradabad = qsMoradabad.Filter("Section_name_hindi__ilike", sectionNameHindi)
    }
    
    //Religion English
    for _, religionEnglish := range query.ReligionEnglish {
        qsRampur = qsRampur.Filter("Religion_english__exact", religionEnglish)
        qsMoradabad = qsMoradabad.Filter("Religion_english__exact", religionEnglish)
    }
    
    // Religion Hindi
    for _, religionHindi := range query.ReligionHindi {
        qsRampur = qsRampur.Filter("Religion_hindi__exact", religionHindi)
        qsMoradabad = qsMoradabad.Filter("Religion_hindi__exact", religionHindi)
    }
    
    // Age
    for _, age := range query.Age {
        qsRampur = qsRampur.Filter("Age__exact", age)
        qsMoradabad = qsMoradabad.Filter("Age__exact", age)
    }
        
    // Get voters for each state
    if len(query.DistrictNameEnglish) == 0 && len(query.DistrictNameHindi) == 0 && len(query.DistrictNumber) == 0 {
        for _, state := range query.StateNumber {
            if state == 27 {
                votersCountRampur, err = qsRampur.Limit(1000).All(&votersRampur)
                votersCountMoradabad, err = qsMoradabad.Limit(1000).All(&votersMoradabad)
            }
            
            if state == 12 {
                votersCountBangalore, err = qsBangalore.Limit(1000).All(&votersBangalore)
                votersCountHubli, err = qsHubli.Limit(1000).All(&votersHubli)
            }
            
        }
    } 
    
    // District Name Hindi
    for _, districtNameHindi := range query.DistrictNameHindi {
        if districtNameHindi == "मुरादाबाद" {
            votersCountMoradabad, err = qsMoradabad.Limit(1000).All(&votersMoradabad)
        }
        
        if districtNameHindi == "रामपुर" {
            votersCountRampur, err = qsRampur.Limit(1000).All(&votersRampur)
        }
    }
        
    // District Name English
    for _, districtNameEnglish := range query.DistrictNameEnglish {
        if districtNameEnglish == "Moradabad" || districtNameEnglish == "moradabad" {
            votersCountMoradabad, err = qsMoradabad.Limit(1000).All(&votersMoradabad)
        }
        
        if districtNameEnglish == "Rampur" || districtNameEnglish == "rampur" {
            votersCountRampur, err = qsRampur.Limit(1000).All(&votersRampur)
        }
    }
    
    // Get voters for each district
    for _, district := range query.DistrictNumber {
        if district == 20 {
            votersCountRampur, err = qsRampur.Limit(1000).All(&votersRampur)
        }
        
        if district == 19 {
            votersCountMoradabad, err = qsMoradabad.Limit(1000).All(&votersMoradabad)
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
        votersCount int64
        votersCountRampur int64
        votersCountMoradabad int64
        votersCountBangalore int64
        votersCountHubli int64
        muslimVotersCount int64
        otherVotersCount int64
        maleVotersCount int64
        femaleVotersCount int64
        muslimMaleVotersCount int64
        muslimFemaleVotersCount int64
        otherMaleVotersCount int64
        otherFemaleVotersCount int64
        muslimVotersP float64
        otherVotersP float64
        maleVotersP float64
        femaleVotersP float64
        muslimMaleVotersP float64
        muslimFemaleVotersP float64
        otherMaleVotersP float64
        otherFemaleVotersP float64
        muslimVotersCountRampur int64
        muslimVotersCountMoradabad int64
        muslimVotersCountBangalore int64
        muslimVotersCountHubli int64
        otherVotersCountRampur int64
        otherVotersCountMoradabad int64
        otherVotersCountBangalore int64
        otherVotersCountHubli int64
        maleVotersCountRampur int64
        maleVotersCountMoradabad int64
        maleVotersCountBangalore int64
        maleVotersCountHubli int64
        femaleVotersCountRampur int64
        femaleVotersCountMoradabad int64
        femaleVotersCountBangalore int64
        femaleVotersCountHubli int64
        muslimMaleVotersCountRampur int64
        muslimMaleVotersCountMoradabad int64
        muslimMaleVotersCountBangalore int64
        muslimMaleVotersCountHubli int64
        muslimFemaleVotersCountRampur int64
        muslimFemaleVotersCountMoradabad int64
        muslimFemaleVotersCountBangalore int64
        muslimFemaleVotersCountHubli int64
        otherMaleVotersCountRampur int64
        otherMaleVotersCountMoradabad int64
        otherMaleVotersCountBangalore int64
        otherMaleVotersCountHubli int64
        otherFemaleVotersCountRampur int64
        otherFemaleVotersCountMoradabad int64
        otherFemaleVotersCountBangalore int64
        otherFemaleVotersCountHubli int64
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
    
    
    // Apply filters for each query string    
    // Ac Number
    for _, acNumber := range query.AcNumber {
        qsRampur = qsRampur.Filter("Ac_number__exact", acNumber)
        qsMoradabad = qsMoradabad.Filter("Ac_number__exact", acNumber)
    }
     
    // Part Number
    for _, partNumber := range query.PartNumber {
        qsRampur = qsRampur.Filter("Part_number__exact", partNumber)
        qsMoradabad = qsMoradabad.Filter("Part_number__exact", partNumber)
    }   
    
    // Section Number
    for _, sectionNumber := range query.SectionNumber {
        qsRampur = qsRampur.Filter("Section_number__exact", sectionNumber)
        qsMoradabad = qsMoradabad.Filter("Section_number__exact", sectionNumber)
    }
        
    // Ac Name English
    for _, acNameEnglish := range query.AcNameEnglish {
        qsRampur = qsRampur.Filter("Ac_name_english__icontains", acNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Ac_name_english__icontains", acNameEnglish)
    }
    
    // Ac Name Hindi
    for _, acNameHindi := range query.AcNameHindi {
        qsRampur = qsRampur.Filter("Ac_name_hindi__icontains", acNameHindi)
        qsMoradabad = qsMoradabad.Filter("Ac_name_hindi__icontains", acNameHindi)
    }
    
    // Section Name English
    for _, sectionNameEnglish := range query.SectionNameEnglish {
        qsRampur = qsRampur.Filter("Section_name_english__icontains", sectionNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Section_name_english__ilike", sectionNameEnglish)
    }
    
    // Section Name Hindi
    for _, sectionNameHindi := range query.SectionNameHindi {
        qsRampur = qsRampur.Filter("Section_name_hindi__ilike", sectionNameHindi)
        qsMoradabad = qsMoradabad.Filter("Section_name_hindi__ilike", sectionNameHindi)
    }
            
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
        votersCount int64
        votersResult int64
        votersCountQuery int64
        votersCountScope int64
        votersCountRampur int64
        votersCountMoradabad int64
        votersCountBangalore int64
        votersCountHubli int64
        muslimVotersCount int64
        otherVotersCount int64
        maleVotersCount int64
        femaleVotersCount int64
        muslimMaleVotersCount int64
        muslimFemaleVotersCount int64
        otherMaleVotersCount int64
        otherFemaleVotersCount int64
        muslimVotersP float64
        otherVotersP float64
        maleVotersP float64
        femaleVotersP float64
        muslimMaleVotersP float64
        muslimFemaleVotersP float64
        otherMaleVotersP float64
        otherFemaleVotersP float64
        muslimVotersCountRampur int64
        muslimVotersCountMoradabad int64
        muslimVotersCountBangalore int64
        muslimVotersCountHubli int64
        otherVotersCountRampur int64
        otherVotersCountMoradabad int64
        otherVotersCountBangalore int64
        otherVotersCountHubli int64
        maleVotersCountRampur int64
        maleVotersCountMoradabad int64
        maleVotersCountBangalore int64
        maleVotersCountHubli int64
        femaleVotersCountRampur int64
        femaleVotersCountMoradabad int64
        femaleVotersCountBangalore int64
        femaleVotersCountHubli int64
        muslimMaleVotersCountRampur int64
        muslimMaleVotersCountMoradabad int64
        muslimMaleVotersCountBangalore int64
        muslimMaleVotersCountHubli int64
        muslimFemaleVotersCountRampur int64
        muslimFemaleVotersCountMoradabad int64
        muslimFemaleVotersCountBangalore int64
        muslimFemaleVotersCountHubli int64
        otherMaleVotersCountRampur int64
        otherMaleVotersCountMoradabad int64
        otherMaleVotersCountBangalore int64
        otherMaleVotersCountHubli int64
        otherFemaleVotersCountRampur int64
        otherFemaleVotersCountMoradabad int64
        otherFemaleVotersCountBangalore int64
        otherFemaleVotersCountHubli int64
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
    
    
    // Apply filters for each query string    
    // Ac Number Query
    for _, acNumber := range queries.Query.AcNumber {
        qsRampur = qsRampur.Filter("Ac_number__exact", acNumber)
        qsMoradabad = qsMoradabad.Filter("Ac_number__exact", acNumber)
    }
    
    // Apply filters for each query string    
    // Ac Number Scope
    for _, acNumber := range queries.Scope.AcNumber {
        qsScopeRampur = qsScopeRampur.Filter("Ac_number__exact", acNumber)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Ac_number__exact", acNumber)
    }
     
    // Part Number Query
    for _, partNumber := range queries.Query.PartNumber {
        qsRampur = qsRampur.Filter("Part_number__exact", partNumber)
        qsMoradabad = qsMoradabad.Filter("Part_number__exact", partNumber)
    }   
    
    // Part Number Scope
    for _, partNumber := range queries.Query.PartNumber {
        qsScopeRampur = qsScopeRampur.Filter("Part_number__exact", partNumber)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Part_number__exact", partNumber)
    }  
    
    // Section Number Query
    for _, sectionNumber := range queries.Query.SectionNumber {
        qsRampur = qsRampur.Filter("Section_number__exact", sectionNumber)
        qsMoradabad = qsMoradabad.Filter("Section_number__exact", sectionNumber)
    }
    
    // Section Number Scope
    for _, sectionNumber := range queries.Query.SectionNumber {
        qsScopeRampur = qsScopeRampur.Filter("Section_number__exact", sectionNumber)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Section_number__exact", sectionNumber)
    }
        
    // Ac Name English Query
    for _, acNameEnglish := range queries.Query.AcNameEnglish {
        qsRampur = qsRampur.Filter("Ac_name_english__icontains", acNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Ac_name_english__icontains", acNameEnglish)
    }
    
    // Ac Name English Scope
    for _, acNameEnglish := range queries.Query.AcNameEnglish {
        qsScopeRampur = qsScopeRampur.Filter("Ac_name_english__icontains", acNameEnglish)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Ac_name_english__icontains", acNameEnglish)
    }
    
    // Ac Name Hindi Query
    for _, acNameHindi := range queries.Query.AcNameHindi {
        qsRampur = qsRampur.Filter("Ac_name_hindi__icontains", acNameHindi)
        qsMoradabad = qsMoradabad.Filter("Ac_name_hindi__icontains", acNameHindi)
    }
    
    // Ac Name Hindi Scope
    for _, acNameHindi := range queries.Query.AcNameHindi {
        qsScopeRampur = qsScopeRampur.Filter("Ac_name_hindi__icontains", acNameHindi)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Ac_name_hindi__icontains", acNameHindi)
    }
    
    // Section Name English Query
    for _, sectionNameEnglish := range queries.Query.SectionNameEnglish {
        qsRampur = qsRampur.Filter("Section_name_english__icontains", sectionNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Section_name_english__ilike", sectionNameEnglish)
    }
    
    // Section Name English Scope
    for _, sectionNameEnglish := range queries.Query.SectionNameEnglish {
        qsScopeRampur = qsScopeRampur.Filter("Section_name_english__icontains", sectionNameEnglish)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Section_name_english__ilike", sectionNameEnglish)
    }
    
    // Section Name Hindi Query
    for _, sectionNameHindi := range queries.Query.SectionNameHindi {
        qsRampur = qsRampur.Filter("Section_name_hindi__ilike", sectionNameHindi)
        qsMoradabad = qsMoradabad.Filter("Section_name_hindi__ilike", sectionNameHindi)
    }
    
    // Section Name Hindi Scope
    for _, sectionNameHindi := range queries.Query.SectionNameHindi {
        qsScopeRampur = qsScopeRampur.Filter("Section_name_hindi__ilike", sectionNameHindi)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Section_name_hindi__ilike", sectionNameHindi)
    }
            
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
    
    // Apply filters for each query string
    // Name English Query
    for _, nameEnglish := range queries.Query.NameEnglish {
        qsRampur = qsRampur.Filter("Name_english__icontains", nameEnglish)
        qsMoradabad = qsMoradabad.Filter("Name_english__icontains", nameEnglish)
    }    

    // Name English Scope
    for _, nameEnglish := range queries.Scope.NameEnglish {
        qsScopeRampur = qsScopeRampur.Filter("Name_english__icontains", nameEnglish)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Name_english__icontains", nameEnglish)
    }
    
    // Name Hindi Query
    for _, nameHindi := range queries.Query.NameHindi {
        qsRampur = qsRampur.Filter("Name_hindi__icontains", nameHindi)
        qsMoradabad = qsMoradabad.Filter("Name_hindi__icontains", nameHindi)
    }    
    
    // Name Hindi Scope
    for _, nameHindi := range queries.Scope.NameHindi {
        qsScopeRampur = qsScopeRampur.Filter("Name_hindi__icontains", nameHindi)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Name_hindi__icontains", nameHindi)
    } 
    
    // Relation Name English Query
    for _, relationNameEnglish := range queries.Query.RelationNameEnglish {
        qsRampur = qsRampur.Filter("Relation_name_english__icontains", relationNameEnglish)
        qsMoradabad = qsMoradabad.Filter("Relation_name_english__icontains", relationNameEnglish)
    }
    
    // Relation Name English Scope
    for _, relationNameEnglish := range queries.Scope.RelationNameEnglish {
        qsScopeRampur = qsScopeRampur.Filter("Relation_name_english__icontains", relationNameEnglish)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Relation_name_english__icontains", relationNameEnglish)
    }    
    
    // Relation Name Hindi Query
    for _, relationNameHindi := range queries.Query.RelationNameHindi {
        qsRampur = qsRampur.Filter("Relation_name_hindi__icontains", relationNameHindi)
        qsMoradabad = qsMoradabad.Filter("Relation_name_hindi__icontains", relationNameHindi)
    }
    
    // Relation Name Hindi Scope
    for _, relationNameHindi := range queries.Scope.RelationNameHindi {
        qsScopeRampur = qsScopeRampur.Filter("Relation_name_hindi__icontains", relationNameHindi)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Relation_name_hindi__icontains", relationNameHindi)
    }
    
    // Gender Query
    for _, gender := range queries.Query.Gender {
        qsRampur = qsRampur.Filter("Gender__exact", gender)
        qsMoradabad = qsMoradabad.Filter("Gender__exact", gender)
    }
    
    // Gender Scope
    for _, gender := range queries.Scope.Gender {
        qsScopeRampur = qsScopeRampur.Filter("Gender__exact", gender)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Gender__exact", gender)
    }
    
    //Religion English Query
    for _, religionEnglish := range queries.Query.ReligionEnglish {
        qsRampur = qsRampur.Filter("Religion_english__exact", religionEnglish)
        qsMoradabad = qsMoradabad.Filter("Religion_english__exact", religionEnglish)
    }
    
    //Religion English Scope
    for _, religionEnglish := range queries.Scope.ReligionEnglish {
        qsScopeRampur = qsScopeRampur.Filter("Religion_english__exact", religionEnglish)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Religion_english__exact", religionEnglish)
    }
    
    // Religion Hindi Query
    for _, religionHindi := range queries.Query.ReligionHindi {
        qsRampur = qsRampur.Filter("Religion_hindi__exact", religionHindi)
        qsMoradabad = qsMoradabad.Filter("Religion_hindi__exact", religionHindi)
    }
    
    // Religion Hindi Scope
    for _, religionHindi := range queries.Scope.ReligionHindi {
        qsScopeRampur = qsScopeRampur.Filter("Religion_hindi__exact", religionHindi)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Religion_hindi__exact", religionHindi)
    }
    
    // Age Query
    for _, age := range queries.Query.Age {
        qsRampur = qsRampur.Filter("Age__exact", age)
        qsMoradabad = qsMoradabad.Filter("Age__exact", age)
    }
    
    // Age Scope
    for _, age := range queries.Scope.Age {
        qsScopeRampur = qsScopeRampur.Filter("Age__exact", age)
        qsScopeMoradabad = qsScopeMoradabad.Filter("Age__exact", age)
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