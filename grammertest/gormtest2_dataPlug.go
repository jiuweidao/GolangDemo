package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func init() {

}

type RptDevicePlugDaily struct {
	ProductCode    string `gorm:"column:product_code;primary_key" json:"product_code"`
	StatDate       int    `gorm:"column:stat_date;" json:"stat_date"`
	PlugTimes      int    `gorm:"column:plug_times;" json:"plug_times"`
	PlugOneFifth   int    `gorm:"column:plug_one_fifth;" json:"plug_one_fifth"`
	PlugTwoFifth   int    `gorm:"column:plug_two_fifth;" json:"plug_two_fifth"`
	PlugThreeFifth int    `gorm:"column:plug_three_fifth;" json:"plug_three_fifth"`
	PlugFourFifth  int    `gorm:"column:plug_four_fifth;" json:"plug_four_fifth"`
	PlugFiveFifth  int    `gorm:"column:total_count;" json:"plug_five_fifth"`
	CreateTime     int64  `gorm:"column:create_time" json:"create_time"`
}

type ActPlug struct {
	PlugTimes int `json:"plug_times"`
	ActPlugDistribution ActPlugDistribution `json:"a"`
}
type ActPlugDistribution struct {
	PlugOneFifth   int `json:"plug_one_fifth"`
	PlugTwoFifth   int `json:"plug_two_fifth"`
	PlugThreeFifth int `json:"plug_three_fifth"`
	PlugFourFifth  int `json:"plug_four_fifth"`
	PlugFiveFifth  int `json:"plug_five_fifth"`
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

	d := &ActPlug{}
	/*row := db.Table("rpt_device_plug_daily").
		Where("stat_date >?", 20180102).
		Select("SUM(plug_times), SUM(plug_one_fifth)," +
		"SUM(plug_three_fifth)," +
		"SUM(plug_four_fifth)," +
		"SUM(plug_five_fifth)").Row() // (*sql.Row)

	row.Scan(d)*/
	db.Raw("SELECT "+
		"SUM(plug_one_fifth) AS plug_one_fifth,"+
		"SUM(plug_two_fifth) AS plug_two_fifth ,"+
		"SUM(plug_three_fifth) AS plug_three_fifth,"+
		"SUM(plug_four_fifth) AS plug_four_fifth ,"+
		"SUM(plug_five_fifth) AS plug_five_fifth  "+
		"FROM `rpt_device_plug_daily`  WHERE stat_date >?",
		20180102).Scan(&d.ActPlugDistribution)
	fmt.Println("d:", d)

}
