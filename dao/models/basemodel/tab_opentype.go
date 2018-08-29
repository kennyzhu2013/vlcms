package models

type TabOpentype struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	CreateTime int    `xorm:"not null comment('添加时间') INT(11)"`
	Status     int    `xorm:"not null comment('状态') INT(11)"`
	OpenName   string `xorm:"not null comment('开放类型名称') VARCHAR(30)"`
}
