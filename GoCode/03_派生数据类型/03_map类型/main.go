package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func mapDefine() {

	// var 变量名 map[key的类型]value的类型

	var m map[string]string
	fmt.Println(m)
}

func mapInit1() {
	m1 := map[string]string{
		"name":    "cdc",
		"city":    "南京",
		"address": "江宁区",
	}

	fmt.Printf("%#v\n", m1) // map[string]string{"address":"江宁区", "city":"南京", "name":"cdc"}

	m2 := make(map[string]int, 8)
	m2["cdc"] = 100
	m2["ctt"] = 100
	m2["cee"] = 100
	m2["ccc"] = 100
	fmt.Printf("%#v\n", m2) // map[string]int{"ccc":100, "cdc":100, "cee":100, "ctt":100}
}

func mapHandle1() {
	userInfo := map[string]string{
		"name":    "cdc",
		"city":    "南京",
		"address": "江宁区",
	}

	// 通过键名获取值
	city := userInfo["city"]
	fmt.Println(city) // 南京

	// 通过键名修改值
	userInfo["address"] = "雨花台"
	fmt.Println(userInfo) // map[address:雨花台 city:南京 name:cdc]

}

func mapHandle2() {
	userInfo := map[string]string{
		"name":    "cdc",
		"city":    "南京",
		"address": "江宁区",
	}

	// 如果键存在，ok为true，v为对应的值；不存在，ok为false，v为值类型的零值

	v, ok := userInfo["name"]
	fmt.Printf("v的值：%v  ok的值：%v\n", v, ok) // v的值：cdc  ok的值：true

	v1, ok1 := userInfo["age"]
	fmt.Printf("v1的值：%v  ok1的值：%v\n", v1, ok1) // v1的值：  ok1的值：false
}

func mapHandle3() {
	userInfo := map[string]string{
		"name":    "cdc",
		"city":    "南京",
		"address": "江宁区",
	}

	for key := range userInfo {
		fmt.Println(key)
	}

	for key, value := range userInfo {
		fmt.Printf("key: %v   value: %v\n", key, value)
	}

	for _, value := range userInfo {
		fmt.Println(value)
	}
}

func mapHandle4() {
	userInfo := map[string]string{
		"name":    "cdc",
		"city":    "南京",
		"address": "江宁区",
	}

	delete(userInfo, "name")
	delete(userInfo, "age") // 如果删除的键不存在也不会引发错误

	fmt.Printf("%#v\n", userInfo) // map[string]string{"address":"江宁区", "city":"南京"}
}

func mapHandle5() {
	// 这里只是对外层的切片进行了初始化
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 需要对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "cdc"
	mapSlice[0]["city"] = "南京"
	mapSlice[0]["address"] = "江宁区"
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["name"] = "tr"
	mapSlice[1]["city"] = "南京"
	mapSlice[1]["address"] = "雨花台"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func mapHandle6() {

	// 这里只对外层的map进行了初始化
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap) // map[]

	key := "name"
	value, ok := sliceMap[key]
	if !ok {
		// 对键对应的值类型进行初始化，即对切片进行初始化
		value = make([]string, 0, 2)
	}
	value = append(value, "cdc", "tr")
	sliceMap[key] = value
	fmt.Printf("%#v\n", sliceMap) // map[string][]string{"name":[]string{"cdc", "tr"}}
}

// 按照学生序号顺序输出对应分值
func mapExe1() {

	rand.Seed(time.Now().UnixNano())

	scoreMap := make(map[string]int, 100)
	for i := 1; i < 100; i++ {
		key := fmt.Sprintf("stu%03d", i) //生成stu开头的字符串
		value := rand.Intn(100)          // //生成0~99的随机整数
		scoreMap[key] = value
	}

	fmt.Println(scoreMap)

	// 将学号取出并存放到切片中
	stuSlice := make([]string, 0, 100)
	for key := range scoreMap {
		stuSlice = append(stuSlice, key)
	}

	fmt.Println(stuSlice)

	// 对切片进行排序
	sort.Strings(stuSlice)
	fmt.Println(stuSlice)

	// 按照排序后的顺序进行输出
	for _, stu := range stuSlice {
		fmt.Printf("学号：%v   分数：%v\n", stu, scoreMap[stu])
	}

}

// 写一个程序，统计一个字符串中每个字母出现的次数
func mapExe2() {

	str := "how do you do"

	strMap := make(map[string]int)
	for _, v := range str {
		key := string(v)

		if key != " " {
			_, ok := strMap[key]
			if !ok {
				strMap[key] = 0
			}
			strMap[key] += 1
		}
	}

	fmt.Printf("%#v\n", strMap)
}

func main() {
	//mapInit1()
	//mapHandle1()
	//mapHandle2()
	//mapHandle3()
	//mapHandle4()
	//mapHandle5()
	//mapHandle6()
	mapExe1()
	//mapExe2()
}
