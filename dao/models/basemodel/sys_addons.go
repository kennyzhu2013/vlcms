package models

type SysAddons struct {
	Id           int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Name         string `xorm:"not null comment('插件名或标识') VARCHAR(40)"`
	Title        string `xorm:"not null default '' comment('中文名') VARCHAR(20)"`
	Description  string `xorm:"comment('插件描述') TEXT"`
	Status       int    `xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Config       string `xorm:"comment('配置') TEXT"`
	Author       string `xorm:"default '' comment('作者') VARCHAR(40)"`
	Version      string `xorm:"default '' comment('版本号') VARCHAR(20)"`
	CreateTime   int    `xorm:"not null default 0 comment('安装时间') INT(10)"`
	HasAdminlist int    `xorm:"not null default 0 comment('是否有后台列表') TINYINT(1)"`
}
