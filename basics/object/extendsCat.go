package main

/*
构造函数嵌套
*/
type Cat struct {
	name string
	age  int
}

type BlackCat struct {
	Cat   //嵌入式继承
	color string
}

//构造函数
func NewCat(name string, age int) *Cat {
	return &Cat{name: name, age: age}
}
func NewBlackCat(name string, age int) *BlackCat {
	mycat := &BlackCat{*NewCat(name, age), "黑色"}

	return mycat
}

func main() {

}
