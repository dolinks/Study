package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\W_test.txt"
	File, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666) //路径,操作模式,权限
	if err != nil {
		fmt.Println(err)
		return
	}
	defer File.Close()
	str := "国足是太监\r\n"
	writer := bufio.NewWriter(File)
	for i := 0; i < 20; i++ {
		writer.WriteString(str)
	}
	writer.Flush() //缓存写入文件
}
