package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/emaaForlin/JasferInventorySoftware/data"
	//"github.com/emaaForlin/JasferInventorySoftware/database"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// getProducts return the products from the data store
func (p *Products) GetProducts(c *gin.Context) {
	p.l.Println("Handle GET Products")

	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	// fetch products from the data store
	lp := data.GetProducts(dbClient)
	// show the products as output
	c.IndentedJSON(http.StatusOK, lp)
}

func (p *Products) AddProducts(c *gin.Context) {
	p.l.Println("Handle POST products")

	dbClient, ok := c.MustGet("dbConn").(*gorm.DB)
	if !ok {
		panic("error connection DB")
	}

	prod := &data.Product{} 
	err := prod.FromJSON(c.Request.Body)
	if err != nil {
		http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod, dbClient)
}

func (p *Products) UpdateProducts(c *gin.Context) {
	p.l.Println("Handle PUT products")
	id, _ := strconv.Atoi(c.Param("id"))

	prod := &data.Product{}
	err := prod.FromJSON(c.Request.Body)

	if err != nil {
		http.Error(c.Writer, "Unable to unmarshal json", http.StatusBadRequest)
		p.l.Println(err)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(c.Writer, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(c.Writer, "Product not found", http.StatusInternalServerError)
		return
	}
}
