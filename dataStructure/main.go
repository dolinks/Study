package main

import (
	"Study/dataStructure/ArrayList"
	"fmt"
)

func main() {
	List := ArrayList.NewArrayList()
	List.Appand("a1")
	List.Appand("b2")
	List.Appand("c3")
	List.Insert(1, "d5")
	fmt.Println(List)
}
