package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/emaaForlin/JasferInventoryServer/database"
	"github.com/emaaForlin/JasferInventoryServer/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

func dbMiddleWare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbConn", db)
		c.Next()
	}
}

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	l := log.New(os.Stdout, "JISoftware-prototype: ", log.LstdFlags)
	//create the product handler
	productHandler := handlers.NewProduct(l)

	// connect database
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	int_port, err := strconv.Atoi(port)
	if err != nil {
		fmt.Errorf("Database port needs to be an int", err)
	}

	db := database.NewConnection(host, int_port, user, pass, dbname)
	client, err := db.Connect()

	if err != nil {
		panic(err)
	}

	// initialize gin router
	router := gin.Default()
	router.Use(CORSMiddleWare())
	router.Use(dbMiddleWare(client))

	// setup redoc middleware with options
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	rm := middleware.Redoc(opts, nil)

	// retrieve products
	router.GET("/products", productHandler.GetProducts)
	// retrieve one product by id
	router.GET("/products/:id", productHandler.GetOneProduct)
	// serve swagger file for redoc
	router.GET("/swagger.yaml", gin.WrapH(http.FileServer(http.Dir("./"))))
	// show docs
	router.GET("/docs", gin.WrapH(rm))
	// edit products
	router.PUT("/products/:id", productHandler.UpdateProduct)
	// add products
	router.POST("/products", productHandler.AddProduct)
	// delete products
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	// all the stuff needed to start serving the page are down here
	// setting up http server
	s := &http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  5 * time.Second,   // max time to read requests from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
	}

	go func() {
		err := s.ListenAndServe()
		l.Printf("Server listening on %s", s.Addr)

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
