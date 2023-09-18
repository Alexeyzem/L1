// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func reverseWord(str string) string {
	strS := strings.Split(str, " ") //разбиваем строку на массив по пробелу
	for i := 0; i < len(strS)/2; i++ {
		strS[i], strS[len(strS)-1-i] = strS[len(strS)-1-i], strS[i] //свапаем слова из начала со словами из конца слайса
	}
	return strings.Join(strS, " ")
}

func main() {
	str := "snow dog sun smart"
	fmt.Println(reverseWord(str))
}
