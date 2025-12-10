package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// для примера возьмём токен, подписанный при помощи секретного ключа secretKey
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFsaWNlIiwicm9sZXMiOlsicmVhZGVyIiwid3JpdGVyIl19.Fppwzfqcl8oxEqcQmDesbzJ77DUDdsqLhh3emlKmv3s"

	// создаём секретный ключ для подписи. Это может быть любое слово.
	// Думайте об этом как о секретном слове в банке.
	secret := []byte("my_secret_key")

	// создаем payload
	claims := jwt.MapClaims{
		"login": "Alice2000",
		"roles": []string{"reader", "writer"},
	}

	// создаем jwt и указываем алгоритм хеширования, payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// получаем подписанный токен
	sighnedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		fmt.Printf("failed to sign jwt: %s\n", err)
	}

	// второй аргумент — функция, которая просто возвращает секретный ключ
	// чтобы было понятней, мысленно вместо функции подставьте возвращаемое значение
	jwtToken, err = jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Printf("failed to parse token: %s\n", err)
		return
	}
	if !jwtToken.Valid {
		fmt.Println("token is invalid")
		return
	}

	// приводим поле Claims к типу jwt.MapClaims
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("failed to type assertion of jwt.MapClaims")
		return
	}
	// Так как jwt.Claims — словарь вида map[string]inteface{}, используем синтаксис получения
	// значения по ключу. Получаем значение ключа "login" и "roles"
	loginRaw := res["login"]
	rolesRaw := res["roles"]
	// loginRaw — интерфейс, так как тип значения в jwt.Claims — интерфейс.
	// Чтобы получить строку, нужно снова сделать приведение типа к строке.
	login, ok := loginRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return
	}
	// обратите внимание, что при создании ролей мы указывали тип []string,
	// однако тут приводим к []interface{}
	// так происходит, потому что json не строго типизированный,
	// из-за чего при парсинге нельзя точно определить тип слайса.
	roles, ok := rolesRaw.([]interface{})
	if !ok {
		fmt.Printf("failed to typecast to []interface")
		return
	}

	fmt.Println("Result token:", sighnedToken)

	// выводим login и roles
	fmt.Println(login)
	fmt.Println(roles)
}
