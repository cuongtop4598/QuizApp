package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type Quiz struct {
	Question string
	Answer   []string
	Result   int
}

func main() {
	tmpl := template.Must(template.ParseFiles("./public/home/index.html"))
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		data := Quiz{
			Question: "Có ai yêu tôi không?",
			Answer:   []string{"Có", "Không"},
			Result:   1,
		}
		tmpl.Execute(ctx.Writer, data)
	})
	engine.Run()
}
