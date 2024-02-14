package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/rinha-backend-2024/internal/core/exception"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			switch err.Err.(type) {
			case *exception.NotFoundException:
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": err.Error(),
					"code":    http.StatusNotFound,
				})
			case *exception.ValidationException:
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": err.Error(),
					"code":    http.StatusUnprocessableEntity,
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "internal server error",
					"code":    http.StatusInternalServerError,
				})
			}
		}
	}
}
