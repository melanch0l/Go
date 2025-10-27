package main

import "fmt"

func main() {
	data := []int{1, 10, 20, 30}
	fmt.Println(variadicFunc(10, 20, 30, 40))
	fmt.Println(variadicFunc(data...)) //splitting slice to variable

}

// turning multiple argument with same type to slice with ...(variadic)
func variadicFunc(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}
