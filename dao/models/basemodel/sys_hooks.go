package models

type SysHooks struct {
	Id          int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Name        string `xorm:"not null default '' comment('钩子名称') unique VARCHAR(40)"`
	Description string `xorm:"comment('描述') TEXT"`
	Type        int    `xorm:"not null default 1 comment('类型') TINYINT(1)"`
	UpdateTime  int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Addons      string `xorm:"not null default '' comment('钩子挂载的插件 '，'分割') VARCHAR(255)"`
	Status      int    `xorm:"not null default 1 TINYINT(1)"`
}
