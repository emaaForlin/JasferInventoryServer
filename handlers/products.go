// Package classification of Products API
//
// Documentation for Products API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
)

type Products struct {
	l *log.Logger
}

type GenericError struct {
	Message string `json:"message"`
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}
