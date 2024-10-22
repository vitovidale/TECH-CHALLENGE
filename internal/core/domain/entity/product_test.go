package entity

import (
  "time"
  "testing"
  "github.com/stretchr/testify/require"
)

func TestProduct_Activate(t *testing.T) {
  p := Product{
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  err := p.Activate()
  require.Error(t, err)
  require.EqualError(t, err, "product already active")
}
