package recordDataobject

// RecordDao 数据转换层(专门用来对接db，es，redis等底层储存)
type RecordDao struct {
	Id         int64
	UserId     int64
	Title      string
	Url        string
	Ip         string
	CreateTime string
	UpdateTime string
}
