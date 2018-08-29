package models

import (
	"time"
)

type SysMember struct {
	Uid           int       `xorm:"not null pk autoincr comment('用户ID') INT(10)"`
	Nickname      string    `xorm:"not null default '' comment('昵称') CHAR(16)"`
	Sex           int       `xorm:"not null default 0 comment('性别') TINYINT(3)"`
	Birthday      time.Time `xorm:"not null default '0000-00-00' comment('生日') DATE"`
	Qq            string    `xorm:"not null default '' comment('qq号') CHAR(10)"`
	Score         int       `xorm:"not null default 0 comment('用户积分') MEDIUMINT(8)"`
	Login         int       `xorm:"not null default 0 comment('登录次数') INT(10)"`
	RegIp         int64     `xorm:"not null default 0 comment('注册IP') BIGINT(20)"`
	RegTime       int       `xorm:"not null default 0 comment('注册时间') INT(10)"`
	LastLoginIp   int64     `xorm:"not null default 0 comment('最后登录IP') BIGINT(20)"`
	LastLoginTime int       `xorm:"not null default 0 comment('最后登录时间') INT(10)"`
	Status        int       `xorm:"not null default 0 comment('会员状态') index TINYINT(4)"`
}
