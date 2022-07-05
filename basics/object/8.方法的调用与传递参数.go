package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Getarea(radius float64) float64 { //参数遵循副本机制
	return math.Pi * radius *
		radius
}

func main() {
	var c Circle
	c.radius = 100
	res := c.area()
	fmt.Println(res)
	fmt.Println(c.Getarea(1000))

}
