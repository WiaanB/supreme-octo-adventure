package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/files"
)

func main() {
	// extract the args passed to the program
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("provide me with something, at least...")
		return
	}

	tasks := files.InstantiateJson()

	switch arguments[0] {
	case "add":
		if len(arguments) == 1 {
			fmt.Println("provide me with a task to add...")
			return
		} else {
			tasks.Add(arguments[1])
		}
	case "update":
		if len(arguments) < 3 {
			fmt.Println("provide me with an id and a task to update pls")
			return
		} else {
			num, err := strconv.Atoi(arguments[1])
			if err != nil {
				fmt.Println("provide me with a valid id pls")
				return
			}
			tasks.Update(num, arguments[2])
		}
	case "delete":
		if len(arguments) == 1 {
			fmt.Println("provide me with an id to delete pls")
			return
		} else {
			num, err := strconv.Atoi(arguments[1])
			if err != nil {
				fmt.Println("provide me with a valid id pls")
				return
			}
			tasks.Delete(num)
		}
	case "list":
		listType := "all"
		if len(arguments) == 2 {
			switch arguments[1] {
			case "todo":
				listType = "todo"
			case "in-progress":
				listType = "in-progress"
			case "done":
				listType = "done"
			default:
				fmt.Println("i don't know what you're looking for, here is all of tasks...")
			}
		}
		tasks.List(listType)
	case "mark-in-progress":
		if len(arguments) == 1 {
			fmt.Println("provide me with an id to mark as 'in-progress' pls")
			return
		} else {
			num, err := strconv.Atoi(arguments[1])
			if err != nil {
				fmt.Println("provide me with a valid id pls")
				return
			}
			tasks.MarkAs(num, "in-progress")
		}
	case "mark-done":
		if len(arguments) == 1 {
			fmt.Println("provide me with an id to mark as 'done' pls")
			return
		} else {
			num, err := strconv.Atoi(arguments[1])
			if err != nil {
				fmt.Println("provide me with a valid id pls")
				return
			}
			tasks.MarkAs(num, "done")
		}
	}

	files.SaveFile(tasks)
}
