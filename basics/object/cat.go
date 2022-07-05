package main

import "fmt"

type cat struct {
	name  string
	age   int
	color string
	eat   string
	voice string
}

func Newcat() *cat {
	return &cat{name: "校花", age: 3, color: "黄色", eat: "老鼠", voice: "瞄"}
}
func Newcatfish() *cat {
	return &cat{name: "校草", age: 3, color: "黄色", eat: "鱼", voice: "瞄"}
}

func main1() {
	var mycat cat
	mycat.eat = "鱼"
	mycat.voice = "瞄"
	mycat.name = "鲁花"
	mycat.age = 3
	mycat.color = "花猫"
	fmt.Println(mycat)
}
func main() {
	mycat1 := Newcat()
	fmt.Println(mycat1)
	mycat2 := Newcatfish()
	fmt.Println(mycat2)
}
