package port

import (
  "context"

  "github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

// ProductRepositoryReader is an interface that wraps all the reading operations for a product. 
type ProductRepositoryReader interface {
  FindProductByID(ctx context.Context, id int) (*domain.Product, error)
  FindAllProducts(ctx context.Context) ([]*domain.Product, error)
}

// ProductRepositoryWriter is an interface that wraps all the writing operations for a product.
type ProductRepositoryWriter interface {
  Create(ctx context.Context, p *domain.Product) error
  Update(ctx context.Context, p *domain.Product) error
  Delete(ctx context.Context, id int) error
}

// ProductRepository is an interface that wraps all the reading and writing operations for a product.
type ProductRepository interface {
  ProductRepositoryReader
  ProductRepositoryWriter
}

// ProductService is an interface that wraps all the operations for a product.
type ProductService interface {
  GetByID(ctx context.Context, id int) (*domain.Product, error)
  GetAll(ctx context.Context) ([]*domain.Product, error)
  Create(ctx context.Context, p *domain.Product) error
  Inactivate(ctx context.Context, id int) error
  Activate(ctx context.Context, id int) error
}
