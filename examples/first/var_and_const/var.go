package main

import "fmt"

// 1. 变量 定义 使用 var 关键字
// 2. 函数外定义变量(包内部变量), 必须要有 var 关键字, 并且函数外定义的变量 未使用不会报错
// 3. 函数内定义的变量 作用域只是在函数内
// 4. go 有类型推断，在定义变量的时候可以省略 类型，
// go 没有全局变量
var ss = "Hello Variable"

// 多个变量定义, 可以只写 一个 var, 多个变量放到()里面
var (
	aa = "Multi Variable aa"
	bb = "Multi Variable bb"
	cc = "Multi Variable cc"
)

func main() {
	// 只定义不初始化, var 变量名 类型
	variableZeroValue()
	// 定义并初始化, var 变量名 类型 = value
	variableInitValue()
	// 定义变量时可以省略类型, var 变量名 = value
	variableTypeOmit()
	// 局部变量定义简写： 变量名 := value
	variableShorter()

	fmt.Println(ss)
}

/**
  定义变量，不进行初始化,
  1 数字类的变量默认是 0,
  2. string 默认值是空串
*/
func variableZeroValue() {
	var a int
	var f float32
	var s string
	fmt.Printf("a:%d, f: %f, b:%s \n", a, f, s)
}

// 定义变量并初始化
func variableInitValue() {
	var a, b int = 3, 4
	// var a, b = 3, 4 // 可以省略调 类型
	var str = "Init Value"
	fmt.Printf("%d,%d,%s\n", a, b, str)
}

// 定义变量可以省略类型，go 语言支持类型推断
func variableTypeOmit() {
	// 可以省略调 类型
	var a, b, bo, str = 3, 4, true, "String deduction"
	fmt.Printf("variableTypeOmit: %d, %d, %v,%s\n", a, b, bo, str)
}

// 变量定义 简单写法，只能用在局部变量
func variableShorter() {
	// 局部变量定义: 可以省略调 类型 和 var
	a, b, bo, str := 3, 4, true, "String deduction"
	// 变量赋值
	b = 5
	fmt.Printf("variableShorter: %d,%d,%v,%s\n", a, b, bo, str)
}
