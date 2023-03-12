package main

import "fmt"

func main() {
	// 初始化布尔类似变量
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false

	b5 := true
	b6 := false

	fmt.Printf("b1: %v\n", b1) // b1: true
	fmt.Printf("b2: %v\n", b2) // b2: false
	fmt.Printf("b3: %v\n", b3) // b3: true
	fmt.Printf("b4: %v\n", b4) // b4: false
	fmt.Printf("b5: %v\n", b5) // b5: true
	fmt.Printf("b6: %v\n", b6) // b6: false

	// 布尔值用在条件判断中
	age := 18
	ok := age >= 18
	if ok {
		fmt.Println("你已经成年")
	} else {
		fmt.Println("你还未成年")
	}

	// 布尔值用在循环语句中
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}

	// 布尔值用在逻辑表达式中
	age1 := 18
	gender := "男"

	if age1 >= 18 && gender == "男" {
		fmt.Println("你是成年男子")
	}
}
