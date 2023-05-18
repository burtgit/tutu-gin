package webValidator

type WebParseValidator struct {
	PageUrl string `json:"pageUrl"  binding:"required"`
}

type UserQrcodeValidator struct {
	Ticket string `json:"ticket"  binding:"required"`
}
