package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
1. 返回值类型写在最后面
2. 可返回多个值
3. 函数可作为参数
4. 没有默认参数，可选惨，没有重载
*/
func main() {
	fmt.Println(div(10, 3))
	fmt.Println(div(14, 3))
	fmt.Println(eval(3, 5, "*"))

	if result, err := eval(3, 5, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("result: ", result)
	}

	// 将函数作为参数
	fmt.Println(apply(pow, 2, 3))

	// 匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 3))

	// 可变参数
	fmt.Printf("sumArgs: %d \n", sumArgs(2, 3, 3, 3))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
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

// 函数参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d) \n", opName, a, b)
	return op(a, b)
}

// 可变参数
func sumArgs(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
