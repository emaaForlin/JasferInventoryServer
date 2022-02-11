package main
import (
	"net/http"
	"context"
	"time"
	"log"
	"fmt"
	"os"
	"os/signal"
	"github.com/gin-gonic/gin"
	"github.com/emaaForlin/JasferInventorySoftware/handlers"
	"github.com/emaaForlin/JasferInventorySoftware/database"
//	"github.com/emaaForlin/JasferInventorySoftware/views"
	"gorm.io/gorm"
	"strconv"
)

func apiMiddleWare(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        c.Set("dbConn", db)
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

	// create the index page view
	//index := views.NewView(l, "index.html", data.GetProducts())

	// initialize gin router
	router := gin.Default()
	router.Use(apiMiddleWare(client))
	// load static templates
	//router.Static("/assets", "./views/assets")
	//router.LoadHTMLFiles("./views/index.html")
	
	// render the index
	//router.GET("/", index.Render)
	// retrieve products
	router.GET("/products", productHandler.GetProducts)
	// edit products
	router.PUT("/products/:id", productHandler.UpdateProducts)
	// add products
	router.POST("/products", productHandler.AddProducts)


	// all the stuff needed to start serving the page are down here
	// setting up http server
	s := &http.Server{
		Addr: ":9090",
		Handler: router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
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

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc)
}