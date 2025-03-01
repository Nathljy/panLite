package api

import (
	"panLite/resp"
	"panLite/service"

	"github.com/gin-gonic/gin"
)

func InitAuth(r *gin.RouterGroup) {
	r.POST("/login", login)
	r.POST("/logout", logout)
}

func login(ctx *gin.Context) {
	service := service.LoginService{}
	if err := ctx.ShouldBind(&service); err != nil {
		resp.ReqInvalidErr.Resp(ctx)
		return
	}
	if Err := service.PlainTextLogin(); Err != nil {
		Err.Resp(ctx)
		return
	}
	resp.Success(ctx)
}

func logout(ctx *gin.Context) {

}
