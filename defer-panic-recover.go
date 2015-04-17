package main

import (
	"fmt"
	"os"
)

func firstFn() {
	fmt.Println("1st exec")
}

func secFn() {
	fmt.Println("2st exec")
}

func main() {
	//defer -> schedules a function call and be run after the function completes
	defer secFn()
	firstFn()

	//panic => unexpectedly fails. e.g. fail fast on error
	// panic("a problem")
	// _, err := os.Create("./tmp/file")

	// if err != nil {
	// 	panic(err)
	// }
}
