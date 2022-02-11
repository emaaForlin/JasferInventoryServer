package data

import (
	"encoding/json"
	//"strings"
	//"github.com/emaaForlin/JasferInventorySoftware/database"
	"fmt"
	"io"
	"time"
	"gorm.io/gorm"
)

/*
type Product struct {
	gorm.Model
	ID          int     	`json:"id"`
	Name        string  	`json:"name"`
	Description string  	`json:"description"`
	Price       float32 	`json:"price"`
	SKU         string  	`json:"sku"`
	CreatedAt   time.Time  	`json:"-"`
	UpdatedAt   time.Time  	`json:"-"`
	DeletedAt 	time.Time	`json:"-"`

}
*/
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


func GetProducts(db *gorm.DB) Products {
	prods := &Product{}
	db.Find(prods)

	return productList
}

func AddProduct(p *Product, db *gorm.DB) {
	p.ID = getNextID()
	p.SKU = generateSKU(p.ID, "AA", "BB")
	p.CreatedAt = time.Now().UTC()//.String()
	
	prod := Product{
		ID: p.ID,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		SKU: p.SKU,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	db.Create(&prod)
	// append to products list
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	oldProd, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	p.CreatedAt = oldProd.CreatedAt
	p.UpdatedAt = time.Now().UTC()//.String()
	productList[pos] = p
	return nil
}

func getNextID() int {
	//lastProd := productList[len(productList)-1]
	return len(productList) + 1
}

func generateSKU(id int, prefix, suffix string) string {
	sku := fmt.Sprintf("%s0%d%s", prefix, id, suffix)
	return sku
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}
/*
func searchBarProduct(desc string) (*Product, int, error){
	for i, p := range productList {
		if strings.Contains(p.Description, desc) || strings.Contains(p.Description, desc) {
			fmt.Println(p, i, nil)
		}
	}
	return nil, -1, ErrProductNotFound
}
*/
var productList = []*Product{}/*
	&Product{
		ID:          1,
		Name:        "Jean",
		Description: "Talle 42",
		Price:       2000,
		SKU:         "CN422000",
	},
	&Product{
		ID:          2,
		Name:        "Remera manga corta",
		Description: "Talle M",
		Price:       1500,
		SKU:         "MCMR1500",
	},
	&Product{
		ID:          3,
		Name:        "Short depor",
		Description: "Talle S",
		Price:       1250,
		SKU:         "SDS1250",
	},
	&Product{
		ID:          4,
		Name:        "Jean",
		Description: "Talle 44",
		Price:       2100,
		SKU:         "CN442100",
	},
	&Product{
		ID:          5,
		Name:        "Remera manga corta",
		Description: "Talle L",
		Price:       1750,
		SKU:         "MCL1500",
	},
	&Product{
		ID:          6,
		Name:        "Short depor",
		Description: "Talle M",
		Price:       1350,
		SKU:         "SDM1250",
	},
}
*/