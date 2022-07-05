package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\test.txt"
	content, err := ioutil.ReadFile(path) //content []byte类型数组
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
