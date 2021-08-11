package util

import "github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
