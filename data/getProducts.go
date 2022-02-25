package data

import "gorm.io/gorm"

func GetProducts(db *gorm.DB) ([]*Product, error) {
	err := db.Find(&productList).Error
	return productList, err
}
