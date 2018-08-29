package models

type SysGametype struct {
	Id    int `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Icon  int `xorm:"not null comment('图标') INT(12)"`
	Cover int `xorm:"not null comment('封面') INT(12)"`
}
