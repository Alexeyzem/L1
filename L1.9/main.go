//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func firstWorker(arr []int, downstream chan int) {
	for _, val := range arr {
		downstream <- val //записываем числа из массива в канал, после чего его закрываем
	}
	close(downstream)
}
func secondWorker(upstream, downstream chan int) {
	for val := range upstream {
		downstream <- val * val // до тех пор, пока канал открыт, вычитываем из него числа и передаем дальше их квадрат
	}
	close(downstream)
}

func printSqrt(upstream chan int) {
	for val := range upstream {
		fmt.Println(val) // до тех пор пока открыт канал от предыдущего воркера, вычитываем данные и печатаем их.
	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 19}
	first := make(chan int)
	second := make(chan int)
	go firstWorker(arr, first)
	go secondWorker(first, second)
	printSqrt(second)

}
