package main

import "fmt"

/*
defer用于注册一个延迟调用(在函数返回之前调用)
典型应用场景是释放资源,比如关闭文件句柄,释放数据库连接
一个函数里有多个defer,则后注册的先执行
defer后可以跟一个func,func内部发生panic,main协程不会exit,其他defer还可以正常执行
defer后不是跟func,而直接跟一条执行语句,则相关变量在注册defer时被拷贝或计算
*/

//例子1
func basic() {
	fmt.Println("A")
	defer fmt.Println(1)
	fmt.Println("B")
	defer fmt.Println(2)
	fmt.Println("C")
}

//例子2
func defer_exe_time() (i int) {
	//如果defer是表达式,在注册defer时变量已经赋值,方法的时候没有赋值
	i = 9
	defer func() {
		fmt.Printf("i=%d\n", i) //打印5 而非9 5先赋给返回值,在执行
	}() //此处括号相当于调用func匿名函数
	defer fmt.Printf("i=%d\n", i)
	return 5
}
func main() {
	basic()
	defer_exe_time()
}
