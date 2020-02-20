package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"studygo/day09/imooc-product/web/controller"
)

func main() {
	app:=iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("/Users/liyaochu/mygo/src/studygo/day10/imooc-product/web/views/movie",".html"))

	//注册控制机器
	mvc.New(app.Party("/hello")).Handle(new(controller.MovieController))
	app.Run(iris.Addr(":8080"))
}
