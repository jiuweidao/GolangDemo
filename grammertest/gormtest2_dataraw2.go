package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"time"
)

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

func main() {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true) //全局禁止使用复数表
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

//	var count int

	/*
	SELECT SUM(`value`) FROM data_raw as d0
 WHERE type='BT_CONNECT_TIMES'

 AND sn in
 (
 SELECT DISTINCT sn
 FROM data_raw
 WHERE create_time< UNIX_TIMESTAMP()

 AND type ="PLAY"
 GROUP BY sn
  HAVING SUM(`value`)>1
 )*/
	/*row1 := db.Table("data_raw").Where("type = ?", "BT_CONNECT_TIMES").
		Select("SUM(`value`)").
		Joins("join ( SELECT DISTINCT data_raw.sn " +
		"FROM data_raw " +
		"WHERE data_raw.create_time< UNIX_TIMESTAMP() " +
		"AND type ='PLAY' " +
		"GROUP BY sn " +
		"HAVING SUM(`value`)>1) act " +
		"on act.sn =data_raw.sn").
		Row()

	row1.Scan(&count)*/



	//var actSn []string


	row := db.Table("data_raw").Select("DISTINCT sn").
		Where("type ='PLAY' AND create_time<?",
		time.Now().Unix()).
		Row()
	//row.Scan(&actSn)
	fmt.Println("actSn:", row)

}
