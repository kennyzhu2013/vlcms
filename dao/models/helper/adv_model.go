/*
@Time : 2018/8/29 17:06 
@Author : kenny zhu
@File : adv_model.go
@Software: GoLand
@Others:
*/
package helper

import "strconv"

func (m *Preferences) adv_lists() string {
	return strconv.Itoa(m.User)
}