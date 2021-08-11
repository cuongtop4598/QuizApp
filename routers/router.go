package routers

import (
	v1 "github.com/cuongtop4598/QuizWithGo/QuizApp/routers/api/v1"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/public", "./public")

	// new template engine
	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "views/fontend",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{"partials/ad"},
		Funcs:        nil,
		DisableCache: true,
	})

	user := r.Group("/users")
	{
		user.GET("/", user.LoginPage)
		user.POST("/", user.Login)
		user.GET("/:id", user.GetUser)
		user.PUT("/:id", user.UpdateUser)
		user.DELETE("/:id", user.DeleteUser)
	}

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	// {
	apiv1.GET("/quizzes/:id", v1.GetQuizByID)
	apiv1.GET("/quizzes/{id}", v1.GetQuizByID)
	apiv1.GET("/quizzes/", v1.GetQuizByID)
	apiv1.GET("/quizzes/save/", v1.SaveAnswer)
	apiv1.GET("/quizzes/result/", v1.GetResult)

	// }

	return r
}
