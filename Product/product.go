package Product

import (
	"fmt"
	"io"
)

// The main attributes of products that can be in the store.

type StoreProduct interface {
	Name() string

	Quantity() int

	Price() string

	Model() string
	
	Color()string

	BrandName() string

	ProductType() string

	Sell(quantity int)

	Engine_Type() string

	FuelType() string

	Horsepower() string
}

type Product struct {
	ID string
	StoreProduct
}

func (c *Product) DisplayProduct(w io.Writer) {
	fmt.Fprintln(w, "ID: ", c.ID)
	fmt.Fprintln(w, "Product Type: ", c.ProductType())
	fmt.Fprintln(w, "Brand: ", c.BrandName())
	fmt.Fprintln(w, "Model: ", c.Model())
	fmt.Fprintln(w, "Product Name: ", c.Name())
	fmt.Fprintln(w, "Quantity: ", c.Quantity())

	// Calling the unique method of the products

	switch u := c.StoreProduct.(type) {

	// After adding a new product, store manager should add a new 'switch Case' for it
	// and call its unique methods with a "fmt.Fprintln(w, )" funtion.

	case *Car:
		fmt.Fprintln(w, "Color: ", u.Color())
		fmt.Fprintln(w, "Engine_type: ", u.Engine_Type())
		fmt.Fprintln(w, "Fuel_type: ", u.FuelType())
		fmt.Fprintln(w, "HorsePower: ", u.Horsepower())

	}

	fmt.Println()

	fmt.Println()
}

// method to confirm if a product in store
func (c *Product) InStock() bool {
	return c.Quantity() > 1
}
