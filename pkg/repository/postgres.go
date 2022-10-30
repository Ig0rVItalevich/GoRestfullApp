package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	categoriesTable         = "categories"
	likesProductTable       = "likes_product"
	likesReviewTable        = "liker_review"
	ordersTable             = "orders"
	ordersProductsTable     = "orders_products"
	productsTable           = "products"
	productsCategoriesTable = "products_categories"
	reviewsTable            = "reviews"
	usersTable              = "users"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode = %s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
