package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var counter int // счетчик

	// открываем файл
	fCounter, err := os.OpenFile("counter.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем о закрытии файла
	defer fCounter.Close()
	b := make([]byte, 16)
	// читаем значение счетчика
	n, err := fCounter.Read(b)
	if n > 0 {
		counter, err = strconv.Atoi(string(b[:n]))
		if err != nil {
			log.Fatal(err)
		}
	}
	// ограничим количество действий
	for i := 0; i < 100; i++ {
		counter++
		s := strconv.Itoa(counter)
		// каждый раз пишем с начала файла
		_, err = fCounter.WriteAt([]byte(s), 0)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("counter =", counter)
}
