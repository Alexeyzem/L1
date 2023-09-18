// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func reverse(str string) string {
	var out []rune
	for _, r := range str { //проходимся по массиву с помощью рэндж так как учитываем символы юникода
		out = append([]rune{r}, out...) //записываем в выходной массив переворачивая
	}
	return string(out)
}

func main() {
	str := "«главрыба"
	fmt.Println(reverse(str))

}
