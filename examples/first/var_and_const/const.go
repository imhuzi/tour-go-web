package main

import "fmt"

const x = 5

// 常量在 Go 中， 也就是 constant。 它们在编译时被创建， 只能是数字、字符串或布尔 值；const x = 42
func main() {
	defEnumsByIota()

	println("x=", x)
}

/**
通过 iota 定义枚举：第一个 iota 表示为 0，因此 a 等于 0，当 iota 再次在新的一行使用时，它的值增加 了 1，因此 b 的值是 1。
*/
func defEnumsByIota() {
	//const (
	//	a = iota
	//	b = iota
	//)
	// 简写: 和上面的效果
	const (
		a = iota
		b
		c
		d
	)

	fmt.Printf("const: %d, %d \n", a, d)
}
