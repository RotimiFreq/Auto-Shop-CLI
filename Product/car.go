package Product

import "fmt"

type Car struct {
	productType     string
	brandName       string
	model           string
	horsepower      string
	color string
	engine_Type     string
	production_date string
	fuelType        string
	quantity        int
	price           float64

}

type Order_Car struct {
	ProductType     string
	BrandName       string
	Model           string
	Horsepower      string
	Engine_Type     string
	Production_Date string
	Color           string
	FuelType        string
	Quantity        int
	Price           float64
}

func CarOrderToCar(O *Order_Car) *Car {
	var car = &Car{
		brandName:       O.BrandName,
		productType:     O.ProductType,
		model:           O.Model,
		horsepower:      O.Horsepower,
		engine_Type:     O.Engine_Type,
		production_date: O.Production_Date,
		fuelType:        O.FuelType,
		quantity:        O.Quantity,
		price:           O.Price,
	}

	return car
}

// Method Implemnted in the product-car


func (c *Car) Name() string {
	return c.BrandName() + " " + c.Model()
}

// Quantity of cars available
func (c *Car) Quantity() int {
	return c.quantity
}

// Price of the car 
func (c *Car) Price() string {
	return fmt.Sprintf("%.2f", c.price)
}

// 
func (c *Car) Horsepower() string {
	return c.horsepower
}

// Returns the model of the car
func (c *Car) Model() string {
	return c.model
}

// color of the car
func (c *Car) Color() string {
	return c.color
}

// The brand of the car 
func (c *Car) BrandName() string {
	return c.brandName
}

// Returns the number of cars available after selling the a particluar quantity
func (c *Car) Sell(quantity int) {
	c.quantity = c.quantity - quantity
}

// Type of the car-engine
func (c *Car) Engine_Type() string {
	return c.engine_Type
}

//  the fuel-type of a Car.
func (c *Car) FuelType() string {
	return c.fuelType
}

// ProductType returns the product-type of a Car.
func (c *Car) ProductType() string {
	return c.productType
}


