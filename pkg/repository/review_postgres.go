package repository

import (
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (r *ReviewPostgres) Create(input perfume.Review) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (content, rating, product_id, user_id) VALUES ($1, $2, $3, $4) RETURNING id", reviewsTable)
	row := r.db.QueryRow(query, input.Content, input.Rating, input.ProductId, input.UserId)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *ReviewPostgres) GetById(id int) (perfume.Review, error) {
	var review perfume.Review

	query := fmt.Sprintf("SELECT id, content, date_of_publication, rating, product_id, user_id FROM %s WHERE id = $1", reviewsTable)
	err := r.db.Get(&review, query, id)

	return review, err
}

func (r *ReviewPostgres) GetAll(productId int) ([]perfume.Review, error) {
	var reviews []perfume.Review

	query := fmt.Sprintf("SELECT id, content, rating, date_of_publication, product_id, user_id FROM %s WHERE product_id = $1", reviewsTable)
	err := r.db.Select(&reviews, query, productId)

	return reviews, err
}

func (r *ReviewPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", reviewsTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *ReviewPostgres) Update(id int, input perfume.UpdateReview) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content = $%d", argId))
		args = append(args, *input.Content)
		argId += 1
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating = $%d", argId))
		args = append(args, *input.Rating)
		argId += 1
	}

	setValuesStr := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", reviewsTable, setValuesStr, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
