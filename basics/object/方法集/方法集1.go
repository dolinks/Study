package main

import "fmt"

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型T 方法集包含全部 receiver T 方法")
}

func main() {
	t1 := T{1}
	fmt.Printf("t1 is:%v\n", t1)
	t1.test()
}
