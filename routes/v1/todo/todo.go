package todo

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sqmatheus/todo-api/controllers/todo"
	"gorm.io/gorm"
)

func ApplyRoutes(group *gin.RouterGroup, db *gorm.DB) {
	ctrl := controller.New(db)

	group.GET("/todos", ctrl.All)
	todo := group.Group("/todo")
	{
		todo.POST("/", ctrl.Create)
		todo.GET("/:id", ctrl.Get)
		todo.DELETE("/:id", ctrl.Del)
		todo.PUT("/:id", ctrl.Update)
	}
}
