package main

import (
	_ "CMS/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {

	mode := beego.BConfig.RunMode
	if mode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//	注册mysql
	user := beego.AppConfig.String(mode +"::mysqluser")
	host := beego.AppConfig.String(mode + "::mysqlurls")
	port := beego.AppConfig.String(mode + "::mysqlport")
	datasename := beego.AppConfig.String(mode + "::mysqldb")
	password := beego.AppConfig.String(mode + "::mysqlpass")
	dataSource := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + datasename + "?charset=utf8"
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default","mysql",dataSource, maxIdle,maxConn)
	// 跨域设置
	//origin := beego.Controller{}
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins:[]string{"http://localhost:8090"},
		AllowMethods:[]string{"POST","PUT","DELETE"},
	}))
	beego.Run()
}
