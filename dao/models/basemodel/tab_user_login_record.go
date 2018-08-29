package models

type TabUserLoginRecord struct {
	Id           int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId       int    `xorm:"comment('游戏id') INT(11)"`
	GameName     string `xorm:"comment(' 游戏名称') VARCHAR(30)"`
	ServerId     int    `xorm:"comment('区服id') INT(11)"`
	ServerName   string `xorm:"comment('区服名称') VARCHAR(30)"`
	UserId       int    `xorm:"comment('用户ID') INT(11)"`
	UserAccount  string `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname string `xorm:"comment('用户昵称') VARCHAR(30)"`
	LoginTime    int    `xorm:"comment('登陆时间') INT(11)"`
	LoginIp      string `xorm:"comment('登陆ip') VARCHAR(20)"`
	Type         int    `xorm:"default 1 comment('类型(1:游戏登陆,2:PC登陆)') TINYINT(2)"`
}
