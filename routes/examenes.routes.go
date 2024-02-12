package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pooldeveloper/go-gorm-restapi/db"
	"github.com/pooldeveloper/go-gorm-restapi/models"
)

func CrearExamen(w http.ResponseWriter, r *http.Request) {
	var examen models.Examen

	json.NewDecoder(r.Body).Decode(&examen)

	examenCreado := db.DB.Create(&examen)

	error := examenCreado.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"mensaje": "Examen creado exitosamente"}`)
}

func ObtenerExamenes(w http.ResponseWriter, r *http.Request) {
	var examenes []models.Examen

	db.DB.Find(&examenes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&examenes)
}

func ObtenerExamen(w http.ResponseWriter, r *http.Request) {
	var examen models.Examen

	parametros := mux.Vars(r)

	db.DB.First(&examen, "examen_id = ?", parametros["id"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&examen)
}
