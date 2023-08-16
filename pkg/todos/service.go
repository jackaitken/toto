package todos

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Get() (Todos, error) {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		return Todos{}, err
	}

	list := allLists[dateToday]

	return list, nil
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

	// Key for new item will be (length of list) + 1
	listSize := len(allLists[dateToday].Items)
	allLists[dateToday].Items[strconv.Itoa(listSize+1)] = todo

	if err = Save(allLists); err != nil {
		return err
	}

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

	if err = Save(allLists); err != nil {
		return err
	}

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

	// Check if todo is already marked as complete
	if strings.Contains(todo, CompletedString) {
		fmt.Println("todo is already marked as complete")
		return nil
	}

	todo = Green + todo + CompletedString + ColorReset
	allLists[dateToday].Items[id] = todo

	if err = Save(allLists); err != nil {
		return err
	}

	return nil
}

func Incomplete(id string) error {
	dateToday := time.Now().UTC().Format(TimeFormat)

	allLists, err := GetAllLists()
	if err != nil {
		fmt.Println("error while getting all lists")
		return err
	}

	todo := ColorReset + allLists[dateToday].Items[id] + ColorReset
	if todo == "" {
		fmt.Printf("no todo with id: %v", id)
		return nil
	}

	completeStringIndex := strings.Index(todo, CompletedString)
	if completeStringIndex == -1 {
		fmt.Println("todo is already marked as incomplete")
		return nil
	} else {
		// Make sure we also remove the coloring
		coloringIndex := strings.Index(todo, Green)
		allLists[dateToday].Items[id] = todo[(coloringIndex + len(Green)):completeStringIndex]
	}

	if err = Save(allLists); err != nil {
		return err
	}

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

	lists := ReorderList(allLists)

	if err = Save(lists); err != nil {
		return err
	}

	return nil
}
