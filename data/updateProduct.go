package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func UpdateProduct(id int, p *Product, db *gorm.DB) error {
	oldProd, pos, err := findProduct(id, db)
	if err != nil {
		return err
	}
	p.ID = id
	p.CreatedAt = oldProd.CreatedAt
	p.UpdatedAt = time.Now().UTC()

	if p.ID != 0 && p.Name != "" && p.Price != 0 && p.SKU != "" && p.CreatedAt != p.UpdatedAt {
		err := db.Save(&p).Error
		productList[pos] = p
		return err
	}
	return fmt.Errorf("required values not found")
}
