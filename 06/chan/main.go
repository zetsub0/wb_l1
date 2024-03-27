// Реализовать все возможные способы остановки выполнения горутины.

package main

import "fmt"

// читаем до тех пор, пока не обнаружим что-то в quit
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// после десяти 11 итераций записываем в канал quit для остановки горутины
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
