package todos

type Todos struct {
	Items Item
}

type AllLists map[string]Todos

type Item map[string]string
