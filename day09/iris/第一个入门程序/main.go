package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	//1.创建app结构体对象返回的是application 指针
	app := iris.New()
	app.Get("/getRequest", func(context context.Context) {
		//处理get请求，请求的url为：/getRequest
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("请求路径"+path)

		//获取get请求的参数
		userName := context.URLParam("username")
		app.Logger().Info(userName)
		//返回html格式给前端
		context.HTML("<h1>"+userName+"</h1>")
	})
	
	//处理post请求
	app.Post("/login", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		c.WriteString("请求路径"+path)
		//处理post请求参数
		name := c.PostValue("name")
		password := c.PostValue("password")
		app.Logger().Info(name,password)
		c.HTML(name)
	})


	//处理json请求
	app.Post("/print", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		c.WriteString("请求路径"+path)
		//
		var person Person
		if 	err := c.ReadJSON(&person);err!=nil{
			panic(err.Error())
		}
       c.Writef("Recieve: %#+v\n",person)
	})

	//handle处理方式
	app.Handle("GET","/userinfo", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
	})

	app.Handle("POST","/find", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
	})

	//正则表达式
	app.Get("/find/{date}/{city}", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		date:=c.Params().Get("date")
		city:=c.Params().Get("city")
		c.WriteString(date+""+city)
	})

	//路由组请求
	//xxx/user/find
	//xxx/user/select
	userParty := app.Party("/user", func(c context.Context) {
		c.Next()
	})
	userParty.Get("/find", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		c.HTML("用户注册")
	})
	userParty.Get("/select", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		c.HTML("用户查询")
	})

	//端口监听
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))





}

type Person struct {
	Name string
	Password string
}