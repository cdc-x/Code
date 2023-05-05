package main

import "fmt"

func A() {
	fmt.Println("function A")
}

func B() {
	defer func(){
		err := recover()

		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in function B")
		}
	}()

	panic("panic in function B")
}

func C() {
	fmt.Println("function C")
}

func main() {
	A()
	B()
	C()
}
