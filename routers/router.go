package router

import (
	"github.com/cuongtop4598/QuizWithGo/QuizApp/middlewares/jwt"
	"github.com/cuongtop4598/QuizWithGo/QuizApp/routers/api"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/public", "./public")

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
}
