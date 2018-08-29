package models

type TabAgent struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	OrderNumber    string  `xorm:"comment('订单号') VARCHAR(30)"`
	PayOrderNumber string  `xorm:"not null comment('支付订单号 ') VARCHAR(30)"`
	GameId         int     `xorm:"comment('游戏ID') INT(11)"`
	GameAppid      string  `xorm:"comment('游戏APPID') VARCHAR(32)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	PromoteId      int     `xorm:"default 0 comment('推广员ID') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广员账号') VARCHAR(30)"`
	UserId         int     `xorm:"comment('用户ID') INT(11)"`
	UserAccount    string  `xorm:"comment('玩家账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"comment('用户昵称') VARCHAR(30)"`
	Amount         float64 `xorm:"default 0.00 comment('支付金额') DOUBLE(10,2)"`
	RealAmount     float64 `xorm:"default 0.00 comment('实际金额') DOUBLE(10,2)"`
	PayWay         int     `xorm:"comment('支付方式 1支付宝 2微信') TINYINT(2)"`
	PayStatus      int     `xorm:"default 0 comment('支付状态') TINYINT(2)"`
	PayType        int     `xorm:"comment('类型') TINYINT(2)"`
	CreateTime     int     `xorm:"comment('时间') INT(11)"`
	Zhekou         int     `xorm:"not null default 0 comment('折扣比例') INT(11)"`
}
