package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{name: "min", age: 26} //pass this value to func
	x := 20
	ptr := &x
	fmt.Printf("value of x %d and address of x %v and  value in the address of pointer pointing at is %d\n", x, ptr, *ptr)
	*ptr = 30 //reset the value where pointer pointed at//bypass pass by value nature
	fmt.Printf("value of x %d and address of x %v and  value in the address of pointer pointing at is %d\n", x, ptr, *ptr)
	fmt.Println("Name before", person.name)
	changeName(&person)
	// The & operator gives you the memory address of a variable.
	// In other words, it produces a pointer to that variable.
	fmt.Printf("Returned info My name is %v and age is %v\n", person.name, person.age)
	person.inputName("May Zune", 23)
	fmt.Printf("Returned info My GF name is %v and age is %v\n", person.name, person.age)

}
func changeName(p *Person) { //one parameter with a pointer type
	// The * operator is used to follow a pointer and access the value at that address.
	// Dereferencing lets you read or modify the value stored in the memory location.
	p.name = "min nyi"
	p.age = 26
	fmt.Printf("My name is %v and age is %v\n", p.name, p.age)
}

// (p *Person) is method receiver in this method format and sturct Person will be accessable to method name inputname
func (p *Person) inputName(name string, age int) {
	p.name = name
	p.age = age
}
