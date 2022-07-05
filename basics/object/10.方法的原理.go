package main

import "fmt"

type Persons struct {
	age int
}

func (p Persons) howOld() int {
	return p.age
}
func (p *Persons) growUp() {
	p.age += 1
}

func main() {
	//qcrao 是值类型
	qcrao := Persons{age: 18}
	//值类型 调用接受者也是值类型的方法
	fmt.Println(qcrao.howOld())
	//值类型 调用接收者是指针类型的方法
	qcrao.growUp()
	fmt.Println(qcrao.howOld())
	(&qcrao).growUp() //值类型取地址方法(&)
	fmt.Println(qcrao.howOld())
	//--------------------------
	//setfno是指针类型
	setfno := &Persons{age: 100}
	//指针类型 调用接收者是值类型方法
	fmt.Println(setfno.howOld())
	//指针类型 调用接受者也是指针类型的方法
	setfno.growUp()
	fmt.Println(setfno.howOld())
	(*setfno).growUp() //质针对性取值类型(*)
	fmt.Println(setfno.howOld())
}
