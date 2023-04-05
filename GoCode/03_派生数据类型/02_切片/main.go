package main

import "fmt"

func sliceDemo1() {
	var slice1 []int
	fmt.Printf("slice1的类型：%T  长度：%v  容量：%v  值：%v\n", slice1, len(slice1), cap(slice1), slice1) // slice1的类型：[]int  长度：0  容量：0

	if slice1 == nil {
		fmt.Println("切片未分配内存，是空的")
	} else {
		fmt.Println("切片已分配内存")
	}

}

func sliceInit1() {

	var numSlice1 = []int{1, 2, 3, 4}
	fmt.Printf("numSlice1的类型：%T  长度：%v  容量：%v  值：%v\n", numSlice1, len(numSlice1), cap(numSlice1), numSlice1) // numSlice1的类型：[]int  长度：4  容量：4  值：[1 2 3 4]

	var numArray = [10]int{1, 2, 3, 4, 5, 100, 200, 300, 400, 500}
	var numSlice2 = numArray[:]   // 获取全部
	var numSlice3 = numArray[3:5] // 获取部分

	numSlice3[0] = 10000
	fmt.Printf("numSlice3: %v, numArray: %v\n", numSlice3, numArray)

	fmt.Printf("numSlice2的类型：%T  长度：%v  容量：%v  值：%v\n", numSlice2, len(numSlice2), cap(numSlice2), numSlice2) // numSlice2的类型：[]int  长度：10  容量：10  值：[1 2 3 4 5 100 200 300 400 500]
	fmt.Printf("numSlice3的类型：%T  长度：%v  容量：%v  值：%v\n", numSlice3, len(numSlice3), cap(numSlice3), numSlice3) // numSlice3的类型：[]int  长度：2  容量：7  值：[4 5]

	numSlice4 := make([]int, 5, 10)
	fmt.Printf("numSlice4的类型：%T  长度：%v  容量：%v  值：%v\n", numSlice4, len(numSlice4), cap(numSlice4), numSlice4) // numSlice4的类型：[]string  长度：5  容量：10  值：[0 0 0 0 0]

}

func sliceDemo2() {

	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := array1[2:6]

	fmt.Printf("slice1的类型：%T\n", slice1)      // slice1的类型：[]int
	fmt.Printf("slice1的值：%v\n", slice1)       // slice1的值：[3 4 5 6]
	fmt.Printf("slice1的长度：%d\n", len(slice1)) // slice1的长度：4
	fmt.Printf("slice1的容量：%d\n", cap(slice1)) // slice1的容量：8

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]                                                 // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)) // s:[2 3] len(s):2 cap(s):4

}

func sliceDemo3() {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := array1[2:6]
	fmt.Printf("slice1的值：%v  长度：%v  容量：%v\n", slice1, len(slice1), cap(slice1)) // slice1的值：[3 4 5 6]  长度：4  容量：8

	slice2 := slice1[2:6]
	fmt.Printf("slice2的值：%v  长度：%v  容量：%v\n", slice2, len(slice2), cap(slice2)) // slice2的值：[5 6 7 8]  长度：4  容量：6
}

func sliceDemo4() {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := array1[2:6:8]
	fmt.Printf("slice1的值：%v  长度：%v  容量：%v\n", slice1, len(slice1), cap(slice1)) // slice1的值：[3 4 5 6]  长度：4  容量：6

	slice2 := array1[2:6:10]
	fmt.Printf("slice2的值：%v  长度：%v  容量：%v\n", slice2, len(slice2), cap(slice2)) // slice2的值：[3 4 5 6]  长度：4  容量：8
}

func sliceHandle1() {
	s1 := []string{"AAA", "BBB", "CCC", "DDD", "EEE", "FFF"}
	fmt.Println(s1[0]) // AAA
	fmt.Println(s1[2]) // CCC

	fmt.Println(s1[2:5]) // [CCC DDD EEE]

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	for index := range s1 {
		fmt.Printf("索引值：%v\n", index)
	}

	for index, value := range s1 {
		fmt.Printf("索引值：%v  元素值：%v\n", index, value)
	}

	for _, value := range s1 {
		fmt.Printf("元素值：%v\n", value)
	}

}

func sliceHandle2() {
	// 声明并初始化一个切片
	s1 := []string{"AAA", "BBB"}
	fmt.Printf("s1的地址：%p  s1的值：%v\n", s1, s1) // s1的地址：0xc0000443c0  s1的值：[AAA BBB]

	s2 := append(s1, "CCC")
	fmt.Printf("s2的地址：%p  s2的值：%v\n", s2, s2) // s2的地址：0xc000020080  s2的值：[AAA BBB CCC]

	s3 := append(s1, "CCC", "DDD", "EEE")
	fmt.Printf("s3的地址：%p  s3的值：%v\n", s3, s3) // s3的地址：0xc000050050  s3的值：[AAA BBB CCC DDD EEE]

	strSlice := []string{"CCC", "DDD", "EEE"}
	s4 := append(s1, strSlice...)
	fmt.Printf("s4的地址：%p  s4的值：%v\n", s4, s4) // s4的地址：0xc0000500a0  s4的值：[AAA BBB CCC DDD EEE]

	var slice1 []int
	//slice1[0] = 100  // panic: runtime error: index out of range [0] with length 0

	slice2 := append(slice1, 1, 2, 3, 4, 5)
	fmt.Println(slice2) // [1 2 3 4 5]

}

func sliceHandle3() {
	var s1 []int

	for i := 1; i <= 10; i++ {
		s1 = append(s1, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", s1, len(s1), cap(s1), s1)
	}

}

func sliceHandle4() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1的值：%v\n", s1)  // s1的值：[1 2 3 4 5]

	//s2 := s1
	//s2[0] = 100
	//fmt.Printf("s2的值：%v\n", s2)  // s2的值：[100 2 3 4 5]
	//fmt.Printf("s1的值：%v\n", s1)  // s1的值：[100 2 3 4 5]

	s3 := make([]int, 5, 5)
	copy(s3, s1)
	s3[0] = 100
	fmt.Printf("s3的值：%v\n", s3)  // s3的值：[100 2 3 4 5]
	fmt.Printf("s1的值：%v\n", s1)  // s1的值：[1 2 3 4 5]


}

func sliceHandle5(){
	s1 := []int{1, 2, 3, 4, 5}
	// 要删除索引为2的元素
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)  // [1 2 4 5]
}

func main() {
	//sliceDemo1()
	//sliceInit1()
	//sliceDemo2()
	//sliceDemo3()
	//sliceDemo4()
	//sliceHandle1()
	//sliceHandle2()
	//sliceHandle3()
	//sliceHandle4()
	sliceHandle5()
}
