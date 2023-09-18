//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection(m1, m2 map[int]struct{}) map[int]struct{} {
	out := make(map[int]struct{}) //выходное множество
	for key := range m1 {
		if _, ok := m2[key]; ok { // если есть одинаковые значения, то записываем в выходное множество
			out[key] = struct{}{}
		}
	}
	return out
}

func main() {
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	m1[2] = struct{}{}
	m2[2] = struct{}{}
	m1[3] = struct{}{}
	m2[6] = struct{}{}
	fmt.Println(intersection(m1, m2))
}
