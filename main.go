package main

import (
	"fmt"
	"os"
)

type Task struct {
	name   string
	status string
	date   string
}

func main() {

	//Welcome message
	fmt.Println("--------------------------------------")
	fmt.Println("\n Welcome to my todo app")
	fmt.Println("\n--------------------------------------")
	fmt.Println()

	//Apps options
	fmt.Println("1. Add a new task")
	fmt.Println("2. Edit a task")
	fmt.Println("3. List all task(s)")
	fmt.Println("4. Remove a task")
	fmt.Println("5. Show task status")
	fmt.Println("\n--------------------------------------")
	fmt.Println("\nQ to quit this application")

	//Take input
	var input string

	fmt.Println("\n--------------------------------------")
	fmt.Print("\nEnter input : ")
	fmt.Scan(&input)

	//Handle input
	switch input {
	case "1":
		add_task()
	case "2":
		edit_task()
	case "3":
		list_tasks()
	case "4":
		remove_task()
	case "5":
		show_task_status()
	case "Q":
		fmt.Println("\nBYE BYE")
		os.Exit(0)
	default:
		fmt.Println("Please choose an option.")
	}
}

//TODO: Implement each functions.

// Adds a task
func add_task() {}

// Edit a task
func edit_task() {}

// List all tasks
func list_tasks() {}

// Remove a task
func remove_task() {}

// Show task status
func show_task_status() {}
