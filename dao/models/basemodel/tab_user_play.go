package models

type TabUserPlay struct {
	Id             int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	UserId         int     `xorm:"comment('用户ID') INT(11)"`
	UserAccount    string  `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname   string  `xorm:"comment('用户昵称') VARCHAR(30)"`
	GameId         int     `xorm:"not null comment('游戏id') INT(11)"`
	GameName       string  `xorm:"comment('游戏名称') VARCHAR(30)"`
	GameAppid      string  `xorm:"comment('游戏appid') VARCHAR(32)"`
	ServerId       int     `xorm:"not null comment('区服id') INT(11)"`
	ServerName     string  `xorm:"comment('区服名称') VARCHAR(30)"`
	RoleId         int     `xorm:"default 0 comment('角色') INT(11)"`
	BindBalance    float64 `xorm:"default 0.00 comment('绑定平台币') DOUBLE(10,2)"`
	RoleName       string  `xorm:"comment('角色名称') VARCHAR(20)"`
	RoleLevel      int     `xorm:"default 0 comment('等级') INT(3)"`
	PromoteId      int     `xorm:"default 0 comment('推广员id') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广员账号') VARCHAR(30)"`
}
