package recordDomain

type Record struct {
	Id         int
	UserId     int64
	Title      string
	Url        string
	Ip         string
	CreateTime string
	UpdateTime string
}

func (r *Record) SetId(id int) *Record {
	r.Id = id
	return r
}

func (r *Record) SetUserId(userId int64) *Record {
	r.UserId = userId
	return r
}

func (r *Record) SetTitle(title string) *Record {
	r.Title = title
	return r
}

func (r *Record) SetUrl(url string) *Record {
	r.Url = url
	return r
}

func (r *Record) SetIp(ip string) *Record {
	r.Ip = ip
	return r
}

func (r *Record) SetCreateTime(createTime string) *Record {
	r.CreateTime = createTime
	return r
}

func (r *Record) SetUpdateTime(updateTime string) *Record {
	r.UpdateTime = updateTime
	return r
}

func NewRecord() *Record {
	return &Record{}
}
