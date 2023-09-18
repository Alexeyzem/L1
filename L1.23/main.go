//Удалить i-ый элемент из слайса
package main

import "fmt"

func deleteI(arr []int, pos int) {
	arr = append(arr[:pos], arr[pos+1:]...) //если важно сохранить порядок.
}
func deleteIWihout(arr []int, pos int) []int {
	arr[0], arr[pos] = arr[pos], arr[0] //если порядок не важен, выполняется быстрее.
	return arr[1:]
}

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	deleteI(arr, 2)
	fmt.Println(arr)
	arr = deleteIWihout(arr, 1)
	fmt.Println(arr)
}
