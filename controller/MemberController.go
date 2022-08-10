package controller

import (
	"Cloud/services"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

//路由
func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.SendSmsCode)
}

//发送验证码方法
func (mc *MemberController) SendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")

	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	memberService := services.MemberService{}

	isSend := memberService.SendCode(phone)

	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
	}

	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败",
	})
}
