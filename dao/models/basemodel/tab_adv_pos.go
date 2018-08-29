package models

type TabAdvPos struct {
	Id     int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name   string `xorm:"not null VARCHAR(50)"`
	Title  string `xorm:"not null default '' comment('广告位置名称') CHAR(80)"`
	Module string `xorm:"not null comment('所在模块 模块/控制器/方法') VARCHAR(100)"`
	Type   int    `xorm:"not null default 1 comment('广告位类型 
1.单图
2.多图
3.文字链接
4.代码') INT(11)"`
	Status  int    `xorm:"not null default 1 comment('状态（0：禁用，1：正常）') TINYINT(2)"`
	Data    string `xorm:"not null comment('额外的数据') VARCHAR(500)"`
	Width   string `xorm:"not null default '' comment('广告位置宽度') CHAR(20)"`
	Height  string `xorm:"not null default '' comment('广告位置高度') CHAR(20)"`
	Margin  string `xorm:"not null comment('边缘') VARCHAR(50)"`
	Padding string `xorm:"not null comment('留白') VARCHAR(50)"`
	Theme   string `xorm:"not null default 'all' comment('适用主题，默认为all，通用') VARCHAR(50)"`
}
