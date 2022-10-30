package service

import (
	"errors"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type OrderService struct {
	repos repository.Order
}

func NewOrderService(repos repository.Order) *OrderService {
	return &OrderService{repos: repos}
}

func (s *OrderService) Create(input perfume.Order) (int, error) {
	return s.repos.Create(input)
}

func (s *OrderService) GetById(id int, userId int) (perfume.Order, error) {
	return s.repos.GetById(id, userId)
}

func (s *OrderService) GetAll(userId int) ([]perfume.Order, error) {
	return s.repos.GetAll(userId)
}

func (s *OrderService) Update(id int, input perfume.UpdateOrder, userId int) error {
	_, err := s.GetById(id, userId)
	if err != nil {
		return err
	}

	if !input.Validate() {
		return errors.New("information for update doesn't exist")
	}

	return s.repos.Update(id, input, userId)
}

func (s *OrderService) Delete(id int, userId int) error {
	_, err := s.GetById(id, userId)
	if err != nil {
		return err
	}

	return s.repos.Delete(id)
}
