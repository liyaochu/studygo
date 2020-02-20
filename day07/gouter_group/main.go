package main
//路由
import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	shopGroup := r.Group("/shopping")
	{
		shopGroup.GET("/index",shopIndexHandler)
		shopGroup.GET("/car",shopCarHandler)
	}
	//利用的是HttpRouter路由原理前缀树
	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index",liveIndexHandler)
		liveGroup.GET("/home",liveHomeHandler)
	}
  r.Run()

}
