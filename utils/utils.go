package Utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

//ReadFile Text文件中的列表
func ReadFile(fileName string) ([]byte, error) {
	// 判断文件是否存在
	isFile, err := PathExists(fileName)
	if isFile == false {
		fmt.Println(fileName, "文件不存在")
		return nil, err
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	return b, err
}

//DelFile 删除文件
func DelFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("删除文件失败,文件不存在", fileName)
	} else {
		fmt.Println(fileName, "文件删除成功")
	}
}

//PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
