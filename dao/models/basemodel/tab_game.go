package models

type TabGame struct {
	Id              int     `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	GameName        string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	Sort            int     `xorm:"default 1 comment('排序') INT(10)"`
	Short           string  `xorm:"comment('游戏简写') VARCHAR(20)"`
	GameTypeId      int     `xorm:"comment('游戏类型id') INT(10)"`
	GameTypeName    string  `xorm:"comment('游戏类型名称') VARCHAR(20)"`
	GameScore       float64 `xorm:"comment('游戏评分') DOUBLE(3)"`
	Features        string  `xorm:"comment('游戏特征') VARCHAR(50)"`
	RecommendLevel  float64 `xorm:"DOUBLE(3)"`
	Version         string  `xorm:"comment('版本号') VARCHAR(10)"`
	GameSize        string  `xorm:"default '0' comment('游戏大小') VARCHAR(10)"`
	Icon            int     `xorm:"comment('游戏图标') INT(11)"`
	Cover           int     `xorm:"comment('游戏封面') INT(11)"`
	Screenshot      string  `xorm:"comment('游戏截图') VARCHAR(255)"`
	Introduction    string  `xorm:"comment('游戏简介') VARCHAR(300)"`
	GameApi         string  `xorm:"comment('接口简写') VARCHAR(50)"`
	AndDowAddress   string  `xorm:"not null comment('安卓游戏下载地址') VARCHAR(255)"`
	IosDowAddress   string  `xorm:"comment('ios游戏下载地址') VARCHAR(255)"`
	GameAddress     string  `xorm:"comment('外部链接游戏地址') VARCHAR(255)"`
	DowNum          int     `xorm:"default 0 comment('游戏下载数量') INT(10)"`
	GameStatus      int     `xorm:"comment('游戏状态(0:关闭,1:开启)') TINYINT(2)"`
	PayStatus       int     `xorm:"default 1 comment('充值状态(0:关闭,1:开启)') TINYINT(2)"`
	DowStatus       int     `xorm:"default 1 comment('下载状态(0:关闭,1:开启)') TINYINT(2)"`
	MixStatus       int     `xorm:"default 0 comment('是否参与混服(0:不参与,1:参与)') TINYINT(2)"`
	Developers      string  `xorm:"comment('开发商') VARCHAR(30)"`
	CreateTime      int     `xorm:"comment('创建时间') INT(11)"`
	Discount        float64 `xorm:"comment('折扣') DOUBLE(4,2)"`
	Language        string  `xorm:"comment('语言') VARCHAR(10)"`
	GameAppid       string  `xorm:"comment('游戏appid') VARCHAR(32)"`
	GameCoinName    string  `xorm:"comment('游戏币名称') VARCHAR(10)"`
	GameCoinRation  string  `xorm:"comment('游戏币比例') VARCHAR(10)"`
	Category        int     `xorm:"default 0 comment('游戏开放类型') TINYINT(2)"`
	Ratio           float64 `xorm:"default 0.00 comment('分成比例') DOUBLE(5,2)"`
	Money           int     `xorm:"default 0 comment('注册单价') INT(11)"`
	RecommendStatus string  `xorm:"not null default '1' comment('推荐状态') VARCHAR(10)"`
}
