package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go-db/app/product"
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

}
