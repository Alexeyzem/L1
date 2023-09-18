// Разработать программу, которая проверяет, что все символы в строке уникальные
//  (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

package main

import (
	"fmt"
	"strings"
)

func Check(str string) bool {
	str = strings.ToLower(str)   //для регистронезависимости
	m := make(map[rune]struct{}) //храним уже встретившиеся символы
	for _, r := range str {
		if _, ok := m[r]; !ok {
			m[r] = struct{}{} //такого символа не было - добавляем в мапу
		} else {
			return false //символ уже встречался
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	if Check(str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
