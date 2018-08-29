package models

type TabRebateList struct {
	Id             int     `xorm:"not null pk autoincr INT(11)"`
	PayOrderNumber string  `xorm:"comment('订单号') VARCHAR(30)"`
	GameId         int     `xorm:"comment('游戏id') INT(11)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	UserId         int     `xorm:"comment('用户id') INT(11)"`
	UserName       string  `xorm:"comment('用户名') VARCHAR(30)"`
	PayAmount      float64 `xorm:"default 0.00 comment('充值金额') DOUBLE(10,2)"`
	Ratio          float64 `xorm:"default 0.00 comment('返利比例') DOUBLE(5,2)"`
	RatioAmount    float64 `xorm:"comment('返利平台币') DOUBLE(10,2)"`
	PromoteId      int     `xorm:"default 0 comment('推广id') INT(11)"`
	PromoteName    string  `xorm:"comment('推广员姓名') VARCHAR(30)"`
	CreateTime     int     `xorm:"comment('创建时间') INT(11)"`
}
