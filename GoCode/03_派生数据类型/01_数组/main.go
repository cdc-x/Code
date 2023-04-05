package main

import "fmt"

func arrayDemo1() {
	var a [3]int
	var b [4] int

	fmt.Printf("a的数据类型：%T\n", a) // a的数据类型：[3]int
	fmt.Printf("b的数据类型：%T\n", b) // b的数据类型：[4]int

}

func arrayDemo2() {
	var a [3]int
	var s [3]bool

	fmt.Printf("a的类型：%T，a的值：%v\n", a, a) // a的类型：[3]int，a的值：[0 0 0]
	fmt.Printf("s的类型：%T，s的值：%v\n", s, s) // s的类型：[3]bool，s的值：[false false false]
}

func arrayDemo3() {
	var numArray1 = [3]int{1, 2, 3} // 使用指定的初始值完成初始化
	var numArray2 = [3]int{1, 2}    // 指定的值的个数与长度不匹配时，默认用该类型的0值填充
	//var numArray3 = [3]int{1, 2, 3, 4} // 指定的值的个数不能大于数组的长度

	var stringArray = [3]string{"cdc", "tt", "tr"}

	fmt.Printf("numArray1：%v\n", numArray1)     // numArray1：[1 2 3]
	fmt.Printf("numArray2：%v\n", numArray2)     // numArray2：[1 2 0]
	fmt.Printf("stringArray：%v\n", stringArray) // stringArray：[cdc tt tr]

}

func arrayDemo4() {
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}

	fmt.Printf("numArray的类型：%T  numArray的值：%v\n", numArray, numArray)     // numArray的类型：[2]int  numArray的值：[1 2]
	fmt.Printf("cityArray的类型：%T  cityArray的值：%v\n", cityArray, cityArray) // cityArray的类型：[3]string  cityArray的值：[北京 上海 深圳]
}

func arrayDemo5() {
	data := [...]bool{0: true, 5: true}

	fmt.Printf("data的类型：%T  data的值：%v\n", data, data) // data的类型：[6]bool  data的值：[true false false false false true]

}

func arrayDemo6() {
	num := [4]int{1, 2, 3, 4}

	// 通过索引取值
	fmt.Println(num[2]) // 3

	// 通过索引修改值
	num[0] = 100
	fmt.Println(num) // [100 2 3 4]
}

func arrayDemo7() {
	num := [4]int{1, 2, 3, 4}

	for i := 0; i < len(num); i++ {
		fmt.Printf("索引值：%v，对应的元素值：%v\n", i, num[i])
	}
}

func arrayDemo8() {
	num := [4]int{100, 200, 300, 400}

	// 使用两个变量接收时，即获取索引，又获取值
	for index, value := range num {
		fmt.Printf("索引值：%v，对应的元素值：%v\n", index, value)
	}

	// 只用一个变量接收时，只能获取到索引值
	for index := range num {
		fmt.Printf("索引值：%v\n", index)
	}

	// 只获取值
	for _, value := range num {
		fmt.Printf("元素值：%v\n", value)
	}

}

func arrayDemo9() {

	cityArray := [2][3]string{
		{"南京", "上海", "北京"},
		{"深圳", "长沙", "重庆"}, // 多维数组使用列表初始化时，最后一行也必须加上逗号
	}

	// cityArray的类型：[2][3]string  cityArray的值：[[南京 上海 北京] [深圳 长沙 重庆]]
	fmt.Printf("cityArray的类型：%T  cityArray的值：%v\n", cityArray, cityArray)

	// 通过索引取值
	fmt.Println(cityArray[0][0]) // 南京
	fmt.Println(cityArray[0][1]) // 上海
	fmt.Println(cityArray[0][2]) // 北京
}

func arrayDemo10() {
	// 错误示例  编译报错：use of [...] array outside of array literal
	//cityArray := [...][...]string{
	//	{"南京", "上海", "北京"},
	//	{"深圳", "长沙", "重庆"}, // 多维数组使用列表初始化时，最后一行也必须加上逗号
	//}
	//
	//fmt.Println(cityArray)

	// 正确示例
	cityArray := [...][3]string{
		{"南京", "上海", "北京"},
		{"深圳", "长沙", "重庆"}, // 多维数组使用列表初始化时，最后一行也必须加上逗号
	}

	fmt.Println(cityArray) // [[南京 上海 北京] [深圳 长沙 重庆]]

}

func arrayDemo11() {
	cityArray := [...][3]string{
		{"南京", "上海", "北京"},
		{"深圳", "长沙", "重庆"}, // 多维数组使用列表初始化时，最后一行也必须加上逗号
	}

	// 根据长度进行遍历
	for i := 0; i < len(cityArray); i++ {
		inner := cityArray[i] // 针对最外层的数组进行遍历，得到的第二层的每一个数组
		fmt.Printf("内层数组：%v\n", inner)

		// 针对第二层的每个数组进行遍历，得到的是每个元素的值
		for j := 0; j < len(inner); j++ {
			fmt.Printf("城市信息：%v\n", inner[j])
		}
	}

	// 使用 for+range 进行遍历
	for _, inner := range cityArray {
		fmt.Printf("内层数组：%v\n", inner)
		for _, city := range inner {
			fmt.Printf("城市信息：%v\n", city)
		}
	}
}

func arrayDemo12() {
	var numArray = [10]int{1, 2, 3, 4, 5, 100, 200, 300, 400, 500}

	ret := numArray[2:5]
	fmt.Println(ret)
	fmt.Printf("ret的类型：%T\n", ret)  // ret的类型：[]int

	//ret = numArray[5:3]
	//fmt.Println(numArray)  // .\main.go:156:16: invalid slice index: 5 > 3

}

func main() {
	//arrayDemo1()
	//arrayDemo2()
	//arrayDemo3()
	//arrayDemo4()
	//arrayDemo5()
	//arrayDemo6()
	//arrayDemo7()
	//arrayDemo8()
	//arrayDemo9()
	//arrayDemo10()
	//arrayDemo11()
	arrayDemo12()

}
