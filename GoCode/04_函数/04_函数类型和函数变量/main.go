package main

import "fmt"

type cal func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {

	var func1 cal
	var func2 cal

	func1 = add  // add函数满足cal类型，所以可以将add赋值给cal类型变量func1
	ret1 := func1(10, 5)
	fmt.Println(ret1)

	func2 = sub  // sub函数满足cal类型，所以可以将sub赋值给cal类型变量func2
	ret2 := func2(10, 5)
	fmt.Println(ret2)
}
