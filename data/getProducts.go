package data

import "gorm.io/gorm"

func GetProducts(db *gorm.DB) ([]*Product, error) {
	db.Find(&productList)
	return productList, db.Error
}
