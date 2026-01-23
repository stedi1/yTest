package main

import (
	"fmt"
	"time"
)

func tik() {
	for i := 0; i < 30; i++ {
		fmt.Print(".")
		time.Sleep(time.Millisecond)
	}
}

func tok() {
	for i := 0; i < 30; i++ {
		fmt.Print("o")
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go tik()
	go tok()

	for _, v := range []string{".", "o"} {
		go func(s string) {
			for i := 0; i < 30; i++ {
				fmt.Print(s)
				time.Sleep(time.Millisecond)
			}
		}(v)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("\nконец main()")
}
