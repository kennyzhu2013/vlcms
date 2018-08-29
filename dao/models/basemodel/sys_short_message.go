package models

type SysShortMessage struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Phone      string `xorm:"not null comment('电话号码') VARCHAR(20)"`
	SendStatus string `xorm:"not null comment('短信发送状态') VARCHAR(10)"`
	SendTime   string `xorm:"not null comment('发送时间') VARCHAR(15)"`
	Smsid      string `xorm:"not null comment('发送短信唯一标识') VARCHAR(40)"`
	CreateTime int    `xorm:"not null comment('记录时间') INT(10)"`
	Pid        string `xorm:"not null comment('渠道id') VARCHAR(40)"`
	Status     int    `xorm:"not null comment('状态') TINYINT(2)"`
	Ratio      int    `xorm:"not null comment('比率') INT(10)"`
}
