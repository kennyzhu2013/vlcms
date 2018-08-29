package models

type TabSpend struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	UserId         int     `xorm:"not null comment(' 用户ID') INT(11)"`
	UserAccount    string  `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"comment('用户昵称') VARCHAR(30)"`
	GameId         int     `xorm:"comment('游戏id') INT(11)"`
	GameAppid      string  `xorm:"comment('游戏appid') VARCHAR(32)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	ServerId       int     `xorm:"comment('区服id') INT(11)"`
	ServerName     string  `xorm:"comment('区服名称') VARCHAR(30)"`
	PromoteId      int     `xorm:"comment('推广员id') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广员账号') VARCHAR(30)"`
	OrderNumber    string  `xorm:"comment('订单号') VARCHAR(30)"`
	PayOrderNumber string  `xorm:"comment('支付订单号') VARCHAR(30)"`
	PropsName      string  `xorm:"comment('道具名称') VARCHAR(30)"`
	PayAmount      float64 `xorm:"comment('支付金额') DOUBLE(10,2)"`
	PayTime        int     `xorm:"comment('支付时间') INT(11)"`
	PayStatus      int     `xorm:"comment('支付状态') TINYINT(2)"`
	PayGameStatus  int     `xorm:"comment('游戏支付状态') TINYINT(2)"`
	Extend         string  `xorm:"comment('通知游戏方扩展(一般是游戏方透传)') VARCHAR(32)"`
	PayWay         int     `xorm:"comment('支付方式') TINYINT(2)"`
	SpendIp        string  `xorm:"comment('支付IP') VARCHAR(20)"`
	IsCheck        int     `xorm:"default 1 comment('是否对账  1参与 2不参与 3参与(已对账) 4不参与(已对账)') INT(11)"`
	SelleStatus    int     `xorm:"not null default 0 comment('CP结算 0未结算1 结算') INT(11)"`
	SelleRatio     float64 `xorm:"default 0.00 comment('cp分成比例') DOUBLE(5,2)"`
	SubStatus      int     `xorm:"not null default 0 comment('子渠道结算状态') INT(11)"`
	SelleTime      string  `xorm:"comment('cp结算时间') VARCHAR(30)"`
}
