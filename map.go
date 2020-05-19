package main

import (
	"fmt"
)

func main()  {
	// map is a collection of key value pairs
	presAge := make(map[string] int)
	presAge["soobin"] = 31
	fmt.Println(presAge["soobin"])

	delete(presAge, "soobin")
	fmt.Println(presAge["soobin"])


}

