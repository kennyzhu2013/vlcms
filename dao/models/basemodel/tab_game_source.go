package models

type TabGameSource struct {
	Id         int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	GameId     int    `xorm:"comment('游戏id') INT(11)"`
	GameName   string `xorm:"comment('游戏名称') VARCHAR(30)"`
	FileId     int    `xorm:"comment('文件id') INT(11)"`
	FileName   string `xorm:"comment('文件名称') VARCHAR(30)"`
	FileUrl    string `xorm:"comment('文件路径') VARCHAR(255)"`
	FileSize   string `xorm:"not null comment('文件大小') VARCHAR(30)"`
	FileType   int    `xorm:"comment('原包类型') TINYINT(2)"`
	OpId       int    `xorm:"comment('操作人id') INT(11)"`
	OpAccount  string `xorm:"comment('操作人名称') VARCHAR(30)"`
	Remark     string `xorm:"comment('备注') VARCHAR(100)"`
	CreateTime int    `xorm:"comment('时间') INT(11)"`
}
