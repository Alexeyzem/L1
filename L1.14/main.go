// Разработать программу, которая в рантайме способна определить тип переменной:
//  int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = make(chan interface{})
	fmt.Println(reflect.TypeOf(i)) //встроеная в язык функция, возвращает тип переменной.

	//проверяем тип интерфейса, указав возможные варианты базового типа его значения.
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan interface{}:
		fmt.Println("chan", v)

	}

}
