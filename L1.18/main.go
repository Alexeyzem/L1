// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//  По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	val int
	sync.Mutex
}

func (c *Counter) Increment() {
	go c.inc() //запускаем в отдельной горутине счетчик
}
func (c *Counter) inc() { //работает до завершения программы
	for {
		c.Lock()
		c.val++
		c.Unlock()
	}
}

func main() {
	var c Counter
	c.Increment()
	time.Sleep(1 * time.Second)
	fmt.Println(c.val)
}
