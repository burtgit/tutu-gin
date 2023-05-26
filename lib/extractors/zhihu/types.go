package zhihu

// minimum field
type video struct {
	PlayList struct {
		FHD resolution `json:"FHD"`
		HD  resolution `json:"HD"`
		SD  resolution `json:"SD"`
		LD  resolution `json:"LD"`
	} `json:"playlist"`
	PlayListV2 struct {
		FHD resolution `json:"FHD"`
		HD  resolution `json:"HD"`
		SD  resolution `json:"SD"`
		LD  resolution `json:"LD"`
	} `json:"playlist_v2"`
	CoverUrl string `json:"cover_url"`
}

// minimum field
type resolution struct {
	Size    int64  `json:"size"`
	Format  string `json:"format"`
	PlayURL string `json:"play_url"`
	Width   int    `json:"width"`
}
