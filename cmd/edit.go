package cmd

import (
	"cli-todos/pkg/todos"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit the title of a todo. toto edit <todo_id> <new todo text>",
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		todoId := args[0]
		newTodoText := strings.Join(args[1:], " ")

		if err := todos.Edit(todoId, newTodoText); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
