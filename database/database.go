package database
import (
	"fmt"
	//"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connections struct {
	host	string
	port	int
	user	string
	pass	string
	dbname	string
}

func NewConnection(h string, p int, user string, pass string, db string) *Connections {
	return &Connections{h,p,user,pass,db}
}

func (c *Connections) Connect() (*gorm.DB, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&autocommit=true&parseTime=True&loc=Local", c.user, c.pass, c.host, c.port, c.dbname)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

