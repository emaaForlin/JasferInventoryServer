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
	p.l.Println("[INFO] Handle GetProduct")
	c.Writer.Header().Add("Content-Type", "application/json")

	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		p.l.Println("[ERROR] Error connection to the database")
		panic("error connection DB")
	}

	// fetch products from the data store
	lp, err := data.GetProducts(dbClient)
	if err != nil {
		p.l.Printf("[ERROR] Something was wrong %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": "Something bad has occurred, check where is the mistake"})
		return
	}
	// show the products as output
	c.IndentedJSON(http.StatusOK, lp)
}
