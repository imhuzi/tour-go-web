package main

import "fmt"

func main() {
	// 初始化并赋值
	a1 :=[3]int{1,2,3}
	fmt.Printf("a1: %v, len: %d, cap: %d \n", a1, len(a1), cap(a1))
	// 初始化一个三个元素的数组 ，所有元素都是 0
	var a2 [3] int
	fmt.Printf("a2: %v, len: %d, cap: %d \n", a2, len(a2), cap(a2))
	// 数组不支持 append 操作
	//a1 = append(a1, a2)
}