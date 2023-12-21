package controller

import (
	"net/http"

	"github.com/GITTIIII/sa-66-example/entity"
	"github.com/gin-gonic/gin"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender

	db, err := entity.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Raw("SELECT * FROM genders").Scan(&genders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genders})
}
