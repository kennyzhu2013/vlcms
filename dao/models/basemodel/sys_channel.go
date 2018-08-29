package models

type SysChannel struct {
	Id         int    `xorm:"not null pk autoincr comment('频道ID') INT(10)"`
	Pid        int    `xorm:"not null default 0 comment('上级频道ID') index INT(10)"`
	Title      string `xorm:"not null comment('频道标题') CHAR(30)"`
	Url        string `xorm:"not null comment('频道连接') CHAR(100)"`
	Sort       int    `xorm:"not null default 0 comment('导航排序') INT(10)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status     int    `xorm:"not null default 0 comment('状态') TINYINT(4)"`
	Target     int    `xorm:"not null default 0 comment('新窗口打开') TINYINT(2)"`
}
