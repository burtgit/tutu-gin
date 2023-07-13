package parserDto

type ParserResultDto struct {
	Title     string
	CoverUrls string
	VideoUrls string
	EncodeUrl string
	IsVideo   bool
	Pics      []string
	Audio     []string
	Formats   []ParseFormat
}

type ParseFormat struct {
	QualityNote string
	Separate    int
	Vext        string
	Video       string
}
