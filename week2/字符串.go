package main

import (
	"fmt"
	"strings"
)

// strings.HasPrefix(s, prefix string) bool 前缀

// strings.HasSuffix(s, suffix string) bool 后缀

func main() {

	var s = "1234567"

	a := strings.HasPrefix(s, "123")

	b := strings.HasSuffix(s, "567")

	fmt.Println(a)
	fmt.Print(b)
}

//https://www.kancloud.cn/kancloud/the-way-to-go/72462
