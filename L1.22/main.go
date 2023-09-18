//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

//исользуем встроенный в go пакет big для работы с большими числами.

func sum(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}
func subtraction(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, big.NewInt(0).Neg(b))
}
func div(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}
func multiply(a *big.Int, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func main() {
	var a, b big.Int
	fmt.Scan(&a, &b)
	fmt.Println(sum(&a, &b))
	fmt.Println(subtraction(&a, &b))
	fmt.Println(div(&a, &b))
	fmt.Println(multiply(&a, &b))

}
