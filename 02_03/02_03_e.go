package main

import "fmt"

// Burger Product
type Burger struct {
	breadType string
	meatType  string
	toppings  string
}

// BurguerBuilder Builder Interface
type BurguerBuilder interface {
	setBreadType()
	setMeatType()
	setToppings()
	getBurger() Burger
}

// RegularBurguerBuilder Concrete Builder 1
type RegularBurguerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newRegularBurguerBuilder() *RegularBurguerBuilder {
	return &RegularBurguerBuilder{}
}

// Set the Bread Type
func (b *RegularBurguerBuilder) setBreadType() {
	b.breadType = "Sesame"
}

// Set the Meat Type
func (b *RegularBurguerBuilder) setMeatType() {
	b.meatType = "beef"
}

// Set the Toppings
func (b *RegularBurguerBuilder) setToppings() {
	b.toppings = "Lettuce, Tomato, Bacon and Cheese."
}

func (b *RegularBurguerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// VeganBurguerBuilder Concrete Builder 2
type VeganBurguerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newVeganBurguerBuilder() *VeganBurguerBuilder {
	return &VeganBurguerBuilder{}
}

// Set the Bread Type
func (b *VeganBurguerBuilder) setBreadType() {
	b.breadType = "Gluten Free"
}

// Set the Meat Type
func (b *VeganBurguerBuilder) setMeatType() {
	b.meatType = "Black Bean"
}

// Set the Toppings
func (b *VeganBurguerBuilder) setToppings() {
	b.toppings = "Lettuce and Tomato"
}

func (b *VeganBurguerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// Director struct
type Director struct {
	builder BurguerBuilder
}

func newDirector(b BurguerBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b BurguerBuilder) {
	d.builder = b
}

func getBuilder(builderType string) BurguerBuilder {
	if builderType == "regular" {
		return newRegularBurguerBuilder()
	}

	if builderType == "vegan" {
		return newVeganBurguerBuilder()
	}

	return nil
}

func (d *Director) buildBurguer() Burger {
	d.builder.setBreadType()
	d.builder.setMeatType()
	d.builder.setToppings()
	return d.builder.getBurger()
}

func main() {
	regularBurguerBuilder := getBuilder("regular")
	veganBurguerBuilder := getBuilder("vegan")

	// Initialize the director
	director := newDirector(regularBurguerBuilder)

	// Create the regular burger
	regularBurguer := director.buildBurguer()

	printBurguerInfo(regularBurguer, "Regular")

	// Build the vegan burger
	director.setBuilder(veganBurguerBuilder)
	veganBurguer := director.buildBurguer()

	printBurguerInfo(veganBurguer, "Vegan")
}

func printBurguerInfo(burger Burger, burgerType string) {
	fmt.Printf("%s Burger Bread Type: %s\n", burgerType, burger.breadType)
	fmt.Printf("%s Burger Meat Type: %s\n", burgerType, burger.meatType)
	fmt.Printf("%s Burger Toppings: %s\n", burgerType, burger.toppings)
}
