package handlers

import (
	"net/http"
	"strconv"

	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (p *Products) GetOneProduct(c *gin.Context) {
	p.l.Println("Handle GET Products")
	c.Writer.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.Error(c.Writer, "ID must be an int", http.StatusBadRequest)
	}
	// obtain db client
	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	// fetch products from the data store
	lp, err := data.GetOneProduct(dbClient, id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	// show the products as output
	c.IndentedJSON(http.StatusOK, lp)
}
