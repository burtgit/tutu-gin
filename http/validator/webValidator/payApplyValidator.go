package webValidator

type PayApplyValidator struct {
	PayType int64 `json:"pay_type"   binding:"required" form:"pay_type"`
	Menu    int64 `json:"menu"   binding:"required" form:"menu"`
}
