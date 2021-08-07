package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	gintemplate "github.com/foolin/gin-template"
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

func main() {
	router := gin.Default()

	// config server
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&httpServer, &http2Server)

	// new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views/fontend",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/ad"},
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.Static("/public", "./public")

	router.GET("/", func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			// Push is supported
			options := &http.PushOptions{
				Header: http.Header{
					"Accept-Encoding": ctx.Request.Header["Accept-Encoding"],
				},
			}
			err := pusher.Push("/public/img/ic-logo.png", options)
			if err != nil {
				log.Fatal(err)
			}
			data := Quiz{
				ID:       1,
				Question: "Có ai yêu tôi không?",
				IsOne:    false,
				Answer:   []string{"Có", "Không", "Có hoặc không", "Vinh yêu Cường", "Thịnh yêu Hiếu"},
				Result:   1,
			}
			gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
				"title": "Quiz",
				"model": data,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "not support server push",
			})
		}

	})
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	log.Fatal(httpServer.ListenAndServeTLS("./server.crt", "./server.key"))
}
