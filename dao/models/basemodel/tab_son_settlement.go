package models

type TabSonSettlement struct {
	Id                  int     `xorm:"not null pk autoincr INT(11)"`
	SettlementNumber    string  `xorm:"VARCHAR(30)"`
	SettlementStartTime int     `xorm:"comment('结算开始时间') INT(11)"`
	SettlementEndTime   int     `xorm:"comment('结算结束时间') INT(11)"`
	GameId              int     `xorm:"comment('游戏id') INT(11)"`
	GameName            string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	PromoteId           int     `xorm:"comment('推广id(子账号id)') INT(11)"`
	PromoteAccount      string  `xorm:"comment('子账号') VARCHAR(30)"`
	Pattern             int     `xorm:"comment('合作模式') INT(11)"`
	SumMoney            float64 `xorm:"comment('充值总额') DOUBLE"`
	RegNumber           int     `xorm:"comment('注册人数') INT(11)"`
	Ratio               float64 `xorm:"comment('分成比例') DOUBLE(11)"`
	Money               float64 `xorm:"comment('注册单价') DOUBLE(11)"`
	JieMoney            float64 `xorm:"comment('结算金额') DOUBLE"`
	CreateTime          int     `xorm:"comment('添加时间') INT(11)"`
}
