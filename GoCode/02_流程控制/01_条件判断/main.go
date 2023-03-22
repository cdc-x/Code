package main

import "fmt"

//func ifDemo1() {
//
//	age := 20
//
//	if age >= 18 {
//		fmt.Println("你已经成年了")
//	}
//
//	fmt.Printf("程序运行结束")
//}
//
//func ifDemo2() {
//	if age := 20; age >= 18 {
//		fmt.Println("你已经成年了")
//	}
//
//	fmt.Printf("程序运行结束")
//
//	fmt.Println(age) // undefined: age
//}
//
//func ifDemo3() {
//	var i = 1
//	if i { // 编译失败
//		fmt.Println("here")
//	}
//	fmt.Printf("程序运行结束")
//}
//
//func ifDemo4() {
//	var num int
//
//	fmt.Println("输入一个数字：")
//	fmt.Scan(&num)
//	fmt.Printf("s 的值是：%v \n", num)
//
//	if num%2 == 0 {
//		fmt.Println("偶数")
//	} else {
//		fmt.Print("奇数")
//	}
//}
//
//func ifDemo5() {
//	// 生成一个 1~100 随机数
//	rand.Seed(time.Now().UnixNano())
//	age := rand.Intn(100)
//
//	if age >= 18 {
//		fmt.Println("你已经成年")
//	} else {
//		fmt.Println("你还未成年")
//	}
//}
//
//func ifDemo6() {
//	// 生成一个 1~100 随机数
//	rand.Seed(time.Now().UnixNano())
//
//	if age := rand.Intn(100); age >= 18 {
//		fmt.Println("你是成年人")
//	} else {
//		fmt.Println("你还未成年")
//	}
//
//	//fmt.Println(age)  age 作用域只在布尔表达式中
//}
//
//func ifDemo7() {
//	// 生成一个 1~100 随机数
//	rand.Seed(time.Now().UnixNano())
//	score := rand.Intn(100)
//
//	if score >= 60 && score <= 70 {
//		fmt.Println("C")
//	} else if score > 70 && score <= 90 {
//		fmt.Println("B")
//	} else {
//		fmt.Println("A")
//	}
//}
//
//func ifDemo8() {
//	// 生成一个 1~100 随机数
//	rand.Seed(time.Now().UnixNano())
//
//	if score := rand.Intn(100); score >= 60 && score <= 70 {
//		fmt.Println("C")
//	} else if score > 70 && score <= 90 {
//		fmt.Println("B")
//	} else {
//		fmt.Println("A")
//	}
//
//	//fmt.Println(score)  score 作用域只在布尔表达式中
//}
//
//func ifDemo9() {
//	var c string
//	fmt.Println("输入一个字符：")
//	fmt.Scan(&c)
//
//	if c == "S" {
//		fmt.Println("输入第二个字符：")
//		fmt.Scan(&c)
//		if c == "a" {
//			fmt.Println("Saturday")
//		} else if c == "u" {
//			fmt.Println("Sunday")
//		} else {
//			fmt.Println("输入错误")
//		}
//	} else if c == "F" {
//		fmt.Println("Friday")
//	} else if c == "M" {
//		fmt.Println("Monday")
//	} else if c == "T" {
//		fmt.Println("输入第二个字符：")
//		fmt.Scan(&c)
//		if c == "u" {
//			fmt.Println("Tuesday")
//		} else if c == "h" {
//			fmt.Println("Thursday")
//		} else {
//			fmt.Println("输入错误")
//		}
//	} else if c == "W" {
//		fmt.Println("Wednesday")
//	} else {
//		fmt.Println("输入错误")
//	}
//}
//
//func ifDemo10() {
//	rand.Seed(time.Now().UnixNano())
//
//	a := rand.Intn(100)
//	b := rand.Intn(100)
//	c := rand.Intn(100)
//
//	fmt.Printf("随机生成数字a的值为：%v \n", a)
//	fmt.Printf("随机生成数字b的值为：%v \n", b)
//	fmt.Printf("随机生成数字c的值为：%v \n", c)
//
//	if a > b {
//		if a > c {
//			fmt.Println("a最大")
//		}
//	} else {
//		if b > c {
//			fmt.Println("b最大")
//		} else {
//			fmt.Println("c最大")
//		}
//	}
//}

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func switchDemo2() {
	switch finger := 3; finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	//fmt.Println(finger) // undefined: finger
}

func switchDemo3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
		fallthrough
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func main() {
	//ifDemo1()
	//ifDemo2()
	//ifDemo3()
	//ifDemo4()
	//ifDemo4()
	//ifDemo5()
	//ifDemo6()
	//ifDemo7()
	//ifDemo8()
	//ifDemo9()
	//ifDemo10()
	//switchDemo1()
	//switchDemo2()
	//switchDemo3()
	//switchDemo4()
	switchDemo5()
}
