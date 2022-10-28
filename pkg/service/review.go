package service

import (
	"errors"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type ReviewService struct {
	repos repository.Review
}

func NewReviewService(repos repository.Review) *ReviewService {
	return &ReviewService{repos: repos}
}

func (s *ReviewService) Create(input perfume.Review) (int, error) {
	return s.repos.Create(input)
}

func (s *ReviewService) GetById(id int) (perfume.Review, error) {
	return s.repos.GetById(id)
}

func (s *ReviewService) GetAll(productId int) ([]perfume.Review, error) {
	return s.repos.GetAll(productId)
}

func (s *ReviewService) Update(id int, input perfume.UpdateReview, userId int) error {
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

func (s *ReviewService) Delete(id int) error {
	return s.repos.Delete(id)
}
