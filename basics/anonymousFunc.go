package main

import (
	"fmt"
)

/*
匿名函数
把函数当作变量类型来使用
*/

var sum = func(a, b int) int {
	return a + b
}

type adds func(a, b int) int

func op(f adds, a int) int {
	b := 2 * a
	return f(a, b)
}
func main() {
	fmt.Println(sum(1, 2))
	slc := make([]adds, 0, 5)
	slc = append(slc, sum, sum, sum)
	fmt.Println(slc)
}

type User struct {
	Name string
	//bye   foo                      //bye的类型是foo,而foo代表一种函数类型
	hello func(name string) string //使用匿名函数赖声明struct字段的类型
}

func foo() {

}
