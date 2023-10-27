package controllers

import (
	"API/database"
	"API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTemperatura(c *gin.Context) {
	var temperatura models.Temperatura
	if err := c.ShouldBindJSON(&temperatura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	temperatura.Data = time.Now()

	database.DB.Create(&temperatura)
	c.JSON(http.StatusCreated, temperatura)
}

func GetTemperatura(c *gin.Context) {
	var temperaturas []models.Temperatura
	database.DB.Find(&temperaturas)

	var response []map[string]interface{}
	for _, temperatura := range temperaturas {
		dataFomatada := temperatura.Data.Format("2006-01-02")

		temperaturaMap := map[string]interface{}{
			"temperatura": temperatura.Temperatura,
			"data":        dataFomatada,
		}

		response = append(response, temperaturaMap)
	}

	c.JSON(http.StatusOK, response)
}
