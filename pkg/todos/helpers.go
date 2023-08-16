package todos

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func GetAllLists() (AllLists, error) {
	lists := make(AllLists)

	// If the file is empty we can just return the new lists
	fileInfo, _ := os.Stat("./todos.json")
	if fileInfo.Size() == 0 {
		return lists, nil
	}

	file, err := os.Open("./todos.json")
	if err != nil {
		return AllLists{}, err
	}

	err = json.NewDecoder(file).Decode(&lists)
	if err != nil {
		return AllLists{}, err
	}

	return lists, nil
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

	if err = Save(allLists); err != nil {
		return err
	}

	return nil
}

func Save(allLists AllLists) error {
	lists, err := json.Marshal(allLists)
	if err != nil {
		fmt.Println("error marshaling file")
		return err
	}

	_ = os.WriteFile("./todos.json", lists, 0644)

	return nil
}

func ParseList(todos Todos) string {
	if len(todos.Items) == 0 {
		return Red + "Today's todo list is empty! Use 'toto add' to create some todos" + ColorReset
	}

	formattedTodos := Yellow + "Todo list for today\n" + ColorReset

	// Maps are unordered, so sort the keys first
	keys := make([]string, 0, len(todos.Items))
	for key := range todos.Items {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		formattedTodo := fmt.Sprintf("%s: %s\n", key, todos.Items[key])
		formattedTodos = formattedTodos + formattedTodo
	}

	return formattedTodos
}

func ReorderList(lists AllLists) AllLists {
	// To get a fresh order we create a new map
	// and then populate this new map with the old values

	dateToday := time.Now().UTC().Format(TimeFormat)
	todayList := lists[dateToday]

	newList := Todos{}
	newList.Items = make(Item)
	curIdx := 1

	for _, value := range todayList.Items {
		strIdx := strconv.Itoa(curIdx)
		newList.Items[strIdx] = value
		curIdx += 1
	}

	delete(lists, dateToday)
	lists[dateToday] = newList

	return lists
}
