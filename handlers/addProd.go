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
	p.l.Printf("[INFO] Handle AddProduct")

	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		p.l.Println("[ERROR] Error connection to the database")
		panic("error connection DB")
	}

	prod := &data.Product{}
	err := prod.FromJSON(c.Request.Body)
	if err != nil {
		p.l.Printf("[ERROR] Something was wrong %s\n", err)
		http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.AddProduct(prod, dbClient)
	if err != nil {
		p.l.Printf("[ERROR] Something was wrong %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": "Something bad has occurred, check where is the mistake"})
		return
	}
}
