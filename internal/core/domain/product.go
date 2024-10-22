package domain

import (
  "time"
  "errors"
)

type Product struct {
  ID int 
  Name string 
  Description string 
  Price float64 
  CreatedAt time.Time 
  UpdatedAt time.Time 
  DeletedAt time.Time 

  category *Category 
}

func (p *Product) IsActive() bool {
  return p.DeletedAt.IsZero()
}

func (p *Product) GetPrice() float64 {
  return p.Price
}

func (p *Product) Inactivate() error {
  if !p.DeletedAt.IsZero() {
    return errors.New("product already inactive")
  }
  p.DeletedAt = time.Now()
  return nil
}

func (p *Product) Activate() error {
  if p.DeletedAt.IsZero() {
    return errors.New("product already active")
  }
  p.DeletedAt = time.Time{}
  return nil
}
