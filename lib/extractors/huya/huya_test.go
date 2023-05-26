package huya

import (
	"testing"

	"github.com/iawia002/lux/extractors"
	"github.com/iawia002/lux/test"
)

func TestHuya(t *testing.T) {
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:   "https://m.v.huya.com/play/872047087.html",
				Title: "12.28 集梦薛小谦【封号斗罗】直播名场面",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New().Extract(tt.args.URL, extractors.Options{})
		})
	}
}
