package models

type TabGiftbag struct {
	Id              int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId          int    `xorm:"comment('游戏id') INT(11)"`
	GameName        string `xorm:"comment('游戏名称') VARCHAR(30)"`
	ServerId        int    `xorm:"comment('区服id') INT(11)"`
	ServerName      string `xorm:"comment('区服名称') VARCHAR(30)"`
	GiftbagName     string `xorm:"comment('礼包名称') VARCHAR(30)"`
	GiftbagType     int    `xorm:"comment('礼包类型') TINYINT(2)"`
	Level           int    `xorm:"comment('领取等级') INT(3)"`
	Sort            int    `xorm:"default 0 comment('排序') INT(10)"`
	Status          int    `xorm:"default 1 comment('状态') TINYINT(2)"`
	CallApi         int    `xorm:"default 0 comment('调用接口') TINYINT(2)"`
	TongServer      int    `xorm:"default 0 comment('是否通服') TINYINT(2)"`
	RecommendStatus int    `xorm:"comment('礼包推荐状态') TINYINT(2)"`
	Novice          string `xorm:"comment('激活码') VARCHAR(3000)"`
	Digest          string `xorm:"comment('摘要') VARCHAR(300)"`
	Desribe         string `xorm:"comment('描述') VARCHAR(300)"`
	StartTime       int    `xorm:"comment('开始时间') INT(11)"`
	EndTime         int    `xorm:"comment('结束时间') INT(11)"`
	CreateTime      int    `xorm:"comment('创建时间') INT(11)"`
}
