package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code int    `json:"-"`
	Msg  string `json:"msg"`
}

const (
	DefaultErrCode = 400
)

func (err *Error) Error() string {
	return err.Msg
}

func (err *Error) Resp(ctx *gin.Context, args ...interface{}) {
	code := DefaultErrCode
	if getCode := err.Code; getCode > 0 {
		code = getCode
	}
	if len(args) == 0 {
		ctx.JSON(code, err)
	}
	getCode, ok := args[0].(int)
	if !ok || getCode <= 0 {
		ctx.JSON(code, err)
	}
	ctx.JSON(getCode, err)
}

var (
	ReqInvalidErr = &Error{
		Code: http.StatusForbidden,
		Msg:  "Req Invalid",
	}
)
