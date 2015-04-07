package main

import (
	"fmt"
	"time"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's a weekday")
	}

	//each into collection
	countries := []string{"Thailand", "UK", "US", "Japan", "Lao"}
	for _, c := range countries {
		fmt.Println("Country:", c)
	}
}
