package models

type SysFile struct {
	Id         int    `xorm:"not null pk autoincr comment('文件ID') INT(10)"`
	Name       string `xorm:"not null default '' comment('原始文件名') CHAR(30)"`
	Savename   string `xorm:"not null default '' comment('保存名称') CHAR(20)"`
	Savepath   string `xorm:"not null default '' comment('文件保存路径') CHAR(30)"`
	Ext        string `xorm:"not null default '' comment('文件后缀') CHAR(5)"`
	Mime       string `xorm:"not null default '' comment('文件mime类型') CHAR(40)"`
	Size       int    `xorm:"not null default 0 comment('文件大小') INT(10)"`
	Md5        string `xorm:"not null default '' comment('文件md5') unique CHAR(32)"`
	Sha1       string `xorm:"not null default '' comment('文件 sha1编码') CHAR(40)"`
	Location   int    `xorm:"not null default 0 comment('文件保存位置') TINYINT(3)"`
	Url        string `xorm:"not null default '' comment('远程地址') VARCHAR(255)"`
	CreateTime int    `xorm:"not null default 0 comment('上传时间') INT(10)"`
}
