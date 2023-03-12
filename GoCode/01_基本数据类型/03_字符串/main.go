package main

import (
	"fmt"
	"strings"
	"bytes"
)

func strDemo1() {

	// 初始化当行字符串
	str1 := "hello \nworld!"

	// 初始化多行字符串
	str2 := `hello\n
			 world`

	fmt.Println(str1)
	fmt.Println(str2)
}

func strDemo2() {
	// 字符串长度
	str1 := "abcdefg"
	fmt.Println(len(str1)) // 7

	// 字符串拼接
	// 方式一，使用加号
	name := "cdc"
	age := "18"
	msg1 := name + " " + age
	fmt.Printf("msg1: %v\n", msg1) // cdc 18

	msg2 := ""
	msg2 += name
	msg2 += " "
	msg2 += age
	fmt.Printf("msg2: %v\n", msg2) // cdc 18

	// 方式二，fmt.Sprintf()
	msg3 := fmt.Sprintf("%s,%s", name, age)
	fmt.Printf("msg: %v\n", msg3)

	// 方式三，strings.Join()
	msg4 := strings.Join([]string{name, age}, " ")
	fmt.Printf("msg: %v\n", msg4)

	// 方式四，buffer.WriteString
	var buffer bytes.Buffer
	buffer.WriteString("cdc")
	buffer.WriteString(" ")
	buffer.WriteString("18")
	fmt.Printf("msg: %v\n", buffer.String())

	// 字符串切割
	str2 := "how do you do"
	splitRet := strings.Split(str2, "e")
	fmt.Printf("%v -- %T -- %d\n", splitRet, splitRet, len(splitRet))

	// 是否包含
	fmt.Println(strings.Contains(str2, "do"))
	fmt.Println(strings.Contains(str2, "hehe"))

	// 判断前后缀
	fmt.Println(strings.HasPrefix(str2, "how"))
	fmt.Println(strings.HasPrefix(str2, "do"))
	fmt.Println(strings.HasSuffix(str2, "how"))
	fmt.Println(strings.HasSuffix(str2, "do"))

	// 字串出现位置
	fmt.Println(strings.Index(str2, "do"))     // 4
	fmt.Println(strings.LastIndex(str2, "do")) // 11

	// 字符串切片
	fmt.Println(str2[:5])  // how d  --> 从开头切到指定索引
	fmt.Println(str2[3:6]) //  do --> 从指定开始索引位置切到指定结束索引位置
	fmt.Println(str2[2:])  // w do you do --> 从指定索引位置切到字符串最后
	fmt.Println(str2[:])   // how do you do --> 从开始切到最后
}

func strDemo3() {
	// 普通占位符
	type website struct {
		url string
	}

	web := website{url: "www.chendacheng.com"}

	name := "cdc"

	fmt.Printf("%v\n", name) // cdc
	fmt.Printf("%+v\n", web) // {url:www.chendacheng.com}
	fmt.Printf("%#v\n", web) // main.website{url:"www.chendacheng.com"}
	fmt.Printf("%T\n", name) // string

	// 布尔占位符
	fmt.Printf("%t\n", true) //输出值的 true 或 false

	// 整数占位符
	fmt.Printf("%b\n", 100)  //二进制表示  1100100
	fmt.Printf("%c\n", 1111) //数值对应的 Unicode 编码字符  ї
	fmt.Printf("%d\n", 10)   //十进制表示  10
	fmt.Printf("%o\n", 8)    //八进制表示  10
	fmt.Printf("%q\n", 22)   //转化为十六进制并附上单引号  '\x16'
	fmt.Printf("%x\n", 1223) //十六进制表示，用a-f表示  4c7
	fmt.Printf("%X\n", 1223) //十六进制表示，用A-F表示  4C7
	fmt.Printf("%U\n", 1024) //Unicode表示  U+0400

	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte  wqdew
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串  "dedede"
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示  6162637a7863
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示  6173647A7863

	// 浮点数和复数的组成部分
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示  1.234500e+01
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示  1.234455E+01
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分  12.345600
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出  12.3456
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出  12.3456

	// 指针占位符
	fmt.Printf("%T\n", &name) //*string 字符串指针
	fmt.Printf("%p\n", &name) //0xc000042270 指针地址

}

func main() {
	//strDemo1()
	//strDemo2()
	strDemo3()
}
