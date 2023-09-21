package webValidator

type PayDetailValidator struct {
	UserId int64  `json:"user_id" form:"user_id" binding:"required"`
	Price  string `json:"price" form:"price" binding:"required"`
	Order  string `json:"order" form:"order" binding:"required"`
	Vip    int64  `json:"vip" form:"vip" binding:"required"`
}
