package main

import "fmt"

func main() {
	type A struct {
		x int
		y int
	}
	type B struct {
		x int
		y int
	}

	var a A
	var b B
	a = A(b) //结构体是用户单独定义的类型,和其他类型进行转换是需要完全相同的字段(名字,个数和类型)
	fmt.Println(a, b)

}
