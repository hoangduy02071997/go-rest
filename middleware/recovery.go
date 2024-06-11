package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		//log.Println("Recovery...")

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

func RecoveryGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					//panic(err) // --> un comment for show log stack trace
					return
				}
				appErr := common.ErrInternalServer(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
