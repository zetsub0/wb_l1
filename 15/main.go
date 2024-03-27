// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
//
// func main() {
//   someFunc()
// }
//

package main

import (
	"fmt"
	"strings"
)

// 1. v - большая строка, и если на неё будет ссылаться глобальная переменная
//    то ГК её не трогает. ее => трата ресурсов впустую.
//    Решение: сделать justString логальной и сделать для someFunc возвращаемое значение;
//			   использовать copy, вместо инциализации подстроки
// 2. justString является подстрокой v, и это будет адекватно работать в контексте ASCII,
//    но если в v будет содержаться, условно, русские символы, то justString сформируется
//    некорректно. Для решения этой проблемы можно использоваться руны.

func main() {
	justString := someFunc()
	fmt.Println(len(justString), "\n", justString)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	res := []rune(v)
	return string(res)
}

func createHugeString(size int) string {
	return strings.Repeat("а", size)
}

// таким образом мы не теряем никаких символов и не оставляем лишнего мусора в памяти
