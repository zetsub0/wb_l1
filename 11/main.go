// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func intersectionSet(set1, set2 []int) []int {
	valMap := map[int]int{}
	l := 0

	// соберем в эту мапу все числа (по одному разу) из первого слайса
	for _, v := range set1 {
		if valMap[v] == 0 {
			valMap[v] = 1
		}
	}

	// добавим в мапу из второго слайса только совпадающие элементы
	for _, v := range set2 {
		if valMap[v] == 1 {
			valMap[v] = 2
			l++
		}
	}

	// сразу создадим слайс нужной длины
	res := make([]int, 0, l)
	for v := range valMap {
		if valMap[v] == 2 {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 4, 5, 5, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 15, 16}
	set2 := []int{12, 12, 12, 13, 14, 15, 16, 2, 19, 20, 18, 16, 17}
	set3 := intersectionSet(set1, set2)
	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(set3)
}
