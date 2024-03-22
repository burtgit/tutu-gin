package youtube

import (
	"fmt"
	"github.com/iawia002/lux/test"
	"regexp"
	"testing"
	"tutu-gin/lib/extractors"
)

func TestYoutube(t *testing.T) {
	tests := []struct {
		name     string
		args     test.Args
		playlist bool
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:   "https://youtu.be/afqvp6O-pJA?si=YgdxpG3Gzo96CHAE",
				Title: "Multifandom Mashup 2017",
			},
		},
		{
			name: "signature test",
			args: test.Args{
				URL:   "https://youtu.be/ZWj-1xQK0jY?list=RDCMUCY4DrFDwxH7VaDHCsXzDE_w",
				Title: "Halo Infinite - E3 2019 - Discover Hope",
			},
		},
		{
			name: "playlist test",
			args: test.Args{
				URL:   "https://www.youtube.com/watch?v=Lt2pwLxJxgA&list=PLIYAO-qLriEtYm7UcXPH3SOJxgqjwRrIw",
				Title: "papi酱 - 你有酱婶儿的朋友吗？",
			},
			playlist: true,
		},
		{
			name: "url_encoded_fmt_stream_map test",
			args: test.Args{
				URL:   "https://youtu.be/1X8JVo8Zofg",
				Title: "QNAP Customer Story | Scorptec",
			},
		},
		{
			name: "stream 404 test 1",
			args: test.Args{
				URL:   "https://m.youtube.com/watch?v=nyDk-rgP0tI",
				Title: "FreeFileSync: Mirror Synchronization",
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

func TestClient(t *testing.T) {
	validURL := `https?://www\.tiktok\.com/(?:embed|@(?P<user_id>[\w\.-]+)?/video)/(?P<id>\d+)`
	re := regexp.MustCompile(validURL)

	match := re.FindStringSubmatch("https://www.tiktok.com/@ayami_nmb48/video/7343566694993923346")
	fmt.Println("Matched groups:")
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			fmt.Printf("%s: %s\n", name, match[i])
		}
	}
}
