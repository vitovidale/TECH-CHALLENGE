package domain

import (
  "time"
  "errors"
)

type Category struct {
  ID uint64 
  Name string 
  CreatedAt time.Time 
  UpdatedAt time.Time 
  DeletedAt time.Time 
}

func (c *Category) IsActive() bool {
  return c.DeletedAt.IsZero()
}

func (c *Category) Inactivate() error {
  if !c.DeletedAt.IsZero() {
    return errors.New("category already inactive")
  }
  c.DeletedAt = time.Now()
  return nil
}

func (c *Category) Activate() error {
  if c.DeletedAt.IsZero() {
    return errors.New("category already active")
  }
  c.DeletedAt = time.Time{}
  return nil
}
