package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model

	ExamenID  uint       `json:"examen_id" gorm:"not null"`
	Nombre    string     `json:"nombre" gorm:"not null"`
	Preguntas []Pregunta `json:"preguntas"`
}
