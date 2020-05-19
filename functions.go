package main

import "fmt"

func main() {
	listNums := []float64{1,2,3,4,5}
	fmt.Println("Sum: ", addThemUp(listNums))

	val1, val2 := next2Values(4)
	fmt.Println("Next2Vals: ", val1, val2)

	// nested function gets Closure
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}

// no access to other functions' variables
func addThemUp(numbers []float64) float64{
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// return multiple values
func next2Values(num int) (int, int){
	return num+1, num+2
}

// unknown # arguments
func subtractThem(args ...int) int {
	finalValue := 0
	for _, val := range args {
		finalValue -= val
	}
	return finalValue
}

// Recursion

