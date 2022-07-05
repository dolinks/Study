package main

/*
闭包
i 自由变量一直存在
*/
import "fmt"

//例子1
func sub() func() {
	i := 10
	fmt.Printf("%p\n", &i)
	b := func() {
		fmt.Printf("i addr %p\n", &i)
		i--
		fmt.Println(i)
	}
	return b
}

//例子2
func add(base int) func(int) int {
	return func(i int) int {
		fmt.Printf("base addr %p\n", &base)
		base += i
		return base
	}
}

func main() {
	//f := sub()
	//f()
	//f()
	//fmt.Println()

	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(10)
	fmt.Println(tmp2(1), tmp2(2))

}
