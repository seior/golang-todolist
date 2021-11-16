package repository

import "todolist/model"

type TodolistRepository struct {
	Todolist []model.Todolist
}

func (todolistRepository *TodolistRepository) Create(todolist model.Todolist) {
	todolistRepository.Todolist = append(todolistRepository.Todolist, todolist)
}

func (todolistRepository *TodolistRepository) Delete(index int) {
	todolistRepository.Todolist = append(todolistRepository.Todolist[:index], todolistRepository.Todolist[index+1:]...)
}

func (todolistRepository *TodolistRepository) Get() *[]model.Todolist {
	return &todolistRepository.Todolist
}

func (todolistRepository *TodolistRepository) GetByIndex(index int) *model.Todolist {
	return &todolistRepository.Todolist[index]
}
