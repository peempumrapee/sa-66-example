package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peempumrapee/sa-66-example/controller"
	"github.com/peempumrapee/sa-66-example/entity"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}
		c.Next()
	}
}

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/sound/getallsound", controller.GetAllSound)
	r.GET("/sound/getsoundbyid/:id", controller.GetSoundByID)

	r.Run()
}
