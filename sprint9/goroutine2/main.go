package main

import (
	"fmt"
	"time"
)

func main() {
	data := [100]int{}

	fmt.Println(data)

	// вставьте сюда код с запуском горутин
	// ...
	for i := range 100 {
		go func(i int) {
			data[i] = i
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// проверим как заполнен массив

	fmt.Println(data)
}
