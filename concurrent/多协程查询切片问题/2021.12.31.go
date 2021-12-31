package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Printf("start time %s\n", time.Now().String())
	defer func(t string) {
		fmt.Printf("end time %s\n", t)
	}(time.Now().String())

	bigSlice := []int{14213125, 3, 4, 5, 2, 3, 1, 23, 1, 23, 45, 64435, 231, 214324, 323, 123, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 12, 31, 23, 12, 3, 12312, 123, 3412414, 1234, 1243, 123, 412, 2341}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	found := make(chan struct{})

	n := 5
	target := 214324

	step := len(bigSlice) / n
	for i := 0; i < n; i++ {
		go find(ctx, bigSlice[i*step:(i+1)*step], target, found)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Timeout! Not Found")
	case <-found:
		fmt.Println("Found it!")
	}

	cancel()
}

func find(ctx context.Context, nums []int, val int, found chan struct{}) {
	for _, num := range nums {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(500 * time.Millisecond)

		if num == val {
			found <- struct{}{}
			return
		}
	}
}
