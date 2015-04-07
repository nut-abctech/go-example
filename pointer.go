package main

import (
	"fmt"
	"unsafe"
)

func setZero(val int) {
	val = 0
}

func setZeroPointer(p *int) {
	*p = 0
}

type MyType struct {
	id   int
	name string
}

func main() {
	val := 10
	fmt.Println("initital : ", val)

	setZero(val) //pass by value func
	fmt.Println("set zero : ", val)

	setZeroPointer(&val) ///& gives the mem address of val to pointer *p
	fmt.Println("set zero pointer : ", val)

	fmt.Println("pointer address : ", &val)

	//example 2
	myTypeValue := MyType{10, "bill"}
	pointer := unsafe.Pointer(&myTypeValue)

	// Display the address and values
	fmt.Printf("Addr: %v id : %d name: %s\n",
		pointer,
		myTypeValue.id,
		myTypeValue.name)

	//copy whole value to func
	changeType(myTypeValue)

	//pass value by address
	changeTypeRef(&myTypeValue)

	fmt.Printf("Addr: %v id : %d name: %s\n",
		pointer,
		myTypeValue.id,
		myTypeValue.name)

	myTypeNew := new(MyType) // allocate Mem and zeroed values (int = 0, string = "")
	fmt.Printf("Use New Addr: %v id : %d name: %s\n",
		pointer,
		myTypeNew.id,
		myTypeNew.name)

	changeTypeRef(myTypeNew)
	fmt.Printf("Use New Addr: %v id : %d name: %s\n",
		pointer,
		myTypeNew.id,
		myTypeNew.name)
}

// myTypeValue is the copy value
func changeType(myTypeValue MyType) {
	myTypeValue.id = 100
	myTypeValue.name = "Michale"

	pointer := unsafe.Pointer(&myTypeValue)

	fmt.Printf("Addr: %v id : %d name: %s\n",
		pointer,
		myTypeValue.id,
		myTypeValue.name)
}

//receive address, myTypeValue is pointer
func changeTypeRef(myTypeValue *MyType) {
	myTypeValue.id = 100
	myTypeValue.name = "Michale Ref"
}
