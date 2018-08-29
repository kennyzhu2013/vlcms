package models

type SysCategory struct {
	Id             int    `xorm:"not null pk autoincr comment('分类ID') INT(10)"`
	Name           string `xorm:"not null comment('标志') unique VARCHAR(30)"`
	Title          string `xorm:"not null comment('标题') VARCHAR(50)"`
	Pid            int    `xorm:"not null default 0 comment('上级分类ID') index INT(10)"`
	Sort           int    `xorm:"not null default 0 comment('排序（同级有效）') INT(10)"`
	ListRow        int    `xorm:"not null default 10 comment('列表每页行数') TINYINT(3)"`
	MetaTitle      string `xorm:"not null default '' comment('SEO的网页标题') VARCHAR(50)"`
	Keywords       string `xorm:"not null default '' comment('关键字') VARCHAR(255)"`
	Description    string `xorm:"not null default '' comment('描述') VARCHAR(255)"`
	TemplateIndex  string `xorm:"not null default '' comment('频道页模板') VARCHAR(100)"`
	TemplateLists  string `xorm:"not null default '' comment('列表页模板') VARCHAR(100)"`
	TemplateDetail string `xorm:"not null default '' comment('详情页模板') VARCHAR(100)"`
	TemplateEdit   string `xorm:"not null default '' comment('编辑页模板') VARCHAR(100)"`
	Model          string `xorm:"not null default '' comment('列表绑定模型') VARCHAR(100)"`
	ModelSub       string `xorm:"not null default '' comment('子文档绑定模型') VARCHAR(100)"`
	Type           string `xorm:"not null default '' comment('允许发布的内容类型') VARCHAR(100)"`
	LinkId         int    `xorm:"not null default 0 comment('外链') INT(10)"`
	AllowPublish   int    `xorm:"not null default 0 comment('是否允许发布内容') TINYINT(3)"`
	Display        int    `xorm:"not null default 0 comment('可见性') TINYINT(3)"`
	Reply          int    `xorm:"not null default 0 comment('是否允许回复') TINYINT(3)"`
	Check          int    `xorm:"not null default 0 comment('发布的文章是否需要审核') TINYINT(3)"`
	ReplyModel     string `xorm:"not null default '' VARCHAR(100)"`
	Extend         string `xorm:"comment('扩展设置') TEXT"`
	CreateTime     int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime     int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status         int    `xorm:"not null default 0 comment('数据状态') TINYINT(4)"`
	Icon           int    `xorm:"not null default 0 comment('分类图标') INT(10)"`
	Groups         string `xorm:"not null default '' comment('分组定义') VARCHAR(255)"`
}
