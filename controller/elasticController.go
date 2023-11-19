package controller

import (
	"net/http"

	"github.com/billzayy/elastic-golang/services"
	"github.com/gin-gonic/gin"
)

func SearchElasticController(c *gin.Context) {
	inputName := c.Param("name")

	result, err := services.SearchDoc(inputName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": result})
}

func SearchAllController(c *gin.Context) {
	result, err := services.SearchAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failure", "data": err})
	}
	c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": result})
}

func TestController(c *gin.Context) {
	result, _ := services.Test()
	c.JSON(http.StatusAccepted, gin.H{"status": "Successful", "data": result})
}