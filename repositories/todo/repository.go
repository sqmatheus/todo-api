package todo

import (
	"errors"
	"github.com/sqmatheus/todo-api/models"
	"gorm.io/gorm"
)

type InputCreateTodo struct {
	Title string `json:"title" binding:"required,gte=8,lte=30"`
}

type InputUpdateTodo struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"omitempty,gte=8,lte=30"`
	Done  *bool  `json:"done" binding:"omitempty"`
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(id string) (*models.TodoModel, error) {
	var todo models.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", id).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return &todo, errors.New("invalid id")
	}

	return &todo, nil
}

func (r *Repository) All() (*[]models.TodoModel, error) {
	var todos []models.TodoModel
	db := r.db.Model(&todos)

	checkTodo := db.Debug().Select("*").Find(&todos)

	if checkTodo.Error != nil {
		return &todos, errors.New("not found")
	}

	return &todos, nil
}

func (r *Repository) Create(inputTodo *InputCreateTodo) (*models.TodoModel, error) {
	var todo models.TodoModel
	db := r.db.Model(&todo)

	todo.Title = inputTodo.Title

	createTodo := db.Debug().Create(&todo)
	db.Commit()

	if createTodo.Error != nil {
		return &todo, errors.New("failed to create a todo")
	}

	return &todo, nil
}

func (r *Repository) Delete(id string) error {
	var todo models.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", id).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return errors.New("invalid id")
	}

	deleteTodo := db.Debug().Delete(&todo)
	db.Commit()

	if deleteTodo.Error != nil {
		return errors.New("failed to delete a todo")
	}

	return nil
}

func (r *Repository) Update(inputTodo *InputUpdateTodo) (*models.TodoModel, error) {
	var todo models.TodoModel
	db := r.db.Model(&todo)

	checkTodo := db.Debug().Select("*").Where("id = ?", inputTodo.ID).Find(&todo)

	if checkTodo.RowsAffected < 1 {
		return &todo, errors.New("invalid id")
	}

	if inputTodo.Title != "" {
		todo.Title = inputTodo.Title
	}

	if inputTodo.Done != nil {
		todo.Done = *inputTodo.Done
	}

	updateTodo := db.Debug().Select("title", "done").Where("id = ?", inputTodo.ID).Updates(todo)
	db.Commit()

	if updateTodo.Error != nil {
		return &todo, errors.New("failed to create a todo")
	}

	return &todo, nil
}
