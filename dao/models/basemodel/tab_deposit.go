package models

type TabDeposit struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	OrderNumber    string  `xorm:"comment('订单号') VARCHAR(30)"`
	PayOrderNumber string  `xorm:"comment('支付订单号') VARCHAR(30)"`
	UserId         int     `xorm:"comment('用户id') INT(11)"`
	UserAccount    string  `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"comment('用户昵称') VARCHAR(30)"`
	PromoteId      int     `xorm:"comment('推广员ID') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广账号') VARCHAR(30)"`
	PayAmount      float64 `xorm:"default 0.00 comment('充值金额') DOUBLE(10,2)"`
	PayStatus      int     `xorm:"default 0 comment('充值状态') TINYINT(2)"`
	PayWay         int     `xorm:"default 0 comment('支付方式') TINYINT(2)"`
	PaySource      int     `xorm:"default 0 comment('支付来源') TINYINT(2)"`
	PayIp          string  `xorm:"comment('充值IP') VARCHAR(20)"`
	CreateTime     int     `xorm:"comment('支付时间') INT(11)"`
}
