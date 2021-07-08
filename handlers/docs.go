// Package classification	 Product API
//
// Documentation of Product API
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//	   Host: localhost:9000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handlers

import "github.com/vermavashish/try/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponse struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of Product returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productNoContent struct {
}

// swagger:parameters deleteProduct listSingle updateProduct
type productIDParameter struct {
	//The id the product
	// in: path
	// required: true
	ID int `json:"id"`
}
