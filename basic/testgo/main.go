package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1
	fmt.Println("01:", n)
	go func() {
		n++
		if n == 2 {
			fmt.Println("02:", n)
			return
		}
		fmt.Println("after return")
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("00003:", n)
}
