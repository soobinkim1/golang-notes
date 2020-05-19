package main // declares package name

import "fmt"// formattig, providing format for input & output


func main() {
	fmt.Println("hello, world")
	var (
		age int = 40
		fav float64 = 1.6180339
		myname string = "soobin"
		isOver40 bool = false
	)
	randNum := 1
	fmt.Println(myname, age, randNum, fav, isOver40)
	fmt.Printf("%.3f \n", fav) // formatted string
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	
	// For Loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i ++
	}
	for j := 0; j < 5; j ++ {
		fmt.Println(j)
	}
	// If
	yourAge := 20
	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can't Vote")
	} else {
		fmt.Println("Below 16")
	}
	// Switch
	switch yourAge{
	case 16: fmt.Println("You are 16")
	default: fmt.Println("You are not 16")
	}
}