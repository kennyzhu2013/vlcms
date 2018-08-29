package models

type TabRepairRecord struct {
	Id             int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	PayOrderNumber string `xorm:"comment('支付订单号') VARCHAR(30)"`
	UserId         int    `xorm:"not null comment(' 用户ID') INT(11)"`
	UserAccount    string `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname   string `xorm:"comment('用户昵称') VARCHAR(30)"`
	GameId         int    `xorm:"comment('游戏id') INT(11)"`
	GameAppid      string `xorm:"comment('游戏appid') VARCHAR(32)"`
	GameName       string `xorm:"comment('游戏名称') VARCHAR(30)"`
	OpId           int    `xorm:"not null comment('操作人id') INT(11)"`
	OpNickname     string `xorm:"comment('操作人昵称') VARCHAR(30)"`
	CreateTime     int    `xorm:"comment('添加时间') INT(11)"`
	Type           int    `xorm:"default 0 comment('区分绑定平台币    1为绑定平台币补单') INT(11)"`
}
