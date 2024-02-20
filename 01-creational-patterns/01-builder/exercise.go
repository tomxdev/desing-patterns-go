package main

import "fmt"

// Burger Product
type Burger struct {
	breadType string
	meatType  string
	toppings  string
}

// BurgerBuilder Builder Interface
type BurgerBuilder interface {
	setBreadType()
	setMeatType()
	setToppings()
	getBurger() Burger
}

// RegularBurgerBuilder Concrete Builder 1
type RegularBurgerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newRegularBurgerBuilder() *RegularBurgerBuilder {
	return &RegularBurgerBuilder{}
}

// Set the Bread Type
func (b *RegularBurgerBuilder) setBreadType() {
	b.breadType = "Sesame"
}

// Set the Meat Type
func (b *RegularBurgerBuilder) setMeatType() {
	b.meatType = "beef"
}

// Set the Toppings
func (b *RegularBurgerBuilder) setToppings() {
	b.toppings = "Lettuce, Tomato, Bacon and Cheese."
}

func (b *RegularBurgerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// VeganBurgerBuilder Concrete Builder 2
type VeganBurgerBuilder struct {
	breadType string
	meatType  string
	toppings  string
}

func newVeganBurgerBuilder() *VeganBurgerBuilder {
	return &VeganBurgerBuilder{}
}

// Set the Bread Type
func (b *VeganBurgerBuilder) setBreadType() {
	b.breadType = "Gluten Free"
}

// Set the Meat Type
func (b *VeganBurgerBuilder) setMeatType() {
	b.meatType = "Black Bean"
}

// Set the Toppings
func (b *VeganBurgerBuilder) setToppings() {
	b.toppings = "Lettuce and Tomato"
}

func (b *VeganBurgerBuilder) getBurger() Burger {
	return Burger{
		breadType: b.breadType,
		meatType:  b.meatType,
		toppings:  b.toppings,
	}
}

// Director struct
type Director struct {
	builder BurgerBuilder
}

func newDirector(b BurgerBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b BurgerBuilder) {
	d.builder = b
}

func getBuilder(builderType string) BurgerBuilder {
	if builderType == "regular" {
		return newRegularBurgerBuilder()
	}

	if builderType == "vegan" {
		return newVeganBurgerBuilder()
	}

	return nil
}

func (d *Director) buildBurger() Burger {
	d.builder.setBreadType()
	d.builder.setMeatType()
	d.builder.setToppings()
	return d.builder.getBurger()
}

func main() {
	regularBurgerBuilder := getBuilder("regular")
	veganBurgerBuilder := getBuilder("vegan")

	// Initialize the director
	director := newDirector(regularBurgerBuilder)

	// Create the regular burger
	regularBurger := director.buildBurger()

	printBurgerInfo(regularBurger, "Regular")

	// Build the vegan burger
	director.setBuilder(veganBurgerBuilder)
	veganBurger := director.buildBurger()

	printBurgerInfo(veganBurger, "Vegan")
}

func printBurgerInfo(burger Burger, burgerType string) {
	fmt.Printf("%s Burger Bread Type: %s\n", burgerType, burger.breadType)
	fmt.Printf("%s Burger Meat Type: %s\n", burgerType, burger.meatType)
	fmt.Printf("%s Burger Toppings: %s\n", burgerType, burger.toppings)
}
