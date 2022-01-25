package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
内置的数据类型(11大类)： Boolean types, Numeric types,String types,Array types,
Slice types,Struct types,Pointer types,Function types,Interface types,Map types,Channel types

1. bool,
2. (u)int: 32位, (u)int8, (u)int16,(u)int32,(u)int64,
3. float32, float64,
4. 复数：complex64(实数部分32位+虚数部分32位), complex128
5. byte  alias for uint8
6. string, 是所有8位字节字符集的集合
7. uintptr 是 指针类型(unsafe.Pointer) 本质是一个 整数类型
8. rune(32位) 本质是一个 int32   -> alias for int32
*/
//1. go 中没有隐式类型转换，所有运算中涉及类型转换的必须强制转换
//2. 特殊值: nil 是 预定义的标识符， 指针(pointer), 管道(channel), 函数(func), 接口(interface), map, or slice 等类型的默认值(或者说0值)
//3. 默认值 int是0, float 是 0.0, string 是空字符串

func main() {
	var a complex64 = 21.0
	println(a)
	euler()
	// 强制类型转换示例: math.Sqrt()三角形勾股定理
	triangle()

	// rune
	var ru rune = rune(65)
	fmt.Println(ru)
}

// 欧拉公式(复数)
func euler() {
	// 定义 一个复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 欧拉公式
	fmt.Printf("euler:%v \n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("euler:%f \n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("euler:%0.3f \n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换 三角形勾股定理
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Printf("求三角形第三边: %d \n", c)
}
