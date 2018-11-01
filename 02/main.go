package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
测试bool
*/
func testBool() {
	var isLast = false
	fmt.Println(isLast)
}

/*
测试数值
*/
func testNumber() {
	var a int32 = 2
	var b float32 = 4.67890
	var c complex64 = 5 + 5i
	fmt.Println(a, b, c)
}

/*
测试字符串
*/
func testSring() {
	var a = `hello`
	var b = ` world`
	fmt.Println(a + b)
}

/*
测试错误
*/
func testError() {
	err := errors.New("fatal error")
	fmt.Println(err)
}

/*
测试array
*/
func testArray() {
	arr := [3]int{1, 2, 3}
	doubleArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(arr[0])
	fmt.Println(doubleArray[1][3])
}

/*
测试slice
*/
func testSlice() {
	aSlice := []int{1, 2, 3}
	fmt.Println(aSlice)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bSlice := arr[2:7]
	fmt.Println(bSlice)
	fmt.Println(len(bSlice))
	fmt.Println(cap(bSlice))

	appSlice := append(bSlice, 12)
	fmt.Println(appSlice)
	fmt.Println(arr)

	cSlice := make([]int, 10)
	copy(cSlice, appSlice)
	fmt.Println(cSlice)
}

/*
测试map
*/
func testMap() {
	mm := make(map[string]string)
	mm["name"] = "chuck"
	mm["age"] = "12"
	fmt.Println(mm)

	delete(mm, "age")
	fmt.Println(mm)

	age, ok := mm["age"]
	fmt.Println(ok)
	fmt.Println(age)
}

func main() {
	var name, age = "chuck", 12
	fmt.Println(name)
	fmt.Println(strconv.Itoa(age))

	testBool()
	testNumber()
	testSring()
	testError()
	testArray()
	testSlice()
	testMap()
}
