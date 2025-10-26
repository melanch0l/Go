package main

import "fmt"

func main() {
	var data = [8]string{"hi", "hey", "hello", "bonjour", "halo", "bien", "ca va", "bye"}
	fmt.Println(data)
	var newData = data[1:3]
	newData[1] = "hi"
	fmt.Println(newData)
	fmt.Println(data)
	var extendData = newData[1:6]
	fmt.Println(extendData)
	fmt.Println(len(data), cap(extendData)) //display full data to the right

	//dynamic slice

	var dynamicData = []string{"1", "2"}
	fmt.Println(dynamicData)
	dynamicData = append(dynamicData, "3", "4", "5") //overwrite dynamicData and created new array
	fmt.Println(dynamicData)
}
