package models

type TabMend struct {
	Id               int    `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	UserId           int    `xorm:"not null comment('用户id') INT(11)"`
	UserAccount      string `xorm:"comment('用户账号') VARCHAR(30)"`
	UserNickname     string `xorm:"comment('用户昵称') VARCHAR(30)"`
	PromoteId        int    `xorm:"comment('推广员id') INT(11)"`
	PromoteAccount   string `xorm:"comment('推广员账号') VARCHAR(30)"`
	PromoteIdTo      int    `xorm:"comment('修改后推广员id') INT(11)"`
	PromoteAccountTo string `xorm:"comment('修改后推广员账号') VARCHAR(30)"`
	Remark           string `xorm:"comment('备注') VARCHAR(100)"`
	CreateTime       int    `xorm:"comment('创建时间') INT(11)"`
	OpId             int    `xorm:"comment('操作人id') INT(11)"`
	OpAccount        string `xorm:"comment('操作人账号') VARCHAR(30)"`
}
