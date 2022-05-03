package service

import (
	"fmt"
)

func (cart Cart) ApplyPromo() (float32, string) {
	var totalDiscountPrice float32
	var promoDescription string

	discountPrice, desc := cart.getBonusItem()
	if desc != "" {
		totalDiscountPrice += discountPrice
		promoDescription += desc + "\n"
	}

	discountPrice, desc = cart.getDiscount()
	if desc != "" {
		totalDiscountPrice += discountPrice
		promoDescription += desc + "\n"
	}

	discountPrice, desc = cart.payLessItem()
	if desc != "" {
		totalDiscountPrice += discountPrice
		promoDescription += desc + "\n"
	}

	return totalDiscountPrice, promoDescription
}

func (cart Cart) getDiscount() (float32, string) {

	//discount promo rule
	discountedItem := "A304SD"
	minPurchasedItem := 3
	var discount float32 = 0.1

	// promo desc
	desc := fmt.Sprintf("Buying more than %v %vs will have a %v %% discount on all %vs", minPurchasedItem, products.ProductList[discountedItem].Name, (discount * 100), products.ProductList[discountedItem].Name)

	item := cart.Items[discountedItem]

	if item >= minPurchasedItem {
		return products.ProductList[discountedItem].Price * float32(item) * float32(discount), desc
	}
	return 0, ""
}

func (cart Cart) getBonusItem() (float32, string) {

	// bonus item promo rule
	baseItem := "43N23P"
	bonusItem := "234234"

	// promo desc
	desc := fmt.Sprintf("Each sale of a %v comes with a free %v", products.ProductList[baseItem].Name, products.ProductList[bonusItem].Name)

	item := cart.Items[baseItem]

	if item > 0 {
		for i := 0; i < item; i++ {
			cart.AddItem(bonusItem)
		}
		return products.ProductList["234234"].Price * float32(item), desc
	}
	return 0, ""
}

func (cart Cart) payLessItem() (float32, string) {

	// pay less item promo rule
	promoItem := "120P90"
	minPurchasedItem := 3
	paidItem := 2

	// promo desc
	desc := fmt.Sprintf("Buy %v %vs for the price of %v", minPurchasedItem, products.ProductList[promoItem].Name, paidItem)

	item := cart.Items[promoItem]

	if item >= minPurchasedItem {
		unpaidItem := item / minPurchasedItem
		return products.ProductList[promoItem].Price * float32(unpaidItem), desc
	}
	return 0, ""
}
