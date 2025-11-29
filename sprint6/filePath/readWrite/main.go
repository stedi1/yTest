package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const content = "Тестовый файл для саписи/чтения. Привет мир!"

func main() {
	// создаем директорию в текущей папке
	err := os.Mkdir("clients", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}
	fileName := filepath.Join("clients", "README")
	// создаём файл с текстом
	err = os.WriteFile(fileName, []byte(content), 0755)
	if err != nil {
		log.Fatal(err)
	}
	// читаем файл
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// сравниваем записанное и прочитанное
	if string(data) != content {
		log.Fatal("что-то пошло не так")
	}
	fmt.Println("Запись и чтение прошли усешно")
}
