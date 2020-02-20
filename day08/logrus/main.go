package main
//logrus 示例
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func  initLogrus()  {
	log.Formatter= &logrus.JSONFormatter{} //记录json格式日志
	file, err := os.Open("./xxx.log")
  if err!=nil{
  	fmt.Println(err)
	  return
  }
   log.Out=file
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	// Only log the warning severity or above.
	log.Level = logrus.InfoLevel
}
func main() {
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		log.WithFields(logrus.Fields{
			"animal": "walrus",
			"size":   10,
		}).Warn("A group of walrus emerges from the ocean")
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
