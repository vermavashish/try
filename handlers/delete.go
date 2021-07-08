package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vermavashish/try/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
//
// responses:
//	201: noContent
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handel Delete Product ", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
