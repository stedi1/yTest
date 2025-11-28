package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 1 * time.Second, // максимальное время ожидания ответа будет 1 секунда
	}
	for {
		request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
		if err != nil {
			fmt.Println("Ошибка формирования запроса:", err)
			break
		}
		request.Header.Set("Custom-Header", "John Doe")
		request.Header.Add("Custom-Header", "1234")
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Ошибка отправки запроса:", err)
			break
		}
		// вызов после defer произойдёт при выходе из функции,
		// поэтому Close() с defer можно указать здесь, так как
		// response.Body.Close() нужно вызвать в любом случае
		defer response.Body.Close()
		// читаем тело ответа
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			return
		}
		fmt.Println(string(body))
		// ждем 2 секунды
		time.Sleep(2 * time.Second)
	}
}
