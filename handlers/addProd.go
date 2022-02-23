package handlers

import (
	"net/http"

	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	err = data.AddProduct(prod, dbClient)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}
