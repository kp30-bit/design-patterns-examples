package decorator

type Beverage interface {
	Cost() int
	Description() string
}

type Coffee struct{}

func (b Coffee) Cost() int {
	return 5
}

func (b Coffee) Description() string {
	return "this is your coffee"
}

type BeverageWithMilk struct {
	Beverage Beverage
}

func (b BeverageWithMilk) Cost() int {
	return b.Beverage.Cost() + 5
}

func (b BeverageWithMilk) Description() string {
	return b.Beverage.Description() + " with milk\n\n"
}
