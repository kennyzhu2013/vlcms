package models

type TabTool struct {
	Id         int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Name       string `xorm:"comment('标识') VARCHAR(30)"`
	Title      string `xorm:"comment('标题') VARCHAR(30)"`
	Config     string `xorm:"comment('配置文件内容') TEXT"`
	Template   string `xorm:"comment('模板内容') TEXT"`
	Type       int    `xorm:"comment('类型') TINYINT(3)"`
	Status     int    `xorm:"comment('状态') TINYINT(3)"`
	CreateTime int    `xorm:"comment('添加时间') INT(11)"`
}
