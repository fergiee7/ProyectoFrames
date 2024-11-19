package controladores

import (
	"net/http"
	"proyecto/database"
	"proyecto/models"

	"github.com/gin-gonic/gin"
)

// Controlador para la página de gestión de materias
func MateriasPage(c *gin.Context) {
	db := database.GetDB()
	var materias []models.Materias

	// Obtener todas las materias de la base de datos
	if result := db.Find(&materias); result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": result.Error.Error()})
		return
	}

	// Renderizar la página de materias con los datos obtenidos
	c.HTML(http.StatusOK, "materias.html", gin.H{
		"materias": materias,
	})
}

// Controlador para agregar una nueva materia
func AgregarMateria(c *gin.Context) {
	db := database.GetDB()
	nombre := c.PostForm("nombre")
	descripcion := c.PostForm("descripcion")

	materia := models.Materias{Nombre: nombre, Descripcion: descripcion}
	db.Create(&materia)

	// Redirigir de vuelta a la página de materias
	c.Redirect(http.StatusFound, "/materias")
}

// Controlador para eliminar una materia
func EliminarMateria(c *gin.Context) {
	db := database.GetDB()
	materiaID := c.Param("id")

	// Eliminar la materia por ID
	if err := db.Delete(&models.Materias{}, materiaID).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error al eliminar la materia")
		return
	}

	// Redirigir de nuevo a la página de gestión de materias
	c.Redirect(http.StatusFound, "/materias")
}

// Controlador para cargar el formulario de edición de una materia
func EditarMateriaPage(c *gin.Context) {
	db := database.GetDB()
	materiaID := c.Param("id")
	var materia models.Materias

	// Obtener la materia por su ID
	if err := db.First(&materia, materiaID).Error; err != nil {
		c.String(http.StatusNotFound, "Materia no encontrada")
		return
	}

	// Renderizar el formulario de edición con los datos actuales
	c.HTML(http.StatusOK, "editar_materia.html", gin.H{
		"materia": materia,
	})
}

// Controlador para actualizar la materia en la base de datos
func ActualizarMateria(c *gin.Context) {
	db := database.GetDB()
	materiaID := c.Param("id")

	var materia models.Materias
	if err := db.First(&materia, materiaID).Error; err != nil {
		c.String(http.StatusNotFound, "Materia no encontrada")
		return
	}

	// Obtener los datos del formulario
	nombre := c.PostForm("nombre")
	descripcion := c.PostForm("descripcion")

	// Actualizar los campos de la materia
	materia.Nombre = nombre
	materia.Descripcion = descripcion

	db.Save(&materia)

	// Redirigir de nuevo a la página de gestión de materias
	c.Redirect(http.StatusFound, "/materias")
}
