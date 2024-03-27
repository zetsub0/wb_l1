// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import "fmt"

//func main() {
//	nums := []int{02, 04, 6, 8, 10}
//	squares := make(chan int)
//
//	go sum(nums, squares)
//
//	fmt.Println(<-squares)
//}

func sum(arr []int, out chan int) {
	res := 0
	for _, v := range arr {
		res += v * v
	}
	out <- res
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println(5 / 2)
	go sum(nums[:len(nums)/2], ch1) // Сумма первой половины слайса
	go sum(nums[len(nums)/2:], ch2) // Сумма второй половины слайса

	fmt.Println(<-ch1 + <-ch2)
}
