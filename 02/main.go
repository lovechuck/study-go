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

func main() {
	var name, age = "chuck", 12
	fmt.Println(name)
	fmt.Println(strconv.Itoa(age))

	testBool()
	testNumber()
	testSring()
	testError()
}
