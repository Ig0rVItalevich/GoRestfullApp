package repository

import (
	"errors"
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(input perfume.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (status, comment, user_id) VALUES ($1,$2, $3) RETURNING id", ordersTable)
	row := tx.QueryRow(query, input.Status, input.Comment, input.UserId)
	var orderId int
	err = row.Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, product := range input.Products {
		query := fmt.Sprintf("INSERT INTO %s (product_id, count, order_id) VALUES ($1, $2, $3)", ordersProductsTable)
		_, err := tx.Exec(query, product.Id, product.Count, orderId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	return orderId, err
}

func (r *OrderRepository) GetById(id int, userId int) (perfume.Order, error) {
	var order perfume.Order
	query := fmt.Sprintf(`SELECT id, status, date_of_publication, comment, user_id, 
           										date_of_completion FROM %s WHERE id = $1`, ordersTable)
	err := r.db.Get(&order, query, id)
	if err != nil {
		return order, err
	}

	if order.UserId != userId {
		return order, errors.New("user does not have permissions for this order")
	}

	var products []perfume.OrderProduct
	query = fmt.Sprintf(`SELECT product_id, count FROM %s WHERE order_id = $1`, ordersProductsTable)
	err = r.db.Select(&products, query, id)
	if err != nil {
		return order, err
	}

	order.Products = products

	return order, nil
}

func (r *OrderRepository) GetAll(userId int) ([]perfume.Order, error) {
	var orders []perfume.Order
	query := fmt.Sprintf(`SELECT id, status, date_of_publication, comment, user_id, 
           										date_of_completion FROM %s WHERE user_id = $1`, ordersTable)
	err := r.db.Select(&orders, query, userId)
	if err != nil {
		return orders, err
	}
	for i := 0; i < len(orders); i++ {
		var products []perfume.OrderProduct
		query := fmt.Sprintf(`SELECT product_id, count FROM %s WHERE order_id = $1`, ordersProductsTable)
		err := r.db.Select(&products, query, orders[i].Id)
		if err != nil {
			return orders, err
		}

		orders[i].Products = products
	}

	return orders, nil
}

func (r *OrderRepository) Update(id int, input perfume.UpdateOrder, userId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if input.Status != nil {
		query := fmt.Sprintf("UPDATE %s SET status = $1 WHERE id = $2", ordersTable)
		_, err := tx.Exec(query, input.Status, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if input.ProductsToBeDeleted != nil {
		for i := 0; i < len(input.ProductsToBeDeleted); i++ {
			var productDB perfume.OrderProduct
			query := fmt.Sprintf("SELECT product_id, count FROM %s WHERE order_id = $1 AND product_id = $2", ordersProductsTable)
			err := r.db.Get(&productDB, query, id, input.ProductsToBeDeleted[i].Id)
			if err != nil {
				logrus.Warn(err.Error())
				continue
			}

			if input.ProductsToBeDeleted[i].Count == 0 || productDB.Count-input.ProductsToBeDeleted[i].Count <= 0 {
				query := fmt.Sprintf("DELETE FROM %s WHERE order_id = $1 AND product_id = $2", ordersProductsTable)
				_, err := tx.Exec(query, id, input.ProductsToBeDeleted[i].Id)
				if err != nil {
					tx.Rollback()
					return err
				}
			} else {
				query := fmt.Sprintf("UPDATE %s SET count = $1 WHERE order_id = $2 AND product_id = $3", ordersProductsTable)
				_, err := tx.Exec(query, productDB.Count-input.ProductsToBeDeleted[i].Count, id, input.ProductsToBeDeleted[i].Id)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	if input.ProductsToBeAdded != nil {
		for i := 0; i < len(input.ProductsToBeAdded); i++ {
			var productDB perfume.OrderProduct
			query := fmt.Sprintf("SELECT product_id, count FROM %s WHERE order_id = $1 AND product_id = $2", ordersProductsTable)
			err := r.db.Get(&productDB, query, id, input.ProductsToBeAdded[i].Id)
			if err != nil {
				query := fmt.Sprintf("INSERT INTO %s (order_id, product_id, count) VALUES ($1, $2, $3)", ordersProductsTable)
				_, err := tx.Exec(query, id, input.ProductsToBeAdded[i].Id, input.ProductsToBeAdded[i].Count)
				if err != nil {
					tx.Rollback()
					return err
				}
			} else {
				query := fmt.Sprintf("UPDATE %s SET count = $1 WHERE order_id = $2 AND product_id = $3", ordersProductsTable)
				_, err := tx.Exec(query, productDB.Count+input.ProductsToBeAdded[i].Count, id, input.ProductsToBeAdded[i].Id)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	tx.Commit()
	return nil
}

func (r *OrderRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", ordersTable)
	_, err := r.db.Exec(query, id)

	return err
}
