package main

import (
	"encoding/base64"
	"net/http"
)

func main() {
	// создаем новый http-запрос
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	login := "Alice2000"
	password := "qwerty1234"
	// соединяем логин и пароль через двоеточие и получаем формат base64
	basicAuthData := base64.StdEncoding.EncodeToString([]byte(login + ":" + password))
	// простовляем заголовок Authorization
	req.Header.Set("Authorization", "Basic "+basicAuthData)
}
