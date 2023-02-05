package main

import (
	_ "beego-guess/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:yourpassword@tcp(127.0.0.1:3306)/gomysql?charset=utf8")
}

func main() {
	beego.Run()
}
