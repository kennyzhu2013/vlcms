package models

type SysLinks struct {
	Id   int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	Mark string `xorm:"not null default '0' comment('所属后台') VARCHAR(50)"`
}
