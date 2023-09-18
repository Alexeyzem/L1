//Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func sleep1(t time.Duration) {
	fmt.Println("start sleep1")
	<-time.After(t) //Ожидает прошествия t времени после чего в канал записывает нынешнее время
	fmt.Println("stop sleep1")
}
func sleep2(t time.Duration) {
	a := time.Now().Add(t) //Время которое будет через t времени
	fmt.Println("start sleep2")
	for time.Now().Sub(a) < 0 {
		// пока ныненшнее время меньше чем время установленное в переменную а крутимся в цикле
	}
	fmt.Println("stop sleep2")
}

func sleep3(t time.Duration) {
	fmt.Println("start sleep3")
	<-time.Tick(t) //Совершает один тик в t времени после чего в канал записывает нынешнее время
	fmt.Println("stop sleep3")
}
func main() {
	sleep1(time.Second)
	sleep2(2 * time.Second)
	sleep3(time.Second)
}
