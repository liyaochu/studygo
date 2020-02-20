package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func liveIndexHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"shopping/index",
	})
}
func liveHomeHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"shopping/car",
	})
}
