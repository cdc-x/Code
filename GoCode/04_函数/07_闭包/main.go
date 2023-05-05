package main

import (
	"fmt"
	"strings"
)

func add(x int) func(int) int {

	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	//f := add(10)
	//fmt.Println(f(10)) // 20
	//fmt.Println(f(20)) // 40
	//fmt.Println(f(30)) // 70
	//fmt.Println(f(40)) // 110

	//txtJudgeFunc := makeSuffix(".txt")
	//jpgJudgeFunc := makeSuffix(".jpg")
	//
	//ret1 := txtJudgeFunc("test")
	//ret2 := jpgJudgeFunc("test")
	//
	//fmt.Println(ret1) // test.txt
	//fmt.Println(ret2) // test.jpg

	//f1, f2 := calc(10)
	//fmt.Println(f1(1))  // 11
	//fmt.Println(f1(2))  // 13
	//fmt.Println(f2(3))  // 10
	//fmt.Println(f2(4))  // 6

}
