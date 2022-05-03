package main

import (
	"github.com/superco01/checkout/service"
)

func main() {
	var srv = service.InitCart()

	service.ProductList()
	srv.AddItem("A304SD")
	srv.AddItem("A304SD")
	srv.AddItem("A304SD")
	srv.AddItem("43N23P")
	srv.AddItem("43N23P")
	srv.AddItem("43N23P")
	srv.AddItem("120P90")
	srv.AddItem("120P90")

	srv.Checkout()

	service.ProductList()
}
