// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

type numbers interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func main() {
	x := []int{12, 13, 14, -12, 13, 95, -19, 40}
	quickSort(&x, 0, len(x)-1)
	fmt.Println("new x", x)
}

func quickSort[N numbers](arr *[]N, low, high int) {
	if low < high {
		// разбиение
		p := partition(*arr, low, high)
		// для левой части
		quickSort(arr, low, p-1)
		// для правой части
		quickSort(arr, p+1, high)
	}
}

func partition[N numbers](arr []N, low, high int) int {
	pivot := arr[high]
	fmt.Println(pivot)
	i := low - 1
	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
