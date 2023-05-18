package webValidator

type WebParseValidator struct {
	PageUrl string `json:"pageUrl"  binding:"required"`
}

type UserQrcodeValidator struct {
	Ticket string `form:"ticket"  binding:"required"`
}
