package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func queryStringHandler(c *gin.Context){
	//获取 query string 参数
	//name := c.Query("name")   //查不到默认就是空
	name := c.DefaultQuery("name", "李耀楚")
	city := c.DefaultQuery("city", "上海")//查不到,给定一个值
     c.JSON(http.StatusOK,gin.H{
     	"name":name,
     	"city":city,
	 })
}

func formHandler(c *gin.Context){
	name := c.PostForm("name")
	city := c.DefaultPostForm("city", "上海")
	c.JSON(http.StatusOK,gin.H{
		"name":name,
		"city":city,
	})
}
func pathHandler(c *gin.Context){
	param := c.Param("action")
	c.JSON(http.StatusOK,gin.H{
		"action":param,
	})
}
func main() {
	r := gin.Default()
	r.GET("/query_string",queryStringHandler)
	//form 表单传参
	r.POST("/form",formHandler)
	//url参数  /book/list
	r.POST("/book/:action ",pathHandler)
   r.Run()
}
