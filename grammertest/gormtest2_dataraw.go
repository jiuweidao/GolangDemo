package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/anker-dev/infra/uuidUtil"
)

//var db2 *gorm.DB

func init() {

}

type DataRaw struct {
	Uuid        string `gorm:"column:uuid;primary_key"`
	ProductCode string `gorm:"column:product_code"`
	Sn          string `gorm:"column:sn"`
	Type        string `gorm:"column:type"`
	LocalHour   int    `gorm:"column:local_hour"`
	LocalMinute int    `gorm:"column:local_minute"`
	CreateTime  int64  `gorm:"column:create_time"`
}

func (t *DataRaw) TableName() string {
	return "data_raw"
}

func main() {
	var err error
	db2, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	/*
	查询所有用户
	 */

	db2.LogMode(true)
	//db2.SingularTable(true)//全局禁止使用复数表
	db2.DB().SetMaxIdleConns(10)
	db2.DB().SetMaxOpenConns(100)

	dr := DataRaw{
		Uuid:        uuidUtil.GetUUID(),
		ProductCode: "2",
		Type:        "3",
		LocalHour:   1,
		LocalMinute: 2,
	}

	tx := db2.Begin()
	defer tx.Rollback()
	if err := tx.Create(dr).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")

	}

}
