package utils

import (
	"sync/atomic"
	"time"
)

// 时间戳：占用41位，记录生成ID的时间戳，精确到毫秒级。
// 机器ID：占用10位，用于标识不同的机器。
// 序列号：占用12位，用于解决同一毫秒内生成多个ID的冲突。

// 时间戳的位数、机器ID的位数、序列号的位数
const (
	timestampBits  = 41                         // 时间戳位数
	machineIDBits  = 10                         // 机器ID位数
	sequenceBits   = 12                         // 序列号位数
	maxMachineID   = -1 ^ (-1 << machineIDBits) // 最大机器ID
	maxSequenceNum = -1 ^ (-1 << sequenceBits)  // 最大序列号
)

type Snowflake struct {
	timestamp   int64 // 时间戳
	machineID   int64 // 机器ID
	sequenceNum int64 // 序列号
}

func NewSnowflake(machineID int64) *Snowflake {
	if machineID < 0 || machineID > maxMachineID {
		panic("machineID must be between 1 and 1023")
	}

	return &Snowflake{
		timestamp:   time.Now().UnixNano() / 1e6,
		machineID:   machineID,
		sequenceNum: 0,
	}
}

func (s *Snowflake) waitNextMillis() int64 {
	currentTimestamp := time.Now().UnixNano() / 1e6
	for currentTimestamp <= s.timestamp {
		currentTimestamp = time.Now().UnixNano() / 1e6
	}
	return currentTimestamp
}

func (s *Snowflake) GenerateID() int64 {
	currentTimestamp := time.Now().UnixNano() / 1e6
	if currentTimestamp == s.timestamp {
		s.sequenceNum = (atomic.AddInt64(&s.sequenceNum, 1)) & maxSequenceNum
		if s.sequenceNum == 0 {
			currentTimestamp = s.waitNextMillis()
		}
	} else {
		s.sequenceNum = 0
	}
	s.timestamp = currentTimestamp
	id := (currentTimestamp << (machineIDBits + sequenceBits)) |
		(s.machineID << sequenceBits) | s.sequenceNum
	return id
}
