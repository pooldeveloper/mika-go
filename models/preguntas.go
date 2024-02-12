package models

import "gorm.io/gorm"

type Pregunta struct {
	gorm.Model

	ExamenID    uint   `json:"examen_id" gorm:"not null"`
	CursoID     uint   `json:"curso_id" gorm:"not null"`
	PreguntaImg string `json:"pregunta_img" gorm:"not null"`
	SolucionImg string `json:"solucion_img" gorm:"not null"`
	Clave       string `json:"clave" gorm:"not null"`
	Numero      string `json:"numero" gorm:"not null"`
}
