package service

import (
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type Authorization interface {
	CreateUser(user perfume.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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
	Update(id int, input perfume.UpdateReview, userId int) error
	Delete(id int) error
}

type LikeProduct interface {
	Create(input perfume.LikeProduct) (int, error)
	GetById(id int) (perfume.LikeProduct, error)
	GetAll(productId int) ([]perfume.LikeProduct, error)
	Update(id int, input perfume.UpdateLikeProduct, userId int) error
	Delete(id int) error
}

type LikeReview interface {
	Create(input perfume.LikeReview) (int, error)
	GetById(id int) (perfume.LikeReview, error)
	GetAll(reviewtId int) ([]perfume.LikeReview, error)
	Update(id int, input perfume.UpdateLikeReview, userId int) error
	Delete(id int) error
}

type Service struct {
	Authorization
	Product
	Order
	Review
	LikeReview
	LikeProduct
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
		Review:        NewReviewService(repos.Review),
		LikeProduct:   NewLikeProductService(repos.LikeProduct),
		LikeReview:    NewLikeReviewService(repos.LikeReview),
	}
}
