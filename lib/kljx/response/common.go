package response

type ClientResponse interface {
	User | Parser | Site | Menu | PayApplyResult | BatchParser | ParserRecord | Version | any
}
type Response[T ClientResponse] struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
}
