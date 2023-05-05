package main

import "fmt"

func init(){
	fmt.Println("init1...")
}

func init()  {
	fmt.Println("init2...")
}

func init()  {
	fmt.Println("init3...")
}

func main() {
	fmt.Println("这里是main函数...")
}
