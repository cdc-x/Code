package main

import "fmt"

func recursionErrorDemo() {
	fmt.Println("aaaaa")
	recursionErrorDemo()

}

// 斐波那契数列
func fibDemo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibDemo(n-1) + fibDemo(n-2)

}

// 求阶乘
func factorialDemo(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialDemo(n-1)
}

func main() {
	//recursionErrorDemo()

	//ret := fibDemo(5)
	//fmt.Println("斐波那契的值：", ret)

	ret := factorialDemo(5)
	fmt.Println("阶乘计算结果：", ret)
}
