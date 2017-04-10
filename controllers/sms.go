package controllers

import (
	"github.com/astaxie/beego"
	"burstsms/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"log"
)

// Make a redis pool
var RedisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle: 5,
	Wait: true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

var SmsQueue = work.NewEnqueuer("burstsms_sms", RedisPool)

type SmsController struct {
	beego.Controller
}

func (c *SmsController) Get() {
  c.Data["title"] = "Fill the form"
	c.TplName = "sms.tpl"
}

func (c * SmsController) Post() {
	phone_number := c.GetString("phone_number")
	text := c.GetString("message_body")

// insert the submited sms to database
	o := orm.NewOrm()
	sms := models.Sms{ Text: text, Number: phone_number }
	id, err := o.Insert(&sms)
	if err != nil {
		log.Fatal(err)
	}

// queue the sms for background processing
	_, err = SmsQueue.Enqueue("send_sms", work.Q{"Number": phone_number, "Text": text})
	if err != nil {
		log.Fatal(err)
	}

	c.Data["title"] = "Message saved! ID: " + strconv.FormatInt(id, 10) + ". another message? " 
	c.TplName = "sms.tpl"
}
