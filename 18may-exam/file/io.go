package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("文件不存在")
		return
	}

	n, err := io.WriteString(file, "zhougang 3213\n")
	if err != nil {
		fmt.Println("写入失败")
	}
	fmt.Printf("写入%d个字节\n", n)

	//  第二种
	var b1 = []byte("heer is word")
	err = ioutil.WriteFile("test2.txt", b1, 0666)
	if err != nil {
		fmt.Println("不能写入test2.txt")
	}

	// 第三种
	file3, err := os.Create("test3.txt")
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	defer file3.Close()
	n3, err := file3.Write([]byte("here is file3\n"))
	if err != nil {
		fmt.Println("写入file3失败")
	}
	fmt.Printf("写入file3 %d 个字节\n", n3)

	n33, _ := file3.Seek(0, os.SEEK_END)
	_, err = file3.WriteAt([]byte("here is file3 append\n"), n33)

	if err != nil {
		fmt.Println("写入file3失败")
	}
	fmt.Printf("写入file3 %d 个字节\n", n33)

	// 第四种
	file4, err := os.Create("test4.txt")
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	defer file4.Close()

	bf := bufio.NewWriter(file4)
	_, err = bf.WriteString("here is file4\n")
	_, err = bf.WriteString("here is file4 append \n")
	if err != nil {
		fmt.Println("写入file4失败")
	}
	bf.Flush()
}
