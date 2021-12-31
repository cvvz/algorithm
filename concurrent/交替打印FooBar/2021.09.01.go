package main

import (
	"fmt"
	"sync"
)

type FooBar struct {
	times int
	chFoo chan struct{}
	chBar chan struct{}
}

func main() {
	var wg sync.WaitGroup
	fb := FooBar{
		times: 5,
		chFoo: make(chan struct{}),
		chBar: make(chan struct{}),
	}
	// 🌟🌟wg.Add 最好放在goroutine外部，不然可能还没有Add，程序就跑完了
	wg.Add(2)
	go func() {
		fb.Foo()
		wg.Done()
	}()
	go func() {
		fb.Bar()
		wg.Done()
	}()
	fb.chFoo <- struct{}{}
	wg.Wait()
}

func (fb *FooBar) Foo() {
	for i := 0; i < fb.times; i++ {
		<-fb.chFoo
		fmt.Print("foo ")
		fb.chBar <- struct{}{}
	}
	// 不加会dead lock
	<-fb.chFoo
}

func (fb *FooBar) Bar() {
	for i := 0; i < fb.times; i++ {
		<-fb.chBar
		fmt.Print("bar ")
		fb.chFoo <- struct{}{}
	}
}
