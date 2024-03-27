// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"time"
)

const N = 10 // число фиббоначи

func fibonacci(ctx context.Context) {
	x, y := 0, 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(x)
			return
		default:
			x, y = y, x+y
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// подстраиваем таймер под число фиббоначи и останавливаем горутину
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go fibonacci(ctx)

	time.Sleep(time.Millisecond * N * 100)
	cancel()

	time.Sleep(time.Millisecond * 100)
}
