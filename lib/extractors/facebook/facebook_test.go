package facebook

import (
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
			name: "normal test",
			args: test.Args{
				URL:     "https://www.facebook.com/reel/833442481868590?fs=e&s=TIeQ9V&mibextid=OCFGub",
				Title:   "Роман Грищук - Підтримка з Японії 🇯🇵 Гурт Yokohama Sisters 👏",
				Size:    1441128,
				Quality: "sd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, extractors.Options{})
			test.CheckError(t, err)
			t.Log(data)
		})
	}
}
