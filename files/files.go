package files

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/task"
)

func InstantiateJson() task.List {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			createJsonFile()
		}
	} else {
		defer closeFile(file)
	}

	return readJsonFile()
}

func SaveFile(tasks task.List) {
	jsonBlob, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("failed to marshal the tasks:", err)
	}

	err = os.WriteFile("tasks.json", jsonBlob, 0666)
	if err != nil {
		fmt.Println("failed to write to the 'tasks.json' file:", err)
	}
}

func createJsonFile() {
	fmt.Println("making a 'tasks.json' file, because it doesn't exist...")
	err := os.WriteFile("tasks.json", []byte("[]"), 0666)
	if err != nil {
		fmt.Println("failed to create the 'tasks.json' file:", err)
	}
}

func readJsonFile() task.List {
	var tasks task.List

	jsonBlob, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("failed to read the 'tasks.json' file:", err)
	}

	err = json.Unmarshal(jsonBlob, &tasks)
	if err != nil {
		fmt.Println("failed to retrieve the JSON information:", err)
	}

	return tasks
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("error closing file:", err)
	}
}
