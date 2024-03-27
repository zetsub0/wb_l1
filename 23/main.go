// Удалить i-ый элемент из слайса.

package main

import "fmt"

func main() {
	a := []int{1, 6, 8, 2, 4, 7, 9, 0, 1, 3, 5, 6}

	i := 0

	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}
