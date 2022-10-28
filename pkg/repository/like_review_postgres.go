package repository

import (
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type LikeReviewPostgres struct {
	db *sqlx.DB
}

func NewLikeReviewPostgres(db *sqlx.DB) *LikeReviewPostgres {
	return &LikeReviewPostgres{db: db}
}

func (r *LikeReviewPostgres) Create(input perfume.LikeReview) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (mark, review_id, user_id) VALUES ($1, $2, $3) RETURNING id", likesReviewTable)
	row := r.db.QueryRow(query, input.Mark, input.ReviewId, input.UserId)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *LikeReviewPostgres) GetById(id int) (perfume.LikeReview, error) {
	var like perfume.LikeReview

	query := fmt.Sprintf("SELECT id, mark, date_of_publication, review_id, user_id FROM %s WHERE id = $1", likesReviewTable)
	err := r.db.Get(&like, query, id)

	return like, err
}

func (r *LikeReviewPostgres) GetAll(reviewId int) ([]perfume.LikeReview, error) {
	var likes []perfume.LikeReview

	query := fmt.Sprintf("SELECT id, mark, date_of_publication, review_id, user_id FROM %s WHERE review_id = $1", likesReviewTable)
	err := r.db.Select(&likes, query, reviewId)

	return likes, err
}

func (r *LikeReviewPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", likesReviewTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *LikeReviewPostgres) Update(id int, input perfume.UpdateLikeReview) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark = $%d", argId))
		args = append(args, *input.Mark)
		argId += 1
	}

	setValuesStr := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", likesReviewTable, setValuesStr, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
