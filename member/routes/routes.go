package routes

import (
	"github.com/gin-gonic/gin"
	"member/controllers"
)

func SetupRoutes(r *gin.Engine) {

	// 註冊會員相關路由
	r.POST("/user", controllers.CreateMember)
	r.GET("/user/:id", controllers.GetMember)
	r.PUT("/user/:id", controllers.UpdateMember)
	r.DELETE("/user/:id", controllers.DeleteMember)

}
