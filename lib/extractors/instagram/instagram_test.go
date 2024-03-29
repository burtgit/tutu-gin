package instagram

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
			name: "video test",
			args: test.Args{
				URL:   "https://www.instagram.com/reel/CXbF8W_JW_n/?igsh=MWV0YTIzOXZrZDI2Ng==",
				Title: "Instagram BlIka1ZFCNr",
				Size:  3003662,
			},
		},
		{
			name: "image test",
			args: test.Args{
				URL:   "https://www.instagram.com/reel/CvAdOjno5LZ/?utm_source=ig_web_copy_link&igsh=MzRlODBiNWFlZA==",
				Title: "Instagram Bl5oVUyl9Yx",
				Size:  250596,
			},
		},
		{
			name: "image album test",
			args: test.Args{
				URL:   "https://www.instagram.com/reel/C4R3yCuO5UQ/?igsh=NTc4MTIwNjQ2YQ==",
				Title: "Instagram Bjyr-gxF4Rb",
				Size:  4599909,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := New().Extract(tt.args.URL, extractors.Options{})
			test.CheckError(t, err)
			fmt.Println(data)
			//test.Check(t, tt.args, data)
		})
	}
}

func TestEncode(t *testing.T) {
	fmt.Println(EncodeBaseN("C21eFeQRGHO"))
}
