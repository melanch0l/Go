package main

import "fmt"

// type alias on map
type value map[string]float32

func (p value) itemsPrice() {
	fmt.Println(p)
}
func main() {
	price := make([]float32, 2, 4)
	price[0] = 1.3
	price = append(price, 1.4)
	price = append(price, 1.5)
	fmt.Println(price)
	proudct := make(map[string]string, 1)
	proudct["name"] = "Mango"
	fmt.Println(proudct["name"])
	itemprice := make(value, 4)
	itemprice["mango"] = 1.2
	itemprice["orange"] = 1.2
	itemprice["banana"] = 0.5
	itemprice["apple"] = 2
	itemprice.itemsPrice()

}
