package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\W_test.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666) //路径,读写 操作模式 覆盖模式,权限
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	str := "国足最傻逼再加10\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n') //循环按行去读
		fmt.Print(readString)
		if err == io.EOF { //文件结束
			break
		}
	}
	writer.Flush()

}
