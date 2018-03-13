package main
import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"fmt"
)



func init() {

}
type FirmwareUpdateScheme struct {
	DeviceMac     string `gorm:"column:device_mac"`
	UpgradeScheme int    `gorm:"column:upgrade_scheme" description:"升级通道"`
}

var db1 *gorm.DB
func main()  {
	var err error
	//db, err = gorm.Open("mysql", "<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local")
	//db1, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	db1, err := gorm.Open("mysql", "root:123456@tcp(docker:3306)/speaker?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	db1.SingularTable(true)//全局禁止使用复数表
	db1.DB().SetMaxIdleConns(10)
	db1.DB().SetMaxOpenConns(100)

	var result []*FirmwareUpdateScheme
	mac := "6C-5A-B5-82-D2-22"
	if err := db1.Where("device_mac = ?", mac).Find(&result).Error; err != nil {
		log.Printf("Failed to get Device Upgrade Schme, with deviceId %s, with error %+v", mac, err)

	}else {
		for _,value :=range result{
			fmt.Printf("DeviceMac:%s,\nUpgradeScheme:%d\n",value.DeviceMac,value.UpgradeScheme)
		}
	}


}