package repository

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user perfume.User) (int, error)
	GetUser(username, password string) (perfume.User, error)
}

type Product interface {
	Create(input perfume.Product) (int, error)
	GetById(id int) (perfume.Product, error)
	GetAll() ([]perfume.Product, error)
	Update(id int, input perfume.UpdateProduct) error
	Delete(id int) error
}

type Order interface {
}

type Review interface {
	Create(input perfume.Review) (int, error)
	GetById(id int) (perfume.Review, error)
	GetAll(productId int) ([]perfume.Review, error)
	Update(id int, input perfume.UpdateReview) error
	Delete(id int) error
}

type LikeProduct interface {
	Create(input perfume.LikeProduct) (int, error)
	GetById(id int) (perfume.LikeProduct, error)
	GetAll(productId int) ([]perfume.LikeProduct, error)
	Update(id int, input perfume.UpdateLikeProduct) error
	Delete(id int) error
}

type LikeReview interface {
	Create(input perfume.LikeReview) (int, error)
	GetById(id int) (perfume.LikeReview, error)
	GetAll(reviewtId int) ([]perfume.LikeReview, error)
	Update(id int, input perfume.UpdateLikeReview) error
	Delete(id int) error
}

type Repository struct {
	Authorization
	Product
	Order
	Review
	LikeReview
	LikeProduct
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductPostgres(db),
		Review:        NewReviewPostgres(db),
		LikeProduct:   NewLikeProductPostgres(db),
		LikeReview:    NewLikeReviewPostgres(db),
	}
}
