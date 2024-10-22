package domain

import (
  "time"
  "testing"
  "github.com/stretchr/testify/require"
)

func TestProduct_Activate(t *testing.T) {
  t.Run("product is inactive", func(t *testing.T) {
    p := Product{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      DeletedAt: time.Now(),
    }
    err := p.Activate()
    require.NoError(t, err)
    require.True(t, p.IsActive())
  })
  t.Run("product is active", func(t *testing.T) {
    p := Product{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }
    err := p.Activate()
    require.Error(t, err)
    require.EqualError(t, err, "product already active")
  })
}

func TestProduct_Inactivate(t *testing.T) {
  t.Run("product is active", func(t *testing.T) {
    p := Product{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }
    err := p.Inactivate()
    require.NoError(t, err)
    require.False(t, p.IsActive())
  })
  t.Run("product is inactive", func(t *testing.T) {
    p := Product{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      DeletedAt: time.Now(),
    }
    err := p.Inactivate()
    require.Error(t, err)
    require.EqualError(t, err, "product already inactive")
  })
}
