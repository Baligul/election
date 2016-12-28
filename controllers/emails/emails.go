// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go

/*
   Send Bulk Emails

   curl -X POST -H "Content-Type: application/json" -d '{}' "http://107.178.208.219:80/api/email/voters/slips?mobile_no=9343352734&token=f8a220f5e8d1741d"

   curl -X POST -H "Content-Type: application/json" -d '{}' "http://localhost:8080/api/emails?mobile_no=9343352734&token=f8a220f5e8d1741d"
*/

package emails

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"

	//modelAccounts "github.com/Baligul/election/models/accounts"

	"github.com/astaxie/beego"
)

func init() {
	//cleanup()
}

type EmailsCtrl struct {
	beego.Controller
}

func (e *EmailsCtrl) SendBulkEmails() {
	var (
		err error
		//num                  int64
		//user                 []*modelAccounts.Account
	)

	from := mail.NewEmail("Example User", "shariq@humansystech.com")
	subject := "Demo"
	to := mail.NewEmail("Example User", "baligcoup8@gmail.com")
	content := mail.NewContent("text/plain", "and easy to do anywhere, even with Go")
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
