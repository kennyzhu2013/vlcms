package models

type SysUcenterSetting struct {
	Id    int    `xorm:"not null pk autoincr comment('设置ID') INT(10)"`
	Type  int    `xorm:"not null default 0 comment('配置类型（1-用户配置）') TINYINT(3)"`
	Value string `xorm:"not null comment('配置数据') TEXT"`
}
