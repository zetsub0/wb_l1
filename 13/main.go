// Поменять местами два числа без создания временной переменной

package main

import "fmt"

func main() {
	a, b := 1, 2
	b, a = a, b
	fmt.Println(a, b)
}

//func main() {
//	a, b := 10, 20
//	a = a * b
//	b = a / b
//	a = a / b
//	fmt.Println(a, b)
//}

//func main() {sff
//	a, b := 10, 20
//	a = a - b
//	b = a + b
//	a = -a + b
//	fmt.Println(a, b)
//}

//func main() {
//	a, b := 10, 20
//	a = a ^ b
//	b = b ^ a
//	a = a ^ b
//	fmt.Println(a, b)
//}
