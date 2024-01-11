package zhihu

import (
	"fmt"
	"testing"
	"tutu-gin/lib/extractors"

	"github.com/iawia002/lux/test"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "video test",
			args: test.Args{
				URL:   "https://www.zhihu.com/zvideo/1447817011515056128",
				Title: `Cursor, GPT-4 驱动的强大代码编辑器 - 知乎`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, extractors.Options{})
			test.CheckError(t, err)
			fmt.Println(data)
			//test.Check(t, tt.args, data[0])
		})
	}
}
