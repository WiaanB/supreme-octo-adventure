package task

import (
	"fmt"
	"time"
)

type Task struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (t *Task) String() string {
	return fmt.Sprintf("id: %d, task: %s, status: %s", t.Id, t.Task, t.Status)
}

type List []Task

func (l *List) Add(task string) {
	if len(task) == 0 {
		fmt.Println("task cannot be empty")
		return
	}

	findLargeId := func() int {
		largestId := 0
		for _, t := range *l {
			if t.Id > largestId {
				largestId = t.Id
			}
		}
		return largestId
	}()

	*l = append(*l, Task{Id: findLargeId + 1, Task: task, Status: "todo", CreatedAt: time.Now().Format(time.Stamp), UpdatedAt: time.Now().Format(time.Stamp)})
	fmt.Println("added task to the list")
}

func (l *List) Update(id int, task string) {
	found := false
	for i, t := range *l {
		if t.Id == id {
			(*l)[i].Task = task
			(*l)[i].UpdatedAt = time.Now().Format(time.Stamp)
			found = true
		}
	}
	if !found {
		fmt.Println("task not found, check the list again??")
	} else {
		fmt.Println("updated task with id:", id)
	}
}

func (l *List) Delete(id int) {
	found := false
	newList := make(List, 0)
	for i, t := range *l {
		if id == t.Id {
			found = true
		} else {
			newList = append(newList, (*l)[i])
		}
	}
	if !found {
		fmt.Println("task not found, check the list again??")
	} else {
		fmt.Println("removed task with id:", id)
		*l = newList
	}
}

func (l *List) MarkAs(id int, status string) {
	found := false
	for i, t := range *l {
		if t.Id == id {
			(*l)[i].Status = status
			found = true
		}
	}
	if !found {
		fmt.Println("task not found, check the list again??")
	} else {
		fmt.Printf("marked task with id: %d as '%s'\n", id, status)
	}
}

func (l *List) List(listType string) {
	switch listType {
	case "todo":
		todos := make([]Task, 0)
		for _, t := range *l {
			if t.Status == "todo" {
				todos = append(todos, t)
			}
		}
		if len(todos) == 0 {
			fmt.Println("no tasks set to 'todo'")
			return
		}
		printList(todos)
	case "in-progress":
		inProgress := make([]Task, 0)
		for _, t := range *l {
			if t.Status == "in-progress" {
				inProgress = append(inProgress, t)
			}
		}
		if len(inProgress) == 0 {
			fmt.Println("no tasks set to 'in-progress'")
			return
		}
		printList(inProgress)
	case "done":
		done := make([]Task, 0)
		for _, t := range *l {
			if t.Status == "done" {
				done = append(done, t)
			}
		}
		if len(done) == 0 {
			fmt.Println("no tasks set to 'done'")
			return
		}
		printList(done)
	default:
		printList(*l)
	}
}

func printList(l List) {
	for _, t := range l {
		fmt.Printf("%s\n", t.String())
	}
}
