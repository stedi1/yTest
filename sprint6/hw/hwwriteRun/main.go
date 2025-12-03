package main

import (
	"log"
	"os"
	"time"
)

func main() {
	// открыть(создать) файл run.txt
	runFile, err := os.OpenFile("run.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// закрываем файл
	defer runFile.Close()
	// получаем время в нужном формате "02.01.06 15:04:05"
	time := time.Now().Format("02.01.06 15:04:05\n")
	// записываем строку в файл
	_, err = runFile.WriteString(time)
	if err != nil {
		log.Fatal(err)
	}
}
