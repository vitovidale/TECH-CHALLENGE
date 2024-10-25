package domain

import (
  "errors"
)

var (
  ErrInternal = errors.New("internal error")
  ErrDataNotFound = errors.New("data not found")

  // category errors
  ErrCategoryAlreadyActive = errors.New("category already active")
  ErrCategoryAlreadyInactive = errors.New("category already inactive")
  ErrCategoryNotFound = errors.New("category not found")
  
  // product errors
  ErrProductAlreadyActive = errors.New("product already active")
  ErrProductAlreadyInactive = errors.New("product already inactive")
  ErrProductAlreadyExists = errors.New("product already exists")
  ErrProductNotFound = errors.New("product not found")
)
