package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app).Handle(new(UserController))

	//路由组处理
	mvc.Configure(app.Party("/user"), func(con *mvc.Application) {
		con.Handle(new(UserController))
	})
	app.Run(iris.Addr(":8080"))
}
//默认匹配
//type : get
func (c *UserController) Get() string {
	iris.New().Logger().Info("get请求")
	return "helloword"
}
//type : post
func (c *UserController) Post() string {
	iris.New().Logger().Info("Post请求")
	return "helloword"
}

/**type : get
   url:localhost:8080/info
   需要在Get后在Info
 */
func (c *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("get请求")
	return mvc.Response{
		Object:map[string]interface{}{
			"code":1,
			"msg":"请求成功",
		},
	}
}
/**type : post
  url:localhost:8080/login
*/
func (c *UserController) PostLogin() string {
	iris.New().Logger().Info("Post请求")
	return "helloword"
}

/**
  人为匹配
 */
func (c *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET","/query","UserInfo")
}
func(c *UserController) UserInfo() mvc.Result{
	iris.New().Logger().Info("userinfo query")
	return mvc.Response{}
}
type UserController struct {
}
