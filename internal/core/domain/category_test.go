package domain

import (
  "time"

  "testing"
  "github.com/stretchr/testify/require"
)

func TestCategory_Activate(t *testing.T) {
  t.Run("category is inactive", func(t *testing.T) {
    c := Category{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      DeletedAt: time.Now(),
    }
    err := c.Activate()
    require.NoError(t, err)
    require.True(t, c.IsActive())
  })
  t.Run("category is active", func(t *testing.T) {
    c := Category{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }
    err := c.Activate()
    require.Error(t, err)
    require.EqualError(t, err, "category already active")
  })
}

func TestCategory_Inactivate(t *testing.T) {
  t.Run("category is active", func(t *testing.T) {
    c := Category{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
    }
    err := c.Inactivate()
    require.NoError(t, err)
    require.False(t, c.IsActive())
  })
  t.Run("category is inactive", func(t *testing.T) {
    c := Category{
      CreatedAt: time.Now(),
      UpdatedAt: time.Now(),
      DeletedAt: time.Now(),
    }
    err := c.Inactivate()
    require.Error(t, err)
    require.EqualError(t, err, "category already inactive")
  })
}
