package middleware

import (
	"github.com/gin-gonic/gin"
)

func StatusHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ip := c.ClientIP()
		// fmt.Println("本地解析客户端地址：" + ip)
		c.Next()
		//if ip == "118.25.251.62" || ip == "119.29.94.101" || ip == "101.36.127.200" {
		//	fmt.Println("检测到异常ip地址")
		//	c.JSON(http.StatusOK, api.ApiSuccessResponse(parserDto.ParserResultDto{
		//		Title:     "请关注微信公众号『考拉解析』进行视频去水印",
		//		CoverUrls: "https://img.alicdn.com/bao/uploaded/i1/O1CN01MVPP541Pyp5ooFMjU_!!2-rate.png",
		//		VideoUrls: "http://gslb.miaopai.com/stream/JHeJmOV8InV6MISxrZAD8xsps4Mga0-IxSTERw__.mp4?vend=miaopai&ssig=6c6f8ac141464a5c6244c0cf7b9751a5&time_stamp=1592927773813&mpflag=32&unique_id=1592924158668318",
		//		IsVideo:   true,
		//	}))
		//	c.Abort()
		//	return
		//} else {
		//	c.Next()
		//}
	}
}
