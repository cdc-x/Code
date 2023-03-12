package main

import "fmt"

func main() {
	var a = '华'
	var b = 'a'
	fmt.Printf("a: %v,%c\n", a, a) // a: 21326,华
	fmt.Printf("b: %v,%c\n", b, b) // b: 97,a

	s := "hello南京"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}

	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) // pig

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2)) // 红萝卜

}
