package routes

import (
	"github.com/gin-gonic/gin"
	api "member/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		member := v1.Group("member")
		{
			member.POST("register", api.RegisterHandler())
			member.POST("login", api.LoginHandler())
		}
	}

	return r
}
