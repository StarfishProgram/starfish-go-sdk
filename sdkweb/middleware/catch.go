package sdkwebmiddleware

import (
	"net/http"

	"github.com/StarfishProgram/starfish-go-sdk/sdk"
	"github.com/StarfishProgram/starfish-go-sdk/sdkcodes"
	"github.com/StarfishProgram/starfish-go-sdk/sdklog"
	"github.com/gin-gonic/gin"
)

func Catch(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		if code, ok := err.(sdkcodes.Code); ok {
			sdklog.AddCallerSkip(3).Warn(code)
			var msg string
			if code.Code() == sdkcodes.Internal.Code() {
				msg = sdkcodes.Internal.Msg()
			} else {
				msg = code.Msg()
			}
			ctx.JSON(http.StatusOK, sdk.AnyMap{
				"code": code.Code(),
				"msg":  msg,
				"i18n": code.I18n(),
				"data": nil,
			})
			ctx.Abort()
			return
		}
		sdklog.AddCallerSkip(2).Error(err)
		ctx.JSON(http.StatusOK, sdk.AnyMap{
			"code": sdkcodes.Internal.Code(),
			"msg":  sdkcodes.Internal.Msg(),
			"i18n": sdkcodes.Internal.I18n(),
			"data": nil,
		})
		ctx.Abort()
	}()
	ctx.Next()
}
