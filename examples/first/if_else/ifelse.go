package main

func main() {
	r := testIfesle(10)
	println("r: ", r)
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
