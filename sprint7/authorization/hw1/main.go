package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// FindPassword возвращает имя из users у кого хеш совпадает с хешем пароля
func FindPassword(users map[string]string, password string) string {
	bytesHash := sha256.Sum256([]byte(password))
	str := hex.EncodeToString(bytesHash[:])

	for user, hash := range users {
		if hash == str {
			return user
		}
	}

	return ""
}

func main() {
	// база данных сотрудников
	var users = map[string]string{
		"Аристова О.А.":  "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
		"Бородин А.Я.":   "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4",
		"Бородина Ж.А.":  "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5",
		"Коваленко Е.Н.": "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
		"Короткова И.И.": "165408fae624cfbfc32054d7a7c18ee087af98dab3cb33fab9ed2e8cf5d5ff5c",
		"Матвеев Н.Г.":   "15b7634d4ffada54ec0ff74711355cc83c2357f9c7bcff1d1c6950662c76c3d8",
	}

	fmt.Println("Кто виноват?", FindPassword(users, "g56-rte-4a9-3gh"))
}
