package testing

//Sum ...
func Sum(numbers []int) int {
	sum := 0
	// 这个 bug 是故意的
	for _, n := range numbers {
		sum += n
	}
	return sum
}
