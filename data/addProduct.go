package data

import (
	"time"

	"gorm.io/gorm"
)

func AddProduct(p *Product, db *gorm.DB) error {
	p.ID = getNextID(db)
	p.CreatedAt = time.Now().UTC()

	prod := Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		SKU:         p.SKU,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   time.Date(1986, time.January, 1, 0, 0, 0, 0, time.Local),
	}
	if prod.ID != 0 && prod.Name != "" && prod.Price != 0 {
		db.Create(&prod)
	}
	return db.Error
}
