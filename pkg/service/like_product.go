package service

import (
	"errors"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type LikeProductService struct {
	repos repository.LikeProduct
}

func NewLikeProductService(repos repository.LikeProduct) *LikeProductService {
	return &LikeProductService{repos: repos}
}

func (s *LikeProductService) Create(input perfume.LikeProduct) (int, error) {
	return s.repos.Create(input)
}

func (s *LikeProductService) GetById(id int) (perfume.LikeProduct, error) {
	return s.repos.GetById(id)
}

func (s *LikeProductService) GetAll(productId int) ([]perfume.LikeProduct, error) {
	return s.repos.GetAll(productId)
}

func (s *LikeProductService) Update(id int, input perfume.UpdateLikeProduct, userId int) error {
	review, err := s.GetById(id)
	if err != nil {
		return err
	}

	if review.UserId != userId {
		return errors.New("no rights to change")
	}

	if !input.Validate() {
		return errors.New("information for update doesn't exist")
	}
	return s.repos.Update(id, input)
}

func (s *LikeProductService) Delete(id int) error {
	return s.repos.Delete(id)
}
