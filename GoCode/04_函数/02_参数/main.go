package main

import "fmt"

func say(name string) {

	fmt.Println("hello, ", name)
}

func say2(age int, name string) {
	fmt.Printf("hello, my name is %v and I'm %v years old\n", name, age)
}

func say3(name, address string, age int) {
	fmt.Printf("hello, my name is %v and I'm %v years old\n", name, age)
	fmt.Println(address)
}

func say4(age int, args ...string) {
	fmt.Println(args)

	for index, value := range args {
		fmt.Println(index, value)
	}
}

func do(array [4]int) {
	array[0] = 100
	fmt.Println("do内部的array：", array)  // do内部的array： [100 2 3 4]
}


func do2(slice []int) {
	slice[0] = 100
	fmt.Println("do内部的slice：", slice)  // do内部的slice： [100 2 3 4]
}


func main() {
	//say("cdc")
	//say2(18, "cdc") // hello, my name is cdc and I'm 18 years old
	//say3("cdc", "南京", 18)
	//say4(18, "cdc", "南京", "中国")

	numArray := [4]int{1, 2, 3, 4}
	do(numArray)
	fmt.Println("main中的array：", numArray)  // main中的array： [1 2 3 4]

	numSlice := []int{1, 2, 3, 4}
	do2(numSlice)
	fmt.Println("main中的slice：", numSlice)  // main中的slice： [100 2 3 4]
}
