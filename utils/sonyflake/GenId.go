package sonyflake

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
)

type sonyFlakeIdGenerator struct {
	machineID uint16
	sonyFlake *sonyflake.Sonyflake
}

func NewSonyFlakeIdGenerator(mID uint16) *sonyFlakeIdGenerator {
	idGenerator := &sonyFlakeIdGenerator{
		machineID: mID,
	}

	var st = sonyflake.Settings{}
	if mID <= 0 {
		st = sonyflake.Settings{
			MachineID: nil,
		}
	} else {
		st = sonyflake.Settings{
			MachineID: func() (uint16, error) {
				return idGenerator.machineID, nil
			},
		}
	}

	sonyFlake = sonyflake.NewSonyflake(st)
	idGenerator.sonyFlake = sonyFlake
	return idGenerator
}

// 获取 机器编码ID的 回调函数
func (s *sonyFlakeIdGenerator) getMachineID() (uint16, error) {
	// machineID 返回nil, 则返回专用IP地址的低16位
	return s.machineID, nil
}

// 获取全局 ID 的函数

func (s *sonyFlakeIdGenerator) GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("需要先初始化以后再执行 GetID 函数 err: %#v \n", err)
		return
	}
	return sonyFlake.NextID()
}
