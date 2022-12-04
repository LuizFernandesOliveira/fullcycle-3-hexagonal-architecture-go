package db_test

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/fullcycle-3-hexagonal-architecture-go/adapters/db"
	"github.com/LuizFernandesOliveira/fullcycle-3-hexagonal-architecture-go/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE IF NOT EXISTS products (
    		"id" string,
    		"name" string,
    		"price" float,
    		"status" string);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	product := `INSERT INTO products(id, name, price, status) values(?, ?, ?, ?)`

	stmt, err := db.Prepare(product)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec("1", "Product 1", 10.00, "enabled")
	stmt.Exec("2", "Product 2", 20.00, "disabled")
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = application.ENABLED
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
