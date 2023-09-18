//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func changeBit(val int64, pos, bit int) int64 {
	if bit < 0 || bit > 1 {
		return val // неккоректное значение бита дали, возвращаем то, что и было
	}
	if pos >= 64 {
		return val // неккоректное значение позиции бита дали, возвращаем то, что и было
	}
	var mask int64 = 1 << pos // маска с 1 на позиции которую нужно изменить
	if bit == 1 {
		return mask | val // изменяем нужную позицию на 1 с помощью битового или
	} else {
		mask = ^mask      //инвертируем маску
		return mask & val // изменяем нужную позицию на 0 с помощью битового не
	}
}

func main() {
	var i int64 = 33
	fmt.Println(changeBit(i, 5, 0))
}
