package handlers

import (
	"net/http"

	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	lp, err := data.GetProducts(dbClient)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	// show the products as output
	c.IndentedJSON(http.StatusOK, lp)
}
