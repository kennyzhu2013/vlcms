package models

type SysUcenterApp struct {
	Id         int    `xorm:"not null pk autoincr comment('应用ID') INT(10)"`
	Title      string `xorm:"not null comment('应用名称') VARCHAR(30)"`
	Url        string `xorm:"not null comment('应用URL') VARCHAR(100)"`
	Ip         string `xorm:"not null default '' comment('应用IP') CHAR(15)"`
	AuthKey    string `xorm:"not null default '' comment('加密KEY') VARCHAR(100)"`
	SysLogin   int    `xorm:"not null default 0 comment('同步登陆') TINYINT(1)"`
	AllowIp    string `xorm:"not null default '' comment('允许访问的IP') VARCHAR(255)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status     int    `xorm:"not null default 0 comment('应用状态') index TINYINT(4)"`
}
