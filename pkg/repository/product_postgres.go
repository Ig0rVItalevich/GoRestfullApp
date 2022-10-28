package repository

import (
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(input perfume.Product) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (title, content, count, cost, rating)
								VALUES ($1, $2, $3, $4, $5) RETURNING id`, productsTable)
	row := r.db.QueryRow(query, input.Title, input.Content, input.Count, input.Cost, input.Rating)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *ProductPostgres) GetById(id int) (perfume.Product, error) {
	var product perfume.Product

	query := fmt.Sprintf("SELECT id, title, content, count, cost, date_of_publication, rating FROM %s WHERE id = $1", productsTable)
	err := r.db.Get(&product, query, id)

	return product, err
}

func (r *ProductPostgres) GetAll() ([]perfume.Product, error) {
	var products []perfume.Product

	query := fmt.Sprintf("SELECT id, title, content, count, cost, date_of_publication, rating FROM %s", productsTable)
	err := r.db.Select(&products, query)

	return products, err
}

func (r *ProductPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", productsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *ProductPostgres) Update(id int, input perfume.UpdateProduct) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId += 1
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content = $%d", argId))
		args = append(args, *input.Content)
		argId += 1
	}

	if input.Count != nil {
		setValues = append(setValues, fmt.Sprintf("count = $%d", argId))
		args = append(args, *input.Count)
		argId += 1
	}

	if input.Cost != nil {
		setValues = append(setValues, fmt.Sprintf("cost = $%d", argId))
		args = append(args, *input.Cost)
		argId += 1
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating = $%d", argId))
		args = append(args, *input.Rating)
		argId += 1
	}

	setValuesStr := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", productsTable, setValuesStr, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
