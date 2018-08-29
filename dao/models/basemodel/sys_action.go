package models

type SysAction struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name       string `xorm:"not null default '' comment('行为唯一标识') CHAR(30)"`
	Title      string `xorm:"not null default '' comment('行为说明') CHAR(80)"`
	Remark     string `xorm:"not null default '' comment('行为描述') CHAR(140)"`
	Rule       string `xorm:"comment('行为规则') TEXT"`
	Log        string `xorm:"comment('日志规则') TEXT"`
	Type       int    `xorm:"not null default 1 comment('类型') TINYINT(2)"`
	Status     int    `xorm:"not null default 0 comment('状态') TINYINT(2)"`
	UpdateTime int    `xorm:"not null default 0 comment('修改时间') INT(11)"`
}
