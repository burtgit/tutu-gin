package initialize

import (
	"tutu-gin/lib/extractors"
	"tutu-gin/lib/extractors/qq"
	"tutu-gin/lib/extractors/zhihu"
)

func InitParse() {
	extractors.Register("qq", qq.New())
	extractors.Register("zhihu", zhihu.New())
}
