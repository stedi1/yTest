package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// получаем имя файла из URL
	dwnFile := r.URL.Path[len("/download/"):]
	if dwnFile == "" {
		http.Error(w, "не указано имя файла", http.StatusBadRequest)
		return
	}
	root, err := os.OpenRoot("downloads")
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	// открываем файл
	dwnFile = filepath.Base(dwnFile)
	file, err := root.Open(dwnFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("файл %s не найден", dwnFile), http.StatusNotFound)
		return
	}
	defer file.Close()

	// устанавливаем заголовки для скачивания
	w.Header().Set("Content-Disposition", "attachment; filename="+dwnFile)
	// можно указать нужный тип файла
	w.Header().Set("Content-Type", "application/octet-stream")
	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "внутреняя ошибка", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", fmt.Sprint(fileStat.Size()))

	// копируем содержимое файла в ответ
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	// получаем файл из формы
	file, handler, err := r.FormFile("attach")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	// закрываем файл
	defer file.Close()

	root, err := os.OpenRoot("uploads")
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	// создаём файл с таким же именем
	dst, err := root.Create(handler.Filename)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// копируем содержимое загруженного файла в новый файл
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}
}
