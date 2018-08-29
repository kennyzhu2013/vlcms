package models

type TabProvide struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	OrderNumber    string  `xorm:"comment('订单号') VARCHAR(30)"`
	PayOrderNumber string  `xorm:"comment('商户订单号') VARCHAR(30)"`
	Cost           int     `xorm:"default 0 comment('是否计算成本') INT(11)"`
	UserId         int     `xorm:"comment('用户ID') INT(11)"`
	UserAccount    string  `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"comment('用户昵称') VARCHAR(30)"`
	GameId         int     `xorm:"comment('游戏id') INT(11)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	PromoteId      int     `xorm:"default 0 comment('推广id') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广员姓名') VARCHAR(30)"`
	ServerId       int     `xorm:"comment('区服id') INT(11)"`
	ServerName     string  `xorm:"comment('区服名称') VARCHAR(30)"`
	Amount         float64 `xorm:"comment('充值金额') DOUBLE(10,2)"`
	Remark         string  `xorm:"comment('备注') VARCHAR(100)"`
	Status         int     `xorm:"default 0 comment('状态') TINYINT(2)"`
	OpId           int     `xorm:"comment('操作人id') INT(11)"`
	OpAccount      string  `xorm:"comment('操作人账号') VARCHAR(30)"`
	CreateTime     int     `xorm:"comment('时间') INT(11)"`
}
