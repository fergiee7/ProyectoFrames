package controladores

import (
	"net/http"
	"proyecto/database"
	"proyecto/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Todos los campos son obligatorios"})
		return
	}

	var user models.Usuarios
	database.GetDB().Where("username = ?", username).First(&user)
	if user.ID != 0 {
		c.HTML(http.StatusConflict, "register.html", gin.H{"error": "Usuario ya existe"})
		return
	}

	newUser := models.Usuarios{Username: username, Password: password, Rol: "administrativo"}
	database.GetDB().Create(&newUser)

	c.Redirect(http.StatusFound, "/login")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var user models.Usuarios

	database.GetDB().Where("username = ? AND password = ?", username, password).First(&user)
	if user.ID == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Credenciales incorrectas"})
		return
	}

	session := sessions.Default(c)
	session.Set("user", user.Username)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Mostrar página de registro

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
