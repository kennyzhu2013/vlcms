package models

type TabAdv struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Title      string `xorm:"not null default '' comment('广告名称') CHAR(80)"`
	PosId      int    `xorm:"not null comment('广告位置') INT(11)"`
	Data       string `xorm:"not null comment('图片地址') TEXT"`
	ClickCount int    `xorm:"not null comment('点击量') INT(11)"`
	Url        string `xorm:"not null comment('链接地址') VARCHAR(500)"`
	Sort       int    `xorm:"not null default 0 comment('排序') INT(3)"`
	Status     int    `xorm:"not null default 1 comment('状态（0：禁用，1：正常）') TINYINT(2)"`
	CreateTime int    `xorm:"not null default 0 comment('开始时间') INT(11)"`
	StartTime  int    `xorm:"INT(11)"`
	EndTime    int    `xorm:"default 0 comment('结束时间') INT(11)"`
	Target     string `xorm:"default '_blank' VARCHAR(20)"`
}
