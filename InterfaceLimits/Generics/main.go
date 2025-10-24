package main

import "fmt"

func main() {
	var a, b int = 1, 2
	var c, d float64 = 1.0, 2.0
	var e, f string = "hello ", "user"

	add(a, b)
	add(c, d)
	add(e, f)
}
func add[T int | float64 | string](a, b T) {
	fmt.Println("Result are :", a+b)
}
