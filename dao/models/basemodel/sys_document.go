package models

type SysDocument struct {
	Id          int    `xorm:"not null pk autoincr comment('文档ID') INT(10)"`
	Uid         int    `xorm:"not null default 0 comment('用户ID') index(idx_status_type_pid) INT(10)"`
	Name        string `xorm:"not null default '' comment('标识') CHAR(40)"`
	Title       string `xorm:"not null default '' comment('标题') CHAR(80)"`
	CategoryId  int    `xorm:"not null comment('所属分类') index(idx_category_status) INT(10)"`
	GroupId     int    `xorm:"not null comment('所属分组') SMALLINT(3)"`
	Description string `xorm:"not null default '' comment('描述') CHAR(140)"`
	Root        int    `xorm:"not null default 0 comment('根节点') INT(10)"`
	Pid         int    `xorm:"not null default 0 comment('所属ID') index(idx_status_type_pid) INT(10)"`
	ModelId     int    `xorm:"not null default 0 comment('内容模型ID') TINYINT(3)"`
	Type        int    `xorm:"not null default 2 comment('内容类型') TINYINT(3)"`
	Position    int    `xorm:"not null default 0 comment('推荐位') SMALLINT(5)"`
	LinkId      int    `xorm:"not null default 0 comment('外链') INT(10)"`
	CoverId     int    `xorm:"not null default 0 comment('封面') INT(10)"`
	Display     int    `xorm:"not null default 1 comment('可见性') TINYINT(3)"`
	Deadline    int    `xorm:"not null default 0 comment('截至时间') INT(10)"`
	Attach      int    `xorm:"not null default 0 comment('附件数量') TINYINT(3)"`
	View        int    `xorm:"not null default 0 comment('浏览量') INT(10)"`
	Comment     int    `xorm:"not null default 0 comment('评论数') INT(10)"`
	Extend      int    `xorm:"not null default 0 comment('扩展统计字段') INT(10)"`
	Level       int    `xorm:"not null default 0 comment('优先级') INT(10)"`
	CreateTime  int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime  int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Admin       int    `xorm:"not null default 0 comment('文档创建者') INT(11)"`
	Status      int    `xorm:"not null default 0 comment('数据状态') index(idx_category_status) index(idx_status_type_pid) TINYINT(4)"`
}
