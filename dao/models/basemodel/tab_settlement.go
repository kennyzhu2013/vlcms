package models

type TabSettlement struct {
	Id               int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	PromoteId        int     `xorm:"comment('推广员ID') INT(11)"`
	GameId           string  `xorm:"comment('游戏ID') VARCHAR(32)"`
	Status           int     `xorm:"default 1 comment('结算状态') TINYINT(2)"`
	TiStatus         int     `xorm:"default -1 comment('提现状态(-1未申请 0 申请中 1 已通过 2 已驳回)') INT(11)"`
	GameName         string  `xorm:"not null comment('游戏名称') VARCHAR(40)"`
	PromoteAccount   string  `xorm:"not null comment('推广员账号') VARCHAR(55)"`
	TotalMoney       float64 `xorm:"comment('充值总额') DOUBLE(10,2)"`
	TotalNumber      int     `xorm:"not null default 0 comment('注册人数') INT(19)"`
	SettlementNumber string  `xorm:"not null comment('结算单号') VARCHAR(60)"`
	Pattern          int     `xorm:"not null default 0 comment('合作模式') INT(11)"`
	Ratio            float64 `xorm:"default 0.00 comment('CPS分成比例(%)') DOUBLE(5,2)"`
	Money            float64 `xorm:"default 0.00 comment('CPA注册单价(元)') DOUBLE(10,2)"`
	SumMoney         float64 `xorm:"default 0.00 comment('结算金额') DOUBLE(10,2)"`
	CreateTime       int     `xorm:"comment('结算时间') INT(11)"`
}
