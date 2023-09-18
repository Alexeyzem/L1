//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	a, b = b, a //внутренний способ языка
	fmt.Println(a, b)
	a = 10
	b = 20
	a = a ^ b
	b = a ^ b
	a = a ^ b //с помощью xor - исключающее или
	fmt.Println(a, b)
	a, b = 100, 200
	a = a + b
	b = a - b
	a = a - b // с помощью суммы этих чисел
	fmt.Println(a, b)
}
