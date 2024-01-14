package response

type Parser struct {
	Title     string        `json:"title"`
	CoverUrls string        `json:"coverUrls"`
	VideoUrls string        `json:"videoUrls"`
	IsVideo   bool          `json:"isVideo"`
	Pics      []string      `json:"pics"`
	Audio     []interface{} `json:"audio"`
	Formats   []struct {
		QualityNote string `json:"qualityNote"`
		Separate    int    `json:"separate"`
		Vext        string `json:"vext"`
		Video       string `json:"video"`
	} `json:"Formats"`
	Header     Header `json:"header"`
	IsOverseas bool   `json:"isOverseas"`
	EncodeUrl  string `json:"encodeUrl"`
}

type Header struct {
	UserAgent string `json:"User-Agent"`
}

type BatchParser struct {
	List struct {
		VideoList []struct {
			Type     string      `json:"type"`
			Pics     []string    `json:"pics"`
			Music    interface{} `json:"music"`
			Size     interface{} `json:"size"`
			Cover    string      `json:"cover"`
			Desc     string      `json:"desc"`
			PlayAddr string      `json:"playAddr,omitempty"`
		} `json:"videoList"`
		Paging string `json:"paging"`
	} `json:"list"`
	Direct bool `json:"direct"`
	Times  struct {
		Token   string `json:"token"`
		Type    int    `json:"type"`
		EndTime string `json:"end_time"`
		Times   int    `json:"times"`
		Message string `json:"message"`
	} `json:"times"`
	Header Header `json:"header"`
}

type ParserRecord struct {
	List []struct {
		Title      string `json:"title"`
		Url        string `json:"url"`
		CreateTime string `json:"create_time"`
	} `json:"list"`
}
