package recordApplication

import (
	"tutu-gin/record/recordApplication/recordDto"
	"tutu-gin/record/recordDomain"
	"tutu-gin/record/recordDomain/recordRepository"
)

type RecordService struct {
}

// Create 创建记录
func (r RecordService) Create(dto *recordDto.CreateDto) (*recordDomain.Record, error) {

	rp := recordRepository.NewRecordRepository()

	record := recordDomain.NewRecord().
		SetTitle(dto.Title).
		SetUrl(dto.Url).
		SetIp(dto.Ip).
		SetUserId(dto.UserId)
	var err error
	record, err = rp.Insert(record)

	return record, err
}
