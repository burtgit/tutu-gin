package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"tutu-gin/lib/kljx"
	"tutu-gin/lib/kljx/response"
)

type Index struct{}

func (i Index) Index(c *gin.Context) {
	if strings.Contains(c.Request.Host, "danaqsy.com") {

		title := "抖音"
		example := "http://v.douyin.com/22TnGD/"
		desc := "抖音 视频"
		path := c.Request.RequestURI
		if c.Request.RequestURI == "/xiaohongshu.html" {
			title = "小红书"
			example = "http://xhslink.com/hlVaar"
			desc = "小红书 视频/图片"
		} else if c.Request.RequestURI == "/kuaishou.html" {
			title = "快手"
			example = "https://v.kuaishou.com/ruZDM9"
			desc = "快手 视频"
		} else if c.Request.RequestURI == "/xigua.html" {
			title = "西瓜"
			example = "https://www.ixigua.com/7223945085413818880"
			desc = "西瓜 视频"
		} else if c.Request.RequestURI == "/weibo.html" {
			title = "微博"
			example = "http://m.weibo.cn/status/4885315180037721?"
			desc = "微博 视频"
		} else if c.Request.RequestURI == "/bilibili.html" {
			title = "哔哩哔哩"
			example = "https://www.bilibili.com/video/BV1WS4y1H7yG"
			desc = "哔哩哔哩 视频"
		} else if c.Request.RequestURI == "/instagram.html" {
			title = "Instagram"
			example = "https://www.instagram.com/p/Ce1wWSQrSua/"
			desc = "ins图片/视频"
		} else if c.Request.RequestURI == "/youtube.html" {
			title = "YouTube"
			example = "https://www.youtube.com/watch?v=UTKS3UKUCUs/"
			desc = "YouTube 视频"
		} else if c.Request.RequestURI == "/tiktok.html" {
			title = "Tiktok"
			example = "https://www.tiktok.com/t/ZPRvFSKCn/"
			desc = "Tiktok 视频"
		}
		c.HTML(http.StatusOK, "dana_index.html", gin.H{
			"title":   title,
			"example": example,
			"desc":    desc,
			"path":    path,
		})
	} else if strings.Contains(c.Request.Host, "tkqsy.com") {
		c.HTML(http.StatusOK, "pipi_index.html", gin.H{})
	} else {

		var token string
		tokens, err := c.Request.Cookie("tokens")

		if err == nil {
			token = tokens.Value
		}

		_, userDetail := kljx.NewClient[response.User]().Apply(kljx.UserInfo, map[string]string{
			"token": token,
		})

		var isVip bool
		if userDetail.EndTime > time.Now().Unix() {
			isVip = true
		} else if userDetail.Times > 0 {
			isVip = true
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"isVip": isVip,
		})
	}
}

func NewIndex() *Index {
	return &Index{}
}
