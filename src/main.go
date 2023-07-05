package main

import (
	"cli-todos/pkg/todos"
	"fmt"
	"log"
)

func main() {
	//err := todos.CreateList()
	//if err != nil {
	//	log.Fatalf("error while creating new list: %v", err)
	//}

	//err := todos.CreateList()
	//if err != nil {
	//	log.Fatalf("error while creating new list: %v", err)
	//}

	err := todos.CreateTodo("finish work")
	if err != nil {
		log.Fatalf("error while creating new todo: %v", err)
	}

	err = todos.Complete("1")
	if err != nil {
		log.Fatalf("error while marking todo as complete: %v", err)
	}

	list, err := todos.Get("2023-07-05")
	if err != nil {
		log.Fatalf("error while getting todo list: %v", err)
	}

	err = todos.Incomplete("1")
	if err != nil {
		log.Fatalf("error while marking todo as incomplete")
	}

	fmt.Println(list)

}
