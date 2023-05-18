package recordDto

type CreateDto struct {
	UserId int64  `json:"userId,omitempty" form:"userId"`
	Ip     string `json:"ip,omitempty" form:"ip"`
	Title  string `json:"title,omitempty" form:"title"`
	Url    string `json:"url,omitempty" form:"url"`
}
