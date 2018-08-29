package models

type SysAuthExtend struct {
	GroupId  int `xorm:"not null pk comment('用户id') unique(group_extend_type) index MEDIUMINT(10)"`
	ExtendId int `xorm:"not null pk comment('扩展表中数据的id') unique(group_extend_type) index MEDIUMINT(8)"`
	Type     int `xorm:"not null pk comment('扩展类型标识 1:栏目分类权限;2:模型权限') unique(group_extend_type) TINYINT(1)"`
}
