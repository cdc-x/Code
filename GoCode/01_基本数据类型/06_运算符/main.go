package main

import "fmt"

func main() {
	a := 5
	b := 2
	c := 3

	// 算术运算符
	fmt.Printf("a+b=%v\n", a+b)  // a+b=7
	fmt.Printf("a-b=%v\n", a-b)  // a-b=3
	fmt.Printf("a*b=%v\n", a*b)  // a*b=10
	fmt.Printf("a/b=%v\n", a/b)  // a/b=2
	fmt.Printf("a%%b=%v\n", a%b) // a%b=1 想要打印%号，可以用%转义，即 %%

	// ++（自增）和 --（自减）在Go语言中是单独的语句，并不是运算符。
	if a > 4 {
		a++
	} else {
		a--
	}
	fmt.Print(a)

	// 关系运算符
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a >= b) // true
	fmt.Println(a < b)  // false
	fmt.Println(a <= b) // false

	// 逻辑运算符
	fmt.Println(a > b && c > b) // true
	fmt.Println(a < b && c > b) // false
	fmt.Println(a > b || c > b) // true
	fmt.Println(a < b || c < b) // false
	fmt.Println(!(a > b)) // false

	// 位运算符
	fmt.Println(a & b)  // 4
	fmt.Println(a | b)  // 7
	fmt.Println(a ^ b)  // 3
	fmt.Println(a >> 1) // 2
	fmt.Println(a << 1) // 10

	// 赋值运算符
	a += 2 // a = a + 2
	a -= 2 // a = a - 2
	a *= 2 // a = a * 2
	a /= 2 // a = a / 2
	a %= 2 // a = a % 2
}
