/*
@Time : 2018/8/29 17:06 
@Author : kenny zhu
@File : adv_model.go
@Software: GoLand
@Others:
*/
package helper

import (
	"github.com/kennyzhu/vlcms/dao/models/basemodel"
	common "github.com/kennyzhu/vlcms/dao/models"
	"time"
)

var AdvHelper struct {
	BaseHelper
}

func init() {
	AdvHelper.DefaultInsertObject = defaultInsertObject
	AdvHelper.DefaultObject       = defaultObject
}

func defaultInsertObject() interface{} {
	obj := new(models.TabAdv)
	obj.StartTime = int( time.Now().UnixNano() )
	obj.EndTime = int( time.Now().UnixNano() )
	return obj
}

func defaultObject() interface{} {
	obj := new(models.TabAdv)
	return obj
}

// 结构体中extends标记对应的结构顺序应和最终生成SQL中对应的表出现的顺序相同..
type AdvAvsPos struct {
	models.TabAdv `xorm:"extends"`
	Name string
}

func (AdvAvsPos) TableName() string {
	return "user"
}

// Get advertise lists.
func Adv_lists(name string, limit, start int) []AdvAvsPos {
	obj := AdvHelper.DefaultObject().(models.TabAdv)
	advList := make([]AdvAvsPos, 0)
	session := common.ORM().Table(&obj).Where("status = ?", 1).Join("INNER", "tab_adv_pos", "tab_adv.pos_id = tab_adv_pos.id")
	session.Desc().Where("name = ?", name).Limit(limit, start).Find(&advList)
	return advList
}