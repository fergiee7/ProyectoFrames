package database

import (
	"fmt"
	"proyecto/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:test@tcp(127.0.0.1:3306)/escuela?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return
	}

	fmt.Println("Conexión exitosa a la base de datos")

	// Migración de tablas sin claves foráneas
	if err := DB.AutoMigrate(&models.Materias{}, &models.Calificaciones{}, &models.Usuarios{}); err != nil {
		fmt.Println("Error al migrar las tablas:", err)
		return
	}
	fmt.Println("Tablas migradas correctamente")

}

func GetDB() *gorm.DB {
	return DB
}
