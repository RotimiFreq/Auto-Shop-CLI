package main

import (
	"Auto_cli/Product"
	"Auto_cli/store"

	
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer = os.Stdout

	availableProducts := make(store.Store)
	soldProducts := make(store.Store)

	var p = []Product.Order_Car{
		{
			ProductType:     "Car",
			BrandName:       "Toyota",
			Model:           "Corolla",
			Color:           "black",
			FuelType:        "petrol",
			Quantity:        35,
			Price:           3250000.00,
			Engine_Type:     "2.0L 4-cyl",
			Horsepower:      "169hp/151 lb-ft",
			Production_Date: "January-2020",
		},
		{
			ProductType:     "Car",
			BrandName:       "Toyota",
			Model:           "Avalon",
			Color:           "red",
			FuelType:        "petrol",
			Quantity:        22,
			Price:           105000000.00,
			Engine_Type:     "2.5L 4-cyl",
			Horsepower:      "340hp/400 lb-ft",
			Production_Date: "January-2020",
		},
	}
	//Demo

	fmt.Fprintln(w, "DEMO for CLI-SHOP")

	// Add products to the store
	fmt.Fprintln(w, "Add product")
	ids := availableProducts.AddProduct(w, &p)
	for i := 0; i < len(ids); i++ {
		fmt.Printf("You added product %v with id %v\n", p[i].BrandName+" "+p[i].Model, ids[i])
	}


	// 	Lists Products
	fmt.Fprintln(w, "List product")
	availableProducts.ListProducts(w)

	// Sell Products
	fmt.Fprintln(w, "Sell product")
	availableProducts.SellProduct(w, ids[0], 10, &soldProducts)

	fmt.Fprintln(w)

	availableProducts.SellProduct(w, ids[1], 20, &soldProducts)

	// List sold items
	fmt.Fprintln(w, "List sold product")
	
	availableProducts.ListSoldItems(w, &soldProducts)

	fmt.Fprintln(w, "End of demo")
}
