package models

type TabUser struct {
	Id             int     `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Account        string  `xorm:"comment('登陆账号') VARCHAR(30)"`
	Password       string  `xorm:"comment('登陆密码') VARCHAR(32)"`
	PromoteId      int     `xorm:"not null default 0 comment('推广id') INT(11)"`
	PromoteAccount string  `xorm:"comment('推广员账号') VARCHAR(30)"`
	ParentId       int     `xorm:"default 0 comment('父类id') INT(11)"`
	ParentName     string  `xorm:"comment('父类名称') VARCHAR(30)"`
	FgameId        int     `xorm:"default 0 INT(11)"`
	FgameName      string  `xorm:"VARCHAR(30)"`
	Nickname       string  `xorm:"comment('昵称') VARCHAR(30)"`
	Sex            int     `xorm:"default 0 comment('性别(0 男 1 女)') INT(11)"`
	Email          string  `xorm:"comment('邮箱') VARCHAR(50)"`
	Phone          string  `xorm:"comment('手机号码') VARCHAR(15)"`
	RealName       string  `xorm:"comment('真实姓名') VARCHAR(20)"`
	Idcard         string  `xorm:"comment('身份证') VARCHAR(20)"`
	VipLevel       int     `xorm:"default 0 comment('vip等级') TINYINT(2)"`
	Cumulative     float64 `xorm:"default 0.00 comment('累计充值') DOUBLE(10,2)"`
	Balance        float64 `xorm:"default 0.00 comment('余额') DOUBLE(10,2)"`
	AntiAddiction  int     `xorm:"default 0 comment('防沉迷') TINYINT(2)"`
	LockStatus     int     `xorm:"default 1 comment('锁定状态') TINYINT(2)"`
	RegisterWay    int     `xorm:"default 0 comment('注册方式') TINYINT(2)"`
	RegisterTime   int     `xorm:"comment('注册时间') INT(11)"`
	LoginTime      int     `xorm:"comment('登陆时间') INT(11)"`
	RegisterIp     string  `xorm:"not null comment('注册ip') VARCHAR(16)"`
	LoginIp        string  `xorm:"not null comment('登陆ip') VARCHAR(16)"`
	IsCheck        int     `xorm:"not null default 1 comment('是否对账  1参与 2不参与 3参与(已对账) 4不参与(已对账)') INT(11)"`
	SubStatus      int     `xorm:"default 0 comment('子渠道结算状态(0未结算 1已结算)') INT(11)"`
	Collection     string  `xorm:"comment('游戏收藏  id') VARCHAR(255)"`
	UserType       int     `xorm:"default 0 comment('用户类型 0：用户，1：混服') SMALLINT(2)"`
	MixUserId      int     `xorm:"default 0 comment('所属混服站点（所属混服用户）') INT(11)"`
	Openid         string  `xorm:"comment('用户标识') VARCHAR(255)"`
}
