package webValidator

type PayCheckValidator struct {
	OrderId int64 `json:"order_id" form:"order_id" binding:"required"`
	PayType int64 `json:"pay_type" form:"pay_type" binding:"required"`
}
