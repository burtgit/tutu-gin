package parserDto

type ParserResultDto struct {
	Title     string
	CoverUrls string
	VideoUrls string
	IsVideo   bool
	Pics      []string
	Audio     []string
	Formats   []string
}
