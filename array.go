package main

import (
	"fmt"
	"time"
)

func main() {
	// Array - fixed num
	var favNums2[5] float64 // unassigned numbers are all 0
	favNums2[0] = 16
	favNums2[1] = 4
	fmt.Println(favNums2)

	favNums3 := [5]float64 {1,2,3,4,5}
	fmt.Println(favNums3)

	favStrings := [2]string {} // unassinged are all empty strings
	favStrings[0] ="hi"
	fmt.Println(favStrings)

	for i,  value := range favNums3 {
		fmt.Println(i, value)
	}

	nums := []int {5,4,3,2,1} // array length is pre-set
	numSlice := nums[3:5]
	numSlice[1] = 2
	fmt.Println(numSlice)

	// make([]int, len, capacity)
	nums_default := make([]int, 5, 10) 
		// allocate zeroed array then return a slice that refers to that array
		// third argument is the capacity
	copy(nums_default, numSlice)
	fmt.Println(nums_default)
	fmt.Println(len(nums_default), cap(nums_default))

	fmt.Println(time.Now())
}