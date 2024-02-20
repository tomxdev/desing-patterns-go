package main

import "fmt"

// ICar Abstract Product 1
type ICar interface {
	setCarType(cartType string)
	printDetails()
}

type Car struct {
	carType string
}

func (c *Car) setCarType(cartType string) {
	c.carType = cartType
}

func (c *Car) printDetails() {
	fmt.Printf("Car Type: %s\n", c.carType)
}

// ICarDetails Abstract Product 2
type ICarDetails interface {
	setTransmission(transmission string)
	setEngine(engine string)
	setGasType(gasType string)
	printDetails()
}

type CarDetails struct {
	transmission string
	engine       string
	gasType      string
}

func (cd *CarDetails) setTransmission(transmission string) {
	cd.transmission = transmission
}

func (cd *CarDetails) setEngine(engine string) {
	cd.engine = engine
}

func (cd *CarDetails) setGasType(gasType string) {
	cd.gasType = gasType
}

func (cd *CarDetails) printDetails() {
	fmt.Printf("Car Transmission: %s\n", cd.transmission)
	fmt.Printf("Car Engine: %s\n", cd.engine)
	fmt.Printf("Car Gas Type: %s\n", cd.gasType)
	fmt.Println()
}

// LuxuryCar Concrete Product A1
type LuxuryCar struct {
	Car
}

// LuxuryCarDetails Concrete Product A2
type LuxuryCarDetails struct {
	CarDetails
}

// HybridCar Concrete Product B1
type HybridCar struct {
	Car
}

// HybridCarDetails Concrete Product B2
type HybridCarDetails struct {
	CarDetails
}

// ICarFactory Abstract Factory Interface
type ICarFactory interface {
	makeCar() ICar
	makeCarDetails() ICarDetails
}

// LuxuryCarFactory Concrete Factory 1
type LuxuryCarFactory struct {
}

func (l *LuxuryCarFactory) makeCar() ICar {
	return &LuxuryCar{
		Car: Car{
			carType: "Luxury",
		},
	}
}

func (l *LuxuryCarFactory) makeCarDetails() ICarDetails {
	return &LuxuryCarDetails{
		CarDetails: CarDetails{
			transmission: "manual",
			engine:       "gas",
			gasType:      "premium",
		},
	}
}

type HybridCarFactory struct {
}

func (l *HybridCarFactory) makeCar() ICar {
	return &HybridCar{
		Car: Car{
			carType: "Hybrid",
		},
	}
}

func (l *HybridCarFactory) makeCarDetails() ICarDetails {
	return &HybridCarDetails{
		CarDetails: CarDetails{
			transmission: "automatic",
			engine:       "hybrid",
			gasType:      "electric",
		},
	}
}

// factory method
func getCarFactory(carType string) ICarFactory {
	if carType == "luxury" {
		return &LuxuryCarFactory{}
	}

	if carType == "hybrid" {
		return &HybridCarFactory{}
	}

	return nil
}

func main() {
	luxuryFactory := getCarFactory("luxury")
	hybridFactory := getCarFactory("hybrid")

	// Build the Luxury Car
	luxuryCar := luxuryFactory.makeCar()
	luxuryCarDetails := luxuryFactory.makeCarDetails()

	luxuryCar.printDetails()
	luxuryCarDetails.printDetails()

	// Build the Hybrid Car
	hybridCar := hybridFactory.makeCar()
	hybridCarDetails := hybridFactory.makeCarDetails()

	hybridCar.printDetails()
	hybridCarDetails.printDetails()
}
