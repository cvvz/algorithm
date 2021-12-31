// GO里面map如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	sp := NewSp()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("1", 1)
		sp.Out("3", 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sp.Rd("3", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("3", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)
		sp.Rd("1", 1*time.Second)

	}()

	wg.Wait()
}

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type spImpl struct {
	sync.Mutex
	entrys map[string]*entry
}

type entry struct {
	value interface{}
	ch    chan struct{}
}

func NewSp() sp {
	return &spImpl{
		entrys: map[string]*entry{},
	}
}

func (s *spImpl) Out(key string, val interface{}) {
	s.Lock()
	defer s.Unlock()

	e, exist := s.entrys[key]
	if !exist {
		s.entrys[key] = &entry{
			value: val,
		}
		return
	}
	e.value = val // 这里是对entry内的value赋值，所以entry一定要是指针类型

	if e.ch != nil {
		close(e.ch)
		e.ch = nil
	}
}

func (s *spImpl) Rd(key string, timeout time.Duration) interface{} {
	s.Lock()
	e, exist := s.entrys[key]
	if exist && e.value != nil {
		fmt.Printf("get key %s, value %v\n", key, e.value)
		s.Unlock()
		return e.value
	}
	if !exist {
		s.entrys[key] = &entry{
			ch: make(chan struct{}),
		}
	}
	ch := s.entrys[key].ch
	s.Unlock()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("time out! key %s\n", key)
		return nil
	case <-ch:
	}

	s.Lock()
	defer s.Unlock()

	fmt.Printf("get key %s, value %v before timeout\n", key, s.entrys[key].value)
	return s.entrys[key].value
}
