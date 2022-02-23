package handlers

import (
	"net/http"
	"strconv"

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
