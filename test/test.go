package main

import (
	"fmt"
	"net/http"
)

func mainHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Получен запрос")
	// curTime := time.Now().Format("02.01.2026 15:04:05")
	// res.Write([]byte(curTime))

	// s := fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s", req.Host, req.URL.Path, req.Method)
	// res.Write([]byte(s))

	s := fmt.Sprintf("%s/%s", req.Host, req.URL.Path)
	res.Write([]byte(s))
}

func main() {
	fmt.Println("Запускаем сервер")
	http.HandleFunc(`/`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}
