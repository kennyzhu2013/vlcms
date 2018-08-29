package models

type SysUcenterAdmin struct {
	Id       int `xorm:"not null pk autoincr comment('管理员ID') INT(10)"`
	MemberId int `xorm:"not null default 0 comment('管理员用户ID') INT(10)"`
	Status   int `xorm:"not null default 0 comment('管理员状态') TINYINT(3)"`
}
