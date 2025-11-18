package factory

type Dessert interface {
	Serve() string
}

type Barfi struct {
}

type Jamun struct {
}

func (d *Barfi) Serve() string {
	return "Barfi served"
}

func (d *Jamun) Serve() string {
	return "Jamun served"
}

func DessertFactory(kind string) Dessert {
	if kind == "Barfi" {
		return &Barfi{}
	} else if kind == "Jamun" {
		return &Jamun{}
	}
	return nil
}
