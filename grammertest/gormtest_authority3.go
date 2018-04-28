package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func init() {

}

type AuthorityUser struct {
	Id         string `gorm:"column:uuid;primary_key" json:"id"`
	UserEmail  string `gorm:"column:user_email" json:"user_email"`
	Authority  string `gorm:"column:authority" json:"authority"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
}

func main() {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker_omc?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true) //全局禁止使用复数表
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	userEmail:="1233"

	/*
	user :=AuthorityUser{}
	db.Where("user_email=?", userEmail).Find(&user)
	fmt.Printf("%v",user)*/

 db.Delete(AuthorityUser{}, "user_email = ?", userEmail)

	/*if err := db.Delete(&AuthorityUser{UserEmail: userEmail}).Error; err != nil {
		log.Printf("Failed to delete firmware package", err)

	}*/
}
