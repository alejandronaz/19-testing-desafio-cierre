package repository

import "app/internal"

// NewRepositoryProductsMock returns a new RepositoryProductsMock.
func NewRepositoryProductsMock() *RepositoryProductsMock {
	return &RepositoryProductsMock{}
}

// RepositoryProductsMock is a mock of the RepositoryProducts interface.
type RepositoryProductsMock struct {
	// FuncSearchProducts is a function that returns a list of products that match the query
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	// Calls is a struct that contains the number of calls to each method.
	Calls struct {
		SearchProducts int
	}
}

func (r *RepositoryProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	r.Calls.SearchProducts++
	return r.FuncSearchProducts(query)
}
