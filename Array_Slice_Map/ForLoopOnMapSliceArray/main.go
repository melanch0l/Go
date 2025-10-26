package main

import (
	"fmt"
)

func main() {
	userName := make([]string, 1, 4)
	userName[0] = "min"
	userName = append(userName, "nyi")
	userName = append(userName, "may")
	userName = append(userName, "zune")
	fmt.Println(userName)
	for _, value := range userName {
		fmt.Println(value)
	}
	birthPlace := make(map[string]string, 2)
	birthPlace["min"] = "Lonekin"
	birthPlace["may"] = "Yangon"
	for key, value := range birthPlace {
		fmt.Printf("Key:%s and Value:%s\n", key, value)
	}

}
