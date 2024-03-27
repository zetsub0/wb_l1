// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

// получаем значение, нужный бит и значение для этого бита
func setBit(num, bit, val int64) int64 {
	mask := int64(1 << bit)
	// если значение = 1, то делаем побайтовую дизъюнкцию
	if val == 1 {
		num |= mask

		// если значение = 0, то делаем побайтовую конъюнкцию
	} else if val == 0 {
		num &= ^mask
	} else {
		fmt.Println("wrong value")
		return 0
	}
	return num
}

func main() {
	var num, bit, val int64
	fmt.Scan(&num, &bit, &val)

	fmt.Println(setBit(num, bit, val))
}
