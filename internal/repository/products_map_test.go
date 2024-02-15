package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProducts(t *testing.T) {
	t.Run("Success search", func(t *testing.T) {
		// arrange
		// - db
		db := map[int]internal.Product{
			1: {Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       100,
					SellerId:    1,
				},
			},
			2: {Id: 2,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 2",
					Price:       200,
					SellerId:    2,
				},
			},
		}
		// - repo
		rp := repository.NewProductsMap(db)

		// act
		products, err := rp.SearchProducts(internal.ProductQuery{Id: 1})

		// assert
		require.NoError(t, err)
		require.Len(t, products, 1)
		require.Equal(t, products[1], db[1])
	})
}
