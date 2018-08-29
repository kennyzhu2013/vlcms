package models

type SysDocumentArticle struct {
	Id       int    `xorm:"not null pk default 0 comment('文档ID') INT(10)"`
	Parse    int    `xorm:"not null default 0 comment('内容解析类型') TINYINT(3)"`
	Content  string `xorm:"not null comment('文章内容') TEXT"`
	Template string `xorm:"not null default '' comment('详情页显示模板') VARCHAR(100)"`
	Bookmark int    `xorm:"not null default 0 comment('收藏数') INT(10)"`
}
