package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pooldeveloper/go-gorm-restapi/db"
	"github.com/pooldeveloper/go-gorm-restapi/models"
)

func CrearPregunta(w http.ResponseWriter, r *http.Request) {
	var pregunta models.Pregunta

	json.NewDecoder(r.Body).Decode(&pregunta)

	preguntaCreada := db.DB.Create(&pregunta)

	error := preguntaCreada.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"mensaje": "Pregunta creada exitosamente"}`)
}

func EnviarRespuestas(w http.ResponseWriter, r *http.Request) {
	var cursos []models.Curso

	parametros := mux.Vars(r)

	db.DB.Preload("Preguntas").Where("examen_id = ?", parametros["id"]).Find(&cursos)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&cursos)
}
