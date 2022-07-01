package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutu-gin/core/exception"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, v := range c.Errors {
			err := v.Err
			if apiError, ok := err.(*exception.ApiException); ok {
				c.JSON(http.StatusOK, gin.H{
					"Code": apiError.Code,
					"Msg":  apiError.Error(),
					"Data": apiError.Data,
				})
			} else if domainError, ok := err.(*exception.DomainException); ok {
				c.JSON(http.StatusOK, gin.H{
					"Code": domainError.Code,
					"Msg":  domainError.Error(),
					"Data": domainError.Data,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"Code": 500,
					"Msg":  "未知错误",
					"Data": err.Error(),
				})
			}

			return
		}
	}
}
