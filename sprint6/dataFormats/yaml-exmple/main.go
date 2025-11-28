package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Artist содержит данные об артисте
type Artist struct {
	ID    int      `yaml:"id"`
	Name  string   `yaml:"name"`
	Genre string   `yaml:"genre"`
	Songs []string `yaml:"songs"`
}

func main() {
	artist := Artist{
		ID:    1,
		Name:  "30 seconds to Mars",
		Genre: "rock",
		Songs: []string{
			`The Kill`,
			`A Beautiful Lie`,
			`Attack`,
			`Live Like A Dream`,
		},
	}

	// сериализуем данные из переменной artist в слайс байт
	// & означает, что мы передаём указатель на artist
	yamlData, err := yaml.Marshal(&artist)
	if err != nil {
		fmt.Println("ошибка при сериализации в yaml:", err)
		return
	}

	var newArtist Artist
	err = yaml.Unmarshal(yamlData, &newArtist)
	if err != nil {
		fmt.Println("ошибка при десериализации:", err)
		return
	}

	// выводим результат
	fmt.Println(string(yamlData))
	fmt.Println("Данные после десериализации ===>")
	fmt.Println("ID:", newArtist.ID)
	fmt.Println("Name:", newArtist.Name)
	fmt.Println("Genre:", newArtist.Genre)
	fmt.Printf("Songs: %v\n", newArtist.Songs)
}
