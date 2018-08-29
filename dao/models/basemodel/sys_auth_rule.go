package models

type SysAuthRule struct {
	Id        int    `xorm:"not null pk autoincr comment('规则id,自增主键') MEDIUMINT(8)"`
	Module    string `xorm:"not null comment('规则所属module') index(module) VARCHAR(20)"`
	Type      int    `xorm:"not null default 1 comment('1-url;2-主菜单') index(module) TINYINT(2)"`
	Name      string `xorm:"not null default '' comment('规则唯一英文标识') CHAR(80)"`
	Title     string `xorm:"not null default '' comment('规则中文描述') CHAR(20)"`
	Status    int    `xorm:"not null default 1 comment('是否有效(0:无效,1:有效)') index(module) TINYINT(1)"`
	Condition string `xorm:"not null default '' comment('规则附加条件') VARCHAR(300)"`
}
