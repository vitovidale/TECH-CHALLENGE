package service

import (
  "context"

  "github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
  "github.com/vitovidale/TECH-CHALLENGE/internal/core/port"
)

type ProductService struct {
 productRepository port.ProductRepository 
}

func NewProductService(productRepository port.ProductRepository) *ProductService {
  return &ProductService{productRepository: productRepository}
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*domain.Product, error) {
  return s.productRepository.FindProductByID(ctx, id)
}

func (s *ProductService) GetAll(ctx context.Context) ([]*domain.Product, error) {
  return s.productRepository.FindAllProducts(ctx)
}

func (s *ProductService) Create(ctx context.Context, p *domain.Product) error {
  return s.productRepository.Create(ctx, p)
}

func (s *ProductService) Inactivate(ctx context.Context, id int) error {
  p, err := s.productRepository.FindProductByID(ctx, id)
  if err != nil {
    return err
  }
  return p.Inactivate()
}

func (s *ProductService) Activate(ctx context.Context, id int) error {
  p, err := s.productRepository.FindProductByID(ctx, id)
  if err != nil {
    return err
  }
  return p.Activate()
}
