package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
)

func ValidateHeader() func(*gin.Context) {
	return func(c *gin.Context) {
		// Sử dụng recovery để handle panic (crash)
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternalServer(err))
				}

				panic(r)
			}
		}()
		c.Next()
	}
}
