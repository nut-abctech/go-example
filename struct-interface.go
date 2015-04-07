package main

import (
	"fmt"
	"math"
	"strconv"
)

type Geometry interface {
	area() float64
}

// Square
type Square struct {
	width, height float64
}

// pointer receivers
func (s *Square) area() float64 {
	return s.width * s.height
}

// Circle
type Circle struct {
	radius float64
}

// value receivers
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Person struct {
	name   string
	age    int
	nick   string
	gender string
}

//method
func (p *Person) Stringify() string {
	return p.name + " :[" + p.nick + "], " + strconv.Itoa(p.age) + " year olds, Gender: " + p.gender
}

func main() {
	person := Person{
		name:   "Pradit",
		age:    10,
		nick:   "nut",
		gender: "Male",
	}
	fmt.Println(person.Stringify())
	pointerPerson := &person
	// use . for dereference automatically
	pointerPerson.age = 30
	fmt.Println(person.Stringify())

	//interface Geometry
	s := Square{width: 3, height: 10}
	c := Circle{radius: 5}
	//need to be pointer because the method defined as a pointer receiver,
	Measure(&s) //**
	Measure(c)  //**
}

func Measure(g Geometry) {
	fmt.Println("area: ", g.area())
}
