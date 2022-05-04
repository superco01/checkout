package service

import (
	"github.com/k0kubun/pp"
	m "github.com/superco01/checkout/model"
)

var products = m.Products{
	ProductList: make(map[string]m.Product),
}

func ProductList() {
	pp.Println(products)
}

func init() {
	products.ProductList["120P90"] = m.Product{
		SKU:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}

	products.ProductList["43N23P"] = m.Product{
		SKU:   "43N23P",
		Name:  "MacBook Pro",
		Price: 5399.99,
		Qty:   5,
	}

	products.ProductList["A304SD"] = m.Product{
		SKU:   "A304SD",
		Name:  "Alexa Speaker",
		Price: 109.50,
		Qty:   10,
	}

	products.ProductList["234234"] = m.Product{
		SKU:   "234234",
		Name:  "Rasberry Pi B",
		Price: 30.00,
		Qty:   2,
	}
}
