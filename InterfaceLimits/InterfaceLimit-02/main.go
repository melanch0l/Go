package main

import "fmt"
func main(){
fmt.Print("enter a value: ")
var a int
fmt.Scan(&a)
operator(a)
}
func operator(a interface{}){
	//extracting type info
	value, isInt := a.(int)
	if isInt{
		fmt.Println("The value is:",value)
		return
	}else{
		fmt.Println("Other type ")
		return
	}
}
