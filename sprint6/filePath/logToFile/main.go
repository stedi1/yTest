package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("console.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// сохраняем идентификатор текущего вывода
	stdout := os.Stdout
	// присваиваем os.Stdout идентификатор открытого файла
	os.Stdout = f
	// строка должна записаться в файл
	fmt.Println("Это вывод в файл console.txt")

	// возвращаем обратно вывод в консоль
	os.Stdout = stdout
	// строка выведется в консоль
	fmt.Println("Это вывод в терминал")
}
