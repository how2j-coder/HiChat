package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch         int64 = 1609459200000 // 2021-01-01 00:00:00 UTC的毫秒时间戳
	timestampBits uint8 = 41            // 时间戳占用位数
	machineIDBits uint8 = 10            // 机器ID占用位数
	sequenceBits  uint8 = 12            // 序列号占用位数

	maxMachineID int64 = -1 ^ (-1 << machineIDBits) // 最大机器ID
	maxSequence  int64 = -1 ^ (-1 << sequenceBits)   // 最大序列号
)

// Snowflake 分布式ID生成器结构体
type Snowflake struct {
	mutex     sync.Mutex
	timestamp int64 // 上次生成ID的时间戳
	machineID int64 // 机器标识ID
	sequence  int64 // 序列号
}

// NewSnowflake 创建Snowflake实例
// machineID范围: 0 <= machineID <= maxMachineID
func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("machine ID out of range")
	}

	return &Snowflake{
		timestamp: 0,
		machineID: machineID,
		sequence:  0,
	}, nil
}

// Generate 生成全局唯一ID
func (s *Snowflake) Generate() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now().UnixMilli()

	// 处理时钟回拨问题
	if now < s.timestamp {
		return 0, errors.New("clock moved backwards")
	}

	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// 当前毫秒序列号已用完，等待下一毫秒
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	// 组合各部分生成最终ID
	id := (now-epoch)<<(machineIDBits+sequenceBits) |
		(s.machineID << sequenceBits) |
		s.sequence

	return id, nil
}

// ParseID 解析雪花ID的组成成分
func ParseID(id int64) (t time.Time, machineID int64, sequence int64) {
	sequence = id & maxSequence
	id >>= sequenceBits
	machineID = id & maxMachineID
	id >>= machineIDBits
	t = time.UnixMilli(id + epoch)
	return
}
