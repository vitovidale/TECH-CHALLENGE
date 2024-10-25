package service

import (
	"context"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/port"
)

type ProductService struct {
	categoryRepository port.CategoryRepository
	productRepository  port.ProductRepository
}

func NewProductService(categoryRepository port.CategoryRepository, productRepository port.ProductRepository) *ProductService {
	return &ProductService{categoryRepository: categoryRepository, productRepository: productRepository}
}

func (s *ProductService) GetByID(ctx context.Context, id uint64) (*domain.Product, error) {
	p, err := s.productRepository.FindProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) GetAll(ctx context.Context) ([]*domain.Product, error) {
	p, err := s.productRepository.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Create(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	category, err := s.categoryRepository.FindCategoryByID(ctx, p.CategoryID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, domain.ErrInternal
	}
	p.Category = category
	err = s.productRepository.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Inactivate(ctx context.Context, id uint64) (*domain.Product, error) {
	p, err := s.productRepository.FindProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = p.Inactivate()
	if err != nil {
		return nil, err
	}
	err = s.productRepository.Update(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Activate(ctx context.Context, id uint64) (*domain.Product, error) {
	p, err := s.productRepository.FindProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = p.Activate()
	if err != nil {
		return nil, err
	}
	err = s.productRepository.Update(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
