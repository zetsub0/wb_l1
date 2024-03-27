// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	s := "а"
	tmp := []rune(s)
	tmpRes := make([]rune, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		tmpRes = append(tmpRes, tmp[i])
	}
	result := string(tmpRes)
	fmt.Println(result)
}
