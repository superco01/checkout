package service

import (
	"fmt"
)

type Cart struct {
	Items map[string]int
}

func InitCart() *Cart {
	return &Cart{
		Items: make(map[string]int),
	}
}

func (cart Cart) AddItem(key string) error {
	product, found := products.ProductList[key]

	if !found {
		return fmt.Errorf("product with SKU number %v not found", key)
	}

	if product.Qty < 1 {
		return fmt.Errorf("product %v with SKU number %v is out of stock", product.Name, key)
	}
	//update product stock in db
	product.Qty -= 1
	products.ProductList[key] = product

	//update product item in cart
	item, found := cart.Items[key]
	if found {
		cart.Items[key] = item + 1
	} else {
		cart.Items[key] = 1
	}

	return nil
}

func (cart Cart) Checkout() (float32, error) {
	fmt.Println("=====Checkout=====")
	var total float32

	for key, v := range cart.Items {
		product := products.ProductList[key]
		total = total + (product.Price * float32(v))
	}

	fmt.Printf("Total price: \t\t $%.2f \n", total)

	discountPrice, promoDesc := cart.ApplyPromo()
	cart.printItems()
	fmt.Println("Promo description: \n" + promoDesc)
	fmt.Printf("Total discount: \t\t $%.2f \n", discountPrice)

	fmt.Printf("Total price after discount: \t $%.2f \n", (total - discountPrice))

	return total, nil
}

func (cart Cart) printItems() {
	for key, v := range cart.Items {
		product, found := products.ProductList[key]
		if found {
			fmt.Printf("%v \t %vx \t $%.2f \n", product.Name, v, product.Price)
		}
	}

}
