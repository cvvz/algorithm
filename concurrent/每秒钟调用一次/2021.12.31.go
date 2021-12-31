// 写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()

			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
