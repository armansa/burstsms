package main

import (
	_ "burstsms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", "mysql", "root:root@/burstsms?charset=utf8", 30)
}

func main() {
	orm.RunCommand()
	beego.Run()
}
