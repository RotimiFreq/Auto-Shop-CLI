package store

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"Auto_cli/Product"
	"github.com/google/uuid"
)


type (

	// Store keep records of both available and sold products
	Store map[string] *Product.Product
	// Display the Sold Items
	SoldItemsDisplay struct {
		Total_Amount_Sales float64     `json:"total_amount_sold"`
		SoldItems  []SoldItems `json:"products_sold"`
	}

	// Sub-model to display sold items
	SoldItems struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		QuantitySold int    `json:"quantity_sold"`
		PriceSold    string `json:"price_sold"`
	}

)

// Method to list all products in store
func (s *Store) ListProducts(w io.Writer) {
	for _, v := range *s {
		if v.InStock() {
			v.DisplayProduct(w)
		}
	}
}
// Method to display products available for store
func (s *Store) NoOfProductsForSale(w io.Writer) {
	quantity := 0

	for _, v := range *s {
		if v.InStock() {
			quantity++
		}
	}

	fmt.Fprintf(w, "There are %v products up for sale\n", quantity)
}

// Method to add products to the store

func (s *Store) AddProduct(w io.Writer, order *[]Product.Order_Car) []string {
	fmt.Fprintln(w, "Adding products to store...")
	ids := make([]string, 0)

	for _, v := range *order {
		car := Product.CarOrderToCar(&v)
		id := uuid.New()

		var p = &Product.Product{
			ID:           id.String(),
			StoreProduct: car,
		}

		(*s)[id.String()] = p
		ids = append(ids, id.String())
	}

	fmt.Fprintf(w, "Added %v products successfully...\n", len(*order))

	return ids
}


// Methods to sell products in the store
func (s *Store) SellProduct(w io.Writer, id string, quantity int, salesStore *Store) {
	if _, ok := (*s)[id]; ok {
		(*s)[id].Sell(quantity)
	} else {
		fmt.Fprintln(w, "Product you specified does not exist")
		return
	}

	if _, ok := (*salesStore)[id]; !ok {
		(*salesStore)[id] = (*s)[id]
		(*salesStore)[id].Sell((*s)[id].Quantity() - quantity)
	} else {
		(*salesStore)[id].Sell(-quantity)
	}

	fmt.Fprintf(w, "Congrats! You just sold %v %v, with id %v\n", quantity, (*s)[id].Name(), id)
}

// Methods to list the sold products
func (s *Store) ListSoldItems(w io.Writer, salesStore *Store) {
	si := make([]SoldItems, 0)
	totalPrice := 0.0

	for _, v := range *salesStore {
		s := SoldItems{
			ID:           v.ID,
			Name:         v.Name(),
			QuantitySold: v.Quantity(),
			PriceSold:    v.Price(),
		}

		si = append(si, s)

		price, _ := strconv.ParseFloat(v.Price(), 64)
		totalPrice += price
	}

	siDisplay := SoldItemsDisplay{
		Total_Amount_Sales: totalPrice,
		SoldItems:  si,
	}

	jsonDisplay, err := json.MarshalIndent(siDisplay, "", " ")
	if err != nil {
		fmt.Fprintln(w, "error: Could not encode JSON")
		return
	}

	fmt.Fprintln(w, string(jsonDisplay))
}
