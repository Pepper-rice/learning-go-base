package main

import (
	"fmt"
)

func main() {
	var a = "inital"

	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	// 声明未初始化，变量会初始化为：零值
	// 零值的意思是数值是0，字符串是空字符串，指针是 nil
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写
	f := "short"
	fmt.Println(f)
}
