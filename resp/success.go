package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResp struct {
	Code int         `json:"-"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (success *SuccessResp) Success(ctx *gin.Context) {
	ctx.JSON(success.Code, success)
}

func Success(ctx *gin.Context, args ...interface{}) {
	if len(args) == 0 {
		ctx.JSON(http.StatusOK, SuccessResp{Msg: "OK"})
		return
	}
	ctx.JSON(http.StatusOK, SuccessResp{
		Msg:  "OK",
		Data: args[0],
	})
}
