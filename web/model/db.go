package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"z-blog/helper"
	"z-blog/modules/setting"
)

var Db *gorm.DB

func init(){
	sec := setting.Cfg.Section("database")

	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		sec.Key("USER").String(),
		sec.Key("PASSWORD").String(),
		sec.Key("HOST").String(),
		sec.Key("PORT").String(),
		sec.Key("DATABASE_NAME").String(),
		sec.Key("TIME_ZONE").String())

	println(connection)
	Db, err = gorm.Open("mysql", connection)

	if err != nil {
		log.Fatal("Failed to init db:", err)
		panic("failed to connect database")
	}

	Db.LogMode(true)

	initAdmin()
}

func initAdmin() {
	user := new(Admin)
	err := Db.First(user).Error
	if  err != nil {
		username := setting.Cfg.Section("admin").Key("USERNAME").String()
		password := setting.Cfg.Section("admin").Key("PASSWORD").String()
		admin := new(Admin)
		admin.Id = 1
		admin.Username = username
		admin.Password = helper.Encrypt(password)
		Db.Create(admin)
	}
}
