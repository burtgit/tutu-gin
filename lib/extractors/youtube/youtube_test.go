package youtube

import (
	"testing"
	"tutu-gin/lib/extractors"

	"github.com/iawia002/lux/test"
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
				URL:   "https://www.youtube.com/watch?v=Gnbch2osEeo",
				Title: "Multifandom Mashup 2017",
			},
		},
		{
			name: "signature test",
			args: test.Args{
				URL:   "https://www.youtube.com/watch?v=ZtgzKBrU1GY",
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
				URL:   "https://youtu.be/DNaOZovrSVo",
				Title: "QNAP Customer Story | Scorptec",
			},
		},
		{
			name: "stream 404 test 1",
			args: test.Args{
				URL:   "https://www.youtube.com/watch?v=MRJ8NnUXacY",
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
