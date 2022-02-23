package data

import "gorm.io/gorm"

func GetProducts(db *gorm.DB) []*Product {
	db.Find(&productList)
	return productList
}
