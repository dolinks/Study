package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	path1 := "D:\\goProject\\src\\Study\\basics\\file\\test.txt"
	path2 := "D:\\goProject\\src\\Study\\basics\\file\\test_copy.txt"
	file, err := ioutil.ReadFile(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(path2, file, 0666) //可读可写,可执行
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝成功")
}
