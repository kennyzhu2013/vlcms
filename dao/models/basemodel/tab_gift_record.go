package models

type TabGiftRecord struct {
	Id           int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId       int    `xorm:"comment('游戏id') INT(11)"`
	GameName     string `xorm:"comment('游戏名称') VARCHAR(30)"`
	ServerId     int    `xorm:"comment('区服') INT(11)"`
	ServerName   string `xorm:"comment('区服名称') VARCHAR(30)"`
	GiftId       int    `xorm:"comment('礼包id') INT(11)"`
	GiftName     string `xorm:"comment('礼包名称') VARCHAR(30)"`
	Status       int    `xorm:"default 1 comment('状态(0:未使用,1:已使用)') TINYINT(2)"`
	Novice       string `xorm:"comment('激活码') VARCHAR(30)"`
	UserId       int    `xorm:"comment('用户id') INT(11)"`
	UserAccount  string `xorm:"comment('用户昵称') VARCHAR(30)"`
	UserNickname string `xorm:"comment('用户昵称') VARCHAR(30)"`
	CreateTime   int    `xorm:"comment('创建时间') INT(11)"`
}
