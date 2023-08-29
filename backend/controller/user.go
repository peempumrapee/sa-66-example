package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peempumrapee/sa-66-example/entity"
)

func CreateUser(c *gin.Context) {

  var user entity.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

    return 
  }

  if err := entity.DB().Save(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetUser(c *gin.Context) {

  var user entity.User

  id := c.Param("id")

  if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

func ListUsers(c *gin.Context) {
  var users []entity.User

  if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

    return
  }
  
  c.JSON(http.StatusOK, gin.H{"data": users})
}

func DeleteUser(c *gin.Context) {
  id := c.Param("id")

  if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

    return
  }
  
  c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateUser(c *gin.Context) {
  var user entity.User
  
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

    return
  }

  if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})

    return
  }

  if err := entity.DB().Save(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}
