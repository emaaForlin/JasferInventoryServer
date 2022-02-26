package data_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/emaaForlin/JasferInventoryServer/data"
	"github.com/emaaForlin/JasferInventoryServer/database"
	"gorm.io/gorm"
)

var myTestProd = data.Product{
	Name:        "Test before compile",
	Description: "This product will be automatically deleted",
	Price:       111111111111111111111111111111111111111,
	SKU:         "TESTRUNN",
	CreatedAt:   time.Now(),
}
var myTestEditedProd = data.Product{
	Name:        "Test edit before compile",
	Description: "This product was edited and it will be automatically deleted",
	Price:       99999999999999999999999999999999999999,
	SKU:         "TESTPASS",
}

func connectDB() *gorm.DB {
	// connect database
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	user := os.Getenv("TEST_DB_USER")
	pass := os.Getenv("TEST_DB_PASS")
	dbname := os.Getenv("TEST_DB_NAME")

	int_port, err := strconv.Atoi(port)
	if err != nil {
		fmt.Errorf("Database port needs to be an int %q", err)
	}

	db := database.NewConnection(host, int_port, user, pass, dbname)
	client, err := db.Connect()

	if err != nil {
		return nil
	}
	return client
}

func TestConnectDB(t *testing.T) {
	client := connectDB()
	if client.Error != nil {
		t.Fatal(client.Error)
	}
}

var client = connectDB()

func TestAddProd(t *testing.T) {
	err := data.AddProduct(&myTestProd, client)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetProds(t *testing.T) {
	_, err := data.GetProducts(client)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOneProd(t *testing.T) {
	_, err := data.GetOneProduct(client, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEditProd(t *testing.T) {
	myTestEditedProd.ID = 1
	myTestEditedProd.CreatedAt = myTestProd.CreatedAt
	myTestEditedProd.UpdatedAt = time.Now()
	err := data.UpdateProduct(myTestEditedProd.ID, &myTestEditedProd, client)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteProd(t *testing.T) {
	err := data.DeleteProduct(1, client)
	if err != nil {
		t.Error(err)
	}
}
