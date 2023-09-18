// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10, 5} // последовательность чисел
	sum := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock() //блокируем доступ к переменной sum, чтобы в нее не писали одновременно две горутины.
			sum += val * val
			mu.Unlock()
		}(val)
	}
	wg.Wait() //дожидаемся выполнения всех горутин.
	fmt.Print(sum)
}
