package models

type TabBill struct {
	Id               int     `xorm:"not null pk autoincr comment('主键') INT(10)"`
	BillNumber       string  `xorm:"not null comment('对账单号') VARCHAR(80)"`
	BillTime         string  `xorm:"not null comment('对账时间') VARCHAR(30)"`
	PromoteId        int     `xorm:"comment('推广员ID') INT(11)"`
	PromoteAccount   string  `xorm:"comment('推广员账号') VARCHAR(30)"`
	GameId           int     `xorm:"comment('游戏ID') INT(11)"`
	GameName         string  `xorm:"not null comment('游戏名称') VARCHAR(30)"`
	TotalMoney       float64 `xorm:"default 0.00 comment('充值总额') DOUBLE(10,2)"`
	TotalNumber      int     `xorm:"not null comment('注册人数') INT(19)"`
	Status           int     `xorm:"not null default 1 comment('状态') TINYINT(2)"`
	BillStartTime    int     `xorm:"not null comment('对账开始时间') INT(19)"`
	BillEndTime      int     `xorm:"not null comment('对账结束时间') INT(19)"`
	CreateTime       int     `xorm:"comment('创建时间') INT(11)"`
	SettlementStatus int     `xorm:"not null default 0 comment('结算状态') TINYINT(2)"`
}
