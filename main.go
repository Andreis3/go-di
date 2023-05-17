package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	//productRepository := product.NewProductRepository(db)
	//
	//productUseCase := product.NewProductUseCase(productRepository)
	//
	//product, err := productUseCase.GetProduct(1)

	usecase := NewUseCase(db)
	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
