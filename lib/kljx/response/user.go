package response

import "time"

type User struct {
	Id          int64     `json:"id"`
	Mobile      string    `json:"mobile"`
	Username    string    `json:"username"`
	Avatar      string    `json:"avatar"`
	Type        int       `json:"type"`
	Openid      string    `json:"openid"`
	Token       string    `json:"token"`
	Status      int       `json:"status"`
	Times       int64     `json:"times"`
	ApplyTimes  int       `json:"apply_times"`
	EndTime     int64     `json:"end_time"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	FormId      string    `json:"form_id"`
	ReadTime    int       `json:"read_time"`
	Message     string    `json:"message"`
	Platform    string    `json:"platform"`
	LoginTimes  int       `json:"login_times"`
	SitesType   int       `json:"sites_type"`
	Tid         int       `json:"tid"`
	Ttid        int       `json:"ttid"`
	Master      int       `json:"master"`
	Tmaster     int       `json:"tmaster"`
	InviteId    int       `json:"invite_id"`
	Money       string    `json:"money"`
	Earn        string    `json:"earn"`
	BatchTime   int       `json:"batch_time"`
	Message2    string    `json:"message2"`
	Times2      int       `json:"times2"`
	BatchApplys int       `json:"batch_applys"`
	IphoneNum   string    `json:"iphoneNum"`
	InviteUrl   string    `json:"invite_url"`
}
