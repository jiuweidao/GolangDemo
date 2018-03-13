package main
import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"fmt"
)


var db *gorm.DB

func init() {

}
type People struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
//	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main()  {
	var err error
	//db, err = gorm.Open("mysql", "<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	//123456@tcp(docker:3306)/speaker?charset=utf8mb4&parseTime=True
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	/*
		创建表
	 */
	if !db.HasTable(&People{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
			CreateTable(&People{}).Error; err != nil {
				panic(err)
		}
	}


	/*
		插入
	 */
	/*like := &People{
		Ip:        "34r34",
		Ua:        "ua",
		Title:     "title",
		CreatedAt: time.Now(),
	}

	if err := db.Create(like).Error; err != nil {
		fmt.Print(err)
	}

*/
	/*
		查询
	 */
	var count int
	lsterrs1 := make([]*People,0)

	if err := db.Where("Ip = ?", "df").Find(&lsterrs1).Error; err != nil {
		fmt.Printf("Failed to get  %s, with error %+v", 1, err)

	}else {
		fmt.Println("get____________-")
		for _,value :=range lsterrs1  {
			fmt.Println(value)
		}
	}


	if err := db.Model(&People{}).Where(&People{Ip: "hello"}).Count(&count).Error;err != nil {
		fmt.Print(err)
	}else {
		lsterrs := make([]*People,0)
		db.Find(&lsterrs)
		fmt.Println("-----------------")
		fmt.Println("db.Find(&lsterrs)")
		for _,value :=range  lsterrs{
			fmt.Println(value.ID)
			fmt.Println(value.Title)
		}
		fmt.Println(lsterrs)

	}

	like1 := People{ID:1}
	if err := db.First(&like1,like1);err!=nil{
		fmt.Println("db.First(&like1,like1)")
		fmt.Println(like1)
	}
//db.Table("users").Select("COALESCE(age,?)", 42).Rows()


	//a:=db.Where("name = ?", "jinzhu").Find(&errs{}).Count()

	fmt.Println("count",count)


	/*
		修改
	 */
	/*db.Model(&err).Update("Ip", "hello")
	db.Model(&err).Updates(People{Ip: "hello", ID: 18})
	db.Model(&err).Updates(People{Ip: "", ID: 0, Title: "erfe"}) // nothing update*/

	fmt.Println("------------")
	//errs := db.Exec("replace into firmware_update_scheme(device_mac, upgrade_scheme) values(?, ?)", scheme.DeviceMac, scheme.UpgradeScheme).Error
	errs := db.Exec("replace into peoples(id, ip,ua,title) values(?,?,?,?)",  30,"df","ds","dfd").Error


	if errs == nil{


	}
}