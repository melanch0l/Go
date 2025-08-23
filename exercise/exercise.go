package exercise

import "fmt"

type Player struct {
	Name      string
	Inventory []Item
}
type Item struct {
	Name string
	Type string
}

func Exercise() {

	p := Player{
		Name: "Min",
		Inventory: []Item{
			{Name: "red pill", Type: "Potion"},
		},
	}

	p.PickUpItem(Item{Name: "sword", Type: "Weapon"})
	p.UseItem("red pill")
	p.DropItem("sword")
	p.UseItem("shield")
}

func (p *Player) PickUpItem(newitem Item) {
	p.Inventory = append(p.Inventory, newitem)
	fmt.Printf("%s is pickup %s and store\n", p.Name, newitem.Name)

}
func (p *Player) DropItem(itemname string) {
	for i, item := range p.Inventory {
		if item.Name == itemname {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s drop  %s from inventory\n", p.Name, itemname)

			return
		}
	}
	fmt.Printf("%s does not have %s in inventory\n", p.Name, itemname)
}
func (p *Player) UseItem(itemname string) {
	for i, item := range p.Inventory {
		// i = index
		// item = value at that index
		if item.Name == itemname {
			if item.Type == "Potion" {
				fmt.Printf("%s is used %s feels fine now\n", p.Name, itemname)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
				// p.Inventory[:i] → all elements before index i
				// p.Inventory[i+1:] → all elements after index i
			} else {
				fmt.Printf("%s is used %s feels strong now\n", p.Name, itemname)
			}
			return
		}
	}
	fmt.Printf("%s dose not have %s in inventory\n", p.Name, itemname)

}
