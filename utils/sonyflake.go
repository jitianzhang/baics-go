package utils

import (
	"github.com/sony/sonyflake"
	"time"
)

func NewSonyFlakeIdGenerator() *sonyflake.Sonyflake {
	Obj, _ := sonyflake.New(sonyflake.Settings{StartTime: time.Now()})
	return Obj
}
