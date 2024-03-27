// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

type numbers interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func main() {
	a := []int{1, 3, 7, 9, 19, 25, 35, 43, 55}
	fmt.Println(binSearch(a, 11))

	fmt.Println(binSearch(a, 55))
}

func binSearch[N numbers](arr []N, target N) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
