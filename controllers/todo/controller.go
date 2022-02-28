package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/sqmatheus/todo-api/models"
	repository "github.com/sqmatheus/todo-api/repositories/todo"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct {
	repository *repository.Repository
}

func New(db *gorm.DB) *Controller {
	return &Controller{repository: repository.New(db)}
}

func (controller *Controller) Create(c *gin.Context) {
	var input repository.InputCreateTodo
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var todo *models.TodoModel
	todo, err = controller.repository.Create(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (controller *Controller) Get(c *gin.Context) {
	id := c.Param("id")

	todo, err := controller.repository.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (controller *Controller) All(c *gin.Context) {
	todos, err := controller.repository.All()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (controller *Controller) Del(c *gin.Context) {
	id := c.Param("id")

	err := controller.repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (controller *Controller) Update(c *gin.Context) {
	var input repository.InputUpdateTodo
	input.ID = c.Param("id")
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Title == "" && input.Done == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'title' or 'done' keys not provided",
		})
		return
	}

	var todo *models.TodoModel
	todo, err = controller.repository.Update(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}
