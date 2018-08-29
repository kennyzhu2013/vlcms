package models

type TabPromote struct {
	Id            int     `xorm:"not null pk autoincr comment('自增主键') INT(11)"`
	Account       string  `xorm:"comment('账号') VARCHAR(30)"`
	Password      string  `xorm:"comment('密码') VARCHAR(32)"`
	SecondPwd     string  `xorm:"comment('二级密码') VARCHAR(32)"`
	Nickname      string  `xorm:"comment('昵称') VARCHAR(30)"`
	MobilePhone   string  `xorm:"comment('手机号') VARCHAR(11)"`
	Email         string  `xorm:"comment('邮箱') VARCHAR(50)"`
	RealName      string  `xorm:"comment('真实姓名') VARCHAR(10)"`
	BankName      string  `xorm:"comment('银行') VARCHAR(50)"`
	BankCard      string  `xorm:"comment('银行卡') VARCHAR(20)"`
	Money         float64 `xorm:"default 0.00 comment('金额') DOUBLE(10,2)"`
	TotalMoney    float64 `xorm:"default 0.00 comment('总金额') DOUBLE(10,2)"`
	BalanceCoin   float64 `xorm:"default 0.00 comment('平台币') DOUBLE(10,2)"`
	PromoteType   int     `xorm:"default 1 comment('推广员类型') INT(2)"`
	Status        int     `xorm:"default 1 comment('状态') INT(11)"`
	ParentId      int     `xorm:"default 0 comment('父类ID') INT(11)"`
	PayLimit      int     `xorm:"default 0 comment('代充额度') INT(11)"`
	RefereeId     int     `xorm:"default 0 comment('推荐人ID') INT(11)"`
	SetPayTime    int     `xorm:"comment('代充额度设置时间') INT(11)"`
	LastLoginTime int     `xorm:"comment('最后登陆时间') INT(11)"`
	CreateTime    int     `xorm:"comment('添加时间') INT(11)"`
	AdminId       int     `xorm:"not null comment('管理员id') INT(11)"`
	Mark1         string  `xorm:"comment('基本信息备注') VARCHAR(255)"`
	BankArea      string  `xorm:"comment('开户地点') VARCHAR(255)"`
	AccountOpenin string  `xorm:"comment('开户网点') VARCHAR(255)"`
	BankAccount   string  `xorm:"comment('银行户名') VARCHAR(255)"`
	Mark2         string  `xorm:"comment('结算信息备注') VARCHAR(255)"`
}
