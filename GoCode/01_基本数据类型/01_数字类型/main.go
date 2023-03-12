package main

import (
	"fmt"
	"unsafe"
	"math"
)

func intDemo1() {

	num1 := 123456
	fmt.Printf("%T\n", num1)

	var num int8
	//num = 111111111 // .\main.go:7:6: constant 111111111 overflows int8
	fmt.Println(num)
}

func intDemo2() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)     // int8 1B -128~127
	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16) // int16 2B -32768~32767
	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32) // int32 4B -2147483648~2147483647
	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64) // int64 8B -9223372036854775808~9223372036854775807

	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)            // uint8 1B 0~255
	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)         // uint16 2B 0~65535
	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)         // uint32 4B 0~4294967295
	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64)) // uint64 8B 0~18446744073709551615

	var ui uint
	ui = uint(math.MaxUint64)                                  //再+1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui) // uint 8B 0~18446744073709551615

	var iMax, iMin int
	iMax = int(math.MaxInt64)                                           //再+1会导致overflows错误
	iMin = int(math.MinInt64)                                           //再-1会导致overflows错误
	fmt.Printf("%T %dB %v~%v\n", iMax, unsafe.Sizeof(iMax), iMin, iMax) // int 8B -9223372036854775808~9223372036854775807

}

func intDemo3() {
	//num1 := 123_456
	//fmt.Println(num1) // 123456

	// 定义一个二进制数
	//num2 := 0b00101101
	//fmt.Println(num2) // 45 --> 0b00101101用十进制表示是45

	// 定义一个八进制数
	num3 := 0377
	fmt.Println(num3) // 255 --> 0377用十进制表示是255

	// 定义一个十六进制数
	num4 := 0xff
	fmt.Println(num4) // 255 --> 0xff用十进制表示是255
}

func intDemo4() {

	a := 10

	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制
	fmt.Printf("%o \n", a) // 12    占位符%o表示八进制
	fmt.Printf("%d \n", a) // 10    占位符%o表示八进制
	fmt.Printf("%x \n", a) // a     占位符%x表示十六进制小写
	fmt.Printf("%X \n", a) // A     占位符%X表示十六进制大写

}

func floatDemo() {
	var f32 float32
	var f64 float64

	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32) // float32 4B -3.4028234663852886e+38~3.4028234663852886e+38
	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64) // float64 8B -1.7976931348623157e+308~1.7976931348623157e+308
}

func floatDemo2() {
	fmt.Printf("%f\n", math.Pi)   // 3.141593
	fmt.Printf("%.2f\n", math.Pi) // 3.14
}

func complexDemo() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1) // (1+2i)
	fmt.Println(c2) // (2+3i)
}

func main() {
	//intDemo1()
	//intDemo2()
	//floatDemo()
	//intDemo3()
	//intDemo4()
	//floatDemo2()
	//complexDemo()

}
