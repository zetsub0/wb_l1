// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

var nums = []int{2, 4, 6, 8, 10}

func main() {
	wg := sync.WaitGroup{}
	for _, v := range nums {
		wg.Add(1)
		go func(x int) {
			fmt.Println(x * x)
			defer wg.Done() // Объявляем о завершении работы горутин

		}(v)

	}
	wg.Wait() // Ждём выполнения горутин
}

//func main() {
//	for _, v := range nums {
//		go func(x int) {
//			fmt.Println(x * x)
//		}(v)
//	}
//	time.Sleep(chan * time.Millisecond)		// Ждём выполнения горутин
//}
