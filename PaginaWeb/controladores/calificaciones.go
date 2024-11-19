package controladores

import (
	"net/http"
	"proyecto/database"
	"proyecto/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Mostrar todas las calificaciones
func CalificacionesPage(c *gin.Context) {
	db := database.GetDB()
	var calificaciones []models.Calificaciones
	db.Find(&calificaciones)

	c.HTML(http.StatusOK, "calificaciones.html", gin.H{
		"calificaciones": calificaciones,
	})
}

// Controlador para agregar una nueva calificación
func AgregarCalificacion(c *gin.Context) {
	materiaIDStr := c.PostForm("materia_id")
	alumnoIDStr := c.PostForm("alumno_id")
	notaStr := c.PostForm("nota")

	materiaID, err := strconv.Atoi(materiaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Materia ID debe ser un número entero"})
		return
	}

	alumnoID, err := strconv.Atoi(alumnoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Alumno ID debe ser un número entero"})
		return
	}

	nota, err := strconv.Atoi(notaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nota debe ser un número entero"})
		return
	}

	calificacion := models.Calificaciones{
		MateriaID:    materiaID,
		AlumnoID:     alumnoID,
		Calificacion: nota,
	}

	if err := database.DB.Create(&calificacion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al agregar calificación"})
		return
	}

	c.Redirect(http.StatusFound, "/calificaciones")
}

// Controlador para eliminar una calificación
func EliminarCalificacion(c *gin.Context) {
	db := database.GetDB()
	calificacionID := c.Param("id")

	if err := db.Delete(&models.Calificaciones{}, calificacionID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al eliminar la calificación"})
		return
	}

	c.Redirect(http.StatusFound, "/calificaciones")
}

func EditarCalificacionPage(c *gin.Context) {
	db := database.GetDB()
	calificacionID := c.Param("id")
	var calificacion models.Calificaciones

	// Buscar la calificación por su ID
	if err := db.First(&calificacion, calificacionID).Error; err != nil {
		c.String(http.StatusNotFound, "Calificación no encontrada")
		return
	}

	// Renderizar la plantilla con la calificación encontrada
	c.HTML(http.StatusOK, "editar_calificaciones.html", gin.H{
		"ID":           calificacion.ID,
		"MateriaID":    calificacion.MateriaID,
		"AlumnoID":     calificacion.AlumnoID,
		"Calificacion": calificacion.Calificacion,
	})
}

func ActualizarCalificacion(c *gin.Context) {
	db := database.GetDB()
	calificacionID := c.Param("id")

	var calificacion models.Calificaciones
	if err := db.First(&calificacion, calificacionID).Error; err != nil {
		c.String(http.StatusNotFound, "Calificación no encontrada")
		return
	}

	// Obtener los datos del formulario
	materiaID, err := strconv.Atoi(c.PostForm("materia_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Materia ID debe ser un número entero"})
		return
	}

	alumnoID, err := strconv.Atoi(c.PostForm("alumno_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Alumno ID debe ser un número entero"})
		return
	}

	nota, err := strconv.Atoi(c.PostForm("nota"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nota debe ser un número entero"})
		return
	}

	// Actualizar la calificación
	calificacion.MateriaID = materiaID
	calificacion.AlumnoID = alumnoID
	calificacion.Calificacion = nota

	if err := db.Save(&calificacion).Error; err != nil {
		c.String(http.StatusInternalServerError, "Error al actualizar calificación")
		return
	}

	// Redirigir a la página de calificaciones
	c.Redirect(http.StatusFound, "/calificaciones")
}
