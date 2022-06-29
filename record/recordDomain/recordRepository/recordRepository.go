package recordRepository

import (
	"tutu-gin/record/recordDomain"
	"tutu-gin/record/recordInfrastructure/recordDataobject"
	"tutu-gin/record/recordInfrastructure/recordMapper"
)

// RecordRepository 仓储层
type RecordRepository struct {
	mapper *recordMapper.RecordDbMapper
}

func (r RecordRepository) Insert(record *recordDomain.Record) (*recordDomain.Record, error) {
	recordDao := &recordDataobject.RecordDao{
		UserId: record.UserId,
		Title:  record.Title,
		Url:    record.Url,
		Ip:     record.Ip,
	}

	var err error
	record.Id, err = r.mapper.Insert(recordDao)
	return record, err
}

func NewRecordRepository() *RecordRepository {
	return &RecordRepository{
		mapper: new(recordMapper.RecordDbMapper),
	}
}
