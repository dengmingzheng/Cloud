package main

import (
	"Cloud/controller"
	"Cloud/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config, err := tools.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	registerRouter(app)

	app.Run(config.AppHost + ":" + config.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
