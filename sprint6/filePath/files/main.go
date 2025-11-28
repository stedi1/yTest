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

func DirInfo(path, shift string) error {
	// получаем список файлов и поддиректорий
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			return err
		}
		if file.IsDir() {
			fmt.Printf("%s%s\n", shift, file.Name())
			// заходим внутрь директории
			err = DirInfo(filepath.Join(path, file.Name()), shift+"  ")
			if err != nil {
				return err
			}
			continue
		}
		// выводим информацию о файле
		fmt.Printf("%s%s %s %d\n", shift, file.Name(),
			fi.ModTime().Format("02.01.06 15:04:05"), fi.Size())
	}
	return nil
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
	if err = DirInfo(curDir, ""); err != nil {
		log.Fatal(err)
	}
}
