package todos

import (
	"encoding/json"
	"os"
)

func GetAllLists() (AllLists, error) {
	lists := make(AllLists)

	file, err := os.Open("./todos_test.json")
	if err != nil {
		return AllLists{}, err
	}

	err = json.NewDecoder(file).Decode(&lists)
	if err != nil {
		return AllLists{}, err
	}

	return lists, nil
}
