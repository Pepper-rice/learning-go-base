package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建一个空的 map
	m := make(map[string]int)

	// 赋值
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))

	// 获取
	v1 := m["k2"]
	fmt.Println("v1:", v1)

	// 长度
	fmt.Println(len(m))

	// 删除
	delete(m, "k2")
	fmt.Println(m)

	// 获取值的第二个返回值，表示是否存在该健
	_, exist := m["k2"]
	fmt.Println(exist)

	n := map[string]string{"fool": "a", "bar": "b"}
	fmt.Println(n)

	notexist := n["a"]
	fmt.Println(notexist)
}
