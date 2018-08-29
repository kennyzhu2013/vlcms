package models

type SysUserdata struct {
	Uid      int `xorm:"not null pk comment('用户id') unique(uid) INT(10)"`
	Type     int `xorm:"not null pk comment('类型标识') unique(uid) TINYINT(3)"`
	TargetId int `xorm:"not null pk comment('目标id') unique(uid) INT(10)"`
}
