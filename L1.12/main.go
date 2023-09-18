//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

//проход по слайсу строк и запись значений как ключ в мапе.
func ArrToSet(arr []string) map[string]struct{} {
	out := make(map[string]struct{})
	for _, v := range arr {
		out[v] = struct{}{}
	}
	return out
}
func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := ArrToSet(arr)
	fmt.Println(m)
}
