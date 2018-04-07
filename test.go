package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "http://learn.szmihua.com/login/weixin/callback"
	estr := url.encode()
	fmt.println(estr)
}
