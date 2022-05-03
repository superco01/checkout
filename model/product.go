package model

type Product struct {
	SKU   string
	Name  string
	Price float32
	Qty   int
}

type Products struct {
	ProductList map[string]Product
}
