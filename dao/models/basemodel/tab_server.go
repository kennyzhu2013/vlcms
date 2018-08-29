package models

type TabServer struct {
	Id              int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	GameId          int    `xorm:"not null comment('游戏id') INT(11)"`
	GameName        string `xorm:"comment('游戏名称') VARCHAR(30)"`
	ServerName      string `xorm:"comment('区服名称') VARCHAR(30)"`
	ServerNum       int    `xorm:"comment('对接区服id') INT(11)"`
	RecommendStatus int    `xorm:"default 1 comment('推荐状态(0:否,1:是)') TINYINT(2)"`
	ShowStatus      int    `xorm:"default 1 comment('显示状态(0:否,1:是)') TINYINT(2)"`
	StopStatus      int    `xorm:"default 0 comment('是否停服(0:否,1:是)') TINYINT(2)"`
	ServerStatus    int    `xorm:"default 0 comment('区服状态(0:正常,1拥挤,2爆满)') TINYINT(2)"`
	Icon            int    `xorm:"comment('区服图标') INT(11)"`
	StartTime       int    `xorm:"comment('开始时间') INT(11)"`
	Desride         string `xorm:"comment('描述') VARCHAR(300)"`
	Prompt          string `xorm:"comment('停服提示') VARCHAR(300)"`
	ParentId        int    `xorm:"comment('父类id') INT(11)"`
	CreateTime      int    `xorm:"comment('创建时间') INT(11)"`
}
