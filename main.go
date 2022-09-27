package main

import "fmt"

type weapon interface {
	getDescription() string
	getCost() float64
}

type SilverSword struct {
	discription string
	cost        float64
}

func NewSilverSword() *SilverSword {
	return &SilverSword{discription: "Silver sword :  is a specialized sword within the Witcher lore that is primarily used to combat supernatural enemies.", cost: 300.5}
}

func NewSilverSword2(discription string, cost float64) *SilverSword {
	return &SilverSword{discription: discription, cost: cost}
}

func (s *SilverSword) getDescription() string {
	return s.discription
}

func (s *SilverSword) getCost() float64 {
	return s.cost
}

type SteelSword struct {
	discription string
	cost        float64
}

func NewSteelSword() *SteelSword {
	return &SteelSword{discription: "Steel Sword :  steel is always meant for fighting humans/nonhumans and normal, non-magical creatures like bears, unlike the books where there is more discrepancy.", cost: 300.5}
}

func NewSteelSword1(discription string, cost float64) *SteelSword {
	return &SteelSword{discription: discription, cost: cost}
}

func (s *SteelSword) getDescription() string {
	return s.discription
}

func (s *SteelSword) getCost() float64 {
	return s.cost
}

type BeastOil struct {
	weapon      weapon
	discription string
	cost        float64
}

func NewBeastOil(weapon2 weapon) *BeastOil {
	return &BeastOil{weapon: weapon2, discription: weapon2.getDescription() + "\nBeast oil: 10% Attack power", cost: weapon2.getCost() + weapon2.getCost()*0.1}
}

func (b *BeastOil) getDescription() string {
	return b.discription
}

func (b *BeastOil) getCost() float64 {
	return b.cost
}

type DraconidOil struct {
	weapon      weapon
	discription string
	cost        float64
}

func NewDraconidOil(weapon2 weapon) *DraconidOil {
	return &DraconidOil{weapon: weapon2, discription: weapon2.getDescription() + "\nDraconid Oil: 10% Attack power", cost: weapon2.getCost() + weapon2.getCost()*0.1}
}

func (b *DraconidOil) getDescription() string {
	return b.discription
}

func (b *DraconidOil) getCost() float64 {
	return b.cost
}
func main() {
	var silverSword weapon
	silverSword = NewSilverSword()
	silverSword = NewBeastOil(silverSword)
	silverSword = NewDraconidOil(silverSword)
	fmt.Println(silverSword.getDescription())
	fmt.Println(silverSword.getCost())

	fmt.Println()

	var steelSword weapon
	steelSword = NewSteelSword()
	steelSword = NewDraconidOil(steelSword)
	steelSword = NewBeastOil(steelSword)
	fmt.Println(steelSword.getDescription())
	fmt.Println(steelSword.getCost())

}
