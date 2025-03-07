package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"update-products/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// UpdateProductHandler maneja la actualización de productos
func UpdateProductHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}
		product.ID = id

		err = models.UpdateProduct(db, &product)
		if err != nil {
			http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}
