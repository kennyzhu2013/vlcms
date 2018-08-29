package models

type TabApply struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId         int     `xorm:"comment('游戏ID') INT(11)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	PromoteId      int     `xorm:"comment('推广员ID') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广账号') VARCHAR(30)"`
	Ratio          float64 `xorm:"default 0.00 comment('分成比例') DOUBLE(5,2)"`
	ApplyTime      int     `xorm:"comment('申请时间') INT(11)"`
	Status         int     `xorm:"comment('审核状态') TINYINT(2)"`
	EnableStatus   int     `xorm:"comment('操作状态') TINYINT(2)"`
	PackUrl        string  `xorm:"comment('游戏包地址') VARCHAR(255)"`
	DowUrl         string  `xorm:"comment('下载地址') VARCHAR(255)"`
	DisposeId      int     `xorm:"comment('操作人') INT(11)"`
	DisposeTime    int     `xorm:"comment('操作时间') INT(11)"`
	RegisterUrl    string  `xorm:"comment('注册登录url') VARCHAR(255)"`
}
