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

func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateRoomReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetRoomSrv()

		l.Create(ctx.Request.Context(), &req)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
		//	return
		//}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}

func DeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteRoomReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetRoomSrv()

		l.Delete(ctx.Request.Context(), &req)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
		//	return
		//}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}

func AddToRoomHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddToRoomReq
		if err := ctx.ShouldBind(&req); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
			return
		}

		l := service.GetRoomSrv()

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

		l := service.GetRoomSrv()

		l.RemoveMember(ctx.Request.Context(), &req)
		//if err != nil {
		//	log.Println(err)
		//	ctx.JSON(http.StatusOK, api.ErrorResponse(ctx, err))
		//	return
		//}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}
