package sdkwebmiddleware

import (
	"github.com/StarfishProgram/starfish-go-sdk/sdk"
	"github.com/StarfishProgram/starfish-go-sdk/sdkcodes"
	"github.com/gin-gonic/gin"
)

func RequestParam[T any](call func(*gin.Context, *T)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var p T
		err := ctx.ShouldBind(&p)
		if err != nil {
			sdk.Assert(false, sdkcodes.RequestParamInvalid.WithMsgf("%s", err.Error()))
		}
		call(ctx, &p)
		ctx.Next()
	}
}
