package response

type Version struct {
	Version     string `json:"version"`
	DownloadUrl string `json:"download_url"`
	Force       bool   `json:"force"`
	Desc        string `json:"desc"`
	NeedUpdate  bool   `json:"need_update"`
}
