package main

import "fmt"

func forDemo1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v \n", i)
	}
}

func forDemo2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v \n", i)
	}
}

func forDemo3() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v \n", i)
		i ++
	}
}

func forDemo4() {
	// 条件永远满足
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v \n", i)
	}

	// 或者直接写一个for关键字
	for {
		fmt.Println("我一直在执行~")
	}
}

func main() {
	forDemo1()
	forDemo2()
	forDemo3()
	forDemo4()
}
