package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func init() {

}

type RptDeviceTotalDaily struct {
	ProductCode string `gorm:"column:product_code;primary_key" json:"product_code"`
	StateDate   int    `gorm:"column:stat_date;primary_key" json:"stat_date"`
	ActiveCount int    `gorm:"column:active_count;" json:"active_count"`
	TotalCount  int    `gorm:"column:total_count;" json:"total_count"`
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`
}
type DeviceNum struct {
	StatDate    string `json:"stat_date"`
	ActiveCount int    `json:"active_count"`
	TotalCount  int    `json:"total_count"`
}
type ActPlayDuration struct {
	StatDate           int `json:"stat_date"`
	DurationLessThan30 int `json:"duration_less_than_30"`
	Duration30To60     int `json:"duration_30_to_60"`
	DurationMoreThan60 int `json:"duration_more_than_60"`
}

func main() {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker_data?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.SingularTable(true) //全局禁止使用复数表
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d := &[]DeviceNum{}

	/*	rows, err:=db.Model(&RptDeviceTotalDaily{}).Select("stat_date,active_count ,total_count").
			Where("stat_date >?",20180102).Rows()

			if err !=nil{
				fmt.Println(err)
				return
			}
		rows.Scan(d)*/
	//row.Scan(&actSn)
	/*defer rows.Close()
	temp:=DeviceNum{}
	for rows.Next() {

		rows.Scan(&temp)
		fmt.Println("temp:", temp)

	}*/

	db.Raw("SELECT stat_date,active_count ,total_count FROM `rpt_device_total_daily`  WHERE stat_date >?",
		20180102).Scan(d)
	fmt.Println("d:", d)

	resp := &[]ActPlayDuration{}

	db.Raw("SELECT stat_date, "+
		"duration_less_than_30 AS durationlessthan30,"+
		"duration_30_to_60 AS duration_30_to_60,"+
		"duration_more_than_60 AS duration_more_than_60 "+
		"FROM `rpt_device_play_duration_daily` "+
		" WHERE stat_date >?",
		20180102).Scan(resp)

	fmt.Println("d:", resp)

}
