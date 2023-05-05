package main

import "fmt"

func demo1() {
	fmt.Println("aaa")
}

func demo2() int {
	age := 18
	return age
}

// 返回多个值
func demo3() (int, string, string) {
	age := 18
	name := "cdc"
	address := "南京"

	return age, name, address
}

// 指定返回的内容，变量不需要声明
func demo4() (name string, age int) {
	age = 18
	name = "cdc"

	return name, age
}

// 指定返回内容，但是未使用
func demo5() (name string, age int) {
	age = 18
	name = "cdc"
	return //  等价于 return name, age
}

// 覆盖命令返回值
func demo6() (name string, age int) {
	a := 18
	b := "cdc"

	return b, a
}

func main() {
	// 函数无返回值时不能使用变量接收
	// demo1() doesn't return a value
	//ret1 := demo1()

	//ret := demo2()
	//fmt.Println(ret)

	//ret1, ret2, ret3 := demo3()
	//fmt.Println("age: ", ret1)
	//fmt.Println("name: ", ret2)
	//fmt.Println("address: ", ret3)

	//ret1, ret2 := demo4()
	//fmt.Println("name: ", ret1)
	//fmt.Println("age: ", ret2)

	//ret1, ret2 := demo5()
	//fmt.Println("name: ", ret1)
	//fmt.Println("age: ", ret2)

	ret1, ret2 := demo6()
	fmt.Println("name: ", ret1)
	fmt.Println("age: ", ret2)

}
