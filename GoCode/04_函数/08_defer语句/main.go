package main

import "fmt"

func deferFun() {

	fmt.Println("statrt ...")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	defer fmt.Println("step4")
	fmt.Println("end ...")
}

func main() {

	//fmt.Println("start connect ...")
	//defer fmt.Println("close connect")
	//fmt.Println("do something ....")

	deferFun()
}


