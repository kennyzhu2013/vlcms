package models

type TabMessage struct {
	Qq            string `xorm:"comment('QQ号') VARCHAR(20)"`
	Id            int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	GameId        int    `xorm:"comment('游戏id') INT(11)"`
	UserId        int    `xorm:"comment('用户id') INT(11)"`
	Title         string `xorm:"comment('标题') VARCHAR(50)"`
	Content       string `xorm:"comment('内容') VARCHAR(300)"`
	Status        int    `xorm:"default 0 comment('修复状态(0:未,1:已)') TINYINT(2)"`
	Type          int    `xorm:"default 0 comment('类型(0:纠错,1:留言)') TINYINT(2)"`
	OpId          int    `xorm:"comment('操作人id') INT(11)"`
	OpAccount     string `xorm:"VARCHAR(20)"`
	CreateTime    int    `xorm:"comment('修改时间') INT(11)"`
	Othercontents string `xorm:"VARCHAR(255)"`
}
