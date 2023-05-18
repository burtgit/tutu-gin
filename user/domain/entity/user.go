package entity

import "time"

type User struct {
	ID          int64     `gorm:"column:id" json:"id"`
	Mobile      string    `gorm:"column:mobile" json:"mobile"`
	Username    string    `gorm:"column:username" json:"username"`
	Avatar      string    `gorm:"column:avatar" json:"avatar"`
	Openid      string    `gorm:"column:openid" json:"openid"`
	Token       string    `gorm:"column:token" json:"token"`
	Status      int64     `gorm:"column:status" json:"status"`
	EndTime     int64     `gorm:"column:end_time" json:"end_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
	Message     string    `gorm:"column:message" json:"message"` //  消息
	Name        int64     `gorm:"column:name" json:"name"`       //  可使用次数
	Times       int64     `gorm:"column:times" json:"times"`     //  可使用次数
	ApplyTimes  int64     `gorm:"column:apply_times" json:"apply_times"`
	FormId      string    `gorm:"column:form_id" json:"form_id"` //  微信消息通知id
	ReadTime    int64     `gorm:"column:read_time" json:"read_time"`
	Platform    string    `gorm:"column:platform" json:"platform"`
	LoginTimes  int64     `gorm:"column:login_times" json:"login_times"` //  登录次数
	Type        int64     `gorm:"column:type" json:"type"`               //  用户类型，普通用户0，考拉团员1，考拉团长2
	SitesType   int64     `gorm:"column:sites_type" json:"sites_type"`
	Tid         int64     `gorm:"column:tid" json:"tid"`                   //  父级
	Ttid        int64     `gorm:"column:ttid" json:"ttid"`                 //  爷级
	Master      int64     `gorm:"column:master" json:"master"`             //  团长
	Tmaster     int64     `gorm:"column:tmaster" json:"tmaster"`           //  关联团长
	InviteId    int64     `gorm:"column:invite_id" json:"invite_id"`       //  最初邀请人
	Money       float64   `gorm:"column:money" json:"money"`               //  余额
	Earn        float64   `gorm:"column:earn" json:"earn"`                 //  已赚
	BatchTime   int64     `gorm:"column:batch_time" json:"batch_time"`     //  批量解析使用时间
	Message2    string    `gorm:"column:message2" json:"message2"`         //  批量解析说明
	Times2      int64     `gorm:"column:times2" json:"times2"`             //  批量解析可用次数
	BatchApplys int64     `gorm:"column:batch_applys" json:"batch_applys"` //  批量解析使用次数
	Iphonenum   string    `gorm:"column:iphonenum" json:"iphonenum"`       //  手机设备号
}

func (u User) TableName() string {
	return "video_user"
}
