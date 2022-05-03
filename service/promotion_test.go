package service

import (
	"log"
	"testing"

	m "github.com/superco01/checkout/model"

	"github.com/stretchr/testify/assert"
)

func TestPromotion_A304SD(t *testing.T) {
	sku := "A304SD"
	var expected float32 = 328.5

	products.ProductList[sku] = m.Product{
		SKU:   sku,
		Name:  "Alexa Speaker",
		Price: 109.50,
		Qty:   10,
	}

	var srv = InitCart()

	for i := 0; i < 3; i++ {
		srv.AddItem(sku)
	}
	log.Printf("cart1: %v", srv.Items)
	res, err := srv.Checkout()

	log.Printf("err: %v", err)
	assert.Equal(t, expected, res)
}

func TestPromotion_43N23P(t *testing.T) {
	sku := "43N23P"
	var expected float32 = 16199.971

	var srv = InitCart()

	for i := 0; i < 3; i++ {
		srv.AddItem(sku)
	}
	res, _ := srv.Checkout()

	assert.Equal(t, expected, res)
}

func TestPromotion_120P90(t *testing.T) {
	sku := "120P90"
	var expected float32 = 149.97

	var srv = InitCart()

	for i := 0; i < 3; i++ {
		srv.AddItem(sku)
	}
	res, _ := srv.Checkout()

	assert.Equal(t, expected, res)
}
