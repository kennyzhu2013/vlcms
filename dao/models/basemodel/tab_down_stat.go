package models

type TabDownStat struct {
	Id         int `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	PromoteId  int `xorm:"comment('推广员id') INT(11)"`
	GameId     int `xorm:"comment('游戏id') INT(11)"`
	Number     int `xorm:"default 1 comment('下载次数') INT(2)"`
	CreateTime int `xorm:"comment('时间') INT(11)"`
	Type       int `xorm:"comment('类型') INT(2)"`
}
