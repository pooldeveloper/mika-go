package models

import "gorm.io/gorm"

type Examen struct {
	gorm.Model

	Nombre        string `json:"nombre" gorm:"not null" `
	ExamenID      uint   `json:"examen_id" gorm:"not null"`
	Instrucciones string `json:"instrucciones" gorm:"not null"`
	Minutos       string `json:"minutos" gorm:"not null"`
}
