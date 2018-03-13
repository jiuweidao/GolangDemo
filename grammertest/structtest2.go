package main

import (
	"fmt"
	"github.com/pborman/uuid"
	"time"
)
/*
uuid是Universally Unique Identifier的缩写，即通用唯一识别码。

uuid的目的是让分布式系统中的所有元素，
都能有唯一的辨识资讯，而不需要透过中央控制端来做辨识资讯的指定。
如此一来，每个人都可以建立不与其它人冲突的 uuid。
 */
type AdminUser struct {
	Id         string `gorm:"column:uuid;primary_key" json:"id"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
	Name       string `gorm:"column:name" json:"name"`
	Avatar     string `gorm:"column:avatar" json:"avatar"`
	Gender     int    `gorm:"column:gender" json:"gender"`
	Status     int    `gorm:"column:status" json:"status"`
	Role       string `gorm:"column:role" json:"role"`
	LockTime   int64  `gorm:"column:lock_time" json:"lock_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time"`
}
func (u *AdminUser) BeforeCreate() error {
	u.Id = uuid.New()
	u.Status = 1
	u.CreateTime = time.Now().Unix()
	return nil
}
func main()  {

	a:=AdminUser{}
	error:=a.BeforeCreate()
	if error ==nil {
		fmt.Println("a.id="+a.Id)
	}





}

