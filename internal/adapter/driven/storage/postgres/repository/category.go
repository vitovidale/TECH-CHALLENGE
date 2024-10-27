package repository

import (
  "context"
  "time"
  
  "github.com/Masterminds/squirrel"
  "github.com/jackc/pgx/v5"

  "github.com/vitovidale/TECH-CHALLENGE/internal/adapter/driven/storage/postgres"
  "github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
)

type CategoryRepository struct {
  db *postgres.DB
}


func NewCategoryRepository(db *postgres.DB) *CategoryRepository {
  return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, p *domain.Category) error {
  query := r.db.QueryBuilder.Insert("categories").
  Columns("id", "name", "created_at", "updated_at", "deleted_at").
    Values(p.ID, p.Name, p.CreatedAt, p.UpdatedAt, p.DeletedAt).
    Suffix("RETURNING id")

  sql, args, err := query.ToSql()
  if err != nil {
    return err
  }

  err = r.db.QueryRow(ctx, sql, args...).Scan(&p.ID)

  if err != nil {
    if dbErr := r.db.GetErrorCode(err); dbErr == "23505" {
      return domain.ErrCategoryAlreadyExists
    }
    return err
  }
  
  return nil
}

func (r *CategoryRepository) Update(ctx context.Context, p *domain.Category) error {
  name := postgres.NullString(p.Name)

  query := r.db.QueryBuilder.Update("categories").
    Set("name", squirrel.Expr("COALESCE(?, name)", name)).
    Set("updated_at", time.Now()).
    Where(squirrel.Eq{"id": p.ID}).
    Suffix("RETURNING id")

  sql, args, err := query.ToSql()
  if err != nil {
    return err
  }

  err = r.db.QueryRow(ctx, sql, args...).Scan(&p.ID)
  if err != nil {
    if dbErr := r.db.GetErrorCode(err); dbErr == "23505" {
      return domain.ErrCategoryNotFound
    }
    return err
  }
  
  return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id uint64) error {
  query := r.db.QueryBuilder.Update("categories").
    Set("deleted_at", time.Now()).
    Where(squirrel.Eq{"id": id}).
    Suffix("RETURNING id")

  sql, args, err := query.ToSql()
  if err != nil {
    return err
  }

  _, err = r.db.Exec(ctx, sql, args...)
  if err != nil {
    return err
  }

  return nil
}


// Read operations on category
func (r *CategoryRepository) FindCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
  var c domain.Category
  query := r.db.QueryBuilder.Select("id", "name", "created_at", "updated_at", "deleted_at").
    From("categories").
    Where(squirrel.Eq{"id": id}).
    Limit(1)

  sql, args, err := query.ToSql()
  if err != nil {
    return nil, err
  }

  err = r.db.QueryRow(ctx, sql, args...).Scan(
    &c.ID,
    &c.Name,
    &c.CreatedAt,
    &c.UpdatedAt,
    &c.DeletedAt,
  )
  if err != nil {
    if err == pgx.ErrNoRows {
      return nil, domain.ErrCategoryNotFound
    }
    return nil, err
  }

  return &c, nil
}

func (r *CategoryRepository) FindAllCategories(ctx context.Context) ([]*domain.Category, error) {
  var c domain.Category
  var categories []*domain.Category

  query := r.db.QueryBuilder.Select("id", "name", "created_at", "updated_at", "deleted_at").
    From("categories").
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
      &c.ID,
      &c.Name,
      &c.CreatedAt,
      &c.UpdatedAt,
      &c.DeletedAt,
    )
    if err != nil {
      return nil, err
    }
    categories = append(categories, &c)
  }

  return categories, nil
}
