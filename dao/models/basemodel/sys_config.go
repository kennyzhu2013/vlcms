package models

type SysConfig struct {
	Id         int    `xorm:"not null pk autoincr comment('配置ID') INT(10)"`
	Name       string `xorm:"not null default '' comment('配置名称') unique VARCHAR(30)"`
	Type       int    `xorm:"not null default 0 comment('配置类型') index TINYINT(3)"`
	Title      string `xorm:"not null default '' comment('配置说明') VARCHAR(50)"`
	Category   int    `xorm:"not null comment('配置分类') TINYINT(3)"`
	Group      int    `xorm:"not null default 0 comment('配置分组') index TINYINT(3)"`
	Extra      string `xorm:"not null default '' comment('配置值') VARCHAR(255)"`
	Remark     string `xorm:"not null default '' comment('配置说明') VARCHAR(100)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status     int    `xorm:"not null default 0 comment('状态') TINYINT(4)"`
	Value      string `xorm:"comment('配置值') TEXT"`
	Sort       int    `xorm:"not null default 0 comment('排序') SMALLINT(3)"`
}
