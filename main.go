package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task

func main() {

	loadTasks()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	if command == "add" {

		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name")
			return
		}

		title := os.Args[2]

		newTask := Task{
			ID:        len(tasks) + 1,
			Title:     title,
			Completed: false,
		}

		tasks = append(tasks, newTask)
		saveTasks()

		fmt.Println("Task added successfully")

	} else if command == "list" {

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		for _, t := range tasks {
			status := "Incomplete"
			if t.Completed {
				status = "Complete"
			}
			fmt.Println(t.ID, "-", t.Title, "-", status)
		}

	} else if command == "delete" {

		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}

		idStr := os.Args[2]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		var updatedTasks []Task

		for _, t := range tasks {
			if t.ID != id {
				updatedTasks = append(updatedTasks, t)
			}
		}

		tasks = updatedTasks
		saveTasks()

		fmt.Println("Task deleted successfully")

	} else if command == "complete" {

		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID")
			return
		}

		idStr := os.Args[2]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Completed = true
			}
		}

		saveTasks()
		fmt.Println("Task marked as complete")
	}
}

func saveTasks() {
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func loadTasks() {

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}
}
