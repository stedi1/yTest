package main

import "fmt"

// хеш-функция
func hash(in string) [4]byte {
	// сумма всех символов
	var result int32

	// проходимся по каждому символу сообщения
	for i := 0; i < len(in); i++ {
		// преобразуем символ в число и добавляем к результату
		result += int32(in[i])
	}

	// преобразуем число в массив байт:
	// используем побитовые операции сдвига
	// чтобы получить первые 8 bit, потом следующие 8 бит и так далее
	return [4]byte{
		byte(0xff & result),
		byte(0xff & (result >> 8)),
		byte(0xff & (result >> 16)),
		byte(0xff & (result >> 24))}
}

func main() {
	fmt.Println(hash("Hello, hashed world!"))
}
