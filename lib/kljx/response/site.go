package response

type Site struct {
	Name         string `json:"name"`
	Wechat       string `json:"wechat"`
	Platform     string `json:"platform"`
	PlatformName string `json:"platform_name"`
	PlatformLogo string `json:"platform_logo"`
	MpQrcode     string `json:"mp_qrcode"`
	Desc         string `json:"desc"`
	BatchDesc    string `json:"batch_desc"`
	Notice       []struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Url     string `json:"url"`
	} `json:"notice"`
	BatchNotice []struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Url     string `json:"url"`
	} `json:"batch_notice"`
	AdPic []struct {
		Img string `json:"img"`
		Url string `json:"url"`
	} `json:"ad_pic"`
	AdText []struct {
		Content string `json:"content"`
		Color   string `json:"color"`
		Url     string `json:"url"`
	} `json:"ad_text"`
	AdBatchPic []struct {
		Img string `json:"img"`
		Url string `json:"url"`
	} `json:"ad_batch_pic"`
	AdBatchText []struct {
		Content string `json:"content"`
		Color   string `json:"color"`
		Url     string `json:"url"`
	} `json:"ad_batch_text"`
}
