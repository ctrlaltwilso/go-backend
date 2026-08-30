package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Product struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

func GetProducts(conn *pgx.Conn) ([]Product, error) {
	rows, _ := conn.Query(context.Background(), "select * from products")
	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func main() {
	// Load env
	godotenv.Load(".env")

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Unable to connect to the product database")
	}

	data, err := GetProducts(conn)
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
