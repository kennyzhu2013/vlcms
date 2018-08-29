package models

type TabBindSpend struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	UserId         int     `xorm:"not null comment('用户id') INT(11)"`
	UserAccount    string  `xorm:"not null comment('用户账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"not null comment('用户昵称') VARCHAR(30)"`
	GameId         int     `xorm:"not null comment('游戏id') INT(11)"`
	GameAppid      string  `xorm:"not null comment('游戏appid') VARCHAR(30)"`
	GameName       string  `xorm:"not null comment('游戏名称') VARCHAR(30)"`
	OrderNumber    string  `xorm:"not null comment('订单号') VARCHAR(30)"`
	PayOrderNumber string  `xorm:"not null comment('商户订单号') VARCHAR(30)"`
	PropsName      string  `xorm:"not null comment('道具名称') VARCHAR(30)"`
	PayAmount      float64 `xorm:"not null default 0.00 comment('金额') DOUBLE(10,2)"`
	PayTime        int     `xorm:"not null comment('支付时间') INT(11)"`
	PayStatus      int     `xorm:"not null default 0 comment('支付状态') TINYINT(2)"`
	PayGameStatus  int     `xorm:"not null default 0 comment('游戏支付状态') TINYINT(2)"`
	Extend         string  `xorm:"comment('通知游戏方扩展(一般是游戏方透传)') VARCHAR(32)"`
	PayWay         int     `xorm:"not null default 0 comment('支付方式') TINYINT(2)"`
	PaySource      int     `xorm:"not null default 0 comment('支付来源') TINYINT(2)"`
	SpendIp        string  `xorm:"not null comment('支付ip') VARCHAR(16)"`
	PromoteAccount string  `xorm:"not null comment('推广员账号') VARCHAR(30)"`
	PromoteId      int     `xorm:"not null default 0 comment('推广员id 0自然注册') INT(11)"`
}
