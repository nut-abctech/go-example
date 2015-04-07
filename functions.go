package main

import (
	"fmt"
)

func main() {
	fmt.Println("plus func ", calculate(1, 21))
	val1, val2 := vals()
	fmt.Println("multiple return vals ", val1, val2)
	_, sec := vals()
	fmt.Println("multiple return vals discard first val ", sec)

}

func calculate(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 1, 2
}
