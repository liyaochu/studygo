package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shopIndexHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"shopping/index",
	})
}
func shopCarHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"shopping/car",
	})
}
