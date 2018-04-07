package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入字符串：")
    str, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("read string,err:", err)
        return
    }
    fmt.Printf("read str succ,ret:%s\n", str)
}