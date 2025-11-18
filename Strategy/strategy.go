package strategy

import "fmt"

type TravelStrategy interface {
	Travel(source, destination string)
}

type Car struct{}
type Bike struct{}
type Walk struct{}

func (t *Car) Travel(source, destination string) {
	fmt.Printf("\nTravelling via car from %v to %v\n", source, destination)
}
func (t *Bike) Travel(source, destination string) {
	fmt.Printf("\nTravelling via bike from %v to %v\n", source, destination)
}
func (t *Walk) Travel(source, destination string) {
	fmt.Printf("\nTravelling via walk from %v to %v\n", source, destination)
}

type TravelWrapper struct {
	strategy TravelStrategy
}

func (t *TravelWrapper) SetStrategy(strategy TravelStrategy) {
	t.strategy = strategy
}

func (t *TravelWrapper) Travel(source, destination string) {
	if t.strategy == nil {
		fmt.Printf("No strategy selected")
		return
	}
	t.strategy.Travel(source, destination)
}
