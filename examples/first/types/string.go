package main

import (
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
  strings 包下相关方法
*/
func stringsPackage() {
	strings.Compare("aa", "aa")
}
