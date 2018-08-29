package models

type SysAuthGroup struct {
	Id          int    `xorm:"not null pk autoincr comment('用户组id,自增主键') MEDIUMINT(8)"`
	Module      string `xorm:"not null default '' comment('用户组所属模块') VARCHAR(20)"`
	Type        int    `xorm:"not null default 0 comment('组类型') TINYINT(4)"`
	Title       string `xorm:"not null default '' comment('用户组中文名称') CHAR(20)"`
	Description string `xorm:"not null default '' comment('描述信息') VARCHAR(80)"`
	Status      int    `xorm:"not null default 1 comment('用户组状态：为1正常，为0禁用,-1为删除') TINYINT(1)"`
	Rules       string `xorm:"not null default '' comment('用户组拥有的规则id，多个规则 , 隔开') VARCHAR(950)"`
}
