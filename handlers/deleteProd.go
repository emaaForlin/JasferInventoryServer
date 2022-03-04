package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryServer/auth"
	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database
// Responses:
//	201: noContent

// DeleteProductS deletes a product from the database
func (p *Products) DeleteProduct(c *gin.Context) {
	if auth.ValidateAPIKey(c) {
		p.l.Println("[INFO] Handle DeleteProduct")

		// obtain db client
		dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
		if !ok {
			p.l.Println("[ERROR] Error connection to the database")
			panic("error connection DB")
		}
		id, _ := strconv.Atoi(c.Param("id"))

		err := data.DeleteProduct(id, dbClient)
		if err == data.ErrProductNotFound {
			p.l.Println("[ERROR] Product not found")
			c.IndentedJSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		if err != nil {
			p.l.Printf("[ERROR] Something was wrong %s\n", err)
			c.IndentedJSON(http.StatusInternalServerError, map[string]string{"error": "This product cannot be deleted"})
		}
	}
}
