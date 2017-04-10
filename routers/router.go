package routers

import (
	"burstsms/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/sms", &controllers.SmsController{})
}
