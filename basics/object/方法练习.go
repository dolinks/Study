package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}
func (p *Person) older() {
	p.Age++
}
func (p Person) add(a, b int) int {
	return a + b
}

//类型实现了String这个方法,那么 fmt.Println 默认会调用这个变量的String()进行输出
func (p Person) String() string {
	return "person" + p.Name + strconv.Itoa(p.Age)
}

//小写本地包访问,方法首字母 大写,可在本地包其他包访问
func main() {
	p := Person{}
	p.Name = "你大爷"
	p.Age = 90
	p.speak()
	p.older()
	fmt.Println(p.add(1, 2))
}
