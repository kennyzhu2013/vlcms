package models

type TabGameType struct {
	Id         int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	TypeName   string `xorm:"comment('游戏类型名称') VARCHAR(20)"`
	Icon       int    `xorm:"not null comment('图标') INT(12)"`
	Cover      int    `xorm:"not null comment('封面') INT(12)"`
	Status     int    `xorm:"default 1 comment('状态(-1:删除,1:正常)') TINYINT(2)"`
	StatusShow int    `xorm:"default 1 comment('显示状态(0:不显示,1:显示)') TINYINT(2)"`
	OpId       int    `xorm:"comment('操作人id') INT(11)"`
	OpNickname string `xorm:"comment('操作人昵称') VARCHAR(30)"`
	CreateTime int    `xorm:"comment('添加时间') INT(11)"`
	Sort       int    `xorm:"default 0 INT(11)"`
}
