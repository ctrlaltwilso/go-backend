package product

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
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

func GetProductById(conn *pgx.Conn, id int) (Product, error) {
	query := "SELECT * FROM products WHERE ID = $1"
	rows, _ := conn.Query(context.Background(), query, id)
	product, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Product])
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func ConnectDb(env string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), env)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
