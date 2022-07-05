package main

import "fmt"

type F struct {
	int
}

func (f F) testT() {
	fmt.Println("类型 *T 方法集包含receiver F 方法")
}
func (f *F) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *F 方法")
}

func main() {
	t1 := F{1}
	t2 := &t1
	fmt.Printf("t2 is :%v\n", t2)
	t2.testT()
	t2.testP()
}
