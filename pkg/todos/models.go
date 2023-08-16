package todos

type Todos struct {
	Items Item `json:"items"`
}

type AllLists map[string]Todos

type Item map[string]string
