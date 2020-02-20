package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
func bookListHandler(c *gin.Context){
   //连接数据库
   //查数据
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":err,
		})
		return
	}
	//返回前端页面
	c.HTML(http.StatusOK,"book/book_list.html",gin.H{
		"code":0,
		"data":bookList,
	})
}
func newBookHandler(c *gin.Context){
  //给用户返回一个添加书籍的处理函数
  c.HTML(http.StatusOK,"book/new_book.html",nil)

}

func createBookHandler(c *gin.Context){
	title := c.PostForm("title")
	price := c.PostForm("price")
	fmt.Println(title,price)
	priceDa, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			
			"msg":"无效的价格参数",
		})
		return
	}
	err = insertBook(title, priceDa)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{

			"msg":"插入数据失败",
		})
		return
	}
	//插入数据成功
	c.Redirect(http.StatusMovedPermanently,"/book/list")
}

func main() {
	//程序启动就加载数据库
	err := initDB()
	if err!=nil{
		return
	}
	r := gin.Default()
    r.LoadHTMLGlob("/Users/liyaochu/mygo/src/studygo/day07/bsm/templates/**/*")


	r.GET("/book/list",bookListHandler)
	r.GET("/book/new",newBookHandler)
	r.POST("/book/create",createBookHandler)
	r.Run()
}
