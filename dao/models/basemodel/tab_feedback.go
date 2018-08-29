package models

type TabFeedback struct {
	Id         int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Username   string `xorm:"comment('用户名') VARCHAR(30)"`
	Nickname   string `xorm:"comment('昵称') VARCHAR(30)"`
	Content    string `xorm:"comment('内容') VARCHAR(300)"`
	Contact    string `xorm:"comment('联系人') VARCHAR(11)"`
	CreateTime int    `xorm:"comment('时间') INT(11)"`
}
