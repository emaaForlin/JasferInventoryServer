package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID			int			`gorm:"primaryKey"			json: "id"`
	Name        string		`gorm:"type:varchar(64)"	json: "name"`
	Description string		`gorm:"type:varchar(64)"	json: "description"`
	Price       float32		`gorm:"type:float"			json: "price"`
	SKU         string 		`gorm:"type:varchar(8)"		json: "sku"`
	CreatedAt   time.Time	`gorm:"type:datetime"		json: "-"`
	UpdatedAt   time.Time	`gorm:"type:datetime"		json: "-"`
//	DeletedAt 	time.Time	`gorm:"type:datetime"		json: "-"`
}

type Products []*Product

var productList []*Product

// ####################################################
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
// ####################################################


func GetProducts(db *gorm.DB) []*Product {
	db.Find(&productList)
	return productList
}

func AddProduct(p *Product, db *gorm.DB) {
	p.ID = getNextID(db)
	p.SKU = generateSKU(p.ID, "AA", "BB")
	p.CreatedAt = time.Now().UTC()
	
	prod := Product{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		SKU: p.SKU,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	if prod.ID != 0 && prod.Name != "" && prod.Price != 0 && prod.SKU != "" {
		db.Create(&prod)
	}
}

func UpdateProduct(id int, p *Product, db *gorm.DB) error {
	oldProd, pos, err := findProduct(id, db)
	if err != nil {
		return err
	}
	p.ID = id
	p.CreatedAt = oldProd.CreatedAt
	p.UpdatedAt = time.Now().UTC()
	
	if p.ID != 0 && p.Name != "" && p.Price != 0 && p.SKU != "" && p.CreatedAt != p.UpdatedAt {
		db.Save(&p)
		productList[pos] = p
		return nil	
	}
	return fmt.Errorf("All values are needed")
}

func DeleteProduct(id int, db *gorm.DB) error {
	_, _, err := findProduct(id, db)
	if err != nil {
		return err
	}
	// this deletes permanently the entry
	db.Unscoped().Delete(&Product{}, id)
	return nil
}

func getNextID(db *gorm.DB) int {
	prod := &Product{}
	var p []Product
	db.Find(&p)

	for i:=1;i<len(productList);i++ {
		if i != p[i-1].ID {
			return i
		}
	}
	db.Last(prod)
	return prod.ID + 1
}

func generateSKU(id int, prefix, suffix string) string {
	// I need to improve this
	sku := fmt.Sprintf("%s0%d%s", prefix, id, suffix)
	return sku
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int, db *gorm.DB) (*Product, int, error) {
	prods := GetProducts(db)
	for i, p := range prods {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}