package initialize

import (
	"tutu-gin/lib/extractors"
	"tutu-gin/lib/extractors/douyin"
	"tutu-gin/lib/extractors/facebook"
	"tutu-gin/lib/extractors/instagram"
	"tutu-gin/lib/extractors/qq"
	"tutu-gin/lib/extractors/twitter"
	"tutu-gin/lib/extractors/youtube"
	"tutu-gin/lib/extractors/zhihu"
)

func InitParse() {
	extractors.Register("qq", qq.New())
	extractors.Register("zhihu", zhihu.New())
	extractors.Register("douyin", douyin.New())
	extractors.Register("iesdouyin", douyin.New())
	extractors.Register("instagram", instagram.New())
	extractors.Register("facebook", facebook.New())
	extractors.Register("youtube", youtube.New())
	extractors.Register("twitter", twitter.New())
}
