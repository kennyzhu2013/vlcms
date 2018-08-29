package models

type SysModel struct {
	Id             int    `xorm:"not null pk autoincr comment('模型ID') INT(10)"`
	Name           string `xorm:"not null default '' comment('模型标识') CHAR(30)"`
	Title          string `xorm:"not null default '' comment('模型名称') CHAR(30)"`
	Extend         int    `xorm:"not null default 0 comment('继承的模型') INT(10)"`
	Relation       string `xorm:"not null default '' comment('继承与被继承模型的关联字段') VARCHAR(30)"`
	NeedPk         int    `xorm:"not null default 1 comment('新建表时是否需要主键字段') TINYINT(1)"`
	FieldSort      string `xorm:"comment('表单字段排序') TEXT"`
	FieldGroup     string `xorm:"not null default '1:基础' comment('字段分组') VARCHAR(255)"`
	AttributeList  string `xorm:"comment('属性列表（表的字段）') TEXT"`
	AttributeAlias string `xorm:"not null default '' comment('属性别名定义') VARCHAR(255)"`
	TemplateList   string `xorm:"not null default '' comment('列表模板') VARCHAR(100)"`
	TemplateAdd    string `xorm:"not null default '' comment('新增模板') VARCHAR(100)"`
	TemplateEdit   string `xorm:"not null default '' comment('编辑模板') VARCHAR(100)"`
	ListGrid       string `xorm:"comment('列表定义') TEXT"`
	ListRow        int    `xorm:"not null default 10 comment('列表数据长度') SMALLINT(2)"`
	SearchKey      string `xorm:"not null default '' comment('默认搜索字段') VARCHAR(50)"`
	SearchList     string `xorm:"not null default '' comment('高级搜索的字段') VARCHAR(255)"`
	CreateTime     int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	UpdateTime     int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status         int    `xorm:"not null default 0 comment('状态') TINYINT(3)"`
	EngineType     string `xorm:"not null default 'MyISAM' comment('数据库引擎') VARCHAR(25)"`
}
