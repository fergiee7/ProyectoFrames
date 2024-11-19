package models

type Calificaciones struct {
	ID           int `gorm:"primaryKey"`
	MateriaID    int
	AlumnoID     int
	Calificacion int
}
