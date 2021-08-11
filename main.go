package main

import (
	"fmt"
	"net/http"

	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/logging"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/setting"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/pkg/util"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/routers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
)

type Quiz struct {
	ID       int
	Question string
	IsOne    bool // if "true" can be chose one answer, otherwise it can be chose more answer
	Answer   []string
	Result   int
}

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
}

func main() {

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	// config server
	server := http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&server, &http2Server)

	//log.Fatal(server.ListenAndServeTLS("./server.crt", "./server.key"))
	server.ListenAndServe()
}
