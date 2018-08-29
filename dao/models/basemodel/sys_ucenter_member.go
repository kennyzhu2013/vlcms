package models

type SysUcenterMember struct {
	Id            int    `xorm:"not null pk autoincr comment('用户ID') INT(10)"`
	Username      string `xorm:"not null comment('用户名') unique CHAR(16)"`
	Password      string `xorm:"not null comment('密码') CHAR(32)"`
	Email         string `xorm:"not null comment('用户邮箱') unique CHAR(32)"`
	Mobile        string `xorm:"not null default '' comment('用户手机') CHAR(15)"`
	RegTime       int    `xorm:"not null default 0 comment('注册时间') INT(10)"`
	RegIp         int64  `xorm:"not null default 0 comment('注册IP') BIGINT(20)"`
	LastLoginTime int    `xorm:"not null default 0 comment('最后登录时间') INT(10)"`
	LastLoginIp   int64  `xorm:"not null default 0 comment('最后登录IP') BIGINT(20)"`
	UpdateTime    int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status        int    `xorm:"default 0 comment('用户状态') index TINYINT(4)"`
}
