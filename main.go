package main
import (
	"net/http"
	"context"
	"time"
	"log"
	"os"
	"os/signal"
	"github.com/gin-gonic/gin"
	"github.com/emaaForlin/JasferInventorySoftware/handlers"
	"github.com/emaaForlin/JasferInventorySoftware/data"
	"github.com/emaaForlin/JasferInventorySoftware/views"
)

func main() {
	l := log.New(os.Stdout, "JISoftware-prototype: ", log.LstdFlags)
	productHandler := handlers.NewProduct(l)

	// index page
	index := views.NewView(l, "index.html", data.GetProducts())

	router := gin.Default()

	router.Static("/assets/bootstrap", "./views/assets/bootstrap")
	router.Static("/assets/js", "./views/assets/js")

	router.LoadHTMLFiles("./views/index.html", "./views/edit.html")
	router.GET("/", index.Render)


	router.GET("/products", productHandler.GetProducts)
	router.PUT("/products/:id", productHandler.UpdateProducts)
	router.POST("/products", productHandler.AddProducts)
	
	s := &http.Server{
		Addr: ":9090",
		Handler: router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	
	//s.ListenAndServe()
	
	l.Printf("Server listening on %s", s.Addr)


	go func() {
		err := s.ListenAndServe()
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