package common

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Result struct {
	Code    int
	Message string
	Data    interface{}
}

func SendResponse(code int, message string, data interface{}) (result Result) {
	result.Code = code
	result.Message = message
	result.Data = data
	return
}

func CheckPhone(phone string) bool {
	result, _ := regexp.MatchString(`^(1[3|5|7|8|9][0-9]\d{4,8})$`, phone)
	if !result {
		return false
	}
	return true
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var builder strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&builder, "%d", numeric[ rand.Intn(r) ])
	}
	return builder.String()
}

func SendCaptchaMsg(phone string, captcha string)  {
	//msg := fmt.Sprintf("【腾讯科技】%v (短信验证码，15分钟内有效)", captcha)
	//send(phone, msg)
}
