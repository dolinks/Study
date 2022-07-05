package main

import (
	"Study/basics/object/SelfInt"
	"fmt"
)

func main() {
	var i SelfInt.Myint = 100
	fmt.Println("main", i)
	i.Adder()
	i.Print()
}
