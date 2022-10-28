package service

import (
	"errors"
	perfume "github.com/Ig0rVItalevich/models"
	"github.com/Ig0rVItalevich/pkg/repository"
)

type ProductService struct {
	repos repository.Product
}

func NewProductService(repos repository.Product) *ProductService {
	return &ProductService{repos: repos}
}

func (s *ProductService) Create(input perfume.Product) (int, error) {
	return s.repos.Create(input)
}

func (s *ProductService) GetById(id int) (perfume.Product, error) {
	return s.repos.GetById(id)
}

func (s *ProductService) GetAll() ([]perfume.Product, error) {
	return s.repos.GetAll()
}

func (s *ProductService) Update(id int, input perfume.UpdateProduct) error {
	if !input.Validate() {
		return errors.New("information for update doesn't exist")
	}
	return s.repos.Update(id, input)
}

func (s *ProductService) Delete(id int) error {
	return s.repos.Delete(id)
}
