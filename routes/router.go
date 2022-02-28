package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sqmatheus/todo-api/routes/v1/todo"
	"gorm.io/gorm"
	"log"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Initialize(engine *gin.Engine, db *gorm.DB) {
	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ping", ping)
			todo.ApplyRoutes(v1, db)
		}
	}

	log.Println("INFO: successfully connected to database")
}
