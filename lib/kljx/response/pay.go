package response

type Menu struct {
	WechatPay bool `json:"wechatPay"`
	MenuList  []struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Price int    `json:"price"`
		Value string `json:"value"`
	} `json:"menuList"`
	BatchMenuList []struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Price int    `json:"price"`
		Value string `json:"value"`
	} `json:"batchMenuList"`
	Message []string `json:"message"`
}

type PayApplyResult struct {
	Money      string `json:"money"`
	OrderId    int    `json:"order_id"`
	PayType    string `json:"pay_type"`
	NeedRemark bool   `json:"need_remark"`
	UserId     int    `json:"user_id"`
	Qrcode     string `json:"qrcode"`
}
