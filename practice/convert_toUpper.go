package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		str      string = "GoLang"
		java_str        = "Java"
		c_str           = "C"
	)

	//全て大文字に変換する
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToUpper(java_str))
	fmt.Println(strings.ToUpper(c_str))

	//全て小文字に変換する
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToLower(java_str))
	fmt.Println(strings.ToLower(c_str))
}
