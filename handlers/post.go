package handlers

import (
	"net/http"

	"github.com/vermavashish/try/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productsResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Println(&prod)
	data.AddProducts(&prod)
}
