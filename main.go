package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	id     int
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
	fmt.Println("6. Create databae")
	fmt.Println("7. Create a table")
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
	case "6":
		create_database()
	case "Q":
		fmt.Println("\nBYE BYE")
		os.Exit(0)
	default:
		fmt.Println("Please choose an option.")
	}
}

//TODO: Implement each functions.

// Adds a task
func add_task() {
	fmt.Println("Adding a task...")
}

// Edit a task
func edit_task() {}

// List all tasks
func list_tasks() {}

// Remove a task
func remove_task() {}

// Show task status
func show_task_status() {}

// Creates a database
func create_database() {
	//Database connection
	database, err := sql.Open("sqlite3", "database/data.db")
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected successfully...")
	}

	//Creating a table
	statment, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (id NOT NULL AUTOINCREMENT, name TEXT, status TEXT, date TEXT)")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := statment.Exec()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Database created successfully...")
		}
	}

}
