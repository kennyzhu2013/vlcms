package models

type TabGiftbagType struct {
	Id         int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	TypeName   string `xorm:"comment('礼包类型名称') VARCHAR(20)"`
	Status     int    `xorm:"default 1 comment('状态(0:禁用,1:启用)') TINYINT(2)"`
	OpId       int    `xorm:"comment('操作人id') INT(11)"`
	OpNickname string `xorm:"comment('操作人昵称') VARCHAR(30)"`
	CreateTime int    `xorm:"comment('添加时间') INT(11)"`
}
