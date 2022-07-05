package main

/*
使用ioutil.readfile 一次性将文件读取
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "D:\\goProject\\src\\Study\\basics\\file\\test.txt"
	File, err := os.Open(path) // 打开文件
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(File)
	fmt.Printf("%T,%v\n", File, File)
	defer File.Close() //关闭文件

	reader := bufio.NewReader(File) //文件句柄
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF { //文件结束
			break
		}
		fmt.Println(str)
	}
}
