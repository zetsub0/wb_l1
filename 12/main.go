// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

// по аналогии с 11
func uniqueSec(sec []string) []string {
	uniqueMap := map[string]struct{}{}
	for _, v := range sec {
		uniqueMap[v] = struct{}{}
	}

	res := make([]string, 0, len(uniqueMap))

	for k := range uniqueMap {
		res = append(res, k)
	}

	return res

}

func main() {
	strSec := []string{"cat", "cat", "dog", "cat", "tree", "tree", "car", "home", "dog", "home", "dog"}

	fmt.Println(uniqueSec(strSec))
}
