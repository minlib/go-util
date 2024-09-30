package sequence

import (
	"sync"
)

// Counter Struct
type Counter struct {
	value int64
	lock  sync.Mutex
}

// NewCounter 创建一个实例化对象
func NewCounter(value int64) *Counter {
	return &Counter{
		value: value,
	}
}

// SetValue 初始化值
func (s *Counter) SetValue(value int64) {
	s.value = value
}

// NextId 生成自增计数器ID，已通过加锁来保证线程安全
func (s *Counter) NextId() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value++
	return s.value
}
