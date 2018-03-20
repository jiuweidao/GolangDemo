package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
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
	db2, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	db2.LogMode(true)
	db2.SingularTable(true) //全局禁止使用复数表
	db2.DB().SetMaxIdleConns(10)
	db2.DB().SetMaxOpenConns(100)

	//dr := DataRaw{}
	//timBegin := timeUtil.CurrentTime()
	type Count struct {
		count int
	}
	var count int
	//db.Where("create_time>?",0).Find(dr).Distinct("City",&result)
	db2.Raw("SELECT count(distinct  sn)FROM `data_raw` WHERE create_time>?", 0).Count(&count)
	fmt.Println("总用户:", count)
	//Find(bson.M{"City":bson.M{}}).Distinct("City",&result)

	a := &[] string{}
	db2.Raw("SELECT distinct  sn FROM `data_raw` WHERE create_time>?", 0).Scan(a)
	fmt.Println("总用户:", a)

	rows, err := db2.Model(&DataRaw{}).Where("create_time>?", 0).Select("distinct  sn").Rows()

	defer rows.Close()
	var sn string
	for rows.Next() {
		rows.Scan(&sn)
		fmt.Println(sn)
	}

	/*
		SELECT distinct  sn
		FROM `data_raw`
		WHERE create_time>0
		GROUP BY sn
		HAVING SUM(`value`)>0*/

	dataraw := &[]DataRaw{}
	//var sns []string
	db2.Model(&DataRaw{}).Where("create_time>? ", 0).
		Select("distinct  sn ").
		Group("sn").Having("SUM(`value`)>0").
		Find(dataraw) //Rows()

	rows.Scan(dataraw)
	fmt.Println(dataraw)
	/*for _,dr := range *dataraw {

		fmt.Println(dr)
	}*/

	//defer rows.Close()

	/*for rows.Next() {
		rows.Scan(&sn)
		fmt.Println(sn)
	}
*/

	type Sns struct {
		sns []string
	}

	rows, err = db2.Model(&DataRaw{}).Where("create_time>? ", 0).
		Select("distinct  sn ").
		Group("sn").Having("SUM(`value`)>0").
		Rows()


		sns1 := Sns{}
	rows.Scan(sns1)
	fmt.Println("____________")
	fmt.Println(sns1)
}
