package common

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(400, err)
			}
		}()
		ctx.Next()
	}
}
