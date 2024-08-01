package v1

import (
	"chat/service"
	"chat/types"
	api "common/api/v1"
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SendMessageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SendMessageReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetMessageSrv()

		resp, err := l.SendMessage(&req)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func TestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, "resp"))
	}
}
