/*
@Time : 2018/8/29 18:49 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others:
*/
package helper

type DefaultDao func() interface{}

type BaseHelper struct {
	// TablePrefix string

	// auto
	DefaultInsertObject DefaultDao
	DefaultObject       DefaultDao
}


