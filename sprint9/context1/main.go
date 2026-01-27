package main

import (
	"context"
	"fmt"
)

// context.WithValue
func action(ctx context.Context) {
	fmt.Println(ctx.Value("func"), ctx.Value("uid"), ctx.Value("user"))
}

func main() {
	ctx := context.WithValue(context.Background(), "func", "main")

	action(context.WithValue(ctx, "uid", 701))
}

// context.WithTimeout
/* var wg sync.WaitGroup

func asteriks(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// выведет context deadline exceeded
			fmt.Println(ctx.Err())
			return
		case <-time.After(400 * time.Millisecond):
			fmt.Print("*")
		}
	}
}

func dots(ctx context.Context) {
	defer wg.Done()
	// указываем ctx в качестве родительского контекста
	ctxAsteriks, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	go asteriks(ctxAsteriks)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(300 * time.Millisecond):
			fmt.Print(".")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg.Add(2)
	go dots(ctx)
	wg.Wait()
} */

// context.WithCancel
/* func process(ctx context.Context) {
	fmt.Println("Начало работы")
	for i := 0; i < 10; i++ {
		if ctx.Err() != nil {
			fmt.Println("Прервали работу")
			return
		}
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Конец работы")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// через 2 секунды вызываем cancel для отмены операции
	time.AfterFunc(2*time.Second, cancel)
	process(ctx)
} */
