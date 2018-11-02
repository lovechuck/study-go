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

/*
测试If
*/
func testIf() {
	num := 10
	if num < 10 {
		fmt.Println(`less than 10`)
	} else if num == 10 {
		fmt.Println(`equal 10`)
	} else {
		fmt.Println(`more than 10`)
	}
}

/*
测试switch
*/
func testSwitch() {
	i := 12
	switch {
	case i < 12:
		fmt.Println("< 12")
	case i == 12:
		fmt.Println("= 12")
	default:
		fmt.Println("> 12")
	}
}

/*测试for*/
func testFor() {
	sum := 0
	for num := 1; num < 10; num++ {
		sum += num
	}
	fmt.Println(sum)

	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println(i)
	}

	arrs := [5]int{1, 2, 3, 4, 5}
	for index, item := range arrs {
		fmt.Println(index, item)
	}

	slices := []string{"a", "b", "c"}
	for index, item := range slices {
		fmt.Println(index, item)
	}

	maps := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, val := range maps {
		fmt.Println(key, val)
	}
}

/*测试函数*/
func testFunc(a int, b int) (int, int) {
	return a + b, a - b
}
func testFunc1(a ...int) {
	fmt.Println(a)
}

/*测试init*/
func init() {
	fmt.Println("init")
}

/*测试全局变量*/
var g = 1

func testGlobal() {
	a := 2
	b := 3
	g := a + b
	fmt.Println(g)
}

/*测试局部变量*/
func testLocal() {
	a := 2
	b := 3
	g := a + b
	fmt.Println(g)
}

/*测试形势参数*/
func testParams(g int) {
	g = g + 1
	fmt.Println(g)
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
	testIf()
	testSwitch()
	testFor()

	a, b := testFunc(1, 2)
	fmt.Println(a, b)
	testFunc1(1, 2, 3, 4)

	testGlobal()
	testLocal()
	testParams(1)
}
