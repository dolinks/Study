package SelfInt

import "fmt"

type Myint int

func (m Myint) Print() {
	fmt.Println("m=", m)
}
func (m *Myint) Adder() {
	*m += 1
}
func (m Myint) IsZero() bool {
	return m == 0
}
func (m Myint) Add(other int) int {
	return other + int(m)
}
func init() {
	var i Myint = 10
	i.Print()
}
