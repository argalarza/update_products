package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Product representa la estructura de un producto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
}

// UpdateProduct actualiza un producto en la base de datos
func UpdateProduct(db *sqlx.DB, product *Product) error {
	query := `
        UPDATE Products 
        SET name = @name, description = @description, price = @price, quantity = @quantity, category = @category, brand = @brand
        WHERE id = @id
    `

	_, err := db.Exec(query,
		sql.Named("name", product.Name),
		sql.Named("description", product.Description),
		sql.Named("price", product.Price),
		sql.Named("quantity", product.Quantity),
		sql.Named("category", product.Category),
		sql.Named("brand", product.Brand),
		sql.Named("id", product.ID),
	)

	return err
}
