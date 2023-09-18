// Разработать программу, которая будет последовательно отправлять значения в канал,
//  а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func foo(arr []int, c chan int) {
	for _, val := range arr { // записываем данные в канал
		c <- val
	}
	close(c)
}
func foo2(c chan int) {
	for v := range c { //вычитываем данные из канала
		fmt.Println(v)
	}
}

func main() {
	n := 0
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < 1000000; i++ {
		arr = append(arr, i)
	}
	c := make(chan int)
	//конкурентно запускаем две функции, записи вычитывания данных из канала
	go foo(arr, c)
	go foo2(c)
	//ждем н секунд и заканчиваем программу.
	time.Sleep(time.Duration(n) * time.Second)
}
