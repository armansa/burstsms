package models

import (
  "github.com/astaxie/beego/orm"
)

type Sms struct {
    Id   int
    Number string `orm:"size(20)"`
    Text string `orm:"size(500)"`
}

func init() {
    orm.RegisterModel(new(Sms))
}
