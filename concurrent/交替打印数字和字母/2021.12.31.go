package main

import "fmt"

func main() {
	num := make(chan int)
	done := make(chan struct{})
	pause := make(chan struct{})

	go func() {
		for i := 1; i <= 26; i += 2 {
			fmt.Printf("%d%d", i, i+1)
			num <- i
			<-pause
		}
		fmt.Print(2728)
		close(done)
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case i := <-num:
				fmt.Printf("%c%c", 'A'+i-1, 'A'+i)
				pause <- struct{}{}
			}
		}
	}()

	<-done
}
