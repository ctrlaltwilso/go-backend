package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"go-db/app/product"

	"github.com/joho/godotenv"
)

func main() {
	// Load env
	godotenv.Load(".env")

	// Connect to the database
	conn, err := product.ConnectDb(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Unable to connect ot the product database")
	}

	data, err := product.GetProducts(conn)
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, product := range data {
		println("ID:", product.ID)
		println("Name:", product.Name)
		println("Description:", product.Description)
		println("Created:", product.CreatedAt.Format("01-02-2006 15:04"))
	}

	oneProduct, err := product.GetProductById(conn, 1)
	if err != nil {
		log.Fatal("Unable to return product by ID")
	}

	println("Return One Product")
	println("ID:", oneProduct.ID)
	println("Name:", oneProduct.Name)
	println("Description:", oneProduct.Description)
	println("Created:", oneProduct.CreatedAt.Format("01-02-2006 15:04"))

	// HTTP handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
		}

		fmt.Fprintf(w, "Welcome to home page")
	})

	// HTTP Server Startup
	adr := os.Getenv("HTTP_PORT")
	var srv http.Server
	srv.Addr = adr
	srv.Handler = mux

	listener, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening on port %s", adr)
	}

	// Shutdown HTTP server
	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Shutting down...")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(idleConnClosed)
	}()

	srv.Serve(listener)

	<-idleConnClosed
}
