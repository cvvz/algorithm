package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(1000)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("%d ", num)
		}
	}()

	wg.Wait()
}
