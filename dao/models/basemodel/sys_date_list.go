package models

import (
	"time"
)

type SysDateList struct {
	Time time.Time `xorm:"DATE"`
}
