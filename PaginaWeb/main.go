package main

import (
	"proyecto/controladores"
	"proyecto/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Iniciar Gin router
	router := gin.Default()

	// Conectar a la base de datos
	database.Connect()

	// Configurar la sesión
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Cargar las plantillas HTML
	router.LoadHTMLGlob("templates/*")

	// Registrar rutas
	router.GET("/register", controladores.RegisterPage)
	router.POST("/register", controladores.Register)
	router.GET("/login", controladores.LoginPage)
	router.POST("/login", controladores.Login)
	router.GET("/dashboard", controladores.Dashboard)
	router.GET("/logout", controladores.Logout)

	// Gestión de materias
	router.GET("/materias", controladores.MateriasPage)
	router.POST("/agregar_materia", controladores.AgregarMateria)
	router.GET("/eliminar_materia/:id", controladores.EliminarMateria)
	router.GET("/editar_materia/:id", controladores.EditarMateriaPage)
	router.POST("/actualizar_materia/:id", controladores.ActualizarMateria)

	// Gestión de calificaciones
	router.GET("/calificaciones", controladores.CalificacionesPage)
	router.POST("/agregar_calificacion", controladores.AgregarCalificacion)
	router.GET("/eliminar_calificacion/:id", controladores.EliminarCalificacion)
	router.GET("/editar_calificacion/:id", controladores.EditarCalificacionPage)
	router.POST("/actualizar_calificacion/:id", controladores.ActualizarCalificacion)

	// Ejecutar servidor en el puerto 8080
	err := router.Run(":8080")
	if err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}
