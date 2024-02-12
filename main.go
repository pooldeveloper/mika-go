package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pooldeveloper/go-gorm-restapi/db"
	"github.com/pooldeveloper/go-gorm-restapi/models"
	"github.com/pooldeveloper/go-gorm-restapi/routes"
	"github.com/rs/cors"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Examen{})
	db.DB.AutoMigrate(models.Curso{})
	db.DB.AutoMigrate(models.Pregunta{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/api/crear-examen", routes.CrearExamen).Methods("POST")
	r.HandleFunc("/api/examenes", routes.ObtenerExamenes).Methods("GET")
	r.HandleFunc("/api/examen/{id}", routes.ObtenerExamen).Methods("GET")

	r.HandleFunc("/api/crear-curso", routes.CrearCurso).Methods("POST")
	r.HandleFunc("/api/cursos/{id}", routes.ObtenerCursos).Methods("GET")

	r.HandleFunc("/api/crear-pregunta", routes.CrearPregunta).Methods("POST")
	r.HandleFunc("/api/enviar-respuestas/{id}", routes.EnviarRespuestas).Methods("POST")

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := corsOptions.Handler(r)

	http.ListenAndServe(":4000", handler)
}
