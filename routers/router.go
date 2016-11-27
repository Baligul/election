// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/Baligul/election/controllers"
	"github.com/Baligul/election/controllers/accounts"
	"github.com/Baligul/election/controllers/email/users"
	"github.com/Baligul/election/controllers/email/voters"
	"github.com/Baligul/election/controllers/groups"
	"github.com/Baligul/election/controllers/tasks"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ElectionController{}, "get:Home")
	beego.Router("/api/voters", &controllers.ElectionController{}, "*:GetVoters")
	beego.Router("/api/statistic", &controllers.ElectionController{}, "post:GetStatistic")
	beego.Router("/api/statistics", &controllers.ElectionController{}, "post:GetStatistics")
	beego.Router("/api/otp", &controllers.ElectionController{}, "post:OTP")
	beego.Router("/api/register", &controllers.ElectionController{}, "post:Register")
	beego.Router("/api/list", &controllers.ElectionController{}, "post:GetList")
	beego.Router("/api/voter", &controllers.ElectionController{}, "post:UpdateVoter")
	beego.Router("/api/group", &groups.GroupCtrl{}, "post:CreateGroup")
	beego.Router("/api/groups", &groups.GroupCtrl{}, "get,post:GetGroups")
	beego.Router("/api/group", &groups.GroupCtrl{}, "put:UpdateGroup")
	beego.Router("/api/group", &groups.GroupCtrl{}, "delete:DeleteGroup")
	beego.Router("/api/account", &accounts.AccountCtrl{}, "post:CreateAccount")
	beego.Router("/api/accounts", &accounts.AccountCtrl{}, "get,post:GetAccounts")
	beego.Router("/api/account", &accounts.AccountCtrl{}, "put:UpdateAccount")
	beego.Router("/api/account", &accounts.AccountCtrl{}, "delete:DeleteAccount")
	beego.Router("/api/email/voters", &voters.VotersCtrl{}, "post:CreateAndEmailPdf")
	beego.Router("/api/email/users", &users.UsersCtrl{}, "post:CreateAndEmailPdf")
	beego.Router("/api/task", &tasks.TaskCtrl{}, "post:CreateTask")
	beego.Router("/api/tasks", &tasks.TaskCtrl{}, "get,post:GetTasks")
	beego.Router("/api/task", &tasks.TaskCtrl{}, "put:UpdateTask")
	beego.Router("/api/task", &tasks.TaskCtrl{}, "delete:DeleteTask")
	beego.Router("/api/taskdetail", &tasks.TaskCtrl{}, "get,post:GetTaskDetail")
	beego.Router("/api/read/json", &controllers.ElectionController{}, "post:ReadJson")
}
