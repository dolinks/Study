package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\W_test.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666) //路径,读写 操作模式 覆盖模式,权限
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	str := "国足最傻逼再加10\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		fmt.Print(readString)
		if err != nil {
			break
		}
	}
}
