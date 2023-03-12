package main

import "fmt"

func main() {
	num := 12
	numFloat := float32(num)
	fmt.Printf("%T -- %v\n", numFloat, numFloat) // float32 -- 12

	a := 'a'
	aStr := string(a)
	fmt.Printf("%T -- %v\n", aStr, aStr) // string -- a
}
