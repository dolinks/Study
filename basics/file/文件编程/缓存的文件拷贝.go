package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(dst, src string) (writerLength int64, err error) {
	open, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(open)
	defer open.Close()

	//写入
	file, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	return io.Copy(writer, reader)
}

func main() {
	path1 := "D:\\goProject\\src\\Study\\basics\\file\\test.txt"
	path2 := "D:\\goProject\\src\\Study\\basics\\file\\缓存文件copy.txt"

	file, err := copyFile(path2, path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file, "拷贝成功")
}
