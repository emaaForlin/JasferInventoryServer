package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryServer/auth"
	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// swagger:route PUT /products/{id} products editProduct
// Modifies a product
//
// Responses:
//	201: noContent
//	404: errorResponse
//
// UpdateProduct modifies a product that already exists
func (p *Products) UpdateProduct(c *gin.Context) {
	if auth.ValidateAPIKey(c) {
		p.l.Println("[INFO] Handle UpdateProduct")
		// obtain db client
		dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
		if !ok {
			p.l.Println("[ERROR] Error connection to the database")
			panic("error connection DB")
		}

		id, _ := strconv.Atoi(c.Param("id"))

		prod := &data.Product{}
		err := prod.FromJSON(c.Request.Body)
		if err != nil {
			p.l.Println("[ERROR] Unable to unmarshal json")
			http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		}

		err = data.UpdateProduct(id, prod, dbClient)
		if err == data.ErrProductNotFound {
			p.l.Println("[ERROR] Product not found")
			c.IndentedJSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
			return
		}
		if err != nil {
			p.l.Printf("[ERROR] Something was wrong %s\n", err)
			c.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": "Something bad has occurred, check the required values"})
			return
		}
	}
}
