package models

type TabRoute struct {
	Id      int    `xorm:"not null pk autoincr comment('路由id') INT(11)"`
	FullUrl string `xorm:"comment('完整url， 如：portal/list/index?id=1') VARCHAR(255)"`
	Url     string `xorm:"comment('实际显示的url') VARCHAR(255)"`
	Status  int    `xorm:"default 1 comment('状态，1：启用 ;0：不启用') TINYINT(1)"`
}
