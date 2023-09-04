package webValidator

type WebParseValidator struct {
	PageUrl string `json:"pageUrl" form:"pageUrl"  binding:"required"`
}

type UserQrcodeValidator struct {
	Ticket string `form:"ticket"  binding:"required"`
}

type ApiParseValidator struct {
	PageUrl string `json:"pageUrl"   binding:"required" form:"pageUrl"`
	Token   string `json:"token"  binding:"required" form:"token"`
}
