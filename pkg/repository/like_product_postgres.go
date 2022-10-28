package repository

import (
	"fmt"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type LikeProductPostgres struct {
	db *sqlx.DB
}

func NewLikeProductPostgres(db *sqlx.DB) *LikeProductPostgres {
	return &LikeProductPostgres{db: db}
}

func (r *LikeProductPostgres) Create(input perfume.LikeProduct) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (mark, product_id, user_id) VALUES ($1, $2, $3) RETURNING id", likesProductTable)
	row := r.db.QueryRow(query, input.Mark, input.ProductId, input.UserId)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *LikeProductPostgres) GetById(id int) (perfume.LikeProduct, error) {
	var like perfume.LikeProduct

	query := fmt.Sprintf("SELECT id, mark, date_of_publication, product_id, user_id FROM %s WHERE id = $1", likesProductTable)
	err := r.db.Get(&like, query, id)

	return like, err
}

func (r *LikeProductPostgres) GetAll(productId int) ([]perfume.LikeProduct, error) {
	var likes []perfume.LikeProduct

	query := fmt.Sprintf("SELECT id, mark, date_of_publication, product_id, user_id FROM %s WHERE product_id = $1", likesProductTable)
	err := r.db.Select(&likes, query, productId)

	return likes, err
}

func (r *LikeProductPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", likesProductTable)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *LikeProductPostgres) Update(id int, input perfume.UpdateLikeProduct) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark = $%d", argId))
		args = append(args, *input.Mark)
		argId += 1
	}

	setValuesStr := strings.Join(setValues, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", likesProductTable, setValuesStr, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
