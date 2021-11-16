package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todolist/model"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestTodolistRepository_Create(t *testing.T) {
	var todolist []model.Todolist
	todolistRepository := TodolistRepository{Todolist: todolist}

	todolistRepository.Create(model.Todolist{
		Name: "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	})

	assert.Equal(t, 1, len(todolistRepository.Todolist))
}

func TestTodolistRepository_Get(t *testing.T) {
	var todolist []model.Todolist
	todolistRepository := TodolistRepository{Todolist: todolist}

	todolistRepository.Create(model.Todolist{
		Name: "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	})

	assert.Equal(t, 1, len(*todolistRepository.Get()))
}

func TestTodolistRepository_GetByIndex(t *testing.T) {
	var todolist []model.Todolist
	todolistRepository := TodolistRepository{Todolist: todolist}

	todolistRepository.Create(model.Todolist{
		Name: "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	})

	assert.Equal(t, "Learn Golang", todolistRepository.GetByIndex(0).Name)
	assert.Equal(t, "Raden Mohamad Rishwan", todolistRepository.GetByIndex(0).Author)
}

func TestTodolistRepository_Delete(t *testing.T) {
	var todolist []model.Todolist
	todolistRepository := TodolistRepository{Todolist: todolist}

	todolistRepository.Create(model.Todolist{
		Name: "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	})

	todolistRepository.Delete(0)

	assert.Equal(t, 0, len(*todolistRepository.Get()))
}