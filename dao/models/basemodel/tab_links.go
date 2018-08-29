package models

type TabLinks struct {
	Id         int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Title      string `xorm:"comment('标题') VARCHAR(50)"`
	LinkUrl    string `xorm:"comment('友链地址') VARCHAR(255)"`
	Status     int    `xorm:"comment('状态') TINYINT(2)"`
	Type       int    `xorm:"comment('类型') TINYINT(2)"`
	Remark     string `xorm:"comment('备注') VARCHAR(255)"`
	CreateTime int    `xorm:"comment('创建时间') INT(11)"`
	Mark       string `xorm:"default '0' comment('0媒体 1渠道') VARCHAR(50)"`
}
