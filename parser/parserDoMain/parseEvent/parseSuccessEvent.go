package parseEvent

import (
	"tutu-gin/parser/parserApplicaition/parserDto"
	"tutu-gin/record/recordApplication"
	"tutu-gin/record/recordApplication/recordDto"
)

type ParseSuccessEvent struct {
	ParserResult *parserDto.ParserResultDto
	UserId       int64
	Ip           string
	Url          string
}

func (p ParseSuccessEvent) PublishEvent() {
	ra := &recordApplication.RecordService{}
	ra.Create(&recordDto.CreateDto{
		UserId: p.UserId,
		Ip:     p.Ip,
		Title:  p.ParserResult.Title,
		Url:    p.Url,
	})
	// log.Println("插入记录")
	// log.Println(err)
}
