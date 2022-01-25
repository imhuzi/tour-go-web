package main

import "fmt"

func main() {
	// 直接初始化了 3个元素的切片
	s1 := []int{1,2,3}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	// 创建了包含3个元素，容量为4的切片
	s2 := make([]int,3,4)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	// s2 目前 是[0,0,0]  append (追加) 一个元素， 变成什么？
	s2 = append(s2,7) // 没有超出容量限制，不会发生扩容
	// [0,0,0,7]
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	// 在追加一个 元素, 触发扩容, 扩容策略？
	s2 = append(s2,8)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	// 只传一个参数 表示 创建一个含有4个元素，容量也是4的切片
	s3 := make([]int,4)
	// 等价于 s3 := make([]int,4,4)
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
}