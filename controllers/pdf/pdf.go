/*
   Create and Send Pdf
   curl -X POST -H "Content-Type: application/json" -d '{"total":1900234,"voters":[{"voter_id":2,"ac_number":3,"part_number":5,"section_number":34,"serial_number_in_part":23,"name_english":"baligul hasan","name_hindi":"बलिग","relation_name_english":"Hasan","relation_name_hindi":"हसन","gender":"M","id_card_number":"SJKJFF132JH","district_name_hindi":"मुरादाबाद","district_name_english":"Moradabad","ac_name_english":"Moradabad Urban","ac_name_hindi":"मुरादाबाद नगर","section_name_english":"Civil Lines","section_name_hindi":"सिविल लाइन्स","religion_english":"Muslim","religion_hindi":"मुसलमान","age":34,"vote":0,"email":"abcd_example@test.com","mobile_no":9898272676,"image":"asdshku32%%7%ahssa*71212","updated_on":""}]}' "http://104.197.6.26:8080/api/pdf?mobile_no=9343352734&token=3964d6b3fb85f787"
*/

package pdf

import (
	//token "crypto/rand"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	//"math/rand"
	//"net/mail"
	//"net/smtp"
	//"strconv"
	//"strings"
	//"time"
	"github.com/jung-kurt/gofpdf"

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
	pdf := gofpdf.New("L", "mm", "A4", "")
	header := []string{"h1", "h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "h10", "h11", "h12", "h13", "h14", "h15"}

	// Colored table
	fancyTable := func() {
		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 0)
		// 	Header
		w := []float64{40, 35, 40, 45, 40, 35, 40, 45, 40, 35, 40, 45, 40, 35, 40}
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
		pdf.SetFont("", "", 0)
		// 	Data
		fill := false
		for _, voter := range voters.Voters {
			pdf.CellFormat(w[0], 6, string(voter.Serial_number_in_part), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], 6, string(voter.Part_number), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[2], 6, voter.Name_english+"("+voter.Name_hindi+")", "LR", 0, "", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}
	fancyTable()
	fileName := "tables"
	err = pdf.OutputFileAndClose(fileName)
	//example.Summary(err, fileName)
	// Output:
	// Successfully generated pdf/Fpdf_CellFormat_tables.pdf
}

/*
func sendEmailWithAttachment(otp int, email string, displayName string, mobileNumber int64) error {
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
}*/
