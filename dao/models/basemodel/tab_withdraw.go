package models

type TabWithdraw struct {
	Id               int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	SettlementNumber string  `xorm:"comment('结算单号') VARCHAR(50)"`
	SumMoney         float64 `xorm:"comment('结算金额') DOUBLE"`
	PromoteId        int     `xorm:"comment('推广员id') INT(11)"`
	PromoteAccount   string  `xorm:"comment('推广姓名') VARCHAR(30)"`
	CreateTime       int     `xorm:"comment('申请时间') INT(11)"`
	EndTime          int     `xorm:"comment('完成时间') INT(11)"`
	Status           int     `xorm:"default 0 comment('提现状态(-1未申请 0 申请中 1 已通过 2 已驳回)') INT(11)"`
}
