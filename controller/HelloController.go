package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

//路由
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

//解析路由
func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"message": context.FullPath(),
	})
}
