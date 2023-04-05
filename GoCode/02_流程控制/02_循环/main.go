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

func forDemo5() {
	s := "hello南京"

	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r) // 104(h) 101(e) 108(l) 108(l) 111(o) 21335(南) 20140(京)
	}
}

func main() {
	forDemo1()
	forDemo2()
	forDemo3()
	forDemo4()
	forDemo5()
}
