package values

import "time"

type Apply struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`
	Date         int64     `json:"date"`
	TotalTimes   int64     `json:"total_times"`
	SuccessTimes int64     `json:"success_times"`
	Platform     string    `json:"platform"`
	CreateTime   time.Time `xorm:"created"`
	UpdateTime   time.Time `xorm:"updated"`
}

func (a Apply) TableName() string {
	return "video_applies"
}
