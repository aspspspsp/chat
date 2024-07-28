package v1

import (
	api "common/api/v1"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/gin-gonic/gin"
	"log"
	"member/services"
	"member/types"
	"net/http"
)

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.RegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := services.GetMemberSrv()
		resp, err := l.Register(ctx.Request.Context(), &req)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

// LoginHandler 用户登陆接口
func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.LoginReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, api.ErrorResponse(ctx, err))
			return
		}

		l := services.GetMemberSrv()
		resp, err := l.Login(ctx.Request.Context(), &req)
		if err != nil {
			log.Fatalf(err.Error())
			ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}
