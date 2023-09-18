//Pеализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		mid := (len(arr) - 1) / 2
		pivot := arr[mid] // выбрали опорный элемент
		var left []int
		var right []int
		for i, value := range arr { //проход по массиву. Все элементы меньше опорного - влево, большие - вправо
			if i != mid {
				if value <= pivot {
					left = append(left, value)
				} else {
					right = append(right, value)
				}
			}
		}
		left = QuickSort(left) //рекурсивно сортируем левую и правую части
		right = QuickSort(right)
		var output []int
		output = append(append(left, arr[mid]), right...) //соединяем
		return output
	}
}
func main() {
	fmt.Println(QuickSort([]int{1, 7, 8, 2, 90, 2, 4, 0, 8}))

	sort.Ints([]int{1, 7, 8, 2, 90, 2, 4, 0, 8}) //встроенный метод в язык
}
