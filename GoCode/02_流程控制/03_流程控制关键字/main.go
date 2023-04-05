package main

import "fmt"

func breakDemo1() {

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}
}

func breakDemo2() {
	i := 2
	switch i {
	case 1:
		fmt.Println("等于1")
		break
	case 2:
		fmt.Println("等于2")
		break // 满足条件后直接退出 switch。break 后所有代码都不会执行
		fallthrough
	case 3:
		fmt.Println("等于3")
		break
	default:
		fmt.Println("不关心")
		break
	}
}

func breakDemo3() {

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 5 {
				break // 只能退出里层的j循环，外层的i循环还是会继续执行
			}

			fmt.Printf("j: %v\n", j)
		}

		fmt.Printf("i: %v\n", i)
	}

breakKey:
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 5 {
				break breakKey // 可以退出两层循环
			}

			fmt.Printf("j: %v\n", j)
		}

		fmt.Printf("i: %v\n", i)
	}
}

func continueDemo1() {

	for i := 0; i <= 10; i++ {
		if i == 2 {
			continue // 遇到i等于2时，直接跳过，开始下一次循环。continue后的代码不会执行
		}

		fmt.Println(i)
	}
}

func continueDemo2() {
forLoop1:
	for i := 0; i <= 5; i++ {
		//forLoop2:
		for j := 0; j <= 5; j++ {
			if i == 2 && j == 2 {
				continue forLoop1 // 当i和j都等于2时，跳过当前的j循环和i循环，直接开始循环i=3
			}

			fmt.Printf("i: %v, j: %v\n", i, j)
		}
	}
}

func gotoDemo1() {

	exitFlag := false

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 2 {
				exitFlag = true
				break
			}
			fmt.Printf("i: %v, j: %v\n", i, j)
		}

		if exitFlag {
			break
		}
	}
}

func gotoDemo2() {

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 2 {
				goto exitFlag
			}
			fmt.Printf("i: %v, j: %v\n", i, j)
		}
	}

exitFlag:
	fmt.Println("end~")
}

func main() {
	//breakDemo1()
	//breakDemo2()
	//breakDemo3()
	//continueDemo1()
	//continueDemo2()
	//gotoDemo1()
	gotoDemo2()
}
