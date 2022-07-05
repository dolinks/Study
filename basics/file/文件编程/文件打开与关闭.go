package main

import (
	"fmt"
	"os"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\test.txt"
	File, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(File)
	fmt.Printf("%T,%v\n", File, File)
	defer File.Close() //关闭文件
}
