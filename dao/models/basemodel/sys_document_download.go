package models

type SysDocumentDownload struct {
	Id       int    `xorm:"not null pk default 0 comment('文档ID') INT(10)"`
	Parse    int    `xorm:"not null default 0 comment('内容解析类型') TINYINT(3)"`
	Content  string `xorm:"not null comment('下载详细描述') TEXT"`
	Template string `xorm:"not null default '' comment('详情页显示模板') VARCHAR(100)"`
	FileId   int    `xorm:"not null default 0 comment('文件ID') INT(10)"`
	Download int    `xorm:"not null default 0 comment('下载次数') INT(10)"`
	Size     int64  `xorm:"not null default 0 comment('文件大小') BIGINT(20)"`
}
