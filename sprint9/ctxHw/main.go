package main

import (
	"context"
	"fmt"
	"time"
)

func tick(ctx context.Context) {
	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 20; i++ {
		// для корректной проверки решения оставьте
		// вызов fmt.Print() на этом же месте
		fmt.Print(i, " ")
		// здесь нужен select
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// продолжаем выполнение цикла
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	tick(ctx)
}
