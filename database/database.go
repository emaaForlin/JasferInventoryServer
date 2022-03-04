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

const creationScript1 string = `CREATE TABLE products (
	ID int unsigned NOT NULL,
	Name varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
	Description varchar(256) DEFAULT NULL,
	Price float DEFAULT NULL,
	SKU varchar(8) DEFAULT NULL,
	created_at timestamp NULL DEFAULT NULL,
	updated_at timestamp NULL DEFAULT NULL,
	deleted_at timestamp NULL DEFAULT NULL,
	PRIMARY KEY (ID) USING BTREE
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;`

const creationScript2 string = `CREATE TABLE authorized_users (
	ID INT(10) NOT NULL AUTO_INCREMENT,
	User VARCHAR(32) NOT NULL COLLATE 'utf8mb4_general_ci',
	Apikey VARCHAR(128) NOT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (ID) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2;`

func NewConnection(h string, p int, user string, pass string, db string) *Connection {
	return &Connection{h, p, user, pass, db}
}

func (c *Connection) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&autocommit=true&parseTime=True&loc=Local", c.user, c.pass, c.host, c.port, c.dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if !checkTableExists("products", db) {
		db.Exec(creationScript1)
	}
	if !checkTableExists("authorized_users", db) {
		db.Exec(creationScript2)
	}
	return db, err
}

func checkTableExists(tableName string, db *gorm.DB) bool {
	var out []string
	db.Raw("SHOW TABLES;").Scan(&out)
	for _, i := range out {
		return i == tableName
	}
	return false
}
