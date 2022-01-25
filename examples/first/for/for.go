package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
for 的 3中形式
1. for init ; condition ; post { } ← 和 C 的 for 一样
2. for condition { } ← 和 while 一样
3. for { } ← 死循环
*/
func main() {
	// 和其他语言一样的 for
	println("sum from 1 to 100:", testFor(100))

	fmt.Println(
		covertToBin(5),  // 101
		covertToBin(13), // 1101
		covertToBin(123433333),
		covertToBin(0), // 0
	)

	printFile("/Users/hz/abc.txt")
}

func testFor(count int) (total int) {
	sum := 0
	for i := 1; i <= count; i++ {
		sum += i
	}
	return sum
}

func covertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	// 和 while 效果一样
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
