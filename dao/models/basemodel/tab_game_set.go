package models

type TabGameSet struct {
	Id             int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId         int    `xorm:"not null comment('游戏ID') INT(11)"`
	LoginNotifyUrl string `xorm:"comment('游戏登陆通知地址') VARCHAR(255)"`
	PayNotifyUrl   string `xorm:"comment('游戏支付通知地址') VARCHAR(255)"`
	GameRoleUrl    string `xorm:"comment('游戏角色获取地址') VARCHAR(255)"`
	GameGiftUrl    string `xorm:"comment('游戏礼包领取地址') VARCHAR(255)"`
	GameKey        string `xorm:"comment('游戏key') VARCHAR(32)"`
	AccessKey      string `xorm:"comment('访问秘钥') VARCHAR(32)"`
	AgentId        string `xorm:"comment('代理id(合作方标示)') VARCHAR(32)"`
	GamePayAppid   string `xorm:"comment('游戏支付APPID') VARCHAR(32)"`
	QqGroup        string `xorm:"comment('qq群') VARCHAR(32)"`
}
