package models

type SysAttachment struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	Uid        int    `xorm:"not null default 0 comment('用户ID') INT(10)"`
	Title      string `xorm:"not null default '' comment('附件显示名') CHAR(30)"`
	Type       int    `xorm:"not null default 0 comment('附件类型') TINYINT(3)"`
	Source     int    `xorm:"not null default 0 comment('资源ID') INT(10)"`
	RecordId   int    `xorm:"not null default 0 comment('关联记录ID') index(idx_record_status) INT(10)"`
	Download   int    `xorm:"not null default 0 comment('下载次数') INT(10)"`
	Size       int64  `xorm:"not null default 0 comment('附件大小') BIGINT(20)"`
	Dir        int    `xorm:"not null default 0 comment('上级目录ID') INT(12)"`
	Sort       int    `xorm:"not null default 0 comment('排序') INT(8)"`
	CreateTime int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime int    `xorm:"not null default 0 comment('更新时间') INT(11)"`
	Status     int    `xorm:"not null default 0 comment('状态') index(idx_record_status) TINYINT(1)"`
}
