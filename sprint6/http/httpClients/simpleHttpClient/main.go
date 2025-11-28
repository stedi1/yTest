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

	for i := 0; i < 2; i++ {
		response, err := client.Post("http://localhost:8080/", "text/plain", nil)
		if err != nil {
			fmt.Println(err)
			break
		}
		// читам тело ответа
		body, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			break
		}
		fmt.Println(string(body))
		// ждем 3 секунды
		time.Sleep(3 * time.Second)
	}
	for {
		request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/time", nil)
		if err != nil {
			fmt.Println("Ошибка формирования запроса:", err)
			break
		}
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
