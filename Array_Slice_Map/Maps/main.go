package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "google.com",
		"Facebook": "facebook.com",
	}
	fmt.Println(websites)
	websites["YouTube"] = "youtube.com"
	fmt.Println(websites)
	delete(websites, "Facebook")
	fmt.Println(websites)
	fmt.Println(websites["Google"])

}
