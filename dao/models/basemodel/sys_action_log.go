package models

type SysActionLog struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	ActionId   int    `xorm:"not null default 0 comment('行为id') index INT(10)"`
	UserId     int    `xorm:"not null default 0 comment('执行用户id') index INT(10)"`
	ActionIp   int64  `xorm:"not null comment('执行行为者ip') index BIGINT(20)"`
	Model      string `xorm:"not null default '' comment('触发行为的表') VARCHAR(50)"`
	RecordId   int    `xorm:"not null default 0 comment('触发行为的数据id') INT(10)"`
	Remark     string `xorm:"not null default '' comment('日志备注') VARCHAR(255)"`
	Status     int    `xorm:"not null default 1 comment('状态') TINYINT(2)"`
	CreateTime int    `xorm:"not null default 0 comment('执行行为的时间') INT(10)"`
}
