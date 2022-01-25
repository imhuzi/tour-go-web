package main

import "fmt"

/**
  1. go 语言中的 switch 会自动 break
  2. switch 后面可以没有表达式
*/
func main() {
	println("加:", testSwitch(10, 2, "+"))
	println("减:", testSwitch(10, 2, "-"))
	println("乘:", testSwitch(10, 2, "*"))
	println("除:", testSwitch(10, 2, "/"))
	println("取余:", testSwitch(11, 2, "%"))

	println(grade(0))
	println(grade(60))
	println(grade(77))
	println(grade(88))
	println(grade(91))
	println(grade(100))
	println(grade(101))
}

func testSwitch(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%":
		result = a % b
	default:
		panic("unsuported operator:" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
