package main

import "fmt"

//var age int

var address = "南京"
var country = "中国"

// 两个数字相加
func addNum(x int, y int) int {
	ret := x + y
	return ret
}

func testLocalVar() {
	var name string
	name = "cdc"
	fmt.Println(name)
}

// 在全局中无法使用局部变量
//name = "ctt"

func main() {
	// 调用函数
	//addRet := addNum(10, 5)
	//fmt.Println("相加的和为：", addRet)

	//age = 18
	//fmt.Println(age)

	//testLocalVar()

	// 在函数内部无法使用其他函数内部的变量
	//name = "ctt"

	// 在代码块中定义的局部变量只能在代码块中使用
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//fmt.Println(i) // 脱离for循环，i 变量不生效
	//
	//if j := 10; j < 100 {
	//	fmt.Println(j)
	//}
	//fmt.Println(j) // 脱离if判断，j 变量不生效
	//
	//switch k := 2; k {
	//case 1:
	//	fmt.Println("aa")
	//case 2:
	//	fmt.Println("bb")
	//default:
	//	fmt.Println("cc")
	//}
	//fmt.Println(k) // 脱离switch判断，k 变量不生效


	var address = "上海"
	fmt.Println(address)  // 上海
	fmt.Println(country)  // 中国
}
