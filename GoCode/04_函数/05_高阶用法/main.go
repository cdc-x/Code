package main

import "fmt"

// 函数可以作为其他函数的参数
func add(x, y int) int {
	return x + y
}

func f(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数可以作为其他函数的返回值
func sub(x, y int) int {
	return x - y
}

func cal(op string) func(int, int) int {
	switch op {
	case "add":
		return add
	case "sub":
		return sub
	default:
		fmt.Println("操作不合法")
		return nil

	}
}

func main() {
	//ret := f(10, 5, add)
	//fmt.Println(ret)

	f := cal("sub")

	if f != nil{
		ret := f(10, 5)
		fmt.Println(ret)
	}
}
