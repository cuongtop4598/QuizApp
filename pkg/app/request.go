package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
