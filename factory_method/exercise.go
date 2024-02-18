package main

import "fmt"

// ICar Product interface
type ICar interface {
	setCarType(carType string)
	setYear(year int)
	printDetails()
}

// Car Abstract Product
type Car struct {
	carType string
	year    int
}

func (c *Car) setCarType(carType string) {
	c.carType = carType
}

func (c *Car) setYear(year int) {
	c.year = year
}

func (c *Car) printDetails() {
	fmt.Printf("Car Type: %s\n", c.carType)
	fmt.Printf("Car Year: %d\n", c.year)
}

// LuxuryCar Concrete Product 1
type LuxuryCar struct {
	Car
}

func newLuxuryCar() ICar {
	return &LuxuryCar{
		Car: Car{
			carType: "Luxury",
			year:    2023,
		},
	}
}

// HybridCar Concrete Product 2
type HybridCar struct {
	Car
}

func newHybridCar() ICar {
	return &HybridCar{
		Car: Car{
			carType: "Hybrid",
			year:    2021,
		},
	}
}

// factory method
func getCar(carType string) ICar {
	if carType == "luxury" {
		return newLuxuryCar()
	}

	if carType == "hybrid" {
		return newHybridCar()
	}

	return nil
}

func main() {
	luxuryCar := getCar("luxury")
	hybridCar := getCar("hybrid")

	luxuryCar.printDetails()
	hybridCar.printDetails()
}
