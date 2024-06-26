// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
//
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

package main

import "fmt"

func main() {
	tmp := []float64{-32, 9, -30.1, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	subSets := make(map[int][]float64)
	_ = subSets

	for _, v := range tmp {
		intV := int(v)
		// intV/10*10 для получения вторых разрядов - 10, 20, 30...
		subSets[intV/10*10] = append(subSets[intV/10*10], v)
	}
	fmt.Println(subSets)
}
