package main

import (
	"fmt"
	"image/color"
	"math"
)

/*
方法声明
和函数类似,只是在函数名字前边多了一个参数,把这个方法绑定到这个参数对应的类型上
可以将方法绑定到任何类型上.(数字,字符串,slice,map,函数)定义附加行为.
同一个包下的任何类型都可以声明方法.只要不是指针类型也不是接口类型
*/
type Point struct {
	X, Y float64
}

//内嵌了一个Point类型以提供字段X,Y.Point类型的方法.
type ColoredPoint struct {
	Point
	Color      color.RGBA
	color.RGBA //结构体类型可以拥有多个匿名字段
}

//普通的函数(包级别的函数)
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//绑定到Point类型的方法(Point.Distance)
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//指针接收者方法
//主调函数会复制每一个实参变量,如果函数需要更新一个变量,或者一个实参太大避免赋值整个实参,
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

//不允许本身是指针的类型进行方法声明
type p *int

// func (p) f(){} 编译器报错:非法的接收者类型

func main() {
	/*通过*Point能够调用(*Point).ScaleBy方法,编译器会对变量进行&p的隐式转换.只有变量才允许这么做.
	 */
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) //{2,4}

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) //{2,4}

	p1 := Point{1, 2}
	(&p1).ScaleBy(2)
	fmt.Println(p1) //{2,4}

	//不能够对一个不能取地址的Point接收者参数调用*Point方法,因为无法获取临时变量的地址
	//Point{1, 2}.ScaleBy(2)

	/*方法变量表达式
	赋予一个方法变量,它是一个函数,把方法绑定到一个接收者上只需要提供实参,不需要提供接收者就能调用
	*/
	distanceFromP := p.Distance //方法变量,Distance是方法()
	distanceFromP(p1)
}
