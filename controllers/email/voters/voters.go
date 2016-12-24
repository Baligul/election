/*
   Create and Send Pdf

   Voter Slips

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":["PILA TALAB"],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' "http://107.178.208.219:80/api/email/voters/slips?mobile_no=9343352734&token=5c3daf85732856f9"

   Voter List
   
   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":["PILA TALAB"],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' "http://107.178.208.219:80/api/email/voters/list?mobile_no=9343352734&token=5c3daf85732856f9"
*/

package voters

import (
	"encoding/json"
	"fmt"
	"github.com/scorredoira/email"
	"net/mail"
	"net/smtp"
	//"os"
	"os/exec"
	"strconv"
	"strings"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/Baligul/election/lib/html"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	//cleanup()
}

type VotersCtrl struct {
	beego.Controller
}

func (e *VotersCtrl) CreateAndEmailPdf() {
	var (
		votersCountRampur    int64
		votersCountMoradabad int64
		votersCountBijnor    int64
		votersCountBangalore int64
		votersCountHubli     int64
		err                  error
		num                  int64
		user                 []*modelAccounts.Account
		filepath             string
		voters               modelVoters.Voters
		votersMoradabad      []*modelVoters.Voter_19
		votersRampur         []*modelVoters.Voter_20
		votersBijnor         []*modelVoters.Voter_21
		votersBangalore      []*modelVoters.Voter
		votersHubli          []*modelVoters.Voter
	)
	mobileNo, _ := e.GetInt("mobile_no")
	token := e.GetString("token")
	key := e.Ctx.Input.Param(":key")

	if key != "list" && key != "slips" {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Wrong parameter passed. Please check your URL again.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

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

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	filepath = createFilePath(query, key)
	// TODO: We need to include back the below condition once our voter data gets perfect
	//if _, err := os.Stat(filepath + ".pdf"); os.IsNotExist(err) || filepath == match {
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
					_, err = qsRampur.Limit(-1).All(&votersRampur)
					if err != nil {
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Db Error Rampur. Unable to get the voters.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}
					votersCountMoradabad, _ = qsMoradabad.Count()
					_, err = qsMoradabad.Limit(-1).All(&votersMoradabad)
					if err != nil {
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Db Error Moradabad. Unable to get the voters.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}
					votersCountBijnor, _ = qsBijnor.Count()
					_, err = qsBijnor.Limit(-1).All(&votersBijnor)
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
					_, err = qsBangalore.Limit(-1).All(&votersBangalore)
					if err != nil {
						responseStatus := modelVoters.NewResponseStatus()
						responseStatus.Response = "error"
						responseStatus.Message = fmt.Sprintf("Db Error Bangalore. Unable to get the voters.")
						responseStatus.Error = err.Error()
						e.Data["json"] = &responseStatus
						e.ServeJSON()
					}
					votersCountHubli, _ = qsHubli.Count()
					_, err = qsHubli.Limit(-1).All(&votersHubli)
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
				_, err = qsMoradabad.Limit(-1).All(&votersMoradabad)
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
				_, err = qsRampur.Limit(-1).All(&votersRampur)
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
				_, err = qsBijnor.Limit(-1).All(&votersBijnor)
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
				_, err = qsMoradabad.Limit(-1).All(&votersMoradabad)
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
				_, err = qsRampur.Limit(-1).All(&votersRampur)
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
				_, err = qsBijnor.Limit(-1).All(&votersBijnor)
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
				_, err = qsMoradabad.Limit(-1).All(&votersMoradabad)
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
				_, err = qsRampur.Limit(-1).All(&votersRampur)
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
				_, err = qsBijnor.Limit(-1).All(&votersBijnor)
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

		// PDF creation code start here
		if key == "slips" {
			err = html.GenerateHtmlFile("templates/voter_slips.html.tmpl", voters, filepath+".html")
		}

		if key == "list" {
			err = html.GenerateHtmlFile("templates/voter_list.html.tmpl", voters, filepath+".html")
		}

		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}

		err = createPdf(filepath)
		if err != nil {
			responseStatus := modelVoters.NewResponseStatus()
			responseStatus.Response = "error"
			responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
			responseStatus.Error = err.Error()
			e.Data["json"] = &responseStatus
			e.ServeJSON()
		}
	//}

	err = sendEmailWithAttachment(user[0].Email, user[0].Display_name, filepath+".pdf")
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	responseStatus := modelVoters.NewResponseStatus()
	responseStatus.Response = "ok"
	responseStatus.Message = fmt.Sprintf("The file has been sent Successfully.")
	e.Data["json"] = &responseStatus
	e.ServeJSON()
}

func createPdf(filepath string) error {
	// PDF creation code start here
	// converts ".html" file to ".pdf"
	err := exec.Command("/usr/local/bin/wkhtmltopdf", filepath+".html", filepath+".pdf").Run()
	if err != nil {
		return err
	}

	return nil
}

func sendEmailWithAttachment(toEmail string, displayName string, filepath string) error {
	// compose the message
	m := email.NewMessage(strings.TrimPrefix(filepath, "Downloads/"), "Dear "+displayName+"!\n\nPlease find attached the required file.\n\nThanks & Regards,\nElectionUBDA Team")
	m.From = mail.Address{Name: "ElectionUBDA Team", Address: "electionubda@gmail.com"}
	m.To = []string{toEmail}

	// add attachments
	if err := m.Attach(filepath); err != nil {
		return err
	}

	// send it
	auth := smtp.PlainAuth("", "electionubda@gmail.com", "hu123*ElectionUBDA", "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}

	return nil
}

func createFilePath(query *modelVoters.Query, key string) string {
	var filepath string
	var match string

	if key == "list" {
		filepath = "Downloads/voters_list"
		match = "Downloads/voters_list"
	}

	if key == "slips" {
		filepath = "Downloads/voters_slips"
		match = "Downloads/voters_slips"
	}

	if len(query.StateNumber) == 1 && query.StateNumber[0] == 27 {
		filepath = "Downloads/UP"
	}

	if len(query.DistrictNameEnglish) == 1 {
		if filepath == match {
			filepath = "Downloads/" + query.DistrictNameEnglish[0]
		} else {
			filepath = filepath + "-" + query.DistrictNameEnglish[0]
		}
	}

	if len(query.AcNameEnglish) == 1 {
		if filepath == match {
			filepath = "Downloads/" + query.AcNameEnglish[0]
		} else {
			filepath = filepath + "-" + query.AcNameEnglish[0]
		}
	}

	if len(query.SectionNameEnglish) == 1 {
		if filepath == match {
			filepath = "Downloads/" + query.SectionNameEnglish[0]
		} else {
			filepath = filepath + "-" + query.SectionNameEnglish[0]
		}
	}

	if len(query.PartNumber) == 1 {
		if filepath == match {
			filepath = "Downloads/PN_" + strconv.Itoa(query.PartNumber[0])
		} else {
			filepath = filepath + "-PN_" + strconv.Itoa(query.PartNumber[0])
		}
	}

	if len(query.SerialNumberInPart) == 1 {
		if filepath == match {
			filepath = "Downloads/SNIP_" + strconv.Itoa(query.SerialNumberInPart[0])
		} else {
			filepath = filepath + "-SNIP_" + strconv.Itoa(query.SerialNumberInPart[0])
		}
	}

	if len(query.ReligionEnglish) == 1 {
		if filepath == match {
			filepath = "Downloads/" + query.ReligionEnglish[0]
		} else {
			filepath = filepath + "-" + query.ReligionEnglish[0]
		}
	}

	if len(query.Vote) == 1 {
		if filepath == match {
			filepath = "Downloads/Vote_" + strconv.Itoa(query.Vote[0])
		} else {
			filepath = filepath + "-Vote_" + strconv.Itoa(query.Vote[0])
		}
	}

	if filepath != match {
		if key == "list" {
			filepath = filepath + "-voters_list"
		}

		if key == "slips" {
			filepath = filepath + "-voters_slips"
		}
	}

	return filepath
}