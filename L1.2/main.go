// Написать программу, которая конкурентно рассчитает значение квадратов чисел
//  взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{} // чтобы дождаться выполнения всех горутин
	arr := []int{2, 4, 6, 8, 10}
	for _, val := range arr {
		wg.Add(1)
		go func(val int) { //запускаем новую горутину
			defer wg.Done()
			fmt.Println(val * val) // выводим квадраты
		}(val)
	}
	wg.Wait() //дожидаемся выполнения всех горутин
}
