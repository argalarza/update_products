package main

import (
	"log"
	"net/http"
	"os"
	"update-products/controllers"
	"update-products/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Conectar con la base de datos SQL Server
	db, err := models.InitDB()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	defer db.Close()

	// Inicializar el enrutador
	r := mux.NewRouter()

	// Rutas para manejar productos
	r.HandleFunc("/products/{id}", controllers.UpdateProductHandler(db)).Methods("PUT")

	// Aplicar CORS: Configurar CORS explícitamente para permitir solicitudes de todos los orígenes
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                 // Permitir solicitudes de cualquier origen
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Métodos permitidos
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"}, // Encabezados permitidos
		ExposedHeaders:   []string{"Content-Type"},                                      // Encabezados expuestos
		AllowCredentials: false,                                                         // No permitir credenciales (cookies, headers de autenticación)
	}).Handler(r)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "3008" // Valor por defecto
	}
	log.Println("🚀 Servidor iniciado en el puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
