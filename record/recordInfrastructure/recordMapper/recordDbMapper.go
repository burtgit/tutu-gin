package recordMapper

import (
	"time"

	"github.com/juju/errors"

	"tutu-gin/core/exception"
	"tutu-gin/core/global"
	"tutu-gin/record/recordInfrastructure/recordDataobject"
)

type RecordDbMapper struct {
	Id         int
	UserId     int64
	Title      string
	Url        string
	Ip         string
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func (m RecordDbMapper) TableName() string {
	return global.SERVICE_CONFIG.DataBase.TablePrefix + "records"
}

func (m *RecordDbMapper) Insert(dao *recordDataobject.RecordDao) (id int, err error) {
	m.UserId = dao.UserId
	m.Title = dao.Title
	m.Url = dao.Url
	m.Ip = dao.Ip

	_, err = global.DB.InsertOne(m)
	if err != nil {
		return 0, errors.Annotate(err, exception.DOMAIN_DB_ACTION_FAIL)
	}

	return m.Id, nil
}
