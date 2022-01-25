package main

import "fmt"

func main() {
	/**
	常见占位符：
	%s  字符串占位符
	%d  数字占位符
	%v  输出结构体
	%+v
	%#v

	文档: /usr/local/go/src/fmt/doc.go
	 */
	name := "Huzi.Wang"
	age := 32
	// Sprint 返回的字符串，所以 大多数情况下 字符串拼接可以使用这个函数
	str := fmt.Sprintf("Hello, I am %s I am %d years old", name, age)
	println(str)
	// 这个 是直接输出到控制台，一般 用于 debug 打印日志
	fmt.Printf("Hello, I am %s I am %d years old]\n", name, age)

	replaceHolder()

}

func replaceHolder() {
	u := &user {
		Name: "Wang.Hui",
		Age: 32,
	}

	fmt.Printf("v => %v \n", u)
	fmt.Printf("v => %+v \n", u)
	fmt.Printf("v => %#v \n", u)
	fmt.Printf("v => %T \n", u)
}

type user struct {
	Name string
	Age int
}