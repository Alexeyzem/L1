//Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	M map[int]int
	sync.RWMutex
}

func NewMap() *myMap {
	return &myMap{
		M: make(map[int]int),
	}
}
func (my *myMap) Record(key, val int) {
	my.Lock() // лочим, чтобы другие горутины не имели право доступа к мапе.
	my.M[key] = val
	my.Unlock()
}
func main() {

	m := NewMap()
	arr := []int{1, 3, 2, 4, 5, 6, 7, 2, 1, 5, 8, 9}
	wg := sync.WaitGroup{}
	for i, val := range arr {
		wg.Add(1)
		go func(i, val int) {
			defer wg.Done()
			m.Record(i, val)
		}(i, val)
	}
	wg.Wait()
	fmt.Println(m.M)
}
