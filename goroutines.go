package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printStm(from string) {
	for i := 0; i < 3; i++ {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		fmt.Println(from, ":", i)
	}
}

func main() {
	printStm("Morning")

	go printStm("go routines")

	//goo exec an anonymous func
	go func(msg string) {
		fmt.Println(msg)
	}("going go")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
