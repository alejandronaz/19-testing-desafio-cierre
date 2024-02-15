package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("Success get", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewRepositoryProductsMock()
		rp.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = make(map[int]internal.Product)
			p[1] = internal.Product{
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "description",
					Price:       10.5,
					SellerId:    1,
				},
			}
			return
		}
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdGet := hd.Get()
		// - request
		req := httptest.NewRequest("GET", "/product?id=1", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		hdGet(res, req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, res.Header())
		expectedBody := `{
			"message": "success",
			"data": {
				"1": {
					"id": 1,
					"description": "description",
					"price": 10.5,
					"seller_id": 1
				}
			}
		}`
		require.JSONEq(t, expectedBody, res.Body.String())
		// verify SearchProducts was called
		require.Equal(t, 1, rp.Calls.SearchProducts)
	})
	t.Run("Error get - invalid id", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewRepositoryProductsMock()
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdGet := hd.Get()
		// - request
		req := httptest.NewRequest("GET", "/product?id=1a", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		hdGet(res, req)

		// assert
		require.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("Error get - internal error", func(t *testing.T) {
		// arrange
		// - repository
		rp := repository.NewRepositoryProductsMock()
		rp.FuncSearchProducts = func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			err = errors.New("internal error")
			return
		}
		// - handler
		hd := handler.NewProductsDefault(rp)
		hdGet := hd.Get()
		// - request
		req := httptest.NewRequest("GET", "/product?id=1", nil)
		// - response
		res := httptest.NewRecorder()

		// act
		hdGet(res, req)

		// assert
		require.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
