package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		str      string = "GoLang"
		java_str        = "java"
		c_str           = "c"
	)

	//全て大文字に変換する
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToUpper(java_str))
	fmt.Println(strings.ToUpper(c_str))
}
