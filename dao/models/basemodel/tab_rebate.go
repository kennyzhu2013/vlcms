package models

type TabRebate struct {
	Id         int     `xorm:"not null pk autoincr INT(11)"`
	GameId     int     `xorm:"comment('游戏id') INT(11)"`
	GameName   string  `xorm:"comment('游戏名称') VARCHAR(50)"`
	Money      int     `xorm:"default 0 comment('金额') INT(11)"`
	Ratio      float64 `xorm:"default 0.00 comment('返利比例') DOUBLE(5,2)"`
	Status     int     `xorm:"comment('状态(0关闭金额限制 1 开启金额限制)') INT(11)"`
	CreateTime int     `xorm:"comment('更新时间') INT(11)"`
}
