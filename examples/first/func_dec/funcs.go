package main

import "fmt"

func main() {
	fmt.Println(div(10, 3))
	fmt.Println(div(14, 3))
	fmt.Println(eval(3, 5, "*"))

	if result, err := eval(3, 5, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("result: ", result)
	}
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a + b, nil
	case "*":
		return a + b, nil
	case "/":
		// 忽略调 余数
		a, _ := div(a, b)
		return a, nil
	default:
		return 0, fmt.Errorf("Not Support Operate: %s", op)
	}
}

/**
带余除法: 多个返回参数
*/
func div(a, b int) (int, int) {
	return a / b, a % b
}

/**
  多个返回参数命名
*/
func div2(a, b int) (q int, r int) {
	q = a / b
	r = a % b
	return
}
