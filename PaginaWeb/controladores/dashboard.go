package controladores

import (
	"net/http"
	"proyecto/database"
	"proyecto/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var materias []models.Materias
	database.GetDB().Find(&materias)

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"user":     user,
		"materias": materias,
	})
}
