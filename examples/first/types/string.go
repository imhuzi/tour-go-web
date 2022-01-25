package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 字符串支持两种格式：""引起来的，和 `` 引起来的
	// `` 通常来输出大段文字
	println("Hello GO! \"I Come\"")
	// `` 支持换行
	println(`He said: "Hello GO" 
    我要出去了`)
	// 长度计算
	stringLen()
	// 多行字符串
	multiLineStr()
	// strings 包下相关方法测试
	stringsPackage()
}

/**
string 长度 计算
*/
func stringLen() {
	// 字节长度
	println(len("你好")) // 6
	// 字符数量
	println(utf8.RuneCountInString("你好"))   // 2
	println(utf8.RuneCountInString("你好aa")) // 4
}

/**
多行字符串
*/
func multiLineStr() {
	s := "Starting part" +
		"Ending part \n"

	s2 := `Starting part 
		Ending part`
	fmt.Println(s, s2)
}

/**
  strings 包下相关方法
*/
func stringsPackage() {
	r := strings.Compare("aa", "aa")
	str, str2 := "aa", "aa"
	r2 := strings.Compare(str, str2)
	rr := strings.EqualFold(str, str2)
	// strings.Compare Result: 0, 0, EqualFold=true
	fmt.Printf("strings.Compare Result: %d, %d, EqualFold=%v", r, r2, rr)
}
