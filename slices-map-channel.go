package main

//the reference types in Go are slice, map and channel.
//See more http://blog.golang.org/go-slices-usage-and-internals
import (
	"fmt"
)

func main() {
	//slices
	s := make([]string, 3)
	fmt.Println("emp: ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("let:", len(s))
	fmt.Println("cap:", len(s))
	//slice operation
	//slice is reference type
	editSlices(s)
	fmt.Println("append slices", s)

	//map
	m := make(map[string]int)
	m["Thailand"] = 66
	m["UK"] = 44
	fmt.Println("Map: ", m)
	v1 := m["UK"]
	fmt.Println("v1", v1)

	delete(m, "UK")
	fmt.Println("Map: ", m)

	countryCode := map[string]int{"Thailand": 66, "UK": 44}
	fmt.Println("countryCode: ", countryCode)
	addMoreCountryCode(countryCode)
	//map is reference type
	fmt.Println("countryCode: ", countryCode)

}

func addMoreCountryCode(m map[string]int) {
	m["Norway"] = 55
}

func editSlices(s []string) {
	s[0] = "abc"
	s[1] = "edf"
}
