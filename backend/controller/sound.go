package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peempumrapee/sa-66-example/entity"
)

func CreateSound(context *gin.Context) {

	var sound entity.Sound

	if err := context.ShouldBindJSON(&sound); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := entity.DB().Save(&sound).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": sound})
}

func GetAllSound(context *gin.Context) {

	var sounds []entity.Sound

	query := `
      SELECT *
      FROM sounds
    `

	if err := entity.DB().Raw(query).Scan(&sounds).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": sounds})
}

func GetSoundByID(context *gin.Context) {

	var sound entity.Sound

	id := context.Param("id")

	query := `
    SELECT *
    FROM sounds
    WHERE id = ?
    `

	if err := entity.DB().Raw(query, id).First(&sound).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": sound})

}
