package models

type SysAuthGroupAccess struct {
	Uid     int `xorm:"not null pk comment('用户id') index unique(uid_group_id) INT(10)"`
	GroupId int `xorm:"not null pk comment('用户组id') index unique(uid_group_id) MEDIUMINT(8)"`
	Houtai  int `xorm:"default 0 comment('所属后台') INT(11)"`
}
