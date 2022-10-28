package service

import (
	"errors"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type LikeReviewService struct {
	repos repository.LikeReview
}

func NewLikeReviewService(repos repository.LikeReview) *LikeReviewService {
	return &LikeReviewService{repos: repos}
}

func (s *LikeReviewService) Create(input perfume.LikeReview) (int, error) {
	return s.repos.Create(input)
}

func (s *LikeReviewService) GetById(id int) (perfume.LikeReview, error) {
	return s.repos.GetById(id)
}

func (s *LikeReviewService) GetAll(productId int) ([]perfume.LikeReview, error) {
	return s.repos.GetAll(productId)
}

func (s *LikeReviewService) Update(id int, input perfume.UpdateLikeReview, userId int) error {
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

func (s *LikeReviewService) Delete(id int) error {
	return s.repos.Delete(id)
}
