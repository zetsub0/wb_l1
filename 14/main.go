// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	default:
		fmt.Println(reflect.TypeOf(i))
	}
}

func main() {
	var a interface{} = make(chan string)
	checkType(a)
	a = 12
	checkType(a)
	a = ""
	checkType(a)

}
