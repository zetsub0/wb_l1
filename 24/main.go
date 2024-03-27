// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"24/point"
	"fmt"
)

func main() {
	p1 := *point.NewPoint(1, 1)
	p2 := *point.NewPoint(6, 4)
	fmt.Println(point.Distance(p1, p2))
}
