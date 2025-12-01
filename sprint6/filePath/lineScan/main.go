package main

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines читает файл построчно и выполняет действие для каждой строки
func ReadLines(fileName string, process func(int, string)) error {
	// открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// создаем сканер для чтения файла
	scanner := bufio.NewScanner(file)

	// читаем файл построчно
	lineNum := 1
	for scanner.Scan() {
		// вызываем функцию обработки для текущей строки
		process(lineNum, scanner.Text())
		lineNum++
	}
	// проверяем ошибки во время чтения
	return scanner.Err()
}

func main() {
	// функция для вывода каждой строки
	process := func(num int, line string) {
		fmt.Printf("%02d: %s\n", num, line)
	}
	// читаем файл построчно
	ReadLines("main.go", process)
}
