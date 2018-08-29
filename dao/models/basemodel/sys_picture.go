package models

type SysPicture struct {
	Id         int    `xorm:"not null pk autoincr comment('主键id自增') INT(10)"`
	Path       string `xorm:"not null default '' comment('路径') VARCHAR(255)"`
	Url        string `xorm:"not null default '' comment('图片链接') VARCHAR(255)"`
	Md5        string `xorm:"not null default '' comment('文件md5') CHAR(32)"`
	Sha1       string `xorm:"not null default '' comment('文件 sha1编码') CHAR(40)"`
	Status     int    `xorm:"not null default 0 comment('状态') TINYINT(2)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
}
