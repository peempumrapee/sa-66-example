package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peempumrapee/sa-66-example/entity"
)

func CreateMember(context *gin.Context) {

	var member entity.Member

	if err := context.ShouldBindJSON(&member); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := entity.DB().Save(&member).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": member})
}

func GetMember(context *gin.Context) {

	var member entity.Member

	id := context.Param("id")

	query := `
    SELECT *
    FROM users
    WHERE id = ?
    `

	if err := entity.DB().Raw(query, id).First(&member).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": member})
}
