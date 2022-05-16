package sequence

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Sequence struct {
	startTimestamp     int64      // 开始时间戳
	workerIdBits       int64      // 机器ID所占的位数
	dataCenterIdBits   int64      // 数据标识ID所占的位数
	maxWorkerId        int64      // 支持的最大机器ID
	maxDataCenterId    int64      // 支持的最大机房ID
	sequenceBits       int64      // 序列在ID中占的位数
	workerIdShift      int64      // 机器ID向左移位数
	dataCenterIdShift  int64      // 机房ID向左移位数
	timestampLeftShift int64      // 时间截向左移位数
	sequenceMask       int64      // 生成序列的掩码最大值
	workerId           int64      // 工作机器ID
	dataCenterId       int64      // 机房ID
	sequence           int64      // 毫秒内序列
	lastTimestamp      int64      // 上次生成ID的时间戳
	lock               sync.Mutex // 锁
}

// New 创建一个实例化对象
func New(dataCenterId int64, workerId int64) Sequence {
	var s = Sequence{}
	// 开始时间戳，默认为 2010-01-01 00:00:00
	s.startTimestamp = 1262275200000
	// 机器ID所占的位数
	s.workerIdBits = 5
	// 数据标识ID所占的位数
	s.dataCenterIdBits = 5
	// 支持的最大机器ID，最大是31
	s.maxWorkerId = -1 ^ (-1 << s.workerIdBits)
	// 支持的最大机房ID，最大是 31
	s.maxDataCenterId = -1 ^ (-1 << s.dataCenterIdBits)
	// 序列在ID中占的位数
	s.sequenceBits = 12
	// 机器ID向左移12位
	s.workerIdShift = s.sequenceBits
	// 机房ID向左移17位
	s.dataCenterIdShift = s.sequenceBits + s.workerIdBits
	// 时间截向左移22位
	s.timestampLeftShift = s.sequenceBits + s.workerIdBits + s.dataCenterIdBits
	// 生成序列的掩码最大值，最大为4095
	s.sequenceMask = -1 ^ (-1 << s.sequenceBits)
	if workerId < 0 || workerId > s.maxWorkerId {
		panic(errors.New(fmt.Sprintf("Worker ID can't be greater than %d or less than 0", s.maxWorkerId)))
	}
	if dataCenterId < 0 || dataCenterId > s.maxDataCenterId {
		panic(errors.New(fmt.Sprintf("DataCenter ID can't be greater than %d or less than 0", s.maxDataCenterId)))
	}
	s.workerId = workerId
	s.dataCenterId = dataCenterId
	// 毫秒内序列(0~4095)
	s.sequence = 0
	// 上次生成 ID 的时间戳
	s.lastTimestamp = -1
	return s
}

// NextId 生成ID，注意此方法已经通过加锁来保证线程安全
func (s *Sequence) NextId() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	timestamp := time.Now().UnixMilli()
	// 如果当前时间小于上一次 ID 生成的时间戳，说明发生时钟回拨，为保证ID不重复抛出异常。
	if timestamp < s.lastTimestamp {
		panic(errors.New(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp)))
	}
	if s.lastTimestamp == timestamp {
		// 同一时间生成的，则序号+1
		s.sequence = (s.sequence + 1) & s.sequenceMask
		// 毫秒内序列溢出：超过最大值
		if s.sequence == 0 {
			// 阻塞到下一个毫秒，获得新的时间戳
			timestamp = time.Now().UnixMilli()
			for timestamp <= s.lastTimestamp {
				timestamp = time.Now().UnixMilli()
			}
		}
	} else {
		// 时间戳改变，序列重置
		s.sequence = 0
	}
	// 保存本次的时间戳
	s.lastTimestamp = timestamp
	// 移位并通过或运算拼到一起
	return ((timestamp - s.startTimestamp) << s.timestampLeftShift) | (s.dataCenterId << s.dataCenterIdShift) | (s.workerId << s.workerIdShift) | s.sequence
}
