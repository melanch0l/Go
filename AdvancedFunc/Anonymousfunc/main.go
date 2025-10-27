package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	//anonymonus func
	transformed := transformNumbers(&numbers, func(numbers int) int {
		return numbers * 2
	})
	fmt.Println(transformed)
	byTwo := divide(2)
	byThree := divide(3) //return func with divisor 3
	quotient1 := transformNumbers(&numbers, byTwo)
	quotient2 := transformNumbers(&numbers, byThree)
	fmt.Println(quotient1)
	fmt.Println(quotient2)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
func divide(divisor int) func(int) int {
	return func(dividend int) int {
		return dividend / divisor //divisor closure
	}
}
