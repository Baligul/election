/*
   Create and Send Pdf
   curl -X POST -H "Content-Type: application/json" -d '{"total":1900234,"voters":[{"voter_id":2,"ac_number":3,"part_number":5,"section_number":34,"serial_number_in_part":23,"name_english":"baligul hasan","name_hindi":"बलिग","relation_name_english":"Hasan","relation_name_hindi":"हसन","gender":"M","id_card_number":"SJKJFF132JH","district_name_hindi":"मुरादाबाद","district_name_english":"Moradabad","ac_name_english":"Moradabad Urban","ac_name_hindi":"मुरादाबाद नगर","section_name_english":"Civil Lines","section_name_hindi":"सिविल लाइन्स","religion_english":"Muslim","religion_hindi":"मुसलमान","age":34,"vote":0,"email":"abcd_example@test.com","mobile_no":9898272676,"image":"asdshku32%%7%ahssa*71212","updated_on":""}]}' "http://104.197.6.26:8080/api/pdf?mobile_no=9343352734&token=cc5b86572d1ad660"
*/

package pdf

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"net/smtp"
	"strconv"
	"github.com/jung-kurt/gofpdf"
	"github.com/scorredoira/email"

	modelAccounts "github.com/Baligul/election/models/accounts"
	modelVoters "github.com/Baligul/election/models/voters"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	//cleanup()
}

type PdfCtrl struct {
	beego.Controller
}

func (e *PdfCtrl) CreateAndSendPdf() {
	var (
		err  error
		num  int64
		user []*modelAccounts.Account
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
	voters := new(modelVoters.Voters)

	err = json.Unmarshal(inputJson, &voters)
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Invalid Json. Unable to parse. Please check your JSON sent as: %s", inputJson)
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	// PDF creation code start here
	//header := []string{"Voter Id", "Name", "Age", "Gender", "Religion", "Mobile No.", "Email", "Relation", "District", "Ac", "Section", "Part No.", "Serial No. in Part", "Vote"}
	header := []string{"Voter Id", "Name", "Age", "Gender", "Religion", "Mobile No."}
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	// Colored table
	fancyTable := func() {
		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("Arial", "B", 16)
		// 	Header
		//w := []float64{40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40}
		w := []float64{40, 40, 40, 40, 40, 40}
		wSum := 0.0
		for _, v := range w {
			wSum += v
		}
		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Arial", "", 0)
		// 	Data
		fill := false
		for _, voter := range voters.Voters {
			pdf.CellFormat(w[0], 6, voter.Id_card_number, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], 6, voter.Name_english, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[2], 6, strconv.Itoa(voter.Age), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[3], 6, voter.Gender, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[4], 6, voter.Religion_english, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[5], 6, strconv.Itoa(voter.Mobile_no), "LR", 0, "", fill, 0, "")
			/*pdf.CellFormat(w[6], 6, voter.Email, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[7], 6, voter.Relation_name_english, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[8], 6, voter.District_name_english, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[9], 6, voter.Ac_name_english+"("+strconv.Itoa(voter.Ac_number)+")", "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[10], 6, voter.Section_name_english+"("+strconv.Itoa(voter.Section_number)+")", "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[11], 6, strconv.Itoa(voter.Part_number), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[12], 6, strconv.Itoa(voter.Serial_number_in_part), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[13], 6, strconv.Itoa(voter.Vote), "LR", 0, "", fill, 0, "")*/
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}
	fancyTable()
	err = pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		responseStatus := modelVoters.NewResponseStatus()
		responseStatus.Response = "error"
		responseStatus.Message = fmt.Sprintf("Could not send the pdf file.")
		responseStatus.Error = err.Error()
		e.Data["json"] = &responseStatus
		e.ServeJSON()
	}

	err = sendEmailWithAttachment(user[0].Email, user[0].Display_name, "hello.pdf")
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

func sendEmailWithAttachment(toEmail string, displayName string, filepath string) error {
	// compose the message
    m := email.NewMessage("Pdf file attached", "Dear " + displayName + "!\n\nPlease find attached the required file.\n\nThanks & Regards,\nElectionUBDA Team")
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