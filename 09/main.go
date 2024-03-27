// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sendChan := make(chan int)
	valChan := make(chan int, len(x))

	// пытаемся читать
	go func(receive, collect chan int) {
		for {
			num := <-receive
			collect <- num * num
		}
	}(sendChan, valChan)

	// пытаемся писать
	for _, v := range x {
		sendChan <- v
	}

	// выводим
	for i := 0; i < len(x); i++ {
		fmt.Println(<-valChan)
	}
}
