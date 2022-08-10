package services

import (
	"Cloud/tools"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(phone string) bool {
	//1-生成验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//2-调用阿里云sdk,完成发送
	sms := tools.GetConfig().Sms
	client, error := dysmsapi.NewClientWithAccessKey(sms.RegionId, sms.AppKey, sms.AppSecret)

	if error != nil {
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = sms.SignName
	request.TemplateCode = sms.TemplateCode
	request.PhoneNumbers = phone

	par, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)

	if err != nil {
		return false
	}

	//3-接收返回结果，判断发送状态
	if response.Code == "OK" {
		return true
	}
	return false
}
