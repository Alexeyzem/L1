//Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func BinarySearch(item int, list []int) bool {
	sort.Ints(list) //если н авход может прийти любой слайс
	low := 0
	high := len(list) - 1
	for low <= high {
		var mid int
		mid = (low + high) / 2 //среднее значение
		guess := list[mid]
		if guess == item {
			return true
		} else if guess > item { //если значение больше того, что нужно, то опускаем верхнюю границу
			high = mid - 1
		} else {
			low = mid + 1 //иначе поднимаем нижнюю границу
		}
	}
	return false //ничего не нашли
}
func main() {
	//возвращает true or false в зависимости есть или нет
	fmt.Println(BinarySearch(6, []int{3, 6, 2, 8, 5})) // работает с любыми массивами, но нужно время на сортировку внутри

	//вовзращает позицию элемента, если не нашел, то вернет позицию, куда можно вставить
	fmt.Println(sort.SearchInts([]int{1, 2, 3, 5, 6, 7, 8}, 9)) //встроенный поиск в го, работает с отсортированными массивами
}
