package xinpianchang

import (
	"fmt"
	"github.com/iawia002/lux/test"
	"testing"
	"tutu-gin/lib/extractors"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "test 1",
			args: test.Args{
				URL:     "https://www.xinpianchang.com/a12204433",
				Title:   "超炫酷视觉系创意短片《遗留》",
				Quality: "720p",
				Size:    79595290,
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
