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

func AddToRoomHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddToRoomReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetMessageSrv()

		l.AddMember(ctx.Request.Context(), &req)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
		//	return
		//}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}

func RemoveToRoomHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.RemoveToRoomReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetMessageSrv()

		l.RemoveMember(ctx.Request.Context(), &req)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
		//	return
		//}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}
