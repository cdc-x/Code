package main

func main() {

	// 使用变量接收匿名函数
	add := func(x, y int) int {
		return x + y
	}
	add(10, 5)

	// 匿名函数作为立即执行函数使用
	func(x, y int) int {
		return x + y
	}(10, 20)
}
