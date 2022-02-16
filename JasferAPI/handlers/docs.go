// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers
import "github.com/emaaForlin/JasferInventoryServer/data"

// NOTE: all this types are created only for documentation porpouses
// these types are not used for any handler

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to delete from database
	// in: path
	// required: true
	ID int	`json:"id"`
}

// swagger:response noContent
type productNoContentWrapper struct {}