package db

import (
	"database/sql"
	"log"

	pb "52HW/gen"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	connStr := "user=username dbname=yourdbname sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveTransaction(transaction *pb.SalesTransaction) error {
	query := `INSERT INTO sales_transactions (transaction_id, product_id, quantity, price, timestamp) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, transaction.TransactionId, transaction.ProductId, transaction.Quantity, transaction.Price, transaction.Timestamp)
	return err
}
