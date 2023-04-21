package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// welcome page func
// getting inputs from console
// create todo
// delete todo by id
// update todo
// list all todos created

var todoList []string

func main() {

	for {
		welcomePage()
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}

		switch input {
		case 1:
			title, description := getUserInputToCreateTodo()
			createTodo(title, description)
		case 2:
			id := getIdFromUser()
			err := deleteItemById(id)
			if err != nil {
				return
			}
		case 3:
			listAllItemsInTodo()
		case 4:
			id := getIdFromUser()
			updateATodo(id)

		default:
			fmt.Println("Please enter the correct input")
		}

	}

}
func welcomePage() {
	fmt.Println("Welcome to our Todo App")
	fmt.Println("------------------------")
	fmt.Println("Type 1 to create a new Todo")
	fmt.Println("Type 2 to delete a  Todo")
	fmt.Println("Type 3 to list all Todos")
	fmt.Println("Type 4 to update a Todo")

}

func createTodo(title string, description string) []string {
	todo := title + " " + description
	todoList = append(todoList, todo)
	fmt.Println(todoList)

	return todoList
}

func getUserInputToCreateTodo() (string, string) {
	var title string
	var description string

	fmt.Println("Please enter title of your task")
	title, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", ""
	}

	fmt.Println("Please enter the description of your task")
	description, er := bufio.NewReader(os.Stdin).ReadString('\n')
	if er != nil {
		return "", ""
	}

	return title, description
}

func listAllItemsInTodo() {
	fmt.Println("List of all todos: ")
	for i := 0; i < len(todoList); i++ {
		value := fmt.Sprintf("%v | %v", i, todoList[i])
		fmt.Println(value)
	}
}

func getIdFromUser() int {
	var id int

	fmt.Println("Please enter id to deleted: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

func updateATodo(id int) {
	title, desc := getUserInputToCreateTodo()
	todo := title + ": " + desc

	for i, _ := range todoList {
		if i == id {
			todoList[i] = todo
		}

	}
	fmt.Println(todoList)
}

func deleteItemById(id int) error {

	if id < 0 || id > len(todoList) {
		fmt.Println("please enter a valid todo id to delete")
		err := errors.New("please enter a valid todo id to delete")
		if err != nil {
			return err
		}
	}
	for i := 0; i < len(todoList); i++ {
		if i == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
		}

	}
	fmt.Println(todoList)
	return nil
}
