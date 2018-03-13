package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm" //对应的beego/orm库 可以通过 go get来获取到本地GOPATH路径下
	"time"
)
//与数据库学生表映射的结构体
type studentinfo struct {
	Id          int    `pk:"auto"`
	Stuname     string `orm:"size(20)"`
	Stuidentify string `orm:"size(30)"`
	Stubirth    time.Time
	Stuclass    string `orm:"size(30)"`
	Stumajor    string `orm:"size(30)"`
}
//数据库连对象需要的信息
var (
	dbuser string = "root"
	dbpwd  string = "691214"
	dbname string = "gosql"
)
//初始化orm
func init() {
	conn := dbuser + ":" + dbpwd + "@/" + dbname + "?charset=utf8"//组合成连接串

	orm.RegisterModel(new(studentinfo))//注册表studentinfo 如果没有会自动创建
	orm.RegisterDriver("mysql", orm.DR_MySQL) //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", conn) //设置conn中的数据库为默认使用数据库
	orm.RunSyncdb("default", false, false)//后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
}

func main() {
	orm.Debug = true //true 打印数据库操作日志信息
	dbObj := orm.NewOrm()  //实例化数据库操作对象
	sql := fmt.Sprintf("insert into studentinfo(Id,Stuname, Stuidentify, Stubirth, Stuclass, Stumajor)"+
		" values(1, 'rjx','xxx319928xxx','%s','信管1班','信息管理与信息系统')", "1992-01-01 11:11:11")
	fmt.Println(sql)
	_, err := dbObj.Raw(sql).Exec()
	if err != nil {
		fmt.Println("插入数据至：t_studentInfo出错")
	}

	sql = fmt.Sprintf("insert into studentinfo(Id, Stuname, Stuidentify, Stubirth, Stuclass, Stumajor)"+
		" values(2, 'qcy','xxx319918xxx','%s','XXX','YYYYYY')", "1992-01-01 11:11:11")
	_, err = dbObj.Raw(sql).Exec()
	if err != nil {
		fmt.Println("插入数据至：t_studentInfo出错")
	}
	//更新数据
	sql = "update studentinfo set Stuname='qcym' where Id= 2"
	_, err = dbObj.Raw(sql).Exec()
	if err != nil {
		fmt.Println("更新t_studentInfo表出错")
	}
	//通过事务方式来进行数据插入
	err = dbObj.Begin()
	sql = fmt.Sprintf("insert into studentinfo(Id, Stuname, Stuidentify, Stubirth, Stuclass, Stumajor)"+
		" values(3, 'loe','xxx319918xxx','%s','zzzz','TTTT')", "1992-01-01 11:11:11")
	_, err = dbObj.Raw(sql).Exec()
	if err != nil {
		dbObj.Rollback()
		fmt.Println("插入t_studentInfo表出错,事务回滚")
	} else {
		dbObj.Commit()
		fmt.Println("插入t_studenInfo表成功,事务提交")
	}
	//查询数据库
	students := make([]studentinfo, 20)
	//var students []studentinfo
	sql = "select Id,Stuname,Stuidentify,Stubirth,Stuclass,Stumajor from studentinfo"
	fmt.Println(sql)
	num, er := dbObj.Raw(sql).QueryRows(&students)
	if er != nil {
		fmt.Println("查询学生信息出错")
	} else {
		fmt.Printf("从t_studenInfo表中共查询到记录:%d条\n", num)
		for index, _ := range students {
			fmt.Printf("第%d个学生个人信息：", index+1)
			fmt.Printf("姓名：%s,身份证号:%s，出生日期:%s,班级:%s,专业:%s", students[index].Stuname, students[index].Stuidentify, students[index].Stubirth, students[index].Stuclass, students[index].Stumajor)
		}
	}

	//


}