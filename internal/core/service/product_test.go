package service

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/vitovidale/TECH-CHALLENGE/internal/core/domain"
	mock_port "github.com/vitovidale/TECH-CHALLENGE/internal/core/port/mock"
	"go.uber.org/mock/gomock"
)

type createProductTestedInput struct {
	product *domain.Product
}

type createProductTestedOutput struct {
	product *domain.Product
	err     error
}

func TestProductService_CreateProduct(t *testing.T) {
	ctx := context.Background()
	categoryID := gofakeit.Number(1, 1000)
	categoryName := gofakeit.ProductCategory()
	category := &domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: gofakeit.Date(),
		UpdatedAt: gofakeit.Date(),
	}

	productID := gofakeit.Number(1, 1000)
	productName := gofakeit.ProductName()
	productPrice := gofakeit.Price(10, 100)
	productInput := &domain.Product{
		ID:         productID,
		Name:       productName,
		Price:      productPrice,
		CategoryID: categoryID,
		CreatedAt:  gofakeit.Date(),
		UpdatedAt:  gofakeit.Date(),
	}

	productOutput := &domain.Product{
		ID:         productID,
		Name:       productName,
		Price:      productPrice,
		CategoryID: categoryID,
		Category:   category,
		CreatedAt:  gofakeit.Date(),
		UpdatedAt:  gofakeit.Date(),
	}

	testCases := []struct {
		title string
		mocks func(
			productRepository *mock_port.MockProductRepository,
			categoryRepository *mock_port.MockCategoryRepository,
		)
		input  createProductTestedInput
		output createProductTestedOutput
	}{
		{
			title: "Create product successfully",
			mocks: func(
				productRepository *mock_port.MockProductRepository,
				categoryRepository *mock_port.MockCategoryRepository,
			) {
				categoryRepository.EXPECT().FindCategoryByID(gomock.Any(), gomock.Eq(categoryID)).Return(category, nil)
				productRepository.EXPECT().Create(gomock.Any(), gomock.Eq(productInput)).Return(nil)
        productRepository.EXPECT().FindProductByID(gomock.Any(), gomock.Eq(productID)).Return(productOutput, nil)
			},
			input: createProductTestedInput{
				product: productInput,
			},
			output: createProductTestedOutput{
				product: productOutput,
				err:     nil,
			},
		},
		{
		  title: "Category not found",
		  mocks: func(
		    productRepository *mock_port.MockProductRepository,
		    categoryRepository *mock_port.MockCategoryRepository,
		  ) {
		    categoryRepository.EXPECT().FindCategoryByID(ctx, categoryID).Return(nil, domain.ErrDataNotFound)
		  },
		  input: createProductTestedInput{
		    product: productInput,
		  },
		  output: createProductTestedOutput{
		    product: nil,
		    err: domain.ErrCategoryNotFound,
		  },
		},
		{
		  title: "Internal error",
		  mocks: func(
		    productRepository *mock_port.MockProductRepository,
		    categoryRepository *mock_port.MockCategoryRepository,
		  ) {
		    categoryRepository.EXPECT().FindCategoryByID(ctx, categoryID).Return(nil, domain.ErrInternal)
		  },
		  input: createProductTestedInput{
		    product: productInput,
		  },
		  output: createProductTestedOutput{
		    product: nil,
		    err: domain.ErrInternal,
		  },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			productRepository := mock_port.NewMockProductRepository(ctrl)
			categoryRepository := mock_port.NewMockCategoryRepository(ctrl)

			tc.mocks(productRepository, categoryRepository)

			service := NewProductService(categoryRepository, productRepository)
			product, err := service.Create(ctx, tc.input.product)
      if err != nil {
        assert.Equal(t, tc.output.err, err, "Error mismatch")
        return
      }
      product, err = service.GetByID(ctx, productID)
      if err != nil {
        assert.Equal(t, tc.output.err, err, "Error mismatch")
        return
      }
			assert.Equal(t, tc.output.product, product, "Product mismatch")
			assert.Equal(t, tc.output.err, err, "Error mismatch")
		})
	}
}