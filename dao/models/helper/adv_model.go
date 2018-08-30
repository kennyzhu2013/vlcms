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

// 支持分表要用ShardGetTableName...
func (AdvAvsPos) TableName() string {
	return common.ShardGetTableName(&models.TabAdv{}, 0)
}

// Get advertise lists.
func Adv_lists(name string, limit, start int) []AdvAvsPos {
	// obj := AdvHelper.DefaultObject().(models.TabAdv)
	advList := make([]AdvAvsPos, 0)
	advTabName := common.ShardGetTableName(&models.TabAdv{}, 0)
	advPosTabName := common.ShardGetTableName(&models.TabAdvPos{}, 0)
	condition := advTabName + ".pos_id = " + advPosTabName + ".id"
	session := common.ORM().Table(advTabName).Where("status = ?", 1).Join("INNER", advPosTabName, condition)
	session.Desc().Where("name = ?", name).Limit(limit, start).Find(&advList)
	return advList
}