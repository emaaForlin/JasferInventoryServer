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
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// swagger:route GET /products products listProducts
// Returns a list of products
// Responses:
//	200: productsResponse
// GetProducts return the products from the database
func (p *Products) GetProducts(c *gin.Context) {
	p.l.Println("Handle GET Products")
	c.Writer.Header().Add("Content-Type", "application/json")

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

// swagger:route POST /products products createProducts
// Create a new product
//
// Responses:
//	200: productsResponse
//	501: errorResponse

// AddProducts adds a new product to the database
func (p *Products) AddProduct(c *gin.Context) {
	p.l.Printf("[DEBUG] Inserting product")

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

// swagger:route PUT /products/{id} products editProduct
// Modifies a product
//
// Responses:
//	201: noContent
//	404: errorResponse
//
// UpdateProduct modifies a product that already exists
func (p *Products) UpdateProduct(c *gin.Context) {
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

// DeleteProductS deletes a product from the database
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
