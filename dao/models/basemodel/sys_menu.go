package models

type SysMenu struct {
	Id     int    `xorm:"not null pk autoincr comment('文档ID') INT(10)"`
	Title  string `xorm:"not null default '' comment('标题') VARCHAR(50)"`
	Pp     int    `xorm:"default 0 comment('商务后台权限') INT(11)"`
	Pid    int    `xorm:"not null default 0 comment('上级分类ID') index INT(10)"`
	Sort   int    `xorm:"not null default 0 comment('排序（同级有效）') INT(10)"`
	Url    string `xorm:"not null default '' comment('链接地址') CHAR(255)"`
	Hide   int    `xorm:"not null default 0 comment('是否隐藏') TINYINT(1)"`
	Tip    string `xorm:"not null default '' comment('提示') VARCHAR(255)"`
	Group  string `xorm:"default '' comment('分组') VARCHAR(50)"`
	IsDev  int    `xorm:"not null default 0 comment('是否仅开发者模式可见') TINYINT(1)"`
	Status int    `xorm:"not null default 0 comment('状态') index TINYINT(1)"`
}
