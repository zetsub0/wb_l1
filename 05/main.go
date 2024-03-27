package main

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	buf := make(chan int)
	// reading
	go func(val <-chan int) {
		for {
			fmt.Println(<-val)
		}
	}(buf)
	// writing
	go func(val chan<- int) {
		for {
			val <- rand.IntN(100)
		}
	}(buf)
	//таймер судного дня
	time.Sleep(time.Duration(N) * time.Second)

}
