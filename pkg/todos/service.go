package todos

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02"
const CompletedString = " --DONE!"

func Get(date string) (Todos, error) {
	allLists, err := GetAllLists()
	if err != nil {
		return Todos{}, err
	}

	list := allLists[date]

	/*

	 */

	return list, nil
}

func CreateList() error {
	// Get all current lists
	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error getting all lists")
		return err
	}

	// Create a new list and add an empty map for the todo items
	newList := Todos{}
	newList.Items = make(Item)

	// Add new list with today's date as the key
	dateToday := time.Now().UTC().Format(TimeFormat)
	allLists[dateToday] = newList

	// Marshal and save the file
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos_test.json", lists, 0644)

	return nil
}

func CreateTodo(todo string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	// Check if a list already exists for today
	if len(allLists[dateToday].Items) == 0 {
		// If not, create a new list for today
		err = CreateList()
		if err != nil {
			fmt.Printf("error while creating a new list: %v", err)
			return err
		}

		// Since we created a new list, we need to get them again...
		allLists, err = GetAllLists()
		if err != nil {
			fmt.Println("error getting all lists")
			return err
		}
	}

	// Key for new item will be length of list + 1
	listSize := len(allLists[dateToday].Items)
	allLists[dateToday].Items[strconv.Itoa(listSize+1)] = todo

	// Marshal and save the file
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos_test.json", lists, 0644)

	return nil
}

func Edit(id string, newText string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	todo := allLists[dateToday].Items[id]
	if todo == "" {
		fmt.Printf("no todo with id: %v", id)
		return nil
	}

	allLists[dateToday].Items[id] = newText

	// Marshal and save the file
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos_test.json", lists, 0644)

	return nil
}

func Complete(id string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	todo := allLists[dateToday].Items[id]
	if todo == "" {
		fmt.Printf("no todo with id: %v", id)
		return nil
	}

	todo = todo + CompletedString
	allLists[dateToday].Items[id] = todo

	// Marshal and save the file
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos_test.json", lists, 0644)

	return nil
}

func Incomplete(id string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	todo := allLists[dateToday].Items[id]
	if todo == "" {
		fmt.Printf("no todo with id: %v", id)
		return nil
	}

	index := strings.Index(todo, "--DONE!")
	if index == -1 {
		fmt.Println("todo is already marked as incomplete")
		return nil
	} else {
		allLists[dateToday].Items[id] = todo[:index]
	}

	// Marshal and save the file
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos_test.json", lists, 0644)

	return nil
}

func Delete(id string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	todo := allLists[dateToday].Items[id]
	if todo == "" {
		fmt.Printf("no todo with id: %v", id)
		return nil
	}

	delete(allLists[dateToday].Items, id)

	return nil
}
