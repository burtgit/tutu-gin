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
				URL:     "https://www.facebook.com/story.php?story_fbid=pfbid0msSLuJwAYMBwTKQZR64hMYcVxhE6atT6knSLW3Q7oHZALnEkUKa9YAfvZ3qRCSNtl&id=100004758931302&mibex&paipv=0&eav=AfahVXSbnq6J6g9sSI-ZGoZD9CT9oE4Jq8JrWxlmz5hmFpW9iZZLs-OeaOaa8Bk54ag&_rdr",
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
