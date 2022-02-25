package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"gorm.io/gorm"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id" gorm:"primaryKey"`

	// The name of the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" gorm:"type:varchar(64)"`

	// The description of the product
	//
	// required: false
	// max length: 1000
	Description string `json:"description" gorm:"type:varchar(64)"`

	// The price of the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" gorm:"type:float"`

	// The intern unique code of the product
	//
	// required: true
	// max length: 8
	SKU string `json:"sku" gorm:"type:varchar(8)"`

	CreatedAt time.Time `json:"-" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"-" gorm:"type:datetime"`
}

type Products []*Product

var productList []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

// for encoding json
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// for decoding json
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func getNextID(db *gorm.DB) (int, error) {
	prod := &Product{}
	var p []Product
	db.Find(&p)

	for i := 1; i < len(productList); i++ {
		if i != p[i-1].ID {
			return i, nil
		}
	}
	db.Last(prod)
	return prod.ID + 1, db.Error
}

func findProduct(id int, db *gorm.DB) (*Product, int, error) {
	prods, _ := GetProducts(db)
	for i, p := range prods {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}
