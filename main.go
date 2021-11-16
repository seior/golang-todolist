package main

import (
	"fmt"
	"strconv"
	"todolist/model"
	"todolist/repository"
	"todolist/utils"
)

func main() {
	var todolists []model.Todolist
	todolistRepository := repository.TodolistRepository{Todolist: todolists}
	fmt.Print("Program Sederhana Todolist" +
		"\n==========================")

	for {
		fmt.Println("\nOption :" +
			"\n1. Lihat Todolist" +
			"\n2. Buat Todolist" +
			"\n3. Hapus Todolist" +
			"\n4. Keluar")

		fmt.Print("Option : ")
		utils.Input.Scan()

		input := utils.Input.Text()

		if input == "1" {
			fmt.Println("\nTodolist : ")
			for index, todolist := range *todolistRepository.Get() {
				fmt.Println(strconv.Itoa(index) + ". " + todolist.Name)
			}
		} else if input == "2" {
			var todolist model.Todolist

			fmt.Println("\nBuat Todolist :  ")

			fmt.Print("Todolist : ")
			utils.Input.Scan()
			todolist.Name = utils.Input.Text()

			fmt.Print("Author : ")
			utils.Input.Scan()
			todolist.Author = utils.Input.Text()

			todolistRepository.Create(todolist)
		} else if input == "3" {
			fmt.Print("\nHapus : ")
			utils.Input.Scan()

			input = utils.Input.Text()

			index, _ := strconv.Atoi(input)

			todolistRepository.Delete(index)
		} else if input == "4" {
			break
		} else {
			fmt.Println("Mohon masukkan angka 1 - 4")
		}
	}

	fmt.Println("Terima Kasih :3")
}
