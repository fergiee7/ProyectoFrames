package models

type Usuarios struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Rol      string
}
