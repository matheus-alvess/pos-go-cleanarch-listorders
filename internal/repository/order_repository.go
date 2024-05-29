package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"pos-go-cleanarch-listorders/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) ListOrders() ([]model.Order, error) {
	rows, err := r.db.Query("SELECT id, price, tax FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.IDField, &order.PriceField, &order.TaxField); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *OrderRepository) CreateOrder(price, tax float64) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO orders (price, tax) VALUES ($1, $2) RETURNING id", price, tax).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	return sql.Open("postgres", connStr)
}
