package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(context *gin.Context){
	context.HTML(http.StatusOK,"login.html",gin.H{
     "msg":"嘿嘿嘿",
	})
}
func indexHandler(context *gin.Context){
	context.HTML(http.StatusOK,"index.html",gin.H{
		"msg":"哈哈哈",
	})
}
func helloHandler(context *gin.Context){
	type People struct {
		Name string `json:"username"`
		Age int
	}
	people := People{
		Name: "liyaochu",
		Age:  18,
	}
	context.JSON(http.StatusOK,people)
}
func main() {
	//创建一个默认的路由引擎
	r:=gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("./templates/*")

	r.GET("/login", loginHandler)

	r.GET("/index", indexHandler)
	r.GET("/hello", helloHandler)

	r.Run()
}
