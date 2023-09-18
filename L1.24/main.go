// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64 // инкапсулированные параметры
}

func (p *Point) Init(arg ...float64) { //инициализация
	if len(arg) == 0 {
		p.x, p.y = 0, 0
		return
	}
	if len(arg) == 1 {
		p.x, p.y = arg[0], 0
		return
	}
	p.x = arg[0]
	p.y = arg[1]
	return
}
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}
func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func FindDistance(p1 Point, p2 Point) float64 { //используем библиотеку math для возведения в квадрат и извлечения корня.
	return math.Sqrt(math.Pow(p1.GetX()-p2.GetX(), 2) + math.Pow(p1.GetY()-p2.GetY(), 2))

}
func main() {
	p1 := Point{}
	p1.SetX(10)
	p1.SetY(20)
	var p2 Point
	p2.Init(5)
	fmt.Println(FindDistance(p1, p2))
}
