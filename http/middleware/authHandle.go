package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"tutu-gin/core/api"
	"tutu-gin/parser/parserApplicaition/parserDto"
	"tutu-gin/user"
	"tutu-gin/user/infrastructure/dto"
)

func AuthHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("tokens")
		if err != nil || len(token) <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"Code": 403,
				"Msg":  "未登录",
				"Data": nil,
			})
			c.Abort()
			return
		}

		userService := user.NewUserService(&dto.ConfigDto{
			Token: token,
		})

		userInfo := userService.Detail()

		if userInfo.Id <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"Code": 403,
				"Msg":  "未登录",
				"Data": nil,
			})
			c.Abort()
			return
		}

		if userInfo.Status == 0 {
			c.JSON(http.StatusOK, api.ApiSuccessResponse(parserDto.ParserResultDto{
				Title:     "请关注微信公众号『大拿解析』进行视频去水印",
				CoverUrls: "https://img.alicdn.com/bao/uploaded/i1/O1CN01MVPP541Pyp5ooFMjU_!!2-rate.png",
				VideoUrls: "http://gslb.miaopai.com/stream/JHeJmOV8InV6MISxrZAD8xsps4Mga0-IxSTERw__.mp4?vend=miaopai&ssig=6c6f8ac141464a5c6244c0cf7b9751a5&time_stamp=1592927773813&mpflag=32&unique_id=1592924158668318",
				IsVideo:   true,
			}))
			c.Abort()
			return
		}

		c.Set("userId", userInfo.Id)
		c.Next()
	}
}
