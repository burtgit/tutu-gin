package exception

type DomainException struct {
	Msg  string
	Code int
	Data interface{}
}

func (d DomainException) Error() string {
	return d.Msg
}

func DomainError(code int, msg string) *DomainException {
	return &DomainException{
		Code: code,
		Msg:  msg,
	}
}

var (
	DOMAIN_NOT_FOUND = DomainError(404, "资源未找到")                 // 资源未找到
	PARSE_FAIL       = DomainError(500, "解析失败，请联系微信kaolajiexi2") // 解析失败
	JSON_PARSE_FAIL  = DomainError(201, "json解析失败")
	DB_ACTION_FAIL   = DomainError(500, "数据库操作失败")
)
