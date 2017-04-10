package main

import (
	_ "burstsms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
  orm.RegisterDriver("mysql", orm.DRMySQL)
	db_name := beego.AppConfig.String("mysql_database")
	user := beego.AppConfig.String("mysql_user")
	pass := beego.AppConfig.String("mysql_pass")
  orm.RegisterDataBase("default", "mysql", user + ":" + pass + "@/" + db_name + "?charset=utf8", 30)
}

func main() {
	orm.RunCommand()
	beego.Run()
}
