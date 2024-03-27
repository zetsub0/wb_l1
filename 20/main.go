// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = input[:len(input)-1]
	fmt.Println(input)
	words := strings.Split(input, " ")
	fmt.Println(words)
	fmt.Println(reverseString(words))
}

func reverseString(arr []string) []string {
	result := make([]string, len(arr))
	for i, v := range arr {
		result[len(result)-i-1] = v
	}
	return result
}
