package models

type Materias struct {
	ID     int `gorm:"primaryKey"`
	Nombre string

	Descripcion string
}
