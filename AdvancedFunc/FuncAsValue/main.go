package main

import "fmt"

func main() {
	num1, num2 := 20, 30
	//add,multiply are passing func as argument
	operate(num1, num2, add)
	operate(num1, num2, multiply)
	num3 := getOperation()
	operate(num1, num2, num3)

}

// operator is func as value
func operate(a, b int, operator func(int, int) int) {
	fmt.Println(operator(a, b))
}
func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}

// return func as value
func getOperation() func(int, int) int {
	return add
}
