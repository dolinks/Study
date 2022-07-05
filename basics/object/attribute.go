package main

import "fmt"

type person struct {
	Name   string
	Age    int
	Score  [5]float64        //数组
	lover  *person           //指针
	mydict map[string]string //字典
	shoes  []int             //切片
}

func main() {
	var p1 *person
	fmt.Println(p1)
	if p1 == nil {
		fmt.Println("没有初始化")
	}
	p1 = new(person)
	p1.shoes = make([]int, 10, 10)
	//slice,map,make分配内存并初始化
	if p1.shoes == nil {
		fmt.Println("shoes没有初始化")
	}
	p1.Name = "你大爷"
	p2 := new(person)
	p2.Name = "你大娘"
	fmt.Println(p1, p2)
}
