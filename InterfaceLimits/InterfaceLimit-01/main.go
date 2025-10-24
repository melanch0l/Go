package main

import (
	"fmt"
)
func main(){
 var a float64 = 5
 operator(a)
}
func operator(a interface{}){
	switch a.(type){
	case int:
		fmt.Println("it is interger")
	case float64:
		fmt.Println("it is float 64")
	default:
		fmt.Println("Maybe")
	}
}