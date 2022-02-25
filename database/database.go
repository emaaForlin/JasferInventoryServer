package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	host   string
	port   int
	user   string
	pass   string
	dbname string
}

const tableScript string = `CREATE TABLE JIS.products (
	ID int unsigned NOT NULL,
	Name varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
	Description varchar(256) DEFAULT NULL,
	Price float DEFAULT NULL,
	SKU varchar(8) DEFAULT NULL,
	created_at timestamp NULL DEFAULT NULL,
	updated_at timestamp NULL DEFAULT NULL,
	deleted_at timestamp NULL DEFAULT NULL,
	PRIMARY KEY (ID) USING BTREE
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`

func NewConnection(h string, p int, user string, pass string, db string) *Connection {
	return &Connection{h, p, user, pass, db}
}

func (c *Connection) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&autocommit=true&parseTime=True&loc=Local", c.user, c.pass, c.host, c.port, c.dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if !checkTableExists("products", db) {
		db.Exec(tableScript)
	}
	return db, err
}

func checkTableExists(tableName string, db *gorm.DB) bool {
	var out string
	db.Raw("SHOW TABLES;").Row().Scan(&out)
	return out == tableName
}
