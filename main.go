package main

import (
	"flag"
	"fmt"
	"runtime/Utils"
	"strings"
)

func main() {
	// 获取命令行的参数
	txtFilepath := flag.String("filename", "C:\\shanchu\\lianze.txt", "输入需要执行删除列表的txt文件 ( 确保删除的文件一行一个 )")
	flag.Parse()
	fmt.Println("port:", *txtFilepath)
	ByteContent, err := Utils.ReadFile(strings.Replace(*txtFilepath, "\\", "\\\\", -1))
	if err == nil {
		for _, v := range strings.Split(string(ByteContent), "\n") {
			if v != "" {
				filename := strings.Replace(v, "\\", "\\\\", -1)
				filename = strings.TrimSpace(filename)
				Utils.DelFile(filename)
			}
		}
	}
}

// https://blog.csdn.net/kjfcpua/article/details/18265441
