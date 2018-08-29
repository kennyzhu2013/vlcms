package models

type SysUrl struct {
	Id         int    `xorm:"not null pk autoincr comment('链接唯一标识') INT(11)"`
	Url        string `xorm:"not null default '' comment('链接地址') unique CHAR(255)"`
	Short      string `xorm:"not null default '' comment('短网址') CHAR(100)"`
	Status     int    `xorm:"not null default 2 comment('状态') TINYINT(2)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
