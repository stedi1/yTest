package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileInfo(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Файл %s не существует\n", path)
			return
		}
		log.Fatal(err)
	}
	fmt.Println("Имя:", fileInfo.Name())
	fmt.Println("Размер:", fileInfo.Size())
	fmt.Println("Является директорией:", fileInfo.IsDir())
	fmt.Println("Время:", fileInfo.ModTime().Format("02.01.06 15:04:05"))
}

func DirInfo(path, shift string) (int64, error) {
	var size int64
	// получаем список файлов и поддиректорий
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			return 0, err
		}
		if file.IsDir() {
			fmt.Printf("%s%s\n", shift, file.Name())
			// заходим внутрь директории
			newSize, err := DirInfo(filepath.Join(path, file.Name()), shift+"  ")
			if err != nil {
				return 0, err
			}
			size += newSize
			continue
		}
		// выводим информацию о файле
		fmt.Printf("%s%s %s %d\n", shift, file.Name(),
			fi.ModTime().Format("02.01.06 15:04:05"), fi.Size())
		size += fi.Size()
	}
	return size, nil
}

func main() {
	// подставим несуществующий файл
	FileInfo("file.go")
	// проверим файл main.go
	FileInfo("main.go")

	fmt.Println()
	// меняем текущую директорию
	err := os.Chdir("../..")
	// получаем новую текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Директория:", curDir)
	size, err := DirInfo(curDir, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Размер всех файлов:", size)
}
