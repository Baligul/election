/*
   Create and Send Pdf

   Voter Slips

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":["PILA TALAB"],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' "http://107.178.208.219:80/api/email/voters/slips?mobile_no=9343352734&token=b4704cf9a3dc3faa"

   Voter List

   curl -X POST -H "Content-Type: application/json" -d '{"state_number":[], "district_number":[20], "voter_id":[],"ac_number":[],"part_number":[],"section_number":[],"serial_number_in_part":[],"name_english":[],"name_hindi":[],"relation_name_english":[],"relation_name_hindi":[],"gender":[],"id_card_number":[], "district_name_hindi":[],"district_name_english":[],"ac_name_english":[],"ac_name_hindi":[],"section_name_english":["PILA TALAB"],"section_name_hindi":[],"religion_english":[],"religion_hindi":[],"age":[],"vote":[],"email":[],"mobile_no":[],"image":[]}' "http://107.178.208.219:80/api/email/voters/list?mobile_no=9343352734&token=b4704cf9a3dc3faa"
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

	"github.com/Baligul/election/formattime"
	"github.com/Baligul/election/lib/html"
	"github.com/Baligul/election/logs"
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

	// Log the request
	_ = logs.WriteLogs("logs/logs.txt", fmt.Sprintf("Mobile no.: %d, Request: POST /api/email/voters/%s, Json: %s", mobileNo, key, e.Ctx.Input.RequestBody))

	if key != "list" && key != "slips" {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Wrong parameter passed. Please check your URL again.")
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Wrong parameter passed. Please check your URL again.")
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	if mobileNo == 0 || token == "" {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Mobile number or token is wrong.")
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
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
		_ = logs.WriteLogs("logs/error_logs.txt", "Update Last Login in Send Email Voters API: "+err.Error())
	}

	inputJson := e.Ctx.Input.RequestBody
	query := new(modelVoters.Query)

	err = json.Unmarshal(inputJson, &query)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
		if len(strings.TrimSpace(strings.Split(sectionNameEnglish, "->")[0])) > 0 {
			condSectionNameEnglish = condSectionNameEnglish.Or("Section_name_english__exact", strings.TrimSpace(strings.Split(sectionNameEnglish, "->")[0]))
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
					// Log the error
					_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
				// Log the error
				_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
			// Log the error
			_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
		voters.File_title = createFileTitle(query)
		err = html.GenerateHtmlFile("templates/voter_list.html.tmpl", voters, filepath+".html")
	}

	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	err = createPdf(filepath)
	if err != nil {
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
		// Log the error
		_ = logs.WriteLogs("logs/error_logs.txt", "Send Email Voters API: "+err.Error())
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
	var voteVal string

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
			filepath = "Downloads/" + strings.TrimSpace(strings.Split(query.SectionNameEnglish[0], "->")[0])
		} else {
			filepath = filepath + "-" + strings.TrimSpace(strings.Split(query.SectionNameEnglish[0], "->")[0])
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
		if query.Vote[0] == 1 {
			voteVal = "Voted"
		}
		if query.Vote[0] == 0 {
			voteVal = "NonVoted"
		}
		if filepath == match {
			filepath = "Downloads/" + voteVal
		} else {
			filepath = filepath + "-" + voteVal
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

func createFileTitle(query *modelVoters.Query) string {
	var (
		fileTitle      string
		match          string
		voteVal        string
		religion       string
		religionPrefix string
	)

	fileTitle = "मतदाताओं की सूची"
	match = "मतदाताओं की सूची"

	if len(query.StateNumber) == 1 && query.StateNumber[0] == 27 {
		fileTitle = "उत्तर प्रदेश"
	}

	if len(query.DistrictNameEnglish) == 1 {
		if fileTitle == match {
			fileTitle = getHindiDistrictName(query.DistrictNameEnglish[0])
		} else {
			fileTitle = getHindiDistrictName(query.DistrictNameEnglish[0]) + ", " + fileTitle
		}
	}

	if len(query.AcNameEnglish) == 1 {
		if fileTitle == match {
			fileTitle = getHindiAcName(query.AcNameEnglish[0])
		} else {
			fileTitle = getHindiAcName(query.AcNameEnglish[0]) + ", " + fileTitle
		}
	}

	if len(query.SectionNameEnglish) == 1 {
		if fileTitle == match {
			fileTitle = getHindiSectionName(strings.TrimSpace(strings.Split(query.SectionNameEnglish[0], "->")[0]))
		} else {
			fileTitle = getHindiSectionName(strings.TrimSpace(strings.Split(query.SectionNameEnglish[0], "->")[0])) + ", " + fileTitle
		}
	}

	if len(query.PartNumber) == 1 {
		if fileTitle == match {
			fileTitle = "भाग संख्या " + strconv.Itoa(query.PartNumber[0])
		} else {
			fileTitle = "भाग संख्या " + strconv.Itoa(query.PartNumber[0]) + ", " + fileTitle
		}
	}

	if len(query.Vote) == 1 {
		if query.Vote[0] == 1 {
			voteVal = "मतदान कर चुके"
		}
		if query.Vote[0] == 0 {
			voteVal = "मतदान नहीं करने वाले"
		}
		if fileTitle == match {
			fileTitle = voteVal
		} else {
			fileTitle = fileTitle + " के " + voteVal
		}
	}

	if voteVal == "" {
		religionPrefix = " के "
	} else {
		religionPrefix = " "
	}

	if len(query.ReligionEnglish) == 1 {
		if query.ReligionEnglish[0] == "Muslim" {
			religion = religionPrefix + "मुसलमान"
		} else if query.ReligionEnglish[0] == "Others" {
			religion = religionPrefix + "अन्य"
		}

		if fileTitle == match {
			fileTitle = religion
		} else {
			fileTitle = fileTitle + religion
		}
	}

	if fileTitle != match {
		fileTitle = fileTitle + " मतदाताओं की सूची"
	}

	if len(query.Age) > 1 {
		fileTitle = fileTitle + " (आयु वर्ग: " + strconv.Itoa(query.Age[0]) + " - " + strconv.Itoa(query.Age[len(query.Age)-1]) + ")"
	}

	return fileTitle
}

func getHindiAcName(acNameEnglish string) string {
	var acNameHindi string

	switch acNameEnglish {
	case "Najibabad":
		acNameHindi = "नजीबाबाद"
	case "Nagina":
		acNameHindi = "नगीना"
	case "Barhapur":
		acNameHindi = "बरहापुर"
	case "Dhampur":
		acNameHindi = "धामपुर"
	case "Nehtaur":
		acNameHindi = "नहटौर"
	case "Bijnor":
		acNameHindi = "बिजनौर"
	case "Chandpur":
		acNameHindi = "चांदपुर"
	case "Noorpur":
		acNameHindi = "नूरपुर"
	case "Kanth":
		acNameHindi = "कांठ"
	case "Thakurdwara":
		acNameHindi = "ठाकुरद्वारा"
	case "Moradabad Rural":
		acNameHindi = "मुरादाबाद देहात"
	case "Moradabad Nagar":
		acNameHindi = "मुरादाबाद नगर"
	case "Kundarki":
		acNameHindi = "कुंदरकी"
	case "Bilari":
		acNameHindi = "बिलारी"
	case "Suar":
		acNameHindi = "सुआर"
	case "Chamraua":
		acNameHindi = "चमरव्वा"
	case "Bilaspur":
		acNameHindi = "बिलासपुर"
	case "Rampur":
		acNameHindi = "रामपुर"
	case "Milak":
		acNameHindi = "मिलक"
	}

	return acNameHindi
}

func getHindiDistrictName(districtNameEnglish string) string {
	var districtNameHindi string

	switch districtNameEnglish {
	case "Moradabad":
		districtNameHindi = "मुरादाबाद"
	case "Rampur":
		districtNameHindi = "रामपुर"
	case "Bijnor":
		districtNameHindi = "बिजनौर"
	}

	return districtNameHindi
}

func getHindiSectionName(sectionNameEnglish string) string {
	var sectionNameHindi string

	switch sectionNameEnglish {
	case ".JATNANGLA":
		sectionNameHindi = "जट नंगला"
	case "[DHURYAI":
		sectionNameHindi = "धुरयाई"
	case "001 CHAKFERI":
		sectionNameHindi = "चकफेरी"
	case "2 AMEER ALI KHAN 1":
		sectionNameHindi = "2 अमीर अली खां १"
	case "AABI HAFEEJPUR":
		sectionNameHindi = "आवी हफीजपुर"
	case "AADAMPUR":
		sectionNameHindi = "आदमपुर"
	case "AADPUR":
		sectionNameHindi = "आदपुर"
	case "AAGA-1":
		sectionNameHindi = "आगा- 1"
	case "AAGA-2":
		sectionNameHindi = "आगा- 2"
	case "AAKILPUR":
		sectionNameHindi = "आकिलपुर"
	case "AAKU":
		sectionNameHindi = "आंकू"
	case "AALAMPUR GABRI AN0":
		sectionNameHindi = "आलमपुर गावंडी आं0"
	case "AANANDIPUR":
		sectionNameHindi = "आनन्‍दीपुर"
	case "AARAJI KANDALA":
		sectionNameHindi = "आराजी कन्‍दला"
	case "AARSAL PARSAL-1":
		sectionNameHindi = "आरसल पारसल 1"
	case "AARSAL PARSAL-2":
		sectionNameHindi = "आरसल पारसल 2"
	case "AARSAL PARSAL-3":
		sectionNameHindi = "आरसल पारसल 3"
	case "AASFABAD CHAMAN":
		sectionNameHindi = "आसफाबाद चमन"
	case "AASPUR NAWADA":
		sectionNameHindi = "आसपुर नवादा"
	case "AASRA KHERA":
		sectionNameHindi = "आसरा खेडा"
	case "ABABKARPUR":
		sectionNameHindi = "अबाबकरपुर"
	case "ABBAS NAGAR SUAR":
		sectionNameHindi = "अब्‍बासनगर स्‍वार"
	case "ABBAS NAGAR":
		sectionNameHindi = "अब्‍बास नगर"
	case "ABBASNAGAR TANDA":
		sectionNameHindi = "अब्‍बासनगर टाण्‍डा"
	case "ABBASPUR":
		sectionNameHindi = "अब्‍बासपुर"
	case "ABDIPUR KALYAN":
		sectionNameHindi = "अब्‍दीपुर कल्‍याण"
	case "ABDIPURHARVANSH":
		sectionNameHindi = "अब्‍दीपुरहरवंश"
	case "ABDUL KHALIKPUR BALRAM":
		sectionNameHindi = "अब्‍दुल खालिकपुर बलराम"
	case "ABDULAJIJPUR":
		sectionNameHindi = "अब्‍दुलअजीजपुर"
	case "ABDULLA MADHYA":
		sectionNameHindi = "अब्‍दुल्‍ला मध्‍य"
	case "ABDULLADAKSHINI":
		sectionNameHindi = "अब्दुल्लादक्षिणी"
	case "ABDULLANAGAR URF DANDA":
		sectionNameHindi = "अब्‍दुल्‍लानगर उर्फ डान्‍डा"
	case "ABDULLANAGAR":
		sectionNameHindi = "अब्‍दुल्‍लानगर"
	case "ABDULLAPUR LEDA":
		sectionNameHindi = "अब्दुल्लापुर लेदा"
	case "ABDULLAPUR":
		sectionNameHindi = "अब्‍दुल्‍लापुर"
	case "ABDULLAPURKURASHI":
		sectionNameHindi = "अब्‍दुल्‍लापुर कुरैशी"
	case "ABDULLAPURVI":
		sectionNameHindi = "अब्दुल्लापूर्वी"
	case "ABDULLPUR DAHANA":
		sectionNameHindi = "अब्‍दुल्‍लापुर दहाना"
	case "ABDULPUR MUNNA":
		sectionNameHindi = "अब्‍दुलपुर मुन्‍ना"
	case "ABDULREHMANPUR JEVAN":
		sectionNameHindi = "अ०रहमानपुर जीवन"
	case "ABHANPUR NARAULI":
		sectionNameHindi = "अभनपुर नरौली"
	case "ABHAYRAJPUR":
		sectionNameHindi = "अभयराजपुर"
	case "ABID NAGAR TEELA JAYANTIPUR":
		sectionNameHindi = "आबिद नगर टीला जयन्‍तीपुर"
	case "ABIDNAGAR URF DHUDLI":
		sectionNameHindi = "आविदनगर उर्फ धूधली"
	case "ABIDPUR":
		sectionNameHindi = "आबिदपुर"
	case "ABILFAJALPURBANI":
		sectionNameHindi = "अबुलफजलपुरबनी"
	case "ABU SAIDPUR":
		sectionNameHindi = "अब्‍बू सईदपुर"
	case "ABUL KHAIRPUR BANGAR":
		sectionNameHindi = "अबुल खैरपुर बंगर"
	case "ABULFAJALPUR KHAS":
		sectionNameHindi = "अबुलफजलपुर खास"
	case "ABUNASRAPUR":
		sectionNameHindi = "अब्‍बुनासरपुर"
	case "ABUPUR KHURD":
		sectionNameHindi = "अबूपुर खुर्द"
	case "ABUSAIDAPUR":
		sectionNameHindi = "अबूसैदपुर"
	case "ACHARAJAN":
		sectionNameHindi = "अचारजान"
	case "ACHARJAN":
		sectionNameHindi = "आचारजान"
	case "ADAL PUR":
		sectionNameHindi = "अदल पुर"
	case "ADAMPUR KAKWADA":
		sectionNameHindi = "आदमपुर ककवाडा"
	case "ADAMPUR":
		sectionNameHindi = "आदमपुर"
	case "ADARSH COLONY ANSHIK":
		sectionNameHindi = "आदर्श कालोनी आंशिक"
	case "ADARSH COLONY":
		sectionNameHindi = "आदर्श कालोनी"
	case "ADARSH NAGAR":
		sectionNameHindi = "आदर्श नगर"
	case "ADHAMPUR":
		sectionNameHindi = "अधमपुर"
	case "ADILNAGAR":
		sectionNameHindi = "आदिलनगर"
	case "ADONGALI":
		sectionNameHindi = "आदोनगली"
	case "ADREHMANPURMEHJANG":
		sectionNameHindi = "अ0रहमानपुरमहजंग"
	case "AFAGANAN AN.":
		sectionNameHindi = "आफगानान आं0"
	case "AFARIDIYAN-1":
		sectionNameHindi = "आफरीदियान-1"
	case "AFARIDIYAN-2":
		sectionNameHindi = "आफरीदियान-2"
	case "AFGALPUR HARRAY":
		sectionNameHindi = "अफजलपुर हर्रे"
	case "AFGANAN 01":
		sectionNameHindi = "अफगानान ०१"
	case "AFGANAN 02":
		sectionNameHindi = "अफगानान 02"
	case "AFGANAN AN.":
		sectionNameHindi = "अफगानान आं०"
	case "AFGANAN PASHCHIM A.":
		sectionNameHindi = "आफगानान पश्चिम आ0"
	case "AFGANAN PASHCHMI A.":
		sectionNameHindi = "अफगानान पश्चिमी आं0"
	case "AFGANAN PURVI A.":
		sectionNameHindi = "अफगानान पूर्वी आं0"
	case "AFGANAN UTTARI":
		sectionNameHindi = "अफगानान उत्‍तरी"
	case "AFGANAN UTTRI A.":
		sectionNameHindi = "अफगानान उत्‍तरी आं0"
	case "AFGANAN":
		sectionNameHindi = "अफगानान"
	case "AFISARS BANGALE":
		sectionNameHindi = "आफिसर्स बंगले"
	case "AFJALPIRBHAU":
		sectionNameHindi = "अफजलपुर भाउ"
	case "AFJALPUR BALDANI":
		sectionNameHindi = "अफजलपुर बल्‍दानी"
	case "AFRASHIYAN GANJ":
		sectionNameHindi = "अफराशियान गंज"
	case "AFZALPUR PATTI":
		sectionNameHindi = "अफजलपुर पटटी"
	case "AGAPUR-1":
		sectionNameHindi = "आगापुर-1"
	case "AGAPUR-2":
		sectionNameHindi = "आगापुर -2"
	case "AGAPUR-3":
		sectionNameHindi = "आगापुर -3"
	case "AGAPUR-4":
		sectionNameHindi = "आगापुर -4"
	case "AGAPUR-5":
		sectionNameHindi = "आगापुर -5"
	case "AGAPUR-6":
		sectionNameHindi = "आगापुर -6"
	case "AGAPUR":
		sectionNameHindi = "आगापुर"
	case "AGLAGA -1":
		sectionNameHindi = "अगलगा 1"
	case "AGLAGA -2":
		sectionNameHindi = "अगलगा 2"
	case "AGRI":
		sectionNameHindi = "अगरी"
	case "AGUPURA PYARA":
		sectionNameHindi = "अगुपुरा प्‍यारा"
	case "AGWANNPUR":
		sectionNameHindi = "अगवानपुर"
	case "AGWANPUR":
		sectionNameHindi = "अगवानपुर"
	case "AHAMAD NAGAR JAITVADHA AANSHIK":
		sectionNameHindi = "अहमद नगर जैतवाड़ा आंशिक"
	case "AHAMAD NAGAR PADMPUR":
		sectionNameHindi = "अहमदनगर पदमपुर"
	case "AHAMADPUR BHARTA":
		sectionNameHindi = "अहमदपुर भरता"
	case "AHAMDABAD KASAURA":
		sectionNameHindi = "अहमदाबाद कसौरा"
	case "AHAMDABAD":
		sectionNameHindi = "अहमदाबाद"
	case "AHAMDAPUR":
		sectionNameHindi = "अहमदपुर"
	case "AHAMDNAGAR NEAR ISMAIL NAGAR":
		sectionNameHindi = "अहमदनगर निकट इस्‍माईलनगर"
	case "AHAMDNAGAR NIKAT FAIZGANJ":
		sectionNameHindi = "अहदमनगर निकट फैजगंज"
	case "AHARAULA":
		sectionNameHindi = "अहरौला"
	case "AHAROLA":
		sectionNameHindi = "अहरोला"
	case "AHIRPUR GAR.AN.":
		sectionNameHindi = "अहीरपुर गै0आ0"
	case "AHLADPUR KHEM URF RAIPUR":
		sectionNameHindi = "अहलादपुर खेम उर्फ रायपुर"
	case "AHMAD KHEL AN.":
		sectionNameHindi = "अहमद खेल आं०"
	case "AHMAD NAGAR JAGEER-1":
		sectionNameHindi = "अहमद नगर जागीर-1"
	case "AHMAD NAGAR JAGEER-2":
		sectionNameHindi = "अहमद नगर जागीर -2"
	case "AHMAD NAGAR KHEDA":
		sectionNameHindi = "अहमद नगर खेडा"
	case "AHMAD NAGAR NEAR BHAGWATIPUR":
		sectionNameHindi = "अहमद नगर नि0 भगवतीपुर"
	case "AHMAD NAGAR NIKAT JIWAI KADIM":
		sectionNameHindi = "अहमदनगर नि0 जिवाई कदीम"
	case "AHMAD NAGAR PAHADI":
		sectionNameHindi = "अहमद नगर पहाडी"
	case "AHMADABAD URF TELIPURA":
		sectionNameHindi = "अहमदाबाद ऊर्फ तेलीपुरा"
	case "AHMADABAD":
		sectionNameHindi = "अहमदावाद"
	case "AHMADGANJ URF MUKATPUR":
		sectionNameHindi = "अहमदगंज उर्फ मुकटपुर"
	case "AHMADKHEL AN.H.NO 1-137":
		sectionNameHindi = "अहमद खेल आ०म०सं०१-१३७"
	case "AHMADKHEL AN.H.NO 138-END":
		sectionNameHindi = "अहमद खेल आं०म०सं० १३८ से अंत तक"
	case "AHMADNAGAR KA A":
		sectionNameHindi = "अहमदनगर कला"
	case "AHMADNAGAR NEAR AGA":
		sectionNameHindi = "अहमदनगर नि0 आगा"
	case "AHMADNAGAR NEAR BHAGWANT NAGAR":
		sectionNameHindi = "अहमदनगर नि0 भगवन्‍तनगर"
	case "AHMADNAGAR NIKAT KUNDANPUR":
		sectionNameHindi = "अहमदनगर नि0 कुन्‍दनपुर"
	case "AHMADNAGAR TARANA":
		sectionNameHindi = "अहमदनगर तराना"
	case "AHMADNAGAR THEGA":
		sectionNameHindi = "अहमदनगर थेगा"
	case "AHMADNAGAR URF MASAIYAN MOUHAMMDI":
		sectionNameHindi = "अहमदनगर ऊर्फ मडैयान मौहम्‍मदी"
	case "AHMADPUR ANANDPUR":
		sectionNameHindi = "अहमदपुर आनन्‍दपुर"
	case "AHMADPUR CHANDRU URF GDANA":
		sectionNameHindi = "अहमदपुर चन्‍दरू उर्फ गडाना"
	case "AHMADPUR MAJEED":
		sectionNameHindi = "अहमदपुर मजीद"
	case "AHMADPUR MOHIUDDINPUR":
		sectionNameHindi = "अहमदपुर मोहिददीनपुर"
	case "AHMADPUR NINGU NANGLA":
		sectionNameHindi = "अहमदपुर निंगू नंगला"
	case "AHMADPUR SADAT":
		sectionNameHindi = "अहमदपुर सादात"
	case "AHMADPUR SULEMAN":
		sectionNameHindi = "अहमदपुर सुलेमान"
	case "AHMADPUR":
		sectionNameHindi = "अहमदपुर"
	case "AHRAULA":
		sectionNameHindi = "अहरौला"
	case "AHRO-1":
		sectionNameHindi = "अहरो-1"
	case "AHRO-2":
		sectionNameHindi = "अहरो-2"
	case "AHRO-3":
		sectionNameHindi = "अहरो-3"
	case "AHRO-4":
		sectionNameHindi = "अहरो-4"
	case "AHROLA":
		sectionNameHindi = "अहरौला"
	case "AIMI-1":
		sectionNameHindi = "ऐमी-1"
	case "AIMI-2":
		sectionNameHindi = "ऐमी-2"
	case "AINCHORA-1":
		sectionNameHindi = "ऐंचोरा-1"
	case "AINCHORA-2":
		sectionNameHindi = "ऐंचोरा-2"
	case "AINJANKHEDA":
		sectionNameHindi = "ऐंजनखेडा"
	case "AJAB NAGAR":
		sectionNameHindi = "अजबनगर"
	case "AJAD NAGAR ANSHIK":
		sectionNameHindi = "आजाद नगर आंशिक"
	case "AJAD NAGAR BHOGPUR MITHONI ANSHIK":
		sectionNameHindi = "आजाद नगर भोगपुर मिठोनी आंशिक"
	case "AJAMATPUR GAIRA0":
		sectionNameHindi = "अजमतपुर गैरा0"
	case "AJAMGAR-CHOPDA":
		sectionNameHindi = "आजमनगर-चोपड़ा"
	case "AJAMPUR GAJI":
		sectionNameHindi = "आजमपुर गाजी"
	case "AJAMPUR MOHAN":
		sectionNameHindi = "आजमपुर मोहन"
	case "AJAMPUR MOHD.ALI URF KHANPUR":
		sectionNameHindi = "आजमपुर मौ0अली उर्फ खानपुर"
	case "AJAMPUR YAR MOHD":
		sectionNameHindi = "आजमपुर यार मौ०"
	case "AJAMPURMANNI":
		sectionNameHindi = "आजमपुरमन्‍नी"
	case "AJAMPURPARMA":
		sectionNameHindi = "आजमपुरपरमा"
	case "AJAYPUR":
		sectionNameHindi = "अजयपुर"
	case "AJDEV":
		sectionNameHindi = "अजदेव"
	case "AJDHAPUR":
		sectionNameHindi = "अजदहापुर"
	case "AJEETPUR-1":
		sectionNameHindi = "अजीतपुर-1"
	case "AJEETPUR-10":
		sectionNameHindi = "अजीतपुर-10"
	case "AJEETPUR-2":
		sectionNameHindi = "अजीतपुर-2"
	case "AJEETPUR-3":
		sectionNameHindi = "अजीतपुर-3"
	case "AJEETPUR-4":
		sectionNameHindi = "अजीतपुर-4"
	case "AJEETPUR-5":
		sectionNameHindi = "अजीतपुर-5"
	case "AJEETPUR-6":
		sectionNameHindi = "अजीतपुर-6"
	case "AJEETPUR-7":
		sectionNameHindi = "अजीतपुर-7"
	case "AJEETPUR-8":
		sectionNameHindi = "अजीतपुर-8"
	case "AJIJPURA":
		sectionNameHindi = "अजीजपुरा"
	case "AJIMULLA NAGAR AN0":
		sectionNameHindi = "अजीमुल्‍ला नगर आं0"
	case "AJIMULLA NAGAR":
		sectionNameHindi = "अजीमुल्‍ला नगर"
	case "AJIMULLANAGAR AN0":
		sectionNameHindi = "अजीमुल्‍ला नगर आं0"
	case "AJIMULLHANAGAR":
		sectionNameHindi = "अजीमुल्‍लानगर"
	case "AJITAPUR DASI":
		sectionNameHindi = "अजीतपुर दासी"
	case "AJITAPUR-9":
		sectionNameHindi = "अजीतपुर 9"
	case "AJJUNAGLI":
		sectionNameHindi = "अज्‍जू नंगली"
	case "AJMAPUR JAMANIMAN":
		sectionNameHindi = "आजमपुर जमनीभान"
	case "AJMAT NAGAR JYODERA":
		sectionNameHindi = "अजमत नगर ज्योडेरा"
	case "AJUPURA JAT":
		sectionNameHindi = "अजुपुरा जट"
	case "AJUPURACHOHAN":
		sectionNameHindi = "अजुपुरा चौहान"
	case "AJUPURARANI":
		sectionNameHindi = "अजुपुरा रानी"
	case "AKAB MOTI MASJID":
		sectionNameHindi = "अकब मोती मस्जिद"
	case "AKAB POST AFIS":
		sectionNameHindi = "अकब पोस्ट आफिस"
	case "AKABRAN":
		sectionNameHindi = "अकाबरान"
	case "AKBARABAD AN.H.NO 1-548":
		sectionNameHindi = "अकबरारबाद आं०म०सं०१ से ५४८"
	case "AKBARABAD AN0 H.NO 549-END":
		sectionNameHindi = "अकबराबाद आं०म०सं० ५४९ से अंत तक"
	case "AKBARABAD AN0":
		sectionNameHindi = "अकबराबाद आं०म०सं० ५४९ से अंत तक"
	case "AKBARABAD G.A.":
		sectionNameHindi = "अकबराबाद गै0आ0"
	case "AKBARAPUR":
		sectionNameHindi = "अकबरपुर"
	case "AKBARPUR ANGAKHEDI":
		sectionNameHindi = "अकबरपुर अंगखेडी"
	case "AKBARPUR ANWLA":
		sectionNameHindi = "अकबरपुर आंवला"
	case "AKBARPUR CHEDRY":
		sectionNameHindi = "अकबरपुर चैदरी"
	case "AKBARPUR CHOGANWA AN0":
		sectionNameHindi = "अकबरपुर चौगावां आं0"
	case "AKBARPUR DEVIDASWALA":
		sectionNameHindi = "अकबरपुर देवीदासवाला"
	case "AKBARPUR DEVMAL":
		sectionNameHindi = "अकबरपुर देवमल"
	case "AKBARPUR GARAV":
		sectionNameHindi = "अकबरपुर गारव"
	case "AKBARPUR KHAS":
		sectionNameHindi = "अकबरपुर खास"
	case "AKBARPUR MAJRA SAINDVAR":
		sectionNameHindi = "अकबरपुर मजरा सैन्‍दवार"
	case "AKBARPUR RADHE GAIR AB.":
		sectionNameHindi = "अकबरपुर राधे गैर आं0"
	case "AKBARPUR SIHALI":
		sectionNameHindi = "अकबरपुर सिहाली"
	case "AKBARPUR TIGRI A.":
		sectionNameHindi = "अकबरपुर तिगरी आ०"
	case "AKBARPUR URF CHOTAPUR":
		sectionNameHindi = "अफजलपुर उर्फ छोटापुर"
	case "AKBARPUR ZOZA":
		sectionNameHindi = "अकबरपुर झौझा"
	case "AKBARPUR":
		sectionNameHindi = "अकबरपुर आ०"
	case "AKBERPUR ASHA URF HAROLI":
		sectionNameHindi = "अकबरपुर आशा उर्फ हरौली"
	case "AKHADA MALLI KHAN":
		sectionNameHindi = "अखाडा मल्ली खां"
	case "AKHBAR FACTORY DHEEMRI":
		sectionNameHindi = "अखवार फैक्‍ट्री धीमरी"
	case "AKHBARABAD":
		sectionNameHindi = "अकवराबाद"
	case "AKHERA":
		sectionNameHindi = "अखेडा"
	case "AKHUN KHE LAN":
		sectionNameHindi = "आखून खे लान"
	case "AKKA BHIKANPUR AITMALI (GAIRA)":
		sectionNameHindi = "अक्‍का भीकनपुर ऐतमाली (गैरा)"
	case "AK‍KA DILARI ANSHIK":
		sectionNameHindi = "अक्‍का डिलारी आंशिक"
	case "AK‍KA DILARI":
		sectionNameHindi = "अक्‍का डिलारी"
	case "AKKA FATTU HAFIJPUR (GAIRA)":
		sectionNameHindi = "अक्‍का फत्‍तू हफीजपुर (गैरा)"
	case "AK‍KA PAN‍DE BHOJPUR ANSHIK":
		sectionNameHindi = "अक्‍का पाण्‍डे भोजपुर आंशिक"
	case "AKLASPUR":
		sectionNameHindi = "अखलासपुर"
	case "AKONDA A.":
		sectionNameHindi = "अकौन्‍धा आ०"
	case "AKONDA":
		sectionNameHindi = "अकौंदा"
	case "ALABALPUR":
		sectionNameHindi = "अलाबलपुर"
	case "ALADEENPUR BHOUGI":
		sectionNameHindi = "अलाउददीनपुर भोगी"
	case "ALADINPUR A.":
		sectionNameHindi = "आलादीनपुर आ0"
	case "ALADINPUR BHATPPURA":
		sectionNameHindi = "अलादीनपुर भटपुरा"
	case "ALAFGUNJ":
		sectionNameHindi = "अलफगंज"
	case "ALAHADADAPUR KARAR":
		sectionNameHindi = "अलहदादपुर करार"
	case "ALAHADADPUR":
		sectionNameHindi = "अलहदादपुर"
	case "ALAMPUR ADWA":
		sectionNameHindi = "आलमपुर एडवा"
	case "ALAMPUR GANGA":
		sectionNameHindi = "आलमपुर गंगा"
	case "ALAMPUR NILA":
		sectionNameHindi = "आलमपुर नीला"
	case "ALAMPUR":
		sectionNameHindi = "आलमपुर"
	case "ALAMPURASU G.ABD":
		sectionNameHindi = "आलमपुर आसू गैर आ०"
	case "ALAMPURI":
		sectionNameHindi = "आलमपुरी"
	case "ALAUDDINPUR DHANRAJ":
		sectionNameHindi = "अलाउददीनपुर धनराज"
	case "ALAUDDINPUR":
		sectionNameHindi = "अलाऊददीनपुर"
	case "ALAUDEENPUR":
		sectionNameHindi = "अलाउदीनपुर"
	case "ALAVALPUR":
		sectionNameHindi = "अलावलपुर"
	case "ALAWALPUR NAINU":
		sectionNameHindi = "अलावलपुर नैनू"
	case "ALAWALPUR UDDA":
		sectionNameHindi = "अलावलपुर उददा"
	case "ALBANAGALA":
		sectionNameHindi = "अलवानगला"
	case "ALEDADPUR":
		sectionNameHindi = "अलेदादपुर"
	case "ALEYALIPUR":
		sectionNameHindi = "आलेअलीपुर"
	case "ALEYARPUR URF HAJIPUR":
		sectionNameHindi = "अल्‍हैयारपुर उफ हाजीपुर"
	case "ALHADADPUR":
		sectionNameHindi = "अल्‍हैदादपुर"
	case "ALHAHEDI":
		sectionNameHindi = "आल्‍हहेडी"
	case "ALHAIPUR AN.":
		sectionNameHindi = "अल्‍हैपुर आं0"
	case "ALHAPUR URF RAMPUR":
		sectionNameHindi = "अल्हापुर उर्फ रामपुर"
	case "ALHAYPUR BEERAN":
		sectionNameHindi = "अल्‍हैपुर वीरान"
	case "ALHAYPUR MOHKAM":
		sectionNameHindi = "अल्‍हैपुर मोहकम"
	case "ALHEDADPUR MUBARAK":
		sectionNameHindi = "अल्‍हेदादपुर मुबारक"
	case "ALI NAGAR JAGEER-1":
		sectionNameHindi = "अलीनगर जागीर 1"
	case "ALI NAGAR JAGEER-2":
		sectionNameHindi = "अलीनगर जागीर 2"
	case "ALI NAGAR TAH":
		sectionNameHindi = "अलीनगर टाह"
	case "ALI NAGAR UTTARI":
		sectionNameHindi = "अलीनगर उत्‍तरी"
	case "ALIGANG":
		sectionNameHindi = "अलीगंज"
	case "ALIGANJ BENJEER-1":
		sectionNameHindi = "अलीगंज बेनजीर-1"
	case "ALIGANJ BENJEER-2":
		sectionNameHindi = "अलीगंज बेनजीर-2"
	case "ALINAGAR JUNUBI":
		sectionNameHindi = "अलीनगर जनूबी"
	case "ALINAGAR KOTA":
		sectionNameHindi = "अलीनगर कोटा"
	case "ALINAGAR PALNI":
		sectionNameHindi = "अलीनगर पालनी"
	case "ALINAGAR SHUMALI":
		sectionNameHindi = "अलीनगर शुमाली"
	case "ALINAGAR":
		sectionNameHindi = "अलीनगर"
	case "ALIPUR DAMODAR":
		sectionNameHindi = "अलीपुर दामोदर"
	case "ALIPUR GANGA":
		sectionNameHindi = "अलीपुर गंगा"
	case "ALIPUR MAKHAN":
		sectionNameHindi = "अलीपुर माखन"
	case "ALIPUR MITTHAN":
		sectionNameHindi = "अलीपुर मिठठन"
	case "ALIPUR NAGLA":
		sectionNameHindi = "अलीपुर नगंला"
	case "ALIPUR SUKHANANDPUR URF LALPUR":
		sectionNameHindi = "अलीपुर सुखानन्‍द उर्फ लालपुर"
	case "ALIPUR THEKA":
		sectionNameHindi = "अलीपुर ठेका"
	case "ALIPURA AN0":
		sectionNameHindi = "अलीपुरा आं0"
	case "ALIPURA JATT":
		sectionNameHindi = "अलीपुरा जट"
	case "ALIPURA KHALSA":
		sectionNameHindi = "अलीपुरा खालसा"
	case "ALIPURA MAJRA KUI":
		sectionNameHindi = "अलीपुरा मजरा कुई"
	case "ALIPURA":
		sectionNameHindi = "अलीपुरा"
	case "ALIPURMAN":
		sectionNameHindi = "अलीपुर मान"
	case "ALIPURMOLAD":
		sectionNameHindi = "अलीपुरमोलहड्"
	case "ALIRAJAPUR":
		sectionNameHindi = "अलीराजापुर"
	case "ALIYABAD MUU0 SABIK":
		sectionNameHindi = "अलियाबाद मुू0 साबिक"
	case "ALIYABAD":
		sectionNameHindi = "अलियाबाद"
	case "ALIYARPUR":
		sectionNameHindi = "अलियारपुर"
	case "ALLAHDADPURKHAJWA":
		sectionNameHindi = "अल्‍हैददादपुर खजुवा"
	case "ALLAPUR BHIKAN":
		sectionNameHindi = "अल्‍लापुर भीकन"
	case "ALLAPUR":
		sectionNameHindi = "अल्‍लापुर"
	case "ALLAUDDINPUR":
		sectionNameHindi = "अलाउददीनपुर"
	case "ALLEHPUR":
		sectionNameHindi = "अल्हैपुर"
	case "ALMAGIRAPUR":
		sectionNameHindi = "आलमगीरपुर"
	case "AMAKHEDA SHAJRAPUR AN.":
		sectionNameHindi = "अमखेडा शंजरपुर आं0"
	case "AMAN NAGAR":
		sectionNameHindi = "अमाननगर"
	case "AMANATABAD":
		sectionNameHindi = "अमानताबाद"
	case "AMANATPUR":
		sectionNameHindi = "अमानतपुर"
	case "AMANNAGAR":
		sectionNameHindi = "अमान नगर"
	case "AMANULLAPUR":
		sectionNameHindi = "अमानुल्‍लापुर"
	case "AMARPUR KASHI":
		sectionNameHindi = "अमरपुर काशी"
	case "AMBARPUR ZAFAR":
		sectionNameHindi = "अम्‍बरपुर जाफर"
	case "AMBARPUR":
		sectionNameHindi = "अम्‍बरपुर"
	case "AMBEDKAR COLONY":
		sectionNameHindi = "अम्‍बेडकर कालौनी"
	case "AMBEDKAR NAGAR H.NO 233-324":
		sectionNameHindi = "अम्‍बेडकर नगर म०सं०२३३ से ३२४तक"
	case "AMBEDKAR NAGAR":
		sectionNameHindi = "अम्‍बेडकर नगर"
	case "AMEER ALI KHAN 02":
		sectionNameHindi = "अमीर अली खां 02"
	case "AMEERPUR MANJU":
		sectionNameHindi = "मौ0 अमीरपुरमन्‍जु"
	case "AMHEDA 2":
		sectionNameHindi = "अम्‍हेडा 2"
	case "AMHEDA":
		sectionNameHindi = "अम्‍हेडा"
	case "AMHEDI GAIR AB.":
		sectionNameHindi = "अम्‍हेडी गैर आं0"
	case "AMINABAD":
		sectionNameHindi = "अमीनाबाद"
	case "AMIPUR BEGA":
		sectionNameHindi = "अमीपुर बेगा"
	case "AMIPUR SUDHA":
		sectionNameHindi = "अमीपुर सुधा"
	case "AMIPUR URF DHARMNAGRI":
		sectionNameHindi = "अमीपुर उर्फ धर्मनगीर"
	case "AMIPUR URF NARYANPUR":
		sectionNameHindi = "अमीपुर उर्फ नरायनपुर"
	case "AMIRAPUR":
		sectionNameHindi = "अमीरपुर"
	case "AMIRPUR GANGU GAIR AB.":
		sectionNameHindi = "अमीरपुर गांगू गैर आं0"
	case "AMIRPUR GANGU":
		sectionNameHindi = "अमीरपुर गांगू"
	case "AMIRPUR MAIDAS GAIR ABAD":
		sectionNameHindi = "अमीरपुर माईदास गेर आबाद"
	case "AMRABAD":
		sectionNameHindi = "अमराबाद"
	case "AMRITSARI":
		sectionNameHindi = "अमृतसरी"
	case "ANANDKHERI":
		sectionNameHindi = "आनन्‍दखेडी"
	case "ANDKHEDA":
		sectionNameHindi = "अण्‍डखेडा"
	case "ANESA NANGLI":
		sectionNameHindi = "अनैसा नंगली"
	case "ANGUR VALI MASJID":
		sectionNameHindi = "अंगूर वाली मस्जिद"
	case "ANGURI BAG":
		sectionNameHindi = "अंगूरी बाग"
	case "ANKHERA":
		sectionNameHindi = "अन्‍डखेडा"
	case "ANOOPNAGAR":
		sectionNameHindi = "अनूपनगर"
	case "ANSARI PARK KATABAGH":
		sectionNameHindi = "अन्‍सारी पार्क कटाबाग"
	case "ANSARIYAN AN.":
		sectionNameHindi = "अन्‍सारियान आं०"
	case "ANSARIYAN D0 WARD -12":
		sectionNameHindi = "अन्सारियान द0 वार्ड-12"
	case "ANSARIYAN D0 WARD-12":
		sectionNameHindi = "अन्सारियान द0 वार्ड-12"
	case "ANSARIYAN P0 WARD -6":
		sectionNameHindi = "अन्सारियान प0 वार्ड-6"
	case "ANSARIYAN U0":
		sectionNameHindi = "अन्सारियान उ0"
	case "ANSARIYAN":
		sectionNameHindi = "अंसारियान"
	case "ANVALAGHAT":
		sectionNameHindi = "आंवला घाट"
	case "ANWA 01":
		sectionNameHindi = "अनवा 01"
	case "ANWA 02":
		sectionNameHindi = "अनवा 02"
	case "ANWARIYA TALIWABAD":
		sectionNameHindi = "अनवरिया तालिवाबाद"
	case "ANWARPUR ADIL GAIR ABAD":
		sectionNameHindi = "अनवरपुर आदिल गैर आबाद"
	case "ANWARPUR CHANDIKA":
		sectionNameHindi = "अनवरपुर चण्‍डि‍का"
	case "ANWARPUR CHATAR":
		sectionNameHindi = "अनवरपुर चतर"
	case "ANYA FAIKTARI ERIYA":
		sectionNameHindi = "अन्य फैक्टरी एरिया"
	case "ANYARI ORF ALINAGAR":
		sectionNameHindi = "अन्‍यारी उर्फ अलीनगर"
	case "AOURANGSHAHPUR GANGDHAR":
		sectionNameHindi = "औरंगशाहपुर गंगाधर"
	case "ARAFPUR KHAJURI":
		sectionNameHindi = "आरफपुर खजूरी"
	case "ARAGI BHASA":
		sectionNameHindi = "आराजी भैंसा"
	case "ARAZI GOPAL JOT":
		sectionNameHindi = "आराजी गोपाल जोत"
	case "ARGUPURA GAIR ABAD":
		sectionNameHindi = "अरगुपुरा गैर आबाद"
	case "ARIFPUR KUN‍DARKI GAIRA0":
		sectionNameHindi = "आरिफपुर कुन्‍दरकी गैरा0"
	case "ARIKHEDA":
		sectionNameHindi = "आरीखेड़ा"
	case "ARJANIPUR":
		sectionNameHindi = "अरजानीपुर"
	case "ASADGAR":
		sectionNameHindi = "असदगढ"
	case "ASADPUR DHAMROLI":
		sectionNameHindi = "असदपुर धमरौली"
	case "ASADULLAPUR AN0":
		sectionNameHindi = "असदुल्‍लापुर आं0"
	case "ASAFABAD BHAWAN URF GANVDI":
		sectionNameHindi = "आसफाबाद भवन उर्फ गांवडी"
	case "ASAFPURSAIDPEER":
		sectionNameHindi = "आसफपुरसैदपीर"
	case "ASALAT GANJ CHHOTI MANDI":
		sectionNameHindi = "असालत गंज छोटी मण्‍डी"
	case "ASALAT NAGAR (KALIJAPUR)":
		sectionNameHindi = "असालत नगर (कालीजपुर)"
	case "ASALAT NAGAR VAGHA":
		sectionNameHindi = "असालत नगर वघा"
	case "ASALATPUR MAJARABILARI":
		sectionNameHindi = "असालतपुर मजराबिलारी"
	case "ASALATPUR NEAR SHIV NAGAR":
		sectionNameHindi = "असालतपुर नि0 शिवनगर"
	case "ASALATPUR URF TOFAPUR":
		sectionNameHindi = "असालतपुर उर्फ तोफापुर"
	case "ASALATPUR":
		sectionNameHindi = "असालतपुर"
	case "ASALATPURA DEHRA ANSHIK":
		sectionNameHindi = "असालतपुरा डेहरा आंशिक"
	case "ASALATPURA MADHYA ANSHIK":
		sectionNameHindi = "असालतपुरा मध्‍य आंशिक"
	case "ASALATPURA PULIA ANSHIK":
		sectionNameHindi = "असालतपुरा पुलिया आंशिक"
	case "ASALATPURA UTTRI":
		sectionNameHindi = "असालतपुरा उत्‍तरी"
	case "ASALEMAPUR":
		sectionNameHindi = "असलेमपुर"
	case "ASDAPUR":
		sectionNameHindi = "असदपुर"
	case "ASDULLAPUR KALYAN":
		sectionNameHindi = "अस्‍दुल्‍लापुर कल्‍याण"
	case "ASDULLAPUR KHALSA":
		sectionNameHindi = "अस्‍दुल्‍ला पुरखालसा"
	case "ASDULLAPUR PRITHVI":
		sectionNameHindi = "असदुल्‍लापुर पृथवी"
	case "ASDULLAPUR":
		sectionNameHindi = "अस्‍दुल्‍लापुर"
	case "ASGARIPUR A0 H.S. NO. 1 TO 133":
		sectionNameHindi = "असगरीपुर आ0 मकान न0 1 सं 133 तक"
	case "ASGARIPUR A0 H.S. NO.134 TO END":
		sectionNameHindi = "असगरीपुर आ0 मकान न0 १३४ से अन्‍त तक"
	case "ASGARIPUR A0":
		sectionNameHindi = "असगरीपुर आं0"
	case "ASHIYANA-2":
		sectionNameHindi = "आशियाना-2"
	case "ASHIYANA.":
		sectionNameHindi = "आशियाना"
	case "ASHIYANA":
		sectionNameHindi = "आशियाना"
	case "ASHOKPUR":
		sectionNameHindi = "अशोकपुर"
	case "ASHRAFPUR":
		sectionNameHindi = "अशरफपुर"
	case "ASLAMPUR JHOJA":
		sectionNameHindi = "असलमपुर झोझा"
	case "ASLAMPURBHULLAN":
		sectionNameHindi = "असलमपुर भुल्‍लन"
	case "ASTABAL KAIMP MAY THANA GANJ":
		sectionNameHindi = "अस्तबल कैम्प मय थाना गंज"
	case "ATAI MOHALLA":
		sectionNameHindi = "अताई मौहल्‍ला"
	case "ATAI NAGAR":
		sectionNameHindi = "अताई नगर"
	case "ATAPUR URF KHATAPUR":
		sectionNameHindi = "अतापुर उर्फ खतापुर"
	case "ATHAI AHIR":
		sectionNameHindi = "अथाई अहीर"
	case "ATHAI JAMRUDDIN H.N. 1 TO 120":
		sectionNameHindi = "अथाई जमरूददीन आ0 म0नं0 १ से १२० तक"
	case "ATHAI JAMRUDDIN H.N. 121 TO END":
		sectionNameHindi = "अथाई जमरूददीन आ0 म0नं0 १२१ से अंत तक"
	case "ATHAI SHAIKH":
		sectionNameHindi = "अथाई शेख"
	case "ATHI":
		sectionNameHindi = "अथाई"
	case "ATRIYA":
		sectionNameHindi = "अटरिया"
	case "ATTA ALLANUR":
		sectionNameHindi = "अटटा अल्लानूर"
	case "AULIYAPUR":
		sectionNameHindi = "औलियापुर"
	case "AURANG TARA":
		sectionNameHindi = "ओरगपुर तारा"
	case "AURANGABAD AN0":
		sectionNameHindi = "औरंगाबाद आं0"
	case "AURANGABAD":
		sectionNameHindi = "औरंगाबाद"
	case "AURANGASHAHAPUR NARAYAN":
		sectionNameHindi = "औरंगशाहपुर नरायण"
	case "AURANGBAD SHAKURPUR":
		sectionNameHindi = "औरगाबाद शकूरपुर"
	case "AURANGJEBPUR CHANDA GAIR ABAD":
		sectionNameHindi = "औरंगजेबपुर चन्‍दा गैर आबाद"
	case "AURANGJEBPUR FAJIL":
		sectionNameHindi = "औरंगजेबपुर फाजिल"
	case "AURANGJEBPUR GULAL":
		sectionNameHindi = "औरंगजेबपुर गुलाल"
	case "AURANGJEBPUR MEHMUD URF MANGOLPURA":
		sectionNameHindi = "औरंगजेबपुर महमूद ऊर्फ मंगोलपुरा"
	case "AURANGJEBPUR SAMBHAL":
		sectionNameHindi = "औरंगजेबपुर सम्‍भल"
	case "AURANGJEBPURHARDAS":
		sectionNameHindi = "औरंगजेबपुर हरदास"
	case "AURANGJEBPURSAHLI":
		sectionNameHindi = "ओरेगजेबपुर शाहली"
	case "AURANGPUR BASANTA AN0":
		sectionNameHindi = "औरंगपुर बसंता आं0"
	case "AURANGPUR BASANTA":
		sectionNameHindi = "औरंगपुर बसंता"
	case "AURANGPUR BHIKKU":
		sectionNameHindi = "औरंगपुर भिक्‍कू"
	case "AURANGPUR BIBI":
		sectionNameHindi = "औरगपुर बीबी"
	case "AURANGPUR FATTA":
		sectionNameHindi = "औरंगपुर फत्‍ता"
	case "AURANGPUR FATTE KHAN":
		sectionNameHindi = "औरंगपुर फत्‍ते खां"
	case "AURANGPUR HAZI":
		sectionNameHindi = "औरगपुर हाजी"
	case "AURANGPUR HIRDAY":
		sectionNameHindi = "औरंगपुर हृदय"
	case "AURANGSHAHPUR BAUVARI":
		sectionNameHindi = "औरंगशाहपुर बनवारी"
	case "AURANPUR NANDLAL":
		sectionNameHindi = "औरंगपुर नंदलाल"
	case "AVANTIKA COLONY":
		sectionNameHindi = "अवंतिका कालौनी"
	case "AVAS VIKAS GANGAPUR":
		sectionNameHindi = "आवास विकास गंगापुर"
	case "AVID NAGAR TEELA ANSHIK":
		sectionNameHindi = "आविद नगर टीला आंशिक"
	case "AVID NAGAR TEELA":
		sectionNameHindi = "आविद नगर टीला"
	case "AWADI RAILWAY STATION":
		sectionNameHindi = "आबादी रेलवे स्‍टेशन"
	case "AWAS VIKAS COLONY":
		sectionNameHindi = "आवास विकास कालोनी"
	case "AZAMGAR URF RATANGAR A.":
		sectionNameHindi = "आजमगढ उर्फ रतनगढ आ०"
	case "AZAMPUR URF ADDOPUR A.":
		sectionNameHindi = "आजमपुर उर्फ आदोपुर आ०"
	case "AZAMPUR URF ADDOPUR A":
		sectionNameHindi = "आजमपुर उर्फ आदोपुर आ०"
	case "AZEEM NAGAR":
		sectionNameHindi = "अजीमनगर"
	case "AZEEMNAGAR":
		sectionNameHindi = "अजीमनगर"
	case "AZMABAD":
		sectionNameHindi = "आजमाबाद"
	case "BAAJAR":
		sectionNameHindi = "बाजार"
	case "BABANPURA":
		sectionNameHindi = "बबनपुरा"
	case "BABARPUR":
		sectionNameHindi = "बाबरपुर"
	case "BABLI":
		sectionNameHindi = "बाबली"
	case "BABOORA":
		sectionNameHindi = "बबूरा"
	case "BABULAL KA FATAK PATPAT SARAI":
		sectionNameHindi = "बाबूलाल का फाटक पटपट सराय"
	case "BABUPURA GHOSIPURA":
		sectionNameHindi = "बाबूपुरा घोसीपुरा"
	case "BACHHAL BHUD":
		sectionNameHindi = "बाछल भूड़"
	case "BACHHIYAI":
		sectionNameHindi = "बछियाई"
	case "BADA SHAH SHAFA ANSHIK":
		sectionNameHindi = "बाडा शाह शफा आंशिक"
	case "BADALI-1":
		sectionNameHindi = "बादली-1"
	case "BADALI-2":
		sectionNameHindi = "बादली-2"
	case "BADASHAHAPUR LAXHMI SEN":
		sectionNameHindi = "बादशाहपुर लक्ष्‍मी सैन"
	case "BADASHAPUR SHAH MAUHM‍MAD":
		sectionNameHindi = "बादशाहपुर शाह मौहमम्‍मद"
	case "BADAVAN AN.":
		sectionNameHindi = "बाडवान आं0"
	case "BADHERA":
		sectionNameHindi = "बढेरा"
	case "BADHPURA MAHESHPUR KHEM":
		sectionNameHindi = "बढपुरा महेशपुर खेम"
	case "BADHPURA SHARKI":
		sectionNameHindi = "बढपुरा शर्की"
	case "BADI MANDI":
		sectionNameHindi = "बडी मण्‍डी"
	case "BADIWALA":
		sectionNameHindi = "बाडीवाला"
	case "BADODE FATEEHULLAPUR":
		sectionNameHindi = "बडौदी फत्‍तेहउल्‍लापुर"
	case "BADPURA SHUMALI":
		sectionNameHindi = "बढपुरा शमाली"
	case "BADPURA":
		sectionNameHindi = "बढपुरा"
	case "BADSHAHPUR TARIKAM":
		sectionNameHindi = "बादशाहपुर तरीकम"
	case "BADSHAHPUR":
		sectionNameHindi = "बादशाहपुर"
	case "BADSHAPUR MAIDAS":
		sectionNameHindi = "बादशाहपुर माईदास"
	case "BADSHAPUR":
		sectionNameHindi = "बादशाहपुर"
	case "BAG PUKHTA":
		sectionNameHindi = "बाग पुख्ता"
	case "BAGADKHA-1":
		sectionNameHindi = "बगडखा-1"
	case "BAGADKHA-2":
		sectionNameHindi = "बगडखा-2"
	case "BAGADKHA-3":
		sectionNameHindi = "बगडखा-3"
	case "BAGADPUR":
		sectionNameHindi = "बागडरपुर"
	case "BAGAMPUR":
		sectionNameHindi = "बेगमपुर"
	case "BAGARPUR":
		sectionNameHindi = "बागडपुर"
	case "BAGH GULABRAI UTTARI":
		sectionNameHindi = "बाग गुलाबराय उत्‍तरी"
	case "BAGHALA":
		sectionNameHindi = "बघाला"
	case "BAGI-1":
		sectionNameHindi = "बगी-1"
	case "BAGI-2":
		sectionNameHindi = "बगी-2"
	case "BAGI":
		sectionNameHindi = "बगी"
	case "BAGIA MAKBARA":
		sectionNameHindi = "बगिया मकबरा"
	case "BAGIA":
		sectionNameHindi = "बगिया"
	case "BAGIAWALA":
		sectionNameHindi = "बगियावाला"
	case "BAGICHA AIMNA":
		sectionNameHindi = "बगीचा ऐमना"
	case "BAGICHA CHHOTE MIAN 2":
		sectionNameHindi = "बगीचा छोटे मियां 2"
	case "BAGICHA CHHOTE MIAN-1":
		sectionNameHindi = "बगीचा छोटे मियां-1"
	case "BAGICHA GAJI MUJFAFAR KHA":
		sectionNameHindi = "बगीचा गाजी मुजफफर खा"
	case "BAGICHA JOKHIRAM":
		sectionNameHindi = "बगीचा जोखीराम"
	case "BAGIYA SAGAR":
		sectionNameHindi = "बगिया सागर"
	case "BAGRAUA":
		sectionNameHindi = "बगरौआ"
	case "BAGRAWWA-1":
		sectionNameHindi = "बगरव्‍वा-1"
	case "BAGRAWWA-2":
		sectionNameHindi = "बगरव्‍वा-2"
	case "BAGWADA A.":
		sectionNameHindi = "बगबाडा आ0"
	case "BAGWADA":
		sectionNameHindi = "बगबाडा"
	case "BAGWAN CHAPEGARAN":
		sectionNameHindi = "बागवान छापेग्रान"
	case "BAGWAN NAKTA":
		sectionNameHindi = "बागवान नक्‍टा"
	case "BAGWANAN":
		sectionNameHindi = "बागवानान"
	case "BAGWANTPUR":
		sectionNameHindi = "भगवन्‍तपुर"
	case "BAHADA GAIR AB.AD":
		sectionNameHindi = "बहेडा गैर आं0"
	case "BAHADAR NAGAR":
		sectionNameHindi = "बहादरनगर"
	case "BAHADARGANJ":
		sectionNameHindi = "बहादुर गंज"
	case "BAHADARPUR BUJURG":
		sectionNameHindi = "बहादुरपुर बुजुर्ग"
	case "BAHADARPUR JATT":
		sectionNameHindi = "बहादुरपुर जट"
	case "BAHADARPUR KHADDAR":
		sectionNameHindi = "बहादरपुर खददर"
	case "BAHADARPUR SHARFUDIN HUSAIN AN.":
		sectionNameHindi = "बहादरपुर शर्फुदीन हुसैन आ०"
	case "BAHADARPUR":
		sectionNameHindi = "बहादरपुर"
	case "BAHADUR GANJ":
		sectionNameHindi = "वहादुर गंज"
	case "BAHADUR NAGAR":
		sectionNameHindi = "बहादुर नगर"
	case "BAHADURAPUR RAJAPUT":
		sectionNameHindi = "बहादुरपुर राजपूत"
	case "BAHADURGAJJU":
		sectionNameHindi = "बहादुरगज्‍जू"
	case "BAHADURGANJ":
		sectionNameHindi = "बहादुरगंज"
	case "BAHADURPUR PATTI":
		sectionNameHindi = "बहादुरपुर पट्टी"
	case "BAHAPUR":
		sectionNameHindi = "बहापुर"
	case "BAHAPURA-2":
		sectionNameHindi = "बहपुरा-2"
	case "BAHAPURI":
		sectionNameHindi = "बहपुरी"
	case "BAHATA SARTHAL":
		sectionNameHindi = "बहटा सरथल"
	case "BAHEDI BRAHAMNAN":
		sectionNameHindi = "बहेडी ब्रहमनान"
	case "BAHEDI":
		sectionNameHindi = "बहेड़ी"
	case "BAHLOLPUR":
		sectionNameHindi = "बहलोलपुर"
	case "BAHMANSORA":
		sectionNameHindi = "बहमनशोरा"
	case "BAHORANPUR KALA ANSHIK":
		sectionNameHindi = "बहोरनपुर कला आंशिक"
	case "BAHORANPUR NARAULI":
		sectionNameHindi = "बहोरनपुर नरौली"
	case "BAHPUR GANGAPUR":
		sectionNameHindi = "वाहपुर गंगापुर"
	case "BAHTOLA":
		sectionNameHindi = "बहटौला"
	case "BAIBALPUR":
		sectionNameHindi = "बैबलपुर"
	case "BAIBALWALA":
		sectionNameHindi = "बैबलवाला"
	case "BAIJANATHAPUR":
		sectionNameHindi = "बैजनाथपुर"
	case "BAINDU KHEDA":
		sectionNameHindi = "बैन्‍दू खेडा"
	case "BAIRAM NAGAR":
		sectionNameHindi = "बैरमनगर"
	case "BAIRAMNAGAR":
		sectionNameHindi = "बैरमनगर"
	case "BAIRAMPUR KHAJURI":
		sectionNameHindi = "बैरमपुर खजूरी"
	case "BAIRMABAD GADI":
		sectionNameHindi = "बैरमाबाद गढी"
	case "BAIRMAPUR":
		sectionNameHindi = "बैरमपुर"
	case "BAIZANI":
		sectionNameHindi = "बैजनी"
	case "BAJAR ABDULLA GANJ":
		sectionNameHindi = "बाजार अब्दुल्ला गंज"
	case "BAJAR KALAN-1":
		sectionNameHindi = "बाजार कलां-1"
	case "BAJAR KALAN-2":
		sectionNameHindi = "बाजार कलां-2"
	case "BAJAR KALAN-3":
		sectionNameHindi = "बाजार कलां-3"
	case "BAJAR KALAN-4":
		sectionNameHindi = "बाजार कलां-4"
	case "BAJAR KALAN-5":
		sectionNameHindi = "बाजार कलां-5"
	case "BAJAR KALAN-6":
		sectionNameHindi = "बाजार कलां-6"
	case "BAJAR KALAN-7":
		sectionNameHindi = "बाजार कलां-7"
	case "BAJAR KALAN-8":
		sectionNameHindi = "बाजार कलां-8"
	case "BAJAR KALAN":
		sectionNameHindi = "बजार कलां"
	case "BAJAR NASARULLA KHAN":
		sectionNameHindi = "बाजार नसरूल्ला खां"
	case "BAJAR PASCHIMI ANSHIK":
		sectionNameHindi = "बाजार पश्चिमी(आं)"
	case "BAJAR PASCHIMI WARD NO-1":
		sectionNameHindi = "बाजार पश्चिमी वार्ड नं0 -1"
	case "BAJAR PASHCHAMI WARD NO-1":
		sectionNameHindi = "बाजार पं0 वार्ड नं0-1"
	case "BAJAR PURVI":
		sectionNameHindi = "बाजार पूर्वी"
	case "BAJAR SAFDARGANJ MAY PAN DARIBA":
		sectionNameHindi = "बाजार सफदरगंज मय पान दरीबा"
	case "BAJARIYA FATEH ALI KHA 2":
		sectionNameHindi = "बजरिया फतेह अली खा 2"
	case "BAJARIYA FATEH ALI KHA-1":
		sectionNameHindi = "बजरिया फतेह अली खा-1"
	case "BAJARIYA HIMMAT KHAN":
		sectionNameHindi = "बजरिया हिम्मत खां"
	case "BAJARIYA JHABBU KHAN MAY SILEYAN-1":
		sectionNameHindi = "बजरिया झब्बू खां मय सिलेयान-1"
	case "BAJARIYA JHABBU KHAN MAY SILEYAN-2":
		sectionNameHindi = "बजरिया झब्बू खां मय सिलेयान -2"
	case "BAJARIYA KADHU":
		sectionNameHindi = "बजरिया कढू"
	case "BAJARIYA KHANASAMA":
		sectionNameHindi = "बजरिया खानसामा"
	case "BAJARIYA MULLA JARIF":
		sectionNameHindi = "बजरिया मुल्ला जरीफ"
	case "BAJHANPUR-01":
		sectionNameHindi = "भजनपुर -01"
	case "BAJHANPUR-02":
		sectionNameHindi = "भजनपुर -0२"
	case "BAJHRABAD URF KHATAI":
		sectionNameHindi = "बाखराबाद उर्फ खटाई"
	case "BAJIDPUR":
		sectionNameHindi = "बाजिदपुर"
	case "BAJIRAPUR JANGIR AN.":
		sectionNameHindi = "बजीरपुर जगीर आं0"
	case "BAJIRAPUR URF MADHIYA":
		sectionNameHindi = "बजीरपुर उर्फ मडैयो"
	case "BAJODI TOLA -1":
		sectionNameHindi = "बाजोड़ी टोला -1"
	case "BAJODI TOLA -2":
		sectionNameHindi = "बाजोड़ी टोला -2"
	case "BAJODI TOLA- 4":
		sectionNameHindi = "बाजोड़ी टोला- 4"
	case "BAJODI TOLA- 5":
		sectionNameHindi = "बाजोड़ी टोला- 5"
	case "BAJODI TOLA-3":
		sectionNameHindi = "बाजोड़ी टोला-3"
	case "BAKAINIYACHANDAPUR":
		sectionNameHindi = "बकैनियाचांदपुर"
	case "BAKALAN":
		sectionNameHindi = "बकालान"
	case "BAKAR KASSABAN":
		sectionNameHindi = "बकर कसावान"
	case "BAKAR NANGLA":
		sectionNameHindi = "बाकर नंगला"
	case "BAKARNAGAR":
		sectionNameHindi = "बाकरनगर"
	case "BAKARPUR ATYAN":
		sectionNameHindi = "बाकरपुर अटायन"
	case "BAKARPUR JAIRAM":
		sectionNameHindi = "बाकरपुर जयराम"
	case "BAKARPUR":
		sectionNameHindi = "बाकरपुर"
	case "BAKENIYA BHAT":
		sectionNameHindi = "बकैनियां भाट"
	case "BAKENIYA":
		sectionNameHindi = "बकैनिया"
	case "BAKHSHANAPUR AN.":
		sectionNameHindi = "बक्‍शनपुर आं0"
	case "BAKNOARI":
		sectionNameHindi = "बकनौरी"
	case "BAKRI KA AHATA DEHRA ASALATPURA":
		sectionNameHindi = "बकरी का अहाता डेहरा असालतपुरा"
	case "BALAKRAM":
		sectionNameHindi = "बालकराम"
	case "BALAPUR":
		sectionNameHindi = "बालापुर"
	case "BALBHADRAPUR":
		sectionNameHindi = "बलभद्रपुर"
	case "BALDEVPURI ANSHIK":
		sectionNameHindi = "बलदेवपुरी आंशिक"
	case "BALDIYA":
		sectionNameHindi = "बल्दिया"
	case "BALIPUR GAWRI":
		sectionNameHindi = "वलीपुर गावडी"
	case "BALIYANAGLI":
		sectionNameHindi = "बलिया नंगली"
	case "BALKHEDA":
		sectionNameHindi = "बलखेडा"
	case "BALKISHANPUR GAIR AB.":
		sectionNameHindi = "बालकिशनपुर गैर आं0"
	case "BALKISHNAPUR":
		sectionNameHindi = "बालकिशनपुर"
	case "BALLA NANGLA":
		sectionNameHindi = "बल्‍ला नंगला"
	case "BALLAM STREET":
		sectionNameHindi = "बल्‍लम स्‍ट्रीट"
	case "BALMIKI MOHALLA CHANDRA NAGAR":
		sectionNameHindi = "वाल्‍मीकि मौहल्‍ला चन्‍द्र नगर"
	case "BALUPURA":
		sectionNameHindi = "बलूपुरा"
	case "BAMANAULI":
		sectionNameHindi = "बमनौली"
	case "BAMANIYAPATTI":
		sectionNameHindi = "बमनिया पटटी वम"
	case "BAMANPURA-1":
		sectionNameHindi = "वमनपुरा-1"
	case "BAMANPURA-2":
		sectionNameHindi = "वमनपुरा-2"
	case "BAMANPURI":
		sectionNameHindi = "बमनपुरी"
	case "BAMNA":
		sectionNameHindi = "बमना"
	case "BAMNAPURI -1":
		sectionNameHindi = "बमनपुरी- 1"
	case "BAMNAPURI -2":
		sectionNameHindi = "बमनपुरी- 2"
	case "BAMNAULA,":
		sectionNameHindi = "बमनौला,"
	case "BAMNOLA":
		sectionNameHindi = "बमनौला"
	case "BAMNOLI":
		sectionNameHindi = "बमनौली"
	case "BAMNOOLI":
		sectionNameHindi = "बमनौली"
	case "BANDAR 01":
		sectionNameHindi = "बन्‍दार 01"
	case "BANDAR 02":
		sectionNameHindi = "वन्‍दार 02"
	case "BANDUKACHIYAN":
		sectionNameHindi = "बन्दूकचियान"
	case "BANGADPUR KISHNA GAIR. AN.AD":
		sectionNameHindi = "बांगडपुर किशना गै0आ0"
	case "BANGALA AJAD KHA":
		sectionNameHindi = "बंगला आजाद खा"
	case "BANGALA KASAM ALI KHAN":
		sectionNameHindi = "बंगला कासम अली खां"
	case "BANGLA GAON ANSHIK":
		sectionNameHindi = "बंगला गॉव आंशिक"
	case "BANIRAMDASS":
		sectionNameHindi = "बनीरामदास"
	case "BANJARAN":
		sectionNameHindi = "बंजारान"
	case "BANK COLONY SIHORA GOVIND":
		sectionNameHindi = "बैंक कालौनी सिहोरा गोविन्‍द"
	case "BANKAVALA":
		sectionNameHindi = "बंकावाला"
	case "BANKHALA":
		sectionNameHindi = "बनखला"
	case "BANKRABAD":
		sectionNameHindi = "बांकराबाद"
	case "BANSFODAN AN0":
		sectionNameHindi = "बांसफोड़ान आं0"
	case "BANSKHERA":
		sectionNameHindi = "बांसखेडा"
	case "BANSNAGLI":
		sectionNameHindi = "बांसनगली"
	case "BANWARIPUR":
		sectionNameHindi = "बनवारीपुर"
	case "BANWATA GANJ":
		sectionNameHindi = "बनवटा गंज"
	case "BARA KHAS":
		sectionNameHindi = "बराखास"
	case "BARA NIKAT GAJEJA":
		sectionNameHindi = "बरा नि0 गजेजा"
	case "BARADARI":
		sectionNameHindi = "बारादरी"
	case "BARAITHA KHIJRAPUR":
		sectionNameHindi = "बरैठा खिजरपुर"
	case "BARAMPUR AN0":
		sectionNameHindi = "बरमपुर आं0"
	case "BARAULI RUSTAMAPUR":
		sectionNameHindi = "बरौली रुस्तमपुर"
	case "BARAVARA KHAS ANSHIK":
		sectionNameHindi = "बरवारा खास आंशिक"
	case "BARBALAN":
		sectionNameHindi = "बरबालान"
	case "BARDA GAUN 01":
		sectionNameHindi = "बडा गांव 01"
	case "BARDA GAUN 02":
		sectionNameHindi = "बडा गांव 02"
	case "BARDA GAUN 03":
		sectionNameHindi = "बडा गांव 03"
	case "BARHAPUR":
		sectionNameHindi = "बढापुर"
	case "BARIPUR":
		sectionNameHindi = "बारीपुर"
	case "BARKALA":
		sectionNameHindi = "बरकला"
	case "BARKAPUR GAIR AB.AD":
		sectionNameHindi = "बरकापुर गैर आं0"
	case "BARKATA":
		sectionNameHindi = "बरकटा"
	case "BARKATPUR GARHI":
		sectionNameHindi = "बरकातपुर गढी"
	case "BARKATPUR":
		sectionNameHindi = "बरकातपुर"
	case "BARKHADA":
		sectionNameHindi = "बेरखेडा"
	case "BARKHEDA BASANT PUR ORF DYANATHPUR":
		sectionNameHindi = "बरखेडा बसन्‍तपुर उर्फ दयानाथपुर"
	case "BARKHERA":
		sectionNameHindi = "बरखेडा"
	case "BARKHURDARPUR GOPAL":
		sectionNameHindi = "बरखुरदारपुर गोपाल"
	case "BARU BHUDA":
		sectionNameHindi = "बारू भूडा"
	case "BARUA -01":
		sectionNameHindi = "बैरूआ -01"
	case "BARUA -02":
		sectionNameHindi = "बैरूआ -02"
	case "BARUKI":
		sectionNameHindi = "बरूकी"
	case "BARWALAN ANSHIK":
		sectionNameHindi = "बरवालान आंशिक"
	case "BARWALAN DAKSHINI ANSHIK":
		sectionNameHindi = "बरवालान दक्षिणी आंशिक"
	case "BARWALAN DAKSHINI":
		sectionNameHindi = "बरवालान दक्षिणी"
	case "BARWALAN UTTRI ANSHIK":
		sectionNameHindi = "बरवालान उत्‍तरी आंशिक"
	case "BARWALAN UTTRI CHAMUNDA GALI ANSHIK":
		sectionNameHindi = "बरवलान उत्‍तरी चमुण्‍डा गली आंशिक"
	case "BARWARA MAJHRA ANSHIK":
		sectionNameHindi = "बरवारा मझरा आंशिक"
	case "BASABANPUR":
		sectionNameHindi = "बसावनपुर"
	case "BASADA":
		sectionNameHindi = "बसेडा"
	case "BASADI":
		sectionNameHindi = "बसेडी"
	case "BASANTPUR MAFI G.A.":
		sectionNameHindi = "बसन्‍तपुर माफी गै0आ0"
	case "BASANTPUR MAFI":
		sectionNameHindi = "बसन्‍तपुर माफी"
	case "BASANTPUR RAMRAI":
		sectionNameHindi = "बसन्‍तपुर रामराय"
	case "BASANTPUR":
		sectionNameHindi = "बसन्‍तपुर"
	case "BASAWANPUR":
		sectionNameHindi = "बसावनपुर"
	case "BASEDA KHADAR":
		sectionNameHindi = "बसेडा खादर"
	case "BASEDA KHEMCHAND":
		sectionNameHindi = "बसेडा खेमचन्‍द"
	case "BASEDA KHURD":
		sectionNameHindi = "बसेडा खुर्द"
	case "BASEDA KOTT":
		sectionNameHindi = "बसेडा कौट"
	case "BASEDA NARAYAN":
		sectionNameHindi = "बसेडा नरायन"
	case "BASEDA UBAR":
		sectionNameHindi = "बसेडा उबार"
	case "BASEDA":
		sectionNameHindi = "बसेडा"
	case "BASERA PATTI GAIRA.":
		sectionNameHindi = "बसेरा पटटी गैरा0"
	case "BASERAKHAS":
		sectionNameHindi = "बसेरा खास"
	case "BASERI":
		sectionNameHindi = "बसेरी"
	case "BASHIRPU":
		sectionNameHindi = "बशीरपुर"
	case "BASIRUHASI":
		sectionNameHindi = "वसीरूहासी"
	case "BASODA":
		sectionNameHindi = "भसौडा"
	case "BASTA A.":
		sectionNameHindi = "बास्‍टा आ०"
	case "BASUWALA":
		sectionNameHindi = "बासूवाला"
	case "BATAPURA":
		sectionNameHindi = "बाटपुरा"
	case "BATHUAKHERA":
		sectionNameHindi = "बथुआखेडा"
	case "BAWAN SARAI":
		sectionNameHindi = "बाबनसराय"
	case "BAWARIYAN":
		sectionNameHindi = "बावरियान"
	case "BAZABALA-1":
		sectionNameHindi = "वजावाला-1"
	case "BAZABALA":
		sectionNameHindi = "वजावाला-2"
	case "BAZANA-2":
		sectionNameHindi = "बैजना-2"
	case "BAZAR DIWAN DAKSHINI":
		sectionNameHindi = "बाजार दिवान दक्षिणी"
	case "BAZAR KALA":
		sectionNameHindi = "बाजार कला"
	case "BAZAR MANGAL":
		sectionNameHindi = "बाजार मंगल"
	case "BAZAR MUFTI ANSHIK":
		sectionNameHindi = "बाजार मुफती आंशिक"
	case "BAZAR SHAMBA":
		sectionNameHindi = "बाजार शम्‍भा"
	case "BAZAR":
		sectionNameHindi = "बाजार"
	case "BAZDARAN":
		sectionNameHindi = "बाजदारान"
	case "BEBIPURA A.":
		sectionNameHindi = "बीबीपुरा आ०"
	case "BEBIPURA":
		sectionNameHindi = "बीबीपुरा आ०"
	case "BEDA,":
		sectionNameHindi = "बेडा,"
	case "BEDAN":
		sectionNameHindi = "बेदान"
	case "BEDARBAKTHPURJADO":
		sectionNameHindi = "बेदाखख्‍तपुरजादों"
	case "BEDPUR":
		sectionNameHindi = "वेदपुर"
	case "BEEJNA ANSHIK":
		sectionNameHindi = "बीजना आंशिक"
	case "BEERAMPUR":
		sectionNameHindi = "वीरमपुर"
	case "BEGAM SARIA":
		sectionNameHindi = "बैगम सराय"
	case "BEGAMPUR BHONAWALA":
		sectionNameHindi = "बेगमपुर भोनावाला"
	case "BEGAMPUR CHAIMAL":
		sectionNameHindi = "बेगमपुर चायमल"
	case "BEGAMPUR HARREY":
		sectionNameHindi = "बेगमपुर हर्रे"
	case "BEGAMPUR RUPCHAND":
		sectionNameHindi = "बेगमपुर रूपचन्‍द"
	case "BEGAMPUR SHADI AN0":
		sectionNameHindi = "बेगमपुर शादी आं०"
	case "BEGAMPUR":
		sectionNameHindi = "बेगमपुर"
	case "BEGMABAD NIKAT ASHOKPUR":
		sectionNameHindi = "बेगमाबाद नि0 अशोकपुर"
	case "BEGMABAD NIKAT KRIMCHA":
		sectionNameHindi = "बेगमाबाद नि0 क्रमचा"
	case "BEGMABAD":
		sectionNameHindi = "बेगमाबाद"
	case "BEGRAJPUR":
		sectionNameHindi = "बेगराजपुर"
	case "BEGUMPUR":
		sectionNameHindi = "बेगमपुर"
	case "BEHADA CHOHAN":
		sectionNameHindi = "बेहडा चौहान"
	case "BEHPURA- 1":
		sectionNameHindi = "बहपुरा-1"
	case "BEHTA":
		sectionNameHindi = "बेहटा"
	case "BEHTRA":
		sectionNameHindi = "बेहटरा"
	case "BELADARAN":
		sectionNameHindi = "बेलदारान"
	case "BELDARAN":
		sectionNameHindi = "बेलदरान"
	case "BELWADA":
		sectionNameHindi = "बेलवाडा"
	case "BENIPUR":
		sectionNameHindi = "बेनीपुर"
	case "BER":
		sectionNameHindi = "बेर"
	case "BERAKHEDA TANDA AN.":
		sectionNameHindi = "बेरखेडा टांडा आं0"
	case "BERAMPUR":
		sectionNameHindi = "बैरमपुर"
	case "BERIYAN AKAB MAHAL SARAY 1":
		sectionNameHindi = "बेरियान अकब महल सराय 1"
	case "BERIYAN AKAB MAHAL SARAY 2":
		sectionNameHindi = "बेरियान अकब महल सराय 2"
	case "BERKHEDA CHAK":
		sectionNameHindi = "बेरखेडा चक"
	case "BERKHEDA":
		sectionNameHindi = "बेरखेडा"
	case "BERKHEDI FEJABAD":
		sectionNameHindi = "बेरखेडी फैजाबाद"
	case "BERKHERA CHOUHAN":
		sectionNameHindi = "बेरखेडा चौहान"
	case "BESRA-2":
		sectionNameHindi = "बीसरा-2"
	case "BHABANIPUR":
		sectionNameHindi = "भवानीपुर"
	case "BHADARPUR KHURD":
		sectionNameHindi = "बहादरपुर खुर्द"
	case "BHADASNA":
		sectionNameHindi = "भदासना"
	case "BHADOLA":
		sectionNameHindi = "भदौला"
	case "BHADORA DEHAT AMBEDKAR NAGAR":
		sectionNameHindi = "भदौरा देहात अम्‍बेडकर नगर"
	case "BHADORA DEHAT ANSHIK":
		sectionNameHindi = "भदौरा देहात आंशिक"
	case "BHAGATPUR RATAN":
		sectionNameHindi = "भगतपुर रतन"
	case "BHAGBANTPUR":
		sectionNameHindi = "भगवन्‍तपुर"
	case "BHAGIJOT":
		sectionNameHindi = "भागीजोत"
	case "BHAGIN":
		sectionNameHindi = "भागैन"
	case "BHAGIYAVALA":
		sectionNameHindi = "भगियावाला"
	case "BHAGOORA":
		sectionNameHindi = "भगौडा"
	case "BHAGTAPUR TANDA":
		sectionNameHindi = "भगतपुर टाण्डा"
	case "BHAGUWALA AN0":
		sectionNameHindi = "भागूवाला आं0"
	case "BHAGUWALA":
		sectionNameHindi = "भागूवाला"
	case "BHAGWAN WALA":
		sectionNameHindi = "भगवान वाला"
	case "BHAGWANPUR HN0 1 TO 125":
		sectionNameHindi = "भगवानपुर म0स0 १ से १२५ तक"
	case "BHAGWANPUR HN0 126 TO END":
		sectionNameHindi = "भगवानपुर म0स0 १२६ से अंत तक"
	case "BHAGWANPUR RAINI":
		sectionNameHindi = "भगवानपुर रैनी"
	case "BHAGWANPUR":
		sectionNameHindi = "भगवानपुर"
	case "BHAGWANPURPRATAP":
		sectionNameHindi = "भगवानपुर प्रताप"
	case "BHAGWANT NAGAR -2":
		sectionNameHindi = "भगवन्‍तनगर 2"
	case "BHAGWANT NAGAR-1":
		sectionNameHindi = "भगवन्‍तनगर 1"
	case "BHAGWATIPUR":
		sectionNameHindi = "भगवतीपुर"
	case "BHAISA":
		sectionNameHindi = "भैंसा"
	case "BHAISAKHANA":
		sectionNameHindi = "भैसखाना"
	case "BHAISIYA ANSHIK":
		sectionNameHindi = "भैसिया आंशिक"
	case "BHAISIYA":
		sectionNameHindi = "भैसिया"
	case "BHAISOD":
		sectionNameHindi = "भैसोड़"
	case "BHAIYANGALA":
		sectionNameHindi = "भैयानगला"
	case "BHAJANPURI":
		sectionNameHindi = "भजनपुरी"
	case "BHAJDAWALAKHALSA":
		sectionNameHindi = "भजडावाला खालसा"
	case "BHAJDI SARAY":
		sectionNameHindi = "भजडी सराय"
	case "BHAMROA-1":
		sectionNameHindi = "भमरौआ -1"
	case "BHAMROA-2":
		sectionNameHindi = "भमरौआ -2"
	case "BHAMROA-3":
		sectionNameHindi = "भमरौआ -3"
	case "BHANAPUR GAJAROLA":
		sectionNameHindi = "भानपुर गजरोला"
	case "BHANDARA":
		sectionNameHindi = "भांडरा"
	case "BHANDPURA-1":
		sectionNameHindi = "भण्‍डपुरा-1"
	case "BHANDPURA-2":
		sectionNameHindi = "भण्‍डपुरा-2"
	case "BHANDRI AANSHIK":
		sectionNameHindi = "भांडरी आंशिक"
	case "BHANERA":
		sectionNameHindi = "भनेडा"
	case "BHARAIKI":
		sectionNameHindi = "भरैकी"
	case "BHARAIRA":
		sectionNameHindi = "भरैरा"
	case "BHARAKHAIDI URF DUGRU":
		sectionNameHindi = "भाराखेडी उर्फ डुगरी"
	case "BHARAKHERI EMMA":
		sectionNameHindi = "भारा खेडी इम्‍मा"
	case "BHARATAVALA":
		sectionNameHindi = "भरतावाला"
	case "BHARATPUR":
		sectionNameHindi = "भरतपुर"
	case "BHARKARA URF AZZAMNAGAR":
		sectionNameHindi = "बेरखेडा उर्फ आजमनगर"
	case "BHARUKA":
		sectionNameHindi = "भरूका"
	case "BHATAN AN":
		sectionNameHindi = "भाटान"
	case "BHATAN":
		sectionNameHindi = "भाटान"
	case "BHATAPURAKUNDARKI":
		sectionNameHindi = "भटपुरा कुन्दरकी"
	case "BHATAWALI ANSHIK":
		sectionNameHindi = "भटावली आंशिक"
	case "BHATGAVAN":
		sectionNameHindi = "भतगवां"
	case "BHATIKHEDA":
		sectionNameHindi = "भाटीखेडा"
	case "BHATOLI":
		sectionNameHindi = "भटोली"
	case "BHATPURA CHAKRPAN":
		sectionNameHindi = "भटपुरा चक्रपान"
	case "BHATPURA TARAN-1":
		sectionNameHindi = "भटपुरा तारन-1"
	case "BHATPURA TARAN-2":
		sectionNameHindi = "भटपुरा तारन-2"
	case "BHATPURA":
		sectionNameHindi = "भटपुरा"
	case "BHATT PURA":
		sectionNameHindi = "भटपुरा"
	case "BHATTI STREET":
		sectionNameHindi = "भटटी स्‍ट्रीट"
	case "BHATTI TOLA-1":
		sectionNameHindi = "भटटी टोला-1"
	case "BHATTI TOLA-2":
		sectionNameHindi = "भटटी टोला-2"
	case "BHATTI TOLA-3":
		sectionNameHindi = "भटटी टोला-3"
	case "BHATYANA KHUSHHALPUR":
		sectionNameHindi = "भटियाना खुशहालपुर"
	case "BHAUKHEDA GAIR ABAD":
		sectionNameHindi = "भाउखेडा गैर आं0"
	case "BHAUPURA":
		sectionNameHindi = "भाऊपुरा"
	case "BHAVANIPUR GANGRAIN":
		sectionNameHindi = "भवानीपुर गगरैंन"
	case "BHAVANIPUR":
		sectionNameHindi = "भवानीपुर"
	case "BHAVANPUR KALA":
		sectionNameHindi = "भवनपुर कला"
	case "BHAWAN":
		sectionNameHindi = "भवन"
	case "BHAWANIPUR AN0":
		sectionNameHindi = "भवानीपुर आ0"
	case "BHAWANIPUR GADDO":
		sectionNameHindi = "भवानीपुर गददो"
	case "BHAWANIPUR KALIYA":
		sectionNameHindi = "भवानीपुर कलिया"
	case "BHAWANIPUR TARKOULA":
		sectionNameHindi = "भवानीपुर तरकौला"
	case "BHAWANIPUR":
		sectionNameHindi = "भवानीपुर"
	case "BHAWARKA-1":
		sectionNameHindi = "भंवरका-1"
	case "BHAWARKA-2":
		sectionNameHindi = "भंवरका-2"
	case "BHAWARKI":
		sectionNameHindi = "भंवरकी"
	case "BHAYAPUR":
		sectionNameHindi = "भायपुर"
	case "BHAYPUR AANSHIK":
		sectionNameHindi = "भायपुर आंशिक"
	case "BHAYPUR ANSHIK":
		sectionNameHindi = "मुन्डिया"
	case "BHAYPUR GAIRA0":
		sectionNameHindi = "भायपुर गैरा0"
	case "BHEEMATHER ANSHIK":
		sectionNameHindi = "भीमाठेर आंशिक"
	case "BHENSIYA JUALAPUR-1":
		sectionNameHindi = "भैंसिया ज्‍वालापुर-1"
	case "BHENSIYA JUALAPUR-2":
		sectionNameHindi = "भैंसिया ज्‍वालापुर-2"
	case "BHENSODI-3":
		sectionNameHindi = "भैसोडी-3"
	case "BHENSORI-1":
		sectionNameHindi = "भैसोडी-1"
	case "BHENSORI-2":
		sectionNameHindi = "भैसोडी-2"
	case "BHESLI JAMALPUR":
		sectionNameHindi = "भैंसली जमालपुर"
	case "BHIDEYKHERA":
		sectionNameHindi = "भिडियाखेडा"
	case "BHIDVARI":
		sectionNameHindi = "भिड़वारी"
	case "BHIKANPUR ASDALPUR":
		sectionNameHindi = "भीकनपुर असदलपुर"
	case "BHIKANPURBAGAH":
		sectionNameHindi = "भीकनपुरवघा"
	case "BHIKMAPUR KULAVADA":
		sectionNameHindi = "भीकमपुर कुलवाड़ा"
	case "BHITAKHEDA":
		sectionNameHindi = "भीतखेड़ा"
	case "BHITAR GAUN 01":
		sectionNameHindi = "भीतर गांव 01"
	case "BHITAR GAUN 02":
		sectionNameHindi = "भीतर गांव 02"
	case "BHOAJIPURA":
		sectionNameHindi = "भौजीपुरा"
	case "BHOANAKPUR":
		sectionNameHindi = "भौनकपुर"
	case "BHOGLI AN":
		sectionNameHindi = "भोगली आं०"
	case "BHOGPUR HNO 1 TO 548":
		sectionNameHindi = "भोगपुर म0स0 १ से ५४८ तक"
	case "BHOGPUR HNO549 TO END":
		sectionNameHindi = "भोगपुर म0स0 549 से अंत तक"
	case "BHOGPUR MITHONI ANSHIK":
		sectionNameHindi = "भोगपुर मिठोनी आंशिक"
	case "BHOGPUR PATTI ABDULLA":
		sectionNameHindi = "भोगपुर पटटी अब्‍दुला"
	case "BHOGPUR PATTI HARSUKK":
		sectionNameHindi = "भोगपुर पटटी हरसुख"
	case "BHOGPUR":
		sectionNameHindi = "भोगपुर"
	case "BHOJPUR BHOPATPUR":
		sectionNameHindi = "भोजपुर भोपतपुर"
	case "BHOJPUR":
		sectionNameHindi = "भोजपुर"
	case "BHOLANAGALA":
		sectionNameHindi = "भोलानगला"
	case "BHOLAPUR KADEEM":
		sectionNameHindi = "भोलापुर कदीम"
	case "BHOOGPUR":
		sectionNameHindi = "भोगपुर"
	case "BHOT BAKKAL -1":
		sectionNameHindi = "भोटबक्‍काल 1"
	case "BHOT BAKKAL -2":
		sectionNameHindi = "भोटबक्‍काल 2"
	case "BHOT BAKKAL -3":
		sectionNameHindi = "भोटबक्‍काल 3"
	case "BHOT BAKKAL -4":
		sectionNameHindi = "भोटबक्‍काल 4"
	case "BHOT-2":
		sectionNameHindi = "भोट-2"
	case "BHOT-3":
		sectionNameHindi = "भोट-3"
	case "BHOT":
		sectionNameHindi = "भोट-1"
	case "BHOUPATPUR":
		sectionNameHindi = "भौपतपुर"
	case "BHOURKA":
		sectionNameHindi = "भौरका"
	case "BHOURKI JADID":
		sectionNameHindi = "भौरकी जदीद"
	case "BHOURKI KADEEM":
		sectionNameHindi = "भौरकी कदीम"
	case "BHUD MARESI":
		sectionNameHindi = "भूड़ मरेसी"
	case "BHUD":
		sectionNameHindi = "भूड"
	case "BHUDA ASALATPURA":
		sectionNameHindi = "भूडा असालतपुरा"
	case "BHUDA":
		sectionNameHindi = "भूड़ा"
	case "BHUDABAS":
		sectionNameHindi = "भूड़ाबास"
	case "BHUDASI 01":
		sectionNameHindi = "भुडासी 01"
	case "BHUDASI 02":
		sectionNameHindi = "भुडासी 0२"
	case "BHUDDI AN.H.NO 1-142":
		sectionNameHindi = "भुडडी आं०म०सं०१से १४२तक"
	case "BHUDDI AN.H.NO 143-END":
		sectionNameHindi = "भुडडी आं०म०सं०१४३ से अन्‍त तक"
	case "BHUDE KA CHAURAHA ASALATPURA":
		sectionNameHindi = "भूडे का चौराहा असालतपुरा"
	case "BHUDHANPUR ORF BILAYATNAGAR":
		sectionNameHindi = "बुढनपुर उर्फ विलायतनगर"
	case "BHUDHPUR,":
		sectionNameHindi = "बूढृपुर,"
	case "BHUKSAURA":
		sectionNameHindi = "भुकसौरा"
	case "BHUVRA AHETMALI-1":
		sectionNameHindi = "भूवरा एहतमाली 1"
	case "BHUVRA AHETMALI-2":
		sectionNameHindi = "भूवरा एहतमाली 2"
	case "BHUVRA MUSTAHAQAM-1":
		sectionNameHindi = "भूवरा मुस्‍तहकम 1"
	case "BHUVRA MUSTAHAQAM-2":
		sectionNameHindi = "भूवरा मुस्‍तहकम 2"
	case "BIBIPUR":
		sectionNameHindi = "बीबीपुर"
	case "BIBRA":
		sectionNameHindi = "बीबरा"
	case "BICHAULA- KUNDARKI":
		sectionNameHindi = "बिचौला- कुन्दरकी"
	case "BICHPURI SANGRAMPUR":
		sectionNameHindi = "विचपुरी संग्रामपुर"
	case "BIDAU":
		sectionNameHindi = "बिढऊ"
	case "BIDPURA":
		sectionNameHindi = "बिढपुरा"
	case "BIDWA NAGALA":
		sectionNameHindi = "बिढवा नगला"
	case "BIHARI NAGAR":
		sectionNameHindi = "बिहारीनगर"
	case "BIHARIPUR":
		sectionNameHindi = "बिहारीपुर"
	case "BIHAT":
		sectionNameHindi = "बीहट"
	case "BIJARKHATA UTTRI-1":
		sectionNameHindi = "बिजारखाता उत्‍तरी 1"
	case "BIJARKHATA UTTRI-2":
		sectionNameHindi = "बिजारखाता उत्‍तरी 2"
	case "BIJARKHATA UTTRI-3":
		sectionNameHindi = "बिजारखाता उत्‍तरी 3"
	case "BIJARKHATA":
		sectionNameHindi = "बिजारखाता"
	case "BIJHDA":
		sectionNameHindi = "बिझडा"
	case "BIJLI":
		sectionNameHindi = "बिजली"
	case "BIJORI":
		sectionNameHindi = "बिजौरी"
	case "BIJPURI LALJI":
		sectionNameHindi = "विचपुरी लालजी"
	case "BIKNAPUR":
		sectionNameHindi = "बिकनपुर"
	case "BILASAPUR GATE":
		sectionNameHindi = "विलासपुर गेट"
	case "BILASPUR GAIR ABAD":
		sectionNameHindi = "बिलासपुर गैर आबाद"
	case "BILASPUR":
		sectionNameHindi = "बिलासपुर"
	case "BINJHAHEDI 1 TO 100":
		sectionNameHindi = "बिन्‍जाहेडी १ से १०० तक"
	case "BINJHAHEDI H.NO 101-END":
		sectionNameHindi = "बिन्‍जाहेडी म०सं०१०१ से अंत तक"
	case "BIRAL":
		sectionNameHindi = "विराल"
	case "BIRBALPUR":
		sectionNameHindi = "बीरबलपुर"
	case "BIRDO NANGLI":
		sectionNameHindi = "बिरदो नंगली"
	case "BISATH":
		sectionNameHindi = "बिसाठ"
	case "BISATIYAN":
		sectionNameHindi = "बिसातियान"
	case "BISHANPURA URF BIDRA A0":
		sectionNameHindi = "बिशनपुरा उर्फ बिडरा आ0"
	case "BISHANPURA":
		sectionNameHindi = "बिशनपुरा"
	case "BISRA-1":
		sectionNameHindi = "बीसरा-1"
	case "BISRI":
		sectionNameHindi = "बीसरी"
	case "BITHUATHER":
		sectionNameHindi = "बिठुआठेर"
	case "BIZPURI":
		sectionNameHindi = "बिचपुरी"
	case "BLOCK - C GOVIND NAGAR":
		sectionNameHindi = "ब्‍लाक - सी गोविन्‍द नगर"
	case "BODDH VIHAR MAJHOLA":
		sectionNameHindi = "बौद्ध विहार मझोला"
	case "BOHRA":
		sectionNameHindi = "बोहरा"
	case "BOODPUR":
		sectionNameHindi = "बूढपुर"
	case "BOREKI":
		sectionNameHindi = "बौरेकी"
	case "BOSENA":
		sectionNameHindi = "बौसेना"
	case "BOVAD VALA":
		sectionNameHindi = "बोवद वाला"
	case "BRAHAMJITPUR CHANDA":
		sectionNameHindi = "बृहमजीतपुर चन्‍दा"
	case "BRAHAMPURI ANSHIK JAYANTIPUR":
		sectionNameHindi = "ब्रह्मपुरी आंशिक जयन्‍तीपुर"
	case "BRAHAMPURI":
		sectionNameHindi = "ब्रहमपुरी"
	case "BRAHMNAN":
		sectionNameHindi = "ब्रहमनान"
	case "BRIJPUR":
		sectionNameHindi = "बृजपुर"
	case "BUAPUR HARAPAL AN.":
		sectionNameHindi = "बुआपुर हरपाल आं0"
	case "BUAPUR KHEM":
		sectionNameHindi = "बुआपुर खेम"
	case "BUAPUR NATTHU AN.":
		sectionNameHindi = "बुआपुर नत्‍थू आं0"
	case "BUCHNAGAL":
		sectionNameHindi = "बूचा नांगल"
	case "BUDA NANGLA":
		sectionNameHindi = "बूढा नंगला"
	case "BUDANPUR A.":
		sectionNameHindi = "बुढनपुर आ०"
	case "BUDANPUR":
		sectionNameHindi = "बुढानपुर"
	case "BUDDHANGAR":
		sectionNameHindi = "बुद्धनगर"
	case "BUDDHI VIHAR ANSHIK":
		sectionNameHindi = "बुध्दि विहार आंशिक"
	case "BUDDHI VIHAR AVAS VIKAS":
		sectionNameHindi = "बुध्दि विहार आवास विकास"
	case "BUDDHI VIHAR, SECTOR":
		sectionNameHindi = "बुध्दि विहार, सेक्‍टर"
	case "BUDDHI VIHAR":
		sectionNameHindi = "बुध्दि विहार"
	case "BUDGARA AN.":
		sectionNameHindi = "बुडगरा आं०"
	case "BUDGARI AN.H.NO.1-498":
		sectionNameHindi = "बुडगरी आं०म०सं०01 से 498 त्‍क्‍"
	case "BUDGARI AN.H.NO.499-END":
		sectionNameHindi = "बुडगरी आं०म०स०499 से अन्‍त तक"
	case "BUDH VIHAR MAJHOLA":
		sectionNameHindi = "बुद्ध विहार मझौला"
	case "BUDHANPUR KHALSA":
		sectionNameHindi = "बुढानपुर खालसा"
	case "BUDHANPUR MU.":
		sectionNameHindi = "बुढानपुर मु."
	case "BUDHAWALA":
		sectionNameHindi = "बुढावाला"
	case "BUDHIYA KI DUKAN":
		sectionNameHindi = "बुढिया की दुकान"
	case "BUDHNAPUR":
		sectionNameHindi = "बुढनपुर"
	case "BUDHPUR MUFTI":
		sectionNameHindi = "बूढपुर मुफती"
	case "BUDHPUR":
		sectionNameHindi = "वुधपुर"
	case "BUDHUPADA":
		sectionNameHindi = "बुद्वूपाडा"
	case "BUDI DARIYAL":
		sectionNameHindi = "बूढि दडियाल"
	case "BUDIKA GAIR AB.AD":
		sectionNameHindi = "बुडिका गैर आ0"
	case "BUDPUR NAWADA":
		sectionNameHindi = "बूढपुर नवादा"
	case "BULAND KVATARS":
		sectionNameHindi = "बुलन्द क्वाटर्स"
	case "BULCHANDPUR":
		sectionNameHindi = "बूलचंदपुर"
	case "BUNDKI":
		sectionNameHindi = "बुन्‍दकी"
	case "BUNDRA KALA":
		sectionNameHindi = "बून्‍दरा कला"
	case "BUNDRA KHURD":
		sectionNameHindi = "बून्‍दरा खुर्द"
	case "BUNDUKACHIYAN AN.":
		sectionNameHindi = "बन्‍दूकचियान आ0"
	case "BURHANNAGAR":
		sectionNameHindi = "बुरहाननगर"
	case "BURHANPUR":
		sectionNameHindi = "बुढनपुर"
	case "BURHANUDDINPUR":
		sectionNameHindi = "बुरहानुददीनपुर"
	case "BURHPUR NAIN SINGH":
		sectionNameHindi = "बुढपुर नैन सिंह"
	case "BURPUR":
		sectionNameHindi = "बूढपुर"
	case "C.P.H. KANTH ROAD":
		sectionNameHindi = "सी0पी0एच0 कांठ रोड"
	case "CHACK GOWARDHAN":
		sectionNameHindi = "चक गोवर्धन"
	case "CHACK MEHDUD SANI":
		sectionNameHindi = "चक महदूद सानी"
	case "CHACK MUSTAFABAD GAIR AB.":
		sectionNameHindi = "चक मुस्‍तफाबाद गैर आं0"
	case "CHAH INCHHARAM":
		sectionNameHindi = "चाह इन्छाराम"
	case "CHAH JADIYAN V JARGARAN":
		sectionNameHindi = "चाह जडियान व जरगरान"
	case "CHAH JATTA":
		sectionNameHindi = "चाह जटटा"
	case "CHAH KESRA SIH":
		sectionNameHindi = "चाह केसरा सिह"
	case "CHAH KHAIYATAN":
		sectionNameHindi = "चाह खैयातान"
	case "CHAH KHAJAN KHAN":
		sectionNameHindi = "चाह खजान खां"
	case "CHAH KHANASAMA":
		sectionNameHindi = "चाह खानसामा"
	case "CHAH MOTE KALLAN":
		sectionNameHindi = "चाह मोटे कल्लन"
	case "CHAH SALOONO":
		sectionNameHindi = "चाह सलूनो"
	case "CHAH SHOR-1":
		sectionNameHindi = "चाह शोर-1"
	case "CHAH SHOR-2":
		sectionNameHindi = "चाह शोर-2"
	case "CHAHASATAI":
		sectionNameHindi = "चाहसतई"
	case "CHAHASHIRI":
		sectionNameHindi = "चाहशीरी"
	case "CHAHLI":
		sectionNameHindi = "चेहली"
	case "CHAHRONAK":
		sectionNameHindi = "चाह रोनक"
	case "CHAHSANG A.":
		sectionNameHindi = "चाहंसग आ०"
	case "CHAHSANG":
		sectionNameHindi = "चाहंसग आ०"
	case "CHAHSHEERI B B B-21":
		sectionNameHindi = "चाहशीरी बी 21"
	case "CHAHSHEERI B B-21":
		sectionNameHindi = "चाहशीरी बी 21"
	case "CHAHSHEERI B B-22":
		sectionNameHindi = "चाहशीरी बी 22"
	case "CHAHSHEERI B B-24":
		sectionNameHindi = "चाहशीरी बी 24"
	case "CHAHSHEERI B-24":
		sectionNameHindi = "चाहशीरी बी 24"
	case "CHAINPUR":
		sectionNameHindi = "चैनपुर"
	case "CHAJJANANGLA":
		sectionNameHindi = "छज्‍जा नंगला"
	case "CHAJJUPURA DOYAM":
		sectionNameHindi = "छज्‍जूपुरा दोयम"
	case "CHAJJUPURA":
		sectionNameHindi = "छज्‍जुपुरा"
	case "CHAJLET":
		sectionNameHindi = "छजलैट"
	case "CHAJUPURA GANDHU":
		sectionNameHindi = "छज्‍जूपुरा गन्‍धू"
	case "CHAJUPURA":
		sectionNameHindi = "छजुपुरा"
	case "CHAK ABDUL REHMAN":
		sectionNameHindi = "चक अब्‍दुल रहमान"
	case "CHAK AMBEDKAR":
		sectionNameHindi = "चक अम्‍बेडकर"
	case "CHAK BEEJNA":
		sectionNameHindi = "चक बीजना"
	case "CHAK GAZROLA":
		sectionNameHindi = "चक गजरौला"
	case "CHAK HAMIRPUR GAIRA0":
		sectionNameHindi = "हमीरपुर गैरा0"
	case "CHAK HARDASPUR":
		sectionNameHindi = "चक हरदासपुर"
	case "CHAK KHARDIYA":
		sectionNameHindi = "चक खरदिया"
	case "CHAK KUKDA GAIR AN.AD":
		sectionNameHindi = "चक कुकडा गैर आं0"
	case "CHAK KUNDESARI":
		sectionNameHindi = "चक कुण्‍डेसरी"
	case "CHAK LADPUR":
		sectionNameHindi = "चक लाडपुर"
	case "CHAK MOHD NAGAR":
		sectionNameHindi = "चक मौ0नगर"
	case "CHAK NO-1 TANDA":
		sectionNameHindi = "चक न0 1 टाण्‍डा"
	case "CHAK NO-10(2) TANDA":
		sectionNameHindi = "चक न0-10,2,टाण्‍डा"
	case "CHAK NO-10(2)TANDA":
		sectionNameHindi = "चक न0-10(2)टाण्‍डा"
	case "CHAK NO-11 TANDA":
		sectionNameHindi = "चक न0-11 टाण्‍डा"
	case "CHAK NO-12 TANDA":
		sectionNameHindi = "चक न0-12 टाण्‍डा"
	case "CHAK NO-13 TANDA":
		sectionNameHindi = "चक न0-13 टाण्‍डा"
	case "CHAK NO-14 TANDA":
		sectionNameHindi = "चक न0-14 टाण्‍डा"
	case "CHAK NO-15 TANDA":
		sectionNameHindi = "चक न0-15 टाण्‍डा"
	case "CHAK NO-16 TANDA -1":
		sectionNameHindi = "चक न0-16 टाण्‍डा-1"
	case "CHAK NO-16TANDA-2":
		sectionNameHindi = "चक न0-16 टाण्‍डा-2"
	case "CHAK NO-17 TANDA":
		sectionNameHindi = "चक न0-17 टाण्‍डा"
	case "CHAK NO-18 (1)TANDA-1":
		sectionNameHindi = "चक न0-18 (1)टाण्‍डा-1"
	case "CHAK NO-18 (2) TANDA":
		sectionNameHindi = "चक न0-18(2) टाण्‍डा"
	case "CHAK NO-18 TANDA":
		sectionNameHindi = "चक न0-18(2) टाण्‍डा"
	case "CHAK NO-19 TANDA":
		sectionNameHindi = "चक न0-19 टाण्‍डा"
	case "CHAK NO-2 TANDA":
		sectionNameHindi = "चक न0-2 टाण्‍डा"
	case "CHAK NO-20(1) TANDA":
		sectionNameHindi = "चक न0-20(1), टाण्‍डा"
	case "CHAK NO-20(2) TANDA":
		sectionNameHindi = "चक न0-20(2)टाण्‍डा"
	case "CHAK NO-21TANDA":
		sectionNameHindi = "चक न0-21 टाण्‍डा"
	case "CHAK NO-22(1)TANDA":
		sectionNameHindi = "चक न0-22(1 )टाण्‍डा"
	case "CHAK NO-22(2) TANDA":
		sectionNameHindi = "चक न0-22,2,टाण्‍डा"
	case "CHAK NO-23 TANDA":
		sectionNameHindi = "चक न0-23 टाण्‍डा"
	case "CHAK NO-24TANDA":
		sectionNameHindi = "चक न0-24 टाण्‍डा"
	case "CHAK NO-25 (2)TANDA":
		sectionNameHindi = "चक न0-25 (2) टाण्‍डा"
	case "CHAK NO-25 (3)TANDA":
		sectionNameHindi = "चक न0-25 (3) टाण्‍डा"
	case "CHAK NO-25 TANDA":
		sectionNameHindi = "चक न0-25 ,1, टाण्‍डा"
	case "CHAK NO-25(1)-2TANDA":
		sectionNameHindi = "चक न0-25(1)-2 टाण्‍डा"
	case "CHAK NO-25(3) TANDA-2":
		sectionNameHindi = "चक न0-25 (3) टाण्‍डा-2"
	case "CHAK NO-26 TANDA-2":
		sectionNameHindi = "चक न0-26 टाण्‍डा-2"
	case "CHAK NO-26TANDA":
		sectionNameHindi = "चक न0-26 टाण्‍डा-1"
	case "CHAK NO-27TANDA":
		sectionNameHindi = "चक न0-27 टाण्‍डा"
	case "CHAK NO-28 TANDA":
		sectionNameHindi = "चक न0-28 टाण्‍डा"
	case "CHAK NO-3":
		sectionNameHindi = "चक न0-3 टाण्‍डा"
	case "CHAK NO-4 TANDA":
		sectionNameHindi = "चक न0-4 टाण्‍डा"
	case "CHAK NO-5 TANDA":
		sectionNameHindi = "चक न0-5 टाण्‍डा"
	case "CHAK NO-6 TANDA":
		sectionNameHindi = "चक न0-6 टाण्‍डा"
	case "CHAK NO-7 TANDA":
		sectionNameHindi = "चक न0-7 टाण्‍डा"
	case "CHAK NO-8 TANDA":
		sectionNameHindi = "चक न0-8 टाण्‍डा"
	case "CHAK NO-9- TANDA(1)":
		sectionNameHindi = "चक न0-9- टाण्‍डा(1)"
	case "CHAK NO-9(2) TANDA":
		sectionNameHindi = "चक न0-9-(2) टाण्‍डा"
	case "CHAK SHAHJANI":
		sectionNameHindi = "चकशाहजानी"
	case "CHAK SUAR -1":
		sectionNameHindi = "चक स्‍वार 1"
	case "CHAK SUAR -2":
		sectionNameHindi = "चक स्‍वार 2"
	case "CHAK SUAR -3":
		sectionNameHindi = "चक स्‍वार 3"
	case "CHAK SUAR -4":
		sectionNameHindi = "चक स्‍वार 4"
	case "CHAK SUAR-5":
		sectionNameHindi = "चक स्‍वार 5"
	case "CHAK SUAR-6":
		sectionNameHindi = "चक स्‍वार 6"
	case "CHAK SUAR-7":
		sectionNameHindi = "चक स्‍वार 7"
	case "CHAK":
		sectionNameHindi = "चक"
	case "CHAKARAJMAL":
		sectionNameHindi = "चकराजमल"
	case "CHAKARPANPUR":
		sectionNameHindi = "चकरपानपुर"
	case "CHAKARPUR BHUD":
		sectionNameHindi = "चकरपुर भूड"
	case "CHAKARPUR KADEEM 02":
		sectionNameHindi = "चकरपुर कदीम 02"
	case "CHAKARPUR KADEEM":
		sectionNameHindi = "चकरपुर कदीम 01"
	case "CHAKARPUR":
		sectionNameHindi = "चकरपुर"
	case "CHAKASHADINGAR":
		sectionNameHindi = "चकशादीनगर"
	case "CHAKASHI":
		sectionNameHindi = "चकासी"
	case "CHAKCHAND G.A.":
		sectionNameHindi = "चकचान्‍द गै0 आ0"
	case "CHAKCHOHADMAL URF BHAJRHAWALA":
		sectionNameHindi = "चक चौहड्मल उ र्फ भ्‍जडावाला"
	case "CHAKFAJLPUR AANSHIK":
		sectionNameHindi = "चकफाजलपुर आंशिक"
	case "CHAKFERI":
		sectionNameHindi = "चकफेरी"
	case "CHAKIYA HAYATNAGAR-1":
		sectionNameHindi = "चकिया हयातनगर-1"
	case "CHAKIYA HAYATNAGAR-2":
		sectionNameHindi = "चकिया हयातनगर-2"
	case "CHAKJAMAL G.A.":
		sectionNameHindi = "चकजमाल गै0आ0"
	case "CHAKJEEM":
		sectionNameHindi = "चक जीम"
	case "CHAKLOHARA":
		sectionNameHindi = "चकलोहरा"
	case "CHAKMANI RAGHUNATHPUR":
		sectionNameHindi = "चकमनी रघुनाथपुर"
	case "CHAKNARANGI":
		sectionNameHindi = "चक नारंगी"
	case "CHAKORANGABAD":
		sectionNameHindi = "चकऔरंगाबाद"
	case "CHAKRAMAIYA G.A.":
		sectionNameHindi = "चकरमैया गै0आ0"
	case "CHAKRUSTAMPUR":
		sectionNameHindi = "चकरूस्‍तमपुर"
	case "CHAKUDAYCHAND":
		sectionNameHindi = "चक उदयचन्‍द"
	case "CHAMANPURA":
		sectionNameHindi = "चमनपुरा"
	case "CHAMARAN -2":
		sectionNameHindi = "चमारान-2"
	case "CHAMARAN-1":
		sectionNameHindi = "चमारान-1"
	case "CHAMARAN":
		sectionNameHindi = "चमारान"
	case "CHAMARPURA":
		sectionNameHindi = "चमरपुरा"
	case "CHAMDE KA GODAM ASALATPURA":
		sectionNameHindi = "चमडे का गोदाम असालतपुरा"
	case "CHAMPATPURCHAKLA":
		sectionNameHindi = "चम्‍पतपुरचकला"
	case "CHAMRA":
		sectionNameHindi = "चमरा"
	case "CHAMRAU ANSHIK":
		sectionNameHindi = "चमरौआ आंशिक"
	case "CHAMROL":
		sectionNameHindi = "चमरोल"
	case "CHAMROLLA":
		sectionNameHindi = "चमरौला"
	case "CHAMROOLA":
		sectionNameHindi = "चमरौला"
	case "CHAMROUA-1":
		sectionNameHindi = "चमरौआ-1"
	case "CHAMROUA-2":
		sectionNameHindi = "चमरौआ-2"
	case "CHAMROUA-3":
		sectionNameHindi = "चमरौआ-3"
	case "CHAMROUA-4":
		sectionNameHindi = "चमरौआ-4"
	case "CHAMROUA-5":
		sectionNameHindi = "चमरौआ-5"
	case "CHAMROUA-6":
		sectionNameHindi = "चमरौआ-6"
	case "CHAMRUNAWADA":
		sectionNameHindi = "चमरूनवादा"
	case "CHANARAN":
		sectionNameHindi = "चमारान"
	case "CHANCHAL PURAYAN":
		sectionNameHindi = "चंचलपुरियान"
	case "CHAND WALI MASJID BHADORA DEHAT":
		sectionNameHindi = "चांद वाली मस्जिद भदौरा देहात"
	case "CHANDA NANGLI":
		sectionNameHindi = "चांदा नंगली"
	case "CHANDAK JATT":
		sectionNameHindi = "चांदक जट"
	case "CHANDAK TURK":
		sectionNameHindi = "चांडक तुर्क"
	case "CHANDAKHEDI":
		sectionNameHindi = "चांदखेडी"
	case "CHAN‍DANPUR ISAPUR":
		sectionNameHindi = "चन्‍दनपुर ईसापुर"
	case "CHANDANPURA":
		sectionNameHindi = "चन्‍दनपुरा"
	case "CHANDAPUR MANGOL":
		sectionNameHindi = "चांदपुर मंगोल"
	case "CHANDAPUR":
		sectionNameHindi = "चांदपुर"
	case "CHANDAYAN":
		sectionNameHindi = "चन्‍दायन"
	case "CHANDELA":
		sectionNameHindi = "चन्‍देला"
	case "CHANDGOYALA":
		sectionNameHindi = "चन्‍दगोयला"
	case "CHANDKHEDI":
		sectionNameHindi = "चांदखेडी"
	case "CHANDLIKA":
		sectionNameHindi = "चान्‍दलीका"
	case "CHANDOK AN0":
		sectionNameHindi = "चन्‍दोक आं0"
	case "CHANDPUR FERU":
		sectionNameHindi = "चांदपुर फैरू"
	case "CHANDPUR GANESH":
		sectionNameHindi = "चांदपुर गनेश"
	case "CHANDPUR NAUBAD":
		sectionNameHindi = "चांदपुर नौआबाद"
	case "CHANDPUR":
		sectionNameHindi = "चांदपुर"
	case "CHANDPURA JADID":
		sectionNameHindi = "चन्‍दपुरा जदीद"
	case "CHANDPURA KADIM":
		sectionNameHindi = "चन्‍दपुरा कदीम"
	case "CHANDPURA KALAN -02":
		sectionNameHindi = "चन्‍द्पुरा कलॉ -02"
	case "CHANDPURA KALAN-1":
		sectionNameHindi = "चन्‍द्पुरा कलॉ -1"
	case "CHANDPURA KHURD":
		sectionNameHindi = "चन्‍द्पुरा खुर्द"
	case "CHANDPURA SALIS":
		sectionNameHindi = "चन्‍दपुरा सालिस"
	case "CHANDPURA":
		sectionNameHindi = "चन्‍दपुरा"
	case "CHANDRA NAGAR ANSHIK":
		sectionNameHindi = "चन्‍द्र नगर आंशिक"
	case "CHANDRBHANPUR KISHOR":
		sectionNameHindi = "चन्‍द्रभानपुर किशोर"
	case "CHANDUPURA NASEBPUR KANNA":
		sectionNameHindi = "चन्‍दूपुरानसीब पुर कन्‍ना"
	case "CHANDUPURA SHIKAMPUR-1":
		sectionNameHindi = "चन्‍दपुरा शीकमपुर-1"
	case "CHANDUPURA SHIKAMPUR-2":
		sectionNameHindi = "चन्‍दुपुरा शीकमपुर-2"
	case "CHANDUPURA SHIKAMPUR-3":
		sectionNameHindi = "चन्‍दुपुरा शीकमपुर-3"
	case "CHANDUPURI":
		sectionNameHindi = "चन्‍दूपुरी"
	case "CHANDWA NAGALA":
		sectionNameHindi = "चन्‍दवा नगला"
	case "CHANGERI AANSHIK":
		sectionNameHindi = "चंगेरी (आंशिक)"
	case "CHANGERI":
		sectionNameHindi = "चंगेरी"
	case "CHANGIPUR":
		sectionNameHindi = "चांगीपुर"
	case "CHAPARRA":
		sectionNameHindi = "छपर्रा"
	case "CHAPEGARAN D.":
		sectionNameHindi = "छापेग्रान द0"
	case "CHAPEGARAN U.":
		sectionNameHindi = "छापेग्रान उ0"
	case "CHAPRHA":
		sectionNameHindi = "चपडा"
	case "CHAR BAGH":
		sectionNameHindi = "चार बाग"
	case "CHARAKH VALI MASJID":
		sectionNameHindi = "चरख वाली मस्जिद"
	case "CHATARBHOJPUR KUSHAL":
		sectionNameHindi = "चतरभोजपुरकुशल"
	case "CHATARPUR NAYAK":
		sectionNameHindi = "चतरपुर नायक"
	case "CHATARPUR":
		sectionNameHindi = "चतरपुर"
	case "CHATKALI":
		sectionNameHindi = "चटकाली"
	case "CHATPUR NAKTAKHEDA":
		sectionNameHindi = "चतपुर नकटाखेडा"
	case "CHATRIYANAGAR H.NO.325-877":
		sectionNameHindi = "क्षत्रीय नगर म०सं०३२५ से ८७७तक"
	case "CHATRUWALA":
		sectionNameHindi = "चतरूवाला"
	case "CHAU KI BASTI ANSHIK":
		sectionNameHindi = "चाऊ की बस्‍ती आंशिक"
	case "CHAUDHARIYAN":
		sectionNameHindi = "चौधरियान"
	case "CHAUDHRI PATTI GAIRA.":
		sectionNameHindi = "चौधरी पटटी गैरा0"
	case "CHAUHANO WALI MILAK LAKDI":
		sectionNameHindi = "चौहानो वाली मिलक लाकडी"
	case "CHAUHANPURA":
		sectionNameHindi = "चौहानपुरा"
	case "CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-1":
		sectionNameHindi = "चौक मौ0 सईद खॉ मय फर्राशखाना कैम्प-1"
	case "CHAUK MOHD. SAEED KHAN MAY FARRASHAKHANA KAIMP-2":
		sectionNameHindi = "चौक मौ0 सईद खॉ मय फर्राशखाना कैम्प-2"
	case "CHAUKONI":
		sectionNameHindi = "चौकोनी"
	case "CHAUKPURI":
		sectionNameHindi = "चौकपुरी"
	case "CHAUNDHRI":
		sectionNameHindi = "चौधेडी"
	case "CHAUPURA-1":
		sectionNameHindi = "चाऊपुरा 1"
	case "CHAUPURA-2":
		sectionNameHindi = "चाऊपुरा 2"
	case "CHAUPURA-3":
		sectionNameHindi = "चाऊपुरा 3"
	case "CHAUPURA-4":
		sectionNameHindi = "चाऊपुरा 4"
	case "CHAUPURA":
		sectionNameHindi = "चाऊपुरा"
	case "CHAURASI GHANTA":
		sectionNameHindi = "चौरासी घंटा"
	case "CHEDRY AKBARPUR":
		sectionNameHindi = "चैदरी अकबरपुर"
	case "CHEEMIN A.":
		sectionNameHindi = "चिम्‍म्‍ान आ०"
	case "CHEHLA A.":
		sectionNameHindi = "चेहला आ0"
	case "CHELLAPUR":
		sectionNameHindi = "चेल्‍लापुर"
	case "CHHACHHARI TEEP":
		sectionNameHindi = "छाछरी टीप"
	case "CHHAJMAL WALA":
		sectionNameHindi = "छजमलवाला"
	case "CHHAJUPURA SAID":
		sectionNameHindi = "छजुपुरा सैद"
	case "CHHAKDA":
		sectionNameHindi = "छकडा"
	case "CHHATRAPUR":
		sectionNameHindi = "छत्रपुर"
	case "CHHATRE WALI GALI SARAI KOTHIWALAN":
		sectionNameHindi = "छत्रे वाली गली सराय कोठीवालान"
	case "CHHIPIYAN MAY BADHEYAN -1":
		sectionNameHindi = "छीपियान मय बढ़ेयान -1"
	case "CHHIPIYAN MAY BADHEYAN -2":
		sectionNameHindi = "छीपियान मय बढ़ेयान -2"
	case "CHHIRAVLI":
		sectionNameHindi = "छिरावली"
	case "CHHITARIYA GAJIR-1":
		sectionNameHindi = "छितरिया जागीर-1"
	case "CHHITARIYA GAJIR-2":
		sectionNameHindi = "छितरिया जागीर-2"
	case "CHHITARIYA GAJIR-3":
		sectionNameHindi = "छितरिया जागीर-3"
	case "CHHITARIYA GAJIR-4":
		sectionNameHindi = "छितरिया जागीर-4"
	case "CHICHOLI":
		sectionNameHindi = "चिचौली"
	case "CHIDIA TOLA LINEPAR ANSHIK":
		sectionNameHindi = "चिडिया टोला लाइनपार आंशिक"
	case "CHIDIYABHAVAN":
		sectionNameHindi = "चिडियाभवन"
	case "CHIDIYABHVAN":
		sectionNameHindi = "चिडि़याभवन"
	case "CHIDIYATHER":
		sectionNameHindi = "चिडियाठेर"
	case "CHIKNA":
		sectionNameHindi = "चिकना"
	case "CHIKTI RAMNAGAR-1":
		sectionNameHindi = "चिकटी रामनगर-1"
	case "CHIKTI RAMNAGAR-2":
		sectionNameHindi = "चिकटी रामनगर-2"
	case "CHIMMAN A.":
		sectionNameHindi = "चिम्‍मन आ०"
	case "CHIPIWADA":
		sectionNameHindi = "छिप्‍पीवाडा"
	case "CHIRAN":
		sectionNameHindi = "चिरान"
	case "CHIRANJI LALA":
		sectionNameHindi = "चिरजीलाल"
	case "CHITAWAR AN":
		sectionNameHindi = "छितावर आं०"
	case "CHITERI":
		sectionNameHindi = "चितेरी"
	case "CHITOUNI-1":
		sectionNameHindi = "छितौनी -1"
	case "CHITOUNI-2":
		sectionNameHindi = "छितौनी -2"
	case "CHITTUPUR":
		sectionNameHindi = "चित्तूपुर"
	case "CHODHARPUR":
		sectionNameHindi = "चौधरपुर"
	case "CHODHRANA AN H.N0 1-124":
		sectionNameHindi = "चौधराना आ०म०सं०१ से १२४"
	case "CHODHRANA AN0H.NO 125-END":
		sectionNameHindi = "चौधराना आं०म०सं१२५ से अन्‍त तक"
	case "CHODHRIYAN":
		sectionNameHindi = "चौधरियान"
	case "CHOHADHPUR NANU":
		sectionNameHindi = "चौहडपुरनानू"
	case "CHOHADPUR":
		sectionNameHindi = "चौहडपुर"
	case "CHOHANAN AN.":
		sectionNameHindi = "चौहानान आं०"
	case "CHOHARPUR GAIR AB.":
		sectionNameHindi = "चौहडपुर गैर आं0"
	case "CHOHARPUR":
		sectionNameHindi = "चौहडपुर"
	case "CHOOKBAJAR UTTRI":
		sectionNameHindi = "चौकबाजार (उत्‍तरी)"
	case "CHOUDHARYAN":
		sectionNameHindi = "चौधरियान"
	case "CHOUDHRIYAN P.":
		sectionNameHindi = "चौधरियान पं0"
	case "CHOUDHRIYAN PU.":
		sectionNameHindi = "चौधरियान पूं0"
	case "CHOUDHRIYAN":
		sectionNameHindi = "चौधरियान"
	case "CHOUHARA":
		sectionNameHindi = "चौहरा"
	case "CHOUKONI":
		sectionNameHindi = "चौकोनी"
	case "CHOWK BAZAR":
		sectionNameHindi = "चौक बाजार"
	case "CHOWKI HASAN KHAN PASHCHIMI":
		sectionNameHindi = "चौकी हसन खॉ पश्चिमी"
	case "CHOWKI HASAN KHAN PURVI":
		sectionNameHindi = "चौकी हसन खॉ पूर्वी"
	case "CHUDI WALI GALI DEHRI":
		sectionNameHindi = "चूडी वाली गली देहरी"
	case "CHUDYAKHEDA GAIR AB.AD":
		sectionNameHindi = "चुडियाखेडा गैर आं0"
	case "CHUHA NGALA":
		sectionNameHindi = "चूहा नगला"
	case "CHUKHANDI":
		sectionNameHindi = "चौखण्‍डी"
	case "CIVIL LINE ANSHIK":
		sectionNameHindi = "सिविल लाइन आंशिक"
	case "CIVIL LINE FIRST":
		sectionNameHindi = "सिविल लाईन प्रथम"
	case "CIVIL LINE HOSPITAL":
		sectionNameHindi = "सिविल लाइन अस्‍पताल"
	case "CIVIL LINE SECOND":
		sectionNameHindi = "सिविल लाईन द्वितीय"
	case "CIVIL LINES PURVI ANSHIK":
		sectionNameHindi = "सिविल लाइन्‍स पूर्वी आंशिक"
	case "CIVIL LINES PURVI":
		sectionNameHindi = "सिविल लाईन्‍स पूर्वी"
	case "COT":
		sectionNameHindi = "कोट"
	case "D.S.I.L DWARIKESH NAGAR":
		sectionNameHindi = "डी0एस0आई0एल द्वारिकेश नगर"
	case "DABKA-1":
		sectionNameHindi = "दबका-1"
	case "DABKA-2":
		sectionNameHindi = "दबका-2"
	case "DABKHEDI SALAR":
		sectionNameHindi = "दबखेडी सालार"
	case "DABTHALA":
		sectionNameHindi = "दबथला"
	case "DADHIYAL EHATMALI-1":
		sectionNameHindi = "दढियाल एहतमाली-1"
	case "DADHIYAL EHATMALI-2":
		sectionNameHindi = "दढियाल एहतमाली-2"
	case "DADHIYAL EHATMALI-3":
		sectionNameHindi = "दढियाल एहतमाली-3"
	case "DADHIYAL EHATMALI-4":
		sectionNameHindi = "दढियाल एहतमली-4"
	case "DADHIYAL EHATMALI-5":
		sectionNameHindi = "दढियाल एहतमाली-5"
	case "DADHIYAL EHATMALI-6":
		sectionNameHindi = "दढियाल एहममाली-6"
	case "DADHIYAL MUSTEKAM -1":
		sectionNameHindi = "दढियाल मुस्‍तेकम-1"
	case "DADHIYAL MUSTEKAM -2":
		sectionNameHindi = "दढियाल मुस्‍तेकम-2"
	case "DADHIYAL MUSTEKAM -3":
		sectionNameHindi = "दढियाल मुस्‍तेकम-3"
	case "DADHIYAL MUSTEKAM-4":
		sectionNameHindi = "दढियाल मुस्‍तेकम-4"
	case "DADHIYAL MUSTEKAM-5":
		sectionNameHindi = "दढियाल मुस्‍तेकम-5"
	case "DADHIYAL MUSTEKAM-6":
		sectionNameHindi = "दढियाल मुस्‍तेकम-6"
	case "DADHIYAL MUSTEKAM-7":
		sectionNameHindi = "दढियाल मुस्‍तेकम-7"
	case "DADUPUR PAINDAPUR AITMALI (GAIRA)":
		sectionNameHindi = "दादूपुर पाइन्‍दापुर ऐतमाली (गैरा)"
	case "DAHIRPUR AN0":
		sectionNameHindi = "दहीरपुर आं0"
	case "DAHLAWALA":
		sectionNameHindi = "देहलावाला"
	case "DAIVALGARH":
		sectionNameHindi = "डाईविलगढ"
	case "DAKKANAGALIYA":
		sectionNameHindi = "ढक्‍का नगलिया"
	case "DAKKHANA KHAS JUNUBI":
		sectionNameHindi = "डाकखाना खास जनूबी"
	case "DAKKHANA KHAS SHUMALI":
		sectionNameHindi = "डाकखाना खास शुमाली"
	case "DAKOULA":
		sectionNameHindi = "ढकौला"
	case "DAKOULI":
		sectionNameHindi = "ढकौली"
	case "DAKSHIN SARAY":
		sectionNameHindi = "दक्षिण सराय"
	case "DAKURIYA":
		sectionNameHindi = "ढकुरिया"
	case "DALAPATPUR ANSHIK":
		sectionNameHindi = "दलपतपुर आंशिक"
	case "DALELNAGAR-1":
		sectionNameHindi = "दलेलनगर;1"
	case "DALELNAGAR-2":
		sectionNameHindi = "दलेलनगर-2"
	case "DALKI":
		sectionNameHindi = "दलकी"
	case "DALLAWALA":
		sectionNameHindi = "दल्‍लावाला"
	case "DALLIWALA":
		sectionNameHindi = "दल्‍लीवाला"
	case "DAMMU NAGLA URF DAMPURA":
		sectionNameHindi = "दम्मू नगला उर्फ दम्मपुरा"
	case "DAN SAHAY KI MILAK":
		sectionNameHindi = "दान सहाय की मिलक"
	case "DANDA":
		sectionNameHindi = "डांडा"
	case "DANDI DURJAN":
		sectionNameHindi = "डांडी दुर्जन"
	case "DANDIYA NAYAMATGANG":
		sectionNameHindi = "डंडिया न्‍यामतगंज"
	case "DANIAPUR":
		sectionNameHindi = "दनियापुर"
	case "DANIYALPUR":
		sectionNameHindi = "दानियालपुर"
	case "DANKARA":
		sectionNameHindi = "दनकरा"
	case "DANKARI":
		sectionNameHindi = "दनकरी"
	case "DANYAPUR SHANKARPUR":
		sectionNameHindi = "दनियापुर शंकरपुर"
	case "DARA NAGAR":
		sectionNameHindi = "दारानगर"
	case "DARAKHT KAITH":
		sectionNameHindi = "दरख्त कैथ"
	case "DARAKHT KHINNI":
		sectionNameHindi = "दरख्त खिन्नी"
	case "DARANAGAR":
		sectionNameHindi = "दारानगर"
	case "DARAPUR":
		sectionNameHindi = "दारापुर"
	case "DARAV NAGAR":
		sectionNameHindi = "दरावनगर"
	case "DARBAR SADAT":
		sectionNameHindi = "दरबार सादात"
	case "DARBARASHAH":
		sectionNameHindi = "दरबाराशाह"
	case "DARBARPUR":
		sectionNameHindi = "दरबारपुर"
	case "DARGAH ANSHIK MAKBARA DAKSHINI":
		sectionNameHindi = "दरगाह आंशिक मकबरा दक्षिणी"
	case "DARGAH MAKBARA":
		sectionNameHindi = "दरगाह मकबरा"
	case "DARGOHPUR":
		sectionNameHindi = "दरगोहपुर"
	case "DARIBAPAN":
		sectionNameHindi = "दरीबापान"
	case "DARIYA GADH":
		sectionNameHindi = "दरिया गढ"
	case "DARIYAPUR A0":
		sectionNameHindi = "दरियापुर आ0"
	case "DARIYAPUR,":
		sectionNameHindi = "दरियापुर,"
	case "DARIYAPUR":
		sectionNameHindi = "दरियापुर"
	case "DARJI GALI MAKBARA DAKSHINI":
		sectionNameHindi = "दर्जी गली मकबरा दक्षिणी"
	case "DARODGRAN":
		sectionNameHindi = "दरूदग्रान"
	case "DARSHANPUR":
		sectionNameHindi = "दर्शनपुर"
	case "DARWAR A.":
		sectionNameHindi = "दरबाडा आ०"
	case "DARWAR":
		sectionNameHindi = "दरवड"
	case "DARWARA A.":
		sectionNameHindi = "दरबाडा आ०"
	case "DAS SARAI":
		sectionNameHindi = "दस सराय"
	case "DATANGAR":
		sectionNameHindi = "दातानगर"
	case "DATTYANA A.":
		sectionNameHindi = "दत्‍तियाना आ०"
	case "DAUDPUR HAJI":
		sectionNameHindi = "दाऊदपुर हाजी"
	case "DAUDPUR NANHERA":
		sectionNameHindi = "दाऊदपुर नन्‍हेडा"
	case "DAUDPUR":
		sectionNameHindi = "दाउदपुर"
	case "DAUKPURI TANDA -1":
		sectionNameHindi = "डौंकपुरी टांडा 1"
	case "DAUKPURI TANDA-2":
		sectionNameHindi = "डौंकपुरी टांडा 2"
	case "DAUKPURI TANDA-3":
		sectionNameHindi = "डौंकपुरी टांडा 3"
	case "DAUKPURI TANDA-4":
		sectionNameHindi = "डौंकपुरी टांडा 4"
	case "DAULAPURI BAMANIYA":
		sectionNameHindi = "दौलपुरी बमनिया"
	case "DAULARI ANSHIK":
		sectionNameHindi = "दौलारी आंशिक"
	case "DAULATBAGH ANSHIK":
		sectionNameHindi = "दौलतबाग आंशिक"
	case "DAULATBAGH BALMIKI BASTI":
		sectionNameHindi = "दौलतबाग बाल्‍मीकि बस्‍ती"
	case "DAULATBAGH NAGPHANI UTTARI ANSHIK":
		sectionNameHindi = "दौलतबाग नागफनी उत्‍तरी आंशिक"
	case "DAULATPUR BILLOCH":
		sectionNameHindi = "दौलतपुर बिल्‍लौच"
	case "DAULATPUR":
		sectionNameHindi = "दौलतपुर"
	case "DAULAVALA":
		sectionNameHindi = "दौलावाला"
	case "DAULLATPUR":
		sectionNameHindi = "दौलतपुर"
	case "DAULTAPUR BHARAKAU":
		sectionNameHindi = "दौलतपुर भरकऊ"
	case "DAWAL KHEDI":
		sectionNameHindi = "दवखेडी मेव"
	case "DAWARPUR GAIR AB.AD":
		sectionNameHindi = "दामरपुर गैर आं0"
	case "DAYALPU GYANPUR":
		sectionNameHindi = "दयालपुर ज्ञानपुर"
	case "DAYALWALA":
		sectionNameHindi = "दयालवाला"
	case "DAYAMNANGLA":
		sectionNameHindi = "दायम नंगला"
	case "DEDANAGALA NAGALA":
		sectionNameHindi = "दीदा नंगला"
	case "DEENDAYAL NAGAR ANSHIK":
		sectionNameHindi = "दीनदयाल नगर आंशिक"
	case "DEENDAYAL NAGAR PHASE-1":
		sectionNameHindi = "दीनदयाल नगर फेस-1"
	case "DEENDAYAL NAGAR PHASE-2":
		sectionNameHindi = "दीनदयाल नगर फेस-2"
	case "DEENPUR":
		sectionNameHindi = "दीनपुर"
	case "DEEVANANDPUR GARHI":
		sectionNameHindi = "देवानन्‍दपुर गढी"
	case "DEHALI DARAVAJA MAY GAUKHANA KAIMP":
		sectionNameHindi = "देहली दरवाजा मय गऊखाना कैम्प"
	case "DEHRA NIKAT OMRI KLA":
		sectionNameHindi = "डेहरा निकट उमरी कलां"
	case "DEHRI GAON ANSHIK":
		sectionNameHindi = "देहरी गांव आंशिक"
	case "DEHRI GHAT":
		sectionNameHindi = "देहरी घाट"
	case "DEHRI JUMMAN":
		sectionNameHindi = "देहरी जुम्‍मन"
	case "DEHRIA ANSHIK":
		sectionNameHindi = "डहरिया आंशिक"
	case "DELPURA KANNA":
		sectionNameHindi = "डेलपुरा कन्‍ना"
	case "DELPURA NAINU":
		sectionNameHindi = "डेलपुरा नैनू"
	case "DERABULANDI":
		sectionNameHindi = "डेहराबुलन्‍दी"
	case "DESONDIWALA DESH":
		sectionNameHindi = "दिसोनदीवालादेश"
	case "DEVAPUR AI0":
		sectionNameHindi = "देवापुर ऐ0"
	case "DEVAPUR ANSHIK":
		sectionNameHindi = "देवापुर आंशिक"
	case "DEVAPUR":
		sectionNameHindi = "देवापुर"
	case "DEVIPUR URF NAGLA":
		sectionNameHindi = "देवीपुर उर्फ नगला"
	case "DEVIPURA (ANSHIK)":
		sectionNameHindi = "देवीपुरा (आंशिक)"
	case "DEVIPURA":
		sectionNameHindi = "देवीपुरा"
	case "DEVRANIYA SHUMALI 1":
		sectionNameHindi = "देवरनिया शुमाली 1"
	case "DEVRANIYA SHUMALI 2":
		sectionNameHindi = "देवरनिया शुमाली 2"
	case "DEVRI BUJURG":
		sectionNameHindi = "देवरी बुजुर्ग"
	case "DEVRI KHURD":
		sectionNameHindi = "देवरी खुर्द"
	case "DEVRI":
		sectionNameHindi = "देवरी"
	case "DEWAPUR":
		sectionNameHindi = "देवापुर"
	case "DEWRANIYA SHARKI":
		sectionNameHindi = "देवरनिया शर्की"
	case "DHADI MEHMOODPUR":
		sectionNameHindi = "दाढी महमूदपुर"
	case "DHAILA AHIR":
		sectionNameHindi = "ढेला अहीर"
	case "DHAILY AHIR":
		sectionNameHindi = "ढेली अहीर"
	case "DHAKI SADHO":
		sectionNameHindi = "ढाकी साधो"
	case "DHAKI":
		sectionNameHindi = "ढाकी"
	case "DHAKIYA 02":
		sectionNameHindi = "ढकिया 02"
	case "DHAKIYA 03":
		sectionNameHindi = "ढकिया 03"
	case "DHAKIYA 04":
		sectionNameHindi = "ढकिया 04"
	case "DHAKIYA 1":
		sectionNameHindi = "ढकिया 1"
	case "DHAKIYA JAMAPUR":
		sectionNameHindi = "ढकिया जमापुर"
	case "DHAKIYAJAT":
		sectionNameHindi = "ढकियाजट"
	case "DHAKIYANRU":
		sectionNameHindi = "ढकियानरू"
	case "DHAKIYAPIRU":
		sectionNameHindi = "ढकियापीरु"
	case "DHAKKA ANSHIK":
		sectionNameHindi = "ढक्‍का आंशिक"
	case "DHAKKA HAJINAGAR 1":
		sectionNameHindi = "ढक्‍का हाजीनगर 1"
	case "DHAKKA HAJINAGAR-2":
		sectionNameHindi = "ढक्‍का हाजीनगर 2"
	case "DHAKKA KARAMCHAND":
		sectionNameHindi = "ढक्‍का कर्मचन्‍द"
	case "DHAKPURA":
		sectionNameHindi = "ढकपुरा"
	case "DHAKYA":
		sectionNameHindi = "ढि‍कया"
	case "DHALLIYA":
		sectionNameHindi = "धल्लिया"
	case "DHAMORA-1":
		sectionNameHindi = "धमोरा-1"
	case "DHAMORA-2":
		sectionNameHindi = "धमोरा-2"
	case "DHAMORA-3":
		sectionNameHindi = "धमोरा-3"
	case "DHAMPUR HUSAINPUR AN.":
		sectionNameHindi = "धामपुर हुसैनपुर आ0"
	case "DHAMROLA":
		sectionNameHindi = "धमरौला"
	case "DHANELI POORVI-1":
		sectionNameHindi = "धनेली पूर्वी-1"
	case "DHANELI POORVI-2":
		sectionNameHindi = "धनेली पूर्वी-2"
	case "DHANELI UTTARI-1":
		sectionNameHindi = "धनेली उत्‍तरी-1"
	case "DHANELI UTTARI-2":
		sectionNameHindi = "धनेली उत्‍तरी-2"
	case "DHANNUPURATUR KHEDA":
		sectionNameHindi = "धन्नूपुरा तुरखेड़ा"
	case "DHANOORA":
		sectionNameHindi = "धनौरा"
	case "DHANOORI-1":
		sectionNameHindi = "धनौरी 1"
	case "DHANORA":
		sectionNameHindi = "धनौरा"
	case "DHANORI KUNWAR":
		sectionNameHindi = "धनौरी कुंवर"
	case "DHANORI":
		sectionNameHindi = "धनौरी"
	case "DHANOURA":
		sectionNameHindi = "धनोरा"
	case "DHANOURI-2":
		sectionNameHindi = "धनौरी 2"
	case "DHANPUR NIKAT SHAHDRA":
		sectionNameHindi = "धनपुर निकट शाहदरा"
	case "DHANSINI AN0":
		sectionNameHindi = "धनसीनी आं0"
	case "DHANUPURA-1":
		sectionNameHindi = "धनुपुरा-1"
	case "DHANUPURA-2":
		sectionNameHindi = "धनूपुरा-2"
	case "DHANUPURA-3":
		sectionNameHindi = "धनुपुरा-3"
	case "DHANUPURA":
		sectionNameHindi = "धनुपुरा"
	case "DHARAK NANGLA":
		sectionNameHindi = "धारक नंगला"
	case "DHARAMDAS":
		sectionNameHindi = "धर्मदास"
	case "DHARAMNAGARI COLONY":
		sectionNameHindi = "धर्मनगरी कालोनी"
	case "DHARAMPUR BHAGWAN GAIR ABAD":
		sectionNameHindi = "धर्मपुर भगवान गैर आबाद"
	case "DHARAMPUR BHOJA URF PUNDRI KALA":
		sectionNameHindi = "धर्मपुर भोजा उर्फ पूण्‍डरी कला"
	case "DHARAMPUR DHANSI GAIR ABAD":
		sectionNameHindi = "धर्मपुर धनसी गैर आबाद"
	case "DHARAMPUR UTTARI":
		sectionNameHindi = "धर्मपुर उत्‍तरी"
	case "DHARAMPUR":
		sectionNameHindi = "धर्मपुर"
	case "DHARAMPURA":
		sectionNameHindi = "धर्मपुरा"
	case "DHARAMSHANANGLI":
		sectionNameHindi = "धर्मशानंगली"
	case "DHARMA GARHI":
		sectionNameHindi = "धर्मा गढी"
	case "DHARMASHALA MAY MAI KA THAN":
		sectionNameHindi = "धर्मशाला मय माई का थान"
	case "DHARMAWALA":
		sectionNameHindi = "धर्मावाला"
	case "DHARMPUR GARVI":
		sectionNameHindi = "धर्मपुर गर्वी"
	case "DHARMPUR KALAN":
		sectionNameHindi = "धर्मपुर कलां"
	case "DHARMPUR URF JOGIWALA":
		sectionNameHindi = "धर्मपुर उर्फ जोगीवाला"
	case "DHARMPUR":
		sectionNameHindi = "धर्मपुर"
	case "DHARMUPOOTA":
		sectionNameHindi = "धरमूपौटा"
	case "DHARUPUR A.":
		sectionNameHindi = "धारूपुर आ०"
	case "DHATRARA":
		sectionNameHindi = "धतरारा"
	case "DHATURA MEGHA NAGLA AANSHIK":
		sectionNameHindi = "धतूरा मेघा नगला आंशिक"
	case "DHATURA MEGHA NAGLA AI0":
		sectionNameHindi = "धतूरा मेघा नगला ऐ0"
	case "DHAWANI BUJURG-1":
		sectionNameHindi = "धावनी बुजुर्ग-1"
	case "DHAWANI BUJURG-2":
		sectionNameHindi = "धावनी बुजुर्ग-2"
	case "DHAWNI HASANPUR-1":
		sectionNameHindi = "धावनी हसनपुर-1"
	case "DHAWNI HASANPUR-2":
		sectionNameHindi = "धावनी हसनपुर-2"
	case "DHEAVI VALI MILAK":
		sectionNameHindi = "धेावी वाली मिलक"
	case "DHEEMRI ANSHIK":
		sectionNameHindi = "धीमरी आंशिक"
	case "DHEEMRI KHWAZA NAGAR ANSHIK":
		sectionNameHindi = "धीमरी ख्‍वाजा नगर आंशिक"
	case "DHELA GUJJAR":
		sectionNameHindi = "ढेला गुजर"
	case "DHELI GUJJAR":
		sectionNameHindi = "ढेली गूजर"
	case "DHEVERPURA A.":
		sectionNameHindi = "धीवरपुरा आ०"
	case "DHIMARKHEDA":
		sectionNameHindi = "धीमरखेडा"
	case "DHIMRI":
		sectionNameHindi = "धीमरी"
	case "DHINGARPUR":
		sectionNameHindi = "धींगरपुर"
	case "DHINGERPUR":
		sectionNameHindi = "धींगरपुर"
	case "DHIRAZNAGAR":
		sectionNameHindi = "धीरजनगर"
	case "DHNORA":
		sectionNameHindi = "धनौरा"
	case "DHOLAGAD":
		sectionNameHindi = "धौलागढ"
	case "DHOLANPUR":
		sectionNameHindi = "ढोलनपुर"
	case "DHOLSAR":
		sectionNameHindi = "ढोलशर"
	case "DHOMANPUR":
		sectionNameHindi = "धोमनपुर"
	case "DHUNDLI":
		sectionNameHindi = "धूधली"
	case "DHUNPURA":
		sectionNameHindi = "धनुपुरा"
	case "DHURARA":
		sectionNameHindi = "धुराडा"
	case "DHURIYAI":
		sectionNameHindi = "धुरियाई"
	case "DIBDIBA-1":
		sectionNameHindi = "दिबदिबा-1"
	case "DIBDIBA-2":
		sectionNameHindi = "दिबदिबा-2"
	case "DIBDIBA-3":
		sectionNameHindi = "दिबदिबा-3"
	case "DIBDIBA-4":
		sectionNameHindi = "दिबदिबा-4"
	case "DIBDIBA-5":
		sectionNameHindi = "दिबदिबा-5"
	case "DIDORA":
		sectionNameHindi = "डिडोरा"
	case "DIDORI":
		sectionNameHindi = "डिडोरी"
	case "DILARI CHAGERI":
		sectionNameHindi = "डिलारी चगेरी"
	case "DILARI":
		sectionNameHindi = "डिलारी"
	case "DILPURA-1":
		sectionNameHindi = "दिलपुरा-1"
	case "DILPURA-2":
		sectionNameHindi = "दिलपुरा-2"
	case "DILRA RAIPUR":
		sectionNameHindi = "डिलरा रायपुर"
	case "DINADARAPUR":
		sectionNameHindi = "दीनदारपुर"
	case "DINAURA":
		sectionNameHindi = "दिनौरा"
	case "DINDARPURA":
		sectionNameHindi = "दिनदारपुरा"
	case "DINGARPUR AANSHIK":
		sectionNameHindi = "डींगरपुर आंशिक"
	case "DINORA":
		sectionNameHindi = "दिनौडा"
	case "DITNAPUR AN.":
		sectionNameHindi = "दीत्‍तनपुर आं0"
	case "DIVAAN":
		sectionNameHindi = "दीवान"
	case "DIVIYANAGLA":
		sectionNameHindi = "दिविया नगला"
	case "DIWAN PARMANAND":
		sectionNameHindi = "दीवान परमानन्‍द"
	case "DLEEPUR":
		sectionNameHindi = "दलीपपुर"
	case "DOARKADASPUR":
		sectionNameHindi = "द्वारका दासपुर"
	case "DODRAJPUR":
		sectionNameHindi = "दोदराज पुर"
	case "DOHAIYA":
		sectionNameHindi = "डो‍हरिया"
	case "DOHARI":
		sectionNameHindi = "डोहरी"
	case "DOHARIYA":
		sectionNameHindi = "डोहरिया"
	case "DOLTABAD":
		sectionNameHindi = "दोलताबाद"
	case "DOMAHALA ROD MAY GHER GULAM NASIR KHAN":
		sectionNameHindi = "दोमहला रोड मय घेर गुलाम नासिर खां"
	case "DOMGHAR":
		sectionNameHindi = "डोमघर"
	case "DOODHLI":
		sectionNameHindi = "दूधली"
	case "DOSTPUR GAIR A.":
		sectionNameHindi = "दोस्‍तपुरा गै0आ0"
	case "DOULATPUR SUKHKHA":
		sectionNameHindi = "दौलतपुर सुक्‍खा"
	case "DUABAT":
		sectionNameHindi = "दुवावट"
	case "DUARIKAPURI":
		sectionNameHindi = "द्वारिकापुरी"
	case "DUDAI-2":
		sectionNameHindi = "डुडई-2"
	case "DUDELA":
		sectionNameHindi = "डूडैला"
	case "DUDHLA":
		sectionNameHindi = "दुधला"
	case "DUGANPUR":
		sectionNameHindi = "दुगनपुर"
	case "DUKAN SONA BINNA":
		sectionNameHindi = "दुकान सोना बिन्‍ना"
	case "DULHAPUR GAIR ABAD":
		sectionNameHindi = "दुल्‍हापुर गैर आबाद"
	case "DULHAPUR PATTI CHAUHAN":
		sectionNameHindi = "दुल्हापुर पट्टी चौहान"
	case "DULHAPUR PATTI JAT":
		sectionNameHindi = "दूल्हापुर पट्टी जाट"
	case "DULHEPUR NIKAT CHANGERI":
		sectionNameHindi = "दूल्‍हेपुर निकट चंगेरी"
	case "DULICHANDPUR":
		sectionNameHindi = "दुलीचन्‍दपुर"
	case "DUMDAMA JIGAR COLONY":
		sectionNameHindi = "दमदमा जिगर कालौनी"
	case "DUMHELA KHALSA":
		sectionNameHindi = "दुम्‍हैला खालसा"
	case "DUNDAI-1":
		sectionNameHindi = "डुडई-1"
	case "DUNDAWALA AEHATMALI , MUSTEHAQAM-1":
		sectionNameHindi = "दून्‍दावाला एहतमाली व मुस्‍तहकम 1"
	case "DUNDAWALA AEHATMALI , MUSTEHAQAM-2":
		sectionNameHindi = "दून्‍दावाला एहतमाली व मुस्‍तहकम 2"
	case "DUNDAWALA AEHATMALI , MUSTEHKAM 3":
		sectionNameHindi = "दून्‍दावाला एहतमाली व मुस्‍तहकम 3"
	case "DUNDAWALA AEHATMALI , MUSTEHKAM 4":
		sectionNameHindi = "दून्‍दावाला एहतमाली व मुस्‍तहकम 4"
	case "DUNDAWALA AEHATMALI , MUSTEHKAM 5":
		sectionNameHindi = "दून्‍दावाला एहतमाली व मुस्‍तहकम 5"
	case "DUNGARPUR":
		sectionNameHindi = "डूंगरपुर"
	case "DUNGERPUR":
		sectionNameHindi = "डुंगरपुर"
	case "DUPEDHA":
		sectionNameHindi = "दुपेड़ा"
	case "DUPTY GANJ ANSHIK":
		sectionNameHindi = "डिप्‍टी गंज आंशिक"
	case "DURGESH NAGAR ANSHIK SITAPURI":
		sectionNameHindi = "दुर्गेश नगर आंशिक सीतापुरी"
	case "DURGESH NAGAR DOBLE FATAK":
		sectionNameHindi = "दुर्गेश नगर डबल फाटक"
	case "DURJANPUR":
		sectionNameHindi = "दुर्जनपुर"
	case "DURNAGLA":
		sectionNameHindi = "दुर्गनगला"
	case "DUWARKAPUR":
		sectionNameHindi = "द्वारकापुर"
	case "DWARIKAPURI":
		sectionNameHindi = "द्वारिकापुरी"
	case "E":
		sectionNameHindi = "E"
	case "EBRAHIMPUR":
		sectionNameHindi = "इब्राहीमपुर"
	case "EDAM EIBZE COLLEGE ROAD PEERZADA DAKSHIN":
		sectionNameHindi = "एडम इब्‍ज कालिज रोड पीरजादा दक्षिण"
	case "EDELPUR JOGIYA SIHALI":
		sectionNameHindi = "अदलपुर जोगिया सिहाली"
	case "EIDGAH AWADI ANSHIK":
		sectionNameHindi = "र्इदगाह आवादी आंशिक"
	case "EIDGAHA":
		sectionNameHindi = "ईदगाह"
	case "EKTA COLONY LINEPAR":
		sectionNameHindi = "एकता कालौनी लाइनपार"
	case "ELAHABAD":
		sectionNameHindi = "इलाहाबाद"
	case "ENAYATPUR":
		sectionNameHindi = "इनायतपुर"
	case "FADASHEKH MAY BANS MANDI":
		sectionNameHindi = "फड़शेख मय बांस मण्डी"
	case "FAGLABAD PARMANANDPUR ANO":
		sectionNameHindi = "फजलाबाद परमानन्‍दपुर आ0"
	case "FAGLABAD PARMANANDPUR":
		sectionNameHindi = "फजलाबाद परमानन्‍दपुर"
	case "FAIJGANJ PASHCHIMI":
		sectionNameHindi = "फैजगंज पश्चिमी"
	case "FAIJGANJ":
		sectionNameHindi = "फैजगंज"
	case "FAIJNAGAR":
		sectionNameHindi = "फैजनगर"
	case "FAIJULLAGANJ":
		sectionNameHindi = "फैजुल्लागंज"
	case "FAIJULLANGAR":
		sectionNameHindi = "फैजुल्लानगर"
	case "FAIZNAGAR":
		sectionNameHindi = "फैजनगर"
	case "FAIZPUR RAMANAND":
		sectionNameHindi = "फैजपुर रामानन्‍द"
	case "FAIZPUR":
		sectionNameHindi = "फैजपुर"
	case "FAIZULLANAGAR-1":
		sectionNameHindi = "फैजुल्‍लानगर-1"
	case "FAIZULLANAGAR-2":
		sectionNameHindi = "फैजुल्‍लानगर-2"
	case "FAJALPUR ANSHIK":
		sectionNameHindi = "फाजलपुर आंशिक"
	case "FAJALPUR DHAKI A. H.N. 1 TO 114":
		sectionNameHindi = "फजलपुर ढाकी आ0 म0नं0 १ से ११४ तक"
	case "FAJALPUR DHAKI A. H.N. 115 TO END":
		sectionNameHindi = "फजलपुर ढाकी आ0 म0नं0 ११५ से अंत तक"
	case "FAJALPUR NEAR DHANOORY":
		sectionNameHindi = "फाजलपुर नि0 धनौरी"
	case "FAJALPUR POUTA":
		sectionNameHindi = "फजलपुर पौटा"
	case "FAJALPUR":
		sectionNameHindi = "फजलपुर"
	case "FAJULANAGAR":
		sectionNameHindi = "फैजुल्‍लानगर"
	case "FAJULLAPUR AN.":
		sectionNameHindi = "फैजुल्‍लापुर आं0"
	case "FAJULLAPURMAHESH":
		sectionNameHindi = "फैजुल्‍लापुरमहेश"
	case "FAKEERGANJ PASHMI":
		sectionNameHindi = "फकीरगंज (पश्चिमी)"
	case "FAKEERGANJ PURBI":
		sectionNameHindi = "फकीरगंज (पूर्वी)"
	case "FAKEERGANJ URF GHOSIPURA":
		sectionNameHindi = "फकीरगंज उर्फ घोसीपुरा"
	case "FAKEERPURA ANSHIK":
		sectionNameHindi = "फकीरपुरा आंशिक"
	case "FALJABAD":
		sectionNameHindi = "फजलाबाद"
	case "FALOUDA":
		sectionNameHindi = "फलौदा"
	case "FALOUDI":
		sectionNameHindi = "फलौदी"
	case "FARASHTOLI":
		sectionNameHindi = "फराश टोली"
	case "FARAZPUR":
		sectionNameHindi = "फरजपुर"
	case "FAREEDNAGAR":
		sectionNameHindi = "फरीदनगर"
	case "FAREEDPUR BHENDI":
		sectionNameHindi = "फरीदपुर भैंडी"
	case "FAREEDPUR BHOGI":
		sectionNameHindi = "फरीदपुर भोगी"
	case "FAREEDPUR HAJI":
		sectionNameHindi = "फरीदपुर हाजी"
	case "FAREEDPUR KASAM":
		sectionNameHindi = "फरीदपुर कासम"
	case "FARHA SAN SHUMALI":
		sectionNameHindi = "फरहा सान शुमाली"
	case "FARHASAN JUNUBI":
		sectionNameHindi = "फरहा सान जनूबी"
	case "FARIDABAD":
		sectionNameHindi = "फरीदाबाद"
	case "FARIDAPUR HAMIR":
		sectionNameHindi = "फरीदपुर हमीर"
	case "FARIDPUR ALAM GAIR AB.":
		sectionNameHindi = "फरीदपुर आलम"
	case "FARIDPUR BHARTA":
		sectionNameHindi = "फरीदपुर भरता"
	case "FARIDPUR BHOGAN":
		sectionNameHindi = "फरीदपुर भोगन"
	case "FARIDPUR BHORO":
		sectionNameHindi = "फरीदपुर भोरू"
	case "FARIDPUR DALLU":
		sectionNameHindi = "फरीदपुर डल्‍लू"
	case "FARIDPUR DARGO":
		sectionNameHindi = "फरीदपुर दर्गो"
	case "FARIDPUR DULLI":
		sectionNameHindi = "फरीदपुर दुल्‍ली"
	case "FARIDPUR KAZI AN":
		sectionNameHindi = "फरीदपुर काजी"
	case "FARIDPUR MALHU":
		sectionNameHindi = "फरीदपुर मल्‍हू"
	case "FARIDPUR MAN":
		sectionNameHindi = "फरीदपुर मान"
	case "FARIDPUR NIZAM":
		sectionNameHindi = "फरीदपुर निजाम"
	case "FARIDPUR QAZI AN":
		sectionNameHindi = "फरीदपुर काजी"
	case "FARIDPUR SADHIRAN":
		sectionNameHindi = "फरीदपुरसधीरन"
	case "FARIDPUR SANSARU URF DHOKALPUR":
		sectionNameHindi = "फरीदपुर संसारू उर्फ धोकलपुर"
	case "FARIDPUR UDDA":
		sectionNameHindi = "फरीदपुर उददा"
	case "FARIEDPUR MIRA":
		sectionNameHindi = "फरीदपुर मीरा"
	case "FARRAASHAN 01":
		sectionNameHindi = "फर्राशान 01"
	case "FARRAASHAN 02":
		sectionNameHindi = "फर्राशान 02"
	case "FARRAASHAN 03":
		sectionNameHindi = "फर्राशान 03"
	case "FARRAASHAN 04":
		sectionNameHindi = "फर्राशान 04"
	case "FARSHTOLI":
		sectionNameHindi = "फराश टोली"
	case "FATAHANGA AN.":
		sectionNameHindi = "फतेहनगर आं0"
	case "FATAHANGAR AN.":
		sectionNameHindi = "फतेहनगर आ0"
	case "FATAHAPUR JAMAL AN0":
		sectionNameHindi = "फतेहपुर जमाल आ0"
	case "FATAN PUR":
		sectionNameHindi = "फतन पुर"
	case "FATEHABAD":
		sectionNameHindi = "फतेहाबाद"
	case "FATEHAPUR LAL B.A":
		sectionNameHindi = "फतेहपुर लाल बी0ए"
	case "FATEHAPUR SALAVAT NAGAR":
		sectionNameHindi = "फतेहपुर सलावत नगर"
	case "FATEHAULLA GANJ MANIHARAN":
		sectionNameHindi = "फतेहउल्ला गंज मनिहारान"
	case "FATEHAULLAGANJ BANJARAN":
		sectionNameHindi = "फतेहउल्लागंज बंजारान"
	case "FATEHAULLAGANJ CHALACHITRA ROAD":
		sectionNameHindi = "फतेहउल्लागंज चलचित्र रोड"
	case "FATEHAULLAGANJ DHOBIYAN":
		sectionNameHindi = "फतेहउल्लागंज धोबियान"
	case "FATEHAULLAGANJ KURESHIYAN":
		sectionNameHindi = "फतेहउल्लागंज कुरेशियान"
	case "FATEHAULLAGANJ NAIBASTI":
		sectionNameHindi = "फतेहउल्लागंज नईबस्ती"
	case "FATEHAULLAGANJ SAIFIYAN":
		sectionNameHindi = "फतेहउल्लागंज सैफियान"
	case "FATEHPUR ASAL":
		sectionNameHindi = "फतेहपुर असल"
	case "FATEHPUR BULANDI":
		sectionNameHindi = "फतेहपुर बलन्‍दी"
	case "FATEHPUR CHAK":
		sectionNameHindi = "फतेहपुर चक"
	case "FATEHPUR KHURD":
		sectionNameHindi = "फतेहपुर खुर्द"
	case "FATEHPUR NOUABAD":
		sectionNameHindi = "फतेहपुर नौआबाद"
	case "FATEHPUR RAJARAM":
		sectionNameHindi = "फतेहपुर राजाराम"
	case "FATEHPUR":
		sectionNameHindi = "फतेहपुर"
	case "FATEHULLAGANJ ANSARIYAN":
		sectionNameHindi = "फतेह उल्लागंज अन्सारियान"
	case "FATEHULLAGANJ JATVAN MANDIR":
		sectionNameHindi = "फतेहउल्लागंज जाटवान मन्दिर"
	case "FATEHULLAGANJ TAHIRA ROD":
		sectionNameHindi = "फतेह उल्लागंज ताहिरा रोड"
	case "FATEHULLAPUR DURG":
		sectionNameHindi = "फतहउल्‍लापुर दुर्ग"
	case "FATEPUR A":
		sectionNameHindi = "फतेहपुर ए"
	case "FATEPUR DHARA":
		sectionNameHindi = "फतेहपुर धारा"
	case "FATEPUR KALA":
		sectionNameHindi = "फतेहपुर कला"
	case "FATEPUR SABHACHAND":
		sectionNameHindi = "फतेहपुर सभाचन्‍द"
	case "FATHAY ULLAH NAGAR":
		sectionNameHindi = "फतेहउल्‍लानगर"
	case "FATHPUR":
		sectionNameHindi = "फतेहपुर"
	case "FATHULLAHPUR POTTA":
		sectionNameHindi = "फतेहउल्‍लापुर पौटा"
	case "FATTANPUR":
		sectionNameHindi = "फततनपुर"
	case "FAT‍TEHPUR KHAS":
		sectionNameHindi = "फत्‍तेहपुर खास"
	case "FATTEHPUR VISHNOI":
		sectionNameHindi = "फत्‍तेहपुर विश्‍नोई"
	case "FATTEHPURI ANSHIK":
		sectionNameHindi = "फत्‍तेहपुरी आंशि‍क"
	case "FATTEPUR":
		sectionNameHindi = "फत्‍तेपुर"
	case "FATTUWALA":
		sectionNameHindi = "फत्‍तूवाला"
	case "FAULADPUR":
		sectionNameHindi = "फौलादपुर"
	case "FAZALALIPUR":
		sectionNameHindi = "फजलअलीपुर"
	case "FAZALPUR FATEHULLA":
		sectionNameHindi = "फजलपुर फतेहउल्‍ला"
	case "FAZALPUR HABIB AN0":
		sectionNameHindi = "फजलपुर हबीब आं0"
	case "FAZALPUR HAIBAT":
		sectionNameHindi = "फजलपुर हैबत"
	case "FAZALPUR KHAS":
		sectionNameHindi = "फजलपुर खास"
	case "FAZALPUR MAN":
		sectionNameHindi = "फजलपुर मान"
	case "FAZALPUR PUSWADA":
		sectionNameHindi = "फाजलपुर पुसवाडा"
	case "FAZALPUR TABELA":
		sectionNameHindi = "फजलपुर तबेला"
	case "FAZALPUR":
		sectionNameHindi = "फजलपुर"
	case "FAZGANJ":
		sectionNameHindi = "फैजगंज"
	case "FAZIPUR KADER":
		sectionNameHindi = "फैजीपुर खादर"
	case "FAZIPUR KASWA":
		sectionNameHindi = "फैजीपुर कस्‍बा"
	case "FEELKHANA ANSHIK":
		sectionNameHindi = "फीलखाना आंशिक"
	case "FIELD":
		sectionNameHindi = "फील्‍ड"
	case "FILAKHANA JADID":
		sectionNameHindi = "फीलखाना जदीद"
	case "FIROZPUR MANDU":
		sectionNameHindi = "फिरोजपुर मंडू"
	case "FIROZPUR MUBARAK URF NAYA GAOI":
		sectionNameHindi = "फिरोजपुर मुबारक उर्फ नया गांव"
	case "FIROZPUR NAROTAM":
		sectionNameHindi = "फिरोजपुर नरोत्‍तम"
	case "FIROZPUR TAREEKAM":
		sectionNameHindi = "फिरोजपुर तरीकम"
	case "FIROZPUR UGRSAIN":
		sectionNameHindi = "फिरोजपुर उग्रसैन"
	case "FIROZPUR":
		sectionNameHindi = "फिरोजपुर"
	case "FOLADPUR":
		sectionNameHindi = "फौलादपुर"
	case "FOLDA PATTI":
		sectionNameHindi = "फोल्‍डा पटटी"
	case "FONDAPUR":
		sectionNameHindi = "फौन्‍दापुर"
	case "FOOLPUR":
		sectionNameHindi = "फूलपुर"
	case "FRIENDS COLONY":
		sectionNameHindi = "फ्रैण्‍डस कालोनी"
	case "FULAVAD":
		sectionNameHindi = "फुलवाड"
	case "FULSANDA GANGADAS":
		sectionNameHindi = "फुलसन्‍दा गंगादास"
	case "FULSANDA HEERA":
		sectionNameHindi = "फुलसन्‍दा हीरा"
	case "FULSANDA KHAKAM":
		sectionNameHindi = "फुलसन्‍दा खाकम"
	case "FULSUNGHA":
		sectionNameHindi = "फुलसुंघा"
	case "FULSUNGHI":
		sectionNameHindi = "फुलसुंघी"
	case "FUTAMAHAL MAY JAIN MANDIR":
		sectionNameHindi = "फूटामहल मय जैन मन्दिर"
	case "GADA":
		sectionNameHindi = "गदा"
	case "GADAIYA NIKAT HARSONAGLA":
		sectionNameHindi = "गदईया नि0 हरसूनगला"
	case "GADAIYA NIKAT HINGANAGLA":
		sectionNameHindi = "गदईया नि0 हींगानगला"
	case "GADAL":
		sectionNameHindi = "गादल"
	case "GADARIYAN":
		sectionNameHindi = "गडरियान"
	case "GADDI NAGLI- 2":
		sectionNameHindi = "गददीनगली 2"
	case "GADDINAGLI-1":
		sectionNameHindi = "गददीनगली 1"
	case "GADHI":
		sectionNameHindi = "गढी"
	case "GADI KHANA ANSHIK KATGHAR":
		sectionNameHindi = "गाडी खाना आंशिक कटघर"
	case "GADIKHANA KATGHAR":
		sectionNameHindi = "गाडीखाना कटघर"
	case "GADIPUR":
		sectionNameHindi = "गद़़ीपुर"
	case "GADIYA NASEEMGANJ":
		sectionNameHindi = "गदईया नसीमगंज"
	case "GADLA":
		sectionNameHindi = "गादला"
	case "GADMAR PATTI TIKA SINGH":
		sectionNameHindi = "गदमर पटटी मोती सिंह"
	case "GADMARPATTI TIKA SINGH":
		sectionNameHindi = "गदमर पटटी टीका सिंह"
	case "GADRIYAN":
		sectionNameHindi = "गडरियान"
	case "GADUBALA":
		sectionNameHindi = "गदूबाला"
	case "GAGAN KI MILAK":
		sectionNameHindi = "गागन की मिलक"
	case "GAGAN NAGALA":
		sectionNameHindi = "गांगन नगला"
	case "GAGNPUR":
		sectionNameHindi = "गगनपुर"
	case "GAIBLIPUR":
		sectionNameHindi = "गैबलीपुर"
	case "GAJARAULA B.A":
		sectionNameHindi = "गजरौला बी0ए"
	case "GAJARAULA JAY SINGH":
		sectionNameHindi = "गजरौला जय सिंह"
	case "GAJEJJA":
		sectionNameHindi = "गजेजा"
	case "GAJGOLA NANALBADI":
		sectionNameHindi = "गजगोला नानकवाडी"
	case "GAJHEDA ALAM":
		sectionNameHindi = "गझेड़ा आलम"
	case "GAJIPUR AN0":
		sectionNameHindi = "गाजीपुर आं0"
	case "GAJIPUR BHAGWAN":
		sectionNameHindi = "गाजीपुर भगवान"
	case "GAJIPUR HIDAYAT URF GARHIBAN":
		sectionNameHindi = "गाजीपुर हिदायत उर्फ गढीबान"
	case "GAJIPUR":
		sectionNameHindi = "गाजीपुर"
	case "GAJIPURSABHACHAND":
		sectionNameHindi = "गांजीपुरसभाचन्‍द"
	case "GAJRAULA ACHPAL":
		sectionNameHindi = "गजरौला अचपल"
	case "GAJRAULA SHIV":
		sectionNameHindi = "गजरौला शिव"
	case "GAJRAULA":
		sectionNameHindi = "गजरौला"
	case "GAJROLA SED":
		sectionNameHindi = "गजरौला सैद"
	case "GAJRULA":
		sectionNameHindi = "गजरौला"
	case "GAJUPUR":
		sectionNameHindi = "गजूपुर"
	case "GAJUPURA":
		sectionNameHindi = "गजुपुरा"
	case "GAKKHARPUR":
		sectionNameHindi = "गक्‍खरपुर"
	case "GALI MIYAN SAHAB":
		sectionNameHindi = "गली मियां साहब"
	case "GALI NO. 1 JAMA MASJID":
		sectionNameHindi = "गली नं0 1 जामा मस्जिद"
	case "GALIBALIPUR MALUK":
		sectionNameHindi = "गालीबअलीपुर मलूक"
	case "GALPURA":
		sectionNameHindi = "गलपुरा"
	case "GALSHAHEED PURVI ANSHIK":
		sectionNameHindi = "गलशहीद पूर्वी आंशिक"
	case "GAMMANAPURA":
		sectionNameHindi = "गम्‍मनपुरा"
	case "GANDA JUD":
		sectionNameHindi = "गैन्‍डा जूड"
	case "GANDHI ASHRAM":
		sectionNameHindi = "गांधी आश्रम"
	case "GANDHI NAGAR ANSHIK":
		sectionNameHindi = "गांधी नगर आंशिक"
	case "GANDHINAGAR":
		sectionNameHindi = "गांधीनगर"
	case "GANDHOR A.":
		sectionNameHindi = "गन्‍धौर आ०"
	case "GANESH GHAT AANSHIK":
		sectionNameHindi = "गनेश घाट आंशिक"
	case "GANESH GHAT AI0":
		sectionNameHindi = "गनेश घाट ऐ0"
	case "GANESHPUR DEVI AITMALI (GAIRA)":
		sectionNameHindi = "गनेशपुर देवी ऐतमाली (गैरा)"
	case "GANGA KHEDI":
		sectionNameHindi = "गांगा खेडी"
	case "GANGADASPUR":
		sectionNameHindi = "गंगदासपुर"
	case "GANGAN NAGLI":
		sectionNameHindi = "गांगन नगली"
	case "GANGAPUR JADID":
		sectionNameHindi = "गंगापुर जदीद"
	case "GANGAPUR KADIM-1":
		sectionNameHindi = "गंगापुर कदीम-1"
	case "GANGAPUR KADIM-2":
		sectionNameHindi = "गंगापुर कदीम-2"
	case "GANGAPUR KALN":
		sectionNameHindi = "गंगापुर कलां"
	case "GANGAPURSHARKI":
		sectionNameHindi = "गंगापुर शर्की"
	case "GANGARAMWALA":
		sectionNameHindi = "गंगारामवाला"
	case "GANGODA JAT":
		sectionNameHindi = "गंगोडा जट"
	case "GANGODA SHEKH":
		sectionNameHindi = "गंगोडा शेख"
	case "GANGOI KHADAR":
		sectionNameHindi = "गगोई खादर"
	case "GANGU NAGLA A.":
		sectionNameHindi = "गांगू नंगला आ०"
	case "GANGU NAGLI GAR. AB.AD":
		sectionNameHindi = "गांगू नंगली गैर आं0"
	case "GANGUWALA":
		sectionNameHindi = "गांगूवाला"
	case "GANJ KADIM -2":
		sectionNameHindi = "गंज कदीम -2"
	case "GANJAKADIM -1":
		sectionNameHindi = "गंजकदीम -1"
	case "GANORA":
		sectionNameHindi = "गनौरा"
	case "GARABPUR":
		sectionNameHindi = "गारबपुर"
	case "GARAVPUR":
		sectionNameHindi = "गारवपुर"
	case "GARDAI KHEDA":
		sectionNameHindi = "गर्दइ खेडा"
	case "GARHMALPUR":
		sectionNameHindi = "गढ़मलपुर"
	case "GAUHARPUR SULTANPUR":
		sectionNameHindi = "गौहरपुर सुल्‍तानपुर"
	case "GAUKHANA KADIM":
		sectionNameHindi = "गऊखाना कदीम"
	case "GAUSPUR TOPARI":
		sectionNameHindi = "गौसपुर टोपरी"
	case "GAVRI BUJURG":
		sectionNameHindi = "गाबडी बुजुर्ग"
	case "GAYATRI NAGAR ANSHIK":
		sectionNameHindi = "गायत्री नगर आंशिक"
	case "GAZIPUR AN":
		sectionNameHindi = "गाजीपुर आं०"
	case "GAZIPUR":
		sectionNameHindi = "गाजीपुर"
	case "GAZROLA":
		sectionNameHindi = "गजरौला"
	case "GEHLUIYA":
		sectionNameHindi = "गहलुईया"
	case "GEHNI":
		sectionNameHindi = "गहनी"
	case "GHADI":
		sectionNameHindi = "गढी"
	case "GHAIR AJAM KHAN":
		sectionNameHindi = "घेर आजम खां"
	case "GHAIR ALLI KHA":
		sectionNameHindi = "घेर अल्ली खा"
	case "GHAIR BAKHSHI":
		sectionNameHindi = "घेर बख्शी"
	case "GHAIR DARIYA KHA":
		sectionNameHindi = "घेर दरिया खा"
	case "GHAIR GILJIYAN":
		sectionNameHindi = "घेर गिल्जियान"
	case "GHAIR KUTVI MIYAN":
		sectionNameHindi = "घेर कुतवी मियां"
	case "GHAIR MILKIYAN":
		sectionNameHindi = "घेर मिल्कियान"
	case "GHAIR MULLA MALUK":
		sectionNameHindi = "घेर मुल्ला मलूक"
	case "GHAIR MUNSHIVALA MAY CHAH SOTIYAN":
		sectionNameHindi = "घेर मुंशीवाला मय चाह सोतियान"
	case "GHAIR SAIFAL KHAN":
		sectionNameHindi = "घेर सैफल खां"
	case "GHAIR SHAH MOHAMMAD KHAN":
		sectionNameHindi = "घेर शाह मोहम्मद खां"
	case "GHAIR SHAIRFUDDIN KHAN":
		sectionNameHindi = "घेर शर्फुददीन खा"
	case "GHAKKARPUR":
		sectionNameHindi = "गक्‍खरपुर"
	case "GHANASHYAMAPUR":
		sectionNameHindi = "घनश्यामपुर"
	case "GHANGHERI":
		sectionNameHindi = "घंघेडी"
	case "GHANSURPUR":
		sectionNameHindi = "धंसूरपुर"
	case "GHANSURPURAMROLI":
		sectionNameHindi = "घनसूरपुरअमरोली"
	case "GHARAMPUR":
		sectionNameHindi = "घारमपुर"
	case "GHASIWALA COLLONY":
		sectionNameHindi = "घासीवाला"
	case "GHASIWALA":
		sectionNameHindi = "घासीवाला"
	case "GHASMANDI":
		sectionNameHindi = "घासमण्‍डी"
	case "GHASSIWALA":
		sectionNameHindi = "धासीवाला"
	case "GHER AJIM KHA MAY AKHUNAJADA":
		sectionNameHindi = "घेर अजीम खा मय आखूनजादा"
	case "GHER ALAF KHAN":
		sectionNameHindi = "घ़ेर अलफ खां"
	case "GHER BAJ KHA MAY FILAVALAN-1":
		sectionNameHindi = "घेर बाज खा मय फीलवालान-1"
	case "GHER BAJ KHA MAY FILAVALAN-2":
		sectionNameHindi = "घेर बाज खा मय फीलवालान-2"
	case "GHER BAJ KHA MAY FILAVALAN-3":
		sectionNameHindi = "घेर बाज खा मय फीलवालान-3"
	case "GHER BECHA KHAN":
		sectionNameHindi = "घेर बेचा खां"
	case "GHER CHANDAN KHAN MAY JEL K VATARS":
		sectionNameHindi = "घेर चन्‍दन खां चन्दन खां मय जेल क् वाटर्स"
	case "GHER FARUKHSHAH KHAN":
		sectionNameHindi = "घेर फरूखशाह खां"
	case "GHER FATEH KHAN":
		sectionNameHindi = "घेर फतह खां"
	case "GHER GAUS MOHAMMAD KHA":
		sectionNameHindi = "घेर गौस मोहम्मद खा"
	case "GHER HASSAN KHA":
		sectionNameHindi = "घेर हस्सन खा"
	case "GHER INAYAT KHAN":
		sectionNameHindi = "घेर इनायत खां"
	case "GHER JANAS KHAN":
		sectionNameHindi = "घेर जानस खां"
	case "GHER JIYAUNNAVI KHA":
		sectionNameHindi = "घेर जियाउन्नवी खा"
	case "GHER KALANDAR KHAN-1":
		sectionNameHindi = "घेर कलन्दर खां-1"
	case "GHER KALANDAR KHAN-2":
		sectionNameHindi = "घेर कलन्दर खां-2"
	case "GHER KATE BAJ KHAN":
		sectionNameHindi = "घेर कटे बाज खां"
	case "GHER MARDAN KHAN":
		sectionNameHindi = "घेर मर्दान खां"
	case "GHER MEERBAZ KHAN MEY GHER HASSN KHAN-1":
		sectionNameHindi = "घेर मीरबाज खां मय घेर हस्‍सन खां 1"
	case "GHER MEERBAZ KHAN MEY GHER HASSN KHAN-2":
		sectionNameHindi = "घेर मीरबाज खां मय घेर हस्‍सन खां 2"
	case "GHER MIAN KHAN -2":
		sectionNameHindi = "घेर मियाँ खां -2"
	case "GHER MIAN KHAN -3":
		sectionNameHindi = "घेर मियाँ खां -3"
	case "GHER MIAN KHAN-1":
		sectionNameHindi = "घेर मियॉ खां-1"
	case "GHER MIYAN KHEL":
		sectionNameHindi = "घेर मियाँ खेल"
	case "GHER MOHABBAT KHA":
		sectionNameHindi = "घेर मोहब्बत खा"
	case "GHER MUBARAK SHAH KHAN":
		sectionNameHindi = "घेर मुबारक शाह खां"
	case "GHER MUHAB KHA":
		sectionNameHindi = "घेर मुहाब खा"
	case "GHER MULLA GAINRAT":
		sectionNameHindi = "घेर मुल्ला गैंरत"
	case "GHER MUNAVVAR KHAN":
		sectionNameHindi = "घेर मुनव्वर खा"
	case "GHER NAJJU KHAN":
		sectionNameHindi = "घेर नज्जू खां"
	case "GHER PIPALWALA-1":
		sectionNameHindi = "घेर पीपलवाला-1"
	case "GHER PIPALWALA-2":
		sectionNameHindi = "घेर पीपलवाला-2"
	case "GHER PURAN SINGH":
		sectionNameHindi = "घेर पूरन सिह"
	case "GHER REHMAT KHAN 1":
		sectionNameHindi = "घेर रहमत खां 1"
	case "GHER REHMAT KHAN 2":
		sectionNameHindi = "घेर रहमत खां 2"
	case "GHER SAIFUDADIN KHAN 2":
		sectionNameHindi = "घेर सैफुददीन खां 2"
	case "GHER SAIFUDADIN KHAN":
		sectionNameHindi = "घेर सैफुददीन खां 1"
	case "GHER SALAVAT KHAN":
		sectionNameHindi = "घ़ेर सलावत खां"
	case "GHER SARABDAL KHAN":
		sectionNameHindi = "घेर सरब्दाल खां"
	case "GHER TONGA-1":
		sectionNameHindi = "घेर तोंगा-1"
	case "GHER TONGA-2":
		sectionNameHindi = "घेर तोंगा-2"
	case "GHER TONGA-3":
		sectionNameHindi = "घेर तोगा-3"
	case "GHER USMAN KHA":
		sectionNameHindi = "धेर उस्मान खा"
	case "GHER YUSUF KHAN":
		sectionNameHindi = "घेर यूसुफ खाँ"
	case "GHONDA AANSHIK":
		sectionNameHindi = "घोन्डा आंशिक"
	case "GHOOSIPURA PURBI":
		sectionNameHindi = "घोसीपुरा (पूर्वी)"
	case "GHOOSIPURA":
		sectionNameHindi = "घोसीपुरा"
	case "GHOR SAKHI":
		sectionNameHindi = "घेर सखी"
	case "GHOSIYAN":
		sectionNameHindi = "घोसीयान"
	case "GHUDIYAPUR":
		sectionNameHindi = "घुडियापुर"
	case "GHYANPUR":
		sectionNameHindi = "ज्ञानपुर"
	case "GIDODA":
		sectionNameHindi = "गिन्‍दौडा"
	case "GILAPURA":
		sectionNameHindi = "गिलपुरा"
	case "GINDOURI":
		sectionNameHindi = "गिन्‍दौरी"
	case "GINNOR DEHMAFI":
		sectionNameHindi = "गिन्‍नोर देहमाफी"
	case "GIRDAWA SAHANPUR URF KHALAPAR":
		sectionNameHindi = "गिरदावा साहनपुर उर्फ खलापार"
	case "GIRDHARPUR":
		sectionNameHindi = "गिरधरपुर"
	case "GODHI ANSHIK":
		sectionNameHindi = "गोधी आंशिक"
	case "GODHI-1":
		sectionNameHindi = "गोधी-1"
	case "GODHI-2":
		sectionNameHindi = "गोधी-2"
	case "GOGALI":
		sectionNameHindi = "गोगली"
	case "GOGALPUR KASAM":
		sectionNameHindi = "गोकलपुर कासम"
	case "GOHAR ALI KHAN":
		sectionNameHindi = "गौहर अली खां"
	case "GOHARPUR":
		sectionNameHindi = "गोहरपुर"
	case "GOHAWER HALLU A0":
		sectionNameHindi = "गौहावर हल्‍लू आं0"
	case "GOHAWER JAIT A0":
		sectionNameHindi = "गौहावर जैत आ0"
	case "GOKAL NAGRI":
		sectionNameHindi = "गोकल नगरी"
	case "GOKALPUR SUNDAR":
		sectionNameHindi = "गोकलपुर सुन्‍दर"
	case "GOLAPANDEY ORF SANDLIPUR":
		sectionNameHindi = "गोलापाण्‍डे उर्फ सन्‍दलीपुर"
	case "GOONGANAGLA":
		sectionNameHindi = "गूंगानगला"
	case "GOPALPUR NATHTHA NANGLA ORF KOKARPUR":
		sectionNameHindi = "गोपालपुर नत्‍था नंगला उर्फ कोकरपुर"
	case "GOPALPUR":
		sectionNameHindi = "गोपालपुर"
	case "GOPIVALA":
		sectionNameHindi = "गोपीवाला"
	case "GOPIWALA":
		sectionNameHindi = "गोपीवाला"
	case "GORA BAZAR":
		sectionNameHindi = "गोरा बाजार"
	case "GORASHAHAGADH":
		sectionNameHindi = "गौराशाहगढ़"
	case "GOSIPUR":
		sectionNameHindi = "घोसीपुरा"
	case "GOSPUR KHALSA":
		sectionNameHindi = "गौसपुर खालसा"
	case "GOSPUR RAI GAIR ABAD":
		sectionNameHindi = "गौसपुर राय गैर आबाद"
	case "GOSPUR SADAT":
		sectionNameHindi = "गौसपुर सादात"
	case "GOSPUR":
		sectionNameHindi = "गौसपुर"
	case "GOT ANSHIK":
		sectionNameHindi = "गोट आंशिक"
	case "GOUSAMPUR -2":
		sectionNameHindi = "गौसमपुर -2"
	case "GOUSAMPUR 01":
		sectionNameHindi = "गौसमपुर 01"
	case "GOVARDHANPUR":
		sectionNameHindi = "गोवर्धनपुर"
	case "GOVEDHANPUR":
		sectionNameHindi = "गोवरधनपुर"
	case "GOVIND NAGAR ANSHIK":
		sectionNameHindi = "गोविन्‍द नगर आंशिक"
	case "GOVIND PURA NEAR PADAMPUR":
		sectionNameHindi = "गोविन्‍दपुरा नि0 पदमपुर"
	case "GOVINDNAGAR MUHAMMADNAGAR":
		sectionNameHindi = "गोविन्‍दनगर मौहम्‍मदनगर"
	case "GOVINDNAGAR":
		sectionNameHindi = "गोविन्‍दनगर"
	case "GOVINDPUR GAIR AB.AD":
		sectionNameHindi = "गोविन्‍दपुर गैर आं0"
	case "GOVIN‍DPUR KALA":
		sectionNameHindi = "गोविन्‍दपुर कला"
	case "GOVINDPUR":
		sectionNameHindi = "गोविन्‍दपुर"
	case "GOVINDPURA NEAR RAYPUR":
		sectionNameHindi = "गोविन्‍दपुरा नि0 रायपुर"
	case "GOVINDPURA":
		sectionNameHindi = "गोविन्‍दपुरा"
	case "GOVLI A.":
		sectionNameHindi = "गोवली आ०"
	case "GOVLI B.A":
		sectionNameHindi = "गोवली बी ए"
	case "GOWARDHANPUR AN0":
		sectionNameHindi = "गोवधर्नपुर आ0"
	case "GOWARDHANPUR NANGLI":
		sectionNameHindi = "गोवर्धनपुर नंगली"
	case "GOYAL NAGAR MANPUR":
		sectionNameHindi = "गोयल नगर मानपुर"
	case "GOYALA B.A.":
		sectionNameHindi = "गोयला बी0ए0"
	case "GOYALA":
		sectionNameHindi = "गोयला ए0"
	case "GUIYAN BAGH ANSHIK":
		sectionNameHindi = "गुईयां बाग आंशिक"
	case "GUIYAN BAGH":
		sectionNameHindi = "गुईयां बाग"
	case "GUJAR TOLA":
		sectionNameHindi = "गूजर टोला"
	case "GUJARATIYAN":
		sectionNameHindi = "गुजरातियान"
	case "GUJARPUR ASU GAIR ABAD":
		sectionNameHindi = "गुजरपुर आसू गैर आबाद"
	case "GUJARPUR JASPAL":
		sectionNameHindi = "गुजरपुर जसपाल"
	case "GUJRAILA":
		sectionNameHindi = "गुजरैला"
	case "GUJRATI":
		sectionNameHindi = "गुजराती"
	case "GULABWADI ANSHIK":
		sectionNameHindi = "गुलाबवाडी आंशिक"
	case "GULAD PEEPALSANA-1":
		sectionNameHindi = "गूलडपीपलसाना 1"
	case "GULAD PEEPALSANA-2":
		sectionNameHindi = "गूलडपीपलसाना 2"
	case "GULADIA":
		sectionNameHindi = "गुलडिया"
	case "GULADIYA BHAT":
		sectionNameHindi = "गुलडिया भाट"
	case "GULADIYA BIJLA":
		sectionNameHindi = "गुलडिया बीजला"
	case "GULADIYA MAFI":
		sectionNameHindi = "गुलडिया माफी"
	case "GULADIYA MURAD":
		sectionNameHindi = "गुलडिया मुराद"
	case "GULADIYA TEULA":
		sectionNameHindi = "गुलडिया टयूला"
	case "GULALPUR":
		sectionNameHindi = "गुलालपुर"
	case "GULALWALI":
		sectionNameHindi = "गुलालवाली"
	case "GULAMGANJ":
		sectionNameHindi = "गुलामगंज"
	case "GULARIYA KALAN":
		sectionNameHindi = "गुलडिया कलां"
	case "GULI TALAB":
		sectionNameHindi = "गुली तालाब"
	case "GULISTAN":
		sectionNameHindi = "गुलिस्‍तान"
	case "GULZAR NAGAR":
		sectionNameHindi = "गुलजार नगर"
	case "GUNIYAKHEDI":
		sectionNameHindi = "गुनियाखेडी ."
	case "GUNIYAPUR":
		sectionNameHindi = "गुनियापुर"
	case "GURADASAPUR HIRA":
		sectionNameHindi = "गुरदासपुर हीरा"
	case "GURER AANSHIK":
		sectionNameHindi = "गुरेर आंशिक"
	case "GURETHA":
		sectionNameHindi = "गुरेठा"
	case "GURHA":
		sectionNameHindi = "गूढा"
	case "GURHATTI":
		sectionNameHindi = "गुरहटटी"
	case "GUSAIWALA":
		sectionNameHindi = "गुसाईवाला"
	case "GUTABLI MUO":
		sectionNameHindi = "गुतावली मु0"
	case "GVAR KHEDA":
		sectionNameHindi = "ग्वार खेड़ा"
	case "GVARAU":
		sectionNameHindi = "ग्वारऊ"
	case "GWALKHEDA":
		sectionNameHindi = "ग्वालखेडा"
	case "GYANI WALI BASTI ANSHIK":
		sectionNameHindi = "ज्ञानी वाली बस्‍ती आंशिक"
	case "GYANI WALI BASTI":
		sectionNameHindi = "ज्ञानी वाली बस्‍ती"
	case "GYASPUR":
		sectionNameHindi = "ग्‍यासपुर"
	case "HADAYPUR":
		sectionNameHindi = "हदयपुर"
	case "HADIPUR SADRUDDIN":
		sectionNameHindi = "हादीपुर सदरूदीन"
	case "HAFIJABAD SHAKI AN.":
		sectionNameHindi = "हाफिजाबाद शर्की आ0"
	case "HAFIJNAGAR":
		sectionNameHindi = "हाफिजनगर"
	case "HAIBATPUR PEERU":
		sectionNameHindi = "हैबतपुर पीरू"
	case "HAIGERPUR BHATT":
		sectionNameHindi = "हैजरपुर भटट"
	case "HAIJARPUR VEERCHAND":
		sectionNameHindi = "हैजरपुर वीरचन्‍द"
	case "HAJI KA KUA ASALATPURA":
		sectionNameHindi = "हाजी का कुआ असालतपुरा"
	case "HAJI MADNI KI GALI ASALATPURA":
		sectionNameHindi = "हा0 मदनी की गली असालतपुरा"
	case "HAJI NAGLA":
		sectionNameHindi = "हाजी नगला"
	case "HAJI YUSUF WALI GALI PEERZADA":
		sectionNameHindi = "हाजी यूसुफ वाली गली पीरजादा"
	case "HAJIMODPUR KOTKADAR 111 TO END":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर 111 से अंत तक"
	case "HAJIMODPUR KOTKADAR AN0 H.NO 1 TO 326":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर आ0 म0स0 1 से 326तक"
	case "HAJIMODPUR KOTKADAR H.NO 1 TO 439":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर म0स0 1 से 439 तक"
	case "HAJIMODPUR KOTKADAR H.NO 1 TO 923":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर म0स0 1 से 923 तक"
	case "HAJIMODPUR KOTKADAR H.NO 440 TO END":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर म0स0 440 से अंत तक"
	case "HAJIMOHD PUR KOTKADAR AN 1 TO 110":
		sectionNameHindi = "हाजीमौ0पुर उ र्फ कोटकादर १ से ११० तक"
	case "HAJINAGAR":
		sectionNameHindi = "हाजीनगर"
	case "HAJIPUR":
		sectionNameHindi = "हाजीपुर"
	case "HAJJAMAN":
		sectionNameHindi = "हज्‍जामान"
	case "HAJRATNAGAR A.":
		sectionNameHindi = "हजरतनगर आं0"
	case "HAJRATPUR":
		sectionNameHindi = "हजरतपुर"
	case "HAKEEMAN":
		sectionNameHindi = "हकीमान"
	case "HAKEEMPUR CHANDAN":
		sectionNameHindi = "हकीमपुर चान्‍दन"
	case "HAKEEMPUR DULLA G.AB.":
		sectionNameHindi = "हकीमपुर दल्‍ला गै0आ0"
	case "HAKEEMPUR MEGHA":
		sectionNameHindi = "हकीमपुर मेघा"
	case "HAKEEMPUR NARAYAN":
		sectionNameHindi = "हकीमपुर नरायन"
	case "HAKEEMPURA":
		sectionNameHindi = "हकीमपुरा"
	case "HAKEEPURA DAKSHIRIN":
		sectionNameHindi = "हकीमपुरा दक्षिणी"
	case "HAKIKATPUR GANGWALI AN.H.NO.1-209":
		sectionNameHindi = "हकीकतपुरगंगवाली आं०म०सं०१ से २०९"
	case "HAKIKATPUR GOVIND":
		sectionNameHindi = "हकीककतपुगोविन्‍द"
	case "HAKIKATPUR MUTHRA":
		sectionNameHindi = "हकीकतपुर मुथरा"
	case "HAKIKATPUR PRAYAG":
		sectionNameHindi = "हकीकतपुर प्रयाग"
	case "HAKIKATPUR SEHSU":
		sectionNameHindi = "हकीकतपुर सहसू"
	case "HAKIKATPUR VERCHAND":
		sectionNameHindi = "हकीकतपुरवीरचन्‍द"
	case "HAKIKATPURGANGWALI AN.":
		sectionNameHindi = "हकीकतपुर गंगवाली आं०"
	case "HAKIKATPURGANGWALI AN.H.NO210-END":
		sectionNameHindi = "हकीकतपुर गंगवाली आं०म०सं०२१० से अंत तक"
	case "HAKIMAAN 01":
		sectionNameHindi = "हकीमान ०१"
	case "HAKIMAAN 02":
		sectionNameHindi = "हकीमान 02"
	case "HAKIMAN AN.":
		sectionNameHindi = "हकीमान आं0"
	case "HAKIMAN":
		sectionNameHindi = "हकीमान"
	case "HAKIMGANJ":
		sectionNameHindi = "हकीमगंज"
	case "HAKIMPUR DISONDI":
		sectionNameHindi = "हकीमपुर दिसौन्‍दी"
	case "HAKIMPUR DURG":
		sectionNameHindi = "हकीमपुर दुर्ग"
	case "HAKIMPUR KAJI":
		sectionNameHindi = "हकीमपुर काजी"
	case "HAKIMPUR SHAKARGANJ AN.":
		sectionNameHindi = "हकीमपुर शकरगंज आ0"
	case "HAKIMPUR":
		sectionNameHindi = "हकीमपुर"
	case "HAKUMATPUR KESHO":
		sectionNameHindi = "हकूमतपुर केशो"
	case "HAKUMATPUR":
		sectionNameHindi = "हकुमतपुर"
	case "HALA NGALA":
		sectionNameHindi = "हला नगला"
	case "HALAVAIYAN":
		sectionNameHindi = "हलवाईयान"
	case "HALDUA MAFI":
		sectionNameHindi = "हल्‍दुआ माफी"
	case "HALDUKHATA":
		sectionNameHindi = "हल्‍दूखाता"
	case "HALLANAGLA":
		sectionNameHindi = "हल्‍ला नंगला"
	case "HALLOWALI":
		sectionNameHindi = "हल्‍लोवाली"
	case "HALUNAAGAR":
		sectionNameHindi = "हलूनागर"
	case "HALWAIYAN":
		sectionNameHindi = "हल्‍वाईयान"
	case "HAMA NANGLI":
		sectionNameHindi = "हामा नंगली"
	case "HAMAJAPUR":
		sectionNameHindi = "हमजापुर"
	case "HAMEEDPUR":
		sectionNameHindi = "हमीदपुर"
	case "HAMEERPUR":
		sectionNameHindi = "हमीरपुर"
	case "HAMIDABAD-1":
		sectionNameHindi = "हामिदाबाद-1"
	case "HAMIDABAD-2":
		sectionNameHindi = "हामिदाबाद-2"
	case "HAMIDNAGAR":
		sectionNameHindi = "हामिदनगर"
	case "HAMIDPUR GANGA GAIR ABAD":
		sectionNameHindi = "हामिदपुर गंगा गैर आबाद"
	case "HAMIDPUR MAKHAN":
		sectionNameHindi = "हामिदपुर माखन"
	case "HAMIDPUR VIDHICHAND GAIR ABAD":
		sectionNameHindi = "हामिदपुरविधिचन्‍द गैर आबाद"
	case "HAMIR PUR -3":
		sectionNameHindi = "हमीरपुर -3"
	case "HAMIRAPUR":
		sectionNameHindi = "हमीरपुर"
	case "HAMIRPUR-1":
		sectionNameHindi = "हमीरपुर-1"
	case "HAMIRPUR-2":
		sectionNameHindi = "हमीरपुर -2"
	case "HAMIRPUR-4":
		sectionNameHindi = "हमीरपुर -4"
	case "HAMIRPUR-5":
		sectionNameHindi = "हमीरपुर-5"
	case "HAMIRPUR-6":
		sectionNameHindi = "हमीरपुर-6"
	case "HAMJAPURKATHAUR":
		sectionNameHindi = "हमजापुर कठैर"
	case "HANSUPURA":
		sectionNameHindi = "हंसुपुरा"
	case "HANUMAN NAGAR LINEPAR":
		sectionNameHindi = "हनुमान नगर लाइनपार"
	case "HARAIYA KURD":
		sectionNameHindi = "हरैया खुर्द"
	case "HARAURA":
		sectionNameHindi = "हरौरा"
	case "HARBALLABHPUR GAIR AB.AD":
		sectionNameHindi = "हरबलभपुर गैर आं0"
	case "HARBANSHPUR DHARAM":
		sectionNameHindi = "हरबन्‍श्‍ापुर धारम"
	case "HARCHANDPUR AN0":
		sectionNameHindi = "हरचन्‍दपुर आं0"
	case "HARDASPUR GADI ZAFRABAD PRATHAM":
		sectionNameHindi = "हरदासपुर गढी जाफराबाद प्रथम"
	case "HARDASPUR GAIR AB.AD":
		sectionNameHindi = "हरदासपुर गैर आं0"
	case "HARDASPUR KOTRA-1":
		sectionNameHindi = "हरदासपुर कोटरा 1"
	case "HARDASPUR KOTRA-2":
		sectionNameHindi = "हरदासपुर कोटरा 2"
	case "HARDASPUR MADE":
		sectionNameHindi = "हरदासपुर माडे"
	case "HARDASPUR":
		sectionNameHindi = "हरदासपुर"
	case "HARDUA":
		sectionNameHindi = "हरदुआ"
	case "HAREBALI AN0":
		sectionNameHindi = "हरेवली आ0"
	case "HARETA-1":
		sectionNameHindi = "हरेटा-1"
	case "HARETA-2":
		sectionNameHindi = "हरेटा -2"
	case "HAREWALI":
		sectionNameHindi = "हरेवली"
	case "HAREYYA KA MAJHRA":
		sectionNameHindi = "हरैया का मझरा"
	case "HAREYYA":
		sectionNameHindi = "हरैया"
	case "HARGANPUR H NO 1 TO 142":
		sectionNameHindi = "हरगनपुर म0स0 १ से १४२ तक"
	case "HARGANPUR H NO 142 TO END":
		sectionNameHindi = "हरगनपुर म0स0 १४२ से अंत तक"
	case "HARGANPUR":
		sectionNameHindi = "हरगनपुर"
	case "HARGAON CHANDAN AN.H.NO 1-130":
		sectionNameHindi = "हरगॉवचान्‍दन आं०म०सं०१ से १३०"
	case "HARGAON CHANDAN AN.H.NO 131-END":
		sectionNameHindi = "हरगॉवचान्‍द आं०म०सं०१३१से अंत तक"
	case "HARIJAN BASTI":
		sectionNameHindi = "हरिजन बस्‍ती"
	case "HARINOORPUR":
		sectionNameHindi = "हरिनूरपुर"
	case "HARIRAM ROAD VARD-13":
		sectionNameHindi = "हरिराम रोड वार्ड-13"
	case "HARISINGHKABHOGLA":
		sectionNameHindi = "हरिसिंह का भोगला"
	case "HARIYA KALAN":
		sectionNameHindi = "हरैया कलां"
	case "HARIYALE":
		sectionNameHindi = "हरियाल"
	case "HARIYANA":
		sectionNameHindi = "हरियाना"
	case "HARJIPUR":
		sectionNameHindi = "हरजीपुर"
	case "HARKISANPUR JANULABDEEN":
		sectionNameHindi = "हरि‍ककशनपुरजैनुलआबेदीन"
	case "HARKISHANPUR":
		sectionNameHindi = "हरि‍कशनपुर"
	case "HARNAGLA":
		sectionNameHindi = "हरनगला"
	case "HARNATHPUR":
		sectionNameHindi = "हरनाथपुर"
	case "HARPAL NAGAR":
		sectionNameHindi = "हरपाल नगर"
	case "HARPUR":
		sectionNameHindi = "हरपुर"
	case "HARRA AHAMADPUR JALAL":
		sectionNameHindi = "हर्रा अहमदपुर जलाल"
	case "HARSAINPUR":
		sectionNameHindi = "हरसैनपुर"
	case "HARSHWARA AN0":
		sectionNameHindi = "हर्षवाड़ा आं0"
	case "HARSONAGLA":
		sectionNameHindi = "हरसूनगला"
	case "HARTHALA ANSHIK":
		sectionNameHindi = "हरथला आंशिक"
	case "HARTHALA SONAKPUR MORADABAD":
		sectionNameHindi = "हरथला सोनकपुर मुरादाबाद"
	case "HARTHALA":
		sectionNameHindi = "हरथला"
	case "HARTHLA UTTARI ANSHIK":
		sectionNameHindi = "हरथला उत्‍तरी आंशिक"
	case "HARVANSGWALA":
		sectionNameHindi = "हरवंशवाला"
	case "HARVOLA WALA":
		sectionNameHindi = "हरवोला वाला"
	case "HASALIPUR MUTHRA":
		sectionNameHindi = "हसनलीपुर मुथरा"
	case "HASAMPUR GOPAL":
		sectionNameHindi = "हाशमपुर गोपाल"
	case "HASAMPUR":
		sectionNameHindi = "हाशमपुर"
	case "HASAN ALIPUR":
		sectionNameHindi = "हसन अलीपुर"
	case "HASAN GANJ":
		sectionNameHindi = "हसनगॅज"
	case "HASAN PUR UTTARI":
		sectionNameHindi = "हसनपुर उत्‍तरी"
	case "HASANGHADI":
		sectionNameHindi = "हसनगढी"
	case "HASANLIPUR BHOGAN":
		sectionNameHindi = "हसनलीपुरभोगन"
	case "HASANLIPUR HEERA":
		sectionNameHindi = "हसनलीपुर हीरा"
	case "HASANLIPUR KHANDU":
		sectionNameHindi = "हसनलीपुर खन्‍डू"
	case "HASANLIPURDHARMA":
		sectionNameHindi = "हसनअलीपुर धर्मा"
	case "HASANLIPURMAAN":
		sectionNameHindi = "हसनलीपुरमान"
	case "HASANPUR BILASPUR":
		sectionNameHindi = "हसनपुर बिलासपुर"
	case "HASANPUR BILLOCH":
		sectionNameHindi = "हसनपुर बिल्‍लौच"
	case "HASANPUR JAT":
		sectionNameHindi = "हसनपुर जट"
	case "HASANPUR KAZI":
		sectionNameHindi = "हसनपुर काजी"
	case "HASANPUR RUP":
		sectionNameHindi = "हसनपुर रुप"
	case "HASANPUR SHARKI":
		sectionNameHindi = "हसनपुर शर्की"
	case "HASANPUR":
		sectionNameHindi = "हसनपुर"
	case "HASANPURA":
		sectionNameHindi = "हसनपुरा"
	case "HASHAMPUR":
		sectionNameHindi = "हाशमपुर"
	case "HASHMINAGAR":
		sectionNameHindi = "हाशमीनगर"
	case "HASNAPUR PALKI":
		sectionNameHindi = "हसनपुर पालकी"
	case "HASNAPUR":
		sectionNameHindi = "हसनपुर"
	case "HASUPURA HARKISHANPUR A. H.N. 1 TO 220":
		sectionNameHindi = "हसुपुरा हरकशिनपुर आ0 म0नं0 १ से २२० तक"
	case "HASUPURA HARKISHANPUR A. H.N. 221 TO END":
		sectionNameHindi = "हसुपुरा हरकशिनपुर आ0 म0नं0 २२1 से अंत तक"
	case "HASUPURA HARKISHANPUR A. H.N. 244 TO END":
		sectionNameHindi = "हसुपुरा हरकशिनपुर आ0 म0नं0 २४४ से अंत तक"
	case "HASUPURA HARKISHANPUR A. H.N. 61 TO 243":
		sectionNameHindi = "हसुपुरा हरकशिनपुर आ0 म0नं0 ६१ से २४३ तक"
	case "HASUPURA HARKISHANPUR":
		sectionNameHindi = "हसुपुरा हरकिशनपुर"
	case "HASUPURI G.A.":
		sectionNameHindi = "हसुपुरी गै0आ0"
	case "HATA KALAN-3":
		sectionNameHindi = "खाता कलां-3"
	case "HATAMPUR":
		sectionNameHindi = "हातमपुर"
	case "HATHAT":
		sectionNameHindi = "हटहट"
	case "HATHI KHANA":
		sectionNameHindi = "हाथी खाना"
	case "HATHI MANDIR":
		sectionNameHindi = "हाथी मन्दिर"
	case "HATHIPUR BAHAUDDIN":
		sectionNameHindi = "हाथीपुर बहाउद्दीन"
	case "HATHIPUR CHANDU":
		sectionNameHindi = "हाथीपुर चन्दू"
	case "HATHIPUR CHITTU":
		sectionNameHindi = "हाथीपुर चित्तू"
	case "HAYAT NAGAR KATGHAR ANSHIK":
		sectionNameHindi = "हयात नगर कटघर आंशिक"
	case "HAYATNAGAR DAKSHNI":
		sectionNameHindi = "हयातनगर दक्षिणी"
	case "HAYATNAGAR":
		sectionNameHindi = "हयातनगर"
	case "HAYATPUR":
		sectionNameHindi = "हयातपुर"
	case "HAZARPUR KASWA":
		sectionNameHindi = "हैज्‍जरपुर कस्‍वा"
	case "HEEMPUR BUJURG A.":
		sectionNameHindi = "हीमपुर बुर्जग आ०"
	case "HEEMPUR DEEPA":
		sectionNameHindi = "हीमपुर दीपा"
	case "HEEMPUR MAKHDUM JAHAN":
		sectionNameHindi = "हीमपुर मखदूम जहॉ"
	case "HEEMPUR MANAK":
		sectionNameHindi = "हीमपुर मानक"
	case "HEERAPUR GUKAL":
		sectionNameHindi = "हीरापुर गोकल"
	case "HEERAPUR MALI":
		sectionNameHindi = "हीरापुर माली"
	case "HELALPUR":
		sectionNameHindi = "हिलालपुर"
	case "HEMPUR":
		sectionNameHindi = "हेमपुर"
	case "HERWELA":
		sectionNameHindi = "हरवेला"
	case "HESAMPUR":
		sectionNameHindi = "हिसामपुर"
	case "HIDAYATPUR AN0":
		sectionNameHindi = "हि‍दायतपुर आं0"
	case "HIMAUPUR RAI":
		sectionNameHindi = "हिमांयुपुर राय"
	case "HIMAYUPUR":
		sectionNameHindi = "हिमायूपुर"
	case "HIMGIRI COLONY ANSHIK":
		sectionNameHindi = "हिमगिरी कालौनी आंशिक"
	case "HIMMATGANJ":
		sectionNameHindi = "हिम्‍मतगंज"
	case "HIMMATPUR 01":
		sectionNameHindi = "हिम्‍मतपुर 01"
	case "HIMMATPUR 02":
		sectionNameHindi = "हिम्‍मतपुर 02"
	case "HIMMATPUR BELA":
		sectionNameHindi = "हिम्‍मतपुर बेला"
	case "HIMPUR PACHATPURA":
		sectionNameHindi = "हीमपुर पछातपुरा"
	case "HIMPUR PIRTHIYA":
		sectionNameHindi = "हीमपुर पिरथिया"
	case "HINDU CHAUDHARI":
		sectionNameHindi = "हिन्‍दू चौधरियान"
	case "HINDU KATERA":
		sectionNameHindi = "हिन्‍दू कटेरा"
	case "HINGANAGLA":
		sectionNameHindi = "हींगानगला"
	case "HIRAN KHEDA AI0":
		sectionNameHindi = "हिरन खेडा ऐ0"
	case "HIRANKHERA":
		sectionNameHindi = "हिरनखेडा"
	case "HIRANPURA":
		sectionNameHindi = "हिरनपुरा"
	case "HIRDERAMPUR":
		sectionNameHindi = "हिरदेरामपुर"
	case "HIRNAKHERI":
		sectionNameHindi = "हिरनाखेडी"
	case "HISAMAPUR":
		sectionNameHindi = "हिसामपुर"
	case "HOLI KA MAIDAN LINEPAR":
		sectionNameHindi = "होली का मैदान लाईनपार"
	case "HOLI":
		sectionNameHindi = "होली"
	case "HONSPURA":
		sectionNameHindi = "होंसपुरा"
	case "HOSPUR":
		sectionNameHindi = "हौसपुर"
	case "HUKAMPUR MACHIKI":
		sectionNameHindi = "हुकमपुर माचकी"
	case "HULSAN GANJ LAKDI FAJALPUR":
		sectionNameHindi = "हुलसन गंज लाकडी फाजलपुर"
	case "HUMAYUMPUR CHANDRAPAL":
		sectionNameHindi = "हिमायपुर चन्‍द्रपाल"
	case "HUMAYUPUR IDDU":
		sectionNameHindi = "हुमायुंपुर इददू"
	case "HURMATNAGAR TANDA -1":
		sectionNameHindi = "हुरमतनगर टाण्‍डा-1"
	case "HURMATNAGAR TANDA -2":
		sectionNameHindi = "हुरमतनगर टाण्‍डा-2"
	case "HURMATNAGAR":
		sectionNameHindi = "हुरमतनगर"
	case "HUSAIN GANJ BHUDA":
		sectionNameHindi = "हुसैनगंज भूडा"
	case "HUSAINABAD AN. H.NO.1-277":
		sectionNameHindi = "हुसैनाबाद आं०म०सं०१ से २७७त"
	case "HUSAINABAD AN.H-NO 278-END":
		sectionNameHindi = "हुसैनाबाद आं०म० सं०२७८ से ंअत तक"
	case "HUSAINABAD AN.H.NO 1-277":
		sectionNameHindi = "हुसैनाबाद आं०म०सं०१ से २७७"
	case "HUSAINAPUR PACHATAUR":
		sectionNameHindi = "हुसैनपुर पचतौर"
	case "HUSAINGANJ GADDINAGLI":
		sectionNameHindi = "हुसैनगंज गददीनगली"
	case "HUSAINGANJ NEAR JANAKPUR":
		sectionNameHindi = "हुसैनगंज नि0 जनकपुर"
	case "HUSAINGANJ":
		sectionNameHindi = "हुसैनगंज"
	case "HUSAINPUR HAMEED":
		sectionNameHindi = "हुसैनपुर हमीद"
	case "HUSAINPUR HAMIR":
		sectionNameHindi = "हुसैनपुर हमीर"
	case "HUSAINPUR MAKHANA":
		sectionNameHindi = "हुसैनपुर माखना"
	case "HUSAINPUR MEERA":
		sectionNameHindi = "हुसैनपुर मीरा"
	case "HUSAINPUR SANKAT SINGH":
		sectionNameHindi = "हुसैनपुर संकट सिंह"
	case "HUSAINPUR SULTAN":
		sectionNameHindi = "हुसैनपुर सुल्‍तान"
	case "HUSAINPUR TAPPA NANGAL":
		sectionNameHindi = "हुसैनपुर टप्‍पा नांगल"
	case "HUSAINPUR":
		sectionNameHindi = "हुसैनपुर"
	case "HUSANPUR KALA":
		sectionNameHindi = "हुसैनपुर कला"
	case "HUSANPUR KHASA P.":
		sectionNameHindi = "हुसैनुपुर खासा प०"
	case "HUSANPUR KHASA PURV":
		sectionNameHindi = "हुसैनुपुर खासा पूर्व"
	case "HUSENPUR KALA":
		sectionNameHindi = "हुसैनपुर कला"
	case "HUSHIANGUNJ NEAR MADARPUR":
		sectionNameHindi = "हुसैनगंज निकट मदारपुर"
	case "HYATNAGAR":
		sectionNameHindi = "हयातनगर"
	case "IBRAHEEMPUR LAL":
		sectionNameHindi = "इब्राहीमपुर लाल"
	case "IBRAHEEMPUR,":
		sectionNameHindi = "इब्राहीमपुर,"
	case "IBRAHIM MARKET":
		sectionNameHindi = "इब्राहिम मार्केट"
	case "IBRAHIMABAD":
		sectionNameHindi = "इब्राहीमाबाद"
	case "IBRAHIMAPUR NARAYAN":
		sectionNameHindi = "इब्राहीमपुर नारायण"
	case "IBRAHIMAPUR":
		sectionNameHindi = "इब्राहीमपुर"
	case "IBRAHIMPUR BAHAUDDIN":
		sectionNameHindi = "इब्राहीमपुर बहाउददीन"
	case "IBRAHIMPUR BAWAN":
		sectionNameHindi = "इब्राहीमपुर बावन"
	case "IBRAHIMPUR JAMALUDDIN":
		sectionNameHindi = "इब्राहीमपुर जमालुददीन"
	case "IBRAHIMPUR KHANDSAL":
		sectionNameHindi = "इब्राहीमपुर खण्‍डसाल"
	case "IBRAHIMPUR RAJU":
		sectionNameHindi = "इब्राहीमपुर राजू"
	case "IBRAHIMPUR SAHDO":
		sectionNameHindi = "इब्राहीमपुर साधो"
	case "IBRAHIMPUR URF KUMHARPURA":
		sectionNameHindi = "इब्राहीमपुर उर्फ कुम्‍हारपुरा"
	case "IDALPUR URF JOGI PURA":
		sectionNameHindi = "ईदलपुर उर्फ जोगीपुरा"
	case "IGRAH":
		sectionNameHindi = "इग्रह"
	case "ILAR RASULABAD":
		sectionNameHindi = "ईलर रसूलाबाद"
	case "ILAR":
		sectionNameHindi = "ईलर"
	case "ILAYACHIPUR KHADAGU":
		sectionNameHindi = "इलायचीपुर खडगू"
	case "IMALIA":
		sectionNameHindi = "इमलिया"
	case "IMAMABAD AN":
		sectionNameHindi = "इमामवाडा आं0"
	case "IMAMBADA -1":
		sectionNameHindi = "इमामबाडा-1"
	case "IMAMBADA -2":
		sectionNameHindi = "इमामबाडा-2"
	case "IMAMBADA -4":
		sectionNameHindi = "इमामबाडा-4"
	case "IMAMBADA -5":
		sectionNameHindi = "इमामबाडा-5"
	case "IMAMBADA-3":
		sectionNameHindi = "इमामबाडा-3"
	case "IMAMBARA":
		sectionNameHindi = "इमामबाडा"
	case "IMAMPURTIGRI":
		sectionNameHindi = "इमामपुर ति‍गरी"
	case "IMARATI NIKAT AKBRABAD":
		sectionNameHindi = "इमरता निकट अकबरावाद"
	case "IMARATPUR UDHO0 AANSHIK":
		sectionNameHindi = "इमरतपुर उधो0 आंशिक"
	case "IMARTI":
		sectionNameHindi = "इमरती"
	case "IMLI ASMAT KHAN":
		sectionNameHindi = "इमली अस्मत खां"
	case "IMLI BATANIYA":
		sectionNameHindi = "इमली बतनिया"
	case "IMLI JHULEWALI":
		sectionNameHindi = "इमली झूलेवाली"
	case "IMRATA RAY":
		sectionNameHindi = "इमरताराय"
	case "IMRATA-1":
		sectionNameHindi = "इमरता;1"
	case "IMRATA-2":
		sectionNameHindi = "इमरता-2"
	case "IMRATA-3":
		sectionNameHindi = "इमरता-3"
	case "IMRATA-4":
		sectionNameHindi = "इमरता-4"
	case "IMRATA-5":
		sectionNameHindi = "इमरता-5"
	case "IMRATAPUR FAKHARUDDINAPUR":
		sectionNameHindi = "इमरतपुर फखरुद्दीनपुर"
	case "IMRATPUR SYODARA":
		sectionNameHindi = "इमरतपुर स्योडारा"
	case "IMRATPUR":
		sectionNameHindi = "इमरतपुर"
	case "INAMPURA":
		sectionNameHindi = "इनामपुरा"
	case "INAYATNAGAR":
		sectionNameHindi = "इनायतनगर"
	case "INAYATPUR H.NO 1 TO END":
		sectionNameHindi = "इनायतपुर म0स0 १ से अंत तक"
	case "INAYATPUR":
		sectionNameHindi = "इनायतपुर"
	case "INCHAWALA":
		sectionNameHindi = "इन्‍छावाला"
	case "INDARPUR":
		sectionNameHindi = "इन्‍दरपुर"
	case "INDRA COLONY GULABWADI":
		sectionNameHindi = "इन्‍द्रा कालोनी गुलाबवाडी"
	case "INDRA-1":
		sectionNameHindi = "इन्‍ड्रा -1"
	case "INDRI":
		sectionNameHindi = "इन्‍ड्री"
	case "IRSAPUR":
		sectionNameHindi = "ईर्सापुर"
	case "ISAKHERA":
		sectionNameHindi = "ईसा खेडा"
	case "ISAMAIL NAGAR NEAR AHMADNAGAR":
		sectionNameHindi = "इस्‍माईलनगर नि0 अहमदनगर"
	case "ISAPUR":
		sectionNameHindi = "ईसापुर"
	case "ISHAKPUR":
		sectionNameHindi = "इसहाकपुर"
	case "ISLAM NAGAR AN0":
		sectionNameHindi = "इस्‍लामनगर आ0"
	case "ISLAM NAGAR ANSHIK DHEEMRI":
		sectionNameHindi = "इस्‍लाम नगर आंशिक धीमरी"
	case "ISLAMABAD AN0 1 TO END":
		sectionNameHindi = "इस्‍लामाबाद आ0 १ से अन्‍त तक"
	case "ISLAMABAD AN0 20-394":
		sectionNameHindi = "इस्‍लामाबाद आ0 20से 394 तक"
	case "ISLAMNAGAR A.":
		sectionNameHindi = "इस्‍लामनगर आं0"
	case "ISLAMNAGAR A.M.N. 1 TO 119":
		sectionNameHindi = "इस्‍लामनगर आं0 म0न0 1 से ११९ तक"
	case "ISLAMNAGAR A.M.N. 1":
		sectionNameHindi = "इस्‍लामनगर आं0 म0न0 1"
	case "ISLAMNAGAR A.M.N. 120 TO END":
		sectionNameHindi = "इस्‍लामनगर आं0 म0न0 १२० से अन्‍त तक"
	case "ISLAMNAGAR A.M.N. 82 TO END":
		sectionNameHindi = "इस्‍लामनगर आं0 म0न0 82 अन्‍त तक"
	case "ISLAMNAGAR AN.":
		sectionNameHindi = "इस्‍लामनगर आं0"
	case "ISLAMNAGAR RAMPURA":
		sectionNameHindi = "इस्‍लामनगर रम्‍पुरा"
	case "ISLAMPUR BEGA":
		sectionNameHindi = "इस्‍लामपुर बेगा"
	case "ISLAMPUR CHAMRA":
		sectionNameHindi = "इस्‍लामपुर चमरा"
	case "ISLAMPUR DAS":
		sectionNameHindi = "इस्‍लामपुर दास"
	case "ISLAMPUR DEEPA":
		sectionNameHindi = "इस्‍लामपुर दीपा"
	case "ISLAMPUR HATTU":
		sectionNameHindi = "इस्‍लामपुर हटटू"
	case "ISLAMPUR LALA":
		sectionNameHindi = "इस्‍लामपुर लाला"
	case "ISLAMPUR LALU":
		sectionNameHindi = "इस्‍लामपुर लालू"
	case "ISLAMPUR MEERA":
		sectionNameHindi = "इस्‍लामपुर मीरा"
	case "ISLAMPUR NEEMDAS":
		sectionNameHindi = "इस्‍लामपुर नीमदास"
	case "ISLAMPUR SADAT":
		sectionNameHindi = "इस्‍लामपुर सादात"
	case "ISLAMPUR SAHLI":
		sectionNameHindi = "इस्‍लामपुर शाहली"
	case "ISLAMPUR SAHU":
		sectionNameHindi = "इस्‍लामपुर साहू"
	case "ISLAMPUR THAMBU CHAU":
		sectionNameHindi = "इस्‍लामपुर थम्‍बू चाउ"
	case "ISLAMPUR VISHNOI":
		sectionNameHindi = "इस्‍लामपुर विश्‍नोई"
	case "ISLAMPUR":
		sectionNameHindi = "इस्‍लामपुर"
	case "ISMAIL ROAD ASALATPURA":
		sectionNameHindi = "इस्‍माईल रोड असालतपुरा"
	case "ISMAIL ROAD BEKRI WALI GALI":
		sectionNameHindi = "इस्‍माईल रोड बेकरी वाली गली"
	case "ISMAILPUR":
		sectionNameHindi = "इस्‍माईलपुर"
	case "ISMAILPURDAMI":
		sectionNameHindi = "ईस्‍माईलपुर दमी"
	case "ISMAILPURSALI BA":
		sectionNameHindi = "इस्‍माईलपुर शाहली बी0ए0"
	case "ISMAYILPUR A":
		sectionNameHindi = "इस्‍माईलपुर आ०"
	case "ISMAYILPUR":
		sectionNameHindi = "इस्‍माईलपुर"
	case "ISSAPUR":
		sectionNameHindi = "इस्‍सापुर"
	case "ISSEPUR AN0":
		sectionNameHindi = "इस्‍सेपुर आं0"
	case "ISWARPUR":
		sectionNameHindi = "ईश्‍वरपुर"
	case "ITANGA BERAMNAGAR":
		sectionNameHindi = "इटंगा बैरमनगर"
	case "ITAWA":
		sectionNameHindi = "इटावा"
	case "JABALPUR GUDAR AN0":
		sectionNameHindi = "जबलपुर गूदड़ आं0"
	case "JABERDASTPUR G.A.":
		sectionNameHindi = "जबरदस्‍तपुर गै0आ0"
	case "JABTA NAGAR":
		sectionNameHindi = "जाब्‍तानगर"
	case "JADOANPUR-2":
		sectionNameHindi = "जादौपुर-2"
	case "JADONPUR-1":
		sectionNameHindi = "जादौपुर-1"
	case "JADOPUR":
		sectionNameHindi = "जादोपुर"
	case "JAFARABAD KURAI":
		sectionNameHindi = "जाफराबाद कुरई"
	case "JAFARABAD":
		sectionNameHindi = "जाफराबाद"
	case "JAFARHUSENPUR":
		sectionNameHindi = "जाफरहुसैनपुर"
	case "JAFARPUR COT":
		sectionNameHindi = "जाफरपुर कोट"
	case "JAFARPUR PARASRAM":
		sectionNameHindi = "जाफरपुर परसराम"
	case "JAFARPUR":
		sectionNameHindi = "जाफरपुर"
	case "JAFRABAD":
		sectionNameHindi = "जाफराबाद"
	case "JAFRAPUR":
		sectionNameHindi = "जाफरपुर"
	case "JAFRULLAPUR":
		sectionNameHindi = "जफरूल्‍लापुर"
	case "JAGANNATHPUR":
		sectionNameHindi = "जगननाथपुर"
	case "JAGANNATHPURGOKAL":
		sectionNameHindi = "जगन्‍नाथपुरगोकल"
	case "JAGANNATHPURMEHRU":
		sectionNameHindi = "जगन्‍नाथपुर मेहरू"
	case "JAGANPUR URF JAGANBALA":
		sectionNameHindi = "जगनपुर उर्फ जगनवाला"
	case "JAGARAMPURA":
		sectionNameHindi = "जगरमपुरा"
	case "JAGATCHANDPUR":
		sectionNameHindi = "जगतचन्‍दपुर"
	case "JAGATPUR RAMRAI":
		sectionNameHindi = "जगतपुर रामराय"
	case "JAGATPUR":
		sectionNameHindi = "जगतपुर"
	case "JAGDEESHPUR":
		sectionNameHindi = "जगदीशपुर"
	case "JAGDISHPUR":
		sectionNameHindi = "जगदीशपुर"
	case "JAGESAR":
		sectionNameHindi = "जगेसर"
	case "JAGUPURA":
		sectionNameHindi = "जगुपुरा"
	case "JAHAGEERPUR CHEKFERI":
		sectionNameHindi = "जहांगीरपुर चकफेरी"
	case "JAHAGERPUR":
		sectionNameHindi = "जहांगीरपुर"
	case "JAHANABAD":
		sectionNameHindi = "जहानाबाद आं0"
	case "JAHANGEERPUR":
		sectionNameHindi = "जहांगीरपुर"
	case "JAHANGIRPUR LALU":
		sectionNameHindi = "जहागीरपुर लालू"
	case "JAHANGIRPUR":
		sectionNameHindi = "जहांगीरपुर"
	case "JAHANGURPURKHASS":
		sectionNameHindi = "जहांगीहीरपुरखास"
	case "JAHENABAD":
		sectionNameHindi = "जहानाबाद"
	case "JAHGEERPUR M.GAYNPUR URF ROOPPUR":
		sectionNameHindi = "जहांगीरपुर मजरा ज्ञानपुर उर्फ रूपपुर"
	case "JAHIDAPUR SIKMAPUR":
		sectionNameHindi = "जाहिदपुर सीकमपुर"
	case "JAHIRABAD G.A.":
		sectionNameHindi = "जहीराबाद गै0आ0"
	case "JAIDOLI":
		sectionNameHindi = "जैडोली"
	case "JAINABAD":
		sectionNameHindi = "जैनाबाद"
	case "JAINULAUDDINPUR":
		sectionNameHindi = "जैनुलाउद्वीनपुर"
	case "JAIRAMPUR":
		sectionNameHindi = "जयरामपुर"
	case "JAISINGH PUR":
		sectionNameHindi = "जयसिंह पुर"
	case "JAITAPUR":
		sectionNameHindi = "जैतपुर"
	case "JAITIYA FIROJAPUR":
		sectionNameHindi = "जैतिया फिरोजपुर"
	case "JAITIYA SADULLAPUR AANSHIK":
		sectionNameHindi = "जैतिया सादुल्लापुर आंशिक"
	case "JAITIYASADULLAPUR AANSHIK":
		sectionNameHindi = "जैतियासादुल्लापुर आंशिक"
	case "JAITOLI 01":
		sectionNameHindi = "जयतोली 01"
	case "JAITOLI 02":
		sectionNameHindi = "जयतोली 02"
	case "JAITPUR BISAHT AANSHIK":
		sectionNameHindi = "जैतपुर बिसाहट आंशिक"
	case "JAITPUR PATTI AANSHIK":
		sectionNameHindi = "जैतपुर पट्टी आंशिक"
	case "JALABPUR GUDAR AN0":
		sectionNameHindi = "जालबपुर गूदड़ आं0"
	case "JALAFNAGLA-1":
		sectionNameHindi = "जालफनगला 1"
	case "JALAFNAGLA-2":
		sectionNameHindi = "जालफनगला 2"
	case "JALALPUR A.":
		sectionNameHindi = "जलीलपुर आ०"
	case "JALALPUR AASRA":
		sectionNameHindi = "जलालपुर आसरा"
	case "JALALPUR BUCHA":
		sectionNameHindi = "जलालपुर बूचा"
	case "JALALPUR DIPA":
		sectionNameHindi = "जलालपुर दीपा"
	case "JALALPUR HASNA":
		sectionNameHindi = "जलालपुर हसना"
	case "JALALPUR KAZI":
		sectionNameHindi = "जलालपुर काजी"
	case "JALALPUR KHADAR":
		sectionNameHindi = "जलालपुर खादर"
	case "JALALPUR KHALSA":
		sectionNameHindi = "जलालपुर खालसा"
	case "JALALPUR KHAS":
		sectionNameHindi = "जलालपुर खास"
	case "JALALPUR SULTAN":
		sectionNameHindi = "जलालपुर सुल्‍तान"
	case "JALALPUR TURK":
		sectionNameHindi = "जलालपुर तुर्क"
	case "JALALPUR URF SHAIKPURA":
		sectionNameHindi = "जलालपुर उर्फ शेखपुरा"
	case "JALALPUR":
		sectionNameHindi = "जलालपुर"
	case "JALIF NAGLA-1":
		sectionNameHindi = "जालिफ नगला-1"
	case "JALIF NAGLA-2":
		sectionNameHindi = "जालिफ नगला-2"
	case "JALIF NAGLA-3":
		sectionNameHindi = "जालिफ नगला-3"
	case "JALPUR AN0":
		sectionNameHindi = "जालपुर आ0"
	case "JALPUR":
		sectionNameHindi = "जालपुर"
	case "JALPURA":
		sectionNameHindi = "जलपुरा"
	case "JAMA MASJID UTTRI ANSHIK":
		sectionNameHindi = "जामा मस्जिद उत्‍तरी आंशिक"
	case "JAMA MASJID":
		sectionNameHindi = "जामा मस्जिद"
	case "JAMALAPUR ALAM":
		sectionNameHindi = "जमालपुर आलम"
	case "JAMALAPUR MAHIPAT":
		sectionNameHindi = "जमालपुर महीपत"
	case "JAMALAPUR UDYACHAND AN.":
		sectionNameHindi = "जमालपुर उदयचंद आं0"
	case "JAMALLUDDINPUR":
		sectionNameHindi = "जमालुदीनपुर"
	case "JAMALPUR BANGAR":
		sectionNameHindi = "जमालपुर बांगर"
	case "JAMALPUR DHIKLI":
		sectionNameHindi = "जमालपुर ढीकली"
	case "JAMALPUR JATT":
		sectionNameHindi = "जमालपुर जट"
	case "JAMALPUR KHOKO":
		sectionNameHindi = "जमालपुर खोको"
	case "JAMALPUR KIRAT A.":
		sectionNameHindi = "जमालपुर कीरत आ0"
	case "JAMALPUR KIRAT":
		sectionNameHindi = "जमालपुर कीरत"
	case "JAMALPUR MAN":
		sectionNameHindi = "जमालपुर मान"
	case "JAMALPUR MUNDA NANGLA":
		sectionNameHindi = "जमालपुर मुण्‍डा नंगला"
	case "JAMALPUR PATHANI":
		sectionNameHindi = "जमालपुर पठानी"
	case "JAMALPUR SHAIKH":
		sectionNameHindi = "जमालपुर शेख"
	case "JAMALPUR":
		sectionNameHindi = "जमालपुर"
	case "JAMALUDDINPUR G.A.":
		sectionNameHindi = "जमालुददीनपुर गै0आ0"
	case "JAMANIYA KHURD":
		sectionNameHindi = "जमनिया खुर्द"
	case "JAMANPUR":
		sectionNameHindi = "जमनपुर"
	case "JAMAPUR":
		sectionNameHindi = "जमापुर"
	case "JAMASHEDAPUR":
		sectionNameHindi = "जमशेदपुर"
	case "JAMHIRI":
		sectionNameHindi = "जम्‍हीरी"
	case "JAMIN GANJ":
		sectionNameHindi = "जामिनगंज"
	case "JAMNA AZAM":
		sectionNameHindi = "जमना आजम"
	case "JAMNAVALA ANSIK":
		sectionNameHindi = "जमनावाला आंशिक"
	case "JANAKPUR":
		sectionNameHindi = "जनकपुर"
	case "JANULABEDDIN AN0":
		sectionNameHindi = "जैनुलआबेदीन आ0"
	case "JANUNAGAR":
		sectionNameHindi = "जनूनागर"
	case "JARGANV":
		sectionNameHindi = "जरगांव"
	case "JARUFSAJAN":
		sectionNameHindi = "जरूफसाजान"
	case "JASAMAUR":
		sectionNameHindi = "जसमौर"
	case "JASMOLI-1":
		sectionNameHindi = "जसमोली-1"
	case "JASMOLI-2":
		sectionNameHindi = "जसमोली-2"
	case "JASMOURA":
		sectionNameHindi = "जसमौरा"
	case "JASPUR":
		sectionNameHindi = "जसपुर"
	case "JASRATH PUR":
		sectionNameHindi = "जसरथ पुर"
	case "JASSUJOT":
		sectionNameHindi = "जस्‍सूजोत"
	case "JASWANTPUR AN0":
		sectionNameHindi = "जसवन्‍तपुर आं0"
	case "JATAN B-2":
		sectionNameHindi = "जाटान बी0 2"
	case "JATAN B-3":
		sectionNameHindi = "जाटान बी 3"
	case "JATAN B-4":
		sectionNameHindi = "जाटान बी 4"
	case "JATAN":
		sectionNameHindi = "जाटान"
	case "JATAPURA JHADU":
		sectionNameHindi = "जटपुरा झाड़ू"
	case "JATAV BASTI ASALATPURA DAKSHIN":
		sectionNameHindi = "जाटव बस्‍ती असालतपुरा दक्षिण"
	case "JATI AN H.NO 1-45":
		sectionNameHindi = "जाति आं०म०सं०१ से ४५तक"
	case "JATI AN H.NO 46-END":
		sectionNameHindi = "जती आं०म०सं०४६ से अन्‍त तक"
	case "JATOOLA":
		sectionNameHindi = "जटौला"
	case "JATPURA AANSHIK":
		sectionNameHindi = "जटपुरा आंशिक"
	case "JATPURA BONDA AN0":
		sectionNameHindi = "जटपुरा बौण्‍डा आं0"
	case "JATPURA KHAS":
		sectionNameHindi = "जटपुरा खास"
	case "JATPURA-1":
		sectionNameHindi = "जटपुरा-1"
	case "JATPURA-2":
		sectionNameHindi = "जटपुरा-2"
	case "JATPURA":
		sectionNameHindi = "जटपुरा"
	case "JATT NANGLA":
		sectionNameHindi = "जट नंगला"
	case "JAULAPUR":
		sectionNameHindi = "जौलपुर"
	case "JAURASI":
		sectionNameHindi = "जौरासी"
	case "JAWAHAR NAGAR":
		sectionNameHindi = "जवाहर नगर"
	case "JAY SINGH JOT":
		sectionNameHindi = "जय सिंह जोत"
	case "JAYANTIPUR ANSHIK":
		sectionNameHindi = "जयन्‍तीपुर आंशिक"
	case "JAYNAGAR":
		sectionNameHindi = "जयनगर"
	case "JEBDA":
		sectionNameHindi = "जेबडा"
	case "JEELAL MANDI BANS":
		sectionNameHindi = "जीलाल मण्‍डी बॉस"
	case "JEETPUR KHAS":
		sectionNameHindi = "जीतपुर खास"
	case "JEETPURA":
		sectionNameHindi = "जीतपुरा"
	case "JEVANPUR":
		sectionNameHindi = "जीवनपुर"
	case "JHABBU KA NALA":
		sectionNameHindi = "झब्‍बू का नाला"
	case "JHADEWALA":
		sectionNameHindi = "झाडेवाला"
	case "JHADPURA":
		sectionNameHindi = "झाडपुरा"
	case "JHAJHANPUR KANTH ROAD":
		sectionNameHindi = "झांझनपुर कांठ रोड"
	case "JHAKDA":
		sectionNameHindi = "झकडा"
	case "JHAKDI BANGAR":
		sectionNameHindi = "झकडी बंगर"
	case "JHAKKAKI":
		sectionNameHindi = "झक्‍काकी"
	case "JHAL":
		sectionNameHindi = "झाल"
	case "JHALRI":
		sectionNameHindi = "झलरी"
	case "JHANABAD AN 1 TO 875":
		sectionNameHindi = "जहानाबाद आ0 १ से ८७५"
	case "JHANDA AN":
		sectionNameHindi = "झण्‍डा आं०"
	case "JHANDA BADE PIR SAHAB-1":
		sectionNameHindi = "झण्डा बड़े पीर साहब-1"
	case "JHANDA BADE PIR SAHAB-2":
		sectionNameHindi = "झण्डा बड़े पीर साहब-2"
	case "JHANDAPUR":
		sectionNameHindi = "झण्‍डापुर"
	case "JHILMILA AN.H.NO 1-85":
		sectionNameHindi = "झिलमिला आं०म०सं०१से ८५ तक"
	case "JHILMILA AN.H.NO 86-END":
		sectionNameHindi = "झिलमिला आं०म०सं०८६से अंत तक"
	case "JHIRAN":
		sectionNameHindi = "झीरन"
	case "JHUJHAPURA URF NAIPURA":
		sectionNameHindi = "जुझारपुरा उर्फ नाईपुरा"
	case "JHULA":
		sectionNameHindi = "झूला"
	case "JHUNAIYA":
		sectionNameHindi = "झुनईया"
	case "JHURAKJHUNDI":
		sectionNameHindi = "झुरकझुन्‍डी"
	case "JIGANIYA":
		sectionNameHindi = "जिगनिया"
	case "JIGAR COLONY MORADABAD":
		sectionNameHindi = "जिगर कालौनी मुरादाबाद"
	case "JILEDARAN 01":
		sectionNameHindi = "जिलेदारान ०१"
	case "JILEDARAN 02":
		sectionNameHindi = "जिलेदारान 02"
	case "JINA INAYAT KHA":
		sectionNameHindi = "जीना इनायत खा"
	case "JITHANIYA JAGIR":
		sectionNameHindi = "जिठनिया जागीर"
	case "JITHANYA SHARKI":
		sectionNameHindi = "जिठनिया शर्की"
	case "JIVAN KI SARAI":
		sectionNameHindi = "जीवन की सराय"
	case "JIWAI JADID":
		sectionNameHindi = "जिवाई जदीद"
	case "JIWAI KADIM-1":
		sectionNameHindi = "जिवाई कदीम-1"
	case "JIWAI KADIM-2":
		sectionNameHindi = "जिवाई कदीम-2"
	case "JIYARAT SHIRIMIYAN-1":
		sectionNameHindi = "ज्‍यारत शीरीमियां-1"
	case "JIYARAT SHIRIMIYAN-2":
		sectionNameHindi = "ज्‍यारत शीरीमियां-2"
	case "JIYARAT SHIRIMIYAN-3":
		sectionNameHindi = "ज्‍यारत शीरीमियां-3"
	case "JIYARAT SHIRIMIYAN-4":
		sectionNameHindi = "ज्‍यारत शीरीमियां-4"
	case "JIYARAT SHIRIMIYAN-5":
		sectionNameHindi = "ज्‍यारत शीरी मियां-5"
	case "JIYARAT SHIRIMIYAN-6":
		sectionNameHindi = "ज्‍यारत शीरीमियां-6"
	case "JMALPUR":
		sectionNameHindi = "जमालपुर"
	case "JOGANPUR GAYATRI NAGAR":
		sectionNameHindi = "जोगनपुर गायत्री नगर"
	case "JOGIDASPUR GAIR AB.":
		sectionNameHindi = "जोगीदासपुर गैर आं0"
	case "JOGIONDA":
		sectionNameHindi = "जोगी औन्‍धा"
	case "JOGIPURA":
		sectionNameHindi = "जोगीपुरा"
	case "JOINPUR":
		sectionNameHindi = "जौनपुर"
	case "JOJHIYAN":
		sectionNameHindi = "झोझीयान"
	case "JOSHIYAN":
		sectionNameHindi = "जोशियान"
	case "JOTAHIMMA AN.":
		sectionNameHindi = "जोतहिम्‍मा आं0"
	case "JUJHELA":
		sectionNameHindi = "जूझैला"
	case "JUL DHAKIYA":
		sectionNameHindi = "जुल ढकिया"
	case "JULAHAN":
		sectionNameHindi = "जुलाहान"
	case "JUMERAT KA BAJAR.":
		sectionNameHindi = "जुमेरात का बजार"
	case "JUTHIYA-1":
		sectionNameHindi = "जुठिया-1"
	case "JUTHIYA-2":
		sectionNameHindi = "जुठिया-2"
	case "JUTHIYA-3":
		sectionNameHindi = "जुठिया-3"
	case "JVALANGAR AI0 D0 OII0 0 LE0KA0-5":
		sectionNameHindi = "ज्वालानगर अी0 डी0 ऒ0 0 ले0का0-5"
	case "JVALANGAR C 0 OII0 D0 V LEBAR KALONI-3":
		sectionNameHindi = "ज्वालानगर सी 0 ऒ0 डी0 व लेबर कालोनी-3"
	case "JVALANGAR C 0 OII0 D0 V LEBAR KALONI-4":
		sectionNameHindi = "ज्वालानगर सी 0 ऒ0 डी0 व लेबर कालोनी-4"
	case "JVALANGAR C0 OII0 D 0 LEVAR KALONI-1":
		sectionNameHindi = "ज्वालानगर सी0 ऒ0 डी 0 लेवर कालोनी-1"
	case "JVALANGAR C0 OII0 D0 LEVAR KALONI-6":
		sectionNameHindi = "ज्वालानगर सी0 ऒ0 डी0 लेवर कालोनी-6"
	case "JVALANGAR C0 OII0 D0 LEVAR KALONI-7":
		sectionNameHindi = "ज्वालानगर सी0 ऒ0 डी0 लेवर कालोनी-7"
	case "JVALANGAR C0 OII0 D0 LEVAR KALONI-9":
		sectionNameHindi = "ज्वालानगर सी0 ऒ0 डी व लेवर कालोनी-9"
	case "JVALANGAR C0 OII0 D0 V LEBAR KALONI-2":
		sectionNameHindi = "ज्वालानगर सी0 ऒ0 डी0 व लेबर कालोनी-2"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-10":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेवर कालोनी राजमाता फार्म-10"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-11":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-11"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-12":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-12"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-13":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-13"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-14":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-14"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-15":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-15"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-16":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-16"
	case "JVALANGAR C0OII0D0 V LEBAR KA0RAJAMATA FARM-17":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेबर का0राजमाता फार्म-17"
	case "JVALANGAR C0OII0D0 V LEVAR KALONI -8":
		sectionNameHindi = "ज्वालानगर सी0ऒ0डी0 व लेवर कालोनी -8"
	case "JWALACHANDI":
		sectionNameHindi = "ज्‍वालाचंडी"
	case "JWALAPURI":
		sectionNameHindi = "ज्वालापुरी"
	case "JWALI LALA AN0":
		sectionNameHindi = "ज्‍वाली लाला आं0"
	case "JYARAT HALKE VALI":
		sectionNameHindi = "ज्यारत हल्के वाली"
	case "JYOHRA-1":
		sectionNameHindi = "ज्‍योहरा-1"
	case "JYOHRA-2":
		sectionNameHindi = "ज्‍योहरा-2"
	case "JYOTHI":
		sectionNameHindi = "ज्‍योठी"
	case "KABIR NAGAR KALONI":
		sectionNameHindi = "कबीर नगर कालोनी"
	case "KABIRNAGAR A.":
		sectionNameHindi = "कबीरनगर आं0"
	case "KABULPUR":
		sectionNameHindi = "कबूलपुर"
	case "KACHAHARI SARAI AN0":
		sectionNameHindi = "कचहरी सराय आं0"
	case "KACHANAL":
		sectionNameHindi = "कचनाल"
	case "KACHCHI BASTI SURAJ NAGAR":
		sectionNameHindi = "कच्‍ची बस्‍ती सूरज नगर"
	case "KACHHERI SARAI AN0":
		sectionNameHindi = "कचहरी सराय आं0"
	case "KACHHPURA":
		sectionNameHindi = "कच्‍छपुरा"
	case "KACHNAL":
		sectionNameHindi = "कचनाल"
	case "KACHODAN DHAKKA":
		sectionNameHindi = "कचोदन ढक्‍का"
	case "KADAPUR":
		sectionNameHindi = "कडापुर"
	case "KADARABAD KHURD":
		sectionNameHindi = "कादराबाद खुर्द"
	case "KADARGANJ":
		sectionNameHindi = "कादरगंज"
	case "KADARPUR TAYYAB":
		sectionNameHindi = "कादरपुर तैययब"
	case "KADARPUR":
		sectionNameHindi = "कादरपुर"
	case "KADARPURNANU":
		sectionNameHindi = "कादरपुरनानू"
	case "KADRABAD":
		sectionNameHindi = "कादराबाद"
	case "KADRAPUR MASTI":
		sectionNameHindi = "कादरपुर मस्ती"
	case "KADRIGANJ":
		sectionNameHindi = "कादरीगंज"
	case "KAFIABAD":
		sectionNameHindi = "काफियाबाद"
	case "KAFOORPUR":
		sectionNameHindi = "काफूरपुर"
	case "KAGANAGLA":
		sectionNameHindi = "कागानगला"
	case "KAHARAN AN.":
		sectionNameHindi = "कहारान आ0"
	case "KAHARANA AN.":
		sectionNameHindi = "कहारान आ0"
	case "KAHRIPUR JANGAL":
		sectionNameHindi = "केहरीपुर जंगल"
	case "KAIRAN":
		sectionNameHindi = "कायरान"
	case "KAISONAGALI TANDA":
		sectionNameHindi = "कैशोनगली टाण्‍डा"
	case "KAITHOLA":
		sectionNameHindi = "कैथोला"
	case "KAJAMPUR":
		sectionNameHindi = "काजमपुर"
	case "KAJARHAI":
		sectionNameHindi = "कजरहाई"
	case "KAJI SARAY":
		sectionNameHindi = "काजी सराय"
	case "KAJIKHEDA":
		sectionNameHindi = "काजीखेडा"
	case "KAJIPUR IMMA":
		sectionNameHindi = "काजीपुर इम्‍मा"
	case "KAJIPUR PACHDU":
		sectionNameHindi = "काजीपुर पचदू"
	case "KAJIPURA KHALSA":
		sectionNameHindi = "काजीपुरा खालसा"
	case "KAJIPURA":
		sectionNameHindi = "काजीपुरा"
	case "KAJISARAI AN.":
		sectionNameHindi = "काजीसराय आं०"
	case "KAJISARAI AN.H.NO 1-270":
		sectionNameHindi = "काजीसराय आं० म०सं० १ से २७०"
	case "KAJISARAI AN.H.NO 1-73":
		sectionNameHindi = "काजीसराय आं०म०सं०१से ७३तक"
	case "KAJISARAI AN.H.NO 271-END":
		sectionNameHindi = "काजीसराय आं०म०सं०२७१ से अन्‍त तक"
	case "KAJISARAI AN.H.NO 74-END":
		sectionNameHindi = "काजीसराय आं०म०सं०७४ से अंत तक"
	case "KAJISARAY A.":
		sectionNameHindi = "काजीसराय आ०"
	case "KAJIYAN AN.H.NO 1-127":
		sectionNameHindi = "काजीयान आं०म०सं०१से १२७"
	case "KAJIYAN AN.H.NO 128-END":
		sectionNameHindi = "काजीयान आं०म०सं०128से अंत तक"
	case "KAJIYAN":
		sectionNameHindi = "काजीयान"
	case "KAJIYAPURA":
		sectionNameHindi = "कजियापुरा"
	case "KAKADPUR":
		sectionNameHindi = "कक्‍डपुर"
	case "KAKAR GHATE":
		sectionNameHindi = "ककर घटा"
	case "KAKKROA-3":
		sectionNameHindi = "ककरौआ-3"
	case "KAKRALA":
		sectionNameHindi = "ककराला"
	case "KAKROA-1":
		sectionNameHindi = "ककरौआ-1"
	case "KAKROA-2":
		sectionNameHindi = "ककरौआ-2"
	case "KAKROA-4":
		sectionNameHindi = "ककरौआ-4"
	case "KAKROA-5":
		sectionNameHindi = "ककरौआ-5"
	case "KALA JHANDA":
		sectionNameHindi = "काला झांडा"
	case "KALA NANGLI GAIR ABAD":
		sectionNameHindi = "कला नंगली गैर आबाद"
	case "KALA PYADA":
		sectionNameHindi = "काला प्‍यादा"
	case "KALAKHEDI":
		sectionNameHindi = "कालाखेडी"
	case "KALAL KHEDA URF SHAHAJADAKHEDA":
		sectionNameHindi = "कलाल खेड़ा उर्फ शहजादखेडा"
	case "KALALAN AN.":
		sectionNameHindi = "कलालान आं०"
	case "KALALAN AN.H.NO 1-171":
		sectionNameHindi = "कलालान आं०म०सं०१ से १७१तक"
	case "KALALAN AN.H.NO 1-350":
		sectionNameHindi = "कलालान आं०म०सं०१से ३५०तक"
	case "KALALAN H.NO 172-END":
		sectionNameHindi = "कलालान आं०म०सं०१७२ से अन्‍त तक"
	case "KALALAN":
		sectionNameHindi = "कलालान"
	case "KALALI":
		sectionNameHindi = "कलाली"
	case "KALAN AN. H.NO 351-END":
		sectionNameHindi = "कलालान आं०म०सं०३५१ से अन्‍त तक"
	case "KALANPUR":
		sectionNameHindi = "केलनपुर"
	case "KALAPUR BUJURG":
		sectionNameHindi = "कलापुर बुजुर्ग"
	case "KALAVALA":
		sectionNameHindi = "कालावाला"
	case "KALAYANPUR":
		sectionNameHindi = "कल्‍यानपुर"
	case "KALGHAR-1":
		sectionNameHindi = "कलघर-1"
	case "KALGHAR-2":
		sectionNameHindi = "कलघर-2"
	case "KALHERI A":
		sectionNameHindi = "कल्‍हेड़ी ए"
	case "KALHERI":
		sectionNameHindi = "कल्‍हेड़ी"
	case "KALIYA NAGALA":
		sectionNameHindi = "कलईया नगला"
	case "KALKATTA":
		sectionNameHindi = "कलकत्ता"
	case "KALLU NAWADA":
		sectionNameHindi = "कल्‍लूनवादा"
	case "KALLUGANJ":
		sectionNameHindi = "कल्‍लूगंज"
	case "KALLUWALA AN0":
		sectionNameHindi = "कल्‍लूवाला आ0"
	case "KALLUWALA":
		sectionNameHindi = "कल्‍लूवाला"
	case "KALPUR":
		sectionNameHindi = "केलपुर"
	case "KALRAKH":
		sectionNameHindi = "कलरख"
	case "KALYANAPUR":
		sectionNameHindi = "कल्यानपुर"
	case "KALYANPUR H.NO 1 TO END":
		sectionNameHindi = "कल्‍याणपुर म0 स0 १ से अन्‍त तक"
	case "KALYANPUR-1":
		sectionNameHindi = "कल्‍यानपुर-1"
	case "KALYANPUR-2":
		sectionNameHindi = "कल्‍यानपुर-2"
	case "KALYANPUR":
		sectionNameHindi = "कल्‍यानपुर"
	case "KAMALA":
		sectionNameHindi = "कमाला"
	case "KAMALAPUR FATEHABAD":
		sectionNameHindi = "कमालपुर फतेहाबाद"
	case "KAMALAPUR":
		sectionNameHindi = "कमालपुर"
	case "KAMALAPURI KHALSA":
		sectionNameHindi = "कमालपुरी खालसा"
	case "KAMALPUR A.":
		sectionNameHindi = "कमालपुर आ०"
	case "KAMALPUR BHOGA":
		sectionNameHindi = "कमालपुर भोगा"
	case "KAMALPUR GAIR ABAD":
		sectionNameHindi = "कमालपुर गैर आबाद"
	case "KAMALPUR KUNWARSAIN GAIRA":
		sectionNameHindi = "कमालपुर कुवॅरसैन गैरा0"
	case "KAMALPUR":
		sectionNameHindi = "कमालपुर"
	case "KAMALPURPURAINI":
		sectionNameHindi = "कमालपुरपुरैनी"
	case "KAMALUDDINPUR HUSANPUR":
		sectionNameHindi = "कमालुद्वीनपुर हुसैनपुर"
	case "KAMBHAUR":
		sectionNameHindi = "कम्‍भौर"
	case "KAMGARPUR AN0":
		sectionNameHindi = "कामगारपुर आं0"
	case "KAMKARAN AN.":
		sectionNameHindi = "कमकरान आं०"
	case "KAMNA UMRA G.A":
		sectionNameHindi = "कमना उमरा गै०आ०"
	case "KAMORA-1":
		sectionNameHindi = "कमोरा-1"
	case "KAMORA-2":
		sectionNameHindi = "कमोरा-2"
	case "KAMRAJPUR AN0":
		sectionNameHindi = "कामराजपुर आं0"
	case "KAMRUDDIN NAGAR":
		sectionNameHindi = "कमरूददीन नगर"
	case "KAMRUDINNAGAR":
		sectionNameHindi = "कमरूदीन नगर"
	case "KAMUA NAGALA":
		sectionNameHindi = "कमुआ नगला"
	case "KANAKPUR AN0":
		sectionNameHindi = "कनकपुर आं0"
	case "KANAKPUR-1":
		sectionNameHindi = "कनकपुर-1"
	case "KANAKPUR-2":
		sectionNameHindi = "कनकपुर-2"
	case "KANAKPUR":
		sectionNameHindi = "कनकपुर"
	case "KANDKHEDI":
		sectionNameHindi = "कन्‍द खेडी"
	case "KANHA NANGLA":
		sectionNameHindi = "कान्‍हा नंगला"
	case "KANHADAWALA":
		sectionNameHindi = "कन्‍हडवाला"
	case "KANHEDI":
		sectionNameHindi = "कन्‍हेडी"
	case "KANJARI SARAI JAIL KE PICHHE":
		sectionNameHindi = "कंजरी सराय जेल के पीछे"
	case "KANKARKHEDA":
		sectionNameHindi = "कांकरखेडा"
	case "KANKARKHERA":
		sectionNameHindi = "कांकरखेड़ा"
	case "KANONGOYAN":
		sectionNameHindi = "कानूनगोयान"
	case "KANOON GOYAN ANSHIK":
		sectionNameHindi = "कानून गोयान आंशिक"
	case "KANOONGOYAN 01":
		sectionNameHindi = "कानूनगोयान 01"
	case "KANOONGOYAN 02":
		sectionNameHindi = "कानूनगोयान 0२"
	case "KANOONGOYAN 03":
		sectionNameHindi = "कानूनगोयान 0३"
	case "KANOVI":
		sectionNameHindi = "कनोवी"
	case "KANUNGOYAN":
		sectionNameHindi = "कानूनगोयान"
	case "KANWALNAINPUR":
		sectionNameHindi = "कंवलनैनपुर"
	case "KAPNERI":
		sectionNameHindi = "कपनेरी"
	case "KARAITHI -1":
		sectionNameHindi = "करैथी -1"
	case "KARAITHI-2":
		sectionNameHindi = "करैथी -२"
	case "KARAL":
		sectionNameHindi = "कराल"
	case "KARAMPUR URF GAERI":
		sectionNameHindi = "करमपुर उर्फ गांवडी"
	case "KARANAVALA JABTI":
		sectionNameHindi = "करनावाला जब्ती"
	case "KARANAVALA KHALSA":
		sectionNameHindi = "करनावाला खालसा"
	case "KARANPUR HARKISHANPUR":
		sectionNameHindi = "करनपुर हरकिशनपुर"
	case "KARANPUR":
		sectionNameHindi = "करनपुर"
	case "KARAUHIDI":
		sectionNameHindi = "करौदी"
	case "KARAVAR":
		sectionNameHindi = "करावर"
	case "KARBALA REHMAT NAGAR":
		sectionNameHindi = "कर्बला रहमत नगर"
	case "KAREEMGUNJ":
		sectionNameHindi = "करीमगंज"
	case "KAREEMPUR MUBARAK":
		sectionNameHindi = "करीमपुर मुबारक"
	case "KAREEMPUR SHARKI":
		sectionNameHindi = "करीमपुर शर्की"
	case "KAREEMPUR":
		sectionNameHindi = "करीमपुर"
	case "KARIMAPUR JABTI":
		sectionNameHindi = "करीमपुर जब्ती"
	case "KARIMGANJ":
		sectionNameHindi = "करीमगंज"
	case "KARIMNAGAR URF ULEDA":
		sectionNameHindi = "करीमनगर उर्फ उलेढा"
	case "KARIMPUR BAMNAULI":
		sectionNameHindi = "करीमपुर बमनौली"
	case "KARIMPUR GARVI":
		sectionNameHindi = "करीमपुर गर्वी"
	case "KARINGA":
		sectionNameHindi = "करींगा"
	case "KARIYANAGLA MUST.":
		sectionNameHindi = "करियानगला मु0"
	case "KARIYANAGLA SANI":
		sectionNameHindi = "करियानगला सानी"
	case "KARKHEDA-1":
		sectionNameHindi = "करखेडा-1"
	case "KARKHEDA-2":
		sectionNameHindi = "करखेडा-2"
	case "KARKHEDI":
		sectionNameHindi = "करखेडी"
	case "KARMASKHERI":
		sectionNameHindi = "करमसखेडी"
	case "KARNAPUR":
		sectionNameHindi = "करनपुर"
	case "KAROLI":
		sectionNameHindi = "करौली"
	case "KARONDA CHODHAR AN.H.NO 131-END":
		sectionNameHindi = "करोदा चोधर आं०म०सं०१३१से अंत तक"
	case "KARONDA CHODHAR AN.H.NO.1-130":
		sectionNameHindi = "करोदा चौधरआं०म०सं १से १३०तक"
	case "KARONDA PACHDU AN.":
		sectionNameHindi = "करोदा पचदू आं०"
	case "KARONDA PACHDU AN.H.NO 1-200":
		sectionNameHindi = "करोदा पचदू आं०म०सं०१-२००तक"
	case "KARONDAPAHDU AN.H.NO 201-END":
		sectionNameHindi = "करोदा पचदू आं०म०सं०२०१ से अंत तक"
	case "KARSARA":
		sectionNameHindi = "करसरा"
	case "KARSAULA":
		sectionNameHindi = "करसौला"
	case "KARSAULI":
		sectionNameHindi = "करसौली"
	case "KASAIBADA":
		sectionNameHindi = "कसाईबाडा"
	case "KASAMGANJ":
		sectionNameHindi = "कासमगंज"
	case "KASAMPUR AHMAD":
		sectionNameHindi = "कासमपुर अहमद"
	case "KASAMPUR BILLOCH":
		sectionNameHindi = "कासमपुर बिल्‍लोच"
	case "KASAMPUR DHANNA PUR MAFI":
		sectionNameHindi = "कासमपुर धन्‍ना माफी"
	case "KASAMPUR GARHI AN0":
		sectionNameHindi = "कासमपुर गढी आ0"
	case "KASAMPUR KRIPARAM URF SUNDARPUR":
		sectionNameHindi = "कासमपुर कृपाराम उर्फ सुंदरपुर"
	case "KASAMPUR LEKHRAJ":
		sectionNameHindi = "कासमपुर लेखराज"
	case "KASAMPUR MANGAL KHEDA":
		sectionNameHindi = "कासमपुर मंगल खेडा"
	case "KASAMPUR NIKAT KANTH":
		sectionNameHindi = "कासमपुर निकट कांठ"
	case "KASAMPUR VIRU KHALSA":
		sectionNameHindi = "कासमपुर वीरु खालसा"
	case "KASBA AN.":
		sectionNameHindi = "कस्‍बा आं0"
	case "KASBA KOTRA AN. H.NO .":
		sectionNameHindi = "कस्‍बा कोटरा आं०"
	case "KASBA KOTRA AN.H.NO 1-100":
		sectionNameHindi = "कस्‍बा कोटरा आं०म०सं०१से १००"
	case "KASBA KOTRA AN.H.NO 101-212":
		sectionNameHindi = "कस्‍बा कोटरा आं०म०सं०१०१ से २12 तक"
	case "KASBA":
		sectionNameHindi = "कस्‍बा"
	case "KASDI TOLA":
		sectionNameHindi = "कासदी टोला"
	case "KASHAVNAGLA":
		sectionNameHindi = "काशवनगला"
	case "KASHIPUR -1":
		sectionNameHindi = "काशीपुर 1"
	case "KASHIPUR -2":
		sectionNameHindi = "काशीपुर 2"
	case "KASHIPUR-1":
		sectionNameHindi = "काशीपुर-1"
	case "KASHIPUR-2":
		sectionNameHindi = "काशीपुर-2"
	case "KASHIPUR-3":
		sectionNameHindi = "काशीपुर-3"
	case "KASHIPUR-4":
		sectionNameHindi = "काशीपुर-4"
	case "KASHIPUR-5":
		sectionNameHindi = "काशीपुर-5"
	case "KASHIPUR-6":
		sectionNameHindi = "काशीपुर-6"
	case "KASHIPUR-7":
		sectionNameHindi = "काशीपुर-7"
	case "KASHIPUR-8":
		sectionNameHindi = "काशीपुर-8"
	case "KASHIRAM NAGAR BLOCK -":
		sectionNameHindi = "काशीराम नगर ब्‍लॉक -"
	case "KASHIRAMPUR":
		sectionNameHindi = "काशीरामपुर"
	case "KASHIWALA":
		sectionNameHindi = "काशीवाला"
	case "KASHMIRI ABUNASARPUR":
		sectionNameHindi = "कश्‍मीरी अबूनासरपुर"
	case "KASHMIRPUR GADI":
		sectionNameHindi = "कशमीरपुर गढी"
	case "KASHRIPUR":
		sectionNameHindi = "केसरीपुर"
	case "KASIYA KUNDA-1":
		sectionNameHindi = "कसिया कुण्‍डा-1"
	case "KASIYA KUNDA-2":
		sectionNameHindi = "कसिया कुण्‍डा-2"
	case "KASMABAD":
		sectionNameHindi = "कासमाबाद"
	case "KASSABAN":
		sectionNameHindi = "कस्‍साबान"
	case "KASSHABAN":
		sectionNameHindi = "कस्‍सावान"
	case "KASWA RAJPUR":
		sectionNameHindi = "कस्‍बा राजपुर"
	case "KATA BAGH ASALATPURA":
		sectionNameHindi = "कटा बाग असालतपुरा"
	case "KATABAGH ANSHIK":
		sectionNameHindi = "कटाबाग आंशिक"
	case "KATAKUIYA MAY FILAKHANA KADIM-1":
		sectionNameHindi = "कटकुईया मय फीलखाना कदीम-1"
	case "KATAKUIYA MAY FILAKHANA KADIM-2":
		sectionNameHindi = "कटकुईया मय फीलखाना कदीम-2"
	case "KATAR SHAHEED ANSHIK":
		sectionNameHindi = "कटार शहीद आंशिक"
	case "KATAR SHAHEED":
		sectionNameHindi = "कटार शहीद"
	case "KATARMAL A.":
		sectionNameHindi = "कटारमल आ०"
	case "KATGHAR BEECH ANSHIK":
		sectionNameHindi = "कटघर बीच आंशिक"
	case "KATGHAR BEECH HOLI KA MEDAN":
		sectionNameHindi = "कटघर बीच होली का मैदान"
	case "KATRA CHETRAM":
		sectionNameHindi = "कटरा चेतराम"
	case "KATRA JALALUDIN MAY GHER HASSAN KHAN":
		sectionNameHindi = "कटरा जलालुदीन मय धेर हस्सन खां"
	case "KATRA NAJ":
		sectionNameHindi = "कटरा नाज"
	case "KATRA PURAN JAT":
		sectionNameHindi = "कटरा पूरन जाट"
	case "KAUSHALGANJ":
		sectionNameHindi = "कौशलगंज"
	case "KAUTALKHEDA":
		sectionNameHindi = "कौतलखेडा"
	case "KAYAMGANJ":
		sectionNameHindi = "कायमगंज"
	case "KAYAMNAGARDODHRAJPUR":
		sectionNameHindi = "कायमनगर दोदराजपुर"
	case "KAYASTAN A.":
		sectionNameHindi = "कायस्‍थान आ०"
	case "KAYASTH SARAI AN 151-END":
		sectionNameHindi = "कायस्‍थ सराय आ०म०सं०१५१ से अन्‍त तक"
	case "KAYASTH SARAI AN.H.NO 1-150":
		sectionNameHindi = "कायस्‍थ सराय आं०म०सं०१ से १५०तक"
	case "KAYASTHAN AN.":
		sectionNameHindi = "कायस्‍थान आ0"
	case "KAYASTHAN MADHYA":
		sectionNameHindi = "कायास्थान मध्य"
	case "KAYASTHAN PASH‍CHIMI AANSHIK":
		sectionNameHindi = "कायास्थान पश्‍चिमी आंशिक"
	case "KAYASTHAN PURVI":
		sectionNameHindi = "कायास्थान पूर्वी"
	case "KAYASTHAN-1":
		sectionNameHindi = "कायस्‍थान-1"
	case "KAYASTHAN-2":
		sectionNameHindi = "कायस्‍थान-2"
	case "KAYASTHAN":
		sectionNameHindi = "कायस्‍थान"
	case "KAYRI FUL GAIR ABAD":
		sectionNameHindi = "कायरी फुल गैर आबाद"
	case "KAYSTAN A.":
		sectionNameHindi = "कायस्‍थान आ०"
	case "KAYSTHAN":
		sectionNameHindi = "कायस्‍थान"
	case "KAZEESORA":
		sectionNameHindi = "काजीशोरा"
	case "KAZI TOLA":
		sectionNameHindi = "काजी टोला"
	case "KAZIJADGAAN A.":
		sectionNameHindi = "काजीजादगान आ०"
	case "KAZIPURA ANSHIK":
		sectionNameHindi = "काजीपुरा आंशिक"
	case "KAZIPURA":
		sectionNameHindi = "काजीपुरा"
	case "KAZIYAN":
		sectionNameHindi = "काजियान"
	case "KEORAR-1":
		sectionNameHindi = "क्‍योरार-1"
	case "KEORAR-2":
		sectionNameHindi = "क्‍योरार-2"
	case "KESARPUR NEAR KAREEMGANJ":
		sectionNameHindi = "केशरपुर निकट करीमगंज"
	case "KESARPUR":
		sectionNameHindi = "केसरपुर"
	case "KESHARPUR NEAR RAYPUR":
		sectionNameHindi = "केसरपुर नि0 रायपुर"
	case "KESHODASPUR GAIR AB.":
		sectionNameHindi = "केशोदसपुर गैर आं0"
	case "KESHONGLI SUAR":
		sectionNameHindi = "केशोनगली स्‍वार"
	case "KESHOPUR":
		sectionNameHindi = "केशोपुर"
	case "KEWALPUR":
		sectionNameHindi = "केवलपुर"
	case "KHABRI GANDU":
		sectionNameHindi = "खाबरी गन्दू"
	case "KHADA G.B":
		sectionNameHindi = "खेडा गै०आ०"
	case "KHADAK":
		sectionNameHindi = "खडक"
	case "KHADANA":
		sectionNameHindi = "खदाना"
	case "KHAI KHEDA":
		sectionNameHindi = "खाई खेडा"
	case "KHAIDIJAT":
		sectionNameHindi = "खेडी जट"
	case "KHAIKHEDHA ANSHIK":
		sectionNameHindi = "खाईखेड़ा आंशिक"
	case "KHAIRABAD":
		sectionNameHindi = "खैराबाद"
	case "KHAIRKHATA ANSHIK":
		sectionNameHindi = "खैरखाता आंशिक"
	case "KHAIRULLA PUR":
		sectionNameHindi = "खैरूल्‍लापुर"
	case "KHAIRULLAPUR AN0":
		sectionNameHindi = "खैरूल्‍लापुर आं0"
	case "KHAIRULLAPUR":
		sectionNameHindi = "खैरूल्‍लापुर"
	case "KHAIYA KHADDAR":
		sectionNameHindi = "खईया खददर"
	case "KHAJRA":
		sectionNameHindi = "खजरा"
	case "KHAJURA JAT":
		sectionNameHindi = "खजूराजट"
	case "KHAJURI":
		sectionNameHindi = "खजूरी"
	case "KHAJURIYA KALAN":
		sectionNameHindi = "खजुरिया कलां"
	case "KHAJURIYA KHURD":
		sectionNameHindi = "खजुरिया खुर्द"
	case "KHAJURIYA":
		sectionNameHindi = "खजुरिया"
	case "KHAKROBAN ST.":
		sectionNameHindi = "खाकरोबान द०"
	case "KHAKROBAN UTTARI":
		sectionNameHindi = "खाकरोबान उत्‍तरी"
	case "KHAKROBAN":
		sectionNameHindi = "खाकरोबान"
	case "KHAKROWAN":
		sectionNameHindi = "खाकरोवान"
	case "KHALILAPUR":
		sectionNameHindi = "ख‍लीलपुर"
	case "KHALILPUR KADDIM":
		sectionNameHindi = "खलीलपुर कददीम"
	case "KHALILPUR":
		sectionNameHindi = "खलीलपुर"
	case "KHALIULLAPUR":
		sectionNameHindi = "खल्‍लीउल्‍लापुर"
	case "KHAMARIYA-1":
		sectionNameHindi = "खमरिया-1"
	case "KHAMARIYA-2":
		sectionNameHindi = "खमरिया-2"
	case "KHAMRI":
		sectionNameHindi = "खमरी"
	case "KHANA PATTI":
		sectionNameHindi = "खाना पटटी"
	case "KHANAPUR LAKKHI":
		sectionNameHindi = "खानपुर लक्खी"
	case "KHANAPURSIRIYA":
		sectionNameHindi = "खानपुरसिरिया"
	case "KHANAWALUDALLAWALI":
		sectionNameHindi = "खानावाली दल्‍लावाली"
	case "KHANDELI 02":
		sectionNameHindi = "खन्‍देली 02"
	case "KHANDELI":
		sectionNameHindi = "खन्‍देली"
	case "KHANDI KHEDA":
		sectionNameHindi = "खाण्‍डीखेडा"
	case "KHANDIA-1":
		sectionNameHindi = "खडिया -1"
	case "KHANDIA-2":
		sectionNameHindi = "खडिया -2"
	case "KHANDIA-3":
		sectionNameHindi = "खडिया -3"
	case "KHANDIYA":
		sectionNameHindi = "खण्डिया"
	case "KHANDSAL":
		sectionNameHindi = "खण्‍डसाल"
	case "KHANDSHAL":
		sectionNameHindi = "खण्‍डसाल"
	case "KHANDUVA":
		sectionNameHindi = "खन्डुवा"
	case "KHANJAHANPUR BHADAR":
		sectionNameHindi = "खानजहांपुर बहादुर"
	case "KHANJIPURA":
		sectionNameHindi = "खंजीपुरा"
	case "KHANKHANAPUR ORF BICHPURI":
		sectionNameHindi = "खानखानापुर उर्फ विचपुरी"
	case "KHANMPUR CHAK GAIR AB.AD":
		sectionNameHindi = "खानमपुर चक गैर आं0"
	case "KHANPUR BILLOCH":
		sectionNameHindi = "खानपुर बिल्‍लौच"
	case "KHANPUR DULLI":
		sectionNameHindi = "खानपुर दुल्‍ली"
	case "KHANPUR JADID-1":
		sectionNameHindi = "खानपुर जदीद-1"
	case "KHANPUR JADID-2":
		sectionNameHindi = "खानपुर जदीद-2"
	case "KHANPUR KADIM":
		sectionNameHindi = "खानपुर कदीम"
	case "KHANPUR KASBA (GAIRA.)":
		sectionNameHindi = "खानपुर कस्‍बा (गैरा0)"
	case "KHANPUR KHADER A.":
		sectionNameHindi = "खानपुर खादर आ०"
	case "KHANPUR KHARVI":
		sectionNameHindi = "खानपुर गर्वी"
	case "KHANPUR MADHO":
		sectionNameHindi = "खानपुर माधो"
	case "KHANPUR MAJHRA GAIRA0":
		sectionNameHindi = "खानपुर मझरा मौ0 गैरा0"
	case "KHANPUR MAJHRA MAU0 GAIRA0":
		sectionNameHindi = "खानपुर मझरा मौ0 गैरा0"
	case "KHANPUR MAJRA GANSURPUR":
		sectionNameHindi = "खानपुर मजरा घंसूरपुर"
	case "KHANPUR MUJAFFARPUR":
		sectionNameHindi = "खानपुर मुजफफरपुर"
	case "KHANPUR RAYDASS G.A":
		sectionNameHindi = "खानपुर रायदास गै०अ०"
	case "KHANPUR UTTARI-1":
		sectionNameHindi = "खानपुर उत्‍तरी 1"
	case "KHANPUR UTTARI-2":
		sectionNameHindi = "खानपुर उत्‍तरी 2"
	case "KHANPUR":
		sectionNameHindi = "खानपुर"
	case "KHANPURMANAK":
		sectionNameHindi = "खानपुर मानक"
	case "KHANUJAT":
		sectionNameHindi = "खानूजट"
	case "KHANUPURA":
		sectionNameHindi = "खनूपुरा"
	case "KHARADIYAN":
		sectionNameHindi = "खरादियान"
	case "KHARAGPUR BAJE ANSHIK HADAYPUR GAIRABAD":
		sectionNameHindi = "खरगपुर बाजे आंशिक हदयपुर गैरा0"
	case "KHARDIYA":
		sectionNameHindi = "खरदिया"
	case "KHARGAPUR JAGTAPUR":
		sectionNameHindi = "खरगपुर जगतपुर"
	case "KHARI":
		sectionNameHindi = "खारी"
	case "KHARPUR JAGEER":
		sectionNameHindi = "खैरपुर जागीर"
	case "KHARSOUL 01":
		sectionNameHindi = "खरसौल 01"
	case "KHARSOUL 02":
		sectionNameHindi = "खरसौल 02"
	case "KHASEPUR":
		sectionNameHindi = "खासेपुर"
	case "KHASOR":
		sectionNameHindi = "खसौर"
	case "KHASPURA A,.":
		sectionNameHindi = "खासपुरा आ0"
	case "KHASPURA A.":
		sectionNameHindi = "खासपुरा आ0"
	case "KHATA CHINTAMAN":
		sectionNameHindi = "खाता चिन्‍तामन"
	case "KHATA KALAN-1":
		sectionNameHindi = "खाता कलां-1"
	case "KHATA KALAN-2":
		sectionNameHindi = "खाता कलां-2"
	case "KHATA KALAN-4":
		sectionNameHindi = "खाता कलां-4"
	case "KHATA":
		sectionNameHindi = "खाता"
	case "KHATAI":
		sectionNameHindi = "खटाई"
	case "KHATAPUR":
		sectionNameHindi = "खतापुर"
	case "KHATKAN PASIYAN V TALAV MASTU KHAN":
		sectionNameHindi = "खटकान पसियान व तालाव मस्तु खा"
	case "KHATPUR":
		sectionNameHindi = "खटपुर"
	case "KHATRIYAN AN.":
		sectionNameHindi = "खातियान आ0"
	case "KHATRIYAN B-12":
		sectionNameHindi = "खत्रियान बी 12"
	case "KHATRIYAN B-13":
		sectionNameHindi = "खत्रियान बी 13"
	case "KHAUD-1":
		sectionNameHindi = "खौद-1"
	case "KHAUD-2":
		sectionNameHindi = "खौद-2"
	case "KHAUNDALPUR-1":
		sectionNameHindi = "खौंदलपुर-1"
	case "KHAUNDALPUR-2":
		sectionNameHindi = "खौंदलपुर-2"
	case "KHAVDIYABHOOD ANSHIK":
		sectionNameHindi = "खवडियाभूड आंशिक"
	case "KHAVRI AVVAL":
		sectionNameHindi = "खावरी अव्वल"
	case "KHAWADIAGHAT":
		sectionNameHindi = "खवडियाघाट"
	case "KHAWAJA AHMADPUR JALAL URF PAITIYA":
		sectionNameHindi = "ख्‍वाजा अहमपुर जलाल उर्फै पैतिया"
	case "KHAZUYA KHEDA":
		sectionNameHindi = "खजुआ खेडा"
	case "KHEDA D. 1":
		sectionNameHindi = "खेडा द0 1"
	case "KHEDA D. 2":
		sectionNameHindi = "खेडा द0 2"
	case "KHEDA GAJROLA":
		sectionNameHindi = "खेडा गजरौला"
	case "KHEDA U. 1":
		sectionNameHindi = "खेडा उ0 1"
	case "KHEDA U. 3":
		sectionNameHindi = "खेडा उ0 3"
	case "KHEDA U.2":
		sectionNameHindi = "खेडा उ0 2"
	case "KHEDA":
		sectionNameHindi = "खेडा"
	case "KHEDATANDA-1":
		sectionNameHindi = "खेडा टाण्‍डा 1"
	case "KHEDATANDA-2":
		sectionNameHindi = "खेडा टाण्‍डा 2"
	case "KHEDATANDA-3":
		sectionNameHindi = "खेडा टाण्‍डा 3"
	case "KHEDATANDA-4":
		sectionNameHindi = "खेडा टाण्‍डा 4"
	case "KHEDI":
		sectionNameHindi = "खेडी"
	case "KHEDKI HEMRAJ COLLONY":
		sectionNameHindi = "खेडकी हेमराज कालोनी"
	case "KHEDKI TAPPA NANGAL":
		sectionNameHindi = "खेडकी टप्‍पा नांग्ल"
	case "KHEJARPUR":
		sectionNameHindi = "खि‍जरपुर"
	case "KHEMPUR-1":
		sectionNameHindi = "खेमपुर 1"
	case "KHEMPUR-2":
		sectionNameHindi = "खेमपुर 2"
	case "KHEMPUR":
		sectionNameHindi = "खेमपुर"
	case "KHERA PARSAL":
		sectionNameHindi = "खेडा पारसल"
	case "KHERA":
		sectionNameHindi = "खेडा"
	case "KHERABAD":
		sectionNameHindi = "खेराबाद"
	case "KHERKI A.":
		sectionNameHindi = "खेडकी आ०"
	case "KHERPUR JALIKA":
		sectionNameHindi = "खैरपुर जालिका"
	case "KHETAPUR":
		sectionNameHindi = "खेतापुर"
	case "KHIJARPUR":
		sectionNameHindi = "खिजरपुर"
	case "KHIJARPURJAGGU":
		sectionNameHindi = "खिजरपुर जग्‍गू"
	case "KHIMOTIYA BAKHTI":
		sectionNameHindi = "खिमोतिया बख्‍ती"
	case "KHIMOTIYA KHEDA":
		sectionNameHindi = "खिमोतिया खेडा"
	case "KHIRKA":
		sectionNameHindi = "खिरका"
	case "KHITABPUR URF KHANUPURA":
		sectionNameHindi = "खिताबपुर उर्फ खनुपुरा"
	case "KHIZARPUR":
		sectionNameHindi = "खिजरपुर"
	case "KHLILPUR AMRU":
		sectionNameHindi = "खलीलपुर अमरू"
	case "KHOD KALA-1":
		sectionNameHindi = "खौदकला 1"
	case "KHOD KALA-2":
		sectionNameHindi = "खौदकला 2"
	case "KHODPURA 1":
		sectionNameHindi = "खौदपुरा 1"
	case "KHODPURA 2":
		sectionNameHindi = "खौदपुरा 2"
	case "KHOKRA":
		sectionNameHindi = "खोकरा"
	case "KHOKRAN UTTRI":
		sectionNameHindi = "खोकरान उत्‍तरी"
	case "KHOSPURA":
		sectionNameHindi = "खोसपुरा"
	case "KHSNDSAR KOHNA":
		sectionNameHindi = "खण्‍डसार कोहना"
	case "KHUBIYA NAGALA":
		sectionNameHindi = "खुबिया नगला"
	case "KHUDAGANJ":
		sectionNameHindi = "खुदागंज"
	case "KHUDAHERI":
		sectionNameHindi = "खुडाहेडी"
	case "KHUDANAGAR":
		sectionNameHindi = "खुदानगर"
	case "KHUJISTANAGAR":
		sectionNameHindi = "खुजिस्‍तानगर"
	case "KHUKUNANGLA":
		sectionNameHindi = "खाकूनंगला"
	case "KHUNTA KHERA":
		sectionNameHindi = "खूंटाखेडा"
	case "KHURADA AN.":
		sectionNameHindi = "खुराडा आं0"
	case "KHURALLAPUR DESH":
		sectionNameHindi = "खैरूल्‍लापुर देश"
	case "KHURDI":
		sectionNameHindi = "खुर्दी"
	case "KHURRAM ALI SARAIH.NO 1-115TAK":
		sectionNameHindi = "खुर्रम अली सराय आं०म०सं०१-११५तक"
	case "KHURRAMALI SARAI AN.H.NO 116-END":
		sectionNameHindi = "खुर्रमअली सराय आं०म०सं०११६से अंत तक"
	case "KHURRAMANAGAR BICHAULA":
		sectionNameHindi = "खुर्रमनगर बिचौला"
	case "KHURRAMPUR DALLU":
		sectionNameHindi = "खुर्रमपुर डल्‍लू"
	case "KHURRAMPUR KHADAK":
		sectionNameHindi = "खुर्रमपुर खडक"
	case "KHURSHEED NAGAR":
		sectionNameHindi = "खुर्शीद नगर"
	case "KHUSHALNAGAR GARHI":
		sectionNameHindi = "खुशहालन नगर गढी"
	case "KHUSHALPUR DEWAN SINGH":
		sectionNameHindi = "खुशहालपुर दीवान सिंह"
	case "KHUSHALPUR MATHERI H.NO 102-END":
		sectionNameHindi = "खुशहालपुर मठेरी म०सं०१०२ से अन्‍त तक"
	case "KHUSHHAL NAGAR":
		sectionNameHindi = "खुशहाल नगर"
	case "KHUSHHALPUR ANSHIK":
		sectionNameHindi = "खुशहालपुर आंशिक"
	case "KHUSHHALPUR MATHERI H.NO 1-101":
		sectionNameHindi = "खुशहालपुर मठेरी म०सं०१ से १०१तक"
	case "KHUSHHALPUR":
		sectionNameHindi = "खुशहालपुर"
	case "KHUSHHAR NAGAR ANSHIK":
		sectionNameHindi = "खुशहाल नगर आंशिक"
	case "KHUSRO BAG":
		sectionNameHindi = "खुसरो बाग"
	case "KHUTIA":
		sectionNameHindi = "खुटिया"
	case "KHVAJAPUR DHANTALA":
		sectionNameHindi = "ख्वाजपुर धन्तला"
	case "KHWAJGIPUR":
		sectionNameHindi = "ख्‍वाजगीपुर"
	case "KHWAZA NAGAR DHEEMRI ANSHIK":
		sectionNameHindi = "ख्‍वाजा नगर धीमरी आंशिक"
	case "KHWAZA NAGAR":
		sectionNameHindi = "ख्‍वाजा नगर"
	case "KILA FORT":
		sectionNameHindi = "किला फोर्ट"
	case "KILA":
		sectionNameHindi = "किला"
	case "KINAN URF KALYAN":
		sectionNameHindi = "किनान उर्फ कल्‍याण"
	case "KINAN URF MADI":
		sectionNameHindi = "किनान उर्फ माडी"
	case "KIRA 01":
		sectionNameHindi = "किरा 01"
	case "KIRA 02":
		sectionNameHindi = "किरा 02"
	case "KIRA 03":
		sectionNameHindi = "किरा 03"
	case "KIRARKHEDI":
		sectionNameHindi = "किरारखेडी"
	case "KIRATPUR":
		sectionNameHindi = "किरतपुर"
	case "KIRATPURI":
		sectionNameHindi = "किरतपुरी"
	case "KIRTO NANGLI":
		sectionNameHindi = "किरतो नंगली"
	case "KISANPUR ATRIYA":
		sectionNameHindi = "किशनपुर अटरिया"
	case "KISANPUR PANCHAKKI":
		sectionNameHindi = "किशन पुर पन्‍चक्‍की"
	case "KISANPURKUNDA":
		sectionNameHindi = "कीशनपरकुण्‍डा"
	case "KISHANPU BHOGAN":
		sectionNameHindi = "किशनपुर भोगन"
	case "KISHANPUR DULLANAGALA":
		sectionNameHindi = "किशनपुर दूलानगला"
	case "KISHANPUR MOLLAGARH":
		sectionNameHindi = "किशनपुर मौलागढ"
	case "KISHANPUR":
		sectionNameHindi = "किशनपुर"
	case "KISHANWAS":
		sectionNameHindi = "किशनवास"
	case "KISHNAPUR GAVDI":
		sectionNameHindi = "किशनपुर गावडी"
	case "KISHNPUR":
		sectionNameHindi = "किशनपुर"
	case "KISHORPUR":
		sectionNameHindi = "किशोरपुर"
	case "KISHROL PASHCHIMI ANSHIK":
		sectionNameHindi = "किसरौल पश्चिमी आंशिक"
	case "KISRAUL PURVI ANSHIK":
		sectionNameHindi = "किसरौल पूर्वी आंशिक"
	case "KISROL PURVI ANSHIK":
		sectionNameHindi = "किसरौल पूर्वी आंशिक"
	case "KISROL":
		sectionNameHindi = "किसरौल"
	case "KISVA NAGLA":
		sectionNameHindi = "किसवा नगला"
	case "KITAPLAI":
		sectionNameHindi = "किटप्लाई"
	case "KITHODA MUFTI":
		sectionNameHindi = "कि थोडा मुफती"
	case "KITHODA RAI":
		sectionNameHindi = "कि थोडा राय"
	case "KIWAD A.":
		sectionNameHindi = "किवाड आ0"
	case "KODUPURA":
		sectionNameHindi = "कोडूपुरा"
	case "KOHARPUR":
		sectionNameHindi = "कोहरपुर"
	case "KOHNA MUGALPURA ANSHIK":
		sectionNameHindi = "कोहना मुगलपुरा आंशिक"
	case "KOHNKOO":
		sectionNameHindi = "कोहनकू"
	case "KOKAPUR":
		sectionNameHindi = "कोकापुर"
	case "KOLASAGAR":
		sectionNameHindi = "कोलासागर"
	case "KONDRI":
		sectionNameHindi = "कोंडरी"
	case "KOOP -2":
		sectionNameHindi = "कूप -2"
	case "KOOP -3":
		sectionNameHindi = "कूप -3"
	case "KOOP-1":
		sectionNameHindi = "कूप-1"
	case "KOPA":
		sectionNameHindi = "कोपा"
	case "KORBAKU":
		sectionNameHindi = "कोरबाकू"
	case "KORIYA":
		sectionNameHindi = "कौडि़या"
	case "KORIYAN WARD -4":
		sectionNameHindi = "कोरियान वार्ड-4"
	case "KORIYAN WARD N0-4":
		sectionNameHindi = "कोरियान वार्ड नं0-4"
	case "KORIYAN WARD NO-4":
		sectionNameHindi = "कोरियान वार्ड नं0-4"
	case "KOSHALY":
		sectionNameHindi = "कौशल्‍या"
	case "KOTA ALINAGAR":
		sectionNameHindi = "कोटा अलीनगर"
	case "KOTAVALAN":
		sectionNameHindi = "कोतवालान"
	case "KOTAVALI KOHANA":
		sectionNameHindi = "कोतवाली कोहना"
	case "KOTAWALI":
		sectionNameHindi = "कोटावाली"
	case "KOTHA JAGEER":
		sectionNameHindi = "कोठा जागीर"
	case "KOTHI KHAS BAG-1":
		sectionNameHindi = "कोठी खास बाग-1"
	case "KOTHI KHAS BAG-2":
		sectionNameHindi = "कोठी खास बाग-2"
	case "KOTHIWAL NAGAR ANSHIK":
		sectionNameHindi = "कोठीवाल नगर आंशिक"
	case "KOTKADAR HAJI MOHDPUR H.NO 327 TO END":
		sectionNameHindi = "कोटकादर हाजी मौ0पुर म0स0 327 से अंत तक"
	case "KOTLA A.":
		sectionNameHindi = "कोटला आ०"
	case "KOTLA NAGLA":
		sectionNameHindi = "कोटला नगला"
	case "KOTLA":
		sectionNameHindi = "कोटला"
	case "KOTRA AN.":
		sectionNameHindi = "कोटरा आ0"
	case "KOTRA MOLVIYAN":
		sectionNameHindi = "कौटरा मौलवियान"
	case "KOTRA TAPPA KESHO":
		sectionNameHindi = "कोटरा टप्‍पा केशो"
	case "KOTRA":
		sectionNameHindi = "कोटरा"
	case "KOTWALI AN H.NO 1-412":
		sectionNameHindi = "कोतवाली आं०म०सं०१ से ४१२तक"
	case "KOTWALI AN H.NO 1-476":
		sectionNameHindi = "कोतवाली आं०म०सं०१ से ४७६"
	case "KOTWALI AN H.NO 477-END":
		sectionNameHindi = "कोतवाली आं०म०सं ४७७ से अन्‍त तक"
	case "KOTWALI AN HNO -1-123":
		sectionNameHindi = "कोतवाली आं०म०सं०१ से १२३तक"
	case "KOTWALI AN":
		sectionNameHindi = "कोतवाली आं०"
	case "KOTWALIAN H.NO 413-END":
		sectionNameHindi = "कोतवाली आं०म०सं०४१३ से अन्‍त तक"
	case "KOYLA-1":
		sectionNameHindi = "कोयला-1"
	case "KOYLA-2":
		sectionNameHindi = "कोयला-2"
	case "KOYLA-3":
		sectionNameHindi = "कोयला-3"
	case "KOYLA-4":
		sectionNameHindi = "कोयला-4"
	case "KOYLI":
		sectionNameHindi = "कोयली"
	case "KRIMCHA-1":
		sectionNameHindi = "क्रमचा-1"
	case "KRIMCHA-2":
		sectionNameHindi = "क्रमचा-2"
	case "KRIMCHA-3":
		sectionNameHindi = "क्रमचा-3"
	case "KRIPYA HAPPU":
		sectionNameHindi = "कृपिया हप्‍पू"
	case "KRIPYA PANDE":
		sectionNameHindi = "कृपिया पाण्‍डे"
	case "KRISHNA NAGAR H.N0 493TO 512":
		sectionNameHindi = "क़ष्‍ण्‍ानगर म0स0 493 से 512तक"
	case "KSSABAN":
		sectionNameHindi = "कस्‍साबान"
	case "KUAKHEDA MAFI":
		sectionNameHindi = "कुआ खेडा माफी"
	case "KUAKHERA H NO 1 TO 176":
		sectionNameHindi = "कुआखेडा आ0 म0स0 १ से १७६ तक"
	case "KUAKHERA H NO177 TO END":
		sectionNameHindi = "कुआखेडा आ0 म0स0 177 से अंत तक"
	case "KUAKHERA":
		sectionNameHindi = "कुंआखेडा"
	case "KUAN KHEDA":
		sectionNameHindi = "कुआं खेड़ा"
	case "KUCHA DEVIDAS":
		sectionNameHindi = "कूचा देवीदास"
	case "KUCHA PARAMESHVARI DAS":
		sectionNameHindi = "कूचा परमेश्‍वरी दास"
	case "KUCHA SHEKH BACHAI":
		sectionNameHindi = "कूचा शेख बचई"
	case "KUCHITA":
		sectionNameHindi = "कुचैटा"
	case "KUDAMEERPUR":
		sectionNameHindi = "कुडामीरपुर"
	case "KUDDUSPUR":
		sectionNameHindi = "कुददुसपुर"
	case "KUDI MANAK":
		sectionNameHindi = "कूड़ी मानक"
	case "KUHUTUBPUR GAWRI":
		sectionNameHindi = "कुतुबपुर गावडी"
	case "KUIYA":
		sectionNameHindi = "कुईया"
	case "KUJAITA":
		sectionNameHindi = "कुजैटा"
	case "KUKDA ISLAMPUR":
		sectionNameHindi = "कूकडा इस्‍लामपुर"
	case "KUKDI KHEDA":
		sectionNameHindi = "कुकडी खेडा"
	case "KUKURJHUNDI":
		sectionNameHindi = "कुकरझुंडी"
	case "KULAVADA":
		sectionNameHindi = "कुलवाड़ा"
	case "KULCHANA":
		sectionNameHindi = "कुलचाना"
	case "KUMAR KUNJ KOTHI SAHU RAMESH KUMAR":
		sectionNameHindi = "कुमार कुंज कोठी साहू रमेश कुमार"
	case "KUMHARIYA JUBLA":
		sectionNameHindi = "कुम्‍हरिया जुबला"
	case "KUMHARIYA KHURD":
		sectionNameHindi = "कुम्‍हरिया खुर्द"
	case "KUMHARIYA KLA":
		sectionNameHindi = "कुम्‍हरिया कला"
	case "KUMHARIYA-1":
		sectionNameHindi = "कुम्‍हरिया-1"
	case "KUMHARIYA-2":
		sectionNameHindi = "कुम्‍हरिया-2"
	case "KUMHERA":
		sectionNameHindi = "कुम्‍हेडा"
	case "KUMRARA":
		sectionNameHindi = "कुमरारा"
	case "KUNA KHEDA KHALSA":
		sectionNameHindi = "कुंआ खेड़ा खालसा"
	case "KUNCHA ATISH BAJAN":
		sectionNameHindi = "कूंचा आतिश बाजान"
	case "KUNCHA BHAGMAL":
		sectionNameHindi = "कूँचा भागमल"
	case "KUNCHA FIRANGAN":
		sectionNameHindi = "कूंचा फिरंगन"
	case "KUNCHA KAJI":
		sectionNameHindi = "कूँचा काजी"
	case "KUNCHA KASSAVAKHANA":
		sectionNameHindi = "कूंचा कस्सावखाना"
	case "KUNCHA LALA MIYAN":
		sectionNameHindi = "कूंचा लाला मियां"
	case "KUNCHA LANGARAKHANA":
		sectionNameHindi = "कूंचा लंगरखाना"
	case "KUNCHA NAKKALAN":
		sectionNameHindi = "कूँचा नक्कालान"
	case "KUNCHA NALVANDAN":
		sectionNameHindi = "कूंचा नालवन्दान"
	case "KUNCHA SEEKALGARANN":
		sectionNameHindi = "कूँचा सेकलगरांन"
	case "KUNDA 1":
		sectionNameHindi = "कुण्‍डा 1"
	case "KUNDA 2":
		sectionNameHindi = "कुण्‍डा 2"
	case "KUNDA 3":
		sectionNameHindi = "कुण्‍डा 3"
	case "KUNDA BAGAIN":
		sectionNameHindi = "कुण्‍डा बागैन"
	case "KUNDA KHURD":
		sectionNameHindi = "कुण्‍डा खुर्द"
	case "KUNDA":
		sectionNameHindi = "कुण्‍डा"
	case "KUNDANPUR ANSHIK":
		sectionNameHindi = "कुन्‍दनपुर आंशिक"
	case "KUNDANPUR":
		sectionNameHindi = "कुन्‍दनपुर"
	case "KUNDESRA":
		sectionNameHindi = "कुण्‍डेसरा"
	case "KUNDESRI":
		sectionNameHindi = "कुन्‍डेसरी"
	case "KUNVARAPUR NANAKAR-1":
		sectionNameHindi = "कुंवरपुर नानकार 1"
	case "KUNVARAPUR NANAKAR-2":
		sectionNameHindi = "कुंवरपुर नानकार 2"
	case "KUNVARPUR PADAMAPUR":
		sectionNameHindi = "कुंवरपुर पदमपुर"
	case "KURI BANGER":
		sectionNameHindi = "कुरी बंगर"
	case "KURI RAWANA AANSHIK":
		sectionNameHindi = "कूरी रवाना (आंशिक)"
	case "KURI RAWANA":
		sectionNameHindi = "कूरी रवाना (आंशिक)"
	case "KURI":
		sectionNameHindi = "कूरी"
	case "KURIRAWANA AANSHIK":
		sectionNameHindi = "कूरी रवाना (आंशिक)"
	case "KURTHIYA":
		sectionNameHindi = "कुर्थिया"
	case "KUSHHALPUR MARKA":
		sectionNameHindi = "खुशहालपुर मड़का"
	case "KUTQUIE":
		sectionNameHindi = "कटकुइ"
	case "KUTUBAPUR AJJU":
		sectionNameHindi = "कुतुबपुर अज्जू"
	case "KUTUBPUR GARHI":
		sectionNameHindi = "कुतुबपुर गढी"
	case "KUTUBPUR NANGLI":
		sectionNameHindi = "कुतुबपुर नंगली"
	case "KUTUBPUR RAIB GAIRA0":
		sectionNameHindi = "कुतुबपुर रायब गैरा0"
	case "KUTUBPUR":
		sectionNameHindi = "कुतुबपुर"
	case "KUWAR BAL GOVIND":
		sectionNameHindi = "कुवर बाल गोविन्‍द"
	case "LADANPUR":
		sectionNameHindi = "लाडनपुर"
	case "LADAPURA":
		sectionNameHindi = "लडापुरा"
	case "LADAWI":
		sectionNameHindi = "लदावली"
	case "LADAWLI":
		sectionNameHindi = "लदावली"
	case "LADHUPURA":
		sectionNameHindi = "लधुपुरा"
	case "LADORA NARAYANPUR":
		sectionNameHindi = "लदौरा नरायनपुर"
	case "LADORI":
		sectionNameHindi = "लदौरी"
	case "LADPUR BIB":
		sectionNameHindi = "लाडपुर बीबी"
	case "LADPUR GAIRA":
		sectionNameHindi = "लाडपुर गैरा"
	case "LADPUR NEAR BATHUAKHERA":
		sectionNameHindi = "लाडपुर नि0 बथुआखेडा"
	case "LADPUR NEAR SEMRA-1":
		sectionNameHindi = "लाडपुर नि0 सेमरा 1"
	case "LADPUR NEAR SEMRA-2":
		sectionNameHindi = "लाडपुर नि0 सेमरा 2"
	case "LADPUR":
		sectionNameHindi = "लाड्पुर"
	case "LADPURA AN.":
		sectionNameHindi = "लाडपुरा आं०"
	case "LADUPURA":
		sectionNameHindi = "लधुपुरा"
	case "LAHAK KALA":
		sectionNameHindi = "लाहक कला"
	case "LAHAK KHURD":
		sectionNameHindi = "लाहक खुर्द"
	case "LAIDERPUR":
		sectionNameHindi = "लैदरपुर"
	case "LAITHER":
		sectionNameHindi = "लाइठेर"
	case "LAJPAT NAGAR ANSHIK":
		sectionNameHindi = "लाजपत नगर आंशिक"
	case "LAJPAT NAGAR":
		sectionNameHindi = "लाजपत नगर"
	case "LAKADI ANSHIK":
		sectionNameHindi = "लाकडी आंशिक"
	case "LAKDI WALAN PASHCHIMI":
		sectionNameHindi = "लाकडी वालान पश्चिमी"
	case "LAKDIWALAN PURVI":
		sectionNameHindi = "लाकडीवालान पूर्वी"
	case "LAKHANPUR":
		sectionNameHindi = "लखनपुर"
	case "LAKHARAN":
		sectionNameHindi = "लकडहारान"
	case "LAKHIMPUR BHIKA":
		sectionNameHindi = "लखीमपुर भीका"
	case "LAKHIMPUR VISHNU":
		sectionNameHindi = "लखीमपुर विश्‍नू"
	case "LAKHIPURA":
		sectionNameHindi = "लखीपुरा"
	case "LAKHIWALA":
		sectionNameHindi = "लक्‍खीवाला"
	case "LAKHMAN NAGALA":
		sectionNameHindi = "लखमन लगला"
	case "LAKHMIPUR":
		sectionNameHindi = "लखमीपुर"
	case "LAKHNAKHEDA":
		sectionNameHindi = "लखनाखेडा"
	case "LAKHNAPUR LADAPUR":
		sectionNameHindi = "लखनपुर लाड़पुर"
	case "LAKMIDHAR":
		sectionNameHindi = "लख्‍मीधर"
	case "LAKSHMANPURI LINEPAR":
		sectionNameHindi = "लक्ष्‍मणपुरी लाइनपार"
	case "LAL KABAR-1":
		sectionNameHindi = "लाल कबर-1"
	case "LAL KABAR-2":
		sectionNameHindi = "लाल कबर-2"
	case "LAL MASJID-1":
		sectionNameHindi = "लाल मस्जिद-1"
	case "LAL MASJID-2":
		sectionNameHindi = "लाल मस्जिद-2"
	case "LAL NAGRI DHEEMRI ANSHIK":
		sectionNameHindi = "लाल नगरी धीमरी आंशिक"
	case "LAL NAGRI DHEEMRI":
		sectionNameHindi = "लाल नगरी धीमरी"
	case "LAL SARAI AN.H.NO 1-102":
		sectionNameHindi = "लाल सराय आं०म०सं०१से १०२"
	case "LAL SARAI AN.H.NO 1-150 TAK":
		sectionNameHindi = "लाल सराय आं०म०सं०१ से १५० तक"
	case "LAL SARAI AN.H.NO 103-355":
		sectionNameHindi = "लाल सराय आं०म०सं०१०३ से ३५५"
	case "LAL SARAI AN.H.NO 356-END":
		sectionNameHindi = "लाल सराय आं०म०सं०३५६ सेअंत तक"
	case "LAL SARAI HNO 1 TO 100":
		sectionNameHindi = "लाल सराय म0स0 १ से १०० तक"
	case "LAL SARAI HNO101 TO END":
		sectionNameHindi = "लाल सराय म0स0१०१ से अंत तक"
	case "LALA NAGLA":
		sectionNameHindi = "लाला नगला"
	case "LALAPUR HAMIR":
		sectionNameHindi = "लालपुर हमीर"
	case "LALAPUR PIPLASANA":
		sectionNameHindi = "लालापुर पीपलसाना"
	case "LALAPUR TITRI":
		sectionNameHindi = "लालपुर तीतरी"
	case "LALATIKAR MAHESH NAGLI":
		sectionNameHindi = "लालाटीकर महेश नगली"
	case "LALAVALA":
		sectionNameHindi = "लालवाला"
	case "LALAVARA AANSHIK":
		sectionNameHindi = "ललवारा आंशिक"
	case "LALBAGH KHALLE ANSHIK":
		sectionNameHindi = "लालबाग खल्‍ले आंशिक"
	case "LALBAGH KHALLE RAMGANGA COLONY":
		sectionNameHindi = "लालबाग खल्‍ले रामगंगा कालोनी"
	case "LALBAGH UTTARI ANSHIK":
		sectionNameHindi = "लालबाग उत्‍तरी आंशिक"
	case "LALPUR BIKKU":
		sectionNameHindi = "लालपुर भिक्‍कू"
	case "LALPUR CHAUHAN":
		sectionNameHindi = "लालपुर चौहान"
	case "LALPUR GANGVARI AANSHIK":
		sectionNameHindi = "लालपुर गंगवारी आंशिक"
	case "LALPUR HAMIR AANSHIK":
		sectionNameHindi = "लालपुर हमीर आंशिक"
	case "LALPUR KALA -1":
		sectionNameHindi = "लालपुर-1"
	case "LALPUR KALA -2":
		sectionNameHindi = "लालपुर-2"
	case "LALPUR KALA -3":
		sectionNameHindi = "लालपुर-3"
	case "LALPUR KALA-4":
		sectionNameHindi = "लालपुर-4"
	case "LALPUR MAN":
		sectionNameHindi = "लालपुर मान"
	case "LALPUR PATTI KHURD":
		sectionNameHindi = "लालपुर पटटी खुर्द"
	case "LALPUR PATTI KUNDAN":
		sectionNameHindi = "लालपुर पटटी कुन्‍दन"
	case "LALPUR PATTI QADAR KHAN":
		sectionNameHindi = "लालपुर पटटी कादर खां"
	case "LALPUR PATTI SAEED GANJ":
		sectionNameHindi = "लालपुर पटटी सईद गंज"
	case "LALPUR PUROHIT":
		sectionNameHindi = "लालपुर पुरोहित"
	case "LALPUR RAI":
		sectionNameHindi = "लालपुर राय"
	case "LALPUR SHOJIMAL A.":
		sectionNameHindi = "लालपुर शौजीमल ए0"
	case "LALPUR SHOJIMAL":
		sectionNameHindi = "लालपुर शौजीमल"
	case "LALPUR":
		sectionNameHindi = "लालपुर"
	case "LALPURI":
		sectionNameHindi = "लालपुरी"
	case "LALSARAI AN.H.NO 151-END":
		sectionNameHindi = "लाल सराय आं०म०सं१५१ से अन्‍त तक"
	case "LALUNAGLA-1":
		sectionNameHindi = "लालूनगला -1"
	case "LALUNAGLA2":
		sectionNameHindi = "लालूनगला2"
	case "LALUWALA ANSHIK":
		sectionNameHindi = "लालूवाला आंशिक"
	case "LALWARA -1":
		sectionNameHindi = "ललवारा -1"
	case "LALWARA 02":
		sectionNameHindi = "ललवारा 02"
	case "LAMBAKHEDA-1":
		sectionNameHindi = "लाम्‍बाखेडा-1"
	case "LAMBAKHEDA-2":
		sectionNameHindi = "लाम्‍बाखेडा-2"
	case "LAMBAKHEDA-3":
		sectionNameHindi = "लाम्‍बाखेडा-3"
	case "LAMBAKHEDA-4":
		sectionNameHindi = "लाम्‍बाखेडा-4"
	case "LAMBAKHERA A.":
		sectionNameHindi = "लाम्‍बाखेडा आं"
	case "LAMBIYA":
		sectionNameHindi = "लम्बिया"
	case "LANGDE KI PULIA ASALATPURA":
		sectionNameHindi = "लंगडे की पुलिया असालतपुरा"
	case "LANGPURI":
		sectionNameHindi = "लैंगपुरी"
	case "LASHKAR GANJ":
		sectionNameHindi = "लश्‍करगंज"
	case "LATEEFPUR CHUKHRI":
		sectionNameHindi = "लतीफपुर चुखैडी"
	case "LATEEFPUR":
		sectionNameHindi = "लतीफपुर"
	case "LATIFAPUR URF SIPAHIVALA AN.":
		sectionNameHindi = "लतीफपुर उर्फ सिपाहीवाला आं0"
	case "LATIFULLAPUR":
		sectionNameHindi = "लतीफुल्‍लापुर"
	case "LATPAT NAGAR ANSHIK":
		sectionNameHindi = "लाजपत नगर आंशिक"
	case "LATRA":
		sectionNameHindi = "लाटरा"
	case "LAXMIPUR KATTAI":
		sectionNameHindi = "लक्ष्‍मीपुर कटटई"
	case "LEHTORA":
		sectionNameHindi = "लहटौरा"
	case "LINDERPUR A.":
		sectionNameHindi = "लिन्‍डरपुर आ०"
	case "LOCOSHED ANSHIK":
		sectionNameHindi = "लोकोशेड आंशिक"
	case "LODHIPUR GULADIYA":
		sectionNameHindi = "लोधीपुर गुलडिया"
	case "LODHIPUR NAYAK-1":
		sectionNameHindi = "लोधीपुर नायक-1"
	case "LODHIPUR NAYAK-2":
		sectionNameHindi = "लोधीपुर नायक-2"
	case "LODHIPUR RAJPOOT":
		sectionNameHindi = "लोधीपुर राजपूत"
	case "LODHIPUR VASOO AI0 GAIRA0":
		sectionNameHindi = "लोधीपुर वासू ऐ0 गैरा0"
	case "LODHIPUR":
		sectionNameHindi = "लोधीपुर"
	case "LODIPUR JAWAHAR NAGAR ANSHIK":
		sectionNameHindi = "लोदीपुर जवाहर नगर आंशिक"
	case "LODIPUR MILAK":
		sectionNameHindi = "लोदीपुर मिलक"
	case "LOGI KALA":
		sectionNameHindi = "लोगी कला"
	case "LOHAPATTI BHAGIRATH":
		sectionNameHindi = "लोहापटटी भागीरथ"
	case "LOHAPATTI BHOLANATH-1":
		sectionNameHindi = "लोहापटटी भोलानाथ-1"
	case "LOHAPATTI BHOLANATH-2":
		sectionNameHindi = "लोहापटटी भोलानाथ-2"
	case "LOHARA TANDA":
		sectionNameHindi = "लोर्हरा टाण्‍डा"
	case "LOHARAN":
		sectionNameHindi = "लोहारान"
	case "LOHARI SARAI AN.H.NO 1-180":
		sectionNameHindi = "लुहारी सराय आ०म०सं०१ से 180"
	case "LOHARI SARAI PU.AN.H.NO 11/2-180/4":
		sectionNameHindi = "लोहारी सराय पूर्व आं०म०सं०११/२ से १८०/४तक"
	case "LOHARI":
		sectionNameHindi = "लौहारी"
	case "LOHARRA INAYATGANJ":
		sectionNameHindi = "लोहर्रा इनायतगंज"
	case "LOHIYAN AN.":
		sectionNameHindi = "लोहियान आ0"
	case "LONGI KHURD":
		sectionNameHindi = "लौंगी खुर्द"
	case "LONGIKALAN":
		sectionNameHindi = "लौंगीकलां"
	case "LUDHPURA":
		sectionNameHindi = "लुधपुरा"
	case "LUFTIPUR,":
		sectionNameHindi = "लुफतीपुर,"
	case "LUHARI SARAI AN.":
		sectionNameHindi = "लुहारी सराय आं०"
	case "LUHARI SARAI AN.H.NO 1-275":
		sectionNameHindi = "लुहारी सराय आ०म०सं०१से २७५"
	case "LUHARI SARAI AN.H.NO 276-END":
		sectionNameHindi = "लुहारी सराय आं०म०सं०२७६ से अन्‍त तक"
	case "LUHARI SARAI H.NO 181-END":
		sectionNameHindi = "लुहारी सराय आं०181 से अंत तक"
	case "LUHARI SARAI PURV AN.H.NO 1-283":
		sectionNameHindi = "लुहारी सराय पूर्व आं०म०सं०१से २८३तक"
	case "LUHARPURA":
		sectionNameHindi = "लुहारपुरा"
	case "LUKMANPURA":
		sectionNameHindi = "लुकमानपुरा"
	case "MACHHARIYA ANSHIK":
		sectionNameHindi = "मछरिया आंशिक"
	case "MACHHMAR":
		sectionNameHindi = "मच्‍छमार"
	case "MACHHUAPURA KATGHAR":
		sectionNameHindi = "मछुआपुरा कटघर"
	case "MACHKI":
		sectionNameHindi = "माचकी"
	case "MADAIRYAN KALI":
		sectionNameHindi = "मड़इर्यान कली"
	case "MADAIYA BHOGPUR MITHONI":
		sectionNameHindi = "मडैया भौगपुर मिठोनी"
	case "MADAIYAN BHAJJAN":
		sectionNameHindi = "मडैययान भज्‍जन"
	case "MADAIYAN JAULAPUR":
		sectionNameHindi = "मडैयान जौलपुर"
	case "MADAIYAN TULSI":
		sectionNameHindi = "मडैयान तुलसी"
	case "MADAIYAN VADE":
		sectionNameHindi = "मडैयान वदे"
	case "MADAMAHDUD":
		sectionNameHindi = "मदामहदूद"
	case "MADANAGLA":
		sectionNameHindi = "मदानगला"
	case "MADANAPUR":
		sectionNameHindi = "मदनापुर"
	case "MADARIPUR KAKRALA":
		sectionNameHindi = "मदारीपुर ककराला"
	case "MADARPUR":
		sectionNameHindi = "मदारपुर"
	case "MADARSA ZYARAT SHAH BULAKI ANSHIK":
		sectionNameHindi = "मदरसा ज्‍यारत शाह बुलाकी आंशिक"
	case "MADEYAM DHOLSAR":
		sectionNameHindi = "मडैयान ढोलसर"
	case "MADEYAN ODEYRAJ":
		sectionNameHindi = "मडैयान उदयराज"
	case "MADEYAN SENDOALI":
		sectionNameHindi = "मडैयान सैंडोली"
	case "MADH PURI":
		sectionNameHindi = "मध पुरी"
	case "MADHOWALA":
		sectionNameHindi = "माधोवाला"
	case "MADHUBA KHALSA":
		sectionNameHindi = "मधुवा खालसा"
	case "MADHUBANI ANSHIK":
		sectionNameHindi = "मधुबनी आंशिक"
	case "MADHUKAR 01":
		sectionNameHindi = "मधूकर 01"
	case "MADHUKAR 02":
		sectionNameHindi = "मधुकर 02"
	case "MADHUKAR-3":
		sectionNameHindi = "मधुकर-3"
	case "MADHUPARI -01":
		sectionNameHindi = "मधूपुरी -01"
	case "MADHUPARI -02":
		sectionNameHindi = "मधूपुरी -02"
	case "MADHUPURA-1":
		sectionNameHindi = "मधुपुरा 1"
	case "MADHUPURA-2":
		sectionNameHindi = "मधुपुरा 2"
	case "MADHUPURA":
		sectionNameHindi = "मधुपुरा"
	case "MADHUPURI":
		sectionNameHindi = "मधुपुरी"
	case "MADHUSUDANPUR URF JALALPUR":
		sectionNameHindi = "मधुसुदनपुर उर्फ जलालपुर"
	case "MADHVA":
		sectionNameHindi = "मढवा"
	case "MADIYAN JHAU":
		sectionNameHindi = "मडैयान झाऊ"
	case "MADPURI":
		sectionNameHindi = "मदपुरी"
	case "MADRASA ALIYA":
		sectionNameHindi = "मदरसा आलिया"
	case "MADRASA KOHANA":
		sectionNameHindi = "मदरसा कोहना"
	case "MADSUDANPUR BHADHAR":
		sectionNameHindi = "मदसूदनपुर बहादर"
	case "MADSUDANPUR NAND URF JHALRA":
		sectionNameHindi = "मदसूदनपुर नंद उर्फ झलरा"
	case "MAGARMAU":
		sectionNameHindi = "मगरमऊ"
	case "MAGHPUR UDAYCHAND":
		sectionNameHindi = "माघपुर उदयचन्‍द"
	case "MAGHPUR":
		sectionNameHindi = "मेघपुर"
	case "MAGHPURI ORF INAYATPUR":
		sectionNameHindi = "मघपुरी उर्फ इनायतपुर"
	case "MAHABATPUR":
		sectionNameHindi = "महावतरपुर"
	case "MAHAGALI":
		sectionNameHindi = "महगाली"
	case "MAHAJANAN PURVI":
		sectionNameHindi = "महाजनान पूर्वी"
	case "MAHAJANAN":
		sectionNameHindi = "महाजनान"
	case "MAHAKMA JUNUBI":
		sectionNameHindi = "महकमा जनूबी"
	case "MAHAL SARAI AN0":
		sectionNameHindi = "महल सराय आं0"
	case "MAHAL SARAI":
		sectionNameHindi = "महल सराय"
	case "MAHALAULI":
		sectionNameHindi = "महलौली"
	case "MAHALKI":
		sectionNameHindi = "महलकी"
	case "MAHAMDABAD A.":
		sectionNameHindi = "महमदाबाद आ0"
	case "MAHAMDABAD":
		sectionNameHindi = "महमदाबाद"
	case "MAHAMUD PUR":
		sectionNameHindi = "महमूद पुर"
	case "MAHAMUDAPUR MAFI":
		sectionNameHindi = "महमूदपुर माफी"
	case "MAHAMUDPUR MAFI":
		sectionNameHindi = "महमूदपुर माफी"
	case "MAHAPUR MOHD ALI":
		sectionNameHindi = "महापुर मौ०अली"
	case "MAHAPUR":
		sectionNameHindi = "माहपुर"
	case "MAHARAIPUR SEKH URF DINORI":
		sectionNameHindi = "महारायपुर शेख उर्फ दिनौडी"
	case "MAHARAJ PUR":
		sectionNameHindi = "महाराजपुर"
	case "MAHARATPURKALA AN0 H.NO 1 TO 147":
		sectionNameHindi = "महारतपुर कला आ0 म0स0 १ से १४७ तक"
	case "MAHARATPURKALA AN0 H.NO 1 TO 684":
		sectionNameHindi = "महारतपुर कला म0स0 १ से ६८४ तक"
	case "MAHARATPURKALA AN0 H.NO 148 TO END":
		sectionNameHindi = "महारतपुर कला आ0 म0स0148 से अंत तक"
	case "MAHARATPURKALA AN0 H.NO 685 TO END":
		sectionNameHindi = "महारतपुर कला म0स0 ६८५ से अंत तक"
	case "MAHARATPURKALA":
		sectionNameHindi = "महारतपुर कला"
	case "MAHAUDPUR MAFI":
		sectionNameHindi = "महमूदपुर माफी"
	case "MAHAWATPUR BILLOCH":
		sectionNameHindi = "महावतपुर बिल्‍लौच"
	case "MAHAWATPUR DALPAT":
		sectionNameHindi = "महावतपुर दलपत"
	case "MAHAWATPUR NATHA":
		sectionNameHindi = "महावतपुर नाथा"
	case "MAHBULAPURDHAKI":
		sectionNameHindi = "महबुल्‍लापुर ढाकी"
	case "MAHDUDNASHO":
		sectionNameHindi = "महदूदनशो"
	case "MAHEMSAPUR":
		sectionNameHindi = "महमसापुर"
	case "MAHENDER NAGAR":
		sectionNameHindi = "महेन्‍द्रनगर"
	case "MAHENDRI SIKANDARPUR":
		sectionNameHindi = "महेन्‍द्री सिकन्‍दरपुर"
	case "MAHESHAPUR BHILA":
		sectionNameHindi = "महेशपुर भीला"
	case "MAHESHPUR KHEM ANSHIK":
		sectionNameHindi = "महेशपुर खेम आंशिक"
	case "MAHESHPUR":
		sectionNameHindi = "महेशपुर"
	case "MAHESHPURA":
		sectionNameHindi = "महेशपुरा"
	case "MAHESHWARI JAT AN H.NO 1-130":
		sectionNameHindi = "महेश्‍वरीजट आं०म०सं०१ से १३०"
	case "MAHESHWARI JAT AN0H.NO 131-END":
		sectionNameHindi = "महेश्‍वरीजट आं०म०सं०१३१ से अन्‍त तक"
	case "MAHESHWARI":
		sectionNameHindi = "महेश्‍वरी"
	case "MAHEVA":
		sectionNameHindi = "महेवा"
	case "MAHMOODANGLA GAILA0":
		sectionNameHindi = "महमूदानगला गैला0"
	case "MAHMOODPUR KASWA":
		sectionNameHindi = "महमूदपुर कस्‍बा"
	case "MAHMOODPUR LAL":
		sectionNameHindi = "महमूदपुर लाल"
	case "MAHMOODPUR":
		sectionNameHindi = "महमूदपुर"
	case "MAHMUDPUR A.":
		sectionNameHindi = "महमूदपुर आं0"
	case "MAHMUDPUR JAGMAL":
		sectionNameHindi = "महमूदपुर जगमल"
	case "MAHMUDPUR MAFI":
		sectionNameHindi = "महमूदपुर माफी"
	case "MAHSANPUR":
		sectionNameHindi = "महसनपुर"
	case "MAHTOSH-1":
		sectionNameHindi = "महतोष-1"
	case "MAHTOSH-2":
		sectionNameHindi = "महतोष-2"
	case "MAHU A.":
		sectionNameHindi = "माहू आ०"
	case "MAHU NANGLI":
		sectionNameHindi = "माहू नंगली"
	case "MAHU":
		sectionNameHindi = "माहू"
	case "MAHUA KHERA SUAR":
		sectionNameHindi = "महुआखेडा स्‍वार"
	case "MAHUAA":
		sectionNameHindi = "महुआ"
	case "MAHUAKHEDA TANDA":
		sectionNameHindi = "महुआखेडा टाण्‍डा"
	case "MAHUNAGAR -1":
		sectionNameHindi = "महूनागर 1"
	case "MAHUNAGAR -2":
		sectionNameHindi = "महूनागर 2"
	case "MAHUNAGAR -3":
		sectionNameHindi = "महूनागर 3"
	case "MAHUNAGAR 01":
		sectionNameHindi = "महुनागर 01"
	case "MAHUNAGAR 02":
		sectionNameHindi = "महुनागर 02"
	case "MAHUPURA":
		sectionNameHindi = "माहुपुरा"
	case "MAIHARI SARAY AN.":
		sectionNameHindi = "मनिहारी सराय आं0"
	case "MAIHAWATPUR HEMRAJ":
		sectionNameHindi = "महाबतपुर हैमराज"
	case "MAINATHER":
		sectionNameHindi = "मैनाठेर"
	case "MAINI":
		sectionNameHindi = "मैनी"
	case "MAJAR BAGADADI SAHAB":
		sectionNameHindi = "मजार बगदादी साहब"
	case "MAJAR CHUPASHAH MIYA":
		sectionNameHindi = "मजार चुपशाह मिया"
	case "MAJAR SHAH DARAGAHI SAHAB":
		sectionNameHindi = "मजार शाह दरगाही साहब"
	case "MAJAR VAJIR ALI SAHAB":
		sectionNameHindi = "मजार वजीर अली साहब"
	case "MAJHEDASAKRU HNO 1 TO 227":
		sectionNameHindi = "मझेडा शकरू म0स0 १ से २२७ तक"
	case "MAJHEDASAKRU HNO 228 TO END":
		sectionNameHindi = "मझेडा शकरू म0स0 २२८ से अंत तक"
	case "MAJHEDASAKRU HNO 37 TO 454":
		sectionNameHindi = "मझेडा शकरू म0 स0 ३७ से ४५४ तक"
	case "MAJHOLA ANSHIK":
		sectionNameHindi = "मझौला आंशिक"
	case "MAJHOLA BILLOCH A0":
		sectionNameHindi = "मझौला बिल्‍लौच आ0"
	case "MAJHOLA PURVI":
		sectionNameHindi = "मझोला पूर्वी"
	case "MAJHOLI DEHAT ANSHIK":
		sectionNameHindi = "मझोली देहात आंशिक"
	case "MAJHOLI H NO 1 TO 50":
		sectionNameHindi = "मझोली म0 स0 1 से 50 तक"
	case "MAJHOLI H NO 51 TO END":
		sectionNameHindi = "मझोली म0 स0 51 से अन्‍त तक"
	case "MAJHOLI KHAS AANSHIK":
		sectionNameHindi = "मझोली खास आंशिक"
	case "MAJHOLI":
		sectionNameHindi = "मझौली"
	case "MAJHRA AMEER KHAN":
		sectionNameHindi = "मझरा अमीर खां"
	case "MAJHRA GHOSIPURA":
		sectionNameHindi = "मझरा घोसीपुरा"
	case "MAJHRA JANRAILVALA":
		sectionNameHindi = "मझरा जनरैलवाला"
	case "MAJHRA KALINAGAR LACHHIWALA MAJHRA KHEMWALA -1":
		sectionNameHindi = "मझरा कलीनगर ,लच्‍छीवाला , मझरा खेमवाला 1"
	case "MAJHRA KALINAGAR LACHHIWALA MAJHRA KHEMWALA -2":
		sectionNameHindi = "मझरा कलीनगर , लच्‍छीवाला व मझरा खेमवाला 2"
	case "MAJHRA MILK (DILARI)":
		sectionNameHindi = "मझरा मिलक (डिलारी)"
	case "MAJHRA SHAHPUR":
		sectionNameHindi = "मझरा शाहपुर"
	case "MAJHRA-BERKHEDA":
		sectionNameHindi = "मझरा-बेरखेडा"
	case "MAJHRA":
		sectionNameHindi = "मझरा"
	case "MAJIDGANJ AN0":
		sectionNameHindi = "मजीदगंज आं0"
	case "MAJIDGANJ":
		sectionNameHindi = "मजीदगंज"
	case "MAJISTRET COLONY KANTH ROAD":
		sectionNameHindi = "मजिस्‍ट्रेट कालोनी कांठ रोड"
	case "MAJRA BHARATPUR":
		sectionNameHindi = "मजरा भरतपुर"
	case "MAJULLANAGAR-1":
		sectionNameHindi = "माजुल्‍लानगर-1"
	case "MAJULLANAGAR-2":
		sectionNameHindi = "माजुल्‍लानगर-2"
	case "MAJULLANAGAR-3":
		sectionNameHindi = "माजुल्‍लानगर-3"
	case "MAJULLANAGAR-4":
		sectionNameHindi = "माजुल्‍लानगर-4"
	case "MAKANPUR":
		sectionNameHindi = "मकनपुर"
	case "MAKBARA DAKSHINI":
		sectionNameHindi = "मकबरा दक्षिणी"
	case "MAKBARA PRATHAM ANSHIK":
		sectionNameHindi = "मकबरा प्रथम आंशिक"
	case "MAKBARA PRATHAM":
		sectionNameHindi = "मकबरा प्रथम"
	case "MAKBARA ROAD PEERZADA":
		sectionNameHindi = "मकबरा रोड पीरजादा"
	case "MAKBARA SABJI MANDI":
		sectionNameHindi = "मकबरा सब्‍जी मण्‍डी"
	case "MAKHDOOMPUR":
		sectionNameHindi = "मखदूमपुर"
	case "MAKRANDPUR":
		sectionNameHindi = "मकरन्दपुर"
	case "MAKSHOODABAD":
		sectionNameHindi = "मकसूदाबाद"
	case "MAKSUDNAGAR DEVIDAS":
		sectionNameHindi = "मकसूदनगर देवीदास"
	case "MAKSUDPUR A.":
		sectionNameHindi = "मकसूदपुर आं0"
	case "MALAKPUR ABDULLA":
		sectionNameHindi = "मलकपुर अब्‍दुल्‍ला"
	case "MALAKPUR DEHRI":
		sectionNameHindi = "मलकपुर देहरी"
	case "MALAKPUR GAIR ABAD":
		sectionNameHindi = "मलकपुर गैर आबाद"
	case "MALAKPUR LAKHMAN":
		sectionNameHindi = "मलकपुर लखमन"
	case "MALAKPUR SEHSU":
		sectionNameHindi = "मलकपुर सहसू"
	case "MALAKPUR":
		sectionNameHindi = "मलकपुर"
	case "MALANKHEDA":
		sectionNameHindi = "मालनखेडा"
	case "MALEHPUR KHAIYYA":
		sectionNameHindi = "मल्‍हपुर खैयया"
	case "MALGOSHA":
		sectionNameHindi = "मलगोशा"
	case "MALHAPUR JANNU":
		sectionNameHindi = "मल्हपुर जन्नू"
	case "MALHAPUR LAKSHMIPUR":
		sectionNameHindi = "मल्हपुर लक्ष्मीपुर"
	case "MALHAPUR SINDHARI":
		sectionNameHindi = "मल्हपुर सिंधारी"
	case "MALHAPURA":
		sectionNameHindi = "मल्‍हपुरा"
	case "MALHIPUR MAHMUDA NAGLA":
		sectionNameHindi = "मल्हीपुर महमूदा नगला"
	case "MALHUPURA HARDODANDI":
		sectionNameHindi = "मल्‍हुपुरा हरदोडांडी"
	case "MALIPUR":
		sectionNameHindi = "मालीपुर"
	case "MALIYAN":
		sectionNameHindi = "मालियान"
	case "MALIYO WALI GALI KANJRI SARAI":
		sectionNameHindi = "मालियो वाली गली कंजरी सराय"
	case "MALKAN":
		sectionNameHindi = "मलकान"
	case "MALKAPUR SEMLI":
		sectionNameHindi = "मलकपुर सेमली"
	case "MALKAPUR":
		sectionNameHindi = "मलकपुर"
	case "MALKI KAMALPUR":
		sectionNameHindi = "मलकी कमालपुर"
	case "MALUKWALI":
		sectionNameHindi = "मलूकवाली"
	case "MALWA":
		sectionNameHindi = "मालवा"
	case "MALWADA URF MANPUR":
		sectionNameHindi = "मलवाडा उर्फ मानपुर"
	case "MAMRAJPUR":
		sectionNameHindi = "मामराजपुर"
	case "MAMURPUR BADOULI":
		sectionNameHindi = "मामूरपुर वढौली"
	case "MANAKCHAND":
		sectionNameHindi = "मानकचन्‍द"
	case "MANAKPUR ANZARIYA":
		sectionNameHindi = "मानकपुर वजरिया"
	case "MANAKPUR KUNDARKI":
		sectionNameHindi = "मानकपुर कुन्दरकी"
	case "MANAONA":
		sectionNameHindi = "मनौना"
	case "MANAPUR SHIVAPURI":
		sectionNameHindi = "मानपुर शिवपुरी"
	case "MANAVALA":
		sectionNameHindi = "मानावाला"
	case "MANAVARPYURSAID HNO 111 TO END":
		sectionNameHindi = "मुनव्‍वरपुरसैद म0स0 १११ से अंत तक"
	case "MANAVARPYURSAID HNO TO 1 TO 110":
		sectionNameHindi = "मुनव्‍वरपुरसैद म0स0 १ से ११० तक"
	case "MANDAUNRI AN.":
		sectionNameHindi = "मडौरी आं0"
	case "MANDAWALI AN0":
		sectionNameHindi = "मण्‍डावली आं0"
	case "MANDAWLI SAIDU":
		sectionNameHindi = "मण्‍डावली सैदू"
	case "MANDBA HASANPUR-2":
		sectionNameHindi = "मंडवा हसनपुर-1"
	case "MANDBA HASANPUR":
		sectionNameHindi = "मंडवा हसनपुर"
	case "MANDEYAN BALLU":
		sectionNameHindi = "मडैयान बल्‍लू"
	case "MANDEYAN RAMI":
		sectionNameHindi = "मडैयान रामी"
	case "MANDEYAN SHADI-1":
		sectionNameHindi = "मडैयान शादी -1"
	case "MANDEYAN SHADI-2":
		sectionNameHindi = "मडैयान शादी -2"
	case "MANDHOLI-1":
		sectionNameHindi = "मढोली-1"
	case "MANDI BANS":
		sectionNameHindi = "मण्‍डी बॉस"
	case "MANDI BAZAR":
		sectionNameHindi = "मन्‍डी बाजार"
	case "MANDI MARKAM GANJ":
		sectionNameHindi = "मण्‍डी मारकमगंज"
	case "MANDOLI-2":
		sectionNameHindi = "मढोली-2"
	case "MANDORA VIP":
		sectionNameHindi = "मन्‍डौरा विप"
	case "MANDORA":
		sectionNameHindi = "मण्‍डौरा"
	case "MANDORI":
		sectionNameHindi = "मन्‍डौरी"
	case "MANDRO JATT":
		sectionNameHindi = "मडौरा जट"
	case "MANGA VALA":
		sectionNameHindi = "मंगा वाला"
	case "MANGADPUR":
		sectionNameHindi = "मंगदपुर"
	case "MANGOLI 01":
		sectionNameHindi = "मंगोली 01"
	case "MANGOLI 02":
		sectionNameHindi = "मंगोली 0२"
	case "MANGOLPURA":
		sectionNameHindi = "मंगोलपुरा"
	case "MANGU CHARKHI":
		sectionNameHindi = "मंगू चर्खी"
	case "MANGUPURA":
		sectionNameHindi = "मंगूपुरा"
	case "MANIHAR KHEDA":
		sectionNameHindi = "मनिहार खेडा"
	case "MANIHARAN ANSHIK":
		sectionNameHindi = "मनिहारान आंशिक"
	case "MANIHARAN":
		sectionNameHindi = "मनिहारन"
	case "MANIHARI SARAI AN H.NO 147 -END":
		sectionNameHindi = "मनिहारी सराय आं०म०सं०१४७ से अन्‍त तक"
	case "MANIHARI SARAI AN0 H.NO 1-146":
		sectionNameHindi = "मनिहारी सराय आ०म०सं०१ से १४६"
	case "MANIRAMPUR":
		sectionNameHindi = "मनीरामपुर"
	case "MANJHLETA":
		sectionNameHindi = "मझलेटा"
	case "MANKARA":
		sectionNameHindi = "मनकरा"
	case "MANKHANDEPUR GARHI":
		sectionNameHindi = "मनकंठपुर गढी"
	case "MANKUA MAKSOODPUR":
		sectionNameHindi = "मनकुआ मकसूदपुर"
	case "MANKUAA":
		sectionNameHindi = "मनकुआ"
	case "MANKULA":
		sectionNameHindi = "मनकूला"
	case "MANNABARPUR CHAK":
		sectionNameHindi = "मनव्‍वरपुर चक"
	case "MANOHARPUR":
		sectionNameHindi = "मनोहरपुर"
	case "MANOHARWALA":
		sectionNameHindi = "मनोहरवाला"
	case "MANOTA":
		sectionNameHindi = "मनौटा"
	case "MANPUR CHHAPAT":
		sectionNameHindi = "मानपुर चापट"
	case "MANPUR DATTRAM MU.":
		sectionNameHindi = "मानपुर दत्तराम मु."
	case "MANPUR MUJAFFARPUR":
		sectionNameHindi = "मानपुर मुजफफरपुर"
	case "MANPUR NARAYANPUR ANSHIK":
		sectionNameHindi = "मानपुर नारायणपुर आंशिक"
	case "MANPUR OJHA-1":
		sectionNameHindi = "मानपुर ओझा-1"
	case "MANPUR OJHA-2":
		sectionNameHindi = "मानपुर ओझा-2"
	case "MANPUR OJHA-3":
		sectionNameHindi = "मानुपर ओझा-3"
	case "MANPUR OJHA-4":
		sectionNameHindi = "मानपुर ओझा-4"
	case "MANPUR PASHCHIMI":
		sectionNameHindi = "मानपुर पश्चिमी"
	case "MANPUR PATTI AI0 GAIRA0":
		sectionNameHindi = "मानपुर पटटी ऐ0 गैरा0"
	case "MANPUR PURVI ANSHIK":
		sectionNameHindi = "मानपुर पूर्वी आंशिक"
	case "MANPUR SABIT":
		sectionNameHindi = "मानपुर सावित"
	case "MANPUR UTTARI-1":
		sectionNameHindi = "मानपुर उत्‍तरी 1"
	case "MANPUR UTTARI-2":
		sectionNameHindi = "मानपुर उत्‍तरी 2"
	case "MANPUR":
		sectionNameHindi = "मानपुर"
	case "MANSAROVAR COLONY":
		sectionNameHindi = "मानसरोवर कालौनी"
	case "MANSHA":
		sectionNameHindi = "मन्‍शा"
	case "MANSHAPUR":
		sectionNameHindi = "मानशाहपुर"
	case "MANSOORI AVID NAGAR":
		sectionNameHindi = "मंसूरी आविद नगर"
	case "MANSOORI COLONY ANSHIK":
		sectionNameHindi = "मंसूरी कालौनी आंशिक"
	case "MANSOORPUR":
		sectionNameHindi = "मन्‍सूरपुर"
	case "MANSUR PUR":
		sectionNameHindi = "मन्सूर पुर"
	case "MANSURAPUR":
		sectionNameHindi = "मन्सूरपुर"
	case "MANSURI COLONY JAYANTIPUR":
		sectionNameHindi = "मंसुरी कालौनी जयन्‍तीपुर"
	case "MANSURPUR-1":
		sectionNameHindi = "मन्‍सूरपुर-1"
	case "MANSURPUR-2":
		sectionNameHindi = "मन्‍सूरपुर-2"
	case "MANULLAPUR":
		sectionNameHindi = "मानुल्‍लापुर"
	case "MANUNAGAR":
		sectionNameHindi = "मनुनगर"
	case "MANUPURA":
		sectionNameHindi = "मनुपुरा"
	case "MANWALA":
		sectionNameHindi = "मानवाला"
	case "MANZOLA GUJJRA A.":
		sectionNameHindi = "मझौला गुर्जर आ०"
	case "MANZOLA GURJJAR A.":
		sectionNameHindi = "मझौला गुर्जर आ०"
	case "MAQBARA AN0":
		sectionNameHindi = "मकबरा आं0"
	case "MAQSUDANPUR HAFIZ":
		sectionNameHindi = "मकसूदपुर हफीज"
	case "MARGHATI":
		sectionNameHindi = "मरघटी"
	case "MARUFPUR":
		sectionNameHindi = "मारूफपुर"
	case "MASEBI RASULAPUR":
		sectionNameHindi = "मसेबी रसूलपुर"
	case "MASEET A.":
		sectionNameHindi = "मसीत आ०"
	case "MASEVI RASULAPUR AANSHIK":
		sectionNameHindi = "मसेवी रसूलपुर आंशिक"
	case "MASJID DAROGA MAHABUB JAN":
		sectionNameHindi = "मस्जिद दरोगा महबूब जान"
	case "MASJID JASAUDI PRADHAN":
		sectionNameHindi = "मस्जिद जसौदी प्रधान"
	case "MASJID KAITHAVALI":
		sectionNameHindi = "मस्जिद कैथवाली"
	case "MASJID KALE KHAN":
		sectionNameHindi = "मस्जिद काले खाँ"
	case "MASJID KALLASH MAY GUIYA TALAB-1":
		sectionNameHindi = "मस्जिद कल्लाश मय गुईया तालाब-1"
	case "MASJID KALLASH MAY GUIYA TALAB-2":
		sectionNameHindi = "मस्जिद कल्लाश मय गुईया तालाब-2"
	case "MASJID MOLANA JAHURUDDIN KHAN 2":
		sectionNameHindi = "मस्जिद मौलाना जहूरूददीन खां 2"
	case "MASJID MOLANA JAHURUDDIN KHAN-1":
		sectionNameHindi = "मस्जिद मौलाना जहूरूददीन खां-1"
	case "MASJID QAZI":
		sectionNameHindi = "मस्जिद काजी"
	case "MASOOMPUR":
		sectionNameHindi = "मासूमपुर"
	case "MASTALLIPUR":
		sectionNameHindi = "मस्तल्लीपुर"
	case "MASUMAPUR":
		sectionNameHindi = "मासूमपुर"
	case "MASURI":
		sectionNameHindi = "मसूरी"
	case "MASWASI 5":
		sectionNameHindi = "मसवासी 5"
	case "MASWASI-1":
		sectionNameHindi = "मसवासी 1"
	case "MASWASI-2":
		sectionNameHindi = "मसवासी 2"
	case "MASWASI-3":
		sectionNameHindi = "मसवासी 3"
	case "MASWASI-4":
		sectionNameHindi = "मसवासी 4"
	case "MATA MANDIR LINEPAR":
		sectionNameHindi = "माता मंदिर लाइनपार"
	case "MATAPUR GAIR AB.AD":
		sectionNameHindi = "मतापुर गैर आं0"
	case "MATHAIYYA MAJHRA":
		sectionNameHindi = "मठैय्या मझरा"
	case "MATHANA":
		sectionNameHindi = "मथाना"
	case "MATHERA CHOHAN":
		sectionNameHindi = "मठेरा चौहान"
	case "MATHURAPUR -2":
		sectionNameHindi = "मथुरापुर -2"
	case "MATHURAPUR -3":
		sectionNameHindi = "मथुरापुर -3"
	case "MATHURAPUR 01":
		sectionNameHindi = "मथुरापुर 01"
	case "MATHURAPUR MOR AN0":
		sectionNameHindi = "मथुरापुर मोर आं0"
	case "MATIPUR URF MAINI AANSHIK":
		sectionNameHindi = "मातीपुर उर्फ मैनी आंशिक"
	case "MATKHEDA-1":
		sectionNameHindi = "माटखेडा 1"
	case "MATKHEDA-2":
		sectionNameHindi = "माटखेडा 2"
	case "MATLABPUR":
		sectionNameHindi = "मतलबपुर"
	case "MATOURA DURGA":
		sectionNameHindi = "मटौरा दर्गा"
	case "MATOURAMAN":
		sectionNameHindi = "मटौरा मान"
	case "MATWALI --1":
		sectionNameHindi = "मतवाली --1"
	case "MATWALI -2":
		sectionNameHindi = "मतवाली -2"
	case "MAU DEHAT ANSHIK":
		sectionNameHindi = "मऊ देहात आंशिक"
	case "MAU. BARADRI MAHAMUD KHAN":
		sectionNameHindi = "मौ. बारादरी महमूद खां"
	case "MAU":
		sectionNameHindi = "मउ"
	case "MAUHADA AN.":
		sectionNameHindi = "मोहडा आं0"
	case "MAUHADI":
		sectionNameHindi = "मोहडी"
	case "MAUHAMMAD IBRAHEEMPUR":
		sectionNameHindi = "मौहम्मद इब्राहीमपुर"
	case "MAUHAMMAD JAMAPUR":
		sectionNameHindi = "मौहम्मद जमापुर"
	case "MAUJAMPUR JAITRA AN.":
		sectionNameHindi = "मौजमपुर जैतरा आं0"
	case "MAUJMAPUR SURAJ":
		sectionNameHindi = "मौजमपुर सूरज"
	case "MAUSMAPUR":
		sectionNameHindi = "मौसमपुर"
	case "MAUZAMPUR SUJAN":
		sectionNameHindi = "मौजमपुर सुजान"
	case "MAYAPURI":
		sectionNameHindi = "मायापुरी"
	case "MAZHARPUR":
		sectionNameHindi = "मजहरपुर"
	case "MAZRA TRILOK PUR":
		sectionNameHindi = "मजरा त्रिलोकपुर"
	case "MDEAYAN GOR":
		sectionNameHindi = "मडैयान गौर"
	case "MDEYAN BUDHPUR":
		sectionNameHindi = "मडैयान वुधपुर"
	case "MEBLA":
		sectionNameHindi = "मेवला"
	case "MEDNIPUR":
		sectionNameHindi = "मेदनीपुर"
	case "MEDPURASULTAN":
		sectionNameHindi = "मेदपुरा सुल्‍तान"
	case "MEDUWALA":
		sectionNameHindi = "मेदूवाला"
	case "MEENA NAGAR ANSHIK JAYANTIPUR":
		sectionNameHindi = "मीना नगर आंशिक जयन्‍तीपुर"
	case "MEENA NAGAR":
		sectionNameHindi = "मीना नगर"
	case "MEENA NGAR JAYANTIPUR":
		sectionNameHindi = "मीना नगर जयन्‍तीपुर"
	case "MEERAMPUR BEGA":
		sectionNameHindi = "मीरमपुर बेगा"
	case "MEERAPUR KHADAR":
		sectionNameHindi = "मीरापुर खादर"
	case "MEERAPUR MEERGANJ-1":
		sectionNameHindi = "मीरापुर मीरगंज 1"
	case "MEERAPUR MODIWALA AN0":
		sectionNameHindi = "मीरापुर मोदीवाला आ0"
	case "MEERAPUR RAZA":
		sectionNameHindi = "मीरापुर रजा"
	case "MEERAPUR SEEKRI":
		sectionNameHindi = "मीरापुर सीकरी"
	case "MEERAPUR VEERAN":
		sectionNameHindi = "मीरापुर वीरान"
	case "MEERAPUR":
		sectionNameHindi = "मीरापुर"
	case "MEEROPUR AGRI":
		sectionNameHindi = "मीरोपुर अगरी"
	case "MEERPUR GHASI":
		sectionNameHindi = "मीरपुर घासी"
	case "MEERPUR MAJHOLI DEHAT":
		sectionNameHindi = "मीरपुर मझोली देहात"
	case "MEERZAPUR BAKENA":
		sectionNameHindi = "मिर्जापुर बकैना"
	case "MEGHANAGLA JADID":
		sectionNameHindi = "मेघानगला जदीद"
	case "MEGHANAGLA KADIM":
		sectionNameHindi = "मेघानगला कदीम"
	case "MEHAKMA UTTARI":
		sectionNameHindi = "महकमा उत्‍तरी"
	case "MEHANDI KA TIRAHA ASALATPURA":
		sectionNameHindi = "मेहंदी का तिराहा असालतपुरा"
	case "MEHANDI NAGAR JUNUBI":
		sectionNameHindi = "मंहन्‍दी नगर जुनूबी"
	case "MEHAR ALIPUR":
		sectionNameHindi = "मेहर अलीपुर"
	case "MEHARPUR":
		sectionNameHindi = "मेहरपुर"
	case "MEHARSHI DAYANAND SURAJ NAGAR":
		sectionNameHindi = "महर्षि दयानन्‍द कालोनी सूरज नगर"
	case "MEHBULLA GANJ ANSHIK":
		sectionNameHindi = "महबुल्‍ला गंज आंशिक"
	case "MEHBULLA GANJ PURVI ANSHIK":
		sectionNameHindi = "महबुल्‍ला गंज पूर्वी आंशिक"
	case "MEHDOOD KALMI":
		sectionNameHindi = "महदूद कलमी"
	case "MEHEL KUMEDAN":
		sectionNameHindi = "महल कुमेदान"
	case "MEHLAKPUR MAFI":
		sectionNameHindi = "महलकपुर माफी"
	case "MEHLAKPUR NIJAMPUR":
		sectionNameHindi = "महलकपुर निजामपुर"
	case "MEHM‍DINGAR":
		sectionNameHindi = "मेहम्‍दीनगर"
	case "MEHMODPUR KAMIL":
		sectionNameHindi = "महमूदपुर कामिल"
	case "MEHMOODPUR KESHO":
		sectionNameHindi = "महमूदपुर केशो"
	case "MEHMOODPUR MAFI":
		sectionNameHindi = "महमूदपुर माफी"
	case "MEHMOODPUR TIGRI ANSHIK":
		sectionNameHindi = "महमूदपुर तिगरी आंशिक"
	case "MEHMOODPURA PACHMI":
		sectionNameHindi = "महमूदपुरा (पश्चिमी)"
	case "MEHMUDA KHADER":
		sectionNameHindi = "महमूदा खादर"
	case "MEHMUDPUR BHAVTA AN.H.NO 1-65":
		sectionNameHindi = "महमूदपुर भावता आं०म०सं०१से ६५तक"
	case "MEHMUDPUR BHAVTA AN.H.NO 66-END":
		sectionNameHindi = "महमूदपुर भावता आं०म०सं०६६ से अंत तक"
	case "MEHMUDPUR BUZURG":
		sectionNameHindi = "महमूदपुर बुजुर्ग"
	case "MEHMUDPUR NARAYAN AN.1-102":
		sectionNameHindi = "महमूदपुर नारायण आं०म०सं०१से १०२ तक"
	case "MEHMUDPUR NARAYANAN.H.NO 103-END":
		sectionNameHindi = "महमूदपुर नरायन आ०म०सं०103 से अंत तक"
	case "MEHMUDPUR":
		sectionNameHindi = "महमूदपुर"
	case "MEHNDI BAGH":
		sectionNameHindi = "मेहन्‍दी बाग"
	case "MEHNDINAGAR SHUMALI":
		sectionNameHindi = "मेंहदीनगर शुमाली"
	case "MEHNDINAGAR":
		sectionNameHindi = "मेंहदीनगर"
	case "MEHNDIPUR":
		sectionNameHindi = "मेंहदीपुर"
	case "MEMAN SADAT AN.H.NO 1-263":
		sectionNameHindi = "मेमन सादात आ०म०सं०१ से २६३तक"
	case "MEMAN SADAT AN.H.NO 1-372":
		sectionNameHindi = "मेमन सादात आ०म०सं०१से ३७२तक"
	case "MEMAN SADAT AN.H.NO 373-END":
		sectionNameHindi = "मेमन सादात आ०म०सं०३७३ से अन्‍त तक"
	case "MEMANSADAT AN.H.NO0264-END":
		sectionNameHindi = "मेमन सादात आ०म०सं०२६४से अन्‍त तक"
	case "MEMRAN":
		sectionNameHindi = "मेमरान"
	case "MEMUDPUR BHIKKA":
		sectionNameHindi = "महमूदपुर भिक्‍का"
	case "MEMUDPUR MILAK":
		sectionNameHindi = "महमूदपुर मिलक"
	case "MENATHER ANSHIK":
		sectionNameHindi = "मैनाठेर आंशिक"
	case "MENDIPUR":
		sectionNameHindi = "मेंहदीपुर"
	case "MENHDINAGAR":
		sectionNameHindi = "मेंहन्‍दीनगर"
	case "MERJAPUR BALA A.":
		sectionNameHindi = "मिर्जापुर वेला आ०"
	case "MERJAPUR BALA":
		sectionNameHindi = "मिर्जापुर वेला आ०"
	case "MERJAPUR KHADAR":
		sectionNameHindi = "मिर्जापुर खादर"
	case "METHA SAID":
		sectionNameHindi = "मीठा सईद"
	case "MEWA JATT A.":
		sectionNameHindi = "मेवा जट आ0"
	case "MEWA NAWADA A.":
		sectionNameHindi = "मेवा नवादा आ0"
	case "MEWALA DHARU":
		sectionNameHindi = "मेवला धारू"
	case "MEWLA KALA":
		sectionNameHindi = "मेवला कला"
	case "MEWLA MAFI":
		sectionNameHindi = "मेवला माफी"
	case "MH. PUR BHAJJA":
		sectionNameHindi = "मौ0पुर भज्‍जा"
	case "MH. PUR URF JOGIPURA":
		sectionNameHindi = "मौ0पुर उर्फ जोगीपुरा"
	case "MIAN GANJ":
		sectionNameHindi = "मियां गंज"
	case "MIJARPUR PALLA":
		sectionNameHindi = "मिर्जापुर पल्‍ला"
	case "MILAK AARIF":
		sectionNameHindi = "मिलक आरिफ"
	case "MILAK ABBU KHAN":
		sectionNameHindi = "मिलक अब्‍बू खां"
	case "MILAK AHLADADPUR KARAR":
		sectionNameHindi = "मिलक अहलादपुर करार"
	case "MILAK AMAWATI":
		sectionNameHindi = "मिलक अमावती"
	case "MILAK ASAD KHAN":
		sectionNameHindi = "मिलक असद खां"
	case "MILAK ASDULLAPUR-1":
		sectionNameHindi = "मिलक असदुल्‍लापुर-1"
	case "MILAK ASDULLAPUR-2":
		sectionNameHindi = "मिलक असदुल्‍लापुर-2"
	case "MILAK ASDULLAPUR-3":
		sectionNameHindi = "मिलक असदुल्‍लापुर-3"
	case "MILAK ASDULLAPUR-4":
		sectionNameHindi = "मिलक असदुल्‍लापुर-4"
	case "MILAK ASDULLAPUR-5":
		sectionNameHindi = "मिलक असदुल्‍लापुर-5"
	case "MILAK AZAM BALI (BHIKANPUR BAGHA)":
		sectionNameHindi = "मिलक आजम वाली (भीकनपुरबघा)"
	case "MILAK BANKERALI":
		sectionNameHindi = "मिलक बांकरअली"
	case "MILAK BENIRAM":
		sectionNameHindi = "मिलक बेनीराम"
	case "MILAK BHOLA SINGH SONAKPUR":
		sectionNameHindi = "मिलक भोला सिंह सोनकपुर"
	case "MILAK BHURE KHAN":
		sectionNameHindi = "मिलक भूरे खां"
	case "MILAK DUNDI":
		sectionNameHindi = "मिलक दुन्‍दी"
	case "MILAK DYUDI INAYAT RASOOL KHAN":
		sectionNameHindi = "मिलक डयूढी इनायत रसूल खां"
	case "MILAK FEJULLANAGAR":
		sectionNameHindi = "मिलक फैजुल्‍लानगर"
	case "MILAK GANGODA JAT AN.":
		sectionNameHindi = "मिलक गंगोडा जट आं0"
	case "MILAK GOVIN‍DPUR KHURD":
		sectionNameHindi = "मिलक गोविन्‍दपुर खुर्द"
	case "MILAK GULAM ALI":
		sectionNameHindi = "मिलक गुलामअली"
	case "MILAK GULAM KHAN":
		sectionNameHindi = "मिलक गुलाम खां"
	case "MILAK HAFIZ BALLI":
		sectionNameHindi = "मिलक हाफिज वल्‍ली"
	case "MILAK HAJIRO BALI":
		sectionNameHindi = "मिलक हजीरो बाली"
	case "MILAK HASHAM IBRAHEEM KHAN":
		sectionNameHindi = "मिलक हाशम इब्रहीम खां"
	case "MILAK HASHAM":
		sectionNameHindi = "मिलक हाशम"
	case "MILAK ICHHCARAM":
		sectionNameHindi = "मिलक इच्‍छाराम"
	case "MILAK JAHAGIRABAD":
		sectionNameHindi = "मिलक जहागीराबाद"
	case "MILAK KALYANPUR":
		sectionNameHindi = "मिलक कल्‍यानपुर"
	case "MILAK KAMRU":
		sectionNameHindi = "मिलक कामरू"
	case "MILAK KATHAIR ASALATNAGAR BAGHA":
		sectionNameHindi = "मिलक कठैर असालतनगर बघा"
	case "MILAK KHANAM":
		sectionNameHindi = "मिलक खानम"
	case "MILAK KHANPUR":
		sectionNameHindi = "मिलक खानपुर"
	case "MILAK KHOOD":
		sectionNameHindi = "मिलक खौद"
	case "MILAK MAFI PAIPURA":
		sectionNameHindi = "मिलक माफी पईपुरा"
	case "MILAK MAJBOOTA":
		sectionNameHindi = "मिलक मजबूता"
	case "MILAK MALANG GAAN":
		sectionNameHindi = "मिलक मलंगान"
	case "MILAK MIRZA FAYYAZ":
		sectionNameHindi = "मिलक मिर्जा फययाज"
	case "MILAK MOHABBATPUR":
		sectionNameHindi = "मिलक मौहब्‍बतपुर"
	case "MILAK MOHAMMAD BAKSH":
		sectionNameHindi = "मिलक मोहम्‍मद बक्‍श"
	case "MILAK MOHMMAD NAKI":
		sectionNameHindi = "मिलक मौहम्‍मद नकी"
	case "MILAK MOHMMAD SHAH KHAN":
		sectionNameHindi = "मिलक मौहम्‍मद शाह खां"
	case "MILAK MUDIA":
		sectionNameHindi = "मिलक मुडिया"
	case "MILAK MUGLA URF MUGLPIUR":
		sectionNameHindi = "मिलक मुगला ऊर्फ मुगलपुरा"
	case "MILAK MUKEEMPUR":
		sectionNameHindi = "मिलक मुकीमपुर"
	case "MILAK MUNDI":
		sectionNameHindi = "मिलक मुण्‍डी"
	case "MILAK NAGALI":
		sectionNameHindi = "मिलक नगली"
	case "MILAK NASIRABAD-1":
		sectionNameHindi = "मिलक नसीराबाद-1"
	case "MILAK NASIRABAD-2":
		sectionNameHindi = "मिलक नसीराबाद-2"
	case "MILAK NASIRABAD-3":
		sectionNameHindi = "मिलक नसीराबाद-3"
	case "MILAK NIBBI SINGH":
		sectionNameHindi = "मिलक निब्‍बी सिंह"
	case "MILAK NOUKHAREED":
		sectionNameHindi = "मिलक नौ खरीद स्‍वार"
	case "MILAK PEEPLILAL":
		sectionNameHindi = "मिलक पीपलीलाल"
	case "MILAK PIPALSANA":
		sectionNameHindi = "मिलक पीपलसाना"
	case "MILAK QAZI":
		sectionNameHindi = "मिलक काजी"
	case "MILAK SAKTALPUR":
		sectionNameHindi = "मिलक सकतलपुर"
	case "MILAK SANYYA":
		sectionNameHindi = "मिलक सनईया"
	case "MILAK SHADI NAGAR":
		sectionNameHindi = "मिलक शादी नगर"
	case "MILAK SIKARAUL":
		sectionNameHindi = "मिलक सिकरौल"
	case "MILAK SIRASWA HARCHAND":
		sectionNameHindi = "मिलक सिरसवां हरचन्‍द"
	case "MILAK TAHABBAR ALI KHAN":
		sectionNameHindi = "मिलक तहब्‍बर अली खां"
	case "MILAK TAJ KHAN":
		sectionNameHindi = "मिलक ताज खां"
	case "MILAK TURKHEDA":
		sectionNameHindi = "मिलक तुरखेडा"
	case "MILAK WADULLA":
		sectionNameHindi = "मिलक वादुल्‍ला"
	case "MILAK YAKOOB ALI KHAN":
		sectionNameHindi = "मिलक याकूब अली खां"
	case "MILAKIYAN AN.":
		sectionNameHindi = "मिल्कियान आं0"
	case "MILAL AFZAL KHAN":
		sectionNameHindi = "मिलक अफजल खां"
	case "MILAN VIHAR ANSHIK":
		sectionNameHindi = "मिलन विहार आंशिक"
	case "MILAN VIHAR":
		sectionNameHindi = "मिलन विहार"
	case "MILETRI LAIN-1":
		sectionNameHindi = "मिलेट्री लाइन-1"
	case "MILETRI LAIN-2":
		sectionNameHindi = "मिलेट्री लाइन-2"
	case "MILK DHOBI VALI":
		sectionNameHindi = "मिलक धोबी वाली"
	case "MILK KHARAGPUR BAJE":
		sectionNameHindi = "मिलक खरगपुर बाजे"
	case "MILK MEVATI":
		sectionNameHindi = "मिलक मेवाती"
	case "MILKIYAN":
		sectionNameHindi = "मिल्‍कीयान"
	case "MIMLA":
		sectionNameHindi = "मीमला"
	case "MIRAMPUR DURG":
		sectionNameHindi = "मीरमपुर दुर्ग"
	case "MIRAPUR BILASPUR":
		sectionNameHindi = "मीरापुर बिलासपुर"
	case "MIRAPUR MAFI":
		sectionNameHindi = "मीरापुर माफी"
	case "MIRAPUR MEERGANJ -2":
		sectionNameHindi = "मीरापुर मीरगंज 2"
	case "MIRAPUR MOHAN CHAK J0 MU0":
		sectionNameHindi = "मीरपुर मोहन चक ज0 मु0"
	case "MIRDAGAN":
		sectionNameHindi = "मिर्दगान"
	case "MIRDGAN B-10":
		sectionNameHindi = "मिर्दगान बी0 10"
	case "MIRDGAN B-11":
		sectionNameHindi = "मिर्दगान बी0 11"
	case "MIRDGAN":
		sectionNameHindi = "मिर्दगान"
	case "MIRGIPUR":
		sectionNameHindi = "मिरगीपुर"
	case "MIRJAPUR CHAKARPUR":
		sectionNameHindi = "मिर्जापुर चकरपुर"
	case "MIRJAPUR DHIKLI,":
		sectionNameHindi = "मिर्जापुर ढीकली,"
	case "MIRJAPUR KARIMUDDIN":
		sectionNameHindi = "मिर्जापुर करीमुददीन"
	case "MIRJAPUR PARAM":
		sectionNameHindi = "मिर्जापुर परम"
	case "MIRZALIPUR BHARA":
		sectionNameHindi = "मिर्जालीपुर भारा"
	case "MIRZALIPUR CHOHAD":
		sectionNameHindi = "मिर्जालीपुरचौहड"
	case "MIRZAPUR BANGAR":
		sectionNameHindi = "मिर्जापुर बंगर"
	case "MIRZAPUR BILASPUR":
		sectionNameHindi = "मिर्जापुर बिलासपुर"
	case "MIRZAPUR MAHESH":
		sectionNameHindi = "मिर्जापुर महेश"
	case "MIRZAPUR RANGILA":
		sectionNameHindi = "मिर्जापुर रंगीला"
	case "MIRZAPUR SAID":
		sectionNameHindi = "मिर्जापुर सैद"
	case "MIRZAPUR SUAR":
		sectionNameHindi = "मिर्जापुर स्‍वार"
	case "MIRZAPUR":
		sectionNameHindi = "मिर्जापुर"
	case "MISHRINAGAR JUNUBI":
		sectionNameHindi = "मिश्रीनगर जनूबी"
	case "MISHRIPUR":
		sectionNameHindi = "मिश्रीपुर"
	case "MISRINAGAR SHUMALI":
		sectionNameHindi = "मिश्रीनगर शुमाली"
	case "MISRIPUR":
		sectionNameHindi = "मिश्रीपुर"
	case "MISTAN GANJ AKAB MISTAN GANJ":
		sectionNameHindi = "मिस्टन गंज अकब मिस्टन गंज"
	case "MITHAI":
		sectionNameHindi = "मिठाई"
	case "MITHAN KUWAR PRATAP SINGH":
		sectionNameHindi = "मिठान कुवंर प्रताप सिंह"
	case "MITHANPUR BALLU URF NAGLA":
		sectionNameHindi = "मिठनपुर बल्लू उर्फ नगला"
	case "MITHANPUR MAHESH":
		sectionNameHindi = "मिठनपुर महेश"
	case "MITHANPUR SHOBHARAM":
		sectionNameHindi = "मिठान शोभराम"
	case "MITHOPUR":
		sectionNameHindi = "मि‍ठठोपुर"
	case "MITTANPUR":
		sectionNameHindi = "मिठनपुर"
	case "MITTARPUR AHROULA 01":
		sectionNameHindi = "मित्‍तरपुर अहरौला ०१"
	case "MITTARPUR AHROULA 02":
		sectionNameHindi = "मित्‍तरपुर अहरौला ०2"
	case "MITTHEPUR":
		sectionNameHindi = "मिटठेपुर"
	case "MIYAZI MOKHA":
		sectionNameHindi = "मियाजी मोखा"
	case "MO ALIPUR SHEKH":
		sectionNameHindi = "मौ०अलीपुर शेख"
	case "MO. ANSARIYAN OMRI KLA":
		sectionNameHindi = "मौ0 अन्‍सारियान उमरी कलां"
	case "MO. AVIDPURA OMRI KLA":
		sectionNameHindi = "मौ0 आविदपुरा उमरी कलां"
	case "MO. BIDATTSHAH OMRI KLA":
		sectionNameHindi = "मौ0 विदतशाह उमरी कलां"
	case "MO. CHODHRIYAN OMRI KLA":
		sectionNameHindi = "मौ0 चौधरियान उमरी कलां"
	case "MO. FAREEDGANJ OMRI KLA":
		sectionNameHindi = "मौ0 फरीदगंज उमरी कलां"
	case "MO. NANGLIYAN OMRI KLA":
		sectionNameHindi = "मौ0 नंगलियान उमरी कलां"
	case "MO. NASERABAD OMRI KLA":
		sectionNameHindi = "मौ0 नसीरावाद उमरी कलां"
	case "MO. PACHEYAN ORMI KLA":
		sectionNameHindi = "मौ0 पर्छयान उमरी कलां"
	case "MO. TEKEDARAN OMRI KLA":
		sectionNameHindi = "मौ0 ठेकेदारान उमरी कलां"
	case "MO.ALIPUR MUKTA":
		sectionNameHindi = "मौ0अलीपुर मुक्‍ता"
	case "MO.ALIPUR SUKHANAND":
		sectionNameHindi = "मौ०अलीपुर सुखानन्‍द"
	case "MO.HUSAINPUR GAR ABAD":
		sectionNameHindi = "मौ0 हुसेनपुर गै0आ0"
	case "MO.KULIPUR":
		sectionNameHindi = "मौ0कुलीपुर"
	case "MO.PUR KHADER":
		sectionNameHindi = "मौ० पुर खादर"
	case "MO.PUR LAL":
		sectionNameHindi = "मौ0पुर लाल"
	case "MOCHIPURA":
		sectionNameHindi = "मोचीपुरा"
	case "MODHA":
		sectionNameHindi = "मौढा"
	case "MODHIA BHOGAN":
		sectionNameHindi = "मोढियो भोगन"
	case "MODHIA DHANSHI":
		sectionNameHindi = "मोढियो धनीस"
	case "MODPURTIRLOK AN HNO 1 TO 163":
		sectionNameHindi = "मौ0पुर ति‍रलोक आ0 म0स0 १ से १६३ तक"
	case "MODPURTIRLOK AN HNO 164 TO END":
		sectionNameHindi = "मौ0पुर ति‍रलोक आ0 म0स0 १६४से अंततक"
	case "MOH.PUR PARMA":
		sectionNameHindi = "मौ0पुर परमा"
	case "MOHABBAT NAGAR KUNDESRA":
		sectionNameHindi = "मोहब्‍बत नगर कुण्‍डेसरा"
	case "MOHABBATNAGAR NEAR SUAR":
		sectionNameHindi = "मौहब्‍बतनगर नि0 स्‍वार"
	case "MOHABBATPUR BHAGWANTPUR AITMALI (GAIRA)":
		sectionNameHindi = "मौहब्‍बतपुर भगवन्‍तपुर ऐतमाली (गैरा)"
	case "MOHALIYA":
		sectionNameHindi = "मोहलिया"
	case "MOHALLA DANG":
		sectionNameHindi = "मौहल्‍ला दॉग"
	case "MOHALLA SAHU":
		sectionNameHindi = "मौहल्‍ला साहू"
	case "MOHALLA SAINI BASTI":
		sectionNameHindi = "मौहल्‍ला सैनी बस्‍ती"
	case "MOHAMADPUR DEOMAL":
		sectionNameHindi = "मौ0पुर देवमल"
	case "MOHAMADPUR LAKHU":
		sectionNameHindi = "मौ0पुर लक्‍खू"
	case "MOHAMMAD HAYATAPUR":
		sectionNameHindi = "मौहम्मद हयातपुर"
	case "MOHAMMAD NAGAR KHUTIA":
		sectionNameHindi = "मोहम्‍मदनगर खुटिया"
	case "MOHAMMAD NAGAR":
		sectionNameHindi = "मोहम्‍मदनगर"
	case "MOHAMMADAPUR":
		sectionNameHindi = "मौहम्मदपुर"
	case "MOHAMMADGUNJ":
		sectionNameHindi = "मौहम्‍मदगंज"
	case "MOHAMMADNAGAR NANKAR":
		sectionNameHindi = "मोहम्‍मदनगर नानकार"
	case "MOHAMMADNAGAR":
		sectionNameHindi = "मौहम्‍मदनगर"
	case "MOHAMMADPUR BASTAUR AANSHIK":
		sectionNameHindi = "मौहम्मदपुर बस्तौर आंशिक"
	case "MOHAMMADPUR DHYAN SINGH":
		sectionNameHindi = "मौहम्‍मदपुर ध्‍यान सिंह"
	case "MOHAMMADPUR JADID":
		sectionNameHindi = "मोहम्‍मदपुर जदीद"
	case "MOHAMMADPUR KADIM":
		sectionNameHindi = "मोहम्‍मदपुर कदीम"
	case "MOHAN PUR":
		sectionNameHindi = "मोहनपुर"
	case "MOHANPUR":
		sectionNameHindi = "मोहनपुर"
	case "MOHBBATGANJ TANDA":
		sectionNameHindi = "मोहब्‍बतगंज टाण्‍डा"
	case "MOHD ALIPUR DWARKA":
		sectionNameHindi = "मौ0 अलीपुर द्वारका"
	case "MOHD ALIPUR PATTAGHAT":
		sectionNameHindi = "मौ0 अलीपुर पटटाघाट"
	case "MOHD ASGARPUR":
		sectionNameHindi = "मौ०असगरपुर"
	case "MOHD ASHIKPUR BHUREY":
		sectionNameHindi = "मौहम्‍मद आशि कपुर भूरे"
	case "MOHD ASHIKPURKAMALNAIN":
		sectionNameHindi = "मौ०आशिकपुर कमलनैन"
	case "MOHD PUR ATA":
		sectionNameHindi = "मौ0 पुर अता"
	case "MOHD PUR RAVA":
		sectionNameHindi = "मौ0पुर रवा"
	case "MOHD PUR SHUMALI":
		sectionNameHindi = "मो0 पुर शुमाली"
	case "MOHD PUR TAKI":
		sectionNameHindi = "मौहम्‍मदपुर तकी"
	case "MOHD SADIKPUR":
		sectionNameHindi = "मौ0 सादिकपुर"
	case "MOHD. ALAMPUR":
		sectionNameHindi = "मौ0 आलमपुर"
	case "MOHD. ALIPUR CHOHAR":
		sectionNameHindi = "मौ0 अलीपुर चौहड"
	case "MOHD. ALIPUR HIRDAY AN0":
		sectionNameHindi = "मौ0 अलीपुर हृदय आं0"
	case "MOHD. ALIPUR INAYAT A.":
		sectionNameHindi = "मौ० अलीपुर इनायत आ०"
	case "MOHD. ALIPUR PARMA":
		sectionNameHindi = "मौ0 अलीपुर परमा"
	case "MOHD. ALIPUR VEERBHAN A":
		sectionNameHindi = "मौ0 अलीपुर वीरभान ए"
	case "MOHD. ALIPUR VEERBHAN":
		sectionNameHindi = "मौ0 अलीपुर वीरभान"
	case "MOHD. AMI KHANPUR":
		sectionNameHindi = "मौ0 अमी खानपुर"
	case "MOHD. AMIPUR":
		sectionNameHindi = "मौ0 अमीपुर"
	case "MOHD. HUSANPUR MAJRA ORANGABAD":
		sectionNameHindi = "मौ० हुसैनपुर मजरा औरगांबाद"
	case "MOHD. TAHARPUR":
		sectionNameHindi = "मौ0 ताहरपुर"
	case "MOHD. TAKIPUR GHASI":
		sectionNameHindi = "मौ0 तकीपुर घासी"
	case "MOHD.ALIPUR ABHYACHAND":
		sectionNameHindi = "मौ0अलीपुर अभयचन्‍द"
	case "MOHD.ALIPUR ALIMUDDIN":
		sectionNameHindi = "मौ0अलीपुर अलीमुददीन"
	case "MOHD.ALIPUR BHIKAN":
		sectionNameHindi = "मौ0अलीपुर भिक्‍कन"
	case "MOHD.ALIPUR INAYAT":
		sectionNameHindi = "मौ0अलीपुर इनायत"
	case "MOHD.ALIPUR LALMAN URF RAMSAHAYVALA":
		sectionNameHindi = "मौ0अलीपुर लालमन उर्फ रामसहायवाला"
	case "MOHD.PUR JAMAL":
		sectionNameHindi = "मौ0पुर जमाल"
	case "MOHD.PUR SADA":
		sectionNameHindi = "मौ0पुर सादा"
	case "MOHD.PUR SULTAN":
		sectionNameHindi = "मौ0पुर सुल्‍तान"
	case "MOHD.PUR VEERU":
		sectionNameHindi = "मौ0पुर वीरू"
	case "MOHDI HAJRATPUR":
		sectionNameHindi = "मौहडी हजरतपुर"
	case "MOHDPUR ROJORE HNO 1 TO 159":
		sectionNameHindi = "मौ0पुर राजौरी आं0 म0 स0 १ से 159"
	case "MOHDPUR ROJORE HNO 1 TO 254":
		sectionNameHindi = "मौ0पुर राजौरी आं0 म0 स0 १ से 254"
	case "MOHDPUR ROJORE HNO 255 TO END":
		sectionNameHindi = "मौ0पुर राजौरी आं0 म0 स0255 से अंत तक"
	case "MOHIDINPUR":
		sectionNameHindi = "मोहीदीनपुर"
	case "MOHIUDDINPUR KHALSA":
		sectionNameHindi = "मुहीउददीनपुर खालसा"
	case "MOHIUDDINPUR":
		sectionNameHindi = "मोहीउददीनपुर"
	case "MOHMADPUR MANDAWALI":
		sectionNameHindi = "मौ0पुर मण्‍डावली"
	case "MOHM‍MADPUR BAS‍TAUR":
		sectionNameHindi = "मौहम्‍मदपुर बस्‍तौर आंश्किा"
	case "MOHNPUR AANSHIK":
		sectionNameHindi = "मोहनपुर आंशिक"
	case "MOHSANPUR CHAMRA":
		sectionNameHindi = "मोहसनपुर चमरा"
	case "MOHSANPUR KALYAN":
		sectionNameHindi = "मोहसनपुर कल्‍याण"
	case "MOHSANPUR":
		sectionNameHindi = "मोहसनपुर"
	case "MOINUDDINPUR":
		sectionNameHindi = "मोईनूददीनपुर"
	case "MOJAMPUR HARVANSH":
		sectionNameHindi = "मौजमपुरहरवंश"
	case "MOJAMPURDAYAL":
		sectionNameHindi = "मौजमपुर दयाल"
	case "MOJIPUR DHARMA":
		sectionNameHindi = "मौजीपुर धर्मा"
	case "MOJJAMPUR SADAT":
		sectionNameHindi = "मौज्‍जमपुर सादात"
	case "MOJJAMPUR TULSI AN0":
		sectionNameHindi = "मौज्‍जमपुर तुलसी आं0"
	case "MOJJAMPUR TULSI":
		sectionNameHindi = "मौज्‍जमपुर तुलसी आं0"
	case "MOLANA MOHAMMAD ALI JOHAR MARG MAY GANGAPUR HATS-9":
		sectionNameHindi = "मौलाना मोहम्‍मद अली जौहर मार्ग मय गंगापुर हट्स-9"
	case "MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HARTS 8":
		sectionNameHindi = "मौलाना मोहम्‍मद अली जौहर मार्ग मय गंगापुर हट्स-8"
	case "MOLANA MOHAMMAD ALI JOHAR MARG MEY GANGAPUR HATS-7":
		sectionNameHindi = "मौलाना मोहम्‍मद अली जौहर मार्ग मय गंगापुर हट्स-7"
	case "MOLHARPUR":
		sectionNameHindi = "मोल्‍हडपुर"
	case "MOLVIYAN A.":
		sectionNameHindi = "मौलवियान आ0"
	case "MOLVIYAN JUNUBI":
		sectionNameHindi = "मोलवियान जनूबी"
	case "MOLVIYAN SHUMALI":
		sectionNameHindi = "मोलवियान शुमाली"
	case "MOLVIYAN":
		sectionNameHindi = "मौलवियान"
	case "MOMIN NAGAR":
		sectionNameHindi = "मोमि‍ननगर"
	case "MOMINPUR AHAMDABAD-1":
		sectionNameHindi = "मोमिनपुर अहमदाबाद -1"
	case "MOMINPUR AHAMDABAD-2":
		sectionNameHindi = "मोमिनपुर अहमदाबाद -2"
	case "MOMINPUR DARGO":
		sectionNameHindi = "मोमिनपुर दरगां"
	case "MOMINPUR DASU":
		sectionNameHindi = "मोमीनपुर दासू"
	case "MOMINPUR GULARI":
		sectionNameHindi = "मोमिनपुर गिलाडी"
	case "MOMINPUR LAL KUWAR":
		sectionNameHindi = "मोमिनपुर लाल कुवर"
	case "MOMINPUR LALU":
		sectionNameHindi = "मोमीनपुर लालू"
	case "MORA MUSTAHKAM PURANI AWADI":
		sectionNameHindi = "मोरा मुस्‍तहकम पुरानी आवादी"
	case "MORA NAYEE AWADI":
		sectionNameHindi = "मोरा नई आवादी"
	case "MORGHANDI":
		sectionNameHindi = "मोरझडी"
	case "MORI GET-1":
		sectionNameHindi = "मोरी गेट-1"
	case "MORI GET-2":
		sectionNameHindi = "मोरी गेट-2"
	case "MORI GET-3":
		sectionNameHindi = "मोरी गेट-3"
	case "MORLA":
		sectionNameHindi = "मोरला"
	case "MORMAKDUMPUR":
		sectionNameHindi = "मोरमकदूमपुर"
	case "MORNA A.":
		sectionNameHindi = "मौरना आं0"
	case "MORNA A0":
		sectionNameHindi = "मौरना आं0"
	case "MOTA DHAK":
		sectionNameHindi = "मोटा ढाक"
	case "MOTHALA":
		sectionNameHindi = "मोथला"
	case "MOTHEPUR KIRAT":
		sectionNameHindi = "मोथेपुर किरत"
	case "MOTI MASJID":
		sectionNameHindi = "मोती मस्जिद"
	case "MOTIPURA":
		sectionNameHindi = "मोतीपुरा"
	case "MU.CHAUDHARIYAN AN.":
		sectionNameHindi = "मु0चौधरियान आं0"
	case "MU.CHAUDHARIYAN":
		sectionNameHindi = "मु0चौधरियान"
	case "MUBANA":
		sectionNameHindi = "मुवाना"
	case "MUBARAEKPUR A.":
		sectionNameHindi = "मुबारकपुर आ०"
	case "MUBARAEKPUR A":
		sectionNameHindi = "मुबारकपुर आ०"
	case "MUBARAEKPUR KALA":
		sectionNameHindi = "मुबारपुर कला"
	case "MUBARAEKPUR POTA":
		sectionNameHindi = "मुबारकपुर पौटा"
	case "MUBARAKPUR AN0":
		sectionNameHindi = "मुबारकपुर आं0"
	case "MUBARAKPUR HADSA":
		sectionNameHindi = "मुबारकपुर हरदास"
	case "MUBARAKPUR KHOSA":
		sectionNameHindi = "मुबारकपुर खोसा"
	case "MUBARAKPUR MADHU URF GADI":
		sectionNameHindi = "मुबारकपुर मधु उर्फ गढी"
	case "MUBARAKPUR MEERA AN.":
		sectionNameHindi = "मुबारकपुर मीरा आं०"
	case "MUBARAKPUR MEERA AN":
		sectionNameHindi = "मुबारकपुर मीरा आं०"
	case "MUBARAKPUR NAWADA":
		sectionNameHindi = "मुबारकपुर नवादा"
	case "MUBARAKPUR RATHE":
		sectionNameHindi = "मुबारकपुर राठे"
	case "MUBARAKPUR SAHARAN":
		sectionNameHindi = "मुबारकपुर सहारन"
	case "MUBARAKPUR TAPPA BHARTI B..A":
		sectionNameHindi = "मुबारकपुर टप्‍पा भारती बी ए"
	case "MUBARAKPUR URF GADAI":
		sectionNameHindi = "मुबारकपुर उर्फ गदाई"
	case "MUBARAKPUR-1":
		sectionNameHindi = "मुबारकपुर-1"
	case "MUBARAKPUR-2":
		sectionNameHindi = "मुबारकपुर-2"
	case "MUBARAKPUR":
		sectionNameHindi = "मुबारकपुर"
	case "MUBARKPUR TALAN":
		sectionNameHindi = "मुबारकपुर तालन"
	case "MUDALA":
		sectionNameHindi = "मुढाला"
	case "MUDHAL":
		sectionNameHindi = "मुढाल"
	case "MUDI":
		sectionNameHindi = "म्‍यूडी"
	case "MUDIA":
		sectionNameHindi = "मुडिया"
	case "MUDIYA KALAN":
		sectionNameHindi = "मुण्डिया कलां"
	case "MUDIYA KHURD":
		sectionNameHindi = "मुण्डिया खुर्द"
	case "MUDIYA MALOOKPUR AANSHIK":
		sectionNameHindi = "मुडि़या मलूकपुर आंशिक"
	case "MUDIYA MOHINUDDINPUR":
		sectionNameHindi = "मुडिया मुहीउददीनपुर"
	case "MUDIYA NEAR MANKARA":
		sectionNameHindi = "मुण्डिया नि0 मनकरा"
	case "MUDIYA RAJA":
		sectionNameHindi = "मुडिया राजा"
	case "MUDIYAJAIN":
		sectionNameHindi = "मुड़ियाजैन"
	case "MUFTI TOLA ANSHIK":
		sectionNameHindi = "मुफती टोला आंशिक"
	case "MUFTI TOLA":
		sectionNameHindi = "मुफती टोला"
	case "MUFTISARAY A":
		sectionNameHindi = "मुफतीसराय आ०"
	case "MUFTIYAN":
		sectionNameHindi = "मुफतीयान"
	case "MUGALPURA PRATHAM ANSHIK":
		sectionNameHindi = "मुगलपुरा प्रथम आंशिक"
	case "MUGALPURA SAKKA AVVAL ANSHIK":
		sectionNameHindi = "मुगलपुरा सक्‍का अव्‍वल आंशिक"
	case "MUGALPURA SAKKA DOYAM":
		sectionNameHindi = "मुगलपुरा सक्‍का दोयम"
	case "MUGALPURA SECOND":
		sectionNameHindi = "मुगलपुरा द्वितीय"
	case "MUGLAN PASCHIMI":
		sectionNameHindi = "मुगलान पश्चिमी"
	case "MUGLAN PURVI":
		sectionNameHindi = "मुगलान पूर्वी"
	case "MUGLUSHAH AN0":
		sectionNameHindi = "मुगलूशाह आं0"
	case "MUHAMMADPUR BAS KHADIYA":
		sectionNameHindi = "मोहम्‍मदपुर बस खडिया"
	case "MUHIUDDIN PUR":
		sectionNameHindi = "मुहीउद्दीन पुर"
	case "MUJAHIDPUR":
		sectionNameHindi = "मुजाहिदपुर,"
	case "MUJAHIPUR":
		sectionNameHindi = "मुजाहिदपुर"
	case "MUJHIYANA":
		sectionNameHindi = "मुझियाना"
	case "MUJJAFFARPUR TANDA":
		sectionNameHindi = "मुजफफरपुर टांडा"
	case "MUKAM":
		sectionNameHindi = "मुकाम"
	case "MUKANDOURRAHMAL":
		sectionNameHindi = "मुकंदपुरराजमल"
	case "MUKANDPUR RAMU":
		sectionNameHindi = "मुकन्‍दपुर रामू"
	case "MUKARANDPUR":
		sectionNameHindi = "मुकरन्‍दपुर"
	case "MUKARPUR AHIR.":
		sectionNameHindi = "मुकरपुर अहीर,"
	case "MUKARPUR GUJAR":
		sectionNameHindi = "मुकरपुर गुजर"
	case "MUKARPUR KHEMA":
		sectionNameHindi = "मुकरपुर खेमा"
	case "MUKARPUR":
		sectionNameHindi = "मुकरपुर"
	case "MUKARRABPUR ANSHIK":
		sectionNameHindi = "मुकर्रबपुर आंशिक"
	case "MUKARRABPUR MANDI BANS":
		sectionNameHindi = "मुकर्रबपुर मण्‍डी बॉस"
	case "MUKARRABPUR URF KUAKHEDA":
		sectionNameHindi = "मुकर्रबपुर उर्फ कुआं खेड़ा"
	case "MUKARRABPUR":
		sectionNameHindi = "मुकर्रबपुर"
	case "MUKEEMPUR BALU":
		sectionNameHindi = "मुकीपुर बालू"
	case "MUKEEMPUR DHARMASI":
		sectionNameHindi = "मुकीपुर धर्मसी"
	case "MUKEEMPUR DHARU":
		sectionNameHindi = "मुकीपुर धारू"
	case "MUKEEMPUR JAMAL URF INAMPUR AN":
		sectionNameHindi = "मुकीमपुर जमाल उर्फ इनामपुरा"
	case "MUKHTYAR PUR NAWADA":
		sectionNameHindi = "मुख्‍त्‍यारपुर नवादा"
	case "MUKIMPUR DUNIYA GAIR ABAD":
		sectionNameHindi = "मुकीमपुर दुनियागैर आबाद"
	case "MUKIMPUR PADARATH":
		sectionNameHindi = "मुकीमपुर पदारथ"
	case "MUKRABPUR":
		sectionNameHindi = "मुकरबपुर"
	case "MUKRANDANPUR MANAK":
		sectionNameHindi = "मकरन्‍दपुर मानक"
	case "MUKRAPURI AN.":
		sectionNameHindi = "मुकरपुरी आं0"
	case "MUKRMPUR":
		sectionNameHindi = "मुकरमपुर"
	case "MUKSUDPUR A.":
		sectionNameHindi = "मकसूदपुर आं0"
	case "MUKTESHWAR MAHADEV":
		sectionNameHindi = "मुक्‍टेश्‍वर महादेव"
	case "MUKUTPUR":
		sectionNameHindi = "मुकुटपुर"
	case "MULLAKHERA":
		sectionNameHindi = "मुल्‍लाखेडा"
	case "MUNDA KHEDI":
		sectionNameHindi = "मूंडा खेडी"
	case "MUNDALA MOHAMMAD JMAPUR":
		sectionNameHindi = "मुण्‍डाला मौ0 जमापुर"
	case "MUNDALA":
		sectionNameHindi = "मुढाला"
	case "MUNDHA PANDE ANSHIK":
		sectionNameHindi = "मूंढापाण्डे आंशिक"
	case "MUNDIA KHEDA":
		sectionNameHindi = "मुडिया खेडा"
	case "MUNDIYA BHIKAM":
		sectionNameHindi = "मुंडिया भीकम"
	case "MUNDIYA KALAN":
		sectionNameHindi = "मुण्डिया कलां"
	case "MUNDIYA RAJA":
		sectionNameHindi = "मुंडिया राजा"
	case "MUNDIYA RASOOLPUR":
		sectionNameHindi = "मुण्डिया रसूलपुर"
	case "MUNDIYA SEDNAGAR":
		sectionNameHindi = "मुण्डिया सैदनगर"
	case "MUNDIYA":
		sectionNameHindi = "मुण्डिया"
	case "MUNDIYAKHRDA":
		sectionNameHindi = "मुडीयाखेडा"
	case "MUNEEMPUR":
		sectionNameHindi = "मुनीमपुर"
	case "MUNEER GANJ AN0":
		sectionNameHindi = "मुनीरगंज आं0"
	case "MUNEERGANJ AN0":
		sectionNameHindi = "मुनीरगंज आं0"
	case "MUNEERGANJ":
		sectionNameHindi = "मुनीरगंज"
	case "MUNGARPUR":
		sectionNameHindi = "मुंगरपुर"
	case "MUNIMAPUR":
		sectionNameHindi = "मुनीमपुर"
	case "MUNIR NAGAR":
		sectionNameHindi = "मुनीरनगर"
	case "MUQAITPUR":
		sectionNameHindi = "मुकयतपुर"
	case "MURADBAKSHPUR":
		sectionNameHindi = "मुरादबख्‍शपुर"
	case "MURADNAGAR":
		sectionNameHindi = "मुरादनगर"
	case "MURADPUR GAIR AB.":
		sectionNameHindi = "मुरादपुर गैर आं0"
	case "MURADPUR":
		sectionNameHindi = "मुरादपुर"
	case "MURAHAT A.":
		sectionNameHindi = "मुराहट आ०"
	case "MURSAINA":
		sectionNameHindi = "मुरसैना"
	case "MURSHADPUR":
		sectionNameHindi = "मुरशदपुर"
	case "MURTAJAPUR BULAKI URF PEDA":
		sectionNameHindi = "मुर्तजापुर बुलाकी उर्फ पेदा"
	case "MURTJANAGAR URF GANGADHAR":
		sectionNameHindi = "मुर्तजानगर उर्फ गंगाधरपुर"
	case "MURTJAPUR URF PUNAPUR AN H.NO 1 TO 220":
		sectionNameHindi = "मुर्तजापुर उ र्फ पूरनपुर आं १ से २२० तक"
	case "MURTJAPUR URF PUNAPUR AN H.NO 221":
		sectionNameHindi = "मुर्तजापुरर्ऊ पूरनपुर म0 स0 १२१ से अंत तक"
	case "MURTUZABAD":
		sectionNameHindi = "मुर्तजाबाद"
	case "MUSARFGANJ":
		sectionNameHindi = "मुसर्रफगंज"
	case "MUSATAFAPUR TAIYAB":
		sectionNameHindi = "मुस्‍तफापुर तय्यब"
	case "MUSHARFABAD":
		sectionNameHindi = "मुशर्फबाद"
	case "MUSHTAFFABAD GARHI":
		sectionNameHindi = "मुस्‍तफाबाद गढी"
	case "MUSLIM KATERA":
		sectionNameHindi = "मुस्लिम कटेरा"
	case "MUSSAYPUR":
		sectionNameHindi = "मुस्‍सेपुर"
	case "MUSSEPUR":
		sectionNameHindi = "मुस्‍सेपुर"
	case "MUSTAFABAD DAUNKPURI":
		sectionNameHindi = "मुस्‍तफाबाद डौंकपुरी"
	case "MUSTAFABAD KHURD":
		sectionNameHindi = "मुस्‍तफाबाद खुर्द"
	case "MUSTAFABAD KLA":
		sectionNameHindi = "मुस्‍तफाबाद कला"
	case "MUSTAFABAD URF GHDANPURA":
		sectionNameHindi = "मुस्‍तफाबाद उर्फ गदनपुरा"
	case "MUSTAFABAD URF TAKLABAD":
		sectionNameHindi = "मुस्‍तफाबाद उर्फ टकलाबाद"
	case "MUSTAFABAD":
		sectionNameHindi = "मुस्‍तफाबाद"
	case "MUSTAFAPUR ANSU":
		sectionNameHindi = "मुस्‍तफापुर आंसू"
	case "MUSTAFAPUR DITIYA":
		sectionNameHindi = "मुस्‍तफापुर दितीय"
	case "MUSTAFAPUR GAIR ABAD":
		sectionNameHindi = "मुस्‍तफापुर गैर आबाद"
	case "MUSTAFAPURBHURAPURCHAKSAIDJAMAL":
		sectionNameHindi = "मुसतफापुर भूरापुर चकसैदजमल"
	case "MUSTAFFAPUR ABBAL":
		sectionNameHindi = "मुस्‍तफापुर अब्‍बल"
	case "MUSTFABAD":
		sectionNameHindi = "मुस्‍तफाबाद"
	case "MUSVI KHANPUR":
		sectionNameHindi = "मुस्‍वी खानपुर"
	case "MUTAAZ HUSAIN AN0":
		sectionNameHindi = "मुमताज हुसैन आ0"
	case "MUTHRAPUR URF KALUWALA":
		sectionNameHindi = "मुथरापुर ऊर्फ कालूवाला"
	case "MUTIYAPURA BAZARPATTI":
		sectionNameHindi = "मुतियापुरा वजरपटटी"
	case "MUTIYAPURA NIKAT BHOT":
		sectionNameHindi = "मुतियापुरा निकट भोट"
	case "MUZAFFARABAD":
		sectionNameHindi = "मुजफफराबाद"
	case "MUZAFFARPUR DEVIDAS":
		sectionNameHindi = "मुजफफरपुर देवीदास"
	case "MUZAFFARPUR KHADER":
		sectionNameHindi = "मुजफरपुर खादर"
	case "MUZAFFARPUR":
		sectionNameHindi = "मुजफफरपुर"
	case "MUZAFFERPUR KESHO":
		sectionNameHindi = "मुजफफरपुर केशो"
	case "MUZAHIDPUR":
		sectionNameHindi = "मुजाहिदपुर"
	case "NAANKAR RANI-1":
		sectionNameHindi = "नानकार रानी 1"
	case "NAANKAR RANI-2":
		sectionNameHindi = "नानकार रानी 2"
	case "NAANKAR RANI-3":
		sectionNameHindi = "नानकार रानी 3"
	case "NAANKAR RANI-4":
		sectionNameHindi = "नानकार रानी 4"
	case "NAANKAR":
		sectionNameHindi = "नानकार"
	case "NABAB NAGAR TAH":
		sectionNameHindi = "नबाबनगर टाह"
	case "NABAB NAGAR TANDA":
		sectionNameHindi = "नवाव नगर टाण्‍डा"
	case "NABABGANJ SHUMALI":
		sectionNameHindi = "नवावगंज शुमाली"
	case "NABABGUNJ":
		sectionNameHindi = "नवावगंज"
	case "NABADA":
		sectionNameHindi = "नबादा"
	case "NABBA NAGLA":
		sectionNameHindi = "नव्‍वा नगला"
	case "NABI GANJ PIPLI":
		sectionNameHindi = "नबीगंज पीपली"
	case "NABIGANJ JADID":
		sectionNameHindi = "नवीगंज जदीद"
	case "NABIGANJ":
		sectionNameHindi = "नवीगंज"
	case "NADARGANJ":
		sectionNameHindi = "नादरगंज"
	case "NADDAFAN":
		sectionNameHindi = "नददाफान"
	case "NADE HIND":
		sectionNameHindi = "नादे हिन्‍द"
	case "NADNA":
		sectionNameHindi = "नदना"
	case "NADNAU":
		sectionNameHindi = "नदनऊ"
	case "NAEEM GANJ":
		sectionNameHindi = "नईम गंज"
	case "NAGAL AN0":
		sectionNameHindi = "नांगल आं0"
	case "NAGALA HASHA":
		sectionNameHindi = "नगला हाशा"
	case "NAGALA SALAR":
		sectionNameHindi = "नगला सालार"
	case "NAGALAGUJAR":
		sectionNameHindi = "नगलागूजर"
	case "NAGALAKAMAL":
		sectionNameHindi = "नगलाकमाल"
	case "NAGALIYA KASAMGANJ":
		sectionNameHindi = "नगलिया कासमगंज"
	case "NAGALIYA MASHMULA":
		sectionNameHindi = "नगलिया मशमूला"
	case "NAGALIYA NARAYAN":
		sectionNameHindi = "नगलिया नारायन"
	case "NAGALIYA":
		sectionNameHindi = "नगलिया"
	case "NAGALIYAJAT":
		sectionNameHindi = "नगलियाजट"
	case "NAGALIYASHAHPUR":
		sectionNameHindi = "नगलियाशाहपुर"
	case "NAGARIYA - 2":
		sectionNameHindi = "नगरिया - 2"
	case "NAGARIYA KALAN -1":
		sectionNameHindi = "नगरिया कलां-1"
	case "NAGARIYA KALAN -2":
		sectionNameHindi = "नगरिया कलां-2"
	case "NAGARIYA-1":
		sectionNameHindi = "नगरिया-1"
	case "NAGLA BANSNAGLI":
		sectionNameHindi = "नगला बांसनगली"
	case "NAGLA GANESH":
		sectionNameHindi = "नगला गनेश"
	case "NAGLA HARDAS":
		sectionNameHindi = "नंगला हरदास"
	case "NAGLA NASSU":
		sectionNameHindi = "नगला नस्सू"
	case "NAGLA TAHAR":
		sectionNameHindi = "नगला ताहर"
	case "NAGLA UDAI-1":
		sectionNameHindi = "नगला उदई-1"
	case "NAGLA UDAI-2":
		sectionNameHindi = "नगला उदई-2"
	case "NAGLI BHAGWANT":
		sectionNameHindi = "नगली भगवन्‍त"
	case "NAGLI":
		sectionNameHindi = "नंगली"
	case "NAGLIYA AAQIL 1":
		sectionNameHindi = "नगलिया आकिल 1"
	case "NAGLIYA AAQIL-2":
		sectionNameHindi = "नगलिया आकिल- 2"
	case "NAGLIYA AAQIL-3":
		sectionNameHindi = "नगलिया आकिल- 3"
	case "NAGLIYA AAQIL-4":
		sectionNameHindi = "नगलिया आकिल- 4"
	case "NAGLIYA AAQIL-5":
		sectionNameHindi = "नगलिया आकिल- 5"
	case "NAGLIYA AAQIL-6":
		sectionNameHindi = "नगलिया आकिल- 6"
	case "NAGLIYA AAQIL-7":
		sectionNameHindi = "नगलिया आकिल- 7"
	case "NAGLIYA AAQIL-8":
		sectionNameHindi = "नगलिया आकिल- 8"
	case "NAGLIYA AAQIL-9":
		sectionNameHindi = "नगलिया आकिल- 9"
	case "NAGPHANI DAKSHINI ANSHIK":
		sectionNameHindi = "नागफनी दक्षिणी आंशिक"
	case "NAGPHANI UTTARI ANSHIK":
		sectionNameHindi = "नागफनी उत्‍तरी आंशिक"
	case "NAGRIYA KHURD":
		sectionNameHindi = "नगरिया खुर्द"
	case "NAHTAURA GAIRA.":
		sectionNameHindi = "नहटौरा गैरा0"
	case "NAI BASTI B-15":
		sectionNameHindi = "नई बस्‍ती बी 15"
	case "NAI BASTI B-23":
		sectionNameHindi = "नई बस्‍ती बी 23"
	case "NAI BASTI B-24":
		sectionNameHindi = "नई बस्‍ती बी 14"
	case "NAI SARAY":
		sectionNameHindi = "नई सराय"
	case "NAIKOFAL URF BILAI":
		sectionNameHindi = "नेकोफाल उर्फ बिलाई"
	case "NAIMGANJ":
		sectionNameHindi = "नईमगंज"
	case "NAIMPUR GOVARDHAN":
		sectionNameHindi = "नईमपुर गोर्वधन"
	case "NAINPURA AN.":
		sectionNameHindi = "नैनपुरा आं०"
	case "NAINU NANGAL":
		sectionNameHindi = "नैनू नांगल"
	case "NAIWALA":
		sectionNameHindi = "नाईवाला"
	case "NAIYAKHEDA":
		sectionNameHindi = "नईयाखेड़ा"
	case "NAIYAMATPUR BUCHA GAIRA.":
		sectionNameHindi = "नैयमतपुर बूचा गैरा0"
	case "NAIYMATPUR BOOCHA GAIRA0":
		sectionNameHindi = "नैयमतपुर बूचा गैरा0"
	case "NAJARPUR":
		sectionNameHindi = "नजरपुर"
	case "NAJIBPURSADAT":
		sectionNameHindi = "नजीबापुर सादात"
	case "NAJIMPUR":
		sectionNameHindi = "नजीमपुर"
	case "NAJRANA":
		sectionNameHindi = "नजराना"
	case "NAJRPUR ANSHIK":
		sectionNameHindi = "नाजरपुर आंशिक"
	case "NAJRPUR PURVI":
		sectionNameHindi = "नाजरपुर पूर्वी"
	case "NAKATPURI KHURD":
		sectionNameHindi = "नकटपुरी खुर्द"
	case "NAKHUNKA":
		sectionNameHindi = "नाखूनका"
	case "NAKIBPUR":
		sectionNameHindi = "नकीबपुर"
	case "NAKIMPUR RAMRAI GAIR ABAD":
		sectionNameHindi = "नकीमपुर रामराय गैर आबाद"
	case "NAKIPUR BAMNOLIGAIR ABAD":
		sectionNameHindi = "नकीपुर बमनोली गैर आबाद"
	case "NAKSANDA BAD":
		sectionNameHindi = "नकसन्‍दाबाद"
	case "NAKTAPURI KALA":
		sectionNameHindi = "नकटपुरी कलां"
	case "NALAPAR MAY KUCHA SAUDAGARAN-2":
		sectionNameHindi = "नालापार मय कुंचा सौदागरान-2"
	case "NALAPAR MAY KUCHASAUDAGARAN-1":
		sectionNameHindi = "नालापार मय कूचासौदागरान-1"
	case "NALAPAR":
		sectionNameHindi = "नालापार"
	case "NALBANDAN AN.H.NO 1-141/2":
		sectionNameHindi = "नालबंदान आं०म०सं०१से १४१/२तक"
	case "NALBANDAN AN.H.NO 142-END":
		sectionNameHindi = "नालबंदान आं०म०सं०१४२से अंत तक"
	case "NALPURA":
		sectionNameHindi = "नलपुरा"
	case "NAMAINI GADDI":
		sectionNameHindi = "नमैनी गद्दी"
	case "NAMAINI UDAIYA":
		sectionNameHindi = "नमैनी उदईया"
	case "NAMDARPUR CHANDAN":
		sectionNameHindi = "नामदार पुर चान्‍दन"
	case "NANAPUR":
		sectionNameHindi = "नानपुर"
	case "NAND COLONY":
		sectionNameHindi = "नन्‍द कालौनी"
	case "NANDAGANVA":
		sectionNameHindi = "नन्‍दगांव"
	case "NANDGAUN":
		sectionNameHindi = "नन्‍दगांव"
	case "NANDNOR":
		sectionNameHindi = "नान्‍दनौर"
	case "NANDPUR":
		sectionNameHindi = "नन्‍दपुर"
	case "NANGAL AN0":
		sectionNameHindi = "नांगल आं0"
	case "NANGAL JAT":
		sectionNameHindi = "नांगल जट"
	case "NANGALA BHAJJA":
		sectionNameHindi = "नंगला भज्‍जा"
	case "NANGALA BUDHVA":
		sectionNameHindi = "नंगला बुद्धवा"
	case "NANGALA GANNA":
		sectionNameHindi = "नंगला गन्‍ना"
	case "NANGALA GUNGA":
		sectionNameHindi = "नंगला गूंगा"
	case "NANGALA JAJAN":
		sectionNameHindi = "नंगला जाजन"
	case "NANGLA BANVEER":
		sectionNameHindi = "नंगला वनवीर"
	case "NANGLA ISLAM":
		sectionNameHindi = "नंगला इस्‍लाम"
	case "NANGLA JOGRAJ":
		sectionNameHindi = "नंगला जोगराज"
	case "NANGLA LADDHA":
		sectionNameHindi = "नंगला लद्दा"
	case "NANGLA NANSAI":
		sectionNameHindi = "नंगला नैनसी"
	case "NANGLA PITHORA":
		sectionNameHindi = "नंगला पिथौरा"
	case "NANGLA SARI":
		sectionNameHindi = "नंगला सारी"
	case "NANGLA UBBAN":
		sectionNameHindi = "नंगला उभ्‍भन"
	case "NANGLI JAJU":
		sectionNameHindi = "नंगली जाजू"
	case "NANGLI LADAN AN.":
		sectionNameHindi = "नंगली लाडन आ0"
	case "NANHEDA":
		sectionNameHindi = "नन्‍हेडा"
	case "NANHUVALA MU.":
		sectionNameHindi = "नन्हूवाला मु."
	case "NANJIWALA URF RAMJIWALA":
		sectionNameHindi = "नानकार उ र्फ रामजीवाला"
	case "NANKAR":
		sectionNameHindi = "नानकार"
	case "NANKAREN":
		sectionNameHindi = "नानकारैन"
	case "NANUPURA":
		sectionNameHindi = "ननुपुरा"
	case "NARAKHEDA":
		sectionNameHindi = "नरखेड़ा"
	case "NARAUDA":
		sectionNameHindi = "नरौदा"
	case "NARAYAN KHAIDI":
		sectionNameHindi = "नारायण खेडी"
	case "NARAYAN NAGALA":
		sectionNameHindi = "नरायन नगला"
	case "NARAYANPUR DEVA":
		sectionNameHindi = "नरायनपुर देवा"
	case "NARAYANPUR GAIR AB.AD":
		sectionNameHindi = "नारायणपुर गैर आं0"
	case "NARAYANPUR ICHCHA":
		sectionNameHindi = "नारायणपुर इच्‍छा"
	case "NARAYANPUR RAMJI":
		sectionNameHindi = "नरायनपुररामजी"
	case "NARAYANPUR RATAN AN0":
		sectionNameHindi = "नारायणपुर रतन आं0"
	case "NARAYANPUR":
		sectionNameHindi = "नारायणपुर"
	case "NARAYANWALA":
		sectionNameHindi = "नारायणवाला"
	case "NARAYNAPUR TEJU":
		sectionNameHindi = "नारायनपुर तेजू"
	case "NARENDRAPUR":
		sectionNameHindi = "नरेन्द्रपुर"
	case "NARGADI":
		sectionNameHindi = "नरगदी"
	case "NARKHEDA-1":
		sectionNameHindi = "नरखेडा-1"
	case "NARKHEDA-2":
		sectionNameHindi = "नरखेडा-2"
	case "NARKHEDA":
		sectionNameHindi = "नरखेडा"
	case "NARKHEDI":
		sectionNameHindi = "नरखेडी"
	case "NARKHERA":
		sectionNameHindi = "नरखेडा"
	case "NARKHERI":
		sectionNameHindi = "नरखेडी"
	case "NARPATNAGAR-1":
		sectionNameHindi = "नरपतनगर 1"
	case "NARPATNAGAR-2":
		sectionNameHindi = "नरपतनगर 2"
	case "NARPATNAGAR-3":
		sectionNameHindi = "नरपतनगर 3"
	case "NARPATNAGAR-4":
		sectionNameHindi = "नरपतनगर 4"
	case "NARPATNAGAR-5":
		sectionNameHindi = "नरपतनगर 5"
	case "NARPATNAGAR-6":
		sectionNameHindi = "नरपतनगर 6"
	case "NARPATNAGAR-7":
		sectionNameHindi = "नरपतनगर 7"
	case "NARPATNAGAR-8":
		sectionNameHindi = "नरपतनगर 8"
	case "NARSARA":
		sectionNameHindi = "नरसरा"
	case "NARSUA":
		sectionNameHindi = "नरसुआ"
	case "NARUKHEDA":
		sectionNameHindi = "नरुखेड़ा"
	case "NARULLAPUR BADLI GAIR ABAD":
		sectionNameHindi = "नरूल्‍लापुर बादली गैर आबाद"
	case "NASARPUR MANSURPUR":
		sectionNameHindi = "नासरपंर मंसूरपुर"
	case "NASEEBPUR":
		sectionNameHindi = "नसीबपुर"
	case "NASEERPUR BHIKKA":
		sectionNameHindi = "नसीरपुर भिक्का"
	case "NASEERPUR DHANNA":
		sectionNameHindi = "नसीरपुर धन्‍ना"
	case "NASEERPUR SHEK":
		sectionNameHindi = "नसीरपुर शेख"
	case "NASEERPUR":
		sectionNameHindi = "नसीरपुर"
	case "NASIMGANJ":
		sectionNameHindi = "नसीमगंज"
	case "NASIRAPUR BANAVARI":
		sectionNameHindi = "नसीरपुर बनवारी"
	case "NASIRI":
		sectionNameHindi = "नसीरी"
	case "NASIRPUR BUZURG":
		sectionNameHindi = "नसीरपुर बुजुर्ग"
	case "NASIRPUR KUNDARKI":
		sectionNameHindi = "नसीरपुर कुन्दरकी"
	case "NASIRPUR MANSUKH":
		sectionNameHindi = "नसीरपुर मनसु्ख"
	case "NASIRPUR MITHARI":
		sectionNameHindi = "नसीरपुर मिठारी"
	case "NASIRPUR NAIN SINGH":
		sectionNameHindi = "नसीरपुर नैन सिंह"
	case "NASRAT NAGAR":
		sectionNameHindi = "नसरतनगर"
	case "NASRATNAGAR":
		sectionNameHindi = "नसरतनगर"
	case "NASRPURAN NAJAMUDDIN":
		sectionNameHindi = "नसीरपुरनजमुददीन"
	case "NASRULLAPUR":
		sectionNameHindi = "नसरूल्‍लापुर"
	case "NATHADOI":
		sectionNameHindi = "नाथाडोई"
	case "NATHEYWALI":
		sectionNameHindi = "नत्‍थेवाली"
	case "NAUDHA":
		sectionNameHindi = "नौधा"
	case "NAUDHNA AN.":
		sectionNameHindi = "नौधना आं0"
	case "NAUGAVAN-1":
		sectionNameHindi = "नौगवां-1"
	case "NAUGAVAN-2":
		sectionNameHindi = "नौगवां-2"
	case "NAUGJA":
		sectionNameHindi = "नौगजा"
	case "NAUSAN ASHEKHAPUR AANSHIK":
		sectionNameHindi = "नौसना शेखपुर आंशिक"
	case "NAUSNA SHEKHPUR AANSHIK":
		sectionNameHindi = "नौसना शेखपुर आंशिक"
	case "NAUSNASYODARA":
		sectionNameHindi = "नौसनास्योडारा"
	case "NAVAB GANJ PARSAL":
		sectionNameHindi = "नवाबगंज पारसल"
	case "NAVAB NAGAR PADPURI":
		sectionNameHindi = "नबाबनगर पदपुरी"
	case "NAVADA KESHO":
		sectionNameHindi = "नवादा केशो"
	case "NAVADA SAIDAPUR JAMAL URF BHOOTPURI":
		sectionNameHindi = "नवादा सैदपुर जलाल उर्फ भूतपुरी"
	case "NAVADASHAHAPUR":
		sectionNameHindi = "नवादाशाहपुर"
	case "NAVADIYA":
		sectionNameHindi = "नवदिया"
	case "NAVAVPURA":
		sectionNameHindi = "नवावपुरा"
	case "NAVEEN NAGAR ANSHIK":
		sectionNameHindi = "नवीन नगर आंशिक"
	case "NAVEEN NAGAR":
		sectionNameHindi = "नवीन नगर"
	case "NAVI GANJ BATHUAKHEDA":
		sectionNameHindi = "नबीगंज नि0 बथुआखेडा"
	case "NAVIGUJ KADEEM":
		sectionNameHindi = "नवीगंज कदीम"
	case "NAVPUR":
		sectionNameHindi = "नावपुर"
	case "NAWAB GAUNJ JUNUBI":
		sectionNameHindi = "नवाब गंज जनूबी"
	case "NAWABGANJ":
		sectionNameHindi = "नवाबगंज"
	case "NAWABNAGAR NEAR PATWAI":
		sectionNameHindi = "नबावनगर निकट पटवाई"
	case "NAWABPURA ANSHIK":
		sectionNameHindi = "नवाबपुरा आंशिक"
	case "NAWABPURA URF SHAKUPURA":
		sectionNameHindi = "नवावपुरा उर्फ शेखुपुरा"
	case "nawabpura":
		sectionNameHindi = "नवाबपुरा"
	case "NAWABPURA":
		sectionNameHindi = "नवाबपुरा"
	case "NAWADA RAWANA URF TELIPURA":
		sectionNameHindi = "नवादा रवाना उर्फ तेलीपुरा"
	case "NAWADA SAHANPUR":
		sectionNameHindi = "नवादा साहनपुर"
	case "NAWADA":
		sectionNameHindi = "नवादा"
	case "NAWAJISPUR AHMAD":
		sectionNameHindi = "नवाजिसपुर अहमद"
	case "NAWALPUR":
		sectionNameHindi = "नवलपुर"
	case "NAYA GANAO":
		sectionNameHindi = "नया गांव"
	case "NAYA GANV NAZIBABAD":
		sectionNameHindi = "नयागांव नजीबाबाद"
	case "NAYA GANVA MAJRA":
		sectionNameHindi = "नया गांव मंजरा"
	case "NAYA GAON ANSHIK GANGAN":
		sectionNameHindi = "नया गांव आंशिक गागन"
	case "NAYA GAON ANSHIK MAU":
		sectionNameHindi = "नया गांव आंशिक मऊ"
	case "NAYA GAON ANSHIK":
		sectionNameHindi = "नया गॉव आंशिक"
	case "NAYA GAON":
		sectionNameHindi = "नया गॉव"
	case "NAYAK NANGLA":
		sectionNameHindi = "नायक नंगला"
	case "NAYAK SARAI":
		sectionNameHindi = "नायक सराय आं0"
	case "NAYAK SARAY":
		sectionNameHindi = "नायक सराय"
	case "NAYAMATPUR":
		sectionNameHindi = "नियामतपुर"
	case "NAYEE BASTI PURVI":
		sectionNameHindi = "नई बस्‍ती पूर्वी"
	case "NAYEE BASTI":
		sectionNameHindi = "नई बस्‍ती"
	case "NAZARPUR BILLOCH":
		sectionNameHindi = "नजरपुर बिल्‍लोच"
	case "NEHTORA":
		sectionNameHindi = "नहटौरा"
	case "NEJO SARAI AN0 H N0 1 TO 70":
		sectionNameHindi = "नेजो सराय आं0 म0स0 1 से 70 तक"
	case "NEJO SARAI AN0 H NO 71 TO END":
		sectionNameHindi = "नेजो सराय आं0 म0स0 71से अन्‍त तक"
	case "NEJO SARAI":
		sectionNameHindi = "नेजो सराय"
	case "NEJOWALI GAWDI":
		sectionNameHindi = "नेजो वाली गॉवडी"
	case "NEKPUR":
		sectionNameHindi = "नेकपुर"
	case "NEMATPUR":
		sectionNameHindi = "नेमतपुर"
	case "NEMATPYRISKUPURA":
		sectionNameHindi = "नेमतपुरइश्‍कीपुरा"
	case "NEMTULLA NAGAR ORF MITHANPUR":
		sectionNameHindi = "नेमतुल्‍ला नगर उर्फ मिठनपुर"
	case "NETA COLONY JAYANTIPUR":
		sectionNameHindi = "नेता कालेनी जयन्‍तीपुर"
	case "NEW CIVIL LINES":
		sectionNameHindi = "न्‍यू सिविल लाईन्‍स"
	case "NEW SAGARPUR, BISOULI -01":
		sectionNameHindi = "नया सागरपुर स्थित विसौली 01"
	case "NEW SAGARPUR, BISOULI -02":
		sectionNameHindi = "नया सागरपुर स्थित विसौली 0२"
	case "NEW SAGARPUR, BISOULI -03":
		sectionNameHindi = "नया सागरपुर स्थित विसौली 03"
	case "NEWVILLAGE NIKAT AKBARABAD":
		sectionNameHindi = "नयागांव निकट अकबरावाद"
	case "NEZE SARAY":
		sectionNameHindi = "नेजे सराय"
	case "NICHALPUR":
		sectionNameHindi = "नीचलपुर"
	case "NIDHI":
		sectionNameHindi = "निधि"
	case "NIJAMATPURA GANJ":
		sectionNameHindi = "निजामतपुरा गज"
	case "NIJAMPURDEVSI":
		sectionNameHindi = "निजामपुर देवसी"
	case "NIJAMPURPADARATH":
		sectionNameHindi = "निजामपुर पदारथ"
	case "NIKAMMASHAN AN.":
		sectionNameHindi = "निकम्‍माशाह आं0"
	case "NILOHA":
		sectionNameHindi = "निलोहा"
	case "NIMRI":
		sectionNameHindi = "नीमरी"
	case "NIN‍DOODU KHAS AN.":
		sectionNameHindi = "नीन्‍दडू खास आ0"
	case "NINDUDU KHAS AN.":
		sectionNameHindi = "नीन्‍दडू खास आं0"
	case "NIPANIYA-1":
		sectionNameHindi = "निपनियां-1"
	case "NIPANIYA-2":
		sectionNameHindi = "निपनियां-2"
	case "NIRMALAPUR":
		sectionNameHindi = "निर्मलपुर"
	case "NIRULLAPUR CHIMMA":
		sectionNameHindi = "नुरूल्‍लापुर चीमा"
	case "NISHVI":
		sectionNameHindi = "निस्‍वी"
	case "NISVI":
		sectionNameHindi = "निस्‍वी"
	case "NISWA":
		sectionNameHindi = "निस्‍वा"
	case "NIWADKHAS ANSHIK":
		sectionNameHindi = "निवाडखास आंशिक"
	case "NIYAMAT NAGAR":
		sectionNameHindi = "नियामत नगर"
	case "NIYAMATPUR":
		sectionNameHindi = "नियामतपुर"
	case "NIZAMPUR KHORA A":
		sectionNameHindi = "निजामपुर खोडा ए0"
	case "NIZAMPUR KHORA B A":
		sectionNameHindi = "निजामपुर खोडा बी0 ए0"
	case "NIZATPUR":
		sectionNameHindi = "निजातपुर"
	case "NOAABAD JAGEER":
		sectionNameHindi = "नौआबादजागीर"
	case "NOABAD DESH":
		sectionNameHindi = "नौआबाद देश"
	case "NOGRA":
		sectionNameHindi = "नौगरा"
	case "NOIABAD JANGAL":
		sectionNameHindi = "नौआबाद जंगल"
	case "NOMI MEDIUM":
		sectionNameHindi = "नौमी मध्‍य"
	case "NOMI SOUTH":
		sectionNameHindi = "नौमी द0"
	case "NOMI UTTARI":
		sectionNameHindi = "नौमी उत्‍तरी"
	case "NOORALAPUR HAFIZ":
		sectionNameHindi = "नरूल्‍लापुर हफीज"
	case "NOORAMPUR":
		sectionNameHindi = "नूरमपुर"
	case "NOORPUR CHHIBRI AN.":
		sectionNameHindi = "नूरपुर छीबरी आ0"
	case "NOORPUR DALO":
		sectionNameHindi = "नूरपुर डालू"
	case "NOORPUR DEHAT":
		sectionNameHindi = "नूरपुर देहात"
	case "NOORPUR DEHRA":
		sectionNameHindi = "नूरपुर देहरा"
	case "NOORPUR HATTI":
		sectionNameHindi = "नूरपुर हटटी"
	case "NOORPUR HAZOORPUR":
		sectionNameHindi = "नूरपुर हजूरपुर"
	case "NOORPUR KHAIDKI":
		sectionNameHindi = "नूरपुर खेडकी"
	case "NOORPURARAB":
		sectionNameHindi = "नूरपुर अरब"
	case "NORAHA":
		sectionNameHindi = "नौरहा"
	case "NORANGPUR":
		sectionNameHindi = "नौरगपुर"
	case "NOROJPUR GAIR ABAD":
		sectionNameHindi = "नोरोजपुर गैर आबाद"
	case "NRAYANPUR CHANGA":
		sectionNameHindi = "नरायनपुर छंगा"
	case "NRENDARPUR 01":
		sectionNameHindi = "नरेन्‍द्रपुर 01"
	case "NRENDARPUR 02":
		sectionNameHindi = "नरेन्‍द्रपुर 02"
	case "NURALLAPUR UDYACHAND":
		sectionNameHindi = "नुरूल्‍लापुर उदयचन्‍द"
	case "NURALPIURBHAGWANT":
		sectionNameHindi = "नूरअलीपुर भगवंत"
	case "NURANI MASJID JAYANTIPUR":
		sectionNameHindi = "नुरानी मस्जिद जयन्‍तीपुर"
	case "NURAPUR CHHIBRI AN.":
		sectionNameHindi = "नूरपुर छीवरी आ0"
	case "NURAPUR":
		sectionNameHindi = "नूरपुर"
	case "NURMOHDPUR":
		sectionNameHindi = "नूरमौहम्‍मदपुर"
	case "NURPUR":
		sectionNameHindi = "नूरपुर"
	case "NURUDDINPUR URF GANJ":
		sectionNameHindi = "नूरुद्दीनपुर उर्फ गंज"
	case "NURUDEENPUR":
		sectionNameHindi = "नुरूद्वीनपुर"
	case "NURULDEHARPUR URF MIRZAPUR":
		sectionNameHindi = "नुरूल दहरपुर ऊर्फ मिर्जापुर"
	case "NURULLA DAKSHIN AANSHIK":
		sectionNameHindi = "नुरुल्ला दक्षिण आंशिक"
	case "NURULLA MADHYA":
		sectionNameHindi = "नुरुल्ला मध्य"
	case "NURULLA PASH‍CHIMI":
		sectionNameHindi = "नुरुल्ला पश्‍चिमी"
	case "NURULLAPUR":
		sectionNameHindi = "नूरूल्‍लापुर"
	case "NYAMTABAD":
		sectionNameHindi = "न्‍यामताबाद"
	case "N‍YAMTPUR IK‍ROTIYA ANSHIK":
		sectionNameHindi = "न्‍यामतपुर इक्‍रोटिया आंशिक"
	case "NYU AVAS VIKAS":
		sectionNameHindi = "न्यू आवास विकास"
	case "OLIYAPUR":
		sectionNameHindi = "ओैलियापुर"
	case "OOTWAN":
		sectionNameHindi = "ऊंटवान"
	case "ORAINGPUR JOGIPURA URF NORANGPUR":
		sectionNameHindi = "औंरगपुर जोगीपुरा उर्फ नौरंगपुर"
	case "ORANGABAD":
		sectionNameHindi = "औरंगाबाद"
	case "ORANGNAGAR KHEDA":
		sectionNameHindi = "औरंगनगर खेडा"
	case "OSI 01":
		sectionNameHindi = "ओसी 01"
	case "OSI 02":
		sectionNameHindi = "ओसी 02"
	case "P.A.C. 24 BATALIAN":
		sectionNameHindi = "पी0ए0सी0 24 बटालियन"
	case "P.T.C. FIRST":
		sectionNameHindi = "पी0टी0सी0 प्रथम"
	case "P.T.C. SECOND":
		sectionNameHindi = "पी0टी0सी0 द्वितीय"
	case "P.V. PAHADPURKALA URF MALASJYA":
		sectionNameHindi = "पहाडपुर कला उर्फ मलेशिया"
	case "PACHOKRA KHANPUR":
		sectionNameHindi = "पचोकरा खानपुर"
	case "PACHPEDA KATGHAR ANSHIK":
		sectionNameHindi = "पचपेडा कटघर आंशिक"
	case "PACHPEDA KATGHAR":
		sectionNameHindi = "पचपेडा कटघर"
	case "PACHTOUR":
		sectionNameHindi = "पचतौर"
	case "PADAMPUR":
		sectionNameHindi = "पदमपुर"
	case "PADANLA":
		sectionNameHindi = "पाडंला"
	case "PADARATHPUR":
		sectionNameHindi = "पदारथपुर"
	case "PADHAN":
		sectionNameHindi = "पाधान"
	case "PADIYA NAGLA":
		sectionNameHindi = "पदिया नगला"
	case "PADLA AN":
		sectionNameHindi = "पाडला आं०"
	case "PADLA":
		sectionNameHindi = "पाडला"
	case "PADLI MANDU":
		sectionNameHindi = "पाडली मान्‍डू"
	case "PADLI":
		sectionNameHindi = "पाडली"
	case "PADLIBAJE":
		sectionNameHindi = "पाडलीबाजे"
	case "PADMAPUR":
		sectionNameHindi = "पदमपुर"
	case "PADPURI NEAR KALYANPUR":
		sectionNameHindi = "पदपुरी निकट कल्‍यानपुर"
	case "PADPURI SUAR":
		sectionNameHindi = "पदपुरी स्‍वार"
	case "PADPURI":
		sectionNameHindi = "पदपुरी"
	case "PAGAMBERPUR":
		sectionNameHindi = "पेगम्‍बरपुर"
	case "PAGAMPUR URF BAGIR GANG":
		sectionNameHindi = "पैगम्‍बरपुर उर्फ वजीरपुर गज"
	case "PAHADI DARAVAJA AN.":
		sectionNameHindi = "पहाडी दरवाजा आ0"
	case "PAHADI GET":
		sectionNameHindi = "पहाड़ी गेट"
	case "PAHADI":
		sectionNameHindi = "पहाडी"
	case "PAHADIDDARVAJA":
		sectionNameHindi = "पहाडी दरवाजा"
	case "PAHADMAU":
		sectionNameHindi = "पहाडमऊ"
	case "PAHADPUR BILASPUR":
		sectionNameHindi = "पहाडपुर बिलासपुर"
	case "PAHADPUR CHANDRASAIN":
		sectionNameHindi = "पहाडपुर चन्‍द्रसैन"
	case "PAHADPUR KHURD":
		sectionNameHindi = "पहाडपुर खुर्द"
	case "PAHADPUR":
		sectionNameHindi = "पहाडपुर"
	case "PAHARPUR BAST GAIR AB.":
		sectionNameHindi = "पहाडपुर बसत गैर आं0"
	case "PAHARPUR CHATAR":
		sectionNameHindi = "पहाडपुर चतर"
	case "PAHARPUR KAMALPUR":
		sectionNameHindi = "पहाडपुर कमालपुर"
	case "PAHARPUR":
		sectionNameHindi = "पहाडपुर"
	case "PAHULI":
		sectionNameHindi = "पाहुली"
	case "PAIBAGH":
		sectionNameHindi = "पाईबाग"
	case "PAIGAMBAR PUR":
		sectionNameHindi = "पैगम्‍बरपुर"
	case "PAIGAMBARPUR SUKHBASI LAL":
		sectionNameHindi = "पैगम्‍बरपुर सुखवासी लाल"
	case "PAIGAMBARPUR":
		sectionNameHindi = "पैगम्‍वरपुर"
	case "PAIGUPURA":
		sectionNameHindi = "पैगूपुरा"
	case "PAIJANIYA A. H.N. 81 TO END":
		sectionNameHindi = "पैंजनिया आ0 म0नं0 ८१ से अंत तक"
	case "PAIJANIYA A.":
		sectionNameHindi = "पैंजनिया आ0 म0नं0 १ से ८० तक"
	case "PAIMPUR":
		sectionNameHindi = "पैमपुर"
	case "PAINDA NAGAR":
		sectionNameHindi = "पाइन्‍दा नगर"
	case "PAINDA PUR":
		sectionNameHindi = "पाइंदा पुर"
	case "PAINDAPUR":
		sectionNameHindi = "पाइन्‍दापुर"
	case "PAIPATPURA MANDI SAMITI":
		sectionNameHindi = "पैपटपुरा मण्‍डी समिति"
	case "PAIPURA":
		sectionNameHindi = "पईपुरा"
	case "PAJABA":
		sectionNameHindi = "पजावा"
	case "PAJIYA":
		sectionNameHindi = "पजईया"
	case "PAJWA-1":
		sectionNameHindi = "पजाबा-1"
	case "PAJWA-2":
		sectionNameHindi = "पजाबा-2"
	case "PAKAHANPUR":
		sectionNameHindi = "पखनपुर"
	case "PAKBDA":
		sectionNameHindi = "पाकबडा"
	case "PAKKA BAGH ANSHIK":
		sectionNameHindi = "पक्‍का बाग आंशिक"
	case "PAKKABAG":
		sectionNameHindi = "पक्‍काबाग"
	case "PALANPUR":
		sectionNameHindi = "पालनपुर"
	case "PALDI":
		sectionNameHindi = "पलडी"
	case "PALI TEKCHAND":
		sectionNameHindi = "पाली टेकचन्‍द"
	case "PALIJAT":
		sectionNameHindi = "पालीजट"
	case "PALIYA":
		sectionNameHindi = "पलिया"
	case "PALKI EMAN":
		sectionNameHindi = "पालकी एमन"
	case "PALLUPURA GHOSI":
		sectionNameHindi = "पल्‍लूपुरा घोसी"
	case "PALNAPUR":
		sectionNameHindi = "पालनपुर"
	case "PALPUR":
		sectionNameHindi = "पालपुर"
	case "PALPURA":
		sectionNameHindi = "पलपुरा"
	case "PAMARGANG":
		sectionNameHindi = "पामरगंज"
	case "PANCHAYTI MANDIR UTTARI":
		sectionNameHindi = "पंचायती मन्दिर उत्‍तरी"
	case "PANCHAYTI MANDIR":
		sectionNameHindi = "पंचायती मन्दिर"
	case "PANDIT NAGLA ANSHIK":
		sectionNameHindi = "पंडित नगला आंशिक"
	case "PANDITAPUR":
		sectionNameHindi = "पंडितपुर"
	case "PANDIYA AANSHIK":
		sectionNameHindi = "पंडिया आंशिक"
	case "PANDIYA":
		sectionNameHindi = "पन्डिया"
	case "PANIYALA A.":
		sectionNameHindi = "पनियाला आ0"
	case "PANJABIYAN":
		sectionNameHindi = "पंजाबियान"
	case "PANJABNAGAR-1":
		sectionNameHindi = "पंजांबनगर-1"
	case "PANJABNAGAR-2":
		sectionNameHindi = "पंजांबनगर -2"
	case "PANJABNAGAR-3":
		sectionNameHindi = "पंजाबनगर-3"
	case "PANJAWNAGAR":
		sectionNameHindi = "पंजाबनगर"
	case "PANUVALA":
		sectionNameHindi = "पानूवाला"
	case "PANVADIYA KVATARS 2":
		sectionNameHindi = "पनवडिया क्वाटर्स 2"
	case "PANVADIYA KVATARS-1":
		sectionNameHindi = "पनवडिया क्वाटर्स-1"
	case "PANVADIYA":
		sectionNameHindi = "पनवडिया"
	case "PANWADIYA":
		sectionNameHindi = "पनवडिया"
	case "PAPAWAR KHURD":
		sectionNameHindi = "पपावर खुर्द"
	case "PAPSARA":
		sectionNameHindi = "पपसरा"
	case "PARAM-1":
		sectionNameHindi = "परम-1"
	case "PARAM-2":
		sectionNameHindi = "परम-2"
	case "PARAM-3":
		sectionNameHindi = "परम-3"
	case "PARAM-4":
		sectionNameHindi = "परम-4"
	case "PARASUPURA VAJE":
		sectionNameHindi = "परसुपुरा वाजे"
	case "PARBATPUR":
		sectionNameHindi = "पर्वतपुर"
	case "PARCHAI":
		sectionNameHindi = "परचई"
	case "PAREWA":
		sectionNameHindi = "परेवा"
	case "PARMANANDPUR GAWDI":
		sectionNameHindi = "परमानन्‍दपुरगॉवडी"
	case "PAROUTA 01":
		sectionNameHindi = "परौता 01"
	case "PARSHUPURA":
		sectionNameHindi = "परशूपुरा"
	case "PARSIPURA":
		sectionNameHindi = "परसीपुरा"
	case "PARTAPUR HARDASPUR":
		sectionNameHindi = "परतापुर हरदासपुर"
	case "PARWATPUR DAMBU":
		sectionNameHindi = "पर्वतपुर दम्‍बू"
	case "PARWATPUR MAKHDUMPUR AN0":
		sectionNameHindi = "पर्वतपुर मखदूमपुर आं0"
	case "PASBAN":
		sectionNameHindi = "पासबान"
	case "PASHUPURA":
		sectionNameHindi = "पशुपुरा"
	case "PASIYAPURA JANOOBI":
		sectionNameHindi = "पसियापुरा जनूबी"
	case "PASIYAPURA PADARATH":
		sectionNameHindi = "पसियापुरा पदारथ"
	case "PASIYAPURA SHUMALI":
		sectionNameHindi = "पसियापुरा शुमाली"
	case "PASIYAPURA":
		sectionNameHindi = "पसियापुरा"
	case "PASUPURA":
		sectionNameHindi = "पशुपुरा"
	case "PATARIYA":
		sectionNameHindi = "पटरिया"
	case "PATAVARIYAN AN.":
		sectionNameHindi = "पटवारियान आं0"
	case "PATEGANJ":
		sectionNameHindi = "पटेगंज"
	case "PATEL NAGAR BHOGPUR MITHONI ANSHIK":
		sectionNameHindi = "पटेल नगर भोगपुर मिठोनी आंशिक"
	case "PATERI":
		sectionNameHindi = "पटेरी"
	case "PATEYPARA A.":
		sectionNameHindi = "पतियापाडा आ०"
	case "PATHANAN":
		sectionNameHindi = "पठानान"
	case "PATHANGI":
		sectionNameHindi = "पाठंगी"
	case "PATHANPURA AN0":
		sectionNameHindi = "पठानपुरा आं0"
	case "PATHWARI":
		sectionNameHindi = "पथवारी"
	case "PATIYA":
		sectionNameHindi = "पटिया"
	case "PATLE V MASJID KHATOTI":
		sectionNameHindi = "पटले व मस्जिद खतोती"
	case "PATNA":
		sectionNameHindi = "पटना"
	case "PATPADA KESHO":
		sectionNameHindi = "पटपडा केशो"
	case "PATTHAR KHEDA":
		sectionNameHindi = "पत्‍थर खेडा"
	case "PATTI ASHOKPUR":
		sectionNameHindi = "पटटी अशोकपुर"
	case "PATTI BASANTPUR":
		sectionNameHindi = "पटटी बसन्‍तपुर"
	case "PATTI FAZILABAD":
		sectionNameHindi = "पटटी फाजिलाबाद"
	case "PATTI KALAN-2":
		sectionNameHindi = "पटटीकला 2"
	case "PATTI KALYANPUR":
		sectionNameHindi = "पटटी कल्‍यानपुर"
	case "PATTI MODHA":
		sectionNameHindi = "पटटी मौढा"
	case "PATTI TOLA":
		sectionNameHindi = "पटटी टोला"
	case "PATTIBALA DASHNI":
		sectionNameHindi = "पटटीवाला (दक्षिणी)"
	case "PATTIBALA PASHMI":
		sectionNameHindi = "पटटीवाला (पश्चिमी)"
	case "PATTIBALA UTTARI":
		sectionNameHindi = "पटटीवाला (उत्‍तरी)"
	case "PATTIKLA-1":
		sectionNameHindi = "पटटीकला 1"
	case "PATWAI -1":
		sectionNameHindi = "पटवाई -1"
	case "PATWAI -2":
		sectionNameHindi = "पटवाई -२"
	case "PATWAI -3":
		sectionNameHindi = "पटवाई -3"
	case "PATWAI 04":
		sectionNameHindi = "पटवाई 04"
	case "PATWARIYAN":
		sectionNameHindi = "पटवारियान"
	case "PATYAPARA A.":
		sectionNameHindi = "पतियापाडा आ०"
	case "PAVARTAPUR MAHANAND":
		sectionNameHindi = "पवर्तपुर महानंद"
	case "PAVTI A.":
		sectionNameHindi = "पावटी आ०"
	case "PAWAR HOUSE MAJHOLA":
		sectionNameHindi = "पावर हाऊस मझौला"
	case "PAYBAND KHERI":
		sectionNameHindi = "पेबन्‍द खेडी"
	case "PEDI":
		sectionNameHindi = "पेदी"
	case "PEEPALSANA ANSHIK":
		sectionNameHindi = "पीपलसाना आंशिक"
	case "PEEPALSANA SHAHI MASJID":
		sectionNameHindi = "पीपलसाना शाही मस्जिद"
	case "PEEPJAN":
		sectionNameHindi = "पीपजान"
	case "PEER SHAHEED KALA":
		sectionNameHindi = "पीर शहीद काला"
	case "PEERGAIB BAJIGRAN":
		sectionNameHindi = "पीरगैब बाजीग्रान"
	case "PEERGAIB":
		sectionNameHindi = "पीरगैब"
	case "PEERZADA DAKSHINI ANSHIK":
		sectionNameHindi = "पीरजादा दक्षिणी आंशिक"
	case "PEERZADA MADHYA ANSHIK":
		sectionNameHindi = "पीरजादा मध्‍य आंशिक"
	case "PEERZADA MADHYA":
		sectionNameHindi = "पीरजादा मध्‍य"
	case "PEETAL BASTI ANSHIK":
		sectionNameHindi = "पीतल बस्‍ती आंशिक"
	case "PEETANONDHA":
		sectionNameHindi = "पित्‍तन औन्‍धा"
	case "PEETPUR":
		sectionNameHindi = "पीतपुर"
	case "PEGA":
		sectionNameHindi = "पैगा"
	case "PEGAMWARPUR":
		sectionNameHindi = "पैगम्‍बरपुर"
	case "PELI VISHNOI":
		sectionNameHindi = "पेली विश्‍नोई"
	case "PEPALSANA":
		sectionNameHindi = "पीपलसाना"
	case "PERZADAGAN":
		sectionNameHindi = "पीरजादगान"
	case "PESH KOTAVALI":
		sectionNameHindi = "पेश कोतवाली"
	case "PHEENA A.":
		sectionNameHindi = "फीना आ०"
	case "PHOOLBAG":
		sectionNameHindi = "फूलबाग"
	case "PHOOLPUR MITHANPUR":
		sectionNameHindi = "फूलपुर मिठनपुर"
	case "PILA TALAB":
		sectionNameHindi = "पीला तालाब"
	case "PILAKPUR SHYORAM":
		sectionNameHindi = "पीलकपुर श्‍योराम"
	case "PILKAPUR GUMANI":
		sectionNameHindi = "पीलकपुर गुमानी"
	case "PILLANA":
		sectionNameHindi = "पिलाना"
	case "PIPAL GANV":
		sectionNameHindi = "पीपल गांव"
	case "PIPAL TOLA":
		sectionNameHindi = "पीपल टोला"
	case "PIPALGAUN-1":
		sectionNameHindi = "पीपलगावं-1"
	case "PIPALGAUN-2":
		sectionNameHindi = "पीपलगावं-2"
	case "PIPALICHAK":
		sectionNameHindi = "पीपलीचक"
	case "PIPALIUMRAPUR":
		sectionNameHindi = "पीपलीउमरपुर"
	case "PIPALIYA AHALA":
		sectionNameHindi = "पिपलिया अहला"
	case "PIPALIYA BHOJ":
		sectionNameHindi = "पिपलिया भोज"
	case "PIPALIYA GOPAL-2":
		sectionNameHindi = "पिपलिया गोपाल-2"
	case "PIPALIYA GOPAL":
		sectionNameHindi = "पिपलिया गोपाल-1"
	case "PIPALIYA JAT":
		sectionNameHindi = "पिपलियाजट"
	case "PIPALIYA MEHTO":
		sectionNameHindi = "पिपलिया महतो"
	case "PIPALIYA MISRA":
		sectionNameHindi = "पिपलिया मिश्र"
	case "PIPALIYA NAV":
		sectionNameHindi = "पिपलिया नाव"
	case "PIPALIYA RAIJADA":
		sectionNameHindi = "पिपलिया रायजादा"
	case "PIPALIYA VIJAYNAGAR":
		sectionNameHindi = "पिपलिया विजयनगर"
	case "PIPALSANA":
		sectionNameHindi = "पीपलसाना"
	case "PIPALWALA":
		sectionNameHindi = "पीपलवाला"
	case "PIPLA JAGIR A.":
		sectionNameHindi = "पीपला जागीर आ0"
	case "PIPLA SHIV NAGAR-2":
		sectionNameHindi = "पिपला शिवनगर-2"
	case "PIPLA SHIVNAGAR-1":
		sectionNameHindi = "पिपला शिवनगर-1"
	case "PIPLATOLA NIKAT GUJRATOLA 2":
		sectionNameHindi = "पीपलटोला निकट गूजरटोला- 2"
	case "PIPLATOLA NIKAT GUJRATOLA-1":
		sectionNameHindi = "पीपलटोला निकट गूजरटोला-1"
	case "PIPLI AHIR MU.":
		sectionNameHindi = "पीपली अहीर मु."
	case "PIPLI GHANSYAM":
		sectionNameHindi = "पीपली घनश्‍याम"
	case "PIPLI NAYAK-1":
		sectionNameHindi = "पीपली नायक-1"
	case "PIPLI NAYAK-2":
		sectionNameHindi = "पीपली नायक-2"
	case "PIPLI NAYAK-3":
		sectionNameHindi = "पीपली नायक-3"
	case "pipli":
		sectionNameHindi = "पीपली"
	case "PIPLI":
		sectionNameHindi = "पीपली"
	case "PIPLIJAT A":
		sectionNameHindi = "पीपली जट आ०"
	case "PIPLIJAT":
		sectionNameHindi = "पीपली जट आ०"
	case "PIR KA BAJAR AN.":
		sectionNameHindi = "पीर का बजार आं0"
	case "PITHANHERI JHOJHA KALA":
		sectionNameHindi = "पित्‍तनहेडी झोझा कला"
	case "PITHANHERIJHOJHA KHURD":
		sectionNameHindi = "पित्‍तनहेडी झोझा खुर्द"
	case "PITTANHERI JIYA AN.H.NO 1-113":
		sectionNameHindi = "पित्‍तनहेडी जिया आं०म०सं०१से ११३ तक"
	case "PITTHANHERI JIYA AN.H.NO 114-END":
		sectionNameHindi = "पित्‍तनहेडी जिया आं०म०सं०११४ से अंत तक"
	case "PITTHAPUR A.":
		sectionNameHindi = "पित्‍थापुर आ0"
	case "PITTHUPURA":
		sectionNameHindi = "पित्‍थुपुरा"
	case "POHAPPUR":
		sectionNameHindi = "पोहपपुर"
	case "POLICE CHOWKI KE PICHHE LAJPAT NAGAR":
		sectionNameHindi = "पुलिस चौकी के पीछे लाजपत नगर"
	case "POLICE LINE":
		sectionNameHindi = "पुलिस लाइन"
	case "POTA":
		sectionNameHindi = "पौटा"
	case "POTI":
		sectionNameHindi = "पौटी"
	case "PRAKASH ENCLAVE":
		sectionNameHindi = "प्रकाश एनक्‍लेब"
	case "PRAKASH NAGAR ANSHIK":
		sectionNameHindi = "प्रकाश नगर आंशिक"
	case "PRAKASH NAGAR":
		sectionNameHindi = "प्रकाश नगर"
	case "PRANPUR":
		sectionNameHindi = "प्रानपुर"
	case "PRATAP URF KHERPUR":
		sectionNameHindi = "प्रताप उर्फ खेरपुर"
	case "PRATAPPUR":
		sectionNameHindi = "प्रतापपुर"
	case "PRATHBIPUR":
		sectionNameHindi = "पिरथीपुर"
	case "PRATHVI GANJ DAKSHNI":
		sectionNameHindi = "पृथ्‍वीगंज (दक्षिणी)"
	case "PRATHVI GANJ PURBI":
		sectionNameHindi = "पृथ्‍वीगंज (पूर्वी)"
	case "PRATHVIPUR BANWARI":
		sectionNameHindi = "पृथ्‍वीपुर बनवारी"
	case "PRATHVIPUR":
		sectionNameHindi = "पृथ्‍वीपुर"
	case "PREAMPUR":
		sectionNameHindi = "प्रेमपुर"
	case "PREET VIHAR ANSHIK MAJHOLA":
		sectionNameHindi = "प्रीत विहार आंशिक मझोला"
	case "PREM NAGAR ANSHIK KANTH ROAD":
		sectionNameHindi = "प्रेम नगर आंशिक कांठ रोड"
	case "PREM NAGAR ANSHIK":
		sectionNameHindi = "प्रेम नगर आंशिक"
	case "PREM NAGAR LINEPAR ANSHIK":
		sectionNameHindi = "प्रेम नगर लाइनपार आंशिक"
	case "PREM NAGAR":
		sectionNameHindi = "प्रेम नगर"
	case "PREMNAGAR":
		sectionNameHindi = "प्रेमनगर"
	case "PREMPUR":
		sectionNameHindi = "प्रेमपुर"
	case "PREMPURI":
		sectionNameHindi = "प्रेमपुरी"
	case "PRHTHVIPUR":
		sectionNameHindi = "प्रथ्‍वीपुर"
	case "PRITHABI GANJ PASHMI":
		sectionNameHindi = "पृथ्‍वीगंज (पश्‍चिमी)"
	case "PRITHVI NAGAR":
		sectionNameHindi = "पृथ्‍वीनगर"
	case "PRITHVIPUR & CHIDIYAKHEDA":
		sectionNameHindi = "पृथ्‍वीपुर ऊर्फ चिडियाखेडा"
	case "PRITHVIPUR GABDI":
		sectionNameHindi = "पृथ्वीपुर गाबड़ी"
	case "PROTA -2":
		sectionNameHindi = "परौता -2"
	case "PSIYAPURA":
		sectionNameHindi = "पसियापुरा"
	case "PUJERI GALI":
		sectionNameHindi = "पुजेरी गली"
	case "PUKHAY WALA":
		sectionNameHindi = "पुक्‍खेवाला"
	case "PUL PUKHTA":
		sectionNameHindi = "पुल पुख्ता"
	case "PULIS LAIN":
		sectionNameHindi = "पुलिस लाइन"
	case "PUNDARI KHURD":
		sectionNameHindi = "पूण्‍डरी खुर्द"
	case "PUNJABIYAN":
		sectionNameHindi = "पंजाबियान"
	case "PURANPUR GARHI A":
		sectionNameHindi = "पूरनपुर गढ़ी ए"
	case "PURANPUR GARHI AN0":
		sectionNameHindi = "पूरनपुर गढ़ी आं0"
	case "PURANPUR NANGLA":
		sectionNameHindi = "पूरनपुर नंगला"
	case "PURANPUR NAROTTAM":
		sectionNameHindi = "पुरनपुर नरोत्‍तम"
	case "PURANPUR SAYANA":
		sectionNameHindi = "पूरनपुर स्‍याना"
	case "PURENA ABDUL REHMANPUR A.":
		sectionNameHindi = "पुरैना अब्‍दुल रहमानपुर आ0"
	case "PURENA":
		sectionNameHindi = "पुरैना"
	case "PURENI DURWESHPUR":
		sectionNameHindi = "पुरैनी दुर्वेशपुर"
	case "PURENIYA JADID":
		sectionNameHindi = "पुरैनियां जदीद"
	case "PURENIYA KALAN":
		sectionNameHindi = "पुरैनियां कलां"
	case "PURENIYA KHURD":
		sectionNameHindi = "पुरैनियां खुर्द"
	case "PURNAPUR":
		sectionNameHindi = "पुरनापुर"
	case "PURSHOTAMPUR URF THEY PUR":
		sectionNameHindi = "पुरषोत्‍त्‍मपुर ऊर्फ थेपुर"
	case "PURSHOTTAMPUR BHAGWAN":
		sectionNameHindi = "पुरशोत्‍तमपुर भगवान"
	case "PURUSHAOTTAMPUR":
		sectionNameHindi = "पुरूषोत्‍तमपुर"
	case "PURUSHOTTAMPUR":
		sectionNameHindi = "पुरूषोत्‍तमपुर"
	case "PUSWADA-1":
		sectionNameHindi = "पुसवाडा 1"
	case "PUSWADA-2":
		sectionNameHindi = "पुसवाडा 2"
	case "PUTLIGHAR LINEPAR ANSHIK":
		sectionNameHindi = "पुतलीघर लाइनपार आंशिक"
	case "PUTTHA 1":
		sectionNameHindi = "पुटठा 1"
	case "PUTTHA 2":
		sectionNameHindi = "पुटठा 2"
	case "PUTTHA":
		sectionNameHindi = "पुटठा"
	case "PUWENA":
		sectionNameHindi = "पूवैना"
	case "QAZIPARA JANUBI B":
		sectionNameHindi = "काजीपाडा जनूबी"
	case "QAZIPARA JANUBI":
		sectionNameHindi = "काजीपाडा जनूबी"
	case "QAZIPARA SHUMALI":
		sectionNameHindi = "काजीपाडा शुमाली"
	case "QAZIWALA":
		sectionNameHindi = "काजीवाला"
	case "QAZIYAN":
		sectionNameHindi = "काजीयान"
	case "QILA 1 TO 80":
		sectionNameHindi = "किला 1 से 80 तक"
	case "QILA H NO 81 TO END":
		sectionNameHindi = "किला म0 स0 81 से अन्‍त तक"
	case "QUILA":
		sectionNameHindi = "किला"
	case "QURESHIYAN AN0":
		sectionNameHindi = "कुरैशियान आं0"
	case "RABNNA":
		sectionNameHindi = "रवन्‍ना"
	case "RADHUVALA":
		sectionNameHindi = "राधूवाला"
	case "RAEPUR BIJJU":
		sectionNameHindi = "रायपुर बिज्‍जू"
	case "RAFATPUR FATEHABAD":
		sectionNameHindi = "रफतपुर फतेहाबाद"
	case "RAFATPUR":
		sectionNameHindi = "रफतपुर"
	case "RAFAYATPUR HULLAS":
		sectionNameHindi = "रफैतपुर हुसलास"
	case "RAFIKPUR":
		sectionNameHindi = "रफीकपुर"
	case "RAFILUL NAGAR URF RAWALI":
		sectionNameHindi = "रफीउलनगर उर्फ रावली"
	case "RAFIPUR MAJARA":
		sectionNameHindi = "रफीपुर मजरा"
	case "RAFIPUR MOHAN":
		sectionNameHindi = "रफीपुर मोहन"
	case "RAFIPUR":
		sectionNameHindi = "रफीपुर"
	case "RAFIULNAGAR URF RAWALI":
		sectionNameHindi = "रफीउलनगर उर्फ रावली"
	case "RAFTPUR":
		sectionNameHindi = "रफतपुर"
	case "RAGHAVRAMPUR SARNGPUR":
		sectionNameHindi = "राघवरामपुर सारगपुर"
	case "RAGHDAN":
		sectionNameHindi = "राघडान"
	case "RAGHORAMPUR":
		sectionNameHindi = "राधोरामपुर"
	case "RAGHUNATHAPUR":
		sectionNameHindi = "रघुनाथपुर"
	case "RAGHUNATHPUR":
		sectionNameHindi = "रघूनाथपुर"
	case "RAHATPUR KHURD AN0":
		sectionNameHindi = "राहतपुर खुर्द आं0"
	case "RAHPANPUR":
		sectionNameHindi = "रहपनपुर"
	case "RAHTA BILLOCH":
		sectionNameHindi = "रहटा बिल्‍लौच"
	case "RAHU NANGLI A.":
		sectionNameHindi = "राहू नंगली आ0"
	case "RAHUKHEDI KAURA":
		sectionNameHindi = "राहूखेडी कौरा"
	case "RAI BHAN WALA":
		sectionNameHindi = "रायभानवाला"
	case "RAIHATAMAFI":
		sectionNameHindi = "रैहटामाफी"
	case "RAIKANAGALA":
		sectionNameHindi = "रैंका नगला"
	case "RAILWAY COLONY AN0":
		sectionNameHindi = "रेलवे कालोनी आं0"
	case "RAILWAY COLONY ANO":
		sectionNameHindi = "रेलवे कालोनी आ0"
	case "RAILWAY COLONY RAHUKHERI KORA":
		sectionNameHindi = "रेलवे कालोनी राहुखेड़ी कोरा"
	case "RAILWAY HARTHALA COLONY":
		sectionNameHindi = "रेलवे हरथला कालौनी"
	case "RAILWAY SATTLEMENT NO. A. SECTOR-1":
		sectionNameHindi = "रेलवे सैटिलमेन्‍ट नो0ए0 सैक्‍टर 1"
	case "RAILWAY SATTLEMENT SECTOR-2":
		sectionNameHindi = "रेलवे सैटिलमेन्‍ट सैक्‍टर 2"
	case "RAILWAY SATTLEMENT SECTOR-4 ANSHIK":
		sectionNameHindi = "रेलवे सैटिलमेन्‍ट सैक्‍टर 4 आंशिक"
	case "RAILWAY SATTLEMENT SECTOR-5 ANSHIK":
		sectionNameHindi = "रेलवे सैटिलमेन्‍ट सैक्‍टर 5 आंशिक"
	case "RAILWAY SATTLEMENT":
		sectionNameHindi = "रेलवे सेटिलमेन्‍ट"
	case "RAILWAY SUTILMENTE SECTOR-7 LINEPAR":
		sectionNameHindi = "रेलवे सैटिलमेन्‍ट सैक्‍टर 7 लाइनपार"
	case "RAILWAY THIRD GATE":
		sectionNameHindi = "रेलवे थर्ड गेट"
	case "RAIPUR BARYSAL":
		sectionNameHindi = "रायपुर बेरीसाल"
	case "RAIPUR BERYSAL":
		sectionNameHindi = "रायपुर बेरीसाल"
	case "RAIPUR GAIR ABAD":
		sectionNameHindi = "रायपुर गेर आबाद"
	case "RAIPUR KHADER":
		sectionNameHindi = "रायपुर खादर"
	case "RAIPUR KHAS AN0":
		sectionNameHindi = "रायपुर खास आं0"
	case "RAIPUR MOJAMPUR NARAYAN":
		sectionNameHindi = "रायपुर मौजमपुर नारायण"
	case "RAIPUR PIYAGI":
		sectionNameHindi = "रायपुर पिरगी"
	case "RAIPUR SADAT AN H.NO 201":
		sectionNameHindi = "रायपुर सादात म0स0 आ0 २०१ तक"
	case "RAIPUR SADAT AN":
		sectionNameHindi = "रायपुर सादात आ0"
	case "RAIPUR SADAT H.NO1 TO 444 END":
		sectionNameHindi = "रायपुर सादात म0स0 1 से 444 अंत तक"
	case "RAIPUR SAMDA":
		sectionNameHindi = "रायपुर समदा"
	case "RAIPUR SUMALKERI":
		sectionNameHindi = "रायपुर सुमालखेडी"
	case "RAIPUR-01":
		sectionNameHindi = "रायपुर -01"
	case "RAIPUR-02":
		sectionNameHindi = "रायपुर -0२"
	case "RAIPUR-1":
		sectionNameHindi = "रायपुर -1"
	case "RAIPUR-2":
		sectionNameHindi = "रायपुर -2"
	case "RAIPUR":
		sectionNameHindi = "रायपुर"
	case "RAIPURI":
		sectionNameHindi = "रायपुरी"
	case "RAIPURSADAT AN0 445 TO END":
		sectionNameHindi = "रायपुर सादात आ0 म0स0 445 अंत तक"
	case "RAIPURSADAT H.NO 1 TO 200":
		sectionNameHindi = "रायपुर सादात आ0 म0 स0 १ से २०० तक"
	case "RAISAN D. 1":
		sectionNameHindi = "रईसान द0 1"
	case "RAISAN D. 2":
		sectionNameHindi = "रईसान द0 2"
	case "RAISAN D. 3":
		sectionNameHindi = "रईसान द0 3"
	case "RAISAN D. 4":
		sectionNameHindi = "रईसान द0 4"
	case "RAISAN D. 5":
		sectionNameHindi = "रईसान द0 5"
	case "RAISAN D. 6":
		sectionNameHindi = "रईसान द0 6"
	case "RAISAN D. 7":
		sectionNameHindi = "रईसान द0 7"
	case "RAISAN D. 8":
		sectionNameHindi = "रईसान द0 8"
	case "RAISAN U. 1,2":
		sectionNameHindi = "रईसान उ0 1,2"
	case "RAISAN U. 2":
		sectionNameHindi = "रईसान उ0 2"
	case "RAJA RAMPUR":
		sectionNameHindi = "राजा रामपुर"
	case "RAJADARA NAVAB GET V MASJID PILU-1":
		sectionNameHindi = "राजदारा नवाब गेट व मस्जिद पीलू-1"
	case "RAJADARA NAVAB GET V MASJID PILU-2":
		sectionNameHindi = "राजदारा नवाब गेट व मस्जिद पीलू-2"
	case "RAJADARA NAVAB GET V MASJID PILU-3":
		sectionNameHindi = "राजदारा नवाब गेट व मस्जिद पीलू-3"
	case "RAJADARA NAVAB GET V MASJID PILU-4":
		sectionNameHindi = "राजदारा नवाब गेट व मस्जिद पीलू-4"
	case "RAJADARA NAVAB GET V MASJID PILU-5":
		sectionNameHindi = "राजदारा नवाब गेट व मस्जिद पीलू-5"
	case "RAJAPUR KALA":
		sectionNameHindi = "राजपुर कला"
	case "RAJARAM PUR PRATAP":
		sectionNameHindi = "राजारामपुर प्रताप"
	case "RAJARAMPUR FAZIL":
		sectionNameHindi = "राजारामपुर फाजिल"
	case "RAJARAMPUR KHADAR":
		sectionNameHindi = "राजारामपुर खादर"
	case "RAJARAMPUR TULSI":
		sectionNameHindi = "राजारामपुर तुलसी"
	case "RAJDEV NANGLI":
		sectionNameHindi = "राजदेव नंगली"
	case "RAJJAD":
		sectionNameHindi = "रज्जड़"
	case "RAJO GALI ANSHIK":
		sectionNameHindi = "राजो गली आंशिक"
	case "RAJODA KANNAR DEV":
		sectionNameHindi = "रजोड़ा कन्नरदेव"
	case "RAJOPUR SADAT AN":
		sectionNameHindi = "राजोपुर सादात आं0"
	case "RAJOPURSADAT AN H.NO 126-END":
		sectionNameHindi = "राजोपुर सादात आं०म०सं०१२६ से अन्‍त तक"
	case "RAJOURA":
		sectionNameHindi = "रजौडा"
	case "RAJPUR CHOMELA GAIR AB.AD":
		sectionNameHindi = "राजपुर चौमेला गैर आं0"
	case "RAJPUR KESARIYAVALA":
		sectionNameHindi = "राजपुर केसरियावाला"
	case "RAJPUR NAVADA":
		sectionNameHindi = "राजपुर नवादा"
	case "RAJPUR PARSU":
		sectionNameHindi = "राजपुर परसू"
	case "RAJPUR":
		sectionNameHindi = "राजपुर"
	case "RAJPURA SUAR":
		sectionNameHindi = "रजपुरा स्‍वार"
	case "RAJPURA TANDA":
		sectionNameHindi = "रजपुरा टाण्‍डा"
	case "RAJPURA":
		sectionNameHindi = "रजपुरा"
	case "RAJPURKOT":
		sectionNameHindi = "राजपुर कोटा"
	case "RAJUPUR MILAK":
		sectionNameHindi = "राजूपुर मिलक"
	case "RAJUPURA H.NO 1 TO 239":
		sectionNameHindi = "रजपुरा म0स0 १ से २३९ तक"
	case "RAM NAGAR":
		sectionNameHindi = "रामनगर"
	case "RAM SAHAYWALI":
		sectionNameHindi = "राम सहायवाली"
	case "RAM TALAIYA ANSHIK":
		sectionNameHindi = "राम तैलया आंशिक"
	case "RAM TALAIYA LINEPAR":
		sectionNameHindi = "रामतलैया लाइनपार"
	case "RAMANAVALA":
		sectionNameHindi = "रमनावाला"
	case "RAMAPUR BALBHADRA JAHEED":
		sectionNameHindi = "रामपुर बलभद्र जहीद"
	case "RAMAPUR BHILA AANSHIK":
		sectionNameHindi = "रामपुर भीला आंशिक"
	case "RAMAPUR GHOGHAR":
		sectionNameHindi = "रामपुर घोघर"
	case "RAMAPUR MAIGAN":
		sectionNameHindi = "रामपुर मैगन"
	case "RAMAPUR PATTI GUJAR":
		sectionNameHindi = "रामपुर पट्टी गूजर"
	case "RAMATHERA":
		sectionNameHindi = "रामठेरा"
	case "RAMBAG":
		sectionNameHindi = "रामबाग"
	case "RAMDASWALI AN0":
		sectionNameHindi = "रामदासवाली आं0"
	case "RAMDASWALI":
		sectionNameHindi = "रामदासवाली"
	case "RAMESHWAR COLONY MAJHOLA":
		sectionNameHindi = "रामेश्‍वर कालौनी मझौला"
	case "RAMGANGA BIHAR ANSHIK":
		sectionNameHindi = "रामगंगा बिहार आंशिक"
	case "RAMGANGA BIHAR COLONY":
		sectionNameHindi = "रामगंगा बिहार कालौनी"
	case "RAMGANGA BIHAR-1":
		sectionNameHindi = "रामगंगा बिहार-1"
	case "RAMGANGA BIHAR-2":
		sectionNameHindi = "रामगंगा बिहार -2"
	case "RAMGANGA BIHAR":
		sectionNameHindi = "रामगंगा बिहार"
	case "RAMGANGA COLONY":
		sectionNameHindi = "रामगंगा कालोनी"
	case "RAMJIWALA":
		sectionNameHindi = "रामजीवाला"
	case "RAMKAYRI":
		sectionNameHindi = "रामकायरी"
	case "RAMKHEDA":
		sectionNameHindi = "रामखेडा"
	case "RAMLELA PURV":
		sectionNameHindi = "रामीला पूर्व"
	case "RAMNAGAR A.":
		sectionNameHindi = "रामनगर आ०"
	case "RAMNAGAR A0":
		sectionNameHindi = "रामनगर आ0"
	case "RAMNAGAR GANGPUR":
		sectionNameHindi = "रामनगर गंगपुर"
	case "RAMNAGAR KHAGGUVALA":
		sectionNameHindi = "रामनगर खग्गूवाला"
	case "RAMNAGAR KHAGUVALA":
		sectionNameHindi = "रामनगर खागूवाला"
	case "RAMNAGAR LATIFPUR-1":
		sectionNameHindi = "रामनगर लतीफपुर-1"
	case "RAMNAGAR LATIFPUR-2":
		sectionNameHindi = "रामनगर लतीफपुर -2"
	case "RAMNAGAR LATIFPUR-3":
		sectionNameHindi = "रामनगर लतीफपुर-3"
	case "RAMNAGAR MAJHRA":
		sectionNameHindi = "रामनगर मझरा"
	case "RAMNAGAR ORF RAMPURA":
		sectionNameHindi = "रामनगर उर्फ रमपुरा"
	case "RAMNAGAR SHAHIDNAGAR":
		sectionNameHindi = "रामनगर शहीदनगर"
	case "RAMNAGAR":
		sectionNameHindi = "रामनगर"
	case "RAMNAGARGOSAI":
		sectionNameHindi = "रामनगरगोसाई"
	case "RAMNAGARIYA":
		sectionNameHindi = "रामनगरिया"
	case "RAMNAGRIYA":
		sectionNameHindi = "रामनगरिया"
	case "RAMNANGLA":
		sectionNameHindi = "रामनंगला"
	case "RAMNATH KALONI MAY GOVAN ASPATAL":
		sectionNameHindi = "रामनाथ कालोनी मय गोवन अस्‍पताल"
	case "RAMOROOPPUR":
		sectionNameHindi = "रामोरूपपुर"
	case "RAMPUR ASAF":
		sectionNameHindi = "रामपुर आसफ"
	case "RAMPUR ASHA":
		sectionNameHindi = "रामपुर आशा"
	case "RAMPUR BAKALI":
		sectionNameHindi = "रामपुर बकली"
	case "RAMPUR BANGAR":
		sectionNameHindi = "रामपुर बंगर"
	case "RAMPUR BANWARI A.":
		sectionNameHindi = "रामपुर बनवारी ए0"
	case "RAMPUR BANWARI":
		sectionNameHindi = "रामपुर बनवारी"
	case "RAMPUR BHOLA":
		sectionNameHindi = "रामपुर भोला"
	case "RAMPUR BISHNA":
		sectionNameHindi = "रामपुर बिश्‍ना"
	case "RAMPUR CHATHA":
		sectionNameHindi = "रामपुर चाठा"
	case "RAMPUR DAS AN.":
		sectionNameHindi = "रामपुर दास आं०"
	case "RAMPUR DHAMMAN":
		sectionNameHindi = "रामपुर धम्‍मन"
	case "RAMPUR DULLI":
		sectionNameHindi = "रामपुर दुल्‍ली"
	case "RAMPUR GARHI":
		sectionNameHindi = "रामपुरगढी"
	case "RAMPUR KALA":
		sectionNameHindi = "रामपुर कला"
	case "RAMPUR KHADER":
		sectionNameHindi = "रामपुर खादर"
	case "RAMPUR KISHNA":
		sectionNameHindi = "रामपुर किशना"
	case "RAMPUR MANGAL":
		sectionNameHindi = "रामपुर मंगल"
	case "RAMPUR MURAR":
		sectionNameHindi = "रामपुर मुरार"
	case "RAMPUR NAJRANA":
		sectionNameHindi = "रामपुर नजराना"
	case "RAMPUR NOUBAD":
		sectionNameHindi = "रामपुर नौआबाद"
	case "RAMPUR PUNA":
		sectionNameHindi = "रामपुर फूना"
	case "RAMPUR THAKRA":
		sectionNameHindi = "रामपुर ठकरा"
	case "RAMPUR VEERAN":
		sectionNameHindi = "रामपुर वीरान"
	case "RAMPUR VIDAR":
		sectionNameHindi = "रामपुर विडार"
	case "RAMPUR":
		sectionNameHindi = "रामपुर"
	case "RAMPURA AN0":
		sectionNameHindi = "रम्‍पुरा आं0"
	case "RAMPURA BUJURGA-1":
		sectionNameHindi = "रम्‍पुरा बुजुर्ग-1"
	case "RAMPURA BUJURGA-2":
		sectionNameHindi = "रम्‍पुरा बुजुर्ग-2"
	case "RAMPURA":
		sectionNameHindi = "रम्‍पुरा"
	case "RAMPURCHAJMAL":
		sectionNameHindi = "रामपुर छजमल"
	case "RAMPURI":
		sectionNameHindi = "रम्‍पुरी"
	case "RAMPURLASHKARI":
		sectionNameHindi = "रामपुर लश्‍करी"
	case "RAMSAHAYWALA":
		sectionNameHindi = "राम सहायवाला"
	case "RAMSARAY":
		sectionNameHindi = "रामसराय"
	case "RAMUVALA SHEKHU":
		sectionNameHindi = "रामूवाला शेखू"
	case "RANA NANGLA":
		sectionNameHindi = "राना नंगला"
	case "RANANAGLA":
		sectionNameHindi = "रानानंगला"
	case "RANAVLI":
		sectionNameHindi = "रनावली"
	case "RANGDAN":
		sectionNameHindi = "रांगडान"
	case "RANGREZAN":
		sectionNameHindi = "रंगरेजान"
	case "RANI NAGAL ANSHIK":
		sectionNameHindi = "रानी नागल आंशिक"
	case "RANI NAGAL":
		sectionNameHindi = "रानी नागल"
	case "RANI NANGAL ANSHIK":
		sectionNameHindi = "रानी नांगल आंशिक"
	case "RANIKOTA":
		sectionNameHindi = "रानीकोटा"
	case "RANIPUR":
		sectionNameHindi = "रानीपुर"
	case "RANIYATHER":
		sectionNameHindi = "रनियाठेर"
	case "RANUA NAGALA":
		sectionNameHindi = "रनंआ नगला"
	case "RAOPUR BAHAMAN A.":
		sectionNameHindi = "राजोपुर बहमन आ०"
	case "RAOPUR BAHAMAN A":
		sectionNameHindi = "राजोपुर बहमन आ०"
	case "RASDANDIYA":
		sectionNameHindi = "रासडंडिया"
	case "RASHEEPUR":
		sectionNameHindi = "रशीदपुर"
	case "RASHIDPUR GARHI":
		sectionNameHindi = "रशीदपुर गढी"
	case "RASOLPUR NANGLA A.":
		sectionNameHindi = "रसूलपुर नंगला आ०"
	case "RASOOLPUR -1":
		sectionNameHindi = "रसूलपुर 1"
	case "RASOOLPUR -2":
		sectionNameHindi = "रसूलपुर 2"
	case "RASOOLPUR -3":
		sectionNameHindi = "रसूलपुर 3"
	case "RASOOLPUR -4":
		sectionNameHindi = "रसलूपुर 4"
	case "RASOOLPUR -5":
		sectionNameHindi = "रसूलपुर 5"
	case "RASOOLPUR -6":
		sectionNameHindi = "रसूलपुर 6"
	case "RASOOLPUR CHOHRA":
		sectionNameHindi = "रसूलपुर चौहरा"
	case "RASOOLPUR DHAKI":
		sectionNameHindi = "रसूलपुर ढाकी"
	case "RASOOLPUR FAREEDPUR-1":
		sectionNameHindi = "रसूलपुर फरीदपुर 1"
	case "RASOOLPUR FAREEDPUR-2":
		sectionNameHindi = "रसूलपुर फरीदपुर 2"
	case "RASOOLPUR GUJAR DHYAN SINGH":
		sectionNameHindi = "रसूलपुर गूजर ध्‍यान सिंह"
	case "RASOOLPUR PRITHI":
		sectionNameHindi = "रसूलपुर प़थी"
	case "RASOOLPUR SAID AN0":
		sectionNameHindi = "रसूलपुर सैद आं0"
	case "RASOOLPUR SAID":
		sectionNameHindi = "रसूलपुर सैद"
	case "RASOOLPUR SUNBARI":
		sectionNameHindi = "रसूलपुर सुनवारी"
	case "RASOOLPUR":
		sectionNameHindi = "रसूलपुर"
	case "RASOOLPURMUZAFAR H.NO 1 TO 954":
		sectionNameHindi = "रसूलपुर मुजफुरपुर म0स0 1 से 854 तक"
	case "RASOOLPURMUZAFAR H.NO 9 TO 276":
		sectionNameHindi = "रसूलपुर मुजफुरपुर म0स0 ९ से २७६ तक"
	case "RASULA":
		sectionNameHindi = "रसूला"
	case "RASULAPUR HAMIR":
		sectionNameHindi = "रसूलपुर हमीर"
	case "RASULAPUR MOHD KULI . AN.":
		sectionNameHindi = "रसूलपुर मौ0कुली आं0"
	case "RASULAPUR MOHD KULI AN.":
		sectionNameHindi = "रसूलपुर मौ0मौ0कुली आं0"
	case "RASULAPUR NAGLI":
		sectionNameHindi = "रसूलपुर नगली"
	case "RASULDARAN":
		sectionNameHindi = "रसूलदारान"
	case "RASULPUR ABAD AN0":
		sectionNameHindi = "रसूलपुर आबाद आ0"
	case "RASULPUR ABAD":
		sectionNameHindi = "रसूलपुर आबाद"
	case "RASULPUR ALIMUDDIN":
		sectionNameHindi = "रसूलपुर अलीमुदीन"
	case "RASULPUR DAUD":
		sectionNameHindi = "रसूलपुर दाउद"
	case "RASULPUR G.A.":
		sectionNameHindi = "रसूलपुर गै0आ0"
	case "RASULPUR IMMA":
		sectionNameHindi = "रसूलपुर इम्‍मा"
	case "RASULPUR JAGAN":
		sectionNameHindi = "रसूलपुर जागन"
	case "RASULPUR KASBA A.":
		sectionNameHindi = "रसूलपुर कस्‍बा आ0"
	case "RASULPUR MAOH KULI . AN.":
		sectionNameHindi = "रसूलपुर मौ0कुली आं0"
	case "RASULPUR METHEY":
		sectionNameHindi = "रसूलपुरि‍मठठे"
	case "RASULPUR PITTANKA":
		sectionNameHindi = "रसूलपुर पित्‍तनका"
	case "RASULPUR PUNA":
		sectionNameHindi = "रसूलपुर फूना"
	case "RASULPUR SAHBUDDIN":
		sectionNameHindi = "रसूलपुर शाहबुददीन"
	case "RASULPUR VEERAN":
		sectionNameHindi = "रसूलपुर वीरान"
	case "RASULPUR VIRAN GAIR AB.":
		sectionNameHindi = "रसूलपुर वीरान गैर आं0"
	case "RASULPUR":
		sectionNameHindi = "रसूलपुर"
	case "RATANPUR KALA AANSHIK":
		sectionNameHindi = "रतनपुर कला आंशिक"
	case "RATANPUR KUHRD":
		sectionNameHindi = "रतनपुर खुर्द"
	case "RATANPUR RIYAYA":
		sectionNameHindi = "रतनपुर रियाया"
	case "RATANPUR":
		sectionNameHindi = "रतनपुर"
	case "RATANPURA SHUMALI 1":
		sectionNameHindi = "रतनपुरा शुमाली 1"
	case "RATANPURA SHUMALI 2":
		sectionNameHindi = "रतनपुरा शुमाली 2"
	case "RATANPURA-1":
		sectionNameHindi = "रतनपुरा 1"
	case "RATANPURA-2":
		sectionNameHindi = "रतनपुरा 2"
	case "RATANPURA":
		sectionNameHindi = "रतनपुरा"
	case "RATHONDA-1":
		sectionNameHindi = "रठौंडा-1"
	case "RATHONDA-2":
		sectionNameHindi = "रठौंडा-2"
	case "RATUANAGALA-1":
		sectionNameHindi = "रतुआ नगला-1"
	case "RATUANAGALA-2":
		sectionNameHindi = "रतुआ नगला-2"
	case "RATUANAGALA-3":
		sectionNameHindi = "रतुआ नगला-3"
	case "RATUPURA 2":
		sectionNameHindi = "रतुपुरा 2"
	case "RATUPURA":
		sectionNameHindi = "रतुपुरा"
	case "RAVANA -01":
		sectionNameHindi = "रवाना -01"
	case "RAVANA -02":
		sectionNameHindi = "रवाना -02"
	case "RAVANA DARGU":
		sectionNameHindi = "रवाना दरगू"
	case "RAVANA HABAT":
		sectionNameHindi = "रवाना हैबत"
	case "RAVANA":
		sectionNameHindi = "रवाना"
	case "RAVANAGHASI":
		sectionNameHindi = "रवानाघासी"
	case "RAVIDASNAGAR":
		sectionNameHindi = "रविदासनगर"
	case "RAVTI":
		sectionNameHindi = "रावटी"
	case "RAWALHERIKHAJIRI":
		sectionNameHindi = "रावलहेडी खजूरी"
	case "RAWANA SHIKARPUR A.":
		sectionNameHindi = "रावाना शिकारपुर आं0"
	case "RAWANA SHIKARPUR":
		sectionNameHindi = "रवाना शिकारपुर आ0"
	case "RAWANALALA & VISHANPURI":
		sectionNameHindi = "रवानालाला ऊर्फ विशनपुरी"
	case "RAWANIPATTI JBHAHAR":
		sectionNameHindi = "रवानी पटटी जवाहर"
	case "RAWANIPATTI UDA":
		sectionNameHindi = "रवाना पटटी ऊदा"
	case "RAWANPUR":
		sectionNameHindi = "रावणपुर"
	case "RAWAPURI AN0":
		sectionNameHindi = "रवापुरी आ0"
	case "RAYAPUR":
		sectionNameHindi = "रायपुर"
	case "RAYPUR -1":
		sectionNameHindi = "रायपुर 1"
	case "RAYPUR BEGA":
		sectionNameHindi = "रायपुर बेगा"
	case "RAYPUR KHURD ORF SYOHARA":
		sectionNameHindi = "रायपुर खुर्द उर्फ स्‍योहारा"
	case "RAYPUR LAKRA":
		sectionNameHindi = "रायपुर लकडा"
	case "RAYPUR MALIYABAD":
		sectionNameHindi = "रायपुर मलियाबाद"
	case "RAYPUR MULAK":
		sectionNameHindi = "रायपुर मलूक"
	case "RAYPUR-2":
		sectionNameHindi = "रायपुर 2"
	case "RAYPUR-3":
		sectionNameHindi = "रायपुर 3"
	case "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 1":
		sectionNameHindi = "रज़ा कालोनी अहमद नगर व शरणार्थी क्‍वाटर्स एंव वी पी कालोनी 1"
	case "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 2":
		sectionNameHindi = "रज़ा कालोनी अहमद नगर व शरणार्थी क्‍वाटर्स एंव वी पी कालोनी 2"
	case "RAZA KALONI AHMAD NGAR SHARDARTHI KWATARS AWM VP KALONI 3":
		sectionNameHindi = "रज़ा कालोनी अहमद नगर व शरणार्थी क्‍वाटर्स एंव वी पी कालोनी 3"
	case "RAZA NAGAR-1":
		sectionNameHindi = "रजानगर 1"
	case "RAZA NAGAR-2":
		sectionNameHindi = "रजानगर 2"
	case "RAZANAGAR":
		sectionNameHindi = "रजानगर"
	case "REHADWA":
		sectionNameHindi = "रेहडवा"
	case "REHAR AN0":
		sectionNameHindi = "रेहड आ0"
	case "REHAT GANJ":
		sectionNameHindi = "रहट गंज"
	case "REHMAN NAGAR":
		sectionNameHindi = "रहमान नगर"
	case "REHMAPUR":
		sectionNameHindi = "रहमापुर"
	case "REHMAPURUSMAPUR":
		sectionNameHindi = "रहमापुर उसमापुर"
	case "REHMAT NAGAR ANSHIK":
		sectionNameHindi = "रहमत नगर आंशिक"
	case "REHMAT NAGAR GALI NO. 1":
		sectionNameHindi = "रहमत नगर गली नं0 1"
	case "REHMAT NAGAR SOCIETY ANSHIK":
		sectionNameHindi = "रहमत नगर सोसायटी आंशिक"
	case "REHPURA":
		sectionNameHindi = "रहपुरा"
	case "REHRA A.":
		sectionNameHindi = "रेहरा आ०"
	case "REHRA A":
		sectionNameHindi = "रेहरा आ०"
	case "REHSENA-1":
		sectionNameHindi = "रहसेना-1"
	case "REHSENA-2":
		sectionNameHindi = "रहसेना-2"
	case "REHTI JAGIR":
		sectionNameHindi = "रहटी जागीर"
	case "REHTOLI":
		sectionNameHindi = "रहटौली"
	case "RELVE STESH N ERIYA":
		sectionNameHindi = "रेलवे स्टेश न एरिया"
	case "REVERI KALAN 01":
		sectionNameHindi = "रेवडी कलां 01"
	case "REVERI KALAN 02":
		sectionNameHindi = "रेवडी कलां ०२"
	case "RIVDI KHURD":
		sectionNameHindi = "रेवडी खुर्द"
	case "ROARA KALAN-1":
		sectionNameHindi = "रौरा कलां-1"
	case "ROARA KALAN-2":
		sectionNameHindi = "रौरा कलां-2"
	case "ROARA KHURD-1":
		sectionNameHindi = "रौरा खुर्द-1"
	case "ROARA KHURD-2":
		sectionNameHindi = "रौरा खुर्द-2"
	case "ROARA KHURD-3":
		sectionNameHindi = "रौरा खुर्द-3"
	case "ROJA":
		sectionNameHindi = "रोजा"
	case "ROJEY WALI PARIA PEERZADA":
		sectionNameHindi = "रौजे वाली पैरिया पीरजादा"
	case "RONDA AANSHIK":
		sectionNameHindi = "रोन्डा आंशिक"
	case "ROONIYA":
		sectionNameHindi = "रौनिया"
	case "ROOPPUR":
		sectionNameHindi = "रूपपुर"
	case "ROOS‍TAMPUR PARMAL GAIRA0":
		sectionNameHindi = "रूस्‍तमपुर परमाल गैरा0"
	case "ROSHANPUR CHIJROLI":
		sectionNameHindi = "रोशनपुर छिजरोली"
	case "ROSHANPUR JAGEER A.":
		sectionNameHindi = "रोशनपुर जागीर आ0"
	case "ROSHANPUR PRATAP":
		sectionNameHindi = "रोशनपुर प्रताप"
	case "ROSHANPUR":
		sectionNameHindi = "रोशनपुर"
	case "RUGDI CHETRAM":
		sectionNameHindi = "रूगडी चेतराम"
	case "RUGDI SAID":
		sectionNameHindi = "रूगडी सैद"
	case "RUKANPUR A.":
		sectionNameHindi = "रूकनपुर आ०"
	case "RUKANPUR":
		sectionNameHindi = "रूकनपुर"
	case "RUPAPUR TANDOLA":
		sectionNameHindi = "रुपपुर टंडोला"
	case "RUPAPUR":
		sectionNameHindi = "रूपापुर"
	case "RUPPUR":
		sectionNameHindi = "रूपपुर"
	case "RUSTAMAPUR BADHAMAR AANSHIK":
		sectionNameHindi = "रुस्तमपुर बढ़मार आंशिक"
	case "RUSTAMNAGAR NEAR CHAPRRA":
		sectionNameHindi = "रूस्‍तमनगर नि0 छपर्रा"
	case "RUSTAMNAGAR NEAR RAYPUR":
		sectionNameHindi = "रूस्‍तमनगर नि0 रायपुर"
	case "RUSTAMNAGAR SAHASAPUR":
		sectionNameHindi = "रुस्तमनगर सहसपुर"
	case "RUSTAMNAGAR SAHASPUR":
		sectionNameHindi = "रुस्तमनगर सहसपुर"
	case "RUSTAMNAGAR":
		sectionNameHindi = "रूस्‍तमनगर"
	case "RUSTAMPUR KHAS":
		sectionNameHindi = "रुस्तमपुर खास"
	case "RUSTAMPUR SHER":
		sectionNameHindi = "रूस्‍तमपुर शेर"
	case "RUSTAMPUR TIGRI ANSHIK":
		sectionNameHindi = "रूस्‍तमपुर तिगरी आंशिक"
	case "RUSTAMPUR WAJID":
		sectionNameHindi = "रूसतमपुर वाजीद"
	case "RUSTAMPUR":
		sectionNameHindi = "रूस्‍तमपुर"
	case "RUSTAMSARAY":
		sectionNameHindi = "रुस्तमसराय"
	case "RUTAMPUR":
		sectionNameHindi = "रूस्‍तमपुर"
	case "SABADLAPUR AN.":
		sectionNameHindi = "सब्‍दलपुर आं0"
	case "SABALGARH":
		sectionNameHindi = "सबलगढ़"
	case "SABALPUR AN0":
		sectionNameHindi = "सबलपुर आं0"
	case "SABALPUR":
		sectionNameHindi = "सबलपुर"
	case "SABDALPUR A.":
		sectionNameHindi = "सब्‍दलपुर आ०"
	case "SABDALPUR A":
		sectionNameHindi = "सब्‍दलपुर आ०"
	case "SABDALPUR KHURDH":
		sectionNameHindi = "सब्‍दलपुर खुर्द"
	case "SABDALPURDEJAHGIR":
		sectionNameHindi = "सब्‍जदलपुर देहजागीर"
	case "SABDALPURI G.A.":
		sectionNameHindi = "सबदलपुरी गै0आ0"
	case "SABHACHANDPUR KESHOAN":
		sectionNameHindi = "सभाचन्‍दपुर केशो आं०"
	case "SABHACHANDPUR MOHAN URF MAKHWADA":
		sectionNameHindi = "सभाचन्‍दपुर मोहन उ र्फ मख्‍ज्ञव्ाडा"
	case "SABHACHANDPUR":
		sectionNameHindi = "सभाचन्‍दपुर"
	case "SABJI MANDI HARTHALA":
		sectionNameHindi = "सब्‍जी मण्‍डी हरथला"
	case "SAB‍JIPUR AANSHIK":
		sectionNameHindi = "सब्‍जीपुर आंशिक"
	case "SABNIGRAN":
		sectionNameHindi = "सब्‍नीग्रान"
	case "SABUWALA":
		sectionNameHindi = "साबुवाला"
	case "SADA SHIVPUR GABRI":
		sectionNameHindi = "सदाशि‍वपुर गावडी"
	case "SADABAD":
		sectionNameHindi = "सादाबाद"
	case "SADADOBAI AN.":
		sectionNameHindi = "सददोबेर आं0"
	case "SADADOBAIR AN.":
		sectionNameHindi = "सददोबेर आं0"
	case "SADAFAL A.":
		sectionNameHindi = "सदाफल‍ आ0"
	case "SADAKPUR MALOOK ORF KHICDHI":
		sectionNameHindi = "सादकपुर मलूकपुर उर्फ खिचडी"
	case "SADAKPUR URF BILASPUR":
		sectionNameHindi = "सादकपुर उर्फ बिलासपुर"
	case "SADANPUR":
		sectionNameHindi = "सदनपुर"
	case "SADAR SAFAKHANA -3":
		sectionNameHindi = "सदर सफाखाना -3"
	case "SADAR SAFAKHANA-1":
		sectionNameHindi = "सदर सफाखाना-1"
	case "SADAR SAFAKHANA-2":
		sectionNameHindi = "सदर सफाखाना-2"
	case "SADARAKHEDA":
		sectionNameHindi = "सदराखेडा"
	case "SADARNAPUR":
		sectionNameHindi = "सदानन्‍दपुर"
	case "SADARPUR MATLABPUR":
		sectionNameHindi = "सदरपुर मतलबपुर"
	case "SADARPUR":
		sectionNameHindi = "सदरपुर"
	case "SADAT 01":
		sectionNameHindi = "सादात 01"
	case "SADAT 02":
		sectionNameHindi = "सादात 02"
	case "SADAT DAKSHINI AANSHIK":
		sectionNameHindi = "सादात दक्षिणी आंशिक"
	case "SADAT NAGAR B.A.":
		sectionNameHindi = "सादातनगर बी0ए0"
	case "SADAT PASH‍CHIMI AANSHIK":
		sectionNameHindi = "सादात पश्‍चिमी आंशिक"
	case "SADAT SAHADRI":
		sectionNameHindi = "सादात सहदरी"
	case "SADAT UT‍TRI":
		sectionNameHindi = "सादात उत्‍तरी"
	case "SADAT":
		sectionNameHindi = "सादात"
	case "SADATABAD":
		sectionNameHindi = "सादताबाद"
	case "SADATNAGAER A0":
		sectionNameHindi = "सादातनगर ए0"
	case "SADATNAGAR":
		sectionNameHindi = "सादातनगर"
	case "SADATPUR":
		sectionNameHindi = "सादातपुर"
	case "SADATPURGADI":
		sectionNameHindi = "सादातपुरगढी"
	case "SADDIKNAGAR":
		sectionNameHindi = "सददीक नगर"
	case "SADHARANPUR":
		sectionNameHindi = "सधारनपुर"
	case "SADKABAD":
		sectionNameHindi = "सादकाबाद"
	case "SADRUDDINNAGAR":
		sectionNameHindi = "सदरूददीननगर"
	case "SADULLA NAGAR":
		sectionNameHindi = "सादुल्‍लानगर"
	case "SADULLAPUR KHANAPUR":
		sectionNameHindi = "सादुल्‍लाखानपुर"
	case "SADULLAPUR":
		sectionNameHindi = "सादुल्‍लापुर"
	case "SADUPURA":
		sectionNameHindi = "सदूपुरा"
	case "SAEDABAD":
		sectionNameHindi = "सईदाबाद"
	case "SAEED NAGAR":
		sectionNameHindi = "सईदनगर"
	case "SAFEELPUR":
		sectionNameHindi = "सफीलपुर"
	case "SAGAR SARAI ANSHIK":
		sectionNameHindi = "सागर सराय आंशिक"
	case "SAHABGANJ":
		sectionNameHindi = "साहबगंज"
	case "SAHABPUR URF HINDU PUR H.NO 151-END":
		sectionNameHindi = "शाहबपुर ऊर्फ हिन्‍दूपुर म०सं०१५१ से अन्‍त तक"
	case "SAHABPUR URF HINDUPUR H.N 1-150":
		sectionNameHindi = "शाहबपुर ऊर्फ हिन्‍दूपुर म०सं० १ से १५०तक"
	case "SAHABPURA RATAN SINGH":
		sectionNameHindi = "शाहबपुरा रतन सिंह"
	case "SAHACHANDAN A.":
		sectionNameHindi = "शाहचन्‍दन आ०"
	case "SAHADPUR GULAL":
		sectionNameHindi = "शाहदपुर गुलाल,"
	case "SAHANPUR NANU":
		sectionNameHindi = "साहनपुर नानू"
	case "SAHANPUR NAWADA":
		sectionNameHindi = "साहनपुर नवादा"
	case "SAHANPUR SANTOKI":
		sectionNameHindi = "साहनपुर सन्‍तोकी"
	case "SAHANPUR":
		sectionNameHindi = "साहनपुर"
	case "SAHARAWALA":
		sectionNameHindi = "सहारावाला"
	case "SAHARIYA JAWHAR":
		sectionNameHindi = "सहरिया जवाहर"
	case "SAHASAPURI":
		sectionNameHindi = "सहसपुरी"
	case "SAHAVIAN KHURD":
		sectionNameHindi = "सहविया खुर्द"
	case "SAHAVIYAN KALAN":
		sectionNameHindi = "सहवियां कलां"
	case "SAHBAZPUR":
		sectionNameHindi = "शहबाजपुर"
	case "SAHJANPUR ROSHAN":
		sectionNameHindi = "शाहजहॉपुर रोशन"
	case "SAHLIPUR AB.SATTAR":
		sectionNameHindi = "शाहलीपुर अ०सत्‍तार"
	case "SAHLIPUR ASHA":
		sectionNameHindi = "शाहलीपुर आशा"
	case "SAHLIPUR HEERA":
		sectionNameHindi = "शाहलीपुर हीरा"
	case "SAHMUZAFARPUR":
		sectionNameHindi = "शाहमुजफफरपुर"
	case "SAHPUR SUKKHA":
		sectionNameHindi = "शाहपुर सुक्‍खा"
	case "SAHRIYA AANSHIK":
		sectionNameHindi = "सहरियॉ आंशिक"
	case "SAHRIYA DRAJ":
		sectionNameHindi = "सहरिया दराज"
	case "SAHRIYA LAKKKHA":
		sectionNameHindi = "सहरिया लक्‍खा"
	case "SAHRIYA NARPAT":
		sectionNameHindi = "सहरिया नरपत"
	case "SAHU NAGLA":
		sectionNameHindi = "साहू नगला"
	case "SAHUKARA-1":
		sectionNameHindi = "साहूकारा-1"
	case "SAHUKARA-2":
		sectionNameHindi = "साहूकारा-2"
	case "SAHUKARA-3":
		sectionNameHindi = "साहूकारा-3"
	case "SAHUKARA-4":
		sectionNameHindi = "साहूकारा-4"
	case "SAHUKARA-5":
		sectionNameHindi = "साहूकारा-5"
	case "SAHUKARA-6":
		sectionNameHindi = "साहूकारा-6"
	case "SAHUNAGLA":
		sectionNameHindi = "साहूनगला"
	case "SAHUPUR":
		sectionNameHindi = "साहूपुर"
	case "SAHUVAN":
		sectionNameHindi = "साहूवान"
	case "SAHUWALAGOSAIWALA HNO 1 TO 105":
		sectionNameHindi = "साहूवाला गोसाईवाला म0स0 १ से १०५ तक"
	case "SAHUWALAGOSAIWALA HNO 106 TO END":
		sectionNameHindi = "साहूवाला गोसाईवाला म0स0 १०६ से अंत तक"
	case "SAHUWAN":
		sectionNameHindi = "साहूवान"
	case "SAHZAHIR AN. H.NO 1-120":
		sectionNameHindi = "शाहजहीर आं०म०सं०१से १२०"
	case "SAHZAHIR AN.H.NO 1-END":
		sectionNameHindi = "शाहजहीर आं०म०सं०१से अन्‍त तक"
	case "SAHZAHIR AN.H.NO 121-END":
		sectionNameHindi = "शाहजहीर आं०म०सं०१२१ से अन्‍त तक"
	case "SAID ALIGANJ BANZARIYA":
		sectionNameHindi = "सैदअलीगंज वंजरिया"
	case "SAID NAGAR TANDA":
		sectionNameHindi = "सईद नगर टाण्‍डा"
	case "SAIDA":
		sectionNameHindi = "सेढा"
	case "SAIDAPUR BIRU":
		sectionNameHindi = "सैदपुर वीरू"
	case "SAIDKHERI":
		sectionNameHindi = "सैदखेडी"
	case "SAIDNAGAR BAZARPTTI":
		sectionNameHindi = "सैदनगर वजरपटटी"
	case "SAIDNAGAR JANUBI":
		sectionNameHindi = "सैदनगर जनूबी"
	case "SAIDNAGAR URF MADAIYAN PUSE":
		sectionNameHindi = "सईदनगर उ र्फ मडैययान पूसै"
	case "SAIDNAGAR URF NAYAGANVA":
		sectionNameHindi = "सैदनगर उर्फ नयागांव"
	case "SAIDPUR HARIAWALA":
		sectionNameHindi = "सैदपुर हि‍रयावाला"
	case "SAIDPUR KHADDAR (GAIRA.)":
		sectionNameHindi = "सैदपुर खददर (गैरा0)"
	case "SAIDPUR MAFI":
		sectionNameHindi = "सैदपुर माफी"
	case "SAIDPUR MIRA":
		sectionNameHindi = "सैदपुर मीरा"
	case "SAIDPUR":
		sectionNameHindi = "सैदपुर"
	case "SAIDPURA GAJJU":
		sectionNameHindi = "सैदपुरा गज्‍जू"
	case "SAIDPURA":
		sectionNameHindi = "सैदपुरा"
	case "SAIDPURI GAIR AB.AD":
		sectionNameHindi = "सैदपुरी गैर आं0"
	case "SAIDPURI":
		sectionNameHindi = "सैदपुरी"
	case "SAIDPURIMAHICHAND":
		sectionNameHindi = "सैदपुरीमहीचन्‍द"
	case "SAIFABAD":
		sectionNameHindi = "सैफाबाद"
	case "SAIFAPUR CHITTU":
		sectionNameHindi = "सैफपुर चित्तू"
	case "SAIFAPUR MALHU":
		sectionNameHindi = "सैफपुर मल्हू"
	case "SAIFAPUR PALLA":
		sectionNameHindi = "सैफपुर पल्ला"
	case "SAIFNI -01":
		sectionNameHindi = "सैफनी -01"
	case "SAIFNI -02":
		sectionNameHindi = "सैफनी -0२"
	case "SAIFNI -03":
		sectionNameHindi = "सैफनी -03"
	case "SAIFNI -04":
		sectionNameHindi = "सैफनी -04"
	case "SAIFNI -05":
		sectionNameHindi = "सैफनी -05"
	case "SAIFNI -06":
		sectionNameHindi = "सैफनी -06"
	case "SAIFNI -07":
		sectionNameHindi = "सैफनी -07"
	case "SAIFNI -08":
		sectionNameHindi = "सैफनी -08"
	case "SAIFNI -09":
		sectionNameHindi = "सैफनी -09"
	case "SAIFNI -10":
		sectionNameHindi = "सैफनी -10"
	case "SAIFNI -11":
		sectionNameHindi = "सैफनी -11"
	case "SAIFPUR BANGAR":
		sectionNameHindi = "सैफपुर बंगर"
	case "SAIFPUR JAGNA":
		sectionNameHindi = "सैफपुर जगना"
	case "SAIFPUR KHADAR":
		sectionNameHindi = "सैफपुर खादर"
	case "SAIHAL AITMALI (GAIRA.)":
		sectionNameHindi = "सैहल ऐतमाली (गैरा0)"
	case "SAIHAL MUSTAHKAM":
		sectionNameHindi = "सैहल मुस्‍तहकम"
	case "SAIJNA":
		sectionNameHindi = "सैजना"
	case "SAIJNI MUSTAHAKAM":
		sectionNameHindi = "सैजनी मुस्तकम"
	case "SAILATHAN":
		sectionNameHindi = "सैलाथान"
	case "SAINAWALA":
		sectionNameHindi = "सैनावाला"
	case "SAINDWAR A":
		sectionNameHindi = "सैन्‍दवार आ०"
	case "SAINDWAR":
		sectionNameHindi = "सैन्‍दवार आ०"
	case "SAINI WALI GALI BHADORA":
		sectionNameHindi = "सैनी वाली गली भदोरा"
	case "SAIWALA":
		sectionNameHindi = "साइवाला"
	case "SAJAWALPUR":
		sectionNameHindi = "सजावलपुर"
	case "SAKAT PUR":
		sectionNameHindi = "सकत पुर"
	case "SAKATPURA":
		sectionNameHindi = "सकटपुरा"
	case "SAKATUBA":
		sectionNameHindi = "सकटुवा"
	case "SAKKA AVVAL MUGALPURA FIRST":
		sectionNameHindi = "सक्‍का अव्‍वल मुगलपुरा प्रथम"
	case "SAKLI":
		sectionNameHindi = "साकली"
	case "SAKTALPUR MILAK":
		sectionNameHindi = "सकतलपुर"
	case "SAK‍TOO NAGLA AANSHIK":
		sectionNameHindi = "सक्‍टू नगला आंशिक"
	case "SALABA":
		sectionNameHindi = "सलावा"
	case "SALABATPUR GIRDHAR":
		sectionNameHindi = "सलावतपुर गि‍रधर"
	case "SALABATPUR":
		sectionNameHindi = "सलावतपुर"
	case "SALAMATABAD":
		sectionNameHindi = "सलामताबाद"
	case "SALAMPUR A.":
		sectionNameHindi = "सलेमपुर आ०"
	case "SALAR PUR":
		sectionNameHindi = "सलार पुर"
	case "SALARABAD":
		sectionNameHindi = "सलाराबाद"
	case "SALARAPUR B.A":
		sectionNameHindi = "सलारपुर बी ए"
	case "SALAVA":
		sectionNameHindi = "सलावा"
	case "SALEM SARAY":
		sectionNameHindi = "सलेमसराय"
	case "SALEMAPUR":
		sectionNameHindi = "सलेमपुर"
	case "SALEMPUR BANGAR":
		sectionNameHindi = "सलेमपुर बंगर"
	case "SALEMPUR GADHELI":
		sectionNameHindi = "सलेमपुर गधेली"
	case "SALEMPUR JHILLA":
		sectionNameHindi = "सलेमपुर झिल्‍ला"
	case "SALEMPUR KHAS":
		sectionNameHindi = "सलेमपुर खास"
	case "SALEMPUR MATHNA URF PURANPUR":
		sectionNameHindi = "सलेमपुर मथना उर्फ पूरनपुर"
	case "SALEMPUR RUKHADIO":
		sectionNameHindi = "सलेमपुर रूखडियो"
	case "SALEMPUR":
		sectionNameHindi = "सलेमपुर"
	case "SALEMSARAY":
		sectionNameHindi = "सलेमसराय"
	case "SALHAPUR":
		sectionNameHindi = "सल्‍लाहपुर"
	case "SALLAHPUR 2":
		sectionNameHindi = "सल्‍लाहपुर 2"
	case "SALLHAPUR 1":
		sectionNameHindi = "सल्‍लाहपुर 1"
	case "SALMABAD":
		sectionNameHindi = "सालमाबाद"
	case "SALVE NAGAR":
		sectionNameHindi = "साल्‍वेनगर"
	case "SAMADARAM SAHAY":
		sectionNameHindi = "समदाराम सहाय"
	case "SAMADPUR BUNIYADPUR":
		sectionNameHindi = "समदपुर वुन्यिादपुर"
	case "SAMANASARAY AN.":
		sectionNameHindi = "समना सराय आं0"
	case "SAMASPUR HUSSAINPUR":
		sectionNameHindi = "समसपुर हुसैनपुर"
	case "SAMASPUR KHORA URF KHODPURA":
		sectionNameHindi = "समसपुर खोडा ऊर्फ खोडपुरा"
	case "SAMASPUR NASEEB H.NO 1 TO 150":
		sectionNameHindi = "समसमपुर नसीब म0स0 १ से १५० तक"
	case "SAMASPUR URF MANAKPUR":
		sectionNameHindi = "समसपुर उर्फ मानकपुर"
	case "SAMASPUR VEERBHAN":
		sectionNameHindi = "समसपुर वीरभान"
	case "SAMASPUR":
		sectionNameHindi = "समसपुर"
	case "SAMASPURNASEEB H.NO 151 TO END":
		sectionNameHindi = "समसमपुरनसीब म0स0 १५१ से अंतक"
	case "SAMASPURSADDO":
		sectionNameHindi = "समसपुरसद् दो"
	case "SAMATHAL":
		sectionNameHindi = "समाथल"
	case "SAMBHLI GATE":
		sectionNameHindi = "सम्‍भली गेट"
	case "SAMDA CHATURBHUJ":
		sectionNameHindi = "समदा चर्तुभुज"
	case "SAMIPUR AN0":
		sectionNameHindi = "समीपुर आं0"
	case "SAMODIYA-1":
		sectionNameHindi = "समोदिया 1"
	case "SAMODIYA-2":
		sectionNameHindi = "समोदिया 2"
	case "SAMRAT ASHOK NAGAR":
		sectionNameHindi = "सम्राट अशोक नगर"
	case "SANAI":
		sectionNameHindi = "सनाई"
	case "SANAYYA CHAK":
		sectionNameHindi = "सनईया चक"
	case "SANAYYA JATT":
		sectionNameHindi = "सनईया जट"
	case "SANAYYA SUKH":
		sectionNameHindi = "सनईया सुख"
	case "SANDALPUR":
		sectionNameHindi = "सन्दलपुर"
	case "SANIPURA":
		sectionNameHindi = "सानीपुरा"
	case "SANIYA AHMAD":
		sectionNameHindi = "सनइया अहमद"
	case "SANIYAN AN-":
		sectionNameHindi = "सानियान आ0"
	case "SANIYAN AN0":
		sectionNameHindi = "सानियान आं0"
	case "SANJARPUL SULTANPUR":
		sectionNameHindi = "संजरपुर सुल्‍तानपुर"
	case "SANKARA":
		sectionNameHindi = "सनकरा"
	case "SANSARPUR":
		sectionNameHindi = "संसारपुर"
	case "SANTO NANGLI A.":
		sectionNameHindi = "संन्‍तो नंगली आ0"
	case "SANTOMALAN AN0":
		sectionNameHindi = "सन्‍तोमालन आं0"
	case "SANWALDAS":
		sectionNameHindi = "सांवलदास"
	case "SARABA-1":
		sectionNameHindi = "सरावा-1"
	case "SARABA-2":
		sectionNameHindi = "सरावा-2"
	case "SARAEKALA":
		sectionNameHindi = "सरायकला"
	case "SARAI ALAM AN0":
		sectionNameHindi = "सराय आलम आं0"
	case "SARAI BAHAUDDIN":
		sectionNameHindi = "सराय बहाउददीन"
	case "SARAI DADUMBAR":
		sectionNameHindi = "सराय डडुम्‍बर"
	case "SARAI GULJARIMAL":
		sectionNameHindi = "सराय गुलजारीमल"
	case "SARAI HUSAINI BEGUM ANSHIK":
		sectionNameHindi = "सराय हुसैनी बेगम आंशिक"
	case "SARAI IMAM":
		sectionNameHindi = "सराय इमाम"
	case "SARAI IMMA":
		sectionNameHindi = "सराय इम्‍मा"
	case "SARAI JALAL GAIR ABAD":
		sectionNameHindi = "सराय जलाल गैर आबाद"
	case "SARAI JEEVAN":
		sectionNameHindi = "सराय जीवन"
	case "SARAI KAZI ANSHIK":
		sectionNameHindi = "सराय काजी आंशिक"
	case "SARAI KHALSA ANSHIK":
		sectionNameHindi = "सराय खालसा आंशिक"
	case "SARAI KISHAN LAL":
		sectionNameHindi = "सराय किशन लाल"
	case "SARAI KOTHIWALAN ANSHIK":
		sectionNameHindi = "सराय कोठीवालान आंशिक"
	case "SARAI MEER AN H-NO 1-143":
		sectionNameHindi = "सराय मीर आं०म०सं०१से १४३"
	case "SARAI MEER AN H.NO 1-265":
		sectionNameHindi = "सराय मीर आं०म०सं०१ से २६५तक"
	case "SARAI MEER AN H.NO 117-END":
		sectionNameHindi = "सराय मीर आं०म०सं०११७ से अन्‍त तक"
	case "SARAI MEER AN H.NO 266-END":
		sectionNameHindi = "सराय मीर आं०म०सं०२६६ से अनत तक"
	case "SARAI MEER AN0 H.NO01-116":
		sectionNameHindi = "सराय मीर आ०म०सं०१ से ११६तक"
	case "SARAI PUKHTA/PAKKI SARAI ANSHIK":
		sectionNameHindi = "सराय पुख्‍ता/पक्‍की सराय आंशिक"
	case "SARAI PURAINI":
		sectionNameHindi = "सराय पुरैनी"
	case "SARAI SAID GHORA":
		sectionNameHindi = "सराय सैद घोड़ा"
	case "SARAI SHARIF NAGAR GAIR ABAD":
		sectionNameHindi = "सराय शरीफ नगर गैर आबाद"
	case "SARAI SHEKH MAHMOOD ANSHIK":
		sectionNameHindi = "सराय शेख महमूद आंशिक"
	case "SARAIMEER AN H.NO 144-END":
		sectionNameHindi = "सराय मीर आं०म०सं०१४४ से अन्‍त तक"
	case "SARAKATHAL MADHO":
		sectionNameHindi = "सरकथल माधो"
	case "SARASWATI VIHAR ANSHIK":
		sectionNameHindi = "सरस्‍वती विहार आंशिक"
	case "SARAY ASNARA":
		sectionNameHindi = "सराय असनारा"
	case "SARAY DARAVAJA-1":
		sectionNameHindi = "सराय दरवाजा-1"
	case "SARAY DARAVAJA-2":
		sectionNameHindi = "सराय दरवाजा-2"
	case "SARAY JOKHA SINGH":
		sectionNameHindi = "सराय जोखा सिंह"
	case "SARAY KADEEM":
		sectionNameHindi = "सराय कदीम"
	case "SARAY KALAN-1":
		sectionNameHindi = "सराय कलां-1"
	case "SARAY KALAN-2":
		sectionNameHindi = "सराय कलां-2"
	case "SARAY KHAJOOR":
		sectionNameHindi = "सराय खजूर"
	case "SARAY MAHESH -1":
		sectionNameHindi = "सराय महेश -1"
	case "SARAY MAHESH -2":
		sectionNameHindi = "सराय महेश- 2"
	case "SARAY PANJU":
		sectionNameHindi = "सराय पन्जू"
	case "SARAY QAZIYAN":
		sectionNameHindi = "सराय काजियान"
	case "SARAY RAJAB ALI":
		sectionNameHindi = "सराय रजब अली"
	case "SARAY SAHADAT YAR KHAN-1":
		sectionNameHindi = "सराय सहादत यार खॉ-1"
	case "SARAY SAHADAT YAR KHAN-2":
		sectionNameHindi = "सराय सहादत यार खॉ-2"
	case "SARAY SHAK HABEB":
		sectionNameHindi = "सराय शेख हबीब"
	case "SARAY SIAU":
		sectionNameHindi = "सराय स्‍याऊ"
	case "SARAY TALE WALI":
		sectionNameHindi = "सराय तले वाली"
	case "SARAYE JAHANGIR -1":
		sectionNameHindi = "सराये जहाँगीर -1"
	case "SARAYE JAHANGIR -2":
		sectionNameHindi = "सराये जहाँगीर -2"
	case "SARAYE PUKHTA":
		sectionNameHindi = "सराये पुख्ता"
	case "SARBARNAGAR":
		sectionNameHindi = "सरवरनगर"
	case "SARDAR NAGAR AANSHIK":
		sectionNameHindi = "सरदार नगर आंशिक"
	case "SARDARPUR":
		sectionNameHindi = "सरदारपुर"
	case "SARDHARPUR":
		sectionNameHindi = "सरधरपुर"
	case "SARFUDEENNAGAR AN":
		sectionNameHindi = "शर्फुदीननगर आ0"
	case "SARITA NAGAR LINEPAR":
		sectionNameHindi = "सरिता नगर लाइनपार"
	case "SARIYAPUR":
		sectionNameHindi = "सडियापुर"
	case "SARKADA CHATRU AN.":
		sectionNameHindi = "सरकडा चतरू आं0"
	case "SARKADA KHAS AANSHIK":
		sectionNameHindi = "सरकड़ा खास आंशिक"
	case "SARKADA KHAS":
		sectionNameHindi = "सरकड़ा खास"
	case "SARKADA KHERI":
		sectionNameHindi = "सरकडा खेडी"
	case "SARKADA":
		sectionNameHindi = "सरकडा"
	case "SARKADAKAREEM":
		sectionNameHindi = "सरकड़ाकरीम"
	case "SARKADAPRAMAPUR MAFI":
		sectionNameHindi = "सरकडापरमपुर माफी"
	case "SARKADI -1":
		sectionNameHindi = "सरकडी 1"
	case "SARKADI -2":
		sectionNameHindi = "सरकडी 2"
	case "SARKADI":
		sectionNameHindi = "सरकडी"
	case "SARKATHAL SAHANI AN.":
		sectionNameHindi = "सरकथल सानी आं0"
	case "SARTHAL":
		sectionNameHindi = "सरथल"
	case "SARUA DHARAMPUR":
		sectionNameHindi = "सेरूआ धर्मपुर"
	case "SATARAN":
		sectionNameHindi = "सतारन"
	case "SATTI KHEDA":
		sectionNameHindi = "सत्ती खेड़ा"
	case "SATTIYAN":
		sectionNameHindi = "सत्‍तीयान"
	case "SATTO NANGLI A.":
		sectionNameHindi = "सत्‍तो नगंली आ0"
	case "SATUNE SANG -1":
		sectionNameHindi = "सतूने संग -1"
	case "SATUNE SANG -2":
		sectionNameHindi = "सतूने संग -2"
	case "SATUNE SANG -3":
		sectionNameHindi = "सतूने संग -3"
	case "SATUNE SANG -4":
		sectionNameHindi = "सतूने संग -4"
	case "SATUNE SANG -5":
		sectionNameHindi = "सतूने संग -5"
	case "SATVAI":
		sectionNameHindi = "सतबई"
	case "SAUDASPUR":
		sectionNameHindi = "श्‍यौदासपुर"
	case "SAUNDA":
		sectionNameHindi = "सौंदा"
	case "SAYADWARA":
		sectionNameHindi = "सैययदवाडा"
	case "SECTOR - 4B BUDDHI VIHAR":
		sectionNameHindi = "सेक्‍टर - 4बी बुध्दि विहार"
	case "SED NAGAR NEAR ASALATPUR":
		sectionNameHindi = "सैदनगर नि0 असालतपुर"
	case "SEDABAD":
		sectionNameHindi = "सैदाबाद"
	case "SEDI":
		sectionNameHindi = "सेढी"
	case "SEDNAGAR MUNDIYA":
		sectionNameHindi = "सैदनगर मुण्डिया"
	case "SEDPURA URF NAIPURA":
		sectionNameHindi = "सैदपुरा उर्फ नाईपुरा"
	case "SEDU KA MAZHARA":
		sectionNameHindi = "सेढू का मझरा"
	case "SEDULLANAGAR":
		sectionNameHindi = "सैदुल्‍लानगर"
	case "SEEDHI SARAI PASHCHIMI ANSHIK":
		sectionNameHindi = "सीधी सराय पश्चिमी आंशिक"
	case "SEEDHI SARAI PASHCHIMI":
		sectionNameHindi = "सीधी सराय पश्चिमी"
	case "SEELPUR":
		sectionNameHindi = "सीलपुर"
	case "SEEMLA":
		sectionNameHindi = "सीमला"
	case "SEEMLI":
		sectionNameHindi = "सीमली"
	case "SEEPAH":
		sectionNameHindi = "सिपाह"
	case "SEER SHIVRAJ SINGH":
		sectionNameHindi = "सिरशीवराज सिंह"
	case "SEERBASHU CHAND AN0":
		sectionNameHindi = "सीरवासुचन्‍द आ0"
	case "SEERBASHU CHAND":
		sectionNameHindi = "सीरवासुचन्‍द"
	case "SEERWASHU CHAND AN0":
		sectionNameHindi = "सीरवासुचन्‍द आ0"
	case "SEETAPURI DAS SARAI ANSHIK":
		sectionNameHindi = "सीतापुरी दस सराय आंशिक"
	case "SEETAPURI DAS SARAI":
		sectionNameHindi = "सीतापुरी दस सराय"
	case "SEH":
		sectionNameHindi = "सेह"
	case "SEHDOLY":
		sectionNameHindi = "सहदौली"
	case "SEJNINANKAR-1":
		sectionNameHindi = "सेजनीनानकार-1"
	case "SEJNINANKAR-2":
		sectionNameHindi = "सेजनीनानकार-2"
	case "SEJNINANKAR-3":
		sectionNameHindi = "सेजनीनानकार-3"
	case "SEJNINANKAR-4":
		sectionNameHindi = "सेजनीनानकार-4"
	case "SEJNINANKAR-5":
		sectionNameHindi = "सेजनीनानकार-5"
	case "SEJNINANKAR-6":
		sectionNameHindi = "सेजनीनानकार-6"
	case "SEKHPURA":
		sectionNameHindi = "शेखुपुरा"
	case "SELA A.":
		sectionNameHindi = "सेला आ0"
	case "SELA KHEDI":
		sectionNameHindi = "सेलाखेडी"
	case "SELPURA":
		sectionNameHindi = "सेलपुरा"
	case "SEMLA":
		sectionNameHindi = "सेमला"
	case "SEMRA NEAR LADPUR-1":
		sectionNameHindi = "सेमरा नि0 लाडपुर 1"
	case "SEMRA NEAR LADPUR-2":
		sectionNameHindi = "सेमरा नि0 लाडपुर 2"
	case "SENDOALI":
		sectionNameHindi = "सैंडोली"
	case "SENJNA":
		sectionNameHindi = "सैंजना"
	case "SENPUR":
		sectionNameHindi = "सैनपुर"
	case "SENTAKHEDA-1":
		sectionNameHindi = "सेन्‍टाखेडा-1"
	case "SENTAKHEDA-2":
		sectionNameHindi = "सेन्‍टाखेडा-2"
	case "SENTAKHEDA-3":
		sectionNameHindi = "सेन्‍टाखेडा-3"
	case "SENTAKHEDA-4":
		sectionNameHindi = "सेन्‍टाखेडा-4"
	case "SENTAKHEDA-5":
		sectionNameHindi = "सेन्‍टाखेडा-5"
	case "SENTAKHEDA-6":
		sectionNameHindi = "सेन्‍टाखेडा-6"
	case "SETHPUR DHANESAR":
		sectionNameHindi = "सेठपुर धनेसर"
	case "SEWAARAMPUR SAKAT":
		sectionNameHindi = "सेवारामपुर संकट सिंह"
	case "SEWARAM AN0":
		sectionNameHindi = "सेवाराम आं0"
	case "SHABAJPUR":
		sectionNameHindi = "शहबाजपुर"
	case "SHADARAUFR THUNAPUR":
		sectionNameHindi = "शाहदराउर्फ थूनापुर"
	case "SHADI NAGAR HAJIRA-2":
		sectionNameHindi = "शादीनगर हजीरा 2"
	case "SHADINAGAR HARDASPUR":
		sectionNameHindi = "शादीनगर हरदासपुर"
	case "SHADINAGAR HAZIRA 1":
		sectionNameHindi = "शादीनगर हजीरा 1"
	case "SHADINAGAR SAKATPURA":
		sectionNameHindi = "शादीनगर सकटपुरा"
	case "SHADINAGAR":
		sectionNameHindi = "शादीनगर"
	case "SHADINGAR NI0AHAMADNAGAR JAGIR":
		sectionNameHindi = "शादीनगर नि0अहमदनगर जागीर"
	case "SHADIPUR AN H.NO 56-END":
		sectionNameHindi = "शादीपुर आ०म०सं०५६ से अन्‍त तक"
	case "SHADIPUR AN":
		sectionNameHindi = "शादीपुर आं०"
	case "SHADIPUR DALLA":
		sectionNameHindi = "शादीपुर डल्‍ला"
	case "SHADIPUR H.NO1-55":
		sectionNameHindi = "शादीपुर आं०म०सं०1सै 55त्‍क्‍"
	case "SHADIPUR HUSAINPUR":
		sectionNameHindi = "शादीपुर हसैनुपुर"
	case "SHADIPUR IMMAR":
		sectionNameHindi = "शादीपुर इम्‍मा"
	case "SHADIPUR KALA":
		sectionNameHindi = "शादीपुर कला"
	case "SHADIPUR":
		sectionNameHindi = "शादीपुर"
	case "SHAFAYAT NAGAR":
		sectionNameHindi = "शफायतनगर"
	case "SHAFIPUR HEERA":
		sectionNameHindi = "शफीपुर हीरा"
	case "SHAFIYABAD AN.":
		sectionNameHindi = "शफियाबाद आं0"
	case "SHAHAALIPUR UDYACHAND":
		sectionNameHindi = "शाहअलीपुर उत्‍तमचन्‍द"
	case "SHAHABAD GET-1":
		sectionNameHindi = "शाहबाद गेट-1"
	case "SHAHABAD GET-2":
		sectionNameHindi = "शाहबाद गेट-2"
	case "SHAHABAJAPUR KALAN":
		sectionNameHindi = "शाहबाजपुर कला"
	case "SHAHABPURA UMRAV SINGH":
		sectionNameHindi = "शाहबपुरा उमराव सिंह"
	case "SHAHAJAD NAGAR":
		sectionNameHindi = "शहजादनगर"
	case "SHAHAJADAPUR TARU":
		sectionNameHindi = "शहजादपुर तारू"
	case "SHAHALAMPUR":
		sectionNameHindi = "शाहआलमपुर"
	case "SHAHALIPUR JAMAN":
		sectionNameHindi = "शाहअलीपुर जमन"
	case "SHAHALIPUR KOTRA AN0":
		sectionNameHindi = "शाहलीपुरकोटरा आं0"
	case "SHAHALIPUR NAKI":
		sectionNameHindi = "शाहअलीपुर नकी"
	case "SHAHALIPUR NICHAL":
		sectionNameHindi = "शाहअलीपुर नीचल"
	case "SHAHALIPUR":
		sectionNameHindi = "शाहलीपुर ए"
	case "SHAHARYARPUR BANGER":
		sectionNameHindi = "शहरयारपुर बंगर"
	case "SHAHBAD":
		sectionNameHindi = "शाहबाद"
	case "SHAHBAJPUR":
		sectionNameHindi = "शहबाजपुर"
	case "SHAHBAZPUR KHANA":
		sectionNameHindi = "शहबाजपुर खाना"
	case "SHAHBAZPUR":
		sectionNameHindi = "शहबाजपुर"
	case "SHAHDARA":
		sectionNameHindi = "शाहदरा"
	case "SHAHDRA NEAR BHOT":
		sectionNameHindi = "शाहदरा निकट भोट"
	case "SHAHDRA NIKAT DHANPUR":
		sectionNameHindi = "शाहदरा निकट धनपुर"
	case "SHAHI MASJID":
		sectionNameHindi = "शाही मस्जिद"
	case "SHAHJADPUR":
		sectionNameHindi = "शहजादपुर"
	case "SHAHJAHAN PUR GARHI":
		sectionNameHindi = "शाहजहांपुर गढी"
	case "SHAHJAHANPUR":
		sectionNameHindi = "शाहजहांपुर"
	case "SHAHJAHAPUR":
		sectionNameHindi = "शाहजहांपुर"
	case "SHAHKARAMPUR GILARA":
		sectionNameHindi = "शाहकरमपुर गिलाडा"
	case "SHAHLIPUR GADDU":
		sectionNameHindi = "शाहलीपुर गडडू"
	case "SHAHLIPUR":
		sectionNameHindi = "शाहलीपुर"
	case "SHAHLIPURKOTRA H NO 1 TO 178":
		sectionNameHindi = "शाहलीपुरकोटरा म0स0 ०१ सें 178 तक"
	case "SHAHLIPURKOTRA H NO 179 TO END":
		sectionNameHindi = "शाहलीपुरकोटरा म0स0 179 अंत तक"
	case "SHAHNAGAR":
		sectionNameHindi = "शाहनगर"
	case "SHAHNAZARPUR KOTT":
		sectionNameHindi = "शाहनजरपुर कोट"
	case "SHAHPUR ABDULBARI":
		sectionNameHindi = "शाहपुर अब्‍दुलवारी"
	case "SHAHPUR ANSHIK":
		sectionNameHindi = "शाहपुर आंशिक"
	case "SHAHPUR BADWA":
		sectionNameHindi = "शाहपुर बडवा"
	case "SHAHPUR DEV 01":
		sectionNameHindi = "शाहपुरदेव 01"
	case "SHAHPUR DEV 02":
		sectionNameHindi = "शाहपुरदेव 02"
	case "SHAHPUR DHAMDI":
		sectionNameHindi = "शाहपुर धमेडी"
	case "SHAHPUR JAMAL A":
		sectionNameHindi = "शाहपुर जमाल"
	case "SHAHPUR KHEDI":
		sectionNameHindi = "शाहपुर"
	case "SHAHPUR LAL":
		sectionNameHindi = "शाहपुर लाल"
	case "SHAHPUR MIRA":
		sectionNameHindi = "शाहपुर मीरा"
	case "SHAHPUR MUBARAKPUR":
		sectionNameHindi = "शाहपुर मुबारकपुर"
	case "SHAHPUR SAIDU":
		sectionNameHindi = "शाहपुर सैदु"
	case "SHAHPUR SAISU":
		sectionNameHindi = "शाहपुर सैसू"
	case "SHAHPUR TIGRI ANSHIK":
		sectionNameHindi = "शाहपुर तिगरी आंशिक"
	case "SHAHPUR":
		sectionNameHindi = "शाहपुर"
	case "SHAHPURA-1":
		sectionNameHindi = "शहपुरा-1"
	case "SHAHPURA-2":
		sectionNameHindi = "शहपुरा-2"
	case "SHAHPURA-3":
		sectionNameHindi = "शहपुरा-3"
	case "SHAHUVAN A":
		sectionNameHindi = "साहूवान आ०"
	case "SHAHWILLAYAT":
		sectionNameHindi = "शाहविलायत"
	case "SHAHWILLYAT":
		sectionNameHindi = "शाहविलायत"
	case "SHAHZADPUR AN0":
		sectionNameHindi = "शहजादपुर आं0"
	case "SHAHZAHAPUR JASRATH":
		sectionNameHindi = "शाहजहांपुर जसरथ"
	case "SHAIKHAN DAKSHIR":
		sectionNameHindi = "शेखान दक्षिण"
	case "SHAIKHAN DAKSHIRAIN":
		sectionNameHindi = "शेखान दक्षिणी"
	case "SHAIKHAN TARAI":
		sectionNameHindi = "शेखान तराई"
	case "SHAIKHAN UTTARI AN.":
		sectionNameHindi = "शेखान उत्‍तरी आं0"
	case "SHAIKHAN UTTARI":
		sectionNameHindi = "शेखान उत्‍तरी"
	case "SHAIKHPUR LALA":
		sectionNameHindi = "शेखपुर लाला"
	case "SHAIKHPURA PITTHA":
		sectionNameHindi = "शेरपुर पित्‍था"
	case "SHAJADPUR AN0":
		sectionNameHindi = "शहजादपुर आ0"
	case "SHAJADPUR":
		sectionNameHindi = "शहजादपुर"
	case "SHAKARPUR BAHADAR":
		sectionNameHindi = "शाकरपुर बहादर"
	case "SHAKARPURI":
		sectionNameHindi = "शकरपुरी"
	case "SHAKHPURI MEENA":
		sectionNameHindi = "शेखपुरी मीना"
	case "SHAKTI NAGAR LINEPAR":
		sectionNameHindi = "शक्ति नगर लाइनपार"
	case "SHALIPUR BA":
		sectionNameHindi = "शाहलीपुर बी0ए"
	case "SHAMSABAD":
		sectionNameHindi = "शमसाबाद"
	case "SHAMSHABAD URF KHABRIYA":
		sectionNameHindi = "शमशाबाद ऊर्फ खवरिया"
	case "SHAMSHABAD":
		sectionNameHindi = "शमशाबाद"
	case "SHAMSI COLONY DHEEMRI":
		sectionNameHindi = "शमसी कालोनी धीमरी"
	case "SHANAGAR KURALI":
		sectionNameHindi = "शाहनगरकुराली"
	case "SHANKARPUR":
		sectionNameHindi = "शकरपुर"
	case "SHAPPAR SHIKOHPUR":
		sectionNameHindi = "शप्‍पर शिकोहपुर"
	case "SHAPUR BASODHI A.":
		sectionNameHindi = "शाहपुर भसौडी आ०"
	case "SHAPUR BASODI A.":
		sectionNameHindi = "शाहपुर भसौडी आ०"
	case "SHAPUR HARRAE":
		sectionNameHindi = "शाहपुर हर्राय"
	case "SHARAY":
		sectionNameHindi = "सराय"
	case "SHAREEFPUR BANGAR A.":
		sectionNameHindi = "शरीफपुर बंगर ए0"
	case "SHAREEFPUR BANGAR B.A.":
		sectionNameHindi = "शरीफपुर बंगर बी0ए0"
	case "SHAREEFPUR":
		sectionNameHindi = "शरीफपुर"
	case "SHARFUDDINPUR KAIMRY":
		sectionNameHindi = "शरफुददीनपुर कैमरी"
	case "SHARGHAR AN0 HN0 1 TO 376":
		sectionNameHindi = "शेरगढ आ0 म0स0 १ से ३७६ तक"
	case "SHARGHAR AN0":
		sectionNameHindi = "शेरगढ आ0"
	case "SHARIF NAGAR":
		sectionNameHindi = "शरीफ नगर"
	case "SHARIFAPUR":
		sectionNameHindi = "शरीफपुर"
	case "SHARIFPUR KHORAJ":
		sectionNameHindi = "शरीफपुर खोराज"
	case "SHARIFULMALAKPUR":
		sectionNameHindi = "शरीफुलमलकपुर"
	case "SHASTRI NAGAR":
		sectionNameHindi = "शास्‍त्री नगर"
	case "SHAUKATBAGH PASHCHIMI":
		sectionNameHindi = "शौकतबाग पश्चिमी"
	case "SHAURAMAPUR":
		sectionNameHindi = "शौरामपुर"
	case "SHAWILATE":
		sectionNameHindi = "शाहविलायत"
	case "SHAYAMABAD":
		sectionNameHindi = "श्‍यामाबाद"
	case "SHAZAD NAGAR":
		sectionNameHindi = "शहजाद नगर"
	case "SHEESHGARAN":
		sectionNameHindi = "शीशग्रान"
	case "SHEHZADNAGAR":
		sectionNameHindi = "शहजादनगर"
	case "SHEKH SARAI":
		sectionNameHindi = "शेख सराय"
	case "SHEKHAN AN.":
		sectionNameHindi = "शेखान आं0"
	case "SHEKHAN SHUMALI":
		sectionNameHindi = "शेखान शुमाली"
	case "SHEKHAN UTTRI A.":
		sectionNameHindi = "शेखान उत्‍तरी आं0"
	case "SHEKHAPUR BHAVDA":
		sectionNameHindi = "शेखपुर भावडा"
	case "SHEKHOOPUR KHAS":
		sectionNameHindi = "शेखूपुर खास"
	case "SHEKHPUR GARHU":
		sectionNameHindi = "शेखपुर गढ़ू आं0"
	case "SHEKHPUR LALA":
		sectionNameHindi = "शेखपुर लाला"
	case "SHEKHPUR THATH":
		sectionNameHindi = "शेखपुर ठाट"
	case "SHEKHPURA ALAM":
		sectionNameHindi = "शेखपुरा आलम"
	case "SHEKHPURA NOABAD":
		sectionNameHindi = "शेखपुरा नौआबाद"
	case "SHEKHPURA TURK H.NO 68-END":
		sectionNameHindi = "शेखपुरा तुर्क आं०म०सं०६८ से अन्‍त तक"
	case "SHEKHPURA":
		sectionNameHindi = "शेखपुरा"
	case "SHEKHPURATURK AN H.NO 1-67":
		sectionNameHindi = "शेखपुरा तुर्क आं०म०सं०१ से ६७तक"
	case "SHEKHSARAI AN":
		sectionNameHindi = "शेख सराय आं०"
	case "SHEKHUPURA":
		sectionNameHindi = "शेखूपुरा"
	case "SHEKPURI CHOHARD":
		sectionNameHindi = "शेखपुरी चौहड"
	case "SHERAPUR BAHALIN":
		sectionNameHindi = "शेरपुर बहलीन"
	case "SHERAPUR KALIYA":
		sectionNameHindi = "शेरपुर कलिया"
	case "SHERAPUR MAFI URF DHANIYAKHEDA":
		sectionNameHindi = "शेरपुर माफी उर्फ धनियाखेड़ा"
	case "SHERAPUR PATTI":
		sectionNameHindi = "शेरपुर पट्टी"
	case "SHERAPUR RAINI":
		sectionNameHindi = "शेरपुर रैनी"
	case "SHERAPUR":
		sectionNameHindi = "शेरपुर"
	case "SHERAWALA":
		sectionNameHindi = "शेरावाला"
	case "SHERNAGAR NARAINI":
		sectionNameHindi = "शेर नगर नरैनी"
	case "SHERPUR ABHI":
		sectionNameHindi = "शेरपुर अभि"
	case "SHERPUR BALLA":
		sectionNameHindi = "शेरपुर बल्‍ला"
	case "SHERPUR ETMADPUR":
		sectionNameHindi = "शेरपुर एतमादपुर"
	case "SHERPUR HARSARAN":
		sectionNameHindi = "शेरपुर हरसरन"
	case "SHERPUR JAMAL GAIR ABAD":
		sectionNameHindi = "शेरपुर जमाल गैर आबाद"
	case "SHERPUR MAFI":
		sectionNameHindi = "शेरपुर माफी"
	case "SHERPUR MAJRA":
		sectionNameHindi = "शेरपुर मजरा"
	case "SHERPUR":
		sectionNameHindi = "शेरपुर"
	case "SHIKANDARPUR LALMAN":
		sectionNameHindi = "सिकन्‍दरपुर लालमन"
	case "SHIKANDRABAD":
		sectionNameHindi = "सिकन्‍दराबाद"
	case "SHIKARPUR MAJHRA MITHANPUR MAHESH":
		sectionNameHindi = "शिकारपुर मजरा मिठनपुर महेश"
	case "SHIKARPUR":
		sectionNameHindi = "शिकारपुर"
	case "SHIKARUR":
		sectionNameHindi = "शिकारपुर"
	case "SHISHGRAN AN. H.NO.1-121":
		sectionNameHindi = "शीशग्रान आं०म०सं०१से १२१तक"
	case "SHISHGRAN AN.H.NO 122-END":
		sectionNameHindi = "शीशग्रान आं०म०सं०१२२ से अन्‍त तक"
	case "SHIV MANDIR GULABWADI":
		sectionNameHindi = "शिव मन्दिर गुलाबवाडी"
	case "SHIV MANDIR SAGAR SARAI":
		sectionNameHindi = "शिव मन्दिर सागर सराय"
	case "SHIV NAGAR NEAR ASALATPUR":
		sectionNameHindi = "शिवनगर नि0 असालतपुर"
	case "SHIV NAGAR NEAR GULADIYA":
		sectionNameHindi = "शिवनगर नि0 गुलडिया"
	case "SHIV NAGAR":
		sectionNameHindi = "शिव नगर"
	case "SHIV PURI":
		sectionNameHindi = "शिवपुरी"
	case "SHIV VIHAR":
		sectionNameHindi = "शिव विहार"
	case "SHIVAJI NAGAR ANSHIK":
		sectionNameHindi = "शिवाजी नगर आंशिक"
	case "SHIVAJI NAGAR LINEPAR ANSHIK":
		sectionNameHindi = "शिवाजी नगर लाइनपार आंशिक"
	case "SHIVAJI NAGAR LINEPAR":
		sectionNameHindi = "शिवाजी नगर लाइनपार"
	case "SHIVNAGAR NEAR LOHARI":
		sectionNameHindi = "शिवनगर नि0 लोहारी"
	case "SHIVNAGAR NEAR RAMNAGAR":
		sectionNameHindi = "शिवनगर नि0 रामनगर"
	case "SHIVNAGAR NIKAT GANGAPUR":
		sectionNameHindi = "शिवनगर नि0 गंगापुर"
	case "SHIVNAGAR":
		sectionNameHindi = "शिवनगर"
	case "SHIVPURI":
		sectionNameHindi = "शि‍वपुरी"
	case "SHIVUPRI":
		sectionNameHindi = "शिवुपरी"
	case "SHIWALA A.":
		sectionNameHindi = "शिवाला आ0"
	case "SHIWALA A.M.N. 1 TO 124":
		sectionNameHindi = "शिवाला आं0 म0 न0 1 से १२४ तक"
	case "SHIWALA A.M.N.125 TO END":
		sectionNameHindi = "शिवाला आ0म0न0 १२५ से अन्‍त तक"
	case "SHOKAT ALI MARG -1":
		sectionNameHindi = "शौकत अली मार्ग-1"
	case "SHOKAT ALI MARG -2":
		sectionNameHindi = "शौकत अली मार्ग-2"
	case "SHOKAT ALI MARG -3":
		sectionNameHindi = "शौकत अली मार्ग-3"
	case "SHOKAT ALI MARG -4":
		sectionNameHindi = "शौकत अली मार्ग-4"
	case "SHOKAT ALI MARG -5":
		sectionNameHindi = "शौकत अली मार्ग-5"
	case "SHOLAN SHEKH":
		sectionNameHindi = "शौलान शेख"
	case "SHRAIRAFI A.":
		sectionNameHindi = "सराय रफी आ०"
	case "SHRAWANPUR":
		sectionNameHindi = "सरवनपुर"
	case "SHRAYRAFI A.":
		sectionNameHindi = "सराय रफी आ०"
	case "SHRI NAGAR":
		sectionNameHindi = "श्री नगर"
	case "SHUJATPUR KHADER":
		sectionNameHindi = "शुजातपुर खादर"
	case "SHUJATPUR TIKAR A.":
		sectionNameHindi = "शुजातपुर टीकर आं0"
	case "SHUKANANDPUR":
		sectionNameHindi = "सुखानन्‍दपुर"
	case "SHUKKHAPUR MARKANDEYAMPURAM":
		sectionNameHindi = "सुक्‍खापुर मारकेन्‍डपुरम"
	case "SHUKLAN":
		sectionNameHindi = "शुक्‍लान"
	case "SHUMAL KHEDI":
		sectionNameHindi = "शुमालखेडी"
	case "SHUMALKHEDA":
		sectionNameHindi = "शुमालखेड़ा"
	case "SHURPUR GAIR AB.AD":
		sectionNameHindi = "शुरपुर गैर आं0"
	case "SHUTRAKHANA MAY JYOTISHYAN-1":
		sectionNameHindi = "शुतरखाना मय ज्योतिष्यान-1"
	case "SHUTRAKHANA MAY JYOTISHYAN-2":
		sectionNameHindi = "शुतरखाना मय ज्योतिष्यान-2"
	case "SHUTRAKHANA MAY JYOTISHYAN-3":
		sectionNameHindi = "शुतरखाना मय ज्योतिष्यान-3"
	case "SHYAMIWALA AN0":
		sectionNameHindi = "श्‍यामीवाला आं0"
	case "SHYAMLI":
		sectionNameHindi = "श्‍यामली"
	case "SHYAMPUR HADIPUR":
		sectionNameHindi = "श्‍यामपुर हादीपुर"
	case "SHYAMPUR":
		sectionNameHindi = "श्‍यामपुर"
	case "SHYAMPURSARVANKA":
		sectionNameHindi = "श्‍यामपुरसरवांका"
	case "SHYAMSINGHPUR URF BHUDI":
		sectionNameHindi = "श्यामसिंहपुर उर्फ भूड़ी"
	case "SIAU A.":
		sectionNameHindi = "स्‍याऊ आ०"
	case "SICKRONDHA":
		sectionNameHindi = "सिकरौन्‍धा"
	case "SIDARUNAJARPUR":
		sectionNameHindi = "सिडरउनजरपुर"
	case "SIDHARUWALA":
		sectionNameHindi = "सिधारूवाला"
	case "SIDHAVLI":
		sectionNameHindi = "सिढावली"
	case "SIDIYAWALI":
		sectionNameHindi = "सिढियावली"
	case "SIGANI":
		sectionNameHindi = "सीगनी"
	case "SIGANKHEDA-1":
		sectionNameHindi = "सींगनखेडा-1"
	case "SIGANKHEDA-2":
		sectionNameHindi = "सींगनखेडा-2"
	case "SIGANKHEDA-3":
		sectionNameHindi = "सींगनखेडा-3"
	case "SIGANKHEDA-4":
		sectionNameHindi = "सींगनखेडा-4"
	case "SIGANKHEDA-5":
		sectionNameHindi = "सींगनखेडा-5"
	case "SIGANKHEDA-6":
		sectionNameHindi = "सींगनखेडा-6"
	case "SIGANKHEDA-7":
		sectionNameHindi = "सींगनखेडा-7"
	case "SIHALI AITMALI (GAIRA)":
		sectionNameHindi = "सिहाली ऐतमाली (गैरा)"
	case "SIHALI ANSHIK":
		sectionNameHindi = "सिहाली आंशिक"
	case "SIHALI KHADDAR":
		sectionNameHindi = "सिहाली खददर"
	case "SIHALI":
		sectionNameHindi = "सिंहाली"
	case "SIHARI LADDA":
		sectionNameHindi = "सिहारी लददा"
	case "SIHARI NANDA":
		sectionNameHindi = "सिहारी नन्दा"
	case "SIHARI-1":
		sectionNameHindi = "सिहारी-1"
	case "SIHARI-2":
		sectionNameHindi = "सिहारी-2"
	case "SIHARIMALA":
		sectionNameHindi = "सिहारीमाला"
	case "SIHARIYA":
		sectionNameHindi = "सिहरिया"
	case "SIHOR":
		sectionNameHindi = "सिहोर"
	case "SIHORA BAJE ANSHIK":
		sectionNameHindi = "सिहोरा बाजे आंशिक"
	case "SIHORA GIRDHAR":
		sectionNameHindi = "सिहोरा गिरधर"
	case "SIHORA GOVIND ANSHIK":
		sectionNameHindi = "सिहोरा गोविन्‍द आंशिक"
	case "SIHORA NIZAMABAD":
		sectionNameHindi = "सिहोरा निजामाबाद"
	case "SIHORA":
		sectionNameHindi = "सिहोरा"
	case "SIJOULA":
		sectionNameHindi = "सिजौला"
	case "SIJOULI":
		sectionNameHindi = "सिंजौली"
	case "SIKANDARABAD-1":
		sectionNameHindi = "सिकन्‍दरावाद-1"
	case "SIKANDARABAD-2":
		sectionNameHindi = "सिकन्‍दरावाद-2"
	case "SIKANDARAPUR PATTI":
		sectionNameHindi = "सिकंदरपुर पटटी"
	case "SIKANDARPUR BASI":
		sectionNameHindi = "सिकन्‍दरपुर बसी"
	case "SIKANDARPUR BUDH SINGH ABAD GAIR AB.":
		sectionNameHindi = "सिकन्‍दरपुर बुद्व सिंह आबाद गैर आं0"
	case "SIKANDARPUR":
		sectionNameHindi = "सिकन्‍दरपुर"
	case "SIKANDER NAGLA":
		sectionNameHindi = "सिकन्‍दर नंगला"
	case "SIKANDRABAD":
		sectionNameHindi = "सिकन्‍द्राबाद"
	case "SIKARAUL":
		sectionNameHindi = "सिकरौल"
	case "SIKERA":
		sectionNameHindi = "सिकेडा"
	case "SIKHADIYAN -1":
		sectionNameHindi = "सिघाडियान-1"
	case "SIKHADIYAN -2":
		sectionNameHindi = "सिघाडियान-2"
	case "SIKHADIYAN -3":
		sectionNameHindi = "सिघाडियान-3"
	case "SIKHADIYAN -4":
		sectionNameHindi = "सिघाडियान-4"
	case "SIKKAMPUR":
		sectionNameHindi = "सिक्‍कमपुर"
	case "SIKKAWALA":
		sectionNameHindi = "सिक्‍कावाला"
	case "SIKMAPUR PANDE":
		sectionNameHindi = "सीकमपुर पाण्डे"
	case "SIKRAURA":
		sectionNameHindi = "सिकरौरा"
	case "SIKRI BUZURG":
		sectionNameHindi = "सीकरी बुजुर्ग"
	case "SIKRI KHURD":
		sectionNameHindi = "सीकरी खुर्द"
	case "SIKRI":
		sectionNameHindi = "सीकरी"
	case "SIKRODDA AN0":
		sectionNameHindi = "सिकरौडा आं0"
	case "SILAI BARA GAON-1":
		sectionNameHindi = "सिलई बडा गांव-1"
	case "SILAI BARA GAON-2":
		sectionNameHindi = "सिलई बडा गांव-2"
	case "SIMARIYA-1":
		sectionNameHindi = "सिमरिया-1"
	case "SIMARIYA-2":
		sectionNameHindi = "सिमरिया-2"
	case "SIMLA KALA":
		sectionNameHindi = "सीमला कला"
	case "SIMLI":
		sectionNameHindi = "सीमली"
	case "SIMRA":
		sectionNameHindi = "सिमरा"
	case "SINGH A.":
		sectionNameHindi = "सिंघा आ०"
	case "SINGHA":
		sectionNameHindi = "सिंघा आ०"
	case "SINGHADIYAN":
		sectionNameHindi = "सिंघाडियान"
	case "SINGHMAN HAZARI ANSHIK":
		sectionNameHindi = "सिंहमन हजारी आंशिक"
	case "SINGHMAN HAZARI":
		sectionNameHindi = "सिंहमन हजारी"
	case "SINGRA-1":
		sectionNameHindi = "सिंगरा-1"
	case "SINGRA-2":
		sectionNameHindi = "सिंगरा-2"
	case "SIRAKTHAL":
		sectionNameHindi = "सरकथल"
	case "SIRAS KHENDA ANSHIK":
		sectionNameHindi = "सिरस खेंडा आंशिक"
	case "SIRASKHEDHA ANSHIK":
		sectionNameHindi = "सिरसखेड़ा आंशिक"
	case "SIRASKHERA":
		sectionNameHindi = "सिरसखेडा"
	case "SIRASWA GAUD ANSHIK":
		sectionNameHindi = "सिरसवां गौड आंशिक"
	case "SIRASWA HARCHAND":
		sectionNameHindi = "सिरसवां हरचन्‍द"
	case "SIRDHANE BANGAR":
		sectionNameHindi = "सिंरधनी बंगर"
	case "SIRKAPATTI KUNDESARI":
		sectionNameHindi = "सीरकापटटी कुण्‍डेसरी"
	case "SIRKAPATTI RUPAPUR":
		sectionNameHindi = "सीरकापटटी रूपापुर"
	case "SIRKOI ANSHIK":
		sectionNameHindi = "सिरकोई आंशिक"
	case "SIRRA":
		sectionNameHindi = "सिर्रा"
	case "SIRSA INAYTAPUR ANSHIK":
		sectionNameHindi = "सिरसा इनायतपुर आंशिक"
	case "SIRSA THAT":
		sectionNameHindi = "सिरसा ठाठ"
	case "SIRSA":
		sectionNameHindi = "सिरसा"
	case "SIRYABALI":
		sectionNameHindi = "सिरि‍यावली"
	case "SISAUNA-1":
		sectionNameHindi = "सिसोना-1"
	case "SISAUNA-2":
		sectionNameHindi = "सिसौना-2"
	case "SISAUNA":
		sectionNameHindi = "सिसौना"
	case "SISONA A.":
		sectionNameHindi = "सिसौना आ०"
	case "SITAURA":
		sectionNameHindi = "सितौरा"
	case "SITLA":
		sectionNameHindi = "सिठला"
	case "SLAPUR SHAFKATPUR":
		sectionNameHindi = "सलारपुर शफक्‍कतपुर"
	case "SNUPTA":
		sectionNameHindi = "सनुपता"
	case "SOCIETY REHMAT NAGAR ANSHIK":
		sectionNameHindi = "सोसाइटी रहमत नगर आंशिक"
	case "SOFATPUR":
		sectionNameHindi = "सौफतपुर"
	case "SOHANA":
		sectionNameHindi = "सोहना"
	case "SOHARABNAGAR":
		sectionNameHindi = "सोहरावनगर"
	case "SONAKPUR DEHAT":
		sectionNameHindi = "सोनकपुर देहात"
	case "SONAKPUR-1":
		sectionNameHindi = "सोनकपुर 1"
	case "SONAKPUR-2":
		sectionNameHindi = "सोनकपुर 2"
	case "SONAKPUR-3":
		sectionNameHindi = "सोनकपुर 3"
	case "SONAKPUR":
		sectionNameHindi = "सोनकपुर"
	case "SOOPA":
		sectionNameHindi = "सूपा"
	case "SOOTKHADI":
		sectionNameHindi = "सोतखेडी"
	case "SOTIYAN":
		sectionNameHindi = "सोतियान"
	case "STATION ROAD-1":
		sectionNameHindi = "स्टेशन रोड-1"
	case "STATION ROAD-2":
		sectionNameHindi = "स्टेशन रोड-2"
	case "STATION ROAD-3 MAY GOVAN COLONY":
		sectionNameHindi = "स्टेशन रोड -3 मय गोवन कालोनी"
	case "SUANAGLA":
		sectionNameHindi = "सुआनगला"
	case "SUAR KHAS DAKSHNI -1":
		sectionNameHindi = "स्‍वार खास दक्षिणी 1"
	case "SUAR KHAS DAKSHNI -2":
		sectionNameHindi = "स्‍वार खास दक्षिणी 2"
	case "SUAR KHAS DAKSHNI -3":
		sectionNameHindi = "स्‍वार खास दक्षिणी 3"
	case "SUAR KHAS UTTARI":
		sectionNameHindi = "स्‍वार खास उत्‍तरी"
	case "SUAR KHURD":
		sectionNameHindi = "स्‍वार खुर्द"
	case "SUAWALA AN0 H.NO 1 TO 408":
		sectionNameHindi = "सुआवाला आ0 म0स0 १ से ४०८ तक"
	case "SUAWALA AN0 H.NO 1 TO458":
		sectionNameHindi = "सुआवाला आ0 म0स0 1 से458 तक"
	case "SUAWALA AN0 H.NO 251 TO 594":
		sectionNameHindi = "सुआवाला आ0 म0स0 २५१ से ५९४ तक"
	case "SUDNIPUR":
		sectionNameHindi = "शुदनीपुर"
	case "SUFIPUR ANGAD":
		sectionNameHindi = "सूफीपुर अंगद"
	case "SUHAG NAGLA":
		sectionNameHindi = "सुहागनगला"
	case "SUHAGAPUR AN.":
		sectionNameHindi = "सुहागपुर आ0"
	case "SUHAVA":
		sectionNameHindi = "सुहावा"
	case "SUJANPUR":
		sectionNameHindi = "सुजानपुर"
	case "SUJATGANJ KHERA":
		sectionNameHindi = "सुजातगंज खेरा"
	case "SUJAYATGUNJ":
		sectionNameHindi = "सुजायतगंज"
	case "SULATANPUR DOST":
		sectionNameHindi = "सुलतानपुर दोस्त"
	case "SULEMASEKOPUR":
		sectionNameHindi = "सुलेमि‍शकोहपुर"
	case "SULEMASIKHOPUR":
		sectionNameHindi = "सुलेमाशिकोहपुर"
	case "SULTAN AJAMPURBANI GNESH":
		sectionNameHindi = "सुल्‍तान आजमपुरवानी गनेश"
	case "SULTANAPUR KHADDAR MUS":
		sectionNameHindi = "सुल्तानपुर खद्दर मुस्त."
	case "SULTANAPUR MUNDA":
		sectionNameHindi = "सुल्तानपुर मुण्डा"
	case "SULTANAPUR":
		sectionNameHindi = "सुल्तानपुर"
	case "SULTANKA GAIR ABAD":
		sectionNameHindi = "सुल्‍तानका गैर आबाद"
	case "SULTANPUR ABAD":
		sectionNameHindi = "सुल्‍तानपुर आबाद"
	case "SULTANPUR BANGAR":
		sectionNameHindi = "सुलतानपुर बंगर"
	case "SULTANPUR KHADER":
		sectionNameHindi = "सुल्‍तानपुर खादर"
	case "SULTANPUR KHAS":
		sectionNameHindi = "सुत्‍तानुपर खास"
	case "SULTANPUR MEV":
		sectionNameHindi = "सुल्‍तानपुर मेव"
	case "SULTANPUR MUNDA":
		sectionNameHindi = "सुल्तानपुर मुण्डा"
	case "SULTANPUR NIKAT ESAPUR":
		sectionNameHindi = "सुल्‍तानपुर निकट ईसापुर"
	case "SULTANPUR SABHACHAND":
		sectionNameHindi = "सुल्‍तानपुर सभाचन्‍द"
	case "SULTANPUR SADAT":
		sectionNameHindi = "सुल्‍तानपुर सादात"
	case "SULTANPUR TAPPA NANGAL":
		sectionNameHindi = "सुल्‍लतानपुर टप्‍पा नांगल"
	case "SULTANPUR VIRAN GAIR AB.":
		sectionNameHindi = "सुल्‍तानपुर विरान गैर आं0"
	case "SULTANPUR":
		sectionNameHindi = "सुल्‍तानपुर"
	case "SUMALKHEDI":
		sectionNameHindi = "शुमालखेडी"
	case "SUNARI":
		sectionNameHindi = "सुनारी"
	case "SUNARKHEDA":
		sectionNameHindi = "सुनारखेडा"
	case "SUNARO WALA":
		sectionNameHindi = "सुनारो वाला"
	case "SUNDAR NAGAR":
		sectionNameHindi = "सुन्दर नगर"
	case "SUNDARPUR":
		sectionNameHindi = "सुन्दरपुर"
	case "SUNDARWALI":
		sectionNameHindi = "सुनदरवाली"
	case "SUNDRA":
		sectionNameHindi = "सून्‍दरा"
	case "SUNGERPUR":
		sectionNameHindi = "सुंगरपुर"
	case "SUNGHAR":
		sectionNameHindi = "सुनगढ"
	case "SURAJ NAGAR ANSHIK":
		sectionNameHindi = "सूरज नगर आंशिक"
	case "SURAJ NAGAR GULABWADI":
		sectionNameHindi = "सूरज नगर गुलाबवाडी"
	case "SURAJ NAGAR":
		sectionNameHindi = "सूरज नगर"
	case "SURAJPUR KALYAN":
		sectionNameHindi = "सूरजपुर कल्यान"
	case "SURAJPUR":
		sectionNameHindi = "सूरजपुर"
	case "SURANANGLA":
		sectionNameHindi = "सुरानंगला"
	case "SURATSINGHPUR URF NEW VILLAGE-1":
		sectionNameHindi = "सूरतसिंहपुर ऊर्फ नयागांव-1"
	case "SURATSINGHPUR URF NEW VILLAGE-2":
		sectionNameHindi = "सूरतसिंहपुर ऊर्फ नयागांव;2"
	case "SURAZPUR":
		sectionNameHindi = "सूरजपुर"
	case "SURJAN NAGAR":
		sectionNameHindi = "सुरजन नगर"
	case "SURJAPUR":
		sectionNameHindi = "सूरजपुर"
	case "SURPUR AANSU GAIR AB.":
		sectionNameHindi = "सुरपुर आंसू गैर आ0"
	case "SURSENA":
		sectionNameHindi = "सुरसैना"
	case "SURYA NAGAR ANSHIK":
		sectionNameHindi = "सूर्य नगर आंशिक"
	case "SWAHEDI BUJURG":
		sectionNameHindi = "सुवाहेडी बुजुर्ग"
	case "SWAHEDI KHURD":
		sectionNameHindi = "स्‍वाहेडी खुर्द"
	case "SWALEPUR":
		sectionNameHindi = "स्‍वालेपुर"
	case "SWAR KHURD-1":
		sectionNameHindi = "स्‍वार खुर्द-1"
	case "SWAR KHURD-2":
		sectionNameHindi = "स्‍वार खुर्द-2"
	case "SYODARA":
		sectionNameHindi = "स्योड़ारा"
	case "SYONDRA":
		sectionNameHindi = "स्‍योन्‍डरा"
	case "TABEBPUR":
		sectionNameHindi = "तबीबपुर"
	case "TABELA ANSHIK":
		sectionNameHindi = "तबेला आंशिक"
	case "TAGALA":
		sectionNameHindi = "तगाला"
	case "TAGALI":
		sectionNameHindi = "तगाली"
	case "TAH KHURD":
		sectionNameHindi = "टाहखुर्द"
	case "TAH MADAN":
		sectionNameHindi = "टाह मदन"
	case "TAH NAYAK":
		sectionNameHindi = "टाह नायक"
	case "TAHAERPUR":
		sectionNameHindi = "ताहरपुर"
	case "TAHARPUR ASKARAN":
		sectionNameHindi = "ताहरपुर असकरन"
	case "TAHARPUR GULAM EMAMEN":
		sectionNameHindi = "ताहरपुर मदद इमामैन"
	case "TAHARPUR ISHAK AN0":
		sectionNameHindi = "ताहरपुर इश्‍हाक आं0"
	case "TAHARPUR SAID":
		sectionNameHindi = "ताहरपुर सैद"
	case "TAHARPUR TAPPA NANGAL":
		sectionNameHindi = "ताहरपुर टप्‍पा नांगल"
	case "TAHARPUR TAPPA UMRI GAIR AB.AD":
		sectionNameHindi = "ताहरपुर टप्‍पा उमरी गैर आं0"
	case "TAHKLA":
		sectionNameHindi = "टाहकला"
	case "TAHRPUR AB‍BAL AANSHIK":
		sectionNameHindi = "ताहरपुर अब्‍बल आंशिक"
	case "TAIIBPUR QAZI":
		sectionNameHindi = "तैयबपुर काजी"
	case "TAILIYAN":
		sectionNameHindi = "तलीयान"
	case "TAIMOORPUR":
		sectionNameHindi = "तैमूरपुर"
	case "TAJPUR A. H.S. NO. 1 TO 796":
		sectionNameHindi = "ताजपुर आ0 मकान न01 से 796 तक"
	case "TAJPUR A. H.S. NO. 797 TAK.":
		sectionNameHindi = "ताजपुर आ0 मकान न0 797 अन्‍त तक"
	case "TAJPUR A.":
		sectionNameHindi = "ताजपुर आ0"
	case "TAJPUR BEHTA":
		sectionNameHindi = "ताजपुर बेहटा"
	case "TAJPUR LAKHAN":
		sectionNameHindi = "ताजपुर लखन"
	case "TAJPUR MAFI ANSHIK":
		sectionNameHindi = "ताजपुर माफी आंशि‍क"
	case "TAJPUR":
		sectionNameHindi = "ताजपुर"
	case "TAKANPURI":
		sectionNameHindi = "तखनपुरी"
	case "TAKHATPUR":
		sectionNameHindi = "तख्‍तपुर"
	case "TAKHTAPUR HASHA":
		sectionNameHindi = "तख्तपुर हाशा"
	case "TAKHTPUR ALLA URF NANKAR AANSHIK":
		sectionNameHindi = "तख्तपुर अल्ला उर्फ नानकार आंशिक"
	case "TAKI SARAY AN.":
		sectionNameHindi = "तकीसराय आं0"
	case "TAKI SARAY":
		sectionNameHindi = "तकी सराय"
	case "TAKIAWALA ANSHIK":
		sectionNameHindi = "तकियावाला आंशिक"
	case "TAKIPUR BEGA":
		sectionNameHindi = "तकीपुर बेगा"
	case "TAKIPUR HARVANSH":
		sectionNameHindi = "तकीपुर हरवंश"
	case "TAKIPUR MANOHAR":
		sectionNameHindi = "तकीपुर मनोहर"
	case "TAKIPUR NAROTTAM":
		sectionNameHindi = "तकीपुर नरोत्‍त्‍ाम"
	case "TAKIPUR TULSI NG.CH.":
		sectionNameHindi = "तकीपुर तुलसी ना० क्षे०"
	case "TAKIPURA PURAN":
		sectionNameHindi = "तकीपुरा पूरन"
	case "TAKIPURA":
		sectionNameHindi = "तकीपुरा"
	case "TAKIYA -1":
		sectionNameHindi = "तकिया-1"
	case "TAKIYA FATEH SHAH KHAN":
		sectionNameHindi = "तकिया फतेह शाह खा"
	case "TAKIYA GARI PASHIMI":
		sectionNameHindi = "तकिया गढी पश्‍िचमी"
	case "TAKIYA GARI PU.":
		sectionNameHindi = "तकिया गढी पू0"
	case "TAKIYA MEMARAN":
		sectionNameHindi = "तकिया मेमारान"
	case "TAKIYA MUBARAK SHAH KHA":
		sectionNameHindi = "तकिया मुबारक शाह खा"
	case "TAKIYA SARVAR SHAH KHAN":
		sectionNameHindi = "तकिया सरवर शाह खां"
	case "TAKIYA-2":
		sectionNameHindi = "तकिया-2"
	case "TAKSAL":
		sectionNameHindi = "टकसाल"
	case "TALAB CHAMARAN":
		sectionNameHindi = "तालाब चमारान"
	case "TALAB KELE VALA":
		sectionNameHindi = "तालाब केले वाला"
	case "TALAB MAJAR BAGADADI SAHAB":
		sectionNameHindi = "तालाब मजार बगदादी साहब"
	case "TALAB MULLA AIRAM":
		sectionNameHindi = "तालाब मुल्ला ऐरम"
	case "TALAB NIHALUDDIN MAY GHER HASSAN KHAN":
		sectionNameHindi = "तालाब निहालुद्दीन मय घेर हस्सन खा"
	case "TALAB WALI MASJID PEERZADA":
		sectionNameHindi = "तालाब वाली मस्जिद पीरजादा"
	case "TALABPUR-1":
		sectionNameHindi = "तालबपुर-1"
	case "TALABPUR-2":
		sectionNameHindi = "तालबपुर-2"
	case "TALIBABAD":
		sectionNameHindi = "तालिबाबाद"
	case "TALIBPUR":
		sectionNameHindi = "तालिबपुर"
	case "TALMAHAWAR":
		sectionNameHindi = "तालमहावर"
	case "TAMBAKUWALAN ANSHIK":
		sectionNameHindi = "तम्‍बाकूवालान आंशिक"
	case "TAMBOLI":
		sectionNameHindi = "तम्‍बोली"
	case "TAMURABAD":
		sectionNameHindi = "तैमूराबाद"
	case "TANDA 01":
		sectionNameHindi = "टाण्‍डा 01"
	case "TANDA 02":
		sectionNameHindi = "टाण्‍डा 02"
	case "TANDA AALAM":
		sectionNameHindi = "टांडा आलम"
	case "TANDA AFJAL MU":
		sectionNameHindi = "टांडा अफजल मु."
	case "TANDA BAIRAGI":
		sectionNameHindi = "टांण्‍डा बैरागी"
	case "TANDA GOLU":
		sectionNameHindi = "टाण्‍डा गोलू"
	case "TANDA":
		sectionNameHindi = "टान्‍डा"
	case "TANDAAMRAPUR":
		sectionNameHindi = "टांडाअमरपुर"
	case "TANDAMAIDAS AN0 1 TO 150":
		sectionNameHindi = "टांडामाईदास आ0 १ से १५० तक"
	case "TANDAMAIDAS AN0 1 TO 452":
		sectionNameHindi = "टांडामाईदास आ0 १ से 452 तक"
	case "TANDAMAIDAS AN0 1 TO 775":
		sectionNameHindi = "टांडामाईदास आ0 1 से 775 तक"
	case "TANDAMAIDAS AN0 1 TO 794":
		sectionNameHindi = "टांडामाईदास आ0 1 से 794 तक"
	case "TANDAMAIDAS AN0 151 TO END":
		sectionNameHindi = "टांडामाईदास आ0151 से अंत तक"
	case "TANGROLA":
		sectionNameHindi = "तंगरौला"
	case "TANGROLI":
		sectionNameHindi = "तंगरौली"
	case "TAPROULA":
		sectionNameHindi = "टपरौला"
	case "TARAGANG":
		sectionNameHindi = "ताजगंज"
	case "TARAI AN.":
		sectionNameHindi = "तराई आं0"
	case "TARAIR AN.":
		sectionNameHindi = "तराई आ0"
	case "TARAKAULI MANDN":
		sectionNameHindi = "तरकौली मदन"
	case "TARAPUR SYODARA":
		sectionNameHindi = "तारापुर स्योडारा"
	case "TARAPUR":
		sectionNameHindi = "तारापुर"
	case "TARAUA":
		sectionNameHindi = "तरूआ"
	case "TAREEKAMPUR ROOPCHAND":
		sectionNameHindi = "तरीकमपुर रूपचन्‍द"
	case "TAREKAMPUR PARAS":
		sectionNameHindi = "तरीकमपुर पारस"
	case "TARFADLAPAT":
		sectionNameHindi = "तरफदलपत"
	case "TARIKAMPUR DARGO":
		sectionNameHindi = "तरीकमपुरदरगो"
	case "TARIKAMPUR GAIR ABAD":
		sectionNameHindi = "तरीकमपुर गैरआबाद"
	case "TARIKAMPUR":
		sectionNameHindi = "तरीकमपुर"
	case "TARKOULA":
		sectionNameHindi = "तरकौला"
	case "TARKOULI EMMA":
		sectionNameHindi = "तरकौली इम्‍मा"
	case "TASHKA -1":
		sectionNameHindi = "ताशका- 1"
	case "TASHKA- 2":
		sectionNameHindi = "ताशका- 2"
	case "TATARPUR LALU AN0":
		sectionNameHindi = "तातारपुर लालू आं0"
	case "TATARPUR":
		sectionNameHindi = "तातारपुर"
	case "TATATVALAN MAY MAJAR TATASHAH MIYAN":
		sectionNameHindi = "टटट्वालान मय मजार टाटशाह मियां"
	case "TAYABPUR":
		sectionNameHindi = "तैययबपुर"
	case "TAYYABPUR GORVA":
		sectionNameHindi = "तैय्यबपुर गौरवा"
	case "TEEP":
		sectionNameHindi = "टीप"
	case "TEERGARAN":
		sectionNameHindi = "तीरग्रान"
	case "TEEVRI AN0":
		sectionNameHindi = "तीवड़ी आं0"
	case "TEHRI KHAWJA-1":
		sectionNameHindi = "टेहरी ख्‍वाजा-1"
	case "TEHRI KHAWJA-2":
		sectionNameHindi = "टेहरी ख्‍वाजा-2"
	case "TELINAGALI":
		sectionNameHindi = "तेली नंगली"
	case "TELIPADA":
		sectionNameHindi = "तेलीपाडा"
	case "TELIPURA":
		sectionNameHindi = "तेलीपुरा"
	case "TEMRA":
		sectionNameHindi = "टेमरा"
	case "TENDERA":
		sectionNameHindi = "टन्‍डेरा"
	case "TEVARKHAS":
		sectionNameHindi = "तेवरखास"
	case "TEVRAKHAS":
		sectionNameHindi = "तेवरखास"
	case "TEVRAPATTI URF KAJIPURA AANSHIK":
		sectionNameHindi = "तेवरपटटी उर्फ काजीपुरा आंशिक"
	case "THAKURADVARA ARYANAGAR":
		sectionNameHindi = "ठाकुरद्वारा आर्यनगर"
	case "THAKURADVARA TALI":
		sectionNameHindi = "ठाकुरद्वारा ताली"
	case "THAKURAN DAKSHINI":
		sectionNameHindi = "ठाकुरान दक्षिणी"
	case "THAKURAN PASCHIMI WARD -7":
		sectionNameHindi = "ठाकुरान प0 वार्ड-7"
	case "THAKURAN PURVI":
		sectionNameHindi = "ठाकुरान पूर्वी"
	case "THAKURAN":
		sectionNameHindi = "ठाकुरान"
	case "THAKURDUWARA":
		sectionNameHindi = "ठाकुरद्वारा"
	case "THANATIN-1":
		sectionNameHindi = "थानाटीन-1"
	case "THANATIN.2":
		sectionNameHindi = "थानाटीन.2"
	case "THANATIN.3":
		sectionNameHindi = "थानाटीन-3"
	case "THANATIN.4":
		sectionNameHindi = "थानाटीन-4"
	case "THANATIN.5":
		sectionNameHindi = "थानाटीन-5"
	case "THANATIN.6":
		sectionNameHindi = "थानाटीन-6"
	case "THANVALA":
		sectionNameHindi = "थांवला"
	case "THAPAL":
		sectionNameHindi = "थापल"
	case "THARPUR MADAD IMAM":
		sectionNameHindi = "ताहरपुर मदद इमाम"
	case "THAT":
		sectionNameHindi = "ठेट"
	case "THATHERA MOHALLA":
		sectionNameHindi = "ठठेरा मौहल्‍ला"
	case "THATJAT AN.":
		sectionNameHindi = "ठाटजट आं0"
	case "THEEKRI ANSHIK":
		sectionNameHindi = "ठीकरी आंशिक"
	case "THEGA":
		sectionNameHindi = "थैगा"
	case "THIRIA VISHNU":
		sectionNameHindi = "ठिरिया विश्‍नू"
	case "THIRIYA JADID":
		sectionNameHindi = "ठिरिया जदीद"
	case "THIRIYA SALEPUR":
		sectionNameHindi = "ठिरया सालेपुर"
	case "THOTHAR MAY SABJI MANDI -2":
		sectionNameHindi = "ठोठर मय सब्जी मण्डी -2"
	case "THOTHAR MAY SABZI MANDI - 1":
		sectionNameHindi = "ठोठर मय सब्ज़ी मण्डी - 1"
	case "THURALA":
		sectionNameHindi = "थुरैला"
	case "TICKET GANJ":
		sectionNameHindi = "टिकट गंज"
	case "TIGRI GAIR ABAD":
		sectionNameHindi = "तिगरी गैर आबाद"
	case "TIGRIANOOP SINGH":
		sectionNameHindi = "तिगरीअनूप सिह"
	case "TIRAH":
		sectionNameHindi = "तिरहा"
	case "TIRLOKPUR":
		sectionNameHindi = "त्रि‍लोकपुर"
	case "TIRPUDI":
		sectionNameHindi = "तिरपुडी"
	case "TISAVA":
		sectionNameHindi = "तिसावा"
	case "TISOTRA AN0":
		sectionNameHindi = "तिसोतरा आं0"
	case "TITARWALA":
		sectionNameHindi = "तीतरवाला"
	case "TITUS SCHOOL":
		sectionNameHindi = "टाइटस स्‍कूल"
	case "TIVARI URF KADARPUR JASWANT":
		sectionNameHindi = "तीबडी उर्फ कादरपुर जसवन्‍त"
	case "TODA":
		sectionNameHindi = "टोडा"
	case "TODARPUR":
		sectionNameHindi = "टोडरपुर"
	case "TODIPURA":
		sectionNameHindi = "टौडीपुरा"
	case "TOPAKHANA ROAD":
		sectionNameHindi = "तोपखाना रोड"
	case "TRAPUDA":
		sectionNameHindi = "तिरपुडा"
	case "TRILOKPUR":
		sectionNameHindi = "त्रिलोकपुर"
	case "TUKHMAPUR HARVANSH AN.":
		sectionNameHindi = "तुख्‍मापुर हरवंश आं०"
	case "TUKHMAPURHARVANSH AN0":
		sectionNameHindi = "तुखमापुर हरवंश आं०"
	case "TUMADIYA-1":
		sectionNameHindi = "तुमडिया-1"
	case "TUMADIYA-2":
		sectionNameHindi = "तुमडिया-2"
	case "TUMADIYA":
		sectionNameHindi = "तुमडिया"
	case "TUMADIYAKALA":
		sectionNameHindi = "तुमडि़याकलां"
	case "TUMADIYAKALAN":
		sectionNameHindi = "तुमडियाकलां"
	case "TUNGRI":
		sectionNameHindi = "टूंगरी"
	case "TURABNAGAR":
		sectionNameHindi = "तुराबनगर"
	case "TURATPUR":
		sectionNameHindi = "तुरतपुर"
	case "TURKHERA":
		sectionNameHindi = "तुरखेडा"
	case "TUTIPUR GAIRA":
		sectionNameHindi = "तुतीपुर गैरा0"
	case "UCHA":
		sectionNameHindi = "उचा"
	case "UDAIPUR JAGIR":
		sectionNameHindi = "उदयपुर जागीर"
	case "UDAVALA":
		sectionNameHindi = "ऊदावाला"
	case "UDAYEPUR":
		sectionNameHindi = "उदयपुर"
	case "UDAYPUR JAGIR":
		sectionNameHindi = "उदयपुर जागीर"
	case "UDAYPUR NARAULI":
		sectionNameHindi = "उदयपुर नरौली"
	case "UDHAMPUR":
		sectionNameHindi = "ऊधमपुर"
	case "UDHO JOT":
		sectionNameHindi = "उधो जोत"
	case "UDMA VALA":
		sectionNameHindi = "उदमा वाला"
	case "UDPURA KATGHAR ANSHIK":
		sectionNameHindi = "उडपुरा कटघर आंशिक"
	case "UDRANAPUR CHAK URF VIRMAPUR":
		sectionNameHindi = "उदरनपुर चक उर्फ वीरमपुर"
	case "UGHANPUR":
		sectionNameHindi = "उघनपुर"
	case "UGHO PURA":
		sectionNameHindi = "ऊघो पुरा"
	case "UGRSAINPUR":
		sectionNameHindi = "उग्रसैनपुर"
	case "ULAKPUR":
		sectionNameHindi = "उलकपुर"
	case "UMARAGOPALPUR":
		sectionNameHindi = "उमरागोपालपुर"
	case "UMARAPUR ASHA AN.":
		sectionNameHindi = "उमरपुर आशा आं0"
	case "UMARPUR BANGAR B.A":
		sectionNameHindi = "उमरपुर बांगर बी ए"
	case "UMARPUR NATTHAN":
		sectionNameHindi = "उमरपुर नत्‍थन"
	case "UMARPUR":
		sectionNameHindi = "उमरपुर"
	case "UMARPURBERKHERA":
		sectionNameHindi = "उमरपुर बेरखेडा"
	case "UMRAPUR KHADAR":
		sectionNameHindi = "उमरपुर खादर"
	case "UMRARA":
		sectionNameHindi = "उमरारा"
	case "UMRI AN.H.NO 1-157":
		sectionNameHindi = "उमरी आं०म०सं०1से 157तक"
	case "UMRI AN.H.NO 158-END":
		sectionNameHindi = "उमरी आं०म०सं०158 से अन्‍त तक"
	case "UMRI BADI":
		sectionNameHindi = "उमरी बढी"
	case "UMRI BUJURG":
		sectionNameHindi = "ऊमरा बुजुर्ग"
	case "UMRI":
		sectionNameHindi = "उमरी"
	case "UMRIPEER A.":
		sectionNameHindi = "उमरीपीर आ०"
	case "UMRIPEER":
		sectionNameHindi = "उमरीपीर आ०"
	case "UNCHA GAUN 01":
		sectionNameHindi = "ऊंचा गांव 01"
	case "UNCHA GAUN 02":
		sectionNameHindi = "ऊंचा गांव 02"
	case "UNCHA TEELA JAYANTIPUR":
		sectionNameHindi = "ऊॅचा टीला जयन्‍तीपुर"
	case "UNCHAKANI":
		sectionNameHindi = "ऊॅचाकानी"
	case "USAMANPUR":
		sectionNameHindi = "उसमानपुर"
	case "UTTAMPUR BEHLOLPUR":
		sectionNameHindi = "उत्‍तमपुर बहलोलपुर"
	case "VAGHI GOVARDHANPUR":
		sectionNameHindi = "वघी गोवर्धनपुर"
	case "VAHAPUR":
		sectionNameHindi = "वाहपुर"
	case "VAHPUR AANSHIK":
		sectionNameHindi = "वाहपुर आंशिक"
	case "VAISHPUR KUDIYA":
		sectionNameHindi = "वैशपुर कुडिया"
	case "VAJIRAPURLAL AN.":
		sectionNameHindi = "वजीरपुर लाल आं0"
	case "VAKIPUR JATNI":
		sectionNameHindi = "वाकीपुर जटनी"
	case "VALAPUR":
		sectionNameHindi = "वालापुर"
	case "VALIPUR AN H.NO 150-END":
		sectionNameHindi = "वलीपुरा आं०म०सं०१५० से अन्‍त तक"
	case "VALIPURA AN.H.NO 1-149":
		sectionNameHindi = "वलीपुरा आं०म०स०१ से १४९"
	case "VAMNAPURI GET-1":
		sectionNameHindi = "वमनपुरी गेट-1"
	case "VAMNAPURI GET-2":
		sectionNameHindi = "वमनपुरी गेट-2"
	case "VANSH GOPALPUR":
		sectionNameHindi = "वंश गोपालपुर"
	case "VARAHILALPUR MUST0":
		sectionNameHindi = "वराहीलालपुर मुस्त0"
	case "VARVARA KHAS AI0":
		sectionNameHindi = "वरवारा खास ऐ0"
	case "VASIHPUR MAN":
		sectionNameHindi = "वैशपुर मान"
	case "VATHUAKHEDA MU.":
		sectionNameHindi = "वथुआखेड़ा मु."
	case "VEERBHANWALA":
		sectionNameHindi = "वीरभानवाला"
	case "VEERPUR":
		sectionNameHindi = "वीरपुर"
	case "VEERSHAH HAZARI ANSHIK":
		sectionNameHindi = "वीरशाह हजारी आंशिक"
	case "VEERSHAH HAZARI KATGHAR":
		sectionNameHindi = "वीरशाह हजारी कटघर"
	case "VENAJIRPURA URF GHATAMPUR 2":
		sectionNameHindi = "वेनजीरपुरा उर्फ द्याटमपुर 2"
	case "VENAJIRPURA URF GHATAMPUR-1":
		sectionNameHindi = "वेनजीरपुरा उर्फ द्याटमपुर-1"
	case "VICHAULA DHUKI":
		sectionNameHindi = "विचौला ढूकी"
	case "VIDYA NAGAR":
		sectionNameHindi = "विद्या नगर"
	case "VIJAI NANGLA":
		sectionNameHindi = "विजय नंगला"
	case "VIJAY NAGAR SHIVPURI":
		sectionNameHindi = "विजय नगर शिवपुरी"
	case "VIJAY RAMAPUR":
		sectionNameHindi = "विजय रामपुर"
	case "VIJAYPUR":
		sectionNameHindi = "विजयपुर"
	case "VIJIYA":
		sectionNameHindi = "विजईया"
	case "VIKAS NAGAR LINEPAR ANSHIK":
		sectionNameHindi = "विकास नगर लाइनपार आंशिक"
	case "VIKRAMPUR":
		sectionNameHindi = "विक्रमपुर"
	case "VILAKUDAN":
		sectionNameHindi = "विलाकूदान"
	case "VILAVALA":
		sectionNameHindi = "विलावाला"
	case "VINAVALA":
		sectionNameHindi = "विनावाला"
	case "VIRAPUR FATEHAULLAPUR":
		sectionNameHindi = "वीरपुर फतेहउल्लापुर"
	case "VIRAPUR VARIYAR URF KHARAG AANSHIK":
		sectionNameHindi = "वीरपुर वरियार उर्फ खरग आंश्कि"
	case "VIRPUR GAIRA0":
		sectionNameHindi = "वीरपुर गैरा0"
	case "VIRPUR THAN ANSHIK":
		sectionNameHindi = "वीरपुर थान आंशिक"
	case "VIRTHALA":
		sectionNameHindi = "वीरथला"
	case "VIRUWALA":
		sectionNameHindi = "बीरूवाला"
	case "VISART NAGAR URF MAJHRA":
		sectionNameHindi = "विसारत नगर उर्फ मझरा"
	case "VISHANPUR JAGEER":
		sectionNameHindi = "विशनपुर जागीर"
	case "VISHANPUR":
		sectionNameHindi = "वि‍शनपुर"
	case "VISHANPURA":
		sectionNameHindi = "विशनपुरा"
	case "VISHARATNAGAR-1":
		sectionNameHindi = "विशारतनगर-1"
	case "VISHARATNAGAR-2":
		sectionNameHindi = "विशारतनगर-2"
	case "VISHISHT VIHAR":
		sectionNameHindi = "विशिष्‍ट विहार"
	case "VISHNOI SARAI AN.":
		sectionNameHindi = "विश्‍नोई सराय आं०"
	case "VISHNOI SARAI AN.H.N0.1-180":
		sectionNameHindi = "विश्‍नोई सराय आं०म०स०१ से १८०तक"
	case "VISHNOI SARAI AN.H.NO 181-END":
		sectionNameHindi = "विश्‍नोई सराय आं०म०सं०१८१ से अन्‍त तक"
	case "VISHNOI SARAI AN":
		sectionNameHindi = "विश्‍नोई सराय आं०"
	case "VISHNU VIHAR-1":
		sectionNameHindi = "विष्णु विहार-1"
	case "VISHNU VIHAR-2":
		sectionNameHindi = "विष्णु विहार-2"
	case "VISHNU VIHAR-3":
		sectionNameHindi = "विष्णु विहार-3"
	case "VUJHPUR AASHA AANSHIK":
		sectionNameHindi = "वुझपुर आशा आंशिक"
	case "VUJHPUR MAN":
		sectionNameHindi = "वुझपुर मान"
	case "WAHADUR GARH":
		sectionNameHindi = "वहादुर गढ"
	case "WAHADUR KS MAZJHAR":
		sectionNameHindi = "बहादुर का मझरा"
	case "WAHADURPUR":
		sectionNameHindi = "बहादुरपुर"
	case "WAHID NAGAR AN0":
		sectionNameHindi = "वाहिद नगर आं0"
	case "WAJIDAPUR":
		sectionNameHindi = "वाजिदपुर"
	case "WAJIDPUR TIGRI":
		sectionNameHindi = "वाजिदपुर तिगरी"
	case "WAJIDPUR":
		sectionNameHindi = "वाजिदपुर"
	case "WAJIDPURLALA":
		sectionNameHindi = "वाजि‍दपुर लाला"
	case "WAJIRPUR BHAGWAN":
		sectionNameHindi = "वजीरपुर भगवान"
	case "WALA KI SARAI":
		sectionNameHindi = "वाला की सराय"
	case "WANSIJOT":
		sectionNameHindi = "बंशीजोत"
	case "WARSI NAGAR ANSHIK":
		sectionNameHindi = "वारसी नगर आंशिक"
	case "WARSI NAGAR GALI NO. 2 JAMA MASJID":
		sectionNameHindi = "वारसी नगर गली नं0 2 जामा मस्जिद"
	case "WAZEEDPUR":
		sectionNameHindi = "वाजिदपुर"
	case "WAZEERNAGAR":
		sectionNameHindi = "वजीरनगर"
	case "WAZIDPUR":
		sectionNameHindi = "वाजीदपुर"
	case "YAKUBPUR BILLOCH":
		sectionNameHindi = "याकूबपुर बिल्‍लौच"
	case "YAKUBPUR GAIR AB.":
		sectionNameHindi = "याकूबपुर गैर आं0"
	case "YAKUTPUR CHHAPARRA":
		sectionNameHindi = "याकूतपुर छपर्रा"
	case "YARMOHDPUR PRITHVI":
		sectionNameHindi = "यारमौ०पुरपृथवी"
	case "YOOSUFPUR NAGALIYA MAJRA KONDRI":
		sectionNameHindi = "यूसुफपुर नगालिया मजरा कोंडरी"
	case "YUSUF COLONY REHMAT NAGAR":
		sectionNameHindi = "यूसुफ कालौनी रहमत नगर"
	case "YUSUF NAGAR":
		sectionNameHindi = "युसुफनगर"
	case "YUSUF PARK REHMAT NAGAR":
		sectionNameHindi = "यूसुफ पार्क रहमत नगर"
	case "YUSUFA":
		sectionNameHindi = "युसुफा"
	case "YUSUFPUR HAMID":
		sectionNameHindi = "युसूफपुर हमीद"
	case "ZABTAGANJ AN0":
		sectionNameHindi = "जाब्‍तागंज आं0"
	case "ZAFARPUR ANSHU":
		sectionNameHindi = "जाफरपुर आंसू"
	case "ZAGATPUR":
		sectionNameHindi = "जगतपुर"
	case "ZAHID NAGAR ANSHIK":
		sectionNameHindi = "जाहिद नगर आंशिक"
	case "ZAHID NAGAR KARULA":
		sectionNameHindi = "जाहिद नगर करूला"
	case "ZAHIDPUR":
		sectionNameHindi = "जाहिदपुर"
	case "ZALPUR":
		sectionNameHindi = "जालपुर"
	case "ZAMALGANJ":
		sectionNameHindi = "जमालगंज"
	case "ZERO GALI WARSI NAGAR":
		sectionNameHindi = "जीरो गली वारसी नगर"
	case "ZIARAT KHURMEWALI 2":
		sectionNameHindi = "ज्यारत खुर्मेवाली 2"
	case "ZIARAT KHURMEWALI-1":
		sectionNameHindi = "ज्यारत खुर्मेवाली-1"
	case "ZRIFPUR CHATAR":
		sectionNameHindi = "जरीफपुर चतर"
	case "ZULFUKARPUR GARHI":
		sectionNameHindi = "जुल्‍फुकारपुर गढ़ी"
	case "ZUZHAILA A.":
		sectionNameHindi = "जुझैला आ०"
	case "ज्ञKOTWALI AN H.NO 124-END":
		sectionNameHindi = "कोतवाली आं०म०सं०१२४ से अन्‍त तक"
	}

	return sectionNameHindi
}
