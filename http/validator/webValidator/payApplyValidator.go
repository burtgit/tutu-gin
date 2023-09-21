package webValidator

type PayApplyValidator struct {
	PayMethod int64 `json:"pay_method"   binding:"required" form:"pay_method"`
	Vip       int64 `json:"vip"   binding:"required" form:"vip"`
}
