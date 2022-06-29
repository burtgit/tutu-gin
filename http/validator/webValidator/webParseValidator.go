package webValidator

type WebParseValidator struct {
	PageUrl string `json:"pageUrl"  binding:"required"`
}
