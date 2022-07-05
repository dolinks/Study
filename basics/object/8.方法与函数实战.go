package main

import "fmt"

type Point struct {
	x int
	y int
	f func(int, int) (int, int) //函数类型,函数的地址,切换函数,切换行为
}

func main1n() {
	p := &Point{10, 10, func(i int, i2 int) (int, int) {
		return i * 2, i2 * 2
	}}
	p.x, p.y = p.f(p.x, p.y)
	fmt.Println(p)
	p.f = func(i int, i2 int) (int, int) {
		return i + 101, i2 + 101
	}
	p.x, p.y = p.f(p.x, p.y)
	fmt.Println(p)
}

type Dog struct {
	name string
	age  int
}

func (d *Dog) older() {
	d.age++
}
func (d *Dog) show() {
	fmt.Println(d.name, "------->", d.age)
}
func main1x() {
	var d Dog = Dog{"雪雪", 1} //构造初始化
	f1 := (*Dog).show        //方法表达式
	f2 := (*Dog).older       //方法表达式
	fmt.Printf("%T,%T\n", f1, f2)
	f2(&d)
	f1(&d)
	fmt.Println(d)
}

func main() {
	var d Dog = Dog{"雪雪", 1} //构造初始化
	f1 := d.show             //保存函数地址,保存原来结构的数据
	f2 := d.older            //保存地址,修改数据
	f2()
	fmt.Println(d)
	f1()
	fmt.Println(d)
}
