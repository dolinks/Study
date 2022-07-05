package main

import (
	"fmt"
	"math"
)

/*
函数声明
func name(parameter-list)(result-list){
	body
}
当两个函数拥有相同的形参列表和返回列表时,认为这两个函数的类型或签名是相同的.而形参和返回值的名字不会影响到函数类型
实参是按照值传递的,所以修改函数的形参变量不会影响到实参.
如果提供的实参包含引用类型,如指针,slice,map,函数,或者通道,name当函数使用形参变量是就有可能会简介地修改实参变量
如果看到函数声明没有函数体,说明函数使用了除了GO以外的语言实现.这样的声明定义了该函数的签名
package math
func Sin(x float64) float64 //使用汇编语言实现
*/

func bypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

/*
匿名函数
在func关键字后面没有函数的名称,称作匿名函数
这些隐藏的变量引用就是我们把函数归类为引用类型饿而且函数变量无法进行比较的原因.
函数变量类似于使用闭包方法实现的变量.函数变量通常称为闭包
*/
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

/*
变长函数
变长函数被调用时候可以有可变的参数个数.fmt.Printf
在参数列表最后的类型名称之前使用省略号"..."表示声明一个变长函数,调用这个函数的时候可以传递该类型任意数目的参数
变长函数的类型和一个带有普通slice参数的函数类型不相同
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(bypot(3, 4))
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(sum(1, 2, 3, 4))
}
