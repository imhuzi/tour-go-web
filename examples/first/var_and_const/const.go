package main

import (
	"fmt"
	"math"
)

// 常量在 Go 中， 也就是 constant。 它们在编译时被创建， 只能是数字、字符串或布尔 值；const x = 42
const x = 5

func main() {
	defEnumsByIota()
	testConst()
	println("x=", x)
}

/**
test const
*/
func testConst() {
	const fileName = "test.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	println("fileName:"+fileName, c)
	const (
		fileName2 = "test.txt"
		a1, b1    = 3, 4
	)
	c = int(math.Sqrt(a1*a1 + b1*b1))
	println("fileName2:"+fileName2, c)

}

/**
通过 iota 自增值 定义枚举：第一个 iota 表示为 0，因此 a 等于 0，当 iota 再次在新的一行使用时，它的值增加 了 1，因此 b 的值是 1。
*/
func defEnumsByIota() {
	//const (
	//	a = 0
	//	b = 1
	//)
	// 和上面的一个效果
	//const (
	//	a = iota
	//	b = iota
	//)
	// 简写: 和上面的效果
	const (
		a = iota
		c
		d
	)

	fmt.Printf("const: %d, %d \n", a, d)
	// b, kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	println(b, kb, mb, gb, tb, pb)
}
