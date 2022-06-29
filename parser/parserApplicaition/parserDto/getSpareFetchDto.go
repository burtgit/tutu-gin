package parserDto

import "tutu-gin/parser/parserDoMain"

type GetSpareFetchDto struct {
	PageUrl  string
	Platform *parserDoMain.Platform
}
