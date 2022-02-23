package handlers

import (
	"net/http"
	"strconv"

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
