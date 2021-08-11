package user

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	gintemplate.HTML(c, 200, "login", gin.H{})
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {

}
