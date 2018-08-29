package models

type SysAttribute struct {
	Id           int    `xorm:"not null pk autoincr INT(10)"`
	Name         string `xorm:"not null default '' comment('字段名') VARCHAR(30)"`
	Title        string `xorm:"not null default '' comment('字段注释') VARCHAR(100)"`
	Field        string `xorm:"not null default '' comment('字段定义') VARCHAR(100)"`
	Type         string `xorm:"not null default '' comment('数据类型') VARCHAR(20)"`
	Value        string `xorm:"not null default '' comment('字段默认值') VARCHAR(100)"`
	Remark       string `xorm:"not null default '' comment('备注') VARCHAR(100)"`
	IsShow       int    `xorm:"not null default 1 comment('是否显示') TINYINT(1)"`
	Extra        string `xorm:"not null default '' comment('参数') VARCHAR(255)"`
	ModelId      int    `xorm:"not null default 0 comment('模型id') index INT(10)"`
	IsMust       int    `xorm:"not null default 0 comment('是否必填') TINYINT(1)"`
	Status       int    `xorm:"not null default 0 comment('状态') TINYINT(2)"`
	UpdateTime   int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
	CreateTime   int    `xorm:"not null default 0 comment('创建时间') INT(11)"`
	ValidateRule string `xorm:"not null default '' VARCHAR(255)"`
	ValidateTime int    `xorm:"not null default 0 TINYINT(1)"`
	ErrorInfo    string `xorm:"not null default '' VARCHAR(100)"`
	ValidateType string `xorm:"not null default '' VARCHAR(25)"`
	AutoRule     string `xorm:"not null default '' VARCHAR(100)"`
	AutoTime     int    `xorm:"not null default 0 TINYINT(1)"`
	AutoType     string `xorm:"not null default '' VARCHAR(25)"`
}
