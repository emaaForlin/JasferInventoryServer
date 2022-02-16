// Package classification of JasferInventoryAPI
//
// Documentation for JasferInventoryAPI
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
	"net/http"
	"strconv"
	"github.com/emaaForlin/JasferInventoryServer/data"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

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

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// Responses:
//	200: productsResponse

// GetProducts return the products from the database
func (p *Products) GetProducts(c *gin.Context) {
	p.l.Println("Handle GET Products")
	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	// fetch products from the data store
	lp := data.GetProducts(dbClient)
	// show the products as output
	c.IndentedJSON(http.StatusOK, lp)
}

func (p *Products) AddProducts(c *gin.Context) {
	p.l.Println("Handle POST products")
	
	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	prod := &data.Product{} 
	err := prod.FromJSON(c.Request.Body)
	if err != nil {
		http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		p.l.Println(err, prod)
		return
	}
	data.AddProduct(prod, dbClient)
}

func (p *Products) UpdateProducts(c *gin.Context) {
	p.l.Println("Handle PUT products")
	
	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	prod := &data.Product{}
	err := prod.FromJSON(c.Request.Body)
	if err != nil {
		http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		p.l.Println(err)
	}

	err = data.UpdateProduct(id, prod, dbClient)
	if err == data.ErrProductNotFound {
		http.Error(c.Writer, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(c.Writer, "Product not found", http.StatusInternalServerError)
		return
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database
// Responses:
//	201: noContent

// DeleteProducts deletes a product from the database
func (p *Products) DeleteProduct(c *gin.Context) {
	p.l.Println("Handle DELETE product")
	
	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}
	id, _ := strconv.Atoi(c.Param("id"))

	err := data.DeleteProduct(id, dbClient)
	if err != nil {
		http.Error(c.Writer, "Unable to delete unexistent product", http.StatusNotFound)
		p.l.Println(err)
	}
}