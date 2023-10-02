package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peempumrapee/sa-66-example/entity"
)

func CreateCreator(context *gin.Context) {

	var creator entity.Creator

	if err := context.ShouldBindJSON(&creator); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := entity.DB().Save(&creator).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": creator})
}

func GetCreator(context *gin.Context) {

	var creator entity.Creator

	id := context.Param("id")

	query := `
    SELECT *
    FROM creators
    WHERE id = ?
    `

	if err := entity.DB().Raw(query, id).First(&creator).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	context.JSON(http.StatusOK, gin.H{"data": creator})
}
