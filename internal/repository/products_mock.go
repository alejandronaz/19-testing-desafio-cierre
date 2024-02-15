package repository

import "app/internal"

// NewRepositoryProductsMock returns a new RepositoryProductsMock.
func NewRepositoryProductsMock() *RepositoryProductsMock {
	return &RepositoryProductsMock{
		// default values
		FuncSearchProducts: func(query internal.ProductQuery) (p map[int]internal.Product, err error) {
			p = make(map[int]internal.Product)
			return
		},
		Spy: Spy{
			MethodCallsCount: make(map[string]int),
			MethodCallsArgs:  make(map[string][]Args),
		},
	}
}

// Args is a type that represents the arguments of a method call
type Args = []any

// Spy allows us to have more control over the code execution
type Spy struct {
	// MethodCalls
	// - key: method name
	// - value: number of times the method was called
	MethodCallsCount map[string]int

	// MethodArgs
	// - key: method name
	// - value: arguments of each call
	/*
		Example -> SearchProducts was called twice with different arguments:
		{
			"SearchProducts": {
				0: []Args{ internal.ProductQuery{Id: 1} },
				1: []Args{ internal.ProductQuery{Id: 2} },
			},
		}
	*/
	MethodCallsArgs map[string][]Args
}

// RepositoryProductsMock is a mock of the RepositoryProducts interface.
type RepositoryProductsMock struct {
	// FuncSearchProducts is a function that returns a list of products that match the query
	FuncSearchProducts func(query internal.ProductQuery) (p map[int]internal.Product, err error)
	Spy
}

func (r *RepositoryProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {

	// 1. increment the number of times the method was called
	r.MethodCallsCount["SearchProducts"]++

	// 2. save the arguments of the method call
	r.MethodCallsArgs["SearchProducts"] = append(r.MethodCallsArgs["SearchProducts"], Args{query})

	return r.FuncSearchProducts(query)
}
