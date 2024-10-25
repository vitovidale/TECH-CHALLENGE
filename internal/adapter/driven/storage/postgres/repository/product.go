package repository

import (
  "context"
  
  "github.com/Masterminds/squirrel"
  "github.com/jackc/pgx/v5"

  "github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driven/storage/postgres"
  "github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

type ProductRepository struct {
  db *postgres.DB
}


func NewProductRepository(db *postgres.DB) *ProductRepository {
  return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *domain.Product) error {
  query := r.db.QueryBuilder.Insert("products").
    Columns("id", "name", "price", "category_id", "created_at", "updated_at", "deleted_at").
    Values(p.ID, p.Name, p.Price, p.CategoryID, p.CreatedAt, p.UpdatedAt, p.DeletedAt).
    Suffix("RETURNING id")

  sql, args, err := query.ToSql()
  if err != nil {
    return err
  }

  err = r.db.QueryRow(ctx, sql, args...).Scan(
    &p.ID,
    &p.Name,
    &p.Price,
    &p.CategoryID,
    &p.CreatedAt,
    &p.UpdatedAt,
    &p.DeletedAt,
  )

  if err != nil {
    if dbErr := r.db.GetErrorCode(err); dbErr == "23505" {
      return domain.ErrProductAlreadyExists
    }
    return err
  }
  
  return nil
}

func (r *ProductRepository) FindProductByID(ctx context.Context, id int) (*domain.Product, error) {
  var p domain.Product
  query := r.db.QueryBuilder.Select("id", "name", "price", "category_id", "created_at", "updated_at", "deleted_at").
    From("products").
    Where(squirrel.Eq{"id": id}).
    Limit(1)

  sql, args, err := query.ToSql()
  if err != nil {
    return nil, err
  }

  err = r.db.QueryRow(ctx, sql, args...).Scan(
    &p.ID,
    &p.Name,
    &p.Price,
    &p.CategoryID,
    &p.CreatedAt,
    &p.UpdatedAt,
    &p.DeletedAt,
  )
  if err != nil {
    if err == pgx.ErrNoRows {
      return nil, domain.ErrProductNotFound
    }
    return nil, err
  }

  return &p, nil
}

func (r *ProductRepository) FindAllProducts(ctx context.Context) ([]*domain.Product, error) {
  var p domain.Product
  var products []*domain.Product

  query := r.db.QueryBuilder.Select("id", "name", "price", "category_id", "created_at", "updated_at", "deleted_at").
    From("products").
    Where(squirrel.Eq{"deleted_at": nil}).
    OrderBy("id")

  sql, args, err := query.ToSql()
  if err != nil {
    return nil, err
  }

  rows, err := r.db.Query(ctx, sql, args...)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    err := rows.Scan(
      &p.ID,
      &p.Name,
      &p.Price,
      &p.CategoryID,
      &p.CreatedAt,
      &p.UpdatedAt,
      &p.DeletedAt,
    )
    if err != nil {
      return nil, err
    }
    products = append(products, &p)
  }

  return products, nil
}
