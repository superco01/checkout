package service

import (
	"testing"

	m "github.com/superco01/checkout/model"

	"github.com/stretchr/testify/assert"
)

func TestCart_AddItem_Success(t *testing.T) {
	sku := "A304SD"
	expected := 2

	var srv = InitCart()

	for i := 0; i < 2; i++ {
		srv.AddItem(sku)
	}

	assert.Equal(t, expected, srv.Items[sku])
}

func TestCart_Checkout_Success(t *testing.T) {
	sku := "A304SD"
	var expected float32 = 219

	var srv = InitCart()

	for i := 0; i < 2; i++ {
		srv.AddItem(sku)
	}
	res, _ := srv.Checkout()

	ProductList()

	assert.Equal(t, expected, res)
}

func TestCart_AddItem_OutOfStock(t *testing.T) {
	products.ProductList["A304SD"] = m.Product{
		SKU:   "A304SD",
		Name:  "Alexa Speaker",
		Price: 109.50,
		Qty:   1,
	}
	expected := "product Alexa Speaker with SKU number A304SD is out of stock"
	sku := "A304SD"
	var res error

	var srv = InitCart()
	for i := 0; i < 2; i++ {
		res = srv.AddItem(sku)
	}

	assert.Equal(t, expected, res.Error())
}

func TestCart_AddItem_ProductNotFound(t *testing.T) {

	sku := "INVALIDSKU"
	expected := "product with SKU number " + sku + " not found"
	var res error

	var srv = InitCart()
	for i := 0; i < 1; i++ {
		res = srv.AddItem(sku)
	}

	assert.Equal(t, expected, res.Error())
}
