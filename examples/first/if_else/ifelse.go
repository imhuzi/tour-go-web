package main

import (
	"fmt"
	"io/ioutil"
)

/**
1. if-else 后面的 {} 是强制的
2. if 后面的 { 必须 和if在同一行
*/
func main() {
	r := testIfesle(10)
	println("r: ", r)

	b := bounded(101)
	println("bounded:", b)

	// 读取文件并打印
	readAndPrintFile("/Users/hz/id_rsa.pub")
	readAndPrintFile("/Users/id_rsa.pub")
}

/**
test if-else , if-else 后面的 {} 是强制的
*/
func testIfesle(x int) (b int) {
	if x > 0 { // if 开始的 {  必须 和 if在同一行
		return x
	} else {
		return 0
	}
}

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

/**
读取文件并打印
*/
func readAndPrintFile(fileName string) {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
	// 可以简写如下, contents变量 只在 if {} 中有效
	//if contents, err := ioutil.ReadFile(fileName);  err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(contents)
	//}

	// 如下也是 ok的
	//contents, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(contents)

}
