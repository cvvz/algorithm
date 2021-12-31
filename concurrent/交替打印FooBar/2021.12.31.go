package main

import "fmt"

func main() {
	PrintFooBar(10)
}

func PrintFooBar(times int) {
	foo := make(chan string)
	bar := make(chan string)
	done := make(chan struct{})

	go func() {
		for i := 0; i < times; i++ {
			fmt.Println(<-foo)
			bar <- "bar"
		}
		close(done)
	}()

	go func() {
		for {
			select {
			case foo <- "foo":
			case b := <-bar:
				fmt.Println(b)
			case <-done:
				return
			}
		}
	}()

	<-done
}
