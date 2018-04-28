package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func init() {

}

type AuthorityRequest struct {
	Email     string   `json:"email"`
	Authority []string `gorm:"column:authority" json:"authority"`
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


	var authoritys AuthorityRequest
	//var a,b string
	rows, err :=db.Raw("select " +
		"DISTINCT a1.user_email as email," +
		"(select group_concat(authority) from authority_user a2 ) " +
		"from  authority_user  a1").Rows()


	defer rows.Close()
	for rows.Next() {

		rows.Scan(&authoritys)
		fmt.Printf("authority:%v  ",authoritys)

	}


}
