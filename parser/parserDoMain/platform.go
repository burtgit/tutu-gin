package parserDoMain

// 这是一个值对象

type Platform struct {
	Code   string
	Name   string
	Domain []string
}

func NewPlatform(code string, name string, domain []string) *Platform {
	return &Platform{
		Code:   code,
		Name:   name,
		Domain: domain,
	}
}
