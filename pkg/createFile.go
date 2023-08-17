package pkg

import (
	"log"
	"os"
)

func Initialize() {
	// Check if the todos.json file already exists. If not create a new one
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	todoFileExists := false

	for _, file := range files {
		if file.Name() == "todos.json" {
			todoFileExists = true
			break
		}
	}

	if !todoFileExists {
		_, _ = os.Create("todos.json")
	}
}
